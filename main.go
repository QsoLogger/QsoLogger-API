package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"

	"github.com/QsoLogger/QsoLogger-API/SQL/InitSQL"
	"github.com/QsoLogger/QsoLogger-API/SQL/QuerySQL"
	"github.com/QsoLogger/QsoLogger-API/configure"
	"github.com/QsoLogger/QsoLogger-API/docs"
	"github.com/QsoLogger/QsoLogger-API/httpAPI"
	"github.com/QsoLogger/QsoLogger-API/sso"
	"github.com/QsoLogger/QsoLogger-API/staticHandler"
)

var cfg *configure.CFG_t

// @title                     Golang Project Template Restful API
// @version                   0.1 no-release
// @description               This is a Golang Project Template with Restful API Documents
// @termsOfService            https://www.project.net/terms
// @contact.name              Dev Team
// @contact.url               https://www.project.net/contact
// @contact.email             contact@project.net
// @license.name              Private License 0.1
// @license.url               https://www.project.net/license
// @host                      localhost:8080
// @BasePath                  /
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// parser flag
	needHelp := flag.Bool("h", false, "show usage text")
	testFlag := flag.String("t", "", "use `\"test string\"` for args testing")
	flag.Parse()

	//if *needHelp == true {
	if *needHelp {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(*testFlag) > 0 {
		fmt.Println(*testFlag)
		os.Exit(1)
	}

	// load configure ini file
	iniFileName := "QsoLogger.ini"
	cfg = &configure.CFG
	err := cfg.Load(iniFileName)
	if err != nil {
		log.Println("Fail to read file:", err)
		os.Exit(1)
	}
	cfg.Dump()

	// init SQL
	InitSQL.InitDB(cfg.DB.Driver, cfg.DB.DataSource, cfg.DB.TablePrefix)
	QuerySQL.LoadSQL(cfg.DB.Driver, cfg.DB.TablePrefix)

	// connect DB before every query
	db, err := sqlx.Connect(cfg.DB.Driver, cfg.DB.DataSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	type dbInfo_t struct {
		Value string `db:"value"`
	}

	if cfg.LogLevel >= 3 {
		log.Println(QuerySQL.SqlMap["get.db.structure.version"])
	}
	info := dbInfo_t{}
	err = db.Get(&info, QuerySQL.SqlMap["get.db.structure.version"])
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("database structure version:", info.Value)

	appHandle := http.NewServeMux()

	appHandle.HandleFunc("/sso", staticHandler.H_Sso)

	appHandle.HandleFunc("/api/user/myInfo", httpAPI.H_UserMyInfo)
	appHandle.HandleFunc("/api/qsoLog/add", httpAPI.H_QsoLogAdd)
	appHandle.HandleFunc("/api/qsoLog/listByBookId", httpAPI.H_QsoLogListByBookId)
	appHandle.HandleFunc("/api/qsoLog/update", httpAPI.H_QsoLogUpdate)
	appHandle.HandleFunc("/api/qsoLog/hide", httpAPI.H_QsoLogHide)
	appHandle.HandleFunc("/api/qsoLog/adminUpdate", httpAPI.H_QsoLogAdminUpdate)
	appHandle.HandleFunc("/api/qsoLog/adminUnhide", httpAPI.H_QsoLogAdminUnhide)
	appHandle.HandleFunc("/api/qsoLog/adminDelete", httpAPI.H_QsoLogAdminDelete)
	appHandle.HandleFunc("/api/", httpAPI.H_Default)
	appHandle.HandleFunc("/api", httpAPI.H_Default)

	appHandle.HandleFunc("/api/sso/getLoginUrl", sso.H_ssoGetLoginUrl)
	appHandle.HandleFunc("/ssoLogin", sso.H_ssoLogin)
	appHandle.HandleFunc("/ssoInfo", sso.H_ssoInfo)

	appHandle.HandleFunc("/swagger/", docHandler)
	appHandle.HandleFunc("/swagger", docHandler)
	appHandle.HandleFunc("/docs/", docHandler)
	appHandle.HandleFunc("/docs", docHandler)
	appHandle.HandleFunc("/doc/", docHandler)
	appHandle.HandleFunc("/doc", docHandler)

	appHandle.HandleFunc("/favicon.ico", staticHandler.H_favicon)
	appHandle.HandleFunc("/", defaultHandler)

	//appBind := fmt.Sprintf("0.0.0.0:%s", "8080")
	appBind := cfg.Listen
	appServer := &http.Server{
		Addr:    appBind,
		Handler: appHandle,
	}
	appServer.SetKeepAlivesEnabled(false)
	err = appServer.ListenAndServe()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	host, port, err := net.SplitHostPort(req.Host)

	log.Println("Handler: / (defaultHandler)")
	log.Println("req.Host:\t\t", req.Host)
	log.Println("net.SplitHostPort:\t", host, port, err)
	log.Println("req.URL.RequestURI():\t", req.URL.RequestURI())
	log.Println("req.RequestURI:\t", req.RequestURI)
	log.Println("req.URL.Path:\t\t", req.URL.Path)
	log.Println("req.URL.RawQuery:\t", req.URL.RawQuery)
}

// @Summary     API Document Handler
// @Description 本API文档
// @Tags        Default
// @Accept      plain
// @Produce     html
// @Success     200  {string} string
// @Router      /docs [get]
func docHandler(res http.ResponseWriter, req *http.Request) {
	if cfg.EnableDoc != true {
		httpAPI.F_404NotFound(res, req)
	} else {
		docs.SwaggerInfo.Host = req.Host
		URI := req.URL.RequestURI()
		as1 := strings.Split(URI, "/")
		if len(as1) < 3 {
			http.Redirect(res, req, "/docs/index.html", http.StatusFound)
		} else {
			as2 := strings.Split(as1[2], "#")
			if len(as2[0]) == 0 {
				http.Redirect(res, req, "/docs/index.html", http.StatusFound)
			} else {
				httpSwagger.WrapHandler(res, req)
			}
		}
	}
}
