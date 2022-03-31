package database

import (
	"context"
	"database/sql"
	"log"

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
		pgdriver.WithDSN("postgres://postgres:@127.0.0.1:5420/test?sslmode=disable"),
		pgdriver.WithUser(env.Process("DB_USER")),
		pgdriver.WithPassword(env.Process("DB_PSWD")),
		pgdriver.WithDatabase(env.Process("DB_NAME")),
	))

	db := bun.NewDB(conn, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	if err := db.Ping(); err != nil {
		log.Fatalf("DB CONN ERROR: %s", err)
	}

	return &DB{
		db,
	}
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

/* TODO
func (db *DB) GetMeasurementsAverage(ctx context.Context) (measurementsAvg *sens.Measurements) {
	measurements := make([]*model.Measurement, 0)
	err := db.DB.NewSelect().Model(&measurements).Limit(50).Scan(context.Background())
	if err != nil {
		log.Println(err)
	}
	measurementsAvg.Moist = MeasurementsAvg()
	return
}
*/

func (db *DB) GetIrrigation(ctx context.Context) []*model.IrrigationHistory {
	irrigationHistory := make([]*model.IrrigationHistory, 0)
	err := db.DB.NewSelect().Model(&irrigationHistory).ModelTableExpr("irrigation_history").Scan(ctx)
	if err != nil {
		log.Println(err)
	}
	return irrigationHistory
}
