package db

import "github.com/comsma/knead/pkg/domain"

type Inspector interface {
	GetTableInfo(name string) (*domain.Table, error)
}
