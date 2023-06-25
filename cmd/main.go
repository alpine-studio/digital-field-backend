package main

import (
	"hackaton/internal/server"
	"hackaton/store"
	"hackaton/util"
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := server.NewConfig("config.yml")
	if err != nil {
		logger.Sugar().Error(err)
	}

	log.Fatalln(server.Start(cfg, logger.Sugar()))

	//check()
}

func check() {
	st := store.New()
	util.SaveToFile(st)
}
