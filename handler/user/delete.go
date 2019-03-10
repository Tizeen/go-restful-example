package user

import (
	. "github.com/Tizeen/go-restful-example/handler"
	"github.com/Tizeen/go-restful-example/model"
	"github.com/Tizeen/go-restful-example/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
