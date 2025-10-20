package userrecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

type UserBaseRecordMapper interface {
	// ToUserRecord maps a User database row to a UserRecord domain model.
	//
	// Parameters:
	//   - user: A pointer to a User representing the database row.
	//
	// Returns:
	//   - A pointer to a UserRecord containing the mapped data, including
	//     ID, FirstName, LastName, VerifiedCode, IsVerified, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.
	ToUserRecord(user *db.User) *record.UserRecord
}

// UserQueryRecordMapper maps user query results to domain models.
type UserQueryRecordMapper interface {
	UserBaseRecordMapper

	// ToUserRecordPagination maps a GetUsersWithPaginationRow database row to a UserRecord domain model.
	//
	// Parameters:
	//   - user: A pointer to a GetUsersWithPaginationRow representing the database row.
	//
	// Returns:
	//   - A pointer to a UserRecord containing the mapped data, including
	//     ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
	ToUserRecordPagination(user *db.GetUsersWithPaginationRow) *record.UserRecord

	// ToUsersRecordPagination maps a slice of GetUsersWithPaginationRow database rows to a slice of UserRecord domain models.
	//
	// Parameters:
	//   - users: A slice of pointers to GetUsersWithPaginationRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to UserRecord containing the mapped data for each user, including
	//     ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
	ToUsersRecordPagination(users []*db.GetUsersWithPaginationRow) []*record.UserRecord

	// ToUserRecordActivePagination maps a GetActiveUsersWithPaginationRow database row to a UserRecord domain model.
	//
	// Parameters:
	//   - user: A pointer to a GetActiveUsersWithPaginationRow representing the database row.
	//
	// Returns:
	//   - A pointer to a UserRecord containing the mapped data, including
	//     ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
	ToUserRecordActivePagination(user *db.GetActiveUsersWithPaginationRow) *record.UserRecord

	// ToUsersRecordActivePagination maps a slice of GetActiveUsersWithPaginationRow database rows to a slice of UserRecord domain models.
	//
	// Parameters:
	//   - users: A slice of pointers to GetActiveUsersWithPaginationRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to UserRecord containing the mapped data for each user, including
	//     ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
	ToUsersRecordActivePagination(users []*db.GetActiveUsersWithPaginationRow) []*record.UserRecord

	// ToUserRecordTrashedPagination maps a GetTrashedUsersWithPaginationRow database row to a UserRecord domain model.
	//
	// Parameters:
	//   - user: A pointer to a GetTrashedUsersWithPaginationRow representing the database row.
	//
	// Returns:
	//   - A pointer to a UserRecord containing the mapped data, including
	//     ID, FirstName, LastName, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.
	ToUserRecordTrashedPagination(user *db.GetTrashedUsersWithPaginationRow) *record.UserRecord

	// ToUsersRecordTrashedPagination maps a slice of GetTrashedUsersWithPaginationRow database rows
	// to a slice of UserRecord domain models.
	//
	// Parameters:
	//   - users: A slice of pointers to GetTrashedUsersWithPaginationRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to UserRecord containing the mapped data for each user, including
	//     ID, FirstName, LastName, Email, Password, CreatedAt, UpdatedAt, and DeletedAt. domain models.
	ToUsersRecordTrashedPagination(users []*db.GetTrashedUsersWithPaginationRow) []*record.UserRecord
}

// UserCommandRecordMapper maps user table rows to domain models for command operations.
type UserCommandRecordMapper interface {
	UserBaseRecordMapper
}
