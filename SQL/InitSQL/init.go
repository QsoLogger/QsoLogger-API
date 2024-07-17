package InitSQL

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var default_map = map[string]string{}
var mysql_map = map[string]string{}
var postgres_map = map[string]string{}
var sqlite3_map = map[string]string{}

var initFlag = 0

func InitDB(dbDriver, dataSource, tablePrefix string) {
	var driverSqlMap = map[string]string{}
	var sqlMap []string

	strGetCurrentVersionSql := "SELECT value AS version FROM dev_tab_prefix_system WHERE param='data_structure_version'"
	strSetCurrentVersionSql := ""

	if initFlag == 0 {
		log.Println("InitSQL Loading ...")
		load_default_map()
		load_mysql_map()
		load_postgres_map()
		load_sqlite3_map()

		var dbHeadVersion int
		dbHeadVersion = 0

		switch dbDriver {
		case "mysql":
			driverSqlMap = mysql_map
			strSetCurrentVersionSql = "UPDATE dev_tab_prefix_system SET value=?,updateTimestamp=UNIX_TIMESTAMP() WHERE param='data_structure_version'"
		case "postgres":
			driverSqlMap = postgres_map
			strSetCurrentVersionSql = "UPDATE dev_tab_prefix_system SET value=?,updateTimestamp=CAST(CURRENT_TIMESTAMP() AS INT) WHERE param='data_structure_version'"
		case "sqlite3":
			driverSqlMap = sqlite3_map
			strSetCurrentVersionSql = "UPDATE dev_tab_prefix_system SET value=?,updateTimestamp=strftime('%s','now') WHERE param='data_structure_version'"
		default:
			log.Fatal("unsupported DB engine")
		}
		for i := 0; i <= 9999; i++ {
			k := fmt.Sprintf("%04d", i)
			v, exist := driverSqlMap[k]
			if !exist {
				k := fmt.Sprintf("%04d", i)
				v, exist = default_map[k]
			}
			if exist {
				// adjust table prefix
				v = strings.ReplaceAll(v, "dev_tab_prefix_", tablePrefix)
				sqlMap = append(sqlMap, v)
				dbHeadVersion = i + 1
			}
		}

		if dbHeadVersion != len(sqlMap) {
			log.Fatal("SQL MAP index Error: index is not same as file count")
		}
		dbHeadVersion--

		db, err := sqlx.Connect(dbDriver, dataSource)
		if err != nil {
			log.Fatalln(err)
		}

		strGetCurrentVersionSql = strings.ReplaceAll(strGetCurrentVersionSql, "dev_tab_prefix_", tablePrefix) //adjust table prefix

		strSetCurrentVersionSql = strings.ReplaceAll(strSetCurrentVersionSql, "dev_tab_prefix_", tablePrefix) //adjust table prefix

		var dbVersion int
		err = db.Get(&dbVersion, strGetCurrentVersionSql)
		if err != nil {
			dbVersion = 0
		}

		if dbVersion < dbHeadVersion {
			log.Println("DB data structure is need to update")
			log.Println("current version:", dbVersion)
			log.Println("head version:", dbHeadVersion)
			ver := dbVersion
			if ver != 0 {
				ver++
			}
			for ; ver <= dbHeadVersion; ver++ {
				log.Printf("apply DB data structure version: %d, SQL:\n%s\n\n", ver, sqlMap[ver])
				_, err = db.Exec(sqlMap[ver])
				if err != nil {
					log.Fatalln(err)
				}
				_, err = db.Exec(strSetCurrentVersionSql, strconv.Itoa(ver))
				if err != nil {
					log.Fatalln(err)
				}
				log.Println("applied version:", ver)
			}
			_, err = db.Exec(strSetCurrentVersionSql, strconv.Itoa(dbHeadVersion))
			if err != nil {
				log.Fatalln(err)
			}
			initFlag = 1
		} else {
			log.Println("DB data structure is up to date")
		}
	}
}
