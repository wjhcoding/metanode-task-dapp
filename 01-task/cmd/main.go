package main

import (
	"net/http"
	"time"

	"github.com/wjhcoding/metanode-task-dapp/01-task/config"
	"github.com/wjhcoding/metanode-task-dapp/01-task/internal/router"
	"github.com/wjhcoding/metanode-task-dapp/01-task/internal/service/blockchain"
	"github.com/wjhcoding/metanode-task-dapp/01-task/pkg/global/log"
)

func main() {
	log.InitLogger(config.GetConfig().AppName, config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	log.Logger.Info("config", log.Any("config", config.GetConfig()))

	log.Logger.Info("start server", log.String("start", "start web sever..."))

	// ✅ 初始化区块链客户端
	blockchain.InitClient()

	newRouter := router.NewRouter()

	s := &http.Server{
		Addr:           ":8888",
		Handler:        newRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if nil != err {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
}
