package engine

import (
	"cloud/app"
	"cloud/statistic"
	"fmt"
)

func Run() {
	app.Logger.Info().Msg("start server...")
	fmt.Println("Running...")
	//handler.Boot()
	statistic.Boot()
}
