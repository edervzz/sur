package adapters

import (
	"sur/domain/entities"
)

type UrlRepositoryStub struct {
	repository map[string]entities.Url
}

func (db *UrlRepositoryStub) Create(url entities.Url) error {
	db.repository[url.Code] = url
	return nil
}

func (db *UrlRepositoryStub) Read(tenantId entities.Tentant, code string) (*entities.Url, error) {
	url := db.repository[code]
	return &url, nil
}

func NewUrlRepositoryStub() *UrlRepositoryStub {
	return &UrlRepositoryStub{
		repository: make(map[string]entities.Url),
	}
}
