package main

import (
	"api_admin_server/internal/routers"
	"api_admin_server/store"
	"fmt"
	"time"
)

func init() {
	log := new(store.Log)
	log.Get().Debug("server started at %v", time.Now())
}

func main() {
	config := new(store.Config)
	err := routers.Init().Run(fmt.Sprintf(":%d", config.Get().Http.Port))
	if err != nil {
		panic(err)
	}
}
