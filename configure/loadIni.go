package configure

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

type db_t struct {
	Driver      string
	DataSource  string
	TablePrefix string
}

type CFG_t struct {
	serviceName string
	Version     string
	LogLevel    int
	Listen      string
	DB          db_t
}

var CFG CFG_t

func (c *CFG_t) Load(iniFileName string) error {
	cfg, err := ini.Load("./" + iniFileName)
	if err != nil {
		log.Printf("fail to read configure file[%s]: %v\n", iniFileName, err)
		cfg, err = ini.Load("etc/" + iniFileName)
		if err != nil {
			log.Printf("fail to read configure file[%s]: %v\n", iniFileName, err)
			cfg, err = ini.Load("/etc/" + iniFileName)
			if err != nil {
				return fmt.Errorf("fail to read configure file[%s]: %v", iniFileName, err)
			}
		}
	}

	if len(cfg.Section("").Key("serviceName").String()) > 0 {
		c.serviceName = cfg.Section("").Key("serviceName").String()
	} else {
		return errors.New("configure file missing section: [serviceName]")
	}

	if len(cfg.Section("").Key("version").String()) > 0 {
		c.Version = cfg.Section("").Key("version").String()
	} else {
		return errors.New("configure file missing section: [version]")
	}

	logLevel, err := cfg.Section("").Key("logLevel").Int()
	if err != nil {
		return fmt.Errorf("invalid logLevel: %v", err)
	}
	if logLevel > 5 {
		c.LogLevel = 5
	} else if logLevel < 0 {
		c.LogLevel = 0
	} else {
		c.LogLevel = logLevel
	}

	if len(cfg.Section("").Key("listen").String()) > 0 {
		c.Listen = cfg.Section("").Key("listen").String()
	} else {
		c.Listen = "127.0.0.1:8080"
	}

	if len(cfg.Section("database").Key("driver").String()) > 0 {
		c.DB.Driver = cfg.Section("database").Key("driver").String()
	} else {
		return errors.New("configure file missing section: [database.driver]")
	}

	if len(cfg.Section("database").Key("dataSource").String()) > 0 {
		c.DB.DataSource = cfg.Section("database").Key("dataSource").String()
	} else {
		return errors.New("configure file missing section: [database.dataSource]")
	}

	if len(cfg.Section("database").Key("tablePrefix").String()) > 0 {
		c.DB.TablePrefix = cfg.Section("database").Key("tablePrefix").String()
	} else {
		return errors.New("configure file missing section: [database.tablePrefix]")
	}

	return nil
}
