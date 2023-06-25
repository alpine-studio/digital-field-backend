package server

import (
	"hackaton/store"
	"hackaton/util"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const (
	timeSleep = time.Second * 30
)

func Start(cfg *Config, logger *zap.SugaredLogger) error {
	st := store.New()
	util.GetFromFile(st)

	go func() {
		for {
			time.Sleep(timeSleep)
			util.SaveToFile(st)
		}
	}()

	serv := newServer(st, logger)

	return http.ListenAndServe(net.JoinHostPort(cfg.Host, cfg.Port), serv)
}
