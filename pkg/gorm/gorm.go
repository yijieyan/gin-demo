package gorm

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type configWrap struct {
	Configs map[string]*Config `mapstructure:"mysql"`
}

type Config struct {
	Driver       string `mapstructure:"driver"`
	DSN          string `mapstructure:"dsn"`
	DialTimeout  int    `mapstructure:"dial_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxLifeConns int    `mapstructure:"max_life_conns"`
	DebugSQL     bool   `mapstructure:"debug_sql"`
}

func CheckStatus(names ...string) error {
	var mysqlNames []string
	if len(names) > 0 {
		mysqlNames = names
	} else {
		for k := range defaultMgr.getConfigs() {
			mysqlNames = append(mysqlNames, k)
		}
	}
	for _, name := range mysqlNames {
		if cli, err := getClient(name); err != nil {
			errMsg := fmt.Sprintf("mysql (%s) is invalid: %s", name, err.Error())
			return errors.New(errMsg)
		} else {
			if sqlDb, err := cli.DB(); err != nil {
				errMsg := fmt.Sprintf("mysql (%s) is config ping failed", name)
				return errors.New(errMsg)
			} else {
				if err = sqlDb.Ping(); err != nil {
					errMsg := fmt.Sprintf("mysql (%s) is ping failed", name)
					return errors.New(errMsg)
				}
			}
		}
	}
	return nil
}

func Client(name string) *gorm.DB {
	cli, err := getClient(name)
	if err != nil {
		return nil
	} else {
		return cli
	}
}

func getClient(name string) (*gorm.DB, error) {
	return defaultMgr.fetchClient(name)
}
