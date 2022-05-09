package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Application struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
}

type Webserver struct {
	Port      uint `yaml:"port"`
	Timeout   uint `yaml:"timeout"`
	RateLimit uint `yaml:"rate_limit"`
}

type Database struct {
	Tipe        string `yaml:"tipe"`
	Driver      string `yaml:"driver"`
	Address     string `yaml:"address"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Port        uint   `yaml:"port"`
	Name        string `yaml:"name"`
	MaxIdle     uint   `yaml:"maxidle"`
	MaxOpen     uint   `yaml:"maxopen"`
	MaxIdleTime uint   `yaml:"maxidletime"`
	MaxLifeTime uint   `yaml:"maxlifetime"`
	Timeout     uint   `yaml:"timeout"`
}

type Jwt struct {
	VerifyKey               string `yaml:"verifykey"`
	Key                     string `yaml:"key"`
	Tokenexpiredtime        uint   `yaml:"tokenexpiredtime"`
	Refreshtokenexpiredtime uint   `yaml:"refreshtokenexpiredtime"`
}

type Timezone struct {
	Timezone string `yaml:"timezone"`
}

type Log struct {
	Level  string   `json:"level"`
	Levels []string `json:"Levels"`
}

type Payment struct {
	IpaymuVa             string `yaml:"ipaymuva"`
	IpaymuKey            string `yaml:"ipaymukey"`
	IpaymuUrl            string `yaml:"ipaymuurl"`
	IpaymuCallbackUrl    string `yaml:"ipaymucallbackurl"`
	IpaymuTranscationUrl string `yaml:"ipaymutranscationurl"`
	SecretToken          string `yaml:"secrettoken"`
	ApiCompleted         string `yaml:"apicompleted"`
	ApiPembayaran        string `yaml:"apipembayaran"`
	ApiCancel            string `yaml:"apicancel"`
}

type Email struct {
	FromEmail         string `yaml:"fromemail"`
	FromEmailPassword string `yaml:"fromemailpassword"`
	LinkVerifyEmail   string `yaml:"linkverifyemail"`
}

type Telegram struct {
	ChatId   string `yaml:"chatid"`
	BotToken string `yaml:"bottoken"`
}

type ApplicationConfiguration struct {
	Application Application
	Webserver   Webserver
	Database    Database
	Jwt         Jwt
	Timezone    Timezone
	Log         Log
	Payment     Payment
	Email       Email
	Telegram    Telegram
}

var lock = sync.Mutex{}
var applicationConfiguration *ApplicationConfiguration

func GetConfig() *ApplicationConfiguration {
	lock.Lock()
	defer lock.Unlock()

	if applicationConfiguration == nil {
		applicationConfiguration = initConfig()
	}

	return applicationConfiguration
}

func initConfig() *ApplicationConfiguration {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config-dev")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var finalConfig ApplicationConfiguration
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		panic(err)
	}
	return &finalConfig
}
