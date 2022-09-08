package main

import (
	"mrkresnofatihdev/apps/gotwittercloneapp/controller"
	"mrkresnofatihdev/apps/gotwittercloneapp/events"
	"sync"
)

func main() {
	var appRunState sync.WaitGroup
	appRunState.Add(1)
	controller.InitHttpServer(&appRunState)
	events.InitRabbitMq(&appRunState)
	appRunState.Wait()
}
