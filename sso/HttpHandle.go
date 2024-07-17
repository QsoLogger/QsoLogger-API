package sso

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/QsoLogger/QsoLogger-API/configure"
)

// @Summary     get sso login URL
// @Description 获取sso登陆地址
// @Tags        sso
// @Accept      plain
// @Produce     plain
// @Param       ssoRef  query  string  false  "callback url"  "example: 当前URL"
// @Success     200  {object}  HttpRes{data=string}  "the login URL"
// @Failure     400  {object}  HttpRes{data=nil}     "API未绑定"
// @Router      /api/sso/getLoginUrl [get]
func H_ssoGetLoginUrl(res http.ResponseWriter, req *http.Request) {
	login_url := getSsoLoginUrl(res, req, appUrlPrefix(req)+"/ssoLogin")
	fmt.Fprintf(res, "{\"data\":\"%s\"}", login_url)
}

// @Summary     Force Redirect to sso login URL
// @Description 跳转到sso登陆, example: http://127.0.0.1:8080/ssoLogin?ssoRef=http://127.0.0.1:8080/docs
// @Tags        sso
// @Accept      plain
// @Produce     plain
// @Param       SSOID  query  string  false  "sso auth id"
// @Param       ssoRef query  string   false  "callback url"  "example: 当前URL"
// @Success     302  {object}  HttpRes{data=nil}    "如果有cb则跳转cb，否则跳转app首页"
// @Failure     302  {object}  HttpRes{data=nil}    "跳转sso服务登陆页"
// @Failure     400  {object}  HttpRes{data=nil}    "API未绑定"
// @Router      /ssoLogin      [get]
func H_ssoLogin(res http.ResponseWriter, req *http.Request) {
	SSOID, _, err := GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		//未登陆
		if configure.CFG.LogLevel >= configure.All {
			log.Println("sso.H_ssoLogin: user Not login")
		}
		ssoRef := req.URL.Query().Get("ssoRef")
		if len(ssoRef) > 0 {
			if configure.CFG.LogLevel >= configure.All {
				log.Println("sso.H_ssoLogin: Param ssoRef:", ssoRef)
			}
			//ssoRef写入cookie
			var ssoRef_jar http.Cookie
			ssoRef_jar.Name = "ssoRef"
			ssoRef_jar.Path = "/"
			//ssoRef_jar.Domain = "ssoServer.com"
			ssoRef_jar.HttpOnly = false
			ssoRef_jar.MaxAge = 3600 // 1小时内ssoRef，超过半小时跳app首页
			ssoRef_jar.Value = ssoRef
			http.SetCookie(res, &ssoRef_jar)
		}

		login_url := getSsoLoginUrl(res, req, appUrlPrefix(req)+"/ssoLogin")
		http.Redirect(res, req, login_url, http.StatusFound)
	} else {
		//已登陆
		if configure.CFG.LogLevel >= configure.All {
			log.Println("sso.H_ssoLogin: user has login, SSOID:", SSOID)
		}
		ssoRef := req.URL.Query().Get("ssoRef")
		if len(ssoRef) == 0 {
			if configure.CFG.LogLevel >= configure.All {
				log.Println("sso.H_ssoLogin: ssoRef Param is not set")
			}
			COOKIE_ssoRef, err := req.Cookie("ssoRef")
			if err == nil {
				ssoRef = COOKIE_ssoRef.Value
				if configure.CFG.LogLevel >= configure.All {
					log.Println("sso.H_ssoLogin: Cookie ssoRef:", ssoRef)
				}
			} else {
				if configure.CFG.LogLevel >= configure.All {
					log.Println("sso.H_ssoLogin: Cookie ssoRef is not set")
				}
			}
		} else {
			if configure.CFG.LogLevel >= configure.All {
				log.Println("sso.H_ssoLogin: Param ssoRef:", ssoRef)
			}
		}

		//ssoRef清除cookie
		var ssoRef_jar http.Cookie
		ssoRef_jar.Name = "ssoRef"
		ssoRef_jar.Path = "/"
		//ssoRef_jar.Domain = "ssoServer.com"
		ssoRef_jar.HttpOnly = false
		ssoRef_jar.MaxAge = -1
		ssoRef_jar.Value = ssoRef
		http.SetCookie(res, &ssoRef_jar)

		if len(ssoRef) > 0 {
			http.Redirect(res, req, ssoRef, http.StatusFound)
		} else {
			http.Redirect(res, req, appUrlPrefix(req), http.StatusFound)
		}
	}
	fmt.Fprintf(res, "{\"code\":0}")
}

// @Summary     get ssoInfo
// @Description 获取sso信息
// @Tags        sso
// @Accept      plain
// @Produce     json
// @Param       SSOID  query  string  false  "sso auth id"
// @Success     200  {object}  HttpRes{data=UserInfo}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /ssoInfo       [get]
func H_ssoInfo(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(res, "{\"code\":401,\"msg\":\"forbiden\"}")
		return
	}

	userInfo_s, err := json.Marshal(ssoInfo)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "{\"code\":500,\"msg\":\"json encode failed"+err.Error()+"\"}")
		return
	}
	res.Write(userInfo_s)
}
