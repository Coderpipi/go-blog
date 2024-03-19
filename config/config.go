package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
)

var Cfg *Config

type Config struct {
	BasePath string `env:"base_path"`
	Env      string `env:"ENV"`
	*App     `yaml:"app"`
	*Logger  `yaml:"logger"`
	*Mysql   `yaml:"mysql"`
}

type App struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

func InitConfig() {
	if len(os.Getenv("ENV")) == 0 {
		// if no ENV specified
		envFile := ".env"
		for i := 0; i < 10; i++ {
			if _, err := os.Stat(envFile); err != nil {
				envFile = "../" + envFile
				continue
			}
			break
		}

		err := godotenv.Load(envFile)
		if err != nil {
			panic(err)
		}
	}

	Cfg = &Config{}
	if err := env.Parse(Cfg); err != nil {
		panic(err)
	}

	file, err := os.ReadFile(fmt.Sprintf("config-%s.yaml", Cfg.Env))

	if err != nil {
		fmt.Println("读取配置文件失败 " + err.Error())
	}

	// 兜底默认配置文件
	file, err = os.ReadFile(fmt.Sprintf("config.yaml"))

	if err != nil {
		panic("读取配置文件失败 " + err.Error())
	}

	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		panic("解析失败: " + err.Error())
	}

	initLogger()

	initMysql()
}
