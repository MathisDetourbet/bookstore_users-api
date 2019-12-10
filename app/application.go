package app

import (
	_ "github.com/MathisDetourbet/bookstore_users-api/config"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run()
}
