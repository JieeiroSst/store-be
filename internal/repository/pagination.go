package repository

import (
	"encoding/base64"
	"errors"
	"strconv"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Paginations interface {
	Pagination(name string, items []interface{}, create domain.CreatePagination) (*domain.PaginationResponse, error)
}

type PaginationRepo struct {
	db *gorm.DB
}

func NewPaginationRepo(db *gorm.DB) *PaginationRepo {
	return &PaginationRepo{
		db: db,
	}
}

func (r *PaginationRepo) Pagination(name string, items []interface{}, create domain.CreatePagination) (*domain.PaginationResponse, error) {
	var total int64
	var tableType domain.TableNameType

	table, ok := tableType.ParseString(name)
	if !ok {
		return nil, errors.New("")
	}

	tableName := table.String()

	r.db.Table(tableName).Count(&total)

	id, err := base64.URLEncoding.DecodeString(create.After)
	if err != nil {
		return nil, err
	}

	query := r.db.Table(tableName).Where("id > ?", id)
	if create.First != "" {
		limit, err := strconv.Atoi(create.First)
		if err != nil {
			return nil, err
		}
		query = query.Limit(limit)
	}

	query.Find(&items)

	return nil, nil
}

// func (r *PaginationRepo)
