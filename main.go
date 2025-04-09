package main

import (
	"703room/703room.com/appcontext"
	"703room/703room.com/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// Init services
	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatalf("[ERROR]-error happened %s", err)
	}
	logger := config.SetupLogger()
	httpClient := config.SetupHTTPClient()

	// App context
	appCtx := appcontext.NewAppContext(db, nil, logger, httpClient)
	fmt.Println(appCtx)
	r.Run()
}
