package main

import (
	"image"
	_ "image/png"
	"log"
	"os"
	"github.com/Drakirus/go-flutter-desktop-embedder"
	"github.com/go-gl/glfw/v3.2/glfw"
	"io/ioutil"
	"encoding/json"
)

const configData = "config.json"

type Configuration struct {
	WIDTH   int      // 宽度
	HEIGHT  int      // 高度
	RATIO   float64  // 像素密度
	ICON    string   // 图标
	FLUTTER string   // 资源路径
	ICU     string   // ICUData
	ARGS    []string // 虚拟机参数
}

type JsonBody struct {
}

var config Configuration

func main() {
	var (
		err error
	)

	parse := makeJsonBody()
	parse.loadConfigFile(configData, &config)

	options := []gutter.Option{
		gutter.ProjectAssetPath(config.FLUTTER),
		gutter.ApplicationICUDataPath(config.ICU),
		gutter.ApplicationWindowDimension(config.WIDTH, config.HEIGHT),
		gutter.OptionWindowInitializer(setIcon),
		gutter.OptionPixelRatio(config.RATIO),
		gutter.OptionVMArguments(config.ARGS),
	}

	if err = gutter.Run(options...); err != nil {
		log.Fatalln(err)
	}

}

func setIcon(window *glfw.Window) error {
	imgFile, err := os.Open(config.ICON)
	if err != nil {
		return err
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}
	window.SetIcon([]image.Image{img})
	return nil
}

func makeJsonBody() *JsonBody {
	return &JsonBody{}
}

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
