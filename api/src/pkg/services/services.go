package services

import (
	"errors"
	"log"
)

type Service interface {
	GetNewCasesPerLocation(location string) ([]NewCasesResponse, error)
}

// User repository is what lets our service do db operations without knowing anything about the implementation
type Repository interface {
	GetNewCasesPerLocation(string) ([]NewCasesResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (service *service) GetNewCasesPerLocation(location string) ([]NewCasesResponse, error) {
	log.Println("- Service - new cases service is being execute")

	if location == "" {
		return []NewCasesResponse{}, errors.New("Service - location required")
	}
	
	data, err := service.repo.GetNewCasesPerLocation(location)

	if err != nil {
		log.Printf("error: %v", err.Error())
		return []NewCasesResponse{}, err
	}

	return data, err
}
