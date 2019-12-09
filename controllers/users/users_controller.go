package users

import (
	"net/http"

	"github.com/MathisDetourbet/bookstore_users-api/app/utils/errors"
	"github.com/MathisDetourbet/bookstore_users-api/domain/users"
	"github.com/MathisDetourbet/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: handle user create error
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
