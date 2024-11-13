package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// 定义命令行参数
	dailyFlag := flag.Bool("d", false, "Download today's wallpaper")     // -d 参数
	startDateFlag := flag.String("start", "", "Start date (YYYY-MM-DD)") // -h 参数的开始日期
	endDateFlag := flag.String("end", "", "End date (YYYY-MM-DD)")       // -h 参数的结束日期

	// 解析命令行参数
	flag.Parse()

	// 判断是否有 -d 参数
	if *dailyFlag {
		// 获取今天的日期
		today := time.Now().Format("2006-01-02")
		fmt.Printf("Downloading today's wallpaper: %s\n", today)
		// 调用下载当天壁纸的函数
		downloadWallpaper(today, today)
	}

	// 判断是否有 -h 参数，且 start 和 end 日期都存在
	if *startDateFlag != "" && *endDateFlag != "" {
		fmt.Printf("Downloading wallpapers from %s to %s\n", *startDateFlag, *endDateFlag)
		// 调用下载指定日期区间壁纸的函数
		downloadWallpaper(*startDateFlag, *endDateFlag)
	}
}

// 模拟下载壁纸的函数
func downloadWallpaper(startDate, endDate string) {
	fmt.Printf("Downloading wallpapers from %s to %s...\n", startDate, endDate)
	// 在这里加入实际下载壁纸的代码
}
