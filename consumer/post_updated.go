package consumer

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/newschain-cache/app"
	"github.com/s4kibs4mi/newschain-cache/contracts"
	"github.com/s4kibs4mi/newschain-cache/data"
	"github.com/s4kibs4mi/newschain-cache/log"
	"os"
)

func BootConsumerPostUpdated(c *contracts.Contracts, startBlock *uint64) error {
	watchOption := &bind.WatchOpts{
		Start: startBlock,
	}
	postUpdatedChan := make(chan *contracts.ContractsPostUpdated)
	postUpdatedSubscription, err := c.WatchPostUpdated(watchOption, postUpdatedChan)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case err := <-postUpdatedSubscription.Err():
				log.Log().Errorln(err)
				os.Exit(-1)
			case event := <-postUpdatedChan:
				log.Log().Infoln(event.Id)
				log.Log().Infoln(event.Title)
				log.Log().Infoln(event.Raw.Index)
				log.Log().Infoln(event.Raw.BlockNumber)
				log.Log().Infoln(event.Raw.TxIndex)

				post, err := data.GetPostByID(app.DB().Post, uuid.MustParse(event.Id))
				if err != nil {
					log.Log().Errorln(err)
					continue
				}

				_, err = post.
					Update().
					SetTitle(event.Title).
					Save(context.Background())
				if err != nil {
					log.Log().Errorln(err)
				}
			}
		}
	}()

	return nil
}
