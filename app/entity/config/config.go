package config

import (
	"fmt"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var CFG = &Config{}

// Config 配置文件结构体
type Config struct {
	Stage   Stage   `json:"stage"`
	Fiber   Fiber   `json:"fiber"`
	Auth    Auth    `json:"auth"`
	Ent     Ent     `json:"ent"`
	Cache   Cache   `json:"cache"`
	Swagger Swagger `json:"swagger"`
	Log     Log     `json:"log"`
}
type Stage struct {
	Status   string `json:"status"`
	User     string `json:"user"`
	Password string `json:"password"`
	Api      Api    `json:"api"`
}
type Api struct {
	PublicGetPath []string `json:"publicGetPath,omitempty"`
	SaPathPrefix  []string `json:"saPathPrefix,omitempty"`
}
type Fiber struct {
	Host         string `json:"host"`
	Port         int16  `json:"port"`
	ReadTimeout  int    `json:"readTimeout"`
	JsonCoder    string `json:"jsonCoder"`
	TimeLocation string `json:"timeLocation"`
}

type Ent struct {
	AutoMigrate        bool   `json:"autoMigrate"`
	DebugMode          bool   `json:"debugMode"`
	WithDropIndex      bool   `json:"withDropIndex"`
	WithDropColumn     bool   `json:"withDropColumn"`
	WithGlobalUniqueID bool   `json:"withGlobalUniqueID"`
	Backend            string `json:"backend"`
	DB                 DB     `json:"DB"`
	ConnectionRetry    bool   `json:"connectionRetry"`
}
type Swagger struct {
	Enable bool `json:"enable"`
}

func LoadConfig(path string) error {
	viper.SetConfigName("config")
	// viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	err = viper.Unmarshal(&CFG)
	return err
}

// NewConfig Fiber 配置
// https://docs.gofiber.io/api/fiber#config
func (f *Fiber) NewConfig() fiber.Config {
	config := fiber.Config{
		ReadTimeout: time.Second * time.Duration(f.ReadTimeout),
	}
	jsonCoder := f.JsonCoder
	if jsonCoder == "sonic" {
		config.JSONEncoder = sonic.Marshal
		config.JSONDecoder = sonic.Unmarshal
	}
	return config
}

func (f *Fiber) Url() string {
	return fmt.Sprintf("%s:%d", f.Host, f.Port)
}
