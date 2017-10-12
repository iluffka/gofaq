package helper

import (
	"gofaq/terminal_app/config"
	"log"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func SetConfig(path string, cnf *config.Settings)  {
	configFile, err := os.Open(path)
	defer configFile.Close()
	ProcessError(err, "load config error: %s")

	b, err := ioutil.ReadAll(configFile)
	ProcessError(err, "config file read error: %s")

	errUM := yaml.Unmarshal(b, cnf)
	ProcessError(errUM, "unmarshal fail: %s")
}

func ProcessError(err error, formatStr string)  {
	if err != nil {
		log.Printf(formatStr, err.Error())
	}
}
