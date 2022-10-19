package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

var Rdb *redis.Client
var rdbPrefix = "lo" // 添加前缀避免redis重复

type config struct {
	Name     string
	Password string
	Port     string
	Database int
}

// read log config
func readConfig() config {
	conf := config{}
	conf.Name = viper.Get("redis.name").(string)
	conf.Port = viper.Get("redis.port").(string)
	conf.Password = viper.Get("redis.password").(string)
	conf.Database = viper.Get("redis.database").(int)
	return conf
}

func InitRedis() {
	conf := readConfig()
	Rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:" + conf.Port,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		Password:     conf.Password, // no password set
		DB:           conf.Database, // use default DB
	})
	_, err := Rdb.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("Fatal error redis init: %s \n", err))
	}
}

func SetKey(key string, val interface{}, ex time.Duration) error {
	key = rdbPrefix + key
	if err := Rdb.Set(key, val, ex).Err(); err != nil {
		return err
	}
	return nil
}
func DelKey(key string) {
	key = rdbPrefix + key
	_ = Rdb.Del(key)
	//fmt.Println(some,"sas")
	return
}

// GetStrKey 获取strKey
func GetStrKey(key string) (string, error) {
	key = rdbPrefix + key
	val, err := Rdb.Get(key).Result()
	if err != nil { //  err == redis.Nil
		return "", err
	}
	return val, nil
}

// GetByteKey 获取字节key
func GetByteKey(key string) ([]byte, error) {
	key = rdbPrefix + key
	val, err := Rdb.Get(key).Bytes()
	if err != nil { //  err == redis.Nil
		return nil, err
	}
	return val, nil
}

// ValidCode 剩余时间小于三十秒重新发送
func ValidCode(key string) bool {
	key = rdbPrefix + key
	ttl := Rdb.TTL(key)
	if ttl.Val() < 30 {
		return false
	}
	return true
}

// BatchHashSet 批量存入哈希值
func BatchHashSet(key string, fields map[string]interface{}) (bool, error) {
	key = rdbPrefix + key
	_, err := Rdb.HMSet(key, fields).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

// BatchHashGet 批量获取哈希值
func BatchHashGet(key string, fields ...string) (map[string]interface{}, error) {
	key = rdbPrefix + key
	resMap := make(map[string]interface{})
	for _, field := range fields {
		var result interface{}
		val, err := Rdb.HGet(key, fmt.Sprintf("%s", field)).Result()
		switch {
		case err == redis.Nil:
			//base.BeeLogs.Info("Key Not Exists")
			resMap[field] = result
		case err != nil:
			//base.BeeLogs.Error(err.Error())
			resMap[field] = result
			return nil, err
		}
		switch val {
		case "":
			resMap[field] = val
		default:
			resMap[field] = result
		}
	}
	return resMap, nil
}

// HashSet 存入单个哈希值
func HashSet(key, field string, data interface{}) error {
	key = rdbPrefix + key
	err := Rdb.HSet(key, field, data).Err()
	if err != nil {
		return err
	}
	return nil
}

// HashGet 获取单个Str哈希值
func HashGet(key, field string) (string, error) {
	key = rdbPrefix + key
	result := ""
	val, err := Rdb.HGet(key, field).Result()
	switch {
	case err == redis.Nil:
		//base.BeeLogs.Info("Key Not Exists")
		return result, nil
	case err != nil:
		//base.BeeLogs.Error(err.Error())
		return result, err
	}
	return val, nil
}
