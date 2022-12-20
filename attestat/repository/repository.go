package repository

import (
	"skillbox/attestat/models"
	"skillbox/attestat/repository/post"
)

// PostRepo explain...
type PostRepo interface {
	//Update(city *cities.CityRequest) (string, error)
	Create(city *cities.CityRequest) (*cities.CityRequest, error)
	Delete(id int) error
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(start, end int) ([]string, error)
	GetFromFoundation(start, end int) ([]string, error)
	GetFull(id int) (*cities.City, error)
	/*
		Fetch(ctx context.Context, Name string) ([]*models.RequestSelect, error)
		Create(ctx context.Context, p *models.RequestSelect) (int64, error)
		Delete(ctx context.Context, id int64) (bool, error)
		Update(ctx context.Context, p *models.RequestMakeFriend) (*models.RequestMakeFriend, error)*/
}

type Service struct {
	PostRepo
}

func NewRepository(db *post.DataBase) *Service {
	return &Service{
		PostRepo: post.NewCityListDB(db),
	}
}
