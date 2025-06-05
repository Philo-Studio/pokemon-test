package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Philo-Ultra/poke-list-api/pkg/handlers"
	"github.com/Philo-Ultra/poke-list-api/pkg/routes"
	"github.com/Philo-Ultra/poke-list-api/pkg/service"

	"github.com/gorilla/mux"
)

var (
	fPort       = flag.String("port", os.Getenv("PORT"), "the port to listen on, if not specified looks for env variable PORT")
	fDBName     = flag.String("dbname", os.Getenv("DB_NAME"), "the name of the database to connect to, if not specified looks for env variable DB_NAME")
	fDBHost     = flag.String("dbhost", os.Getenv("DB_HOST"), "the hostname of the database to connect to, if not specified looks for env variable DB_HOST")
	fDBPort     = flag.String("dbport", os.Getenv("DB_PORT"), "the port of the database to connect to, if not specified looks for env variable DB_PORT")
	fDBUsername = flag.String("dbusername", os.Getenv("DB_USER"), "the user for the database, if not specified looks for env variable DB_USER")
	fDBPassword = flag.String("dbpassword", os.Getenv("DB_PASS"), "the password for the database, if not specified looks for env variable DB_PASS")
)

// DBConnectionString constructs a PostgreSQL connection string from the given parameters
func DBConnectionString(user, pass, name, host, port string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, pass, name, host, port)
}

func main() {
	flag.Parse()

	if *fDBName == "" || *fDBHost == "" || *fDBPort == "" || *fDBUsername == "" || *fDBPassword == "" || *fPort == "" {
		fmt.Println("error: missing required environment variables or flags")
		fmt.Println()
		fmt.Println("DB_HOST", *fDBHost)
		fmt.Println("DB_PORT", *fDBPort)
		fmt.Println("DB_USERNAME", *fDBUsername)
		fmt.Println("DB_PASSWORD", *fDBPassword)
		fmt.Println("PORT", *fPort)
		flag.Usage()
		os.Exit(1)
	}

	connStr := DBConnectionString(*fDBUsername, *fDBPassword, *fDBName, *fDBHost, *fDBPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("error: could not connect to postgres database: (%v)\n", err)
		os.Exit(1)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("error: could not ping postgres database: (%v)\n", err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	svc := service.NewService(db)
	handler := handlers.NewHandler(*svc)

	routes.AddRoutes(r, handler)

	http.ListenAndServe(":"+*fPort, r)
}
