// @generated

package arguments

// CompanyGetByIDArgs ...
type CompanyGetByIDArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}

// CompanyCountArgs ...
type CompanyCountArgs struct {
	ID        int64  `graphql:"id" validate:"required,min=1"`
	Name      string `graphql:"name" validate:"omitempty"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
}

// CompanyListArgs ...
type CompanyListArgs struct {
	ID        int64  `graphql:"id" validate:"omitempty,min=1"`
	Name      string `graphql:"name" validate:"omitempty"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
	Page      int64  `graphql:"page" validate:"required,min=1"`
	PageSize  int64  `graphql:"pageSize" validate:"required,min=1,max=40"`
}

// CompanyInsertArgs ...
type CompanyInsertArgs struct {
	Name      string `graphql:"name" validate:"omitempty"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
}

// CompanyUpdateArgs ...
type CompanyUpdateArgs struct {
	ID        int64  `graphql:"id" validate:"required,min=1"`
	Name      string `graphql:"name" validate:"omitempty"`
	Status    string `graphql:"status" validate:"omitempty,oneof=active inactive"`
	CreatedBy string `graphql:"createdBy" validate:"omitempty"`
	UpdatedBy string `graphql:"updatedBy" validate:"omitempty"`
}

// CompanyDeleteArgs ...
type CompanyDeleteArgs struct {
	ID int64 `graphql:"id" validate:"required,min=1"`
}
