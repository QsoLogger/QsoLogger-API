package sso

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/QsoLogger/QsoLogger-API/configure"
)

// @Summary		get ssoInfo
// @Description	获取sso信息
// @Tags			sso
func GetUserInfo(res http.ResponseWriter, req *http.Request) (string, UserInfo, error) {
	cfg := configure.CFG

	var SSOID string

	type ssoUserInfo struct {
		Code     int      `json:"code"`
		Msg      string   `json:"msg"`
		UserInfo UserInfo `json:"userInfo"`
	}

	SSOID = req.URL.Query().Get("SSOID")
	if len(SSOID) <= 0 {
		if cfg.LogLevel >= configure.All {
			log.Println("sso.getUserInfo: no query param, checking cookie")
		}
		COOKIE_SSOID, err := req.Cookie("SSOID")
		if err != nil {
			if cfg.LogLevel >= configure.All {
				log.Println("sso.getUserInfo: no query param, no cookie")
			}
			return "", UserInfo{}, errors.New("missing SSOID")
		}
		SSOID = COOKIE_SSOID.Value
	}

	if len(SSOID) <= 0 {
		return "", UserInfo{}, errors.New("invalide SSOID")
	}

	var userInfo UserInfo
	v, exist := userCache.Load(SSOID)
	if !exist {
		userInfo_res, err := httpGetByJson(authUrlPrefix + "/ssoUserInfo?SSOID=" + SSOID)
		if err != nil {
			if cfg.LogLevel >= configure.All {
				log.Println("sso.GetUserInfo:", err)
			}
			return "", userInfo, err
		}
		var dSsoUserInfo ssoUserInfo
		err = json.Unmarshal(userInfo_res, &dSsoUserInfo)
		if err != nil {
			if cfg.LogLevel >= configure.All {
				log.Println("sso.GetUserInfo:", err)
			}
			return "", userInfo, err
		}
		if dSsoUserInfo.Code != 0 {
			if cfg.LogLevel >= configure.All {
				log.Println("sso.GetUserInfo:", dSsoUserInfo.Msg)
			}
			return "", userInfo, errors.New(dSsoUserInfo.Msg)
		}
		userInfo = dSsoUserInfo.UserInfo

		userCache.Store(SSOID, userInfo)
	} else {
		userInfo = v.(UserInfo)
	}

	var SSOID_jar http.Cookie
	SSOID_jar.Name = "SSOID"
	SSOID_jar.Path = "/"
	//SSOID_jar.Domain = "ssoServer.com"
	SSOID_jar.HttpOnly = false
	SSOID_jar.MaxAge = 1576800000 // 60*60*24*365*50
	SSOID_jar.Value = SSOID
	http.SetCookie(res, &SSOID_jar)

	return SSOID, userInfo, nil
}
