package sso

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
)

type UserInfo struct {
	UserName string `json:"userName"`
	Type     int    `json:"type"`
}

var userCache sync.Map

func GetUserInfo(res http.ResponseWriter, req *http.Request) (string, UserInfo, error) {
	var SSOID string

	SSOID = req.URL.Query().Get("SSOID")
	if len(SSOID) <= 0 {
		log.Println("getUserInfo: no query param, checking cookie")
		COOKIE_SSOID, err := req.Cookie("SSOID")
		if err != nil {
			log.Println("getUserInfo: no query param, no cookie")
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
		userInfo_res, err := httpGet(authUrlPrefix + "/ssoUserInfo?SSOID=" + SSOID)
		if err != nil {
			log.Println(err)
			return "", userInfo, err
		}
		var dSsoUserInfo ssoUserInfo
		err = json.Unmarshal(userInfo_res, &dSsoUserInfo)
		if err != nil {
			log.Println(err)
			return "", userInfo, err
		}
		if dSsoUserInfo.Code != 0 {
			log.Println(dSsoUserInfo.Msg)
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
	//SSOID_jar.Domain = "jikedata.com"
	SSOID_jar.HttpOnly = false
	SSOID_jar.MaxAge = 1576800000 // 60*60*24*365*50
	SSOID_jar.Value = SSOID
	http.SetCookie(res, &SSOID_jar)

	return SSOID, userInfo, nil
}
