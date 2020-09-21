package sql

//***sql queries***

//GetUserSQL returns users SQL queries
func GetUserSQL() map[string]string {
	return map[string]string{
		//"InsertUser": "INSERT",
		"GetUser": "SELECT id FROM users WHERE name = $1 and password = $2",
	}
}
