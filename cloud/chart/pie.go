package chart

import (
	"log"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func pieBase(title string, data map[string]interface{}) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: title})
	pie.Add("pie", data)
	return pie
}

func pieShowLabel() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-显示 Label"})
	pie.Add("pie", genKvData(), charts.LabelTextOpts{Show: true})
	return pie
}

func pieLabelFormatter() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-Label 格式"})
	pie.Add("pie", genKvData(), charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"})
	return pie
}

func pieRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-Radius"})
	pie.Add("pie", genKvData(),
		charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
		charts.PieOpts{Radius: []string{"40%", "75%"}},
	)
	return pie
}

func pieRoseArea() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-玫瑰图(Area)"})
	pie.Add("pie", genKvData(),
		charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
		charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "area"},
	)
	return pie
}

func pieRoseRadius(title, name string, data map[string]interface{}) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: title})
	fn := `function (params) {
		return params.name + ' : ' + params.percent+'%';
}`
	pie.Add(name, data,
		charts.LabelTextOpts{Show: true, Formatter: charts.FuncOpts(fn)},
		charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "radius"},
	)
	return pie
}

func pieRoseAreaRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-玫瑰图(Area/Radius)"})
	pie.Add("area", genKvData(),
		charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "area", Center: []string{"25%", "50%"}},
	)
	pie.Add("radius", genKvData(),
		charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
		charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "radius", Center: []string{"75%", "50%"}},
	)
	return pie
}

func pieInPie(title, outName, inName string, outData, inData map[string]interface{}) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: title})
	fn := `function (params) {
		return params.name + ' : ' + params.percent+'%';
}`
	pie.Add(inName, inData,
		charts.LabelTextOpts{Show: true, Formatter: charts.FuncOpts(fn)},
		charts.PieOpts{Radius: []string{"50%", "55%"}, RoseType: "area"},
	)
	pie.Add(outName, outData,
		charts.PieOpts{Radius: []string{"0%", "45%"}, RoseType: "radius"},
	)
	return pie
}

func PieHandler(pageTitle, title, outName, inName string, outData, inData map[string]interface{}, title2, roseName string, roseData map[string]interface{}) error {
	page := charts.NewPage(orderRouters("pie")...)
	page.PageTitle = pageTitle
	page.Add(
		/*pieBase(title, data),*/
		/*	pieShowLabel(),
			pieLabelFormatter(),
			pieRadius(),
			pieRoseArea(),*/
		pieRoseRadius(title2, roseName, roseData),
		/*pieRoseAreaRadius(),*/
		pieInPie(title, outName, inName, inData, outData),
	)
	f, err := os.Create(getRenderPath("pie.html"))
	if err != nil {
		log.Println(err)
	}
	return page.Render(f)
}
