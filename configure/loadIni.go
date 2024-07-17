package configure

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type db_t struct {
	Driver      string `toml:"driver"`
	DataSource  string `toml:"dataSource"`
	TablePrefix string `toml:"tablePrefix"`
}

type CFG_t struct {
	ServiceName string            `toml:"-"`
	Version     string            `toml:"-"`
	EnableDoc   bool              `toml:"enableDoc"`
	LogLevel    int               `toml:"logLevel"`
	Listen      string            `toml:"listen"`
	DB          db_t              `toml:"database"`
	MapA        map[string]db_t   `toml:"mapA"`
	MapB        map[string]string `toml:"mapB"`
}

const (
	None = iota
	Error
	Notice
	Info
	Most
	All
)

var CFG CFG_t

func (c *CFG_t) Load(iniFileName string) error {
	if c.MapA == nil {
		c.MapA = make(map[string]db_t)
	}
	if c.MapB == nil {
		c.MapB = make(map[string]string)
	}
	cData, err := os.ReadFile("./" + iniFileName)
	if err != nil {
		cData, err = os.ReadFile("etc/" + iniFileName)
		if err != nil {
			cData, err = os.ReadFile("../etc/" + iniFileName)
			if err != nil {
				cData, err = os.ReadFile("/etc/" + iniFileName)
				if err != nil {
					return fmt.Errorf("fail to read configure file[%s]: %v", iniFileName, err)
				}
			}
		}
	}

	err = toml.Unmarshal(cData, c)
	if err != nil {
		return fmt.Errorf("fail to Unmarshal configure file: %v", err)
	}

	if c.LogLevel > All {
		c.LogLevel = All
	} else if c.LogLevel < None {
		c.LogLevel = None
	}

	if len(c.Listen) <= 0 {
		c.Listen = "127.0.0.1:8080"
	}

	if len(c.DB.Driver) <= 0 {
		c.DB.Driver = "sqlite3"
		c.DB.DataSource = "/tmp/app.db"
	}

	if len(c.DB.Driver) > 0 && len(c.DB.DataSource) <= 0 {
		return errors.New("configure file missing section: [database.dataSource]")
	}

	if len(c.DB.TablePrefix) <= 0 {
		c.DB.TablePrefix = "devTest_"
	}

	c.ServiceName = "golang App Template"
	c.Version = "0.0.1"

	c.MapA["AA"] = db_t{Driver: "Drive_AA", DataSource: "DataSource_AA", TablePrefix: "TablePrefix_AA"}
	c.MapA["BB"] = db_t{Driver: "Drive_BB", DataSource: "DataSource_BB", TablePrefix: "TablePrefix_BB"}

	c.MapB["keyA"] = "valueA"
	c.MapB["keyB"] = "valueB"

	return nil
}

func (c *CFG_t) Dump() {
	b, err := toml.Marshal(c)
	if err == nil {
		log.Printf("\n++++++++\n%s\n++++++++\n", string(b))
	}
}
