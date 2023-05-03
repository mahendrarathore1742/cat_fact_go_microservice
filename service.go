package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactServer(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*CatFact, error) {

	res, err := http.Get("https://catfact.ninja/fact")

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	Fact := &CatFact{}

	if err := json.NewDecoder(res.Body).Decode(Fact); err != nil {
		return nil, err
	}

	return Fact, nil
}
