package constants

type sortProviderSignCtx string

const (
	OptionsContextKey sortProviderSignCtx = "provider_sign_options"

	Host = "http://localhost:8000"

	UserRoleID   = 1
	AdminRoleID  = 2
	SuperAdminID = 3

	JWTsecret       = "$3cr3t"
	AccessTokenTTL  = 120
	RefreshTokenTTL = 43200 // 30 days
)
