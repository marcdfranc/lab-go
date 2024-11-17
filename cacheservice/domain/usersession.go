package domain

type Preference struct {
	Language string
	Theme    string
}

type UserSession struct {
	Id           string
	Username     string
	Email        string
	Preferences  Preference
	Expiration   string
	RequestCount int
}
