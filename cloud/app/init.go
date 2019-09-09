package app

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/rs/zerolog"
	"os"
)

var (
	Conf   Config         // 系统配置
	Logger zerolog.Logger // 全局日志
	DB     *xorm.Engine   // 全局 DB 实例
)

//-----------------------------初始化配置----------------------------------------------------------------------

func Init() error {
	var err error

	Conf, err = LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	//-----------------------------配置日志及日志文件存储----------------------------------------------------------------------
	leveledLogger := NewLeveledLogger(Conf.ConfPath + "/logs")
	if Conf.Debug {
		Logger = zerolog.New(zerolog.MultiLevelWriter(leveledLogger, zerolog.ConsoleWriter{Out: os.Stderr})).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	} else {
		Logger = zerolog.New(leveledLogger).Level(zerolog.InfoLevel).With().Timestamp().Logger()
	}
	//调用进程ID
	Logger.Info().Int("pid", os.Getpid()).Msg("app booted.")

	Logger.Info().Interface("config", Conf).Msg("loaded config.")

	//-----------------------------连接数据库----------------------------------------------------------------------

	Logger.Info().Msg("start load db...")
	DB, err = LoadDB(Conf.Mysql)
	if err != nil {
		fmt.Println(err)
	}
	Logger.Info().Msg("loaded db.")

	//-----------------------------映射数据表----------------------------------------------------------------------

	Logger.Info().Msg("migrate db...")
	err = Migrate()
	if err != nil {
		fmt.Println(err)
	}
	Logger.Info().Msgf("applied migrations...")

	err = Cron.StartCorn()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
