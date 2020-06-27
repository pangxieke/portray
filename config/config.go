package config

import (
	"errors"
	"github.com/pangxieke/portray/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"strings"
)

type ServerConfig struct {
	Port    uint
	LogFile string
	Env     string
}

type MgoConfig struct {
	Urls string
	DB   string
}

type RedisConfig struct {
	Address string
}

type MySQLConfig struct {
	Host     string
	Port     uint
	User     string
	Password string
	DB       string
	Debug    bool
}

type OSSConfig struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
}

type GRPCConfig struct {
	Beauty string
}

var (
	Server ServerConfig
	Mgo    MgoConfig
	Redis  RedisConfig
	OSS    OSSConfig
	MySQL  MySQLConfig
	GRPC   GRPCConfig
)

func Init(configPath ...string) (err error) {
	if err = setUp(configPath...); err != nil {
		return
	}

	initFuncs := []func() error{
		initServer,
		initMgo,
		initRedis,
		initOSS,
		initMySQL,
		initGRPC,
	}
	for _, val := range initFuncs {
		if err = val(); err != nil {
			return err
		}
	}

	return
}

func setUp(configPaths ...string) (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	err = viper.ReadInConfig()
	if err != nil {
		log.Info("Failed to read config file (but environment config still affected), err = ", err)
		err = nil
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return
}

func initServer() (err error) {
	Server.LogFile = viper.GetString("server.log")
	if Server.LogFile == "" {
		return errors.New("server.log should not be empty")
	}
	Server.Port = viper.GetUint("server.port")
	if Server.Port == 0 {
		return errors.New("server.port should not be empty")
	}
	return
}

func initMgo() (err error) {
	Mgo.Urls = viper.GetString("mgo.urls")
	Mgo.DB = viper.GetString("mgo.db")

	if Mgo.Urls == "" {
		return errors.New("mgo.host should not be empty")
	}
	if Mgo.DB == "" {
		return errors.New("mgo.db should not be empty")
	}
	return
}

func initRedis() (err error) {
	Redis.Address = viper.GetString("redis.address")
	if Redis.Address == "" {
		return errors.New("redis.address should not be empty")
	}
	return
}

func initMySQL() (err error) {
	MySQL.Host = viper.GetString("mysql.host")
	MySQL.Port = viper.GetUint("mysql.port")
	MySQL.User = viper.GetString("mysql.user")
	MySQL.Password = viper.GetString("mysql.password")
	MySQL.DB = viper.GetString("mysql.db")
	if MySQL.DB == "" {
		return errors.New("mysql.db should not be empty")
	}
	MySQL.Debug = viper.GetBool("mysql.debug")
	return
}

// aliyun oss
func initOSS() (err error) {
	OSS.Endpoint = viper.GetString("oss.endpoint")
	if OSS.Endpoint == "" {
		return errors.New("oss.endpoint should not be empty")
	}
	OSS.AccessKeyId = viper.GetString("oss.access_key_id")
	if OSS.AccessKeyId == "" {
		return errors.New("oss.access_key_id should not be empty")
	}
	OSS.AccessKeySecret = viper.GetString("oss.access_key_secret")
	if OSS.AccessKeySecret == "" {
		return errors.New("oss.access_key_secret should not be empty")
	}
	OSS.Bucket = viper.GetString("oss.bucket")
	if OSS.Bucket == "" {
		return errors.New("oss.bucket should not be empty")
	}
	return
}

func initGRPC() (err error) {
	addresses := map[string]*string{
		"grpc.beauty": &GRPC.Beauty,
	}
	for k, v := range addresses {
		a := viper.GetString(k)
		if err = checkReachable(k, a); err != nil {
			return
		}
		*v = a
	}
	return
}

func checkReachable(name, address string) error {
	if address == "" {
		return errors.New(name + " should not be empty")
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return errors.New(name + " address is unreachable, address:" + address)
	}
	defer conn.Close()
	return nil
}
