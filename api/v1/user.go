package v1

import (
	"fmt"
	"net/http"
	"pika/internal/service"

	"github.com/gin-gonic/gin"
)

//GetUserInfoByIDRequest - request
type GetUserInfoByIDRequest struct {
	Request string `json:"request_id"`
	UserID  uint64 `json:"user_id"`
}

//GetUserInfoByID - Get user info by user_id
func GetUserInfoByID(c *gin.Context) {
	var req GetUserInfoByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
			"data": nil,
		})
		return
	}
	user, err := service.NewUserService().GetUserInfoByID(req.UserID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": user,
	})
}
