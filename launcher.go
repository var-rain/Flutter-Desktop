package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type RunConfig struct {
	Name           string `json:"name"`           // Flutter
	Request        string `json:"request"`        // attach
	DeviceId       string `json:"deviceId"`       // flutter-tester
	ObservatoryUri string `json:"observatoryUri"` // dynamic
	Type           string `json:"type"`           // dart
}

type Launcher struct {
	Version        string        `json:"version"`        // 0.2.0
	Configurations []interface{} `json:"configurations"` // RunConfig
}

var url string
var path string
var launchConfig string

const file string = "/.vscode/launch.json"

func main() {
	cmd := exec.Command("./flutter_engine.exe")
	out, _ := cmd.StdoutPipe()
	err, _ := cmd.StderrPipe()
	defer out.Close()
	defer err.Close()
	e := cmd.Start()
	if e != nil {
		fmt.Println(e)
	}
	go sync(out)
	go sync(err)
	wait := cmd.Wait()
	if wait != nil {
		fmt.Println(wait)
	}
}

// 获取输出
func sync(read io.ReadCloser) {
	buf := make([]byte, 1024, 1024)
	for {
		length, _ := read.Read(buf)
		if length > 0 {
			bytes := buf[:length]
			msg := string(bytes)
			fmt.Print(msg)
			if url == "" || path == "" {
				subStrings(msg)
			} else {
				if launchConfig == "" {
					makeJson()
					saveConfig()
				}
			}
		}
	}
}

// 获取参数
func subStrings(text string) {
	begin := "http://"
	if strings.Contains(text, begin) {
		index := strings.Index(text, begin)
		temp := text[index : len(text)-1]
		temp = strings.TrimSuffix(temp, "\r")
		url = strings.TrimSuffix(temp, "\n")
	} else {
		start := "Project:"
		index := strings.Index(text, start)
		temp := text[index+len(start)+1 : len(text)-1]
		temp = strings.TrimSuffix(temp, "\r")
		path = strings.TrimSuffix(temp, "\n")
	}
}

// 保存配置文件
func saveConfig() {
	content := []byte(launchConfig)
	saveJsonConfig := path + file
	err := os.MkdirAll(filepath.Dir(saveJsonConfig), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(saveJsonConfig, content, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 生成启动配置
func makeJson() {
	launcher := &Launcher{
		Version: "0.2.0",
		Configurations: []interface{}{
			&RunConfig{
				Name:           "Flutter",
				Request:        "attach",
				DeviceId:       "flutter-tester",
				Type:           "dart",
				ObservatoryUri: url,
			},
		},
	}
	bytes, _ := json.Marshal(launcher)
	launchConfig = string(bytes)
}
