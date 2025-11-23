package blockchain

import (
	"log"

	"github.com/wjhcoding/metanode-task-dapp/01-task/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

// InitClient 初始化以太坊 RPC 客户端
func InitClient() {
	cfg := config.GetConfig()

	client, err := ethclient.Dial(cfg.Blockchain.RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum RPC: %v", err)
	}

	Client = client
	log.Println("Connected to Ethereum RPC:", cfg.Blockchain.RPC_URL)
}
