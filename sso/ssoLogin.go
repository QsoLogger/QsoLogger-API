package sso

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ssoUserInfo struct {
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
	UserInfo UserInfo `json:"userInfo"`
}

func AskLogin(res http.ResponseWriter, req *http.Request, app_url string) {
	app_url_hex := hex.EncodeToString([]byte(app_url))
	login_url := fmt.Sprintf("%s/ssoLogin/%s", authUrlPrefix, app_url_hex)
	http.Redirect(res, req, login_url, http.StatusFound)

}

func H_ssoUserInfo(res http.ResponseWriter, req *http.Request) {
	var dSsoUserInfo ssoUserInfo
	_, userInfo, err := GetUserInfo(res, req)
	if err != nil {
		log.Println(err)
		dSsoUserInfo.Code = 101
		dSsoUserInfo.Msg = err.Error()
	} else {
		dSsoUserInfo.UserInfo = userInfo
	}

	dSsoUserInfo_res, err := json.Marshal(dSsoUserInfo)
	if err != nil {
		res.Write([]byte("{\"code\":1,\"msg\":\"json encode failed\"}"))
		log.Println(err)
		return
	}
	res.Write(dSsoUserInfo_res)

}
