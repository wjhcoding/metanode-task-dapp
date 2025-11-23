package blockchain

import (
	"context"
	"math/big"
)

// BlockInfo 用于 API 层返回
type BlockInfo struct {
	Number       uint64   `json:"number"`
	Hash         string   `json:"hash"`
	Time         uint64   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

// QueryBlock 查询指定区块
func QueryBlock(blockNumber int64) (*BlockInfo, error) {
	if Client == nil {
		return nil, ErrClientNotInitialized
	}

	bn := big.NewInt(blockNumber)
	block, err := Client.BlockByNumber(context.Background(), bn)
	if err != nil {
		return nil, err
	}

	txHashes := make([]string, 0, len(block.Transactions()))
	for _, tx := range block.Transactions() {
		txHashes = append(txHashes, tx.Hash().Hex())
	}

	return &BlockInfo{
		Number:       block.NumberU64(),
		Hash:         block.Hash().Hex(),
		Time:         block.Time(),
		Transactions: txHashes,
	}, nil
}

// ErrClientNotInitialized 当 Client 没初始化时返回
var ErrClientNotInitialized = &CustomError{"Ethereum client not initialized"}

// CustomError 自定义错误类型
type CustomError struct {
	Msg string
}

func (e *CustomError) Error() string {
	return e.Msg
}
