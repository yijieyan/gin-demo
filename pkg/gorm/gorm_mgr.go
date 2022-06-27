package gorm

import (
	"fmt"
	"gin_demo/pkg/conf"
	"github.com/google/martian/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type gormMgr struct {
	clientsMutex sync.Mutex
	clients      map[string]*gorm.DB
	configs      map[string]*Config
	configsOnce  sync.Once
}

var defaultMgr = &gormMgr{
	clients: make(map[string]*gorm.DB),
	configs: make(map[string]*Config),
}

func (mgr *gormMgr) fetchClient(name string) (*gorm.DB, error) {
	client := mgr.getClient(name)
	if client != nil {
		return client, nil
	}
	return mgr.addClient(name)
}

func (mgr *gormMgr) getClient(name string) *gorm.DB {
	mgr.clientsMutex.Lock()
	defer mgr.clientsMutex.Unlock()
	return mgr.clients[name]
}

func (mgr *gormMgr) addClient(name string) (*gorm.DB, error) {
	configs := mgr.getConfigs()
	cfg := configs[name]
	if cfg == nil {
		return nil, fmt.Errorf("mysql(%s) not configed", name)
	}
	db, err := mgr.newClient(cfg)
	if err == nil {
		mgr.clientsMutex.Lock()
		mgr.clients[name] = db
		mgr.clientsMutex.Unlock()
	}
	return db, nil
}

func (mgr *gormMgr) newClient(cfg *Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		sqlDb, err := db.DB()
		if err != nil {
			return nil, err
		}
		if cfg.MaxOpenConns > 0 { //设置打开数据库连接的最大数量
			sqlDb.SetMaxOpenConns(cfg.MaxOpenConns)
		}
		if cfg.MaxIdleConns > 0 { //设置空闲连接池中连接的最大数量
			sqlDb.SetMaxIdleConns(cfg.MaxIdleConns)
		}
		if cfg.MaxLifeConns > 0 { //设置了连接可复用的最大时间
			sqlDb.SetConnMaxLifetime(time.Duration(cfg.MaxLifeConns) * time.Second)
		}

		err = sqlDb.Ping()
		if err != nil {
			return nil, err
		}
		return db, nil
	}
}

func (mgr *gormMgr) getConfigs() map[string]*Config {
	mgr.configsOnce.Do(func() {
		cfgWrap := new(configWrap)
		if err := conf.Unmarshal(cfgWrap); err != nil {
			log.Errorf("unmarshal mysql configs failed:%v", err)
			return
		} else {
			for _, v := range cfgWrap.Configs {
				mgr.fillDefaultConfig(v)
			}
			mgr.configs = cfgWrap.Configs
		}
	})
	return mgr.configs
}

func (mgr *gormMgr) fillDefaultConfig(conf *Config) {
	if conf.Driver == "" {
		conf.Driver = "mysql"
	}
	if conf.DialTimeout == 0 {
		conf.DialTimeout = 3000
	}
	if conf.ReadTimeout == 0 {
		conf.ReadTimeout = 3000
	}
	if conf.WriteTimeout == 0 {
		conf.WriteTimeout = 3000
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 128
	}
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 5
	}
}
