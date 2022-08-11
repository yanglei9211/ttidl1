package util

import "github.com/cloudwego/hertz/pkg/app"

// 返回包装
func WriteResponse(c *app.RequestContext, data interface{}) {
	//mpData, _ := StructToMap(data, "json", "")
	retData := map[string]interface{}{}
	retData["Status"] = 1
	retData["Data"] = data
	c.JSON(200, retData)
}

func WriteException() {

}
