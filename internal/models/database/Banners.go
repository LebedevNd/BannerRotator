package database

import (
	"database/sql"
	"github.com/LebedevNd/BannerRotator/internal/models"
)

type BannerModel struct {
	DB *sql.DB
}

func (m *BannerModel) Insert(description string) (int, error) {
	return 0, nil
}

func (m *BannerModel) Get(id int) (*models.Banner, error) {
	return nil, nil
}
