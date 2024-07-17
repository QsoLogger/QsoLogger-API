package httpAPI

import (
	"fmt"
	"net/http"

	"github.com/QsoLogger/QsoLogger-API/sso"
)

func H_UserMyInfo(res http.ResponseWriter, req *http.Request) {
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
