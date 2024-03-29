package app

import (
	"cloud/entity"
	_ "github.com/Go-SQL-Driver/mysql"
)

func Migrate() error {
	if err := DB.Sync2(new(entity.ArtList), new(entity.SingerTopList), new(entity.User)); err != nil {
		Logger.Error().Msg("db migration error.")
		return err
	}
	return nil
}
