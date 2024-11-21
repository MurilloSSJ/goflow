package database

func GetDatabase(driver string, host string, port int, username string, password string, database string) Database {
	switch driver {
		case "postgres":
			return &Postgres{
				Host:     "localhost",
				Port:     "5432",
				Username: "postgres",
				Password: "password",
				Database: "postgres",
			}
		default:
			return nil
	}
}