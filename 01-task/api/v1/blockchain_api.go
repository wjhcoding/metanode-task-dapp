package v1

import (
	"net/http"
	"strconv"

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
