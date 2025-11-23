package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/wjhcoding/metanode-task-dapp/01-task/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// SendTransaction 发送 ETH 转账交易，返回 txHash
func SendTransaction(to string, amountEth float64) (string, error) {
	cfg := config.GetConfig()

	privateKey, err := crypto.HexToECDSA(cfg.Blockchain.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(int64(amountEth * 1e18)) // 转 ETH

	gasLimit := cfg.Blockchain.GasLimit
	gasTipCap := big.NewInt(cfg.Blockchain.GasTipCapGwei * 1e9)
	gasFeeCap := big.NewInt(cfg.Blockchain.GasFeeCapGwei * 1e9)

	toAddress := common.HexToAddress(to)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(11155111), // Sepolia
		Nonce:     nonce,
		To:        &toAddress,
		Value:     value,
		Gas:       gasLimit,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      []byte{},
	})

	signer := types.NewLondonSigner(big.NewInt(11155111))
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return "", err
	}

	err = Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	log.Println("TX Hash:", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}
