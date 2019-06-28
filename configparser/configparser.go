package configparser

import (
	"encoding/json"
	"io/ioutil"
)

const (
	INI  = 1
	JSON = 2
)

type Config struct{
	filename string
	filetype int
    res map[string]interface{}  `result of config parser`
}

func InitConfig(filename string,filetype int) *Config {
    config := &Config{}
	config.filename = filename
	config.res = make( map[string]interface{})
	// parse ini config into map
    if(filetype == INI) {

	}
	// parse json config into map
	if(filetype == JSON) {
		data, err := ioutil.ReadFile(filename)
		if(err != nil) {
			panic(err);
		}
		err = json.Unmarshal(data,&(config.res))
		if(err != nil) {
			panic(err);
		}
	}
	return config
}


func (config * Config) Set(key string ,value string) {
    config.res[key] = value
}

func (config * Config) Get(key string) interface{} {
	return config.res[key]
}


