module github.com/Philo-Ultra/poke-list-api

go 1.21.5

require github.com/gorilla/mux v1.8.1

require github.com/lib/pq v1.10.9

replace github.com/Philo-Ultra/poke-list-api/pkg/service => ./pkg/service

replace github.com/Philo-Ultra/poke-list-api/pkg/handlers => ./pkg/handlers

replace github.com/Philo-Ultra/poke-list-api/pkg/routes => ./pkg/routes

replace github.com/Philo-Ultra/poke-list-api/pkg/models => ./pkg/models
