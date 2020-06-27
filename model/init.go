package model

import (
	"fmt"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pangxieke/portray/config"
	"github.com/pkg/errors"
)

var (
	RedisCluster *redis.ClusterClient
	RedisClient  *redis.Client
	db           *gorm.DB
	Mgo          *StorageMgo
)

type StorageMgo struct {
	db      string
	session *mgo.Session
}

func Init() (err error) {
	initFuncs := []func() error{
		initRedis,
		//initMgo,
		initMySQL,
	}
	for _, val := range initFuncs {
		if err = val(); err != nil {
			return err
		}
	}

	return
}

func InitForTest(d *gorm.DB) (err error) {
	db = d
	if err := Migrate(); err != nil {
		return errors.Wrapf(err, "Failed to migrate database")
	}
	db.LogMode(true)
	return
}

func initRedis() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		return errors.Wrapf(err, "Failed to connect redis")
	}
	if pong != "PONG" {
		return errors.Wrapf(err, "Failed to ping redis")
	}
	return
}

// redis cluster
func initRedisCluster() (err error) {
	opt := redis.ClusterOptions{
		Addrs:    strings.Split(config.Redis.Address, ","),
		Password: "",
		PoolSize: 10,
	}
	RedisCluster = redis.NewClusterClient(&opt)

	pong, err := RedisCluster.Ping().Result()
	if err != nil {
		return errors.Wrapf(err, "Failed to connect redis")
	}
	if pong != "PONG" {
		return errors.Wrapf(err, "Failed to ping redis")
	}
	return
}

func initMySQL() (err error) {
	host := config.MySQL.Host
	port := config.MySQL.Port
	username := config.MySQL.User
	password := config.MySQL.Password
	name := config.MySQL.DB
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, name)
	db, err = gorm.Open("mysql", args)
	if err != nil {
		return errors.Wrapf(err, "Failed to open database")
	}

	db.LogMode(config.MySQL.Debug)
	return
}

func initMgo() (err error) {
	url := config.Mgo.Urls
	db := config.Mgo.DB
	if url == "" || db == "" {
		panic(fmt.Sprintf("mongodb init error, url=%s, db=%s", url, db))
	}

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	Mgo = NewStorage(db, session)
	return
}

func NewStorage(db string, session *mgo.Session) *StorageMgo {
	return &StorageMgo{
		db:      db,
		session: session,
	}
}
