package record

// ResetTokenRecord represents a password reset token used for user authentication recovery.
// These tokens are typically short-lived and sent to users via email for password reset flows.
type ResetTokenRecord struct {
	ID        int64  `json:"id"`         // Unique identifier for the reset token record
	Token     string `json:"token"`      // The actual reset token value (typically hashed in storage)
	UserID    int64  `json:"user_id"`    // ID of the user this token belongs to
	ExpiredAt string `json:"expired_at"` // Timestamp when this token expires and becomes invalid
}
