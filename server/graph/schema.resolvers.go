package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/model"
)

func (r *mutationResolver) CreateMeasurement(ctx context.Context, input *model.NewMeasurement) (*model.Measurement, error) {
	return r.DB.CreateMeasurement(ctx, input), nil
}

func (r *mutationResolver) CreateSetting(ctx context.Context, input *model.NewSetting) (*model.Settings, error) {
	return r.DB.CreateSetting(ctx, input), nil
}

func (r *queryResolver) GetMeasurement(ctx context.Context, id string) (*model.Measurement, error) {
	return r.DB.GetMeasurement(ctx), nil
}

func (r *queryResolver) GetMeasurements(ctx context.Context) ([]*model.Measurement, error) {
	return r.DB.GetMeasurements(ctx), nil
}

func (r *queryResolver) GetSettings(ctx context.Context) ([]*model.Settings, error) {
	return r.DB.GetSettings(ctx), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
