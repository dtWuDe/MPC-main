package errors

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code string, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// Auth Errors
var (
	ErrGoogleOauthFailed = NewAppError("GOOGLE_OAUTH_FAILED", "google oauth failed", 400)
	ErrUnauthorized      = NewAppError("UNAUTHORIZED", "unauthorized", 401)
	ErrInvalidPassword   = NewAppError("INVALID_PASSWORD", "invalid password", 401)
	ErrEmailAlreadyInUse = NewAppError("EMAIL_ALREADY_IN_USE", "email already in use", 409)
)

// User Errors
var (
	ErrUserNotFound   = NewAppError("USER_NOT_FOUND", "user not found", 404)
	ErrInvalidRequest = NewAppError("INVALID_REQUEST", "invalid request", 400)
	ErrWalletNotFound = NewAppError("WALLET_NOT_FOUND", "wallet not found", 404)
)

// Asset Errors
var (
	ErrChainNotFound        = NewAppError("CHAIN_NOT_FOUND", "chain not found", 404)
	ErrTokenNotFound        = NewAppError("TOKEN_NOT_FOUND", "token not found", 404)
	ErrInvalidChainID       = NewAppError("INVALID_CHAIN_ID", "invalid chain id", 400)
	ErrInvalidSymbol        = NewAppError("INVALID_SYMBOL", "invalid symbol", 400)
	ErrUnsupportedTokenType = NewAppError("UNSUPPORTED_TOKEN_TYPE", "unsupported token type", 400)
)

// Transaction Errors
var (
	ErrTransactionNotFound = NewAppError("TRANSACTION_NOT_FOUND", "transaction not found", 404)
	ErrTransactionFailed   = NewAppError("TRANSACTION_FAILED", "transaction failed", 400)
	ErrNotImplemented      = NewAppError("NOT_IMPLEMENTED", "not implemented", 501)
	ErrInvalidWallet       = NewAppError("INVALID_WALLET", "invalid wallet", 400)
	ErrInvalidAmount       = NewAppError("INVALID_AMOUNT", "invalid amount", 400)
	ErrInvalidAddress      = NewAppError("INVALID_ADDRESS", "invalid address", 400)
	ErrInssuficientBalance = NewAppError("INSUFFICIENT_BALANCE", "insufficient balance", 400)
)

// B2B Organization Errors
var (
	ErrOrganizationNotFound = NewAppError("ORGANIZATION_NOT_FOUND", "organization not found", 404)
	ErrOrganizationInactive = NewAppError("ORGANIZATION_INACTIVE", "organization is inactive", 403)
	ErrOrganizationLimitExceeded = NewAppError("ORGANIZATION_LIMIT_EXCEEDED", "organization limit exceeded", 429)
)

// API Key Errors
var (
	ErrAPIKeyNotFound = NewAppError("API_KEY_NOT_FOUND", "API key not found", 404)
	ErrInvalidAPIKey  = NewAppError("INVALID_API_KEY", "invalid API key", 401)
	ErrAPIKeyExpired  = NewAppError("API_KEY_EXPIRED", "API key has expired", 401)
	ErrAPIKeyInactive = NewAppError("API_KEY_INACTIVE", "API key is inactive", 401)
	ErrInsufficientPermissions = NewAppError("INSUFFICIENT_PERMISSIONS", "insufficient permissions", 403)
)

// Rate Limiting Errors
var (
	ErrRateLimitExceeded = NewAppError("RATE_LIMIT_EXCEEDED", "rate limit exceeded", 429)
	ErrQuotaExceeded     = NewAppError("QUOTA_EXCEEDED", "quota exceeded", 429)
)

// Member Errors
var (
	ErrMemberNotFound     = NewAppError("MEMBER_NOT_FOUND", "member not found", 404)
	ErrMemberAlreadyExists = NewAppError("MEMBER_ALREADY_EXISTS", "member already exists", 409)
	ErrInvalidRole        = NewAppError("INVALID_ROLE", "invalid role", 400)
	ErrMemberInactive     = NewAppError("MEMBER_INACTIVE", "member is inactive", 403)
)

// Usage Tracking Errors
var (
	ErrUsageNotFound = NewAppError("USAGE_NOT_FOUND", "usage data not found", 404)
)
