package userrecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// userQueryRecordMapper is a struct that implements the UserQueryRecordMapper interface.
type userQueryRecordMapper struct{}

// NewUserQueryRecordMapper creates a new instance of userQueryRecordMapper, which implements the
// UserQueryRecordMapper interface. This mapper is responsible for mapping user query results
// from the database to domain models, facilitating data transformation and access in the application.
func NewUserQueryRecordMapper() UserQueryRecordMapper {
	return &userQueryRecordMapper{}
}

// ToUserRecord maps a User database row to a UserRecord domain model.
//
// Parameters:
//   - user: A pointer to a User representing the database row.
//
// Returns:
//   - A pointer to a UserRecord containing the mapped data, including
//     ID, FirstName, LastName, VerifiedCode, IsVerified, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userQueryRecordMapper) ToUserRecord(user *db.User) *record.UserRecord {
	var deletedAt *string

	if user.DeletedAt.Valid {
		formatedDeletedAt := user.DeletedAt.Time.Format("2006-01-02")

		deletedAt = &formatedDeletedAt
	}

	if user.VerificationCode == "" {
		user.VerificationCode = "null"
	}

	return &record.UserRecord{
		ID:           int(user.UserID),
		FirstName:    user.Firstname,
		LastName:     user.Lastname,
		VerifiedCode: user.VerificationCode,
		IsVerified:   user.IsVerified.Valid,
		Email:        user.Email,
		Password:     user.Password,
		CreatedAt:    user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:    user.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:    deletedAt,
	}
}

// ToUserRecordPagination maps a GetUsersWithPaginationRow database row to a UserRecord domain model.
//
// Parameters:
//   - user: A pointer to a GetUsersWithPaginationRow representing the database row.
//
// Returns:
//   - A pointer to a UserRecord containing the mapped data, including
//     ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userQueryRecordMapper) ToUserRecordPagination(user *db.GetUsersWithPaginationRow) *record.UserRecord {
	var deletedAt *string

	if user.DeletedAt.Valid {
		formatedDeletedAt := user.DeletedAt.Time.Format("2006-01-02")

		deletedAt = &formatedDeletedAt
	}

	return &record.UserRecord{
		ID:         int(user.UserID),
		FirstName:  user.Firstname,
		LastName:   user.Lastname,
		Email:      user.Email,
		Password:   user.Password,
		IsVerified: user.IsVerified.Valid,
		CreatedAt:  user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:  user.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:  deletedAt,
	}
}

// ToUsersRecordPagination maps a slice of GetUsersWithPaginationRow database rows to a slice of UserRecord domain models.
//
// Parameters:
//   - users: A slice of pointers to GetUsersWithPaginationRow representing the database rows.
//
// Returns:
//   - A slice of pointers to UserRecord containing the mapped data for each user, including
//     ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userQueryRecordMapper) ToUsersRecordPagination(users []*db.GetUsersWithPaginationRow) []*record.UserRecord {
	var userRecords []*record.UserRecord

	for _, user := range users {
		userRecords = append(userRecords, s.ToUserRecordPagination(user))
	}

	return userRecords
}

// ToUserRecordActivePagination maps a GetActiveUsersWithPaginationRow database row to a UserRecord domain model.
//
// Parameters:
//   - user: A pointer to a GetActiveUsersWithPaginationRow representing the database row.
//
// Returns:
//   - A pointer to a UserRecord containing the mapped data, including
//     ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userQueryRecordMapper) ToUserRecordActivePagination(user *db.GetActiveUsersWithPaginationRow) *record.UserRecord {
	var deletedAt *string

	if user.DeletedAt.Valid {
		formatedDeletedAt := user.DeletedAt.Time.Format("2006-01-02")

		deletedAt = &formatedDeletedAt
	}

	return &record.UserRecord{
		ID:         int(user.UserID),
		FirstName:  user.Firstname,
		LastName:   user.Lastname,
		Email:      user.Email,
		Password:   user.Password,
		IsVerified: user.IsVerified.Valid,
		CreatedAt:  user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:  user.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:  deletedAt,
	}
}

// ToUsersRecordActivePagination maps a slice of GetActiveUsersWithPaginationRow database rows to a slice of UserRecord domain models.
//
// Parameters:
//   - users: A slice of pointers to GetActiveUsersWithPaginationRow representing the database rows.
//
// Returns:
//   - A slice of pointers to UserRecord containing the mapped data for each user, including
//     ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userQueryRecordMapper) ToUsersRecordActivePagination(users []*db.GetActiveUsersWithPaginationRow) []*record.UserRecord {
	var userRecords []*record.UserRecord

	for _, user := range users {
		userRecords = append(userRecords, s.ToUserRecordActivePagination(user))
	}

	return userRecords
}

// ToUserRecordTrashedPagination maps a GetTrashedUsersWithPaginationRow database row to a UserRecord domain model.
//
// Parameters:
//   - user: A pointer to a GetTrashedUsersWithPaginationRow representing the database row.
//
// Returns:
//   - A pointer to a UserRecord containing the mapped data, including
//     ID, FirstName, LastName, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userQueryRecordMapper) ToUserRecordTrashedPagination(user *db.GetTrashedUsersWithPaginationRow) *record.UserRecord {
	var deletedAt *string

	if user.DeletedAt.Valid {
		formatedDeletedAt := user.DeletedAt.Time.Format("2006-01-02")

		deletedAt = &formatedDeletedAt
	}

	return &record.UserRecord{
		ID:        int(user.UserID),
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

// ToUsersRecordTrashedPagination maps a slice of GetTrashedUsersWithPaginationRow database rows
// to a slice of UserRecord domain models.
//
// Parameters:
//   - users: A slice of pointers to GetTrashedUsersWithPaginationRow representing the database rows.
//
// Returns:
//   - A slice of pointers to UserRecord containing the mapped data for each user, including
//     ID, FirstName, LastName, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userQueryRecordMapper) ToUsersRecordTrashedPagination(users []*db.GetTrashedUsersWithPaginationRow) []*record.UserRecord {
	var userRecords []*record.UserRecord

	for _, user := range users {
		userRecords = append(userRecords, s.ToUserRecordTrashedPagination(user))
	}

	return userRecords
}
