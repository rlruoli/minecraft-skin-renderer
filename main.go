package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"time"

	skinrender "github.com/mineatar-io/skin-render"
	"github.com/thep0y/go-logger/log"
)

func main() {
	if len(os.Args) > 1 {
		var opk skinrender.Options
		opk.Overlay = true
		opk.Slim = false
		fmt.Println("请输入将要输出的图像大小")
		fmt.Scan(&opk.Scale)
		if opk.Scale < 1 {
			exit("输入错误，请输入大于零的整数")
		}
		fmt.Println("输入序号以选择渲染模式")
		fmt.Println("1.渲染脑壳子")
		fmt.Println("2.渲染身体")
		fmt.Println("3.渲染脸")
		fmt.Println("4.渲染正视图")
		fmt.Println("5.渲染背影")
		fmt.Println("6.渲染左侧")
		fmt.Println("7.渲染右侧")
		var mod int
		fmt.Scan(&mod)

		start := time.Now()
		os.Mkdir("out", 644)
		for i := len(os.Args); i > 1; i-- {
			path := os.Args[i-1]
			log.Info(path)
			file, _ := os.Open(path)
			filename := filepath.Base(path)
			imagedata, _, err := image.Decode(file)
			if err != nil {
				log.Error(err)
			}

			out := image.NewNRGBA(image.Rect(0, 0, 0, 0))
			switch mod {
			case 1:
				out = skinrender.RenderHead(convertToNRGBA(imagedata), opk)
			case 2:
				out = skinrender.RenderBody(convertToNRGBA(imagedata), opk)
			case 3:
				out = skinrender.RenderFace(convertToNRGBA(imagedata), opk)
			case 4:
				out = skinrender.RenderFrontBody(convertToNRGBA(imagedata), opk)
			case 5:
				out = skinrender.RenderBackBody(convertToNRGBA(imagedata), opk)
			case 6:
				out = skinrender.RenderLeftBody(convertToNRGBA(imagedata), opk)
			case 7:
				out = skinrender.RenderRightBody(convertToNRGBA(imagedata), opk)
			default:
				exit("模式错误")
			}

			file, err = os.Create("out/" + filename)
			if err != nil {
				panic(err)
			}
			file.Close()

			if err := png.Encode(file, out); err != nil {
				panic(err)
			}
		}
		exit("执行完成，耗时：" + time.Since(start).String())
	} else {
		exit("请将皮肤文件拖到本程序上打开")
	}
}

func convertToNRGBA(img image.Image) *image.NRGBA {
	bounds := img.Bounds()
	nrgba := image.NewNRGBA(image.Rectangle{Max: bounds.Size()})

	draw.Draw(nrgba, nrgba.Bounds(), img, bounds.Min, draw.Src)
	return nrgba
}

func exit(info string) {
	fmt.Println(info)
	fmt.Println("五秒后退出")
	time.Sleep(time.Second * 5)
	os.Exit(0)
}
