package record

// UserRecord represents a user account in the system.
// Contains core user information and authentication details.
type UserRecord struct {
	ID              int     `json:"id"`               // Unique identifier for the user
	FirstName       string  `json:"firstname"`        // User's first name
	LastName        string  `json:"lastname"`         // User's last name
	Email           string  `json:"email"`            // User's email address (used for authentication)
	Password        string  `json:"password"`         // Hashed password (should never be exposed in responses)
	VerifiedCode    string  `json:"verified_code"`    // Verification code for email confirmation
	IsVerified      bool    `json:"is_verified"`      // Flag indicating if user has verified their email
	ConfirmPassword string  `json:"confirm_password"` // Used only during registration/update (not stored)
	CreatedAt       string  `json:"created_at"`       // Timestamp when user account was created
	UpdatedAt       string  `json:"updated_at"`       // Timestamp when user account was last updated
	DeletedAt       *string `json:"deleted_at"`       // Timestamp when account was deactivated (nil if active)
}
