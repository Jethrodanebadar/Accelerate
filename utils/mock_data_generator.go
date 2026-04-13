package utils

import (
	"math/rand"
	"strconv"
	"sync"
)

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Players = map[int]Player{}
var Teams = map[int]Team{}

var Mutex = &sync.RWMutex{}

func generateRandomAge() int {
	return rand.Intn(40) + 21
}

func InitMockData() {
	for i := 1; i <= 100; i++ {
		Players[i] = Player{
			ID:   i,
			Name: "player" + strconv.Itoa(i),
			Age:  generateRandomAge(),
		}
	}

	for i := 1; i <= 10; i++ {
		Teams[i] = Team{
			ID:   i,
			Name: "team" + strconv.Itoa(i),
		}
	}
}
