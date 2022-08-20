package main

import (
	log "log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/controller"
	utils "mrkresnofatihdev/apps/gotwittercloneapp/utils"
	http "net/http"
)

func main() {
	server := utils.ApplicationServer{}
	server.AddController(&controller.PlayerController{})
	server.Initialize()

	application.ResolveApplicationsOnClose()

	log.Print("Server running!")
	log.Fatal(http.ListenAndServe(":8000", server.MainRouter))

}
