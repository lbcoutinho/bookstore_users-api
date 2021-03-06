package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lbcoutinho/bookstore_users-api/domain/users"
	"github.com/lbcoutinho/bookstore_users-api/services"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var user users.User
	// Unmarshal request body into user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UserService.Create(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(isPublicRequest(c)))
}

func Get(c *gin.Context) {
	userId, idErr := getUserIdFromPath(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UserService.Get(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(isPublicRequest(c)))
}

func Update(c *gin.Context) {
	userId, idErr := getUserIdFromPath(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	// Unmarshal request body into user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserService.Update(user, isPartial)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(isPublicRequest(c)))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserIdFromPath(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.Delete(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(isPublicRequest(c)))
}

func getUserIdFromPath(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("User id should be a number")
	}
	return userId, nil
}

func isPublicRequest(c *gin.Context) bool {
	return c.GetHeader("X-Public") == "true"
}
