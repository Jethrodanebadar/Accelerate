package main

import (
	"accelerate/utils"
	"log"
	"net/http"
)

const port = ":8080"


func main() {
	utils.InitMockData()

	http.HandleFunc("/players", utils.PlayersHandler)
	http.HandleFunc("/players/", utils.PlayerByIDHandler)

	http.HandleFunc("/teams", utils.TeamsHandler)

	log.Println("Backend running on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}