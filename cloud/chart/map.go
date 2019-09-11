package chart

import (
	"log"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func mapBase() *charts.Map {
	mc := charts.NewMap("china")
	mc.SetGlobalOptions(charts.TitleOpts{Title: "Map-示例图"})
	mc.Add("map", mapData)
	return mc
}

func mapShowLabel() *charts.Map {
	mc := charts.NewMap("china")
	mc.SetGlobalOptions(charts.TitleOpts{Title: "Map-展示 Label"})
	mc.Add("map", mapData, charts.LabelTextOpts{Show: true})
	return mc
}

func mapVisualMap(title, title2 string, mapData map[string]float32) *charts.Map {
	mc := charts.NewMap("china")
	mc.SetGlobalOptions(
		charts.TitleOpts{Title: title, Subtitle: title2},
		charts.VisualMapOpts{Calculable: true, Min: 0, Max: 615, Text: []string{"max", "min"}},
	)
	mc.Add("city", mapData)
	return mc
}

func mapGuangdong() *charts.Map {
	mc := charts.NewMap("广东")
	mc.SetGlobalOptions(
		charts.TitleOpts{Title: "Map-广东地图"},
		charts.VisualMapOpts{Calculable: true,
			InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
	)
	mc.Add("map", guangdongData)
	return mc
}

func mapShantou() *charts.Map {
	mc := charts.NewMap("汕头")
	mc.SetGlobalOptions(
		charts.TitleOpts{Title: "Map-汕头地图"},
		charts.VisualMapOpts{Calculable: true,
			InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
	)
	mc.Add("map", shantouData)
	return mc
}

func mapTheme() *charts.Map {
	mc := charts.NewMap("china")
	mc.SetGlobalOptions(
		charts.InitOpts{Theme: charts.ThemeType.Macarons},
		charts.TitleOpts{Title: "Map-设置风格"},
		charts.VisualMapOpts{Calculable: true, Max: 150},
	)
	mc.Add("map", mapData)
	return mc
}

func MapHandler(title, subTitle string, mapData map[string]float32) {
	page := charts.NewPage(orderRouters("map")...)
	page.Add(
		/*mapBase(),
		mapShowLabel(),*/
		mapVisualMap(title, subTitle, mapData),
		/*mapGuangdong(),
		mapShantou(),
		mapTheme(),*/
	)
	f, err := os.Create(getRenderPath("map.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(f)
}
