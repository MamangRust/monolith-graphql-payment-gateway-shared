package userrecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

type userCommandRecordMapper struct{}

func NewUserCommandRecordMapper() UserCommandRecordMapper {
	return &userCommandRecordMapper{}
}

// ToUserRecord maps a User database row to a UserRecord domain model.
//
// Parameters:
//   - user: A pointer to a User representing the database row.
//
// Returns:
//   - A pointer to a UserRecord containing the mapped data, including
//     ID, FirstName, LastName, VerifiedCode, IsVerified, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userCommandRecordMapper) ToUserRecord(user *db.User) *record.UserRecord {
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
