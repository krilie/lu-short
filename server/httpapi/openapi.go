package httpapi

import (
	"context"
	"lu-short/common/utils/id_util"
	"lu-short/module/service"
)
import "github.com/gin-gonic/gin"

type HttpOpenApi struct {
	OpenApiSvc *service.UnionSvcOpenApi
}

func NewHttpOpenApi(openApiSvc *service.UnionSvcOpenApi) *HttpOpenApi {
	return &HttpOpenApi{OpenApiSvc: openApiSvc}
}

// https://t.we.com/eRexvjLkE :key
func (openApi *HttpOpenApi) Redirect(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.String(404, "no page find here")
		return
	}
	redirect, err := openApi.OpenApiSvc.LuShortService.Dao.GetReDirectByKey(context.WithValue(context.Background(), "ginCtx", c), key)
	if err != nil {
		c.String(404, "some err here")
		return
	}
	cookie, err := c.Cookie("id")
	if err != nil {
		cookie = id_util.GetUuid()
		c.SetCookie("id", cookie, 6000000, "/", "", true, true)
	}
	println(cookie)
	c.Redirect(302, redirect.OriUrl)
	return
}
