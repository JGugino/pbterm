package pb

import "fmt"

type CollectionType int

const (
	BaseCollection CollectionType = iota
	AuthCollection
	ViewCollection
)

func (c CollectionType) String() string {
	switch c {
	case 0:
		return "base"
	case 1:
		return "auth"
	case 2:
		return "view"
	default:
		return "base"
	}
}

type CollectionAuthAlerts struct {
	Enabled       bool                    `json:"enabled"`
	EmailTemplate CollectionEmailTemplate `json:"emailTemplate"`
}

type OAuthProviderOptions struct {
	Name         string `json:"name"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	AuthURL      string `json:"authURL"`
	TokenURL     string `json:"tokenURL"`
	UserInfoURL  string `json:"userInfoURL"`
	DisplayName  string `json:"displayName"`
	PKCE         string `json:"pkce"`
	Extra        string `json:"extra"`
}

type CollectionOAuthOptions struct {
	Enabled      bool `json:"enabled"`
	MappedFields struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatarURL"`
	} `json:"mappedFields"`
	Providers OAuthProviderOptions
}

type CollectionTokenOptions struct {
	Duration int    `json:"duration"`
	Secret   string `json:""`
}

type CollectionEmailTemplate struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type CollectionOptions struct {
	Name       string               `json:"name"`
	Type       CollectionType       `json:"type"`
	Fields     []map[string]any     `json:"fields"`
	System     bool                 `json:"system"`
	ListRule   string               `json:"listRule"`
	ViewRule   string               `json:"viewRule"`
	CreateRule string               `json:"createRule"`
	UpdateRule string               `json:"updateRule"`
	DeleteRule string               `json:"deleteRule"`
	ViewQuery  string               `json:"viewQuery"`
	AuthAlerts CollectionAuthAlerts `json:"authAlerts"`

	//Auth Options
	OAuth2       CollectionOAuthOptions `json:"oauth2"`
	PasswordAuth struct {
		Enabled        bool     `json:"enabled"`
		IdentityFields []string `json:"identityFields"`
	} `json:"passwordAuth"`
	MFA struct {
		Enabled  bool `json:"enabled"`
		Duration int  `json:"duration"`
		Rule     bool `json:"rule"`
	} `json:"mfa"`
	OTP struct {
		Enabled       bool                    `json:"enabled"`
		Duration      int                     `json:"duration"`
		Length        int                     `json:"length"`
		EmailTemplate CollectionEmailTemplate `json:"emailTemplate"`
	} `json:"otp"`

	//Auth Tokens
	AuthToken          CollectionTokenOptions `json:"authToken"`
	PasswordResetToken CollectionTokenOptions `json:"passwordResetToken"`
	EmailChangeToken   CollectionTokenOptions `json:"emailChangeToken"`
	VerificationToken  CollectionTokenOptions `json:"verificationToken"`
	FileToken          CollectionTokenOptions `json:"fileToken"`

	//Email Templates
	VerificationTemplate       CollectionEmailTemplate `json:"verificationTemplate"`
	ResetPasswordTemplate      CollectionEmailTemplate `json:"resetPasswordTemplate"`
	ConfirmEmailChangeTemplate CollectionEmailTemplate `json:"confirmEmailChangeTemplate"`
}

type PBCollection struct {
	BaseURL string
}

// ### CREATE COLLECTION ###
func (collection *PBCollection) CreateNewCollection(token string, options CollectionOptions) {
	apiUrl := fmt.Sprintf("%s/api/collections", collection.BaseURL)

	res, err := SendAuthenticatedHTTPRequest("POST", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)
}

// ### UPDATE COLLECTION ###
func (collection *PBCollection) UpdateCollection(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/%s", collection.BaseURL, desiredCollection)
	res, err := SendAuthenticatedHTTPRequest("PATCH", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)

}

// ### VIEW COLLECTION ###

func (collection *PBCollection) ScaffoldCollections(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/meta/scaffolds", collection.BaseURL)
	res, err := SendAuthenticatedHTTPRequest("GET", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)

}

func (collection *PBCollection) ViewCollection(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/%s", collection.BaseURL, desiredCollection)
	res, err := SendAuthenticatedHTTPRequest("GET", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)

}

func (collection *PBCollection) ListCollections(token string) {
	apiUrl := fmt.Sprintf("%s/api/collections", collection.BaseURL)
	res, err := SendAuthenticatedHTTPRequest("GET", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)

}

// ### DELETE COLLECTION ###
func (collection *PBCollection) DeleteCollection(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/%s", collection.BaseURL, desiredCollection)
	res, err := SendAuthenticatedHTTPRequest("DELETE", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)
}

func (collection *PBCollection) TruncateCollection(token string, desiredCollection string) {
	apiUrl := fmt.Sprintf("%s/api/collections/%s/truncate", collection.BaseURL, desiredCollection)
	res, err := SendAuthenticatedHTTPRequest("DELETE", apiUrl, map[string]string{}, map[string]any{}, token)

	if err != nil {
		return
	}

	fmt.Println(res)
}
