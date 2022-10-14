package service

import "attendance_backend/src/repository"

type (
	IndexService interface {
	}
	indexService struct {
		repo repository.Factory
	}
)

func GetIndexService(repo repository.Factory) IndexService {
	return &indexService{repo: repo}
}
