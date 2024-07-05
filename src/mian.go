package main

import (
	"github.com/Longfei1/lorca"
	"log"
)

func main() {
	ui, err := lorca.New("https://baidu.com", "", 1024, 768, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// 绑定 Go 函数到 JS 的 'go' 对象
	_ = ui.Bind("hello", func() string {
		return "Hello from Go!"
	})

	// 等待 UI 窗口关闭
	<-ui.Done()
}
