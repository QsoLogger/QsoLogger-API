package httpAPI

import (
	"encoding/json"
	"net/http"

	"github.com/QsoLogger/QsoLogger-API/sso"
)

// @Summary     get UserInfo
// @Description 获取用户信息
// @Tags        User
// @Accept      plain
// @Produce     json
// @Param       SSOID  query   string  false  "sso auth id"
// @Success     200  {object}  HttpRes{data=sso.UserInfo}
// @Failure     401  {object}  HttpRes{data=nil}  "未登陆或SSO验证失败"
// @Failure     404  {object}  HttpRes{data=nil}  "API未绑定"
// @Router      /api/user/myInfo [get]
func H_UserMyInfo(res http.ResponseWriter, req *http.Request) {
	SSOID, ssoInfo, err := sso.GetUserInfo(res, req)
	if err != nil || len(SSOID) == 0 {
		F_401Unauthorized(res, req)
		return
	}

	// if req.Method != "POST" {
	// 	H_Default(res, req)
	// 	return
	// }

	data := ssoInfo
	httpRes := HttpRes{}
	httpRes.Data = data

	httpRes_s, err := json.Marshal(httpRes)
	if err != nil {
		F_500ServerError(res, req, "json encode failed: "+err.Error())
		return
	}
	res.Write(httpRes_s)
}
