package database

import (
	"context"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/model"
	"log"
)

func (db *DB) CreateSettings(ctx context.Context, input *model.NewSettings) *model.Setting {
	*input.ID = 0
	_, err := db.DB.NewInsert().Model(input).ModelTableExpr("settings").Exec(ctx)
	if err != nil {
		log.Println(err)
	}
	return &model.Setting{
		ID:                 input.ID,
		LimitsTrigger:      input.LimitsTrigger,
		WaterLevelLimit:    input.WaterLevelLimit,
		DefaultWaterAmount: input.DefaultWaterAmount,
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
		WaterAmountLimit:   input.WaterAmountLimit,
	}
}

func (db *DB) GetSettings(ctx context.Context) []*model.Setting {
	settings := make([]*model.Setting, 0)
	err := db.DB.NewSelect().Model(&settings).Scan(ctx)
	if err != nil {
		log.Println(err)
	}
	return settings
}

func (db *DB) UpdateSettings(ctx context.Context, input *model.NewSettings) *model.Setting {
	// values := db.DB.NewValues(input)
	// modl := make([]*model.Setting, 0)
	wellthisiskindadumb := 0
	settings := model.Setting{
		ID:                 &wellthisiskindadumb,
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
		DefaultWaterAmount: input.DefaultWaterAmount,
	}
	_, err := db.DB.NewUpdate().Model(&settings).Where("id = ?", 0).Exec(ctx)
	if err != nil {
		log.Println(err)
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
		DefaultWaterAmount: input.DefaultWaterAmount,
	}
}

// GetSettingByColumn
// usage:
//	var columns = []string{"limits_trigger", "water_level_limit", "water_amount_limit", "moist_limit", "scheduled_trigger"}
//	data := db.GetSettingByColumn(columns)
//	fmt.Println(*data.LimitsTrigger)
func (db *DB) GetSettingByColumn(columns []string) model.Setting {
	var settings []model.Setting
	err := db.DB.NewSelect().Model(&settings).Column(columns...).Limit(1).Scan(context.Background())
	if err != nil {
		log.Println(err)
	}
	settingsRow := settings[0]
	return model.Setting{
		ID:                 settingsRow.ID,
		LimitsTrigger:      settingsRow.LimitsTrigger,
		WaterLevelLimit:    settingsRow.WaterLevelLimit,
		WaterAmountLimit:   settingsRow.WaterAmountLimit,
		MoistLimit:         settingsRow.MoistLimit,
		ScheduledTrigger:   settingsRow.ScheduledTrigger,
		HourRange:          settingsRow.HourRange,
		Location:           settingsRow.Location,
		IrrigationDuration: settingsRow.IrrigationDuration,
		ChartType:          settingsRow.ChartType,
		Language:           settingsRow.Language,
		Theme:              settingsRow.Theme,
		Lat:                settingsRow.Lat,
		Lon:                settingsRow.Lon,
		DefaultWaterAmount: settingsRow.DefaultWaterAmount,
	}
}

// CheckSettings
// checks if settings are already present
func (db *DB) CheckSettings() (isSettingsPresent bool) {
	var s []model.Setting
	err := db.DB.NewSelect().Model(&s).Limit(1).Scan(context.Background())
	if err != nil {
		log.Println(err)
	}
	isSettingsPresent = true
	if s == nil {
		isSettingsPresent = false
	}
	return
}
