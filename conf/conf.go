package conf

import "github.com/gorilla/sessions"

//configuration
var (
	//PathDB     = getUserPath() + "database/"
	Store = sessions.NewCookieStore([]byte("test"))

//StaticPath = "./static"
)

//configuration
const (
//NameDB = "test.db"
)
