package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/QsoLogger/QsoLogger-API/SQL/InitSQL"
	"github.com/QsoLogger/QsoLogger-API/SQL/QuerySQL"
	"github.com/QsoLogger/QsoLogger-API/configure"
	"github.com/QsoLogger/QsoLogger-API/httpAPI"
	"github.com/QsoLogger/QsoLogger-API/staticHandler"
)

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
	cfg := &configure.CFG
	err := cfg.Load(iniFileName)
	if err != nil {
		log.Println("Fail to read file:", err)
		os.Exit(1)
	}

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

	appHandle.HandleFunc("/ping", pingHandler)
	appHandle.HandleFunc("/ping/", pingHandler)
	appHandle.HandleFunc("/favicon.ico", staticHandler.H_favicon)

	appHandle.HandleFunc("/sso", staticHandler.H_Sso)
	appHandle.HandleFunc("/api/user/myInfo", httpAPI.H_UserMyInfo)
	appHandle.HandleFunc("/api/qsoLog/add", httpAPI.H_QsoLogAdd)
	appHandle.HandleFunc("/api/qsoLog/list", httpAPI.H_QsoLogList)
	appHandle.HandleFunc("/api/qsoLog/update", httpAPI.H_QsoLogUpdate)
	appHandle.HandleFunc("/api/qsoLog/hide", httpAPI.H_QsoLogHide)
	appHandle.HandleFunc("/api/qsoLog/adminUpdate", httpAPI.H_QsoLogAdminUpdate)
	appHandle.HandleFunc("/api/qsoLog/adminUnhide", httpAPI.H_QsoLogAdminUnhide)
	appHandle.HandleFunc("/api/qsoLog/adminDelete", httpAPI.H_QsoLogAdminDelete)
	appHandle.HandleFunc("/api/", httpAPI.H_Default)
	appHandle.HandleFunc("/api", httpAPI.H_Default)

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
	log.Println("Handler: / (defaultHandler)")

	reqPath := req.URL.Path
	if len(strings.Split(reqPath, "/")) == 3 {
		goGet := req.URL.Query().Get("go-get")
		if goGet == "1" {
			hostName := req.Host
			hostPort := 2222
			hostNamePort := ""
			if hostPort != 22 {
				hostNamePort = hostName + ":" + strconv.Itoa(hostPort)
			} else {
				hostNamePort = hostName + ":"
			}
			// mod meta: <meta name="go-import" content="root-path vcs repo-url">
			// mod meta doc: https://go.dev/ref/mod#vcs-find
			FMT := `<meta name="go-import" content="%s%s git ssh://git@%s%s.git" />`
			fmt.Fprintf(res, FMT, hostName, reqPath, hostNamePort, reqPath)

			return
		}
	}

}

func pingHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Handler: /ping (pingHandler)")
	url := "http://www.baidu.com/"
	http.Redirect(res, req, url, http.StatusFound)
}
