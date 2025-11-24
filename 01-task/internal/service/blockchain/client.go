package blockchain

import (
	"log"
	"math/big"

	"github.com/wjhcoding/metanode-task-dapp/01-task/internal/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client
var ChainID *big.Int

// InitClient 初始化以太坊 RPC 客户端
func InitClient() {
	cfg := config.GetConfig()

	client, err := ethclient.Dial(cfg.Blockchain.RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum RPC: %v", err)
	}

	Client = client

	ChainID = big.NewInt(cfg.Blockchain.ChainID)

	log.Println("Connected to Ethereum RPC:", cfg.Blockchain.RPC_URL)
}
