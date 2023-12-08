package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"time"

	"github.com/wcharczuk/go-chart"
)

func generateStockChart(w http.ResponseWriter, r *http.Request) {
	// 创建一个新的图表
	graph := chart.Chart{
		Title: "股票价格",
		XAxis: chart.XAxis{
			Name:      "日期",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "价格",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
	}

	rand.Seed(time.Now().UnixNano())
	// 假设您有一个包含股票价格的时间序列数据
	// 这里我们用简单的示例数据来说明
	// prices := []float64{100.5, 101.2, 99.8, 102.1, 103.5}
	prices := []float64{
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0,
		rand.Float64()*(200.0-100.0) + 100.0}

	// 创建一个新的线图系列
	series := chart.ContinuousSeries{
		Name:    "股票价格",
		XValues: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, // 时间序列
		YValues: prices,                                   // 价格数据
	}

	// 将线图系列添加到图表中
	graph.Series = []chart.Series{series}

	// 设置图表的大小和分辨率
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	graph.Background = chart.Style{
		Padding: chart.Box{
			Top: 40,
		},
	}
	graph.Width = 800
	graph.Height = 600
	graph.DPI = 96

	// 将图表渲染为图像数据
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		http.Error(w, "生成股票图表时发生错误", http.StatusInternalServerError)
		return
	}

	// 设置HTTP响应头部为image/png
	w.Header().Set("Content-Type", "image/png")

	// 将图像数据写入HTTP响应主体
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		http.Error(w, "写入HTTP响应时发生错误", http.StatusInternalServerError)
		return
	}
}

func main() {
	// 注册处理程序
	http.HandleFunc("/stock_chart", generateStockChart)

	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
}
