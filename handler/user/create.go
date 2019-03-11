package user

import (
	. "github.com/Tizeen/go-restful-example/handler"
	"github.com/Tizeen/go-restful-example/model"
	"github.com/Tizeen/go-restful-example/pkg/errno"
	"github.com/Tizeen/go-restful-example/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {

	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest

	// 如果请求没有数据，返回错误
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 验证请求的数据结构是否正确
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密密码数据
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 添加用户
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 添加到返回数据结构中的 data 字段
	rsp := CreateResponse{
		Username: r.Username,
	}

	SendResponse(c, nil, rsp)

}
