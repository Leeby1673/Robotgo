package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("按 Ctrl+C 停止運行")

	// 設定週期按鍵 A 的時間間隔為 10 分鐘
	intervalA := 10 * time.Minute

	// 設定週期按鍵 B 的時間間隔為 30 分鐘
	intervalB := 30 * time.Minute

	// 使用 ticker 進行週期性觸發
	tickerA := time.NewTicker(intervalA)
	tickerB := time.NewTicker(intervalB)

	for {
		select {
		case <-tickerA.C:
			// 模擬按下鍵盤按鍵 A
			fmt.Println("按下按鍵 A")
			robotgo.KeyTap("a")

		case <-tickerB.C:
			// 模擬按下鍵盤按鍵 B
			fmt.Println("按下按鍵 B")
			robotgo.KeyTap("b")
		}
	}
}
