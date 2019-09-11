package statistic

import (
	"cloud/app"
	"cloud/chart"
	"cloud/entity"
	"cloud/util"
	"github.com/pkg/errors"
	"strconv"
)

var Sum SumStatistic

type SumStatistic struct{}

func (s SumStatistic) SingerMap() (err error) {
	res, err := app.DB.Table(new(entity.ArtList)).QueryString(`select city,count(*) from art_list where city != "" group by city`)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	var data map[string]float32
	data = make(map[string]float32)
	for _, re := range res {
		var fnum float64
		fnum, err = strconv.ParseFloat(re["count(*)"], 32)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		data[re["city"]] = float32(fnum)
	}
	chart.MapHandler("Singer distribution", "only registration singer(Cloud-music)", data)
	return
}

func (s SumStatistic) SingerTops() (err error) {
	var (
		names []string
		xData []int
		yData []int
	)
	title1 := "热度:"
	title2 := "讨论人数:"
	var records []entity.SingerTopList
	err = app.DB.Table(new(entity.SingerTopList)).Asc("ranking").Find(&records)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	title := "Singer Top List(" + records[0].CreateTime.Format()[:10] + ")"
	for i, re := range records {
		names = append(names, strconv.Itoa(i+1)+":"+re.Name)
		xData = append(xData, re.Score)
		yData = append(yData, re.TopicPerson)
		if i == 19 {
			break
		}
	}
	names = util.ReverseStringSlice(names)
	xData = util.ReverseIntSlice(xData)
	yData = util.ReverseIntSlice(yData)
	chart.BarHandler(names, title, title1, title2, xData, yData)
	return nil
}

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
