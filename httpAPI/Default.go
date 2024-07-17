package httpAPI

import (
	"net/http"
)

// @Summary     Default Handler
// @Description 所有未定义的请求，默认返回404,即API未定义或未绑定
// @Tags        Default
// @Accept      plain
// @Produce     json
// @Success     404  {object}  HttpRes{data=nil}
// @Router      /api [get]
func H_Default(res http.ResponseWriter, req *http.Request) {
	F_404NotFound(res, req)
}
