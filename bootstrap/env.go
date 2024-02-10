package bootstrap

type Env struct {
	MongoURI           string
	DBName             string
	AccessTokenSecret  string
	AccessTokenExpiry  int
	RefreshTokenSecret string
	RefreshTokenExpiry int
}

func NewEnv() *Env {
	return &Env{
		MongoURI:           "mongodb://localhost:27017",
		DBName:             "golang-tz",
		AccessTokenSecret:  "golang-tz-access-token-secret",
		AccessTokenExpiry:  24,
		RefreshTokenSecret: "golang-tz-refresh-token-secret",
		RefreshTokenExpiry: 72,
	}
}
