package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/wjhcoding/metanode-task-dapp/01-task/config"
	"github.com/wjhcoding/metanode-task-dapp/01-task/internal/service/blockchain/contract"
)

/*
✅ 下一步可选（按需要再迭代）
用 bind.WaitMined 在接口里等收据，返回最新 count（已在 counter.go 预留）。
把合约地址写进配置文件，避免前端每次传。
加 Swagger 注解（项目已用 swag init 即可生成）。
拆分 service 层，把「等待收据」做成异步任务，通过轮询或 websocket 推送给前端。
支持 EIP-1559 动态费率、本地内存队列防重放、日志落库。
完成 1~4 项，你就拥有了一个「能查、能写、带登录鉴权」的 Counter 微服务雏形，剩下的只是锦上添花。
*/

// CounterGet 查询合约当前 count（view 方法）
func CounterGet(contractAddr common.Address) (uint64, error) {
	c, err := contract.NewContract(contractAddr, Client)
	if err != nil {
		return 0, fmt.Errorf("new contract caller: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	raw, err := c.GetCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, fmt.Errorf("call GetCount: %w", err)
	}
	return raw.Uint64(), nil
}

// CounterInc 发送交易调用 increment()，返回 txHash
// 私钥、链 ID、Gas 价格等都从 cfg 里拿，保持和 tx.go 一致。
func CounterInc(contractAddr common.Address) (string, error) {
	return counterSend(
		"increment",
		contractAddr,
		func(c *contract.Contract, opts *bind.TransactOpts) (*types.Transaction, error) {
			return c.Increment(opts)
		})
}

// CounterDec 发送交易调用 decrement()
func CounterDec(contractAddr common.Address) (string, error) {
	return counterSend(
		"decrement",
		contractAddr,
		func(c *contract.Contract, opts *bind.TransactOpts) (*types.Transaction, error) {
			return c.Decrement(opts)
		})
}

func counterSend(method string, contractAddr common.Address,
	call func(*contract.Contract, *bind.TransactOpts) (*types.Transaction, error),
) (string, error) {

	cfg := config.GetConfig()
	priv, err := crypto.HexToECDSA(cfg.Blockchain.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("bad private key: %w", err)
	}

	// 构造 transactOpts（复用你在 tx.go 里的参数）
	opts, err := bind.NewKeyedTransactorWithChainID(priv, big.NewInt(cfg.Blockchain.ChainID))
	if err != nil {
		return "", fmt.Errorf("new transactor: %w", err)
	}
	nonce, err := Client.PendingNonceAt(context.Background(), opts.From)
	if err != nil {
		return "", fmt.Errorf("nonce: %w", err)
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasTipCap = big.NewInt(cfg.Blockchain.GasTipCapGwei * 1e9)
	opts.GasFeeCap = big.NewInt(cfg.Blockchain.GasFeeCapGwei * 1e9)
	opts.GasLimit = 0 // 让客户端估算

	c, err := contract.NewContract(contractAddr, Client)
	if err != nil {
		return "", fmt.Errorf("new contract transactor: %w", err)
	}

	tx, err := call(c, opts)
	if err != nil {
		return "", fmt.Errorf("%s tx: %w", method, err)
	}

	log.Printf("[%s] tx sent: %s", method, tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}
