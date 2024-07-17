package httpAPI

import (
	"fmt"
	"net/http"

	"github.com/QsoLogger/QsoLogger-API/sso"
)

// @Summary     Add QSO Log
// @Description 添加QSO日志
// @Tags        QsoLog
// @Accept      json
// @Produce     json
// @Param       SSOID  query   string  false    "sso auth id"
// @Param       QsoLog body    QsoLog  true     "Qso log详细信息"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/qsoLog/add [post]
func H_QsoLogAdd(res http.ResponseWriter, req *http.Request) {
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

// @Summary     list QSO Log
// @Description 添加QSO日志
// @Tags        QsoLog
// @Accept      json
// @Produce     json
// @Param       SSOID     query   string  false    "sso auth id"
// @Param       logBookId query   int     true     "Qso logBook的Id"
// @Success     200  {object}  HttpRes{data=[]QsoLog}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/qsoLog/listByBookId [post]
func H_QsoLogListByBookId(res http.ResponseWriter, req *http.Request) {
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

// @Summary     update QSO Log
// @Description 更新QSO日志
// @Tags        QsoLog
// @Accept      json
// @Produce     json
// @Param       SSOID  query   string  false    "sso auth id"
// @Param       QsoLog body    QsoLog  true     "Qso log详细信息"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/qsoLog/update [post]
func H_QsoLogUpdate(res http.ResponseWriter, req *http.Request) {
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

// @Summary     hide QSO Log
// @Description 隐藏QSO日志（用户不能物理删除，只能隐藏）
// @Tags        QsoLog
// @Accept      json
// @Produce     json
// @Param       SSOID  query   string  false    "sso auth id"
// @Param       logId  body    []int   true     "Qso logId数组"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/qsoLog/hide [post]
func H_QsoLogHide(res http.ResponseWriter, req *http.Request) {
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

// @Summary     Admin update QSO Log
// @Description 管理员更新其他用户的QSO日志
// @Tags        QsoLog
// @Accept      json
// @Produce     json
// @Param       SSOID  query   string  false    "sso auth id"
// @Param       QsoLog body    QsoLog  true     "Qso log详细信息"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/admin/qsoLog/update [post]
func H_QsoLogAdminUpdate(res http.ResponseWriter, req *http.Request) {
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

// @Summary     admin unHide QSO Log
// @Description 管理员恢复显示QSO日志
// @Tags        QsoLog
// @Accept      json
// @Produce     json
// @Param       SSOID  query   string  false    "sso auth id"
// @Param       logId  body    []int   true     "Qso logId数组"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/admin/qsoLog/unHide [post]
func H_QsoLogAdminUnhide(res http.ResponseWriter, req *http.Request) {
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

// @Summary     admin delete QSO Log
// @Description 管理员物理删除QSO日志（不可恢复）
// @Tags        QsoLog
// @Accept      json
// @Produce     json
// @Param       SSOID  query   string  false    "sso auth id"
// @Param       logId  body    []int   true     "Qso logId数组"
// @Success     200  {object}  HttpRes{data=nil}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/admin/qsoLog/delete [post]
func H_QsoLogAdminDelete(res http.ResponseWriter, req *http.Request) {
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
