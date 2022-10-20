package ports

import "sur/domain/entities"

type IUrlRepository interface {
	Create(url entities.Url) error
	Read(tenantId entities.Tentant, code string) (*entities.Url, error)
}
