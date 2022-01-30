package database

import (
	"database/sql"
	"github.com/LebedevNd/BannerRotator/internal/models"
)

type SlotModel struct {
	DB *sql.DB
}

func (m *SlotModel) Insert(description string) (int, error) {
	return 0, nil
}

func (m *SlotModel) Get(id int) (*models.Slot, error) {
	return nil, nil
}
