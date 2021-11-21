package database

import (
	"context"
	"database/sql"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
)

type DB struct {
	DB *bun.DB
}

func Connect() *DB {
	conn := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN("postgres://postgres:@localhost:5420/test?sslmode=disable"),
		pgdriver.WithUser("root"),
		pgdriver.WithPassword("k0k0s"),
		pgdriver.WithDatabase("planthub"),
	))

	db := bun.NewDB(conn, pgdialect.New())

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}

func (db *DB) Save(input *model.NewMeasurement, ctx context.Context) *model.Measurement {
	//_, err := db.NewInsert().Model(&input).TableExpr("measurements").Exec()
	_, err := db.DB.NewInsert().Model(input).ModelTableExpr("measurements").Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Measurement{
		Moisture:       input.Moisture,
		Temperature:    input.Temperature,
		Humidity:       input.Humidity,
		WaterLevel:     input.WaterLevel,
		WaterOverdrawn: input.WaterOverdrawn,
	}
}
