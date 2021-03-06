// @generated

package feature

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/validator"
)

// iCoreRepository ...
type iCoreRepository interface {
	getByID(ctx context.Context, params arguments.FeatureGetByIDArgs) (models.Feature, error)
	list(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error)
	count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error)
	insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error)
	update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error)
	delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error)
}

// GetByID ...
func (h HandlerImpl) GetByID(ctx context.Context, params arguments.FeatureGetByIDArgs) (models.Feature, error) {
	var (
		repository = models.Feature{}
	)
	repository, err := h.iRepository.getByID(ctx, params)
	if err != nil {
		return repository, err
	}
	return repository, nil
}

// Count ...
func (h HandlerImpl) Count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error) {
	var (
		count int64
	)
	if err := validator.Struct(params); err != nil {
		return count, err
	}
	count, err := h.iRepository.count(ctx, params)
	if err != nil {
		return count, err
	}
	return count, nil
}

// List ...
func (h HandlerImpl) List(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error) {
	var (
		repositories = []models.Feature{}
	)
	if err := validator.Struct(params); err != nil {
		return repositories, err
	}
	repositories, err := h.iRepository.list(ctx, params)
	if err != nil {
		return repositories, err
	}
	return repositories, nil
}

// Insert ...
func (h HandlerImpl) Insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error) {
	var (
		repository = models.Feature{}
	)
	if err := validator.Struct(params); err != nil {
		return repository, err
	}
	repository, err := h.iRepository.insert(ctx, params)
	if err != nil {
		return repository, err
	}
	return repository, nil
}

// Update ...
func (h HandlerImpl) Update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error) {
	var (
		repository = models.Feature{}
	)
	if err := validator.Struct(params); err != nil {
		return repository, err
	}
	repository, err := h.iRepository.update(ctx, params)
	if err != nil {
		return repository, err
	}
	return repository, nil
}

// Delete ...
func (h HandlerImpl) Delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error) {
	var (
		id int64
	)
	if err := validator.Struct(params); err != nil {
		return id, err
	}
	id, err := h.iRepository.delete(ctx, params)
	if err != nil {
		return id, err
	}
	return id, nil
}

//go:generate mockery -name=iRepository -output=mocks -case=underscore
