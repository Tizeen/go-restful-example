package user

import (
	. "github.com/Tizeen/go-restful-example/handler"
	"github.com/Tizeen/go-restful-example/pkg/errno"
	"github.com/Tizeen/go-restful-example/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r ListRequest

	// 无法绑定GET请求中的数据
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})

}
