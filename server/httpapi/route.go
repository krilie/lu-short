package httpapi

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type HttpApi struct {
	HttpOpenApi *HttpOpenApi
}

func NewHttpApi(httpOpenApi *HttpOpenApi) *HttpApi {
	return &HttpApi{HttpOpenApi: httpOpenApi}
}

func (api *HttpApi) SetRouterAndStartHttpService(ctx context.Context, port int) func(waitDuration time.Duration) error {
	rootRouter := gin.Default()

	rootRouter.GET("/", api.HttpOpenApi.Redirect)

	// 开始服务
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: rootRouter,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			} else {
				panic(err.Error())
			}
			return
		}
	}()

	return func(waitDuration time.Duration) error {
		ctxTimeout, cancelFunc := context.WithTimeout(ctx, waitDuration)
		defer cancelFunc()
		// shutdown
		err := srv.Shutdown(ctxTimeout)
		if err != nil {
			return err
		} else {
			return nil
		}
	}

}
