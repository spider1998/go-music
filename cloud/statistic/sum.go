package statistic

import (
	"cloud/app"
	"cloud/chart"
	"cloud/entity"
	"github.com/pkg/errors"
)

var Sum SumStatistic

type SumStatistic struct{}

func (s *SumStatistic) CNSingerForMusic() (err error) {
	var (
		mSin      []entity.ArtList
		mMusicSum int
		mAlbumSum int
		wSin      []entity.ArtList
		wMusicSum int
		wAlbumSum int
	)

	err = app.DB.Table(new(entity.ArtList)).Where("cat = ?", "1001").Find(&mSin)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	for _, msi := range mSin {
		mMusicSum += msi.MusicSize
		mAlbumSum += msi.AlbumSize
	}
	err = app.DB.Table(new(entity.ArtList)).Where("cat = ?", "1002").Find(&wSin)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	for _, wsi := range wSin {
		wMusicSum += wsi.MusicSize
		wAlbumSum += wsi.AlbumSize
	}
	outData := map[string]interface{}{
		"华语男歌手": mMusicSum,
		"华语女歌手": wMusicSum,
	}
	inData := map[string]interface{}{
		"华语男歌手": mAlbumSum,
		"华语女歌手": wAlbumSum,
	}
	roseData := map[string]interface{}{
		"男歌手": len(mSin),
		"女歌手": len(wSin),
	}
	err = chart.PieHandler("Sum", "华语歌手音乐-专辑占比", "音乐", "专辑", outData,
		inData, "华语歌手性别比例", "sum", roseData)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
