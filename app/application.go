package app

import (
	// load .env file
	_ "github.com/MathisDetourbet/bookstore_users-api/config"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication by mapping urls and run the router
func StartApplication() {
	mapUrls()
	router.Run()
}
