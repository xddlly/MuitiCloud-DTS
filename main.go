package main

import (
	"MultiCloud-DTS/configparser"
	"fmt"
)

func main() {
	//var logLevel = flag.String("l", "info", "for log level")
	//var confPath = flag.String("f", "./conf/config.ini", "for conf path")
	//initAll(*logLevel,*confPath);
	// 全量dump数据
	//log := log.InitLog("test.log","info")
	// 解析配置文件
	config := configparser.InitConfig("./config/config.json",configparser.JSON);

	master := config.Get("master")
	fmt.Println(master.(map[string]interface{})["host"])
	config.Save()
}


// 初始化
func initAll(logLevel string,confPath string){
    // 初始化日志

    // 初始化信号

    // 初始化偏移量

    // 初始化mysql连接
}

// 初始化日志
func initLog() {

}

// 初始化信号
func initSignal() {

}

// 初始化偏移量
func initMasterInfo() {

}

// 初始化mysql
func initMysql() {

}