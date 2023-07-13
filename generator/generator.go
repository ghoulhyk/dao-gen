package generator

import (
	"daogen/conf"
	"daogen/utils"
	"fmt"
	"path/filepath"
	"time"
)

// Generate
// 23-07-12	1.0.0
func Generate(basicPath string, confFilePath string) {
	startTime := time.Now()
	fmt.Println("╭＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝╮")
	fmt.Println("│　　　　　　　　　　　　　　　　　│")
	fmt.Println("│　　ＤＡＴＥ　２３－０７－１２　　│")
	fmt.Println("│　　ＶＥＲＳＩＯＮ　１．０．０　　│")
	fmt.Println("│　　　　　　　　　　　　　　　　　│")
	fmt.Println("╰＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝╯")
	fmt.Println("配置文件：", confFilePath)
	fmt.Println("生成文件基础目录：", basicPath)
	basicPath = filepath.Clean(basicPath)
	if !utils.Exist(confFilePath) {
		panic("配置文件不存在")
	}
	_, err := conf.Init(confFilePath)
	if err != nil {
		panic(fmt.Sprintf("配置文件解析失败[%v]", err.Error()))
	}
	tableDefPath := filepath.Join(basicPath, conf.GetIns().Path2basic.Schema)
	tableList := decodeTableList(tableDefPath)

	initialize(basicPath, tableList)
	generateFixedTmpl(basicPath)
	generateFragmentTmpl(basicPath, tableList)
	generateOrmFixedTmpl(basicPath)
	generateOrmFragmentTmpl(basicPath, tableList)

	finishTime := time.Now()
	fmt.Printf("生成完成:[%v]\n", finishTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("耗时:[%v秒]\n", float32(finishTime.UnixMilli()-startTime.UnixMilli())/1000)
}
