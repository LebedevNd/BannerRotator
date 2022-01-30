package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/LebedevNd/BannerRotator/internal/app"
	"github.com/LebedevNd/BannerRotator/internal/models/database"
	internalhttp "github.com/LebedevNd/BannerRotator/internal/server"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	// comment for lint.
	_ "github.com/jackc/pgx/stdlib"
)

const configFile string = "/configs/config.json"

func main() {
	config, err := NewConfig(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := connectToDb(config.Database, context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	bannerRotator := &app.App{
		BannerModel: &database.BannerModel{DB: db},
		SlotModel:   &database.SlotModel{DB: db},
		GroupModel:  &database.GroupModel{DB: db},
	}

	server := internalhttp.NewServer(*bannerRotator, config.Server.Host, config.Server.Port)
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			fmt.Println("failed to stop http server: " + err.Error())
		}
	}()

	fmt.Println("banner_rotator is up...")

	if err := server.Start(ctx); err != nil {
		fmt.Println("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

func connectToDb(database DatabaseConf, ctx context.Context) (*sql.DB, error) {
	dsn := "user=" + database.Username +
		" dbname=" + database.Database +
		" sslmode=disable password=" + database.Password +
		" host=" + database.Host +
		" port=" + strconv.Itoa(database.Port)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	return db, err
}
