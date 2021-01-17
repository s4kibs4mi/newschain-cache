package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/newschain-cache/app"
	"github.com/s4kibs4mi/newschain-cache/config"
	"github.com/s4kibs4mi/newschain-cache/consumer"
	"github.com/s4kibs4mi/newschain-cache/contracts"
	"github.com/s4kibs4mi/newschain-cache/data"
	"github.com/s4kibs4mi/newschain-cache/ent"
	"github.com/s4kibs4mi/newschain-cache/log"
	"github.com/spf13/cobra"
	"os"
	"sync"
	"time"
)

var (
	workerCmd = &cobra.Command{
		Use:   "worker",
		Short: "worker starts the worker",
		Run:   startWorker,
	}
)

func startWorker(cmd *cobra.Command, args []string) {
	address := common.HexToAddress(config.Ethereum().ContractAddress)
	newsChainContract, err := contracts.NewContracts(address, app.Ethereum())
	if err != nil {
		log.Log().Errorln(err)
		os.Exit(-1)
	}

	tx, err := app.DB().Tx(context.Background())
	if err != nil {
		log.Log().Errorln(err)
		os.Exit(-1)
	}

	var latestBlockNumber uint64
	latestBlock, err := data.GetLatestBlock(tx.LatestBlock)
	if err != nil {
		if !ent.IsNotFound(err) {
			tx.Rollback()
			log.Log().Errorln(err)
			os.Exit(-1)
		}
	} else {
		latestBlockNumber = uint64(latestBlock.BlockNumber)
	}

	eventsIterator, err := newsChainContract.ContractsFilterer.FilterPostCreated(&bind.FilterOpts{
		Start: latestBlockNumber + 1,
	})
	if err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		os.Exit(-1)
	}

	for eventsIterator.Next() {
		event := eventsIterator.Event
		log.Log().Infoln(event.Id)
		log.Log().Infoln(event.Title)
		log.Log().Infoln(event.Raw.Index)
		log.Log().Infoln(event.Raw.BlockNumber)
		log.Log().Infoln(event.Raw.TxIndex)

		_, err := app.DB().
			Post.
			Create().
			SetID(uuid.MustParse(event.Id)).
			SetTitle(event.Title).
			SetCreatedAt(time.Unix(0, event.CreatedAt.Int64())).
			Save(context.Background())
		if err != nil {
			log.Log().Errorln(err)
		}

		latestBlockNumber = event.Raw.BlockNumber
	}

	if latestBlock == nil {
		latestBlock, err = tx.LatestBlock.
			Create().
			SetBlockNumber(uint32(latestBlockNumber)).
			Save(context.Background())
		if err != nil {
			tx.Rollback()
			log.Log().Errorln(err)
			os.Exit(-1)
		}
	} else {
		latestBlock, err = latestBlock.
			Update().
			SetBlockNumber(uint32(latestBlockNumber)).
			Save(context.Background())
		if err != nil {
			tx.Rollback()
			log.Log().Errorln(err)
			os.Exit(-1)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Log().Errorln(err)
		os.Exit(-1)
	}

	latestBlockNumber = uint64(latestBlock.BlockNumber)
	if err := consumer.BootConsumerPostCreated(newsChainContract, &latestBlockNumber); err != nil {
		log.Log().Errorln(err)
		os.Exit(-1)
	}
	if err := consumer.BootConsumerPostUpdated(newsChainContract, &latestBlockNumber); err != nil {
		log.Log().Errorln(err)
		os.Exit(-1)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	log.Log().Infoln("Consumer is waiting for events...")
	wg.Wait()
}
