package app

import "github.com/LebedevNd/BannerRotator/internal/models/database"

type App struct {
	*database.BannerModel
	*database.SlotModel
	*database.GroupModel
}
