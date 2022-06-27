package rds

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type configWrap struct {
	Configs map[string]*Config `mapstructure:"redis"`
}

type Config struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"database"`
}

var ctx = context.Background()

func CheckStatus(names ...string) error {
	var redisName []string
	if len(names) > 0 {
		redisName = names
	} else {
		for k := range defaultMgr.getConfigs() {
			redisName = append(redisName, k)
		}
	}
	fmt.Println(redisName)
	for _, name := range redisName {
		db, err := Client(name)
		if err != nil {
			fmt.Printf("redis(%s) is invalid:%s", name, err)
			return errors.Wrapf(err, "redis(%s) is invalid", name)
		} else {
			_, err := db.Ping(ctx).Result()
			if err != nil {
				fmt.Printf("redis(%s) is ping fail:%s", name, err)
				return errors.Wrapf(err, "redis(%s) is ping fail", name)
			}
		}
	}
	return nil
}

func Client(name string) (*redis.Client, error) {
	cli, err := defaultMgr.fetchClient(name)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// Get 获取redis的值
func Get(name string, key string) (string, error) {
	db, err := Client(name)
	if err != nil {
		return "", err
	}
	val, err := db.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Set 获取redis的值
func Set(name string, key string, val string) error {
	db, err := Client(name)
	if err != nil {
		return err
	}
	err = db.Set(ctx, key, val, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Del 删除redis的值
func Del(name string, key string) error {
	db, err := Client(name)
	if err != nil {
		return err
	}
	err = db.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
