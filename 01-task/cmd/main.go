package main

import (
	"net/http"
	"os"
	"time"

	"github.com/wjhcoding/metanode-task-dapp/01-task/internal/config"
	"github.com/wjhcoding/metanode-task-dapp/01-task/internal/router"
	"github.com/wjhcoding/metanode-task-dapp/01-task/internal/service/blockchain"
	"github.com/wjhcoding/metanode-task-dapp/01-task/pkg/global/log"
)

func main() {

	// 1. 加载配置
	// 获取环境变量：dev / prod / test
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// 统一入口：Load(env)
	_, err := config.Load(env)
	if err != nil {
		panic(err)
	}

	// 2. 初始化日志
	log.InitLogger(config.GetConfig().AppName, config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	log.Logger.Info("config loaded", log.Any("config", config.GetConfig()))

	log.Logger.Info("start server", log.String("start", "start web sever..."))

	// 3. 初始化区块链
	blockchain.InitClient()

	// 4. 初始化路由
	newRouter := router.NewRouter()

	// 5. 启动 HTTP 服务
	server := &http.Server{
		Addr:           ":8888",
		Handler:        newRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if nil != err {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
}
