package consumer

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/newschain-cache/app"
	"github.com/s4kibs4mi/newschain-cache/contracts"
	"github.com/s4kibs4mi/newschain-cache/data"
	"github.com/s4kibs4mi/newschain-cache/ent"
	"github.com/s4kibs4mi/newschain-cache/log"
	"os"
	"time"
)

func BootConsumerPostCreated(c *contracts.Contracts, startBlock *uint64) error {
	watchOption := &bind.WatchOpts{
		Start: startBlock,
	}
	postCreatedChan := make(chan *contracts.ContractsPostCreated)
	postCreatedSubscription, err := c.WatchPostCreated(watchOption, postCreatedChan)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case err := <-postCreatedSubscription.Err():
				log.Log().Errorln(err)
				os.Exit(-1)
			case event := <-postCreatedChan:
				log.Log().Infoln(event.Id)
				log.Log().Infoln(event.Title)
				log.Log().Infoln(event.Raw.Index)
				log.Log().Infoln(event.Raw.BlockNumber)
				log.Log().Infoln(event.Raw.TxIndex)

				tx, err := app.DB().Tx(context.Background())
				if err != nil {
					log.Log().Errorln(err)
					os.Exit(-1)
				}

				_, err = tx.
					Post.
					Create().
					SetID(uuid.MustParse(event.Id)).
					SetTitle(event.Title).
					SetCreatedAt(time.Unix(0, event.CreatedAt.Int64())).
					Save(context.Background())
				if err != nil {
					tx.Rollback()
					log.Log().Errorln(err)
					os.Exit(-1)
				}

				latestBlock, err := data.GetLatestBlock(tx.LatestBlock)
				if err != nil {
					if !ent.IsNotFound(err) {
						tx.Rollback()
						log.Log().Errorln(err)
						os.Exit(-1)
					}
				}

				if latestBlock == nil {
					latestBlock, err = tx.LatestBlock.
						Create().
						SetBlockNumber(uint32(event.Raw.BlockNumber)).
						Save(context.Background())
					if err != nil {
						tx.Rollback()
						log.Log().Errorln(err)
						os.Exit(-1)
					}
				} else {
					latestBlock, err = latestBlock.
						Update().
						SetBlockNumber(uint32(event.Raw.BlockNumber)).
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
			}
		}
	}()

	return nil
}
