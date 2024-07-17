package staticHandler

import _ "embed"
import "net/http"

//go:embed sso.html
var Sso []byte

func H_Sso(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.Write(Sso)
}
