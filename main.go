package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/pkg/errors"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const configData = "config.json"

type Configuration struct {
	Project   string   // 项目路径
	Mode      int      // 窗口模式 0 = WindowModeDefault 1 = WindowModeBorderless 2 = WindowModeBorderlessFullscreen
	Width     int      // 宽度
	Height    int      // 高度
	MinWidth  int      // 最小宽度
	MinHeight int      // 最小高度
	MaxWidth  int      // 最大宽度
	MaxHeight int      // 最大高度
	Ratio     float64  // 像素密度
	ARGS      []string // 虚拟机参数
}

type JsonBody struct {
}

var config Configuration
var root string

func main() {
	var (
		err error
	)

	parse := makeJsonBody()
	parse.loadConfigFile(configData, &config)

	if err = flutter.Run(setOptions()...); err != nil {
		log.Fatalln(err)
	}
}

// 设置项目路径
func setProjectPath() {
	if config.Project == "" {
		root = "."
	} else {
		root = config.Project
	}
	fmt.Println("Project: " + filepath.ToSlash(root))
}

// 设置启动参数
func setOptions() []flutter.Option {
	setProjectPath()
	return []flutter.Option{
		flutter.ProjectAssetsPath(root + "/build/flutter_assets"),
		flutter.ApplicationICUDataPath("icudtl.dat"),
		flutter.WindowInitialDimensions(config.Width, config.Height),
		flutter.WindowDimensionLimits(config.MinWidth, config.MinHeight, config.MaxWidth, config.MaxHeight),
		flutter.WindowIcon(setIcon),
		flutter.ForcePixelRatio(config.Ratio),
		flutter.OptionVMArguments(config.ARGS),
		setWindowMode(config.Mode),
	}
}

// 设置窗口模式
func setWindowMode(mode int) flutter.Option {
	if mode == 1 {
		return flutter.WindowMode(flutter.WindowModeBorderless)
	} else if mode == 2 {
		return flutter.WindowMode(flutter.WindowModeBorderlessFullscreen)
	} else {
		return flutter.WindowMode(flutter.WindowModeDefault)
	}
}

// 设置应用程序图标
func setIcon() ([]image.Image, error) {
	execPath, err := os.Executable()
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve executable path")
	}
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to eval symlinks for executable path")
	}
	imgFile, err := os.Open(root + "/assets/icon.png")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open assets/icon.png")
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode image")
	}
	return []image.Image{img}, nil
}

// 创建JSON结构体
func makeJsonBody() *JsonBody {
	return &JsonBody{}
}

// 解析配置文件
func (jst *JsonBody) loadConfigFile(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
