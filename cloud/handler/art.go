package handler

import (
	"cloud/api"
	"cloud/app"
	"cloud/entity"
	"cloud/service"
	"cloud/util"
	"github.com/go-xorm/xorm"
	"github.com/rs/xid"
	"strconv"
	"strings"
	"time"
)

func GetSingerCity() error {

	//have account
	var sum int
	var users []entity.ArtList
	var nUsers []entity.ArtList
	err := app.DB.Table(new(entity.ArtList)).Where("account_id != ?", 0).Find(&users)
	if err != nil {
		app.Logger.Error().Msg(err.Error())
		return err
	}
	args := map[string]string{}
	for i, user := range users {
		url := "http://localhost:3000/user/detail?uid=" + strconv.Itoa(user.AccountID)
		var res entity.Account
		res, err = api.GetAccount(args, url)
		if err != nil {
			app.Logger.Error().Msg(err.Error())
			return err
		}
		users[i].City = entity.Province[res.Profile.Province]
		sum += 1
		app.Logger.Info().Msg("-----------get " + strconv.Itoa(sum) + "-----------")

	}
	_, err = app.DB.Table(new(entity.User)).Insert(&users)
	if err != nil {
		app.Logger.Error().Msg(err.Error())
		return err
	}
	_, err = app.DB.Transaction(func(sess *xorm.Session) (_ interface{}, err error) {

		var accounts []entity.User
		var accs []entity.User
		//查询没有account_id的人
		err = sess.Table(new(entity.ArtList)).Where("account_id = ?", 0).Find(&nUsers)
		if err != nil {
			app.Logger.Error().Msg(err.Error())
			return
		}
		//获取没有城市
		err = sess.Table(new(entity.User)).Where("city = ?", "").Find(&accounts)
		if err != nil {
			app.Logger.Error().Msg(err.Error())
			return
		}
		//获取有城市的人
		err = sess.Table(new(entity.User)).Where("city != ?", "").Find(&accs)
		if err != nil {
			app.Logger.Error().Msg(err.Error())
			return
		}
		userids := util.SliceAnyToSliceInterface(util.Map(accs, "UserID"))
		var art entity.ArtList
		_, err = sess.Table(new(entity.ArtList)).In("user_id", userids...).Delete(&art)
		if err != nil {
			app.Logger.Error().Msg(err.Error())
			return
		}
		_, err = sess.Table(new(entity.ArtList)).Insert(&accs)
		if err != nil {
			app.Logger.Error().Msg(err.Error())
			return
		}

		//合并所有没有城市的人
		for _, ac := range accounts {
			nUsers = append(nUsers, entity.ArtList(ac))
		}

		args := map[string]string{}
		for _, nuser := range nUsers {
			url := "http://localhost:3000/artist/desc?id=" + strconv.Itoa(nuser.UserID)
			sig := false
			var res entity.SingerDesc
			res, err = api.GetSingerDesc(args, url)
			if err != nil {
				app.Logger.Error().Msg(err.Error())
				return
			}
			if res.BriefDesc != "" {
				wor := strings.Split(res.BriefDesc, "，")
				for _, wo := range wor {
					for _, ci := range entity.City {
						if strings.Contains(wo, ci) {
							if ci == "台北" {
								ci = "台湾"
							}
							nuser.City = ci
							sig = true
							_, err = sess.Table(new(entity.ArtList)).ID(nuser.ID).Update(&nuser)
							if err != nil {
								app.Logger.Error().Msg(err.Error())
								return
							}
							break
						}
						if sig {
							break
						}
					}
					if sig {
						break
					}
				}

				if sig {
					sum += 1
					app.Logger.Info().Msg("-----------get " + strconv.Itoa(sum) + "-----------")
					continue
				}
			}
			if len(res.Introduction) != 0 {
				for _, intro := range res.Introduction {
					if intro.Ti == "早年经历" {
						wor := strings.Split(intro.Txt, "，")
						for _, ci := range entity.City {
							for _, wo := range wor {
								if strings.Contains(wo, ci) {
									if ci == "台北" {
										ci = "台湾"
									}
									nuser.City = ci
									_, err = sess.Table(new(entity.ArtList)).ID(nuser.ID).Update(&nuser)
									if err != nil {
										app.Logger.Error().Msg(err.Error())
										return
									}
									sig = true
									break
								}
							}
							if sig {
								break
							}
						}
						if sig {
							break
						}
					}
				}
				if sig {
					sum += 1
					app.Logger.Info().Msg("-----------get " + strconv.Itoa(sum) + "-----------")
					continue
				}
			}
			continue
		}
		app.Logger.Info().Msg("-----------get " + strconv.Itoa(sum) + "-----------***Over***")

		return
	})
	if err != nil {
		return err
	}

	return nil
}

func GetSingerTops() error {
	url := "http://localhost:3000/toplist/artist"
	args := map[string]string{}
	res, err := api.GetSingerTops(args, url)
	if err != nil {
		app.Logger.Error().Msg(err.Error())
		return err
	}
	var tops []entity.SingerTopList
	for i, re := range res.List.Artists {
		var top entity.SingerTopList
		top.ID = xid.New().String()
		top.UserID = re.ID
		top.Name = re.Name
		top.Alias = re.Alias
		top.LastRank = re.LastRank
		top.Ranking = i + 1
		top.LastRank = re.LastRank + 1
		top.Score = re.Score
		top.TopicPerson = re.TopicPerson
		tops = append(tops, top)
	}
	err = service.ArtList.SaveSingerTops(tops)
	if err != nil {
		app.Logger.Error().Msg(err.Error())
		return err
	}
	return nil
}

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
