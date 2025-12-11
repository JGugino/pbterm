package pb

type AuthMethodResponse struct {
	Password PwdAuthMethod    `json:"password"`
	OAuth2   OAuth2AuthMethod `json:"oauth2"`
	MFA      MFAAuthMethod    `json:"mfa"`
	OTP      OTPAuthMethod    `json:"otp"`
}

type PwdAuthMethod struct {
	Enabled        bool     `json:"enabled"`
	IdentityFields []string `json:"identityFields"`
}

type OAuth2AuthMethod struct {
	Enabled   bool            `json:"enabled"`
	Providers []OAuthProvider `json:"providers"`
}

type MFAAuthMethod struct {
	Enabled  bool `json:"enabled"`
	Duration int  `json:"duration"`
}

type OTPAuthMethod struct {
	Enabled  bool `json:"enabled"`
	Duration int  `json:"duration"`
}

type OAuthProvider struct {
	Name                string `json:"name"`
	DisplayName         string `json:"displayName"`
	State               string `json:"state"`
	AuthURL             string `json:"authURL"`
	CodeVerifier        string `json:"codeVerifier"`
	CodeChallenge       string `json:"codeChallenge"`
	CodeChallengeMethod string `json:"codeChallengeMethod"`
}

type AuthSuccessResponse struct {
	Token  string         `json:"token"`
	Record map[string]any `json:"record"`
}

type PocketBaseErrorResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

type PocketBaseListResponse struct {
	Page       int              `json:"page"`
	PerPage    int              `json:"perPage"`
	TotalItems int              `json:"totalItems"`
	TotalPages int              `json:"totalPages"`
	Items      []map[string]any `json:"items"`
}

type PocketBaseListOptions struct {
	Page      int
	PerPage   int
	Sort      string
	Filter    string
	Expand    string
	Fields    string
	SkipTotal bool
}
