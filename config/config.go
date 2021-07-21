package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config of app
var Config map[string]interface{}

// read settings file
func init() {
	setting, _ := os.Open("unitalk.json")
	defer setting.Close()
	byteValue, _ := ioutil.ReadAll(setting)
	json.Unmarshal([]byte(byteValue), &Config)
}
