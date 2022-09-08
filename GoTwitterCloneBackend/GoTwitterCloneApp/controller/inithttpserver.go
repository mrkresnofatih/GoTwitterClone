package controller

import (
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
	"sync"
)

var httpServer *http.Server

func InitHttpServer(runState *sync.WaitGroup) {
	go func() {
		muxServer := utils.ApplicationServer{}
		muxServer.AddController(&FollowController{})
		muxServer.AddController(&PlayerController{})
		muxServer.AddController(&TweetController{})
		muxServer.Initialize()

		httpServer = &http.Server{
			Addr: ":8000",
			Handler: muxServer.MainRouter,
		}

		log.Println("HttpServer started")

		log.Println(httpServer.ListenAndServe())
		runState.Done()
	}()
}
