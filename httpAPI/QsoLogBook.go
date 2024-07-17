package httpAPI

import (
	"fmt"
	"net/http"

	"github.com/QsoLogger/QsoLogger-API/sso"
)

// @Summary     Add QSO LogBook
// @Description 添加QSO日志本
// @Tags        QsoLogBook
// @Accept      json
// @Produce     json
// @Param       SSOID      query   string      false    "sso auth id"
// @Param       QsoLogBook body    QsoLogBook  true     "Qso logBook详细信息"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/qsoLogBook/add [post]
func H_QsoLogBookAdd(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := sso.GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		F_401Unauthorized(res, req)
		return
	}

	if req.Method != "POST" {
		F_404NotFound(res, req)
		return
	}

	fmt.Fprintf(res, "{\"Type\":\"%d\", \"ssoName\":\"%s\"}", ssoInfo.Type, ssoInfo.UserName)
}

// @Summary     list QSO LogBook
// @Description 添加QSOBook日志本
// @Tags        QsoLogBook
// @Accept      json
// @Produce     json
// @Param       SSOID      query   string  false    "sso auth id"
// @Success     200  {object}  HttpRes{data=[]QsoLogBook}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/qsoLogBook/list [post]
func H_QsoLogBookList(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := sso.GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(res, "{\"code\":401,\"msg\":\"forbiden\"}")
		return
	}

	if req.Method != "POST" {
		H_Default(res, req)
		return
	}

	fmt.Fprintf(res, "{\"Type\":\"%d\", \"ssoName\":\"%s\"}", ssoInfo.Type, ssoInfo.UserName)
}

// @Summary     update QSO LogBook
// @Description 更新QSO日志本
// @Tags        QsoLogBook
// @Accept      json
// @Produce     json
// @Param       SSOID      query   string      false    "sso auth id"
// @Param       QsoLogBook body    QsoLogBook  true     "Qso logBook详细信息"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/qsoLogBook/update [post]
func H_QsoLogBookUpdate(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := sso.GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(res, "{\"code\":401,\"msg\":\"forbiden\"}")
		return
	}

	if req.Method != "POST" {
		H_Default(res, req)
		return
	}

	fmt.Fprintf(res, "{\"Type\":\"%d\", \"ssoName\":\"%s\"}", ssoInfo.Type, ssoInfo.UserName)
}

// @Summary     hide QSO LogBook
// @Description 隐藏QSO日志本（用户不能物理删除，只能隐藏）
// @Tags        QsoLogBook
// @Accept      json
// @Produce     json
// @Param       SSOID      query   string  false    "sso auth id"
// @Param       logBookId  body    []int   true     "Qso logBookId数组"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/qsoLogBook/hide [post]
func H_QsoLogBookHide(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := sso.GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(res, "{\"code\":401,\"msg\":\"forbiden\"}")
		return
	}

	if req.Method != "POST" {
		H_Default(res, req)
		return
	}

	fmt.Fprintf(res, "{\"Type\":\"%d\", \"ssoName\":\"%s\"}", ssoInfo.Type, ssoInfo.UserName)
}

// @Summary     Admin update QSO LogBook
// @Description 管理员更新其他用户的QSO日志本
// @Tags        QsoLogBook
// @Accept      json
// @Produce     json
// @Param       SSOID      query   string  false    "sso auth id"
// @Param       QsoLogBook body    QsoLogBook  true     "Qso logBook详细信息"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/admin/qsoLogBook/update [post]
func H_QsoLogBookAdminUpdate(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := sso.GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(res, "{\"code\":401,\"msg\":\"forbiden\"}")
		return
	}

	if req.Method != "POST" {
		H_Default(res, req)
		return
	}

	fmt.Fprintf(res, "{\"Type\":\"%d\", \"ssoName\":\"%s\"}", ssoInfo.Type, ssoInfo.UserName)
}

// @Summary     admin unHide QSO LogBook
// @Description 管理员恢复显示QSO日志本
// @Tags        QsoLogBook
// @Accept      json
// @Produce     json
// @Param       SSOID      query   string  false    "sso auth id"
// @Param       logIdBook  body    []int   true     "Qso logIdBook数组"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/admin/qsoLogBook/unHide [post]
func H_QsoLogBookAdminUnhide(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := sso.GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(res, "{\"code\":401,\"msg\":\"forbiden\"}")
		return
	}

	if req.Method != "POST" {
		H_Default(res, req)
		return
	}

	fmt.Fprintf(res, "{\"Type\":\"%d\", \"ssoName\":\"%s\"}", ssoInfo.Type, ssoInfo.UserName)
}

// @Summary     admin delete QSO LogBook
// @Description 管理员物理删除QSO日志本（不可恢复）,删除前，需要校验该日志本下不可以有日志
// @Tags        QsoLogBook
// @Accept      json
// @Produce     json
// @Param       SSOID      query   string  false    "sso auth id"
// @Param       logBookId  body    []int   true     "Qso logBookId数组"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/admin/qsoLogBook/delete [post]
func H_QsoLogBookAdminDelete(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := sso.GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(res, "{\"code\":401,\"msg\":\"forbiden\"}")
		return
	}

	if req.Method != "POST" {
		H_Default(res, req)
		return
	}

	fmt.Fprintf(res, "{\"Type\":\"%d\", \"ssoName\":\"%s\"}", ssoInfo.Type, ssoInfo.UserName)
}
