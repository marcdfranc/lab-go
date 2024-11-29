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

type Keyvalue struct {
	Key   string
	Value string
}
