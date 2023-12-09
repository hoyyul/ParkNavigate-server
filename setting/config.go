package setting

import (
	"ParkNavigate/config"
	"ParkNavigate/global"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "/config/config.yaml"

func InitConfig() {
	if global.Config != nil {
		return
	}
	c := &config.Config{}
	workDir, _ := os.Getwd()
	yamlConf, err := os.ReadFile(workDir + ConfigFile)
	if err != nil {
		log.Fatalf("Get yamlConf error: %s\n", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Get yaml Unmarchsal error: %s\n", err)
	}
	log.Println("Configuration file loads successfully.")
	global.Config = c
}
