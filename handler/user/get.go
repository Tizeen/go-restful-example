package user

import (
	. "github.com/Tizeen/go-restful-example/handler"
	"github.com/Tizeen/go-restful-example/model"
	"github.com/Tizeen/go-restful-example/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")

	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
