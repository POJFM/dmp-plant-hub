package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DB struct {
	DB *bun.DB
}

func Connect() *DB {
	// BUNDB

	conn := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN("postgres://postgres:@"+env.Process("DB_URL")+"/test?sslmode=disable"),
		pgdriver.WithUser(env.Process("DB_USER")),
		pgdriver.WithPassword(env.Process("DB_PSWD")),
		pgdriver.WithDatabase(env.Process("DB_NAME")),
	))

	db := bun.NewDB(conn, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	waitForDB(db)

	return &DB{
		db,
	}
}

func waitForDB(db *bun.DB) {
	var err error
	for i := 0; i < 120; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("Successfully connected to DB!")
			return
		}
		log.Printf("Failed to connect DB. Retrying in 10s. Number of retries: %v", i)
		time.Sleep(10 * time.Second)
	}
	log.Fatalf("DB CONN ERROR: %s", err)
}

func (db *DB) CreateMeasurement(ctx context.Context, input *model.NewMeasurement) *model.Measurement {
	_, err := db.DB.NewInsert().Model(input).ModelTableExpr("measurements").Exec(ctx)
	if err != nil {
		log.Println(err)
	}
	return &model.Measurement{
		Hum:            input.Hum,
		Temp:           input.Temp,
		Moist:          input.Moist,
		WithIrrigation: input.WithIrrigation,
	}
}

func (db *DB) GetMeasurements(ctx context.Context) []*model.Measurement {
	measurements := make([]*model.Measurement, 0)
	err := db.DB.NewSelect().Model(&measurements).Scan(ctx)
	if err != nil {
		log.Println(err)
	}
	return measurements
}

func (db *DB) CreateIrrigation(ctx context.Context, input *model.NewIrrigation) *model.IrrigationHistory {
	_, err := db.DB.NewInsert().Model(input).ModelTableExpr("irrigation_history").Exec(ctx)
	if err != nil {
		log.Println(err)
	}
	return &model.IrrigationHistory{
		WaterLevel:     input.WaterLevel,
		WaterAmount:    input.WaterAmount,
		WaterOverdrawn: input.WaterOverdrawn,
	}
}

func (db *DB) GetIrrigation(ctx context.Context) []*model.IrrigationHistory {
	irrigationHistory := make([]*model.IrrigationHistory, 0)
	err := db.DB.NewSelect().Model(&irrigationHistory).ModelTableExpr("irrigation_history").Scan(ctx)
	if err != nil {
		log.Println(err)
	}
	return irrigationHistory
}
