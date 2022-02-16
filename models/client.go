package models

type OauthClients struct {
	ClientID     string //`gorm:"varchar(80);primary_key"`
	ClientSecret string //`gorm:"varchar(80);"`
	ClientIP     string //`gorm:"varchar(16);"`
	GrantType    string //`gorm:"varchar(80);"`
	Scope        string //`gorm:"varchar(1600);"`
}

// NewClient create to client model instance
func NewClient() *OauthClients {
	return &OauthClients{}
}

func (c *OauthClients) GetClientID() string {
	return c.ClientID
}

func (c *OauthClients) GetClientSecret() string {
	return c.ClientSecret
}

func (c *OauthClients) GetClientIP() string {
	return c.ClientIP
}

func (c *OauthClients) GetGrantType() string {
	return c.GrantType
}

func (c *OauthClients) GetScope() string {
	return c.Scope
}
