package configparser

import (
	"encoding/json"
	"io/ioutil"
	"errors"
	"os"
)

const (
	UNKNOWN = 0
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
	config.filetype = UNKNOWN
	config.res = make( map[string]interface{})
	// parse ini config into map
    if(filetype == INI) {
        panic(errors.New("暂不支持的类型"))
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


func (config *Config) Save()  error{
	f, err := os.OpenFile(config.filename + "_bak", os.O_WRONLY|os.O_TRUNC|os.O_CREATE|os.O_SYNC, os.ModePerm)
	defer f.Close()
	if(err != nil) {
		return err
	}
	res , err :=json.Marshal(config.res)
	f.WriteString(string(res))
    return nil
}