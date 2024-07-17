package staticHandler

import _ "embed"
import "net/http"

//go:embed favicon.ico
var Favicon []byte

func H_favicon(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "image/x-icon")
	res.Header().Set("Cache-Control", "public, max-age=7776000")
	res.Write(Favicon)
}
