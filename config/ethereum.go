package config

import "github.com/spf13/viper"

type EthereumCfg struct {
	Node            string
	ContractAddress string
}

var ethereumCfg EthereumCfg

func Ethereum() EthereumCfg {
	return ethereumCfg
}

func LoadEthereum() {
	mu.Lock()
	defer mu.Unlock()

	ethereumCfg = EthereumCfg{
		Node:            viper.GetString("ethereum.node"),
		ContractAddress: viper.GetString("ethereum.contract_address"),
	}
}
