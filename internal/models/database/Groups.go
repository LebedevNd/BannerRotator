package database

import (
	"database/sql"
	"github.com/LebedevNd/BannerRotator/internal/models"
)

type GroupModel struct {
	DB *sql.DB
}

func (m *GroupModel) Insert(description string) (int, error) {
	return 0, nil
}

func (m *GroupModel) Get(id int) (*models.Group, error) {
	return nil, nil
}
