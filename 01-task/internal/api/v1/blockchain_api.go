package v1

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/wjhcoding/metanode-task-dapp/01-task/internal/service/blockchain"
	"github.com/wjhcoding/metanode-task-dapp/01-task/pkg/common/response"
)

// QueryBlock 查询指定区块
func QueryBlock(c *gin.Context) {
	numStr := c.Query("num") // ?num=6000000
	blockNum, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailMsg("invalid block number"))
		return
	}

	blockInfo, err := blockchain.QueryBlock(blockNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Ok("block query success", blockInfo))

}

// SendTransaction 发送交易
func SendTransaction(c *gin.Context) {
	var req struct {
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.FailMsg(err.Error()))
		return
	}

	txHash, err := blockchain.SendTransaction(req.To, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.OkData(map[string]string{
		"txHash": txHash,
	}))
}

// CounterGet godoc
// @Summary 查询 Counter 当前值
// @Param addr query string true "合约地址"
func CounterGet(c *gin.Context) {
	addr := c.Query("addr")
	if !common.IsHexAddress(addr) {
		c.JSON(http.StatusBadRequest, response.FailMsg("invalid address"))
		return
	}
	val, err := blockchain.CounterGet(common.HexToAddress(addr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.OkData(gin.H{"count": val}))
}

// CounterInc godoc
// @Summary 调用 increment()
func CounterInc(c *gin.Context) {
	doCounterMutate(c, "inc")
}

// CounterDec godoc
// @Summary 调用 decrement()
func CounterDec(c *gin.Context) {
	doCounterMutate(c, "dec")
}

// 统一处理 inc/dec
func doCounterMutate(c *gin.Context, op string) {
	type Req struct {
		Addr string `json:"addr" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.FailMsg(err.Error()))
		return
	}
	if !common.IsHexAddress(req.Addr) {
		c.JSON(http.StatusBadRequest, response.FailMsg("invalid address"))
		return
	}
	var hash string
	var err error
	switch op {
	case "inc":
		hash, err = blockchain.CounterInc(common.HexToAddress(req.Addr))
	case "dec":
		hash, err = blockchain.CounterDec(common.HexToAddress(req.Addr))
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.OkData(gin.H{"txHash": hash}))
}
