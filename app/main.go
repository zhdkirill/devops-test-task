package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

type Data struct {
	Title string
}

func getDb() *redis.Client {
	dbHost := getConfig("ERVCP_DB_HOST", "")
	dbPort := getConfig("ERVCP_DB_PORT", "")
	dbPw := getConfig("ERVCP_DB_PW", "")

	if dbHost == "" || dbPort == "" {
		panic("Database Connection not configured!")
	}

	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Password: dbPw,
		DB:       0,
	})

	return db
}

func getConfig(key string, def string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return def
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	count := 0
	db := getDb()
	val, _ := db.Get("count").Result()

	if val != "" {
		count, _ = strconv.Atoi(val)
	}

	count = count + 1
	db.Set("count", count, 0).Result()

	data := Data{Title: strconv.Itoa(count)}
	if count == 100 {
		data = Data{Title: fmt.Sprintf("Congratulations, You Are The %dth visitor to this site!", count)}
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	port := getConfig("ERVCP_PORT", "8080")

	fmt.Printf("ERVCP listening on port %s", port)

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
