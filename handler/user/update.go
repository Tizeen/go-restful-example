package user

import (
	. "github.com/Tizeen/go-restful-example/handler"
	"github.com/Tizeen/go-restful-example/model"
	"github.com/Tizeen/go-restful-example/pkg/errno"
	"github.com/Tizeen/go-restful-example/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// 获取url中的id并转换成int
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	// 检查body的数据是否存在
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = uint64(userId)

	// 验证数据结构是否合法
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
