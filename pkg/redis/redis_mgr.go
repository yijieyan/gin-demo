package rds

import (
	"fmt"
	"gin_demo/pkg/conf"
	"github.com/go-redis/redis/v8"
	"github.com/google/martian/log"
	"sync"
)

type redisMgr struct {
	clientsMutex sync.Mutex
	rdsMap       map[string]*redis.Client
	configs      map[string]*Config
	configOnce   sync.Once
}

var defaultMgr = redisMgr{
	rdsMap:  make(map[string]*redis.Client),
	configs: make(map[string]*Config),
}

func (mgr *redisMgr) fetchClient(name string) (*redis.Client, error) {
	cli := mgr.getClient(name)
	if cli != nil {
		return cli, nil
	}
	return mgr.addClient(name)
}

func (mgr *redisMgr) getClient(name string) *redis.Client {
	mgr.clientsMutex.Lock()
	defer mgr.clientsMutex.Unlock()
	return mgr.rdsMap[name]
}

func (mgr *redisMgr) addClient(name string) (*redis.Client, error) {
	configs := mgr.getConfigs()
	cfg := configs[name]
	if cfg == nil {
		return nil, fmt.Errorf("redis(%s) config not found", name)
	}
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
	})

	mgr.clientsMutex.Lock()
	mgr.rdsMap[name] = cli
	mgr.clientsMutex.Unlock()
	return cli, nil

}

func (mgr *redisMgr) getConfigs() map[string]*Config {
	mgr.loadConfigs()
	return mgr.configs
}

func (mgr *redisMgr) loadConfigs() {
	mgr.configOnce.Do(func() {
		cfgWrap := new(configWrap)
		if err := conf.Unmarshal(cfgWrap); err != nil {
			log.Errorf("unmarshal redis configs failed.error:%s ", err)
			return
		}
		mgr.configs = cfgWrap.Configs
	})
}
