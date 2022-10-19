package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// SuccessResponse 直接返回成功
func SuccessResponse(c *gin.Context, data interface{}) {
	ReturnJson(c, http.StatusOK, 200, "", data)
}

// FailResponse Fail 失败的业务逻辑
func FailResponse(c *gin.Context, dataCode int, err error) {
	ReturnJson(c, http.StatusBadRequest, dataCode, err.Error(), "")
	c.Abort()
}

// ErrorParam 参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, 502, "参数检验失败", wrongParam)
	c.Abort()
}
