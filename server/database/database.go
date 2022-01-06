package database

import (
	"context"
	"database/sql"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
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

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}

func (db *DB) CreateMeasurement(ctx context.Context, input *model.NewMeasurement) *model.Measurement {
	//_, err := db.NewInsert().Model(&input).TableExpr("measurements").Exec()
	_, err := db.DB.NewInsert().Model(input).ModelTableExpr("measurements").Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Measurement{
		Hum:            input.Hum,
		Temp:           input.Temp,
		Moist:          input.Moist,
		WithIrrigation: input.WithIrrigation,
	}
}

func (db *DB) CreateSettings(ctx context.Context, input *model.NewSettings) *model.Setting {
	_, err := db.DB.NewInsert().Model(input).ModelTableExpr("settings").Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Setting{
		LimitsTrigger:      input.LimitsTrigger,
		WaterLevelLimit:    input.WaterLevelLimit,
		WaterAmountLimit:   input.WaterAmountLimit,
		MoistLimit:         input.MoistLimit,
		ScheduledTrigger:   input.ScheduledTrigger,
		HourRange:          input.HourRange,
		Location:           input.Location,
		IrrigationDuration: input.IrrigationDuration,
		ChartType:          input.ChartType,
		Language:           input.Language,
		Theme:              input.Theme,
		Lat:                input.Lat,
		Lon:                input.Lon,
	}
}

// TODO: DIS IS FOKEN RETARDED, WE NEED TO FIX THIS CRAP FUCK SHIT
func (db *DB) UpdateSettings(ctx context.Context, input *model.NewSettings) *model.Setting {
	values := db.DB.NewValues(input)
	_, err := db.DB.NewUpdate().With("_data", values).Model(input).TableExpr("_data").Bulk().Where("settings.id = 1").Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Setting{
		LimitsTrigger:      input.LimitsTrigger,
		WaterLevelLimit:    input.WaterLevelLimit,
		WaterAmountLimit:   input.WaterAmountLimit,
		MoistLimit:         input.MoistLimit,
		ScheduledTrigger:   input.ScheduledTrigger,
		HourRange:          input.HourRange,
		Location:           input.Location,
		IrrigationDuration: input.IrrigationDuration,
		ChartType:          input.ChartType,
		Language:           input.Language,
		Theme:              input.Theme,
		Lat:                input.Lat,
		Lon:                input.Lon,
	}
}

func (db *DB) GetMeasurements(ctx context.Context) []*model.Measurement {
	measurements := make([]*model.Measurement, 0)
	err := db.DB.NewSelect().Model(&measurements).Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return measurements
}

func (db *DB) GetSettings(ctx context.Context) []*model.Setting {
	settings := make([]*model.Setting, 0)
	err := db.DB.NewSelect().Model(&settings).Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return settings
}

func (db *DB) GetIrrigation(ctx context.Context) []*model.Irrigation {
	irrigationHistory := make([]*model.Irrigation, 0)
	err := db.DB.NewSelect().Model(&irrigationHistory).Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return irrigationHistory
}
