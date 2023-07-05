package main

import (
	"fmt"
	"os"
	"practice/app"
	"practice/config"
)

type Options map[string]option

type option struct {
	LoggerName string `json:"logger_name"`
	LogPath    string `json:"log_path"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Level      string `json:"level"`
	Stdout     bool   `json:"stdout"`
}

func main() {
	app, err := app.InitApp(config.ConfigPath("/Users/wangluyu/OneDrive - 北京小熊博望科技有限公司/workspace/project/practice-GO/config/conf.yaml"), "Test")
	if err != nil {
		fmt.Printf("failed to create logger: %s\n", err)
		os.Exit(2)
	}
	app.Logger["Test"].Info("Hello", "name", "wong")

	//v, _ := config.NewConfig(config.ConfigPath("/Users/wangluyu/OneDrive - 北京小熊博望科技有限公司/workspace/project/practice-GO/config/conf.yaml"))
	//a := v.Get("log")
	//for _, item := range a.([]interface{}) {
	//	j, _ := json.Marshal(item)
	//	o := new(option)
	//	json.Unmarshal(j, &o)
	//	fmt.Println(o)
	//}
}
