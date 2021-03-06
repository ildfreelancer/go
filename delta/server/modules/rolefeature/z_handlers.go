// @generated

package rolefeature

import (
	"context"

	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/validator"
)

// iCoreRepository ...
type iCoreRepository interface {
	getByID(ctx context.Context, params arguments.RoleFeatureGetByIDArgs) (models.RoleFeature, error)
	list(ctx context.Context, params arguments.RoleFeatureListArgs) ([]models.RoleFeature, error)
	count(ctx context.Context, params arguments.RoleFeatureCountArgs) (int64, error)
	insert(ctx context.Context, params arguments.RoleFeatureInsertArgs) (models.RoleFeature, error)
	update(ctx context.Context, params arguments.RoleFeatureUpdateArgs) (models.RoleFeature, error)
	delete(ctx context.Context, params arguments.RoleFeatureDeleteArgs) (int64, error)
}

// GetByID ...
func (h HandlerImpl) GetByID(ctx context.Context, params arguments.RoleFeatureGetByIDArgs) (models.RoleFeature, error) {
	var (
		repository = models.RoleFeature{}
	)
	repository, err := h.iRepository.getByID(ctx, params)
	if err != nil {
		return repository, err
	}
	return repository, nil
}

// Count ...
func (h HandlerImpl) Count(ctx context.Context, params arguments.RoleFeatureCountArgs) (int64, error) {
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
func (h HandlerImpl) List(ctx context.Context, params arguments.RoleFeatureListArgs) ([]models.RoleFeature, error) {
	var (
		repositories = []models.RoleFeature{}
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
func (h HandlerImpl) Insert(ctx context.Context, params arguments.RoleFeatureInsertArgs) (models.RoleFeature, error) {
	var (
		repository = models.RoleFeature{}
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
func (h HandlerImpl) Update(ctx context.Context, params arguments.RoleFeatureUpdateArgs) (models.RoleFeature, error) {
	var (
		repository = models.RoleFeature{}
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
func (h HandlerImpl) Delete(ctx context.Context, params arguments.RoleFeatureDeleteArgs) (int64, error) {
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
