package httpAPI

import (
	"fmt"
	"net/http"
)

// @Summary     404 Handler
// @Description API未定义或未绑定
// @Tags        Default
// @Accept      plain
// @Produce     json
// @Success     404  {object}  HttpRes{data=nil}
func F_404NotFound(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(res, "{\"code\":404,\"msg\":\"API Not Found\"}")
}

// @Summary     401 Handler
// @Description 未登陆或SSO验证失败
// @Tags        Default
// @Accept      plain
// @Produce     json
// @Success     401  {object}  HttpRes{data=nil}
func F_401Unauthorized(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(res, "{\"code\":401,\"msg\":\"Unauthorized, sso Login please\"}")
}

// @Summary     500 Handler
// @Description 各种过程错误，如json解析失败、DB解析失败等
// @Tags        Default
// @Accept      plain
// @Produce     json
// @Success     500  {object}  HttpRes{data=nil}
func F_500ServerError(res http.ResponseWriter, _ *http.Request, msg string) {
	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(res, "{\"code\":500,\"msg\":\"InternalServerError: -->[%s]<--\"}", msg)
}
