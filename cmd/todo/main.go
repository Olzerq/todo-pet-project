package todo

import "github.com/olzerq/todo-pet-project/internal/repository/postgres"

func main() {
	dbConfig := postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "task_app",
		Password: "Example",
		DBName:   "task_manager",
		SSLMode:  "disable",
	}
	db, err := postgres.NewPostgreSQL(dbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
