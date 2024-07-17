package QuerySQL

import (
	"log"
	"regexp"
	"strings"
)

var default_map = map[string]string{}
var mysql_map = map[string]string{}
var postgres_map = map[string]string{}
var sqlite3_map = map[string]string{}

var SqlMap map[string]string

func LoadSQL(dbDriver, tablePrefix string) {
	SqlMap = make(map[string]string)
	if len(SqlMap) <= 0 {
		log.Println("QuerySQL Loading ...")

		var driverSqlMap = map[string]string{}

		load_default_map()
		load_mysql_map()
		load_postgres_map()
		load_sqlite3_map()

		switch dbDriver {
		case "mysql":
			driverSqlMap = mysql_map
		case "postgres":
			driverSqlMap = postgres_map
		case "sqlite3":
			driverSqlMap = sqlite3_map
		default:
			log.Fatal("unsupported DB engine")
		}
		reg := regexp.MustCompile("\\$[0-9]+")
		for k, v := range default_map {
			v2, exist := driverSqlMap[k]
			if exist {
				v = v2
			}
			v = strings.ReplaceAll(v, "dev_tab_prefix_", tablePrefix)
			if dbDriver != "sqlite3" {
				v = reg.ReplaceAllString(v, "?")
			}
			SqlMap[k] = v
		}
	}
}
