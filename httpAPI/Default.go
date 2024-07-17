package httpAPI

import (
	"fmt"
	"net/http"
)

func H_Default(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(res, "{\"code\":404,\"msg\":\"not found\"}")
}
