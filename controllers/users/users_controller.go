package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MathisDetourbet/bookstore_users-api/domain/users"
	"github.com/MathisDetourbet/bookstore_users-api/services"
	"github.com/MathisDetourbet/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

const (
	requestParamUserID     = "user_id"
	requestParamUserStatus = "status"
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
		return
	}

	// Add `Location` header and return the URI of the new resource just created
	uri := fmt.Sprintf("/users/%d", result.ID)
	c.Writer.Header().Add("Location", uri)

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {
	userID, idErr := getUserID(c.Param(requestParamUserID))
	if idErr != nil {
		c.JSON(idErr.Status, idErr.Message)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	userID, idErr := getUserID(c.Param(requestParamUserID))
	if idErr != nil {
		c.JSON(idErr.Status, idErr.Message)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, &user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	if isPartial {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
	}
}

func DeleteUser(c *gin.Context) {
	userID, idErr := getUserID(c.Param(requestParamUserID))
	if idErr != nil {
		c.JSON(idErr.Status, idErr.Message)
		return
	}

	if deleteErr := services.DeleteUser(userID); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func Search(c *gin.Context) {
	status := c.Query(requestParamUserStatus)

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
