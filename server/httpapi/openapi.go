package httpapi

import (
	"lu-short/common/utils/id_util"
	"lu-short/module/service"
	"lu-short/server/httpapi/ginutil"
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
	// customerId customerId
	customerId, err := c.Cookie("cId")
	if err != nil {
		customerId = id_util.GetUuid()
		c.SetCookie("id", customerId, 6000000, "/", "", true, true)
	}
	// customerId customerId
	deviceId, err := c.Cookie("dId")
	if err != nil {
		deviceId = id_util.GetUuid()
		c.SetCookie("id", deviceId, 6000000, "/", "", true, true)
	}
	// remote ip devices
	ip := c.ClientIP()
	// agent
	agent := c.Request.UserAgent()
	// 跳转
	ginWrap := ginutil.NewGinWrap(c, openApi.OpenApiSvc.LuShortService.Log)
	oriUrl, err := openApi.OpenApiSvc.LuShortService.Redirect(ginWrap.GetAppContext(), key, customerId, ip, agent, deviceId)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	c.Redirect(302, oriUrl)
	return
}
