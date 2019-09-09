package handler

import (
	"cloud/api"
	"cloud/app"
	"cloud/entity"
	"cloud/service"
	"github.com/rs/xid"
	"strconv"
	"time"
)

func GetArtList() error {
	var wRes []entity.ArtlistResult
	var mRes []entity.ArtlistResult
	var sum int
	app.Logger.Info().Msg("-------------------start hot men-------------------")
	mhot, err := api.GetArtList(map[string]string{
		"cat":   "1001",
		"limit": "100",
	}, "http://localhost:3000/artist/list")
	if err != nil {
		app.Logger.Error().Msg(err.Error())
		return err
	}
	sum += len(mhot.Artists)
	mRes = append(mRes, mhot.Artists...)

	app.Logger.Info().Msg("-------------------start hot women-------------------")
	whot, err := api.GetArtList(map[string]string{
		"cat":   "1002",
		"limit": "100",
	}, "http://localhost:3000/artist/list")
	if err != nil {
		app.Logger.Error().Msg(err.Error())
		return err
	}
	sum += len(whot.Artists)
	wRes = append(wRes, whot.Artists...)
	app.Logger.Info().Msg("-------------------start men-------------------")

	for i := 1; i <= 26; i++ {
		CmenRes, err := api.GetArtList(map[string]string{
			"cat":     "1001",
			"initial": entity.ArtAZ[i],
			"limit":   "100",
		}, "http://localhost:3000/artist/list")
		if err != nil {
			app.Logger.Error().Msg(err.Error())
			return err
		}
		if len(CmenRes.Artists) == 0 {
			break
		}
		println(len(CmenRes.Artists))
		mRes = append(mRes, CmenRes.Artists...)
		sum += len(CmenRes.Artists)
		time.Sleep(1 * time.Second)
		app.Logger.Info().Msg("**********get " + strconv.Itoa(sum) + "....")
	}
	var mReqs []entity.ArtList
	for _, art := range mRes {
		var req entity.ArtList
		req.ID = xid.New().String()
		req.Name = art.Name
		req.AccountID = art.AccountID
		req.UserID = art.ID
		req.AlbumSize = art.AlbumSize
		req.Cat = "1001"
		req.MusicSize = art.MusicSize
		req.PicUrl = art.PicUrl
		mReqs = append(mReqs, req)
	}
	err = service.ArtList.SaveArtList(mReqs)
	if err != nil {
		app.Logger.Error().Msg(err.Error())
		return err
	}
	app.Logger.Info().Msg("-------------------start women-------------------")
	for j := 1; j <= 26; j++ {
		CwomenRes, err := api.GetArtList(map[string]string{
			"cat":     "1002",
			"initial": entity.ArtAZ[j],
			"limit":   "100",
		}, "http://localhost:3000/artist/list")
		if err != nil {
			app.Logger.Error().Msg(err.Error())
			return err
		}
		wRes = append(wRes, CwomenRes.Artists...)
		sum += len(CwomenRes.Artists)
		time.Sleep(1 * time.Second)
		app.Logger.Info().Msg("**********get " + strconv.Itoa(sum) + "....")
	}
	var wReqs []entity.ArtList
	for _, art := range wRes {
		var req entity.ArtList
		req.ID = xid.New().String()
		req.Name = art.Name
		req.AccountID = art.AccountID
		req.UserID = art.ID
		req.AlbumSize = art.AlbumSize
		req.Cat = "1002"
		req.MusicSize = art.MusicSize
		req.PicUrl = art.PicUrl
		wReqs = append(wReqs, req)
	}
	app.Logger.Info().Msg("-------------------start save-------------------")
	err = service.ArtList.SaveArtList(wReqs)
	if err != nil {
		app.Logger.Error().Msg(err.Error())
		return err
	}

	return nil

}
