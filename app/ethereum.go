package app

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/s4kibs4mi/newschain-cache/config"
)

var eClient *ethclient.Client

func ConnectToEthereum() error {
	c, err := ethclient.Dial(config.Ethereum().Node)
	if err != nil {
		return err
	}

	eClient = c
	return nil
}

func Ethereum() *ethclient.Client {
	return eClient
}
