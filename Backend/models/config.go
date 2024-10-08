package models

type AccessConfig struct {
	Port                string
	AccessTokenSecret   string
	AccessTokenLifetime int
}

type RefreshConfig struct {
	Port                 string
	RefreshTokenSecret   string
	RefreshTokenLifetime int
}
