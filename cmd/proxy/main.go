package main

import (
	"accelerate/proxy"
	"io"
	"log"
	"net/http"
)

const backendURL = "http://localhost:8080"
const port = ":9090"

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Proxy: received request → forwarding to backend")

	key := r.Method + ":" +  r.URL.RequestURI()
	if r.Method == http.MethodGet{
		if cached, exists := proxy.MEMORY_CACHE[key]; exists {
			log.Println("Cache HIT: ", key)
			w.Write(cached.Body)
			return
		}
	}
	if r.Method != http.MethodGet{
		delete(proxy.MEMORY_CACHE, "GET:" + r.URL.RequestURI())
	}

	req, err := http.NewRequest(r.Method, backendURL+r.RequestURI, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Backend unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		http.Error(w, "Error reading backend response ", http.StatusInternalServerError )
		return
	}

	headersCopy := make(http.Header)

	for k, v := range resp.Header {
		copiedValues := make([]string, len(v))
		copy(copiedValues, v)
		headersCopy[k] = copiedValues
	}

	if r.Method == http.MethodGet{
		proxy.MEMORY_CACHE[key] = proxy.CacheEntry{
			Body: body,
			StatusCode: resp.StatusCode,
			Headers: headersCopy,
		}
	}

	for k, v := range resp.Header {
		for _, val := range v {
			w.Header().Add(k, val)
		}
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
	log.Println("Cache MISS: ", key)
	for i :=range proxy.MEMORY_CACHE{
		log.Println(i)
	}
	
}

func main() {
	log.Println("Proxy running on", port)

	http.HandleFunc("/", ProxyHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}