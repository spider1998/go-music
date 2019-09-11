package statistic

import (
	"cloud/app"
)

func Boot() {
	/*err := Sum.CNSingerForMusic()
	if err != nil {
		app.Logger.Error().Msg("error chart: " + err.Error())
	}*/
	/*err := Sum.SingerTops()
	if err != nil {
		app.Logger.Error().Msg("error chart: " + err.Error())
	}*/
	err := Sum.SingerMap()
	if err != nil {
		app.Logger.Error().Msg("error chart: " + err.Error())
	}

}
