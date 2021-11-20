package graph

import "github.com/SPSOAFM-IT18/dmp-plant-hub/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *database.DB
}
