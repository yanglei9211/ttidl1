package util

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"runtime"
	"time"
)

// 接口耗时
func MidShowMethodCostTime() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		t := time.Now()
		c.Next(ctx)
		latency := time.Since(t)
		hlog.Info("info mark")
		//fmt.Println(c.ClientIP(), c.Response.StatusCode(), latency)
		//fmt.Println(string(c.URI().Path()), string(c.URI().Host()), c.URI().QueryArgs(), string(c.Request.Method()))
		method := string(c.Request.Method())

		hlog.Info(fmt.Sprintf("%d\t|%s\t|%s\t|%s\t|%s",
			c.Response.StatusCode(),
			latency.String(),
			c.ClientIP(),
			method,
			string(c.URI().Path())))

		// logging query or body
		hlog.Info(fmt.Sprintf("query args: %s", c.URI().QueryArgs()))
		hlog.Info(fmt.Sprintf("body args: %s", string(c.Request.Body())))
		hlog.Info("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		if method == "GET" {
			hlog.Info(fmt.Sprintf("query args: %s", c.URI().QueryArgs()))
		} else if method == "POST" {
			hlog.Info(fmt.Sprintf("body args: %s", string(c.Request.Body())))
		} else {
			// todo more method
		}
	}
}

// 错误处理
func ErrHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)
		fmt.Println("\n\n---------------------------\n\n")
		fmt.Println("in err handler")
		if length := len(c.Errors); length > 0 {
			e := c.Errors[length-1]
			err := e.Err
			if err != nil {
				fmt.Println("=================")
				fmt.Println("异常处理")
				fmt.Println(err.Error())
				ret := map[string]string{}
				ret["emsg"] = err.Error()
				ret["stack"] = ""
				WriteException(c, ret)
				return
			}
		}
	}
}

func LocalRecover() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1<<13)
				runtime.Stack(buf, false)
				strbuf := string(buf)
				perror := err.(error)
				ret := map[string]string{}
				ret["emsg"] = perror.Error()
				ret["stack"] = strbuf
				WriteException(c, ret)
			}
		}()
		c.Next(ctx)
	}
}
