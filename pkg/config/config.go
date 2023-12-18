package config

import (
	"encoding/json"
	"log"
	"os"
)

type Cfg struct {
	SOARname   string
	SOASerial  string
	SOARefresh string
	SOARetry   string
	SOAExpire  string
	SOATTL     string
	NS         []string
}

var (
	Config     *Cfg
	ConfigFile string
)

func Init(configFile string) error {
	ConfigFile = configFile
	content := []byte(`{}`)
	_, err := os.Stat(ConfigFile)
	if !os.IsNotExist(err) {
		content, err = os.ReadFile(ConfigFile)
		if err != nil {
			return err
		}
	}
	if len(content) == 0 {
		content = []byte(`{}`)
	}
	err = json.Unmarshal(content, &Config)
	if err != nil {
		return err
	}

	log.Printf("Finished reading config.")
	return nil
}
