// @generated

package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/utils"
)

// getByID ...
func (r userImpl) getByID(ctx context.Context, params arguments.UserGetByIDArgs) (models.User, error) {
	var (
		user          = models.User{}
		selectBuilder = sq.Select("*").From("user").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return user, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	row.Scan(
		&user.ID,
		&user.Email,
		&user.CompanyID,
		&user.Status,
		&user.CreatedBy,
		&user.UpdatedBy,
	)
	return user, nil
}

// setArgsToListSelectBuilder ...
func setArgsToListSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.UserListArgs) sq.SelectBuilder {
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Email != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"email": params.Email})
	}
	if params.CompanyID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"company_id": params.CompanyID})
	}
	if params.Status != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": params.Status})
	}
	if params.CreatedBy != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"created_by": params.CreatedBy})
	}
	if params.UpdatedBy != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"updated_by": params.UpdatedBy})
	}
	if params.PageSize != 0 {
		selectBuilder = selectBuilder.Limit(uint64(params.PageSize))
	}
	if params.Page != 0 {
		offset := utils.CalculateOffsetForPage(params.Page, params.PageSize)
		selectBuilder = selectBuilder.Offset(uint64(offset))
	}
	return selectBuilder
}

// list ...
func (r userImpl) list(ctx context.Context, params arguments.UserListArgs) ([]models.User, error) {
	var (
		users         = []models.User{}
		selectBuilder = sq.Select("*").From("user")
	)
	selectBuilderWithArgs := setArgsToListSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	if err != nil {
		return users, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return users, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.CompanyID,
			&user.Status,
			&user.CreatedBy,
			&user.UpdatedBy,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// setArgsToCountSelectBuilder ...
func setArgsToCountSelectBuilder(selectBuilder sq.SelectBuilder, params arguments.UserCountArgs) sq.SelectBuilder {
	if params.ID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"id": params.ID})
	}
	if params.Email != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"email": params.Email})
	}
	if params.CompanyID != 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"company_id": params.CompanyID})
	}
	if params.Status != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": params.Status})
	}
	if params.CreatedBy != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"created_by": params.CreatedBy})
	}
	if params.UpdatedBy != "" {
		selectBuilder = selectBuilder.Where(sq.Eq{"updated_by": params.UpdatedBy})
	}
	return selectBuilder
}

// count ...
func (r userImpl) count(ctx context.Context, params arguments.UserCountArgs) (int64, error) {
	var (
		count         int64
		selectBuilder = sq.Select("COUNT(id)").From("user")
	)
	selectBuilderWithArgs := setArgsToCountSelectBuilder(selectBuilder, params)
	sql, args, err := selectBuilderWithArgs.ToSql()
	if err != nil {
		return count, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return count, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// insert ...
func (r userImpl) insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error) {
	var (
		user          = models.User{}
		insertBuilder = sq.Insert("user").Columns(
			"email",
			"company_id",
			"status",
			"created_by",
			"updated_by",
		).Values(
			params.Email,
			params.CompanyID,
			params.Status,
			params.CreatedBy,
			params.UpdatedBy,
		)
	)
	sql, args, err := insertBuilder.ToSql()
	if err != nil {
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return user, err
	}
	row, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return user, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return user, err
	}
	user = models.User{
		ID: id, Email: params.Email,
		CompanyID: params.CompanyID,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	return user, nil
}

// setArgsToUpdateBuilder ...
func setArgsToUpdateBuilder(updateBuilder sq.UpdateBuilder, params arguments.UserUpdateArgs) sq.UpdateBuilder {
	if params.Email != "" {
		updateBuilder = updateBuilder.Set("email", params.Email)
	}
	if params.CompanyID != 0 {
		updateBuilder = updateBuilder.Set("company_id", params.CompanyID)
	}
	if params.Status != "" {
		updateBuilder = updateBuilder.Set("status", params.Status)
	}
	if params.CreatedBy != "" {
		updateBuilder = updateBuilder.Set("created_by", params.CreatedBy)
	}
	if params.UpdatedBy != "" {
		updateBuilder = updateBuilder.Set("updated_by", params.UpdatedBy)
	}

	return updateBuilder
}

// update ...
func (r userImpl) update(ctx context.Context, params arguments.UserUpdateArgs) (models.User, error) {
	var (
		user          = models.User{}
		updateBuilder = sq.Update("user").Where(sq.Eq{"id": params.ID})
	)
	updateBuilderWithArgs := setArgsToUpdateBuilder(updateBuilder, params)

	sql, args, err := updateBuilderWithArgs.ToSql()
	if err != nil {
		return user, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return user, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return user, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return user, err
	}
	if rowAffected == 0 {
		return user, fmt.Errorf("not found record by id %d", params.ID)
	}
	user = models.User{
		ID:        params.ID,
		Email:     params.Email,
		CompanyID: params.CompanyID,
		Status:    params.Status,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}
	return user, nil
}

// delete ...
func (r userImpl) delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error) {
	var (
		id            int64
		deleteBuilder = sq.Delete("user").Where(sq.Eq{"id": params.ID})
	)
	sql, args, err := deleteBuilder.ToSql()
	if err != nil {
		return id, err
	}
	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return id, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return id, err
	}
	rowAffected, err := result.RowsAffected()
	if rowAffected == 0 {
		return id, fmt.Errorf("not found record by id %d", params.ID)
	}
	return params.ID, nil
}
