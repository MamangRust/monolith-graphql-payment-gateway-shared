package record

// RefreshTokenRecord represents a refresh token used for maintaining user sessions.
// Refresh tokens allow users to obtain new access tokens without re-authenticating.
type RefreshTokenRecord struct {
	ID        int    `json:"id"`         // Unique identifier for the refresh token record
	UserID    int    `json:"user_id"`    // ID of the user this token belongs to
	Token     string `json:"token"`      // The actual refresh token value (typically hashed in storage)
	ExpiredAt string `json:"expired_at"` // Timestamp when this token expires and becomes invalid
	CreatedAt string `json:"created_at"` // Timestamp when the token was created
	UpdatedAt string `json:"updated_at"` // Timestamp when the token was last updated
}
