package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("PG_CONN_STRING")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	schema := `CREATE TABLE IF NOT EXISTS "adsblol" (
		data jsonb
	)`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalln(err)
	}

	ticker := time.NewTicker(10 * time.Second)

	q := `INSERT INTO "adsblol" (data) VALUES ($1)`

	for {
		select {
		case <-ticker.C:
			resp, err := http.Get("https://api.adsb.lol/v2/point/34.05/-118.24/250")
			if err != nil {
				fmt.Println("cannot get data from api.adsb.lol", err)
				continue
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("cannot read response body", err)
			} else {
				_, err = db.Exec(q, body)
				if err != nil {
					fmt.Println("cannot store body", err)
				}
			}

			resp.Body.Close()
		}
	}
}
