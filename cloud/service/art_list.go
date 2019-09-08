package service

import (
	"cloud/app"
	"cloud/entity"
	"fmt"
	"github.com/pkg/errors"
)

var ArtList ArtListService

type ArtListService struct{}

func (a *ArtListService) SaveArtList(req []entity.ArtList)(err error) {
	if len(req) == 0{
		return
	}
	_,err = app.DB.Table(new(entity.ArtList)).Insert(&req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	var response []string
	var olds entity.ArtList
	var ee entity.ArtList
	_,err = app.DB.Where("id != ","").Get(&ee)
	fmt.Println(ee)

	res,err := app.DB.QueryString("select min(id) from art_list group by name")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	for _,rid := range res{
		response = append(response,rid["min(id)"])
	}
	_,err = app.DB.Table(new(entity.ArtList)).NotIn("id",response).Delete(olds)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}