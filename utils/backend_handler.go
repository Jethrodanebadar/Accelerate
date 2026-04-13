package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		Mutex.RLock()
		defer Mutex.RUnlock()
		time.Sleep(2 * time.Second)
		json.NewEncoder(w).Encode(Players)

	case http.MethodPost:
		var p Player
		json.NewDecoder(r.Body).Decode(&p)

		Mutex.Lock()
		defer Mutex.Unlock()

		p.ID = len(Players) + 1
		Players[p.ID] = p

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PlayerByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/players/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		Mutex.RLock()
		defer Mutex.RUnlock()
		time.Sleep(2 * time.Second)
		player, ok := Players[id]
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(player)

	case http.MethodPut:
		var updated Player
		json.NewDecoder(r.Body).Decode(&updated)

		Mutex.Lock()
		defer Mutex.Unlock()

		updated.ID = id
		Players[id] = updated

		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		Mutex.Lock()
		defer Mutex.Unlock()

		delete(Players, id)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		Mutex.RLock()
		defer Mutex.RUnlock()
		time.Sleep(2 * time.Second)
		json.NewEncoder(w).Encode(Teams)

	case http.MethodPost:
		var t Team
		json.NewDecoder(r.Body).Decode(&t)

		Mutex.Lock()
		defer Mutex.Unlock()

		t.ID = len(Teams) + 1
		Teams[t.ID] = t

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}