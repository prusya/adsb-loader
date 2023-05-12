package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func connectDb() {
	var err error
	db, err = sqlx.Connect("postgres", os.Getenv("PG_CONN_STRING"))
	if err != nil {
		log.Panicln("cannot connect to postgres", err)
	}
}

func migrateDb() {
	schema1 := `CREATE TABLE IF NOT EXISTS "adsblol" (
		data jsonb
	)`
	_, err := db.Exec(schema1)
	if err != nil {
		log.Panicln("cannot create table adsblol", err)
	}

	schema2 := `CREATE TABLE IF NOT EXISTS "adsbone" (
		data jsonb
	)`
	_, err = db.Exec(schema2)
	if err != nil {
		log.Panicln("cannot create table adsbone", err)
	}
}

func closeDb() {
	err := db.Close()
	if err != nil {
		log.Println("", err)
	}
}

var q1 = `INSERT INTO "adsblol" (data) VALUES ($1)`

func createAdsblolRecord(data []byte) {
	_, err := db.Exec(q1, data)
	if err != nil {
		log.Println("cannot create adsblol record:", err)
		log.Println(string(data))
	}
}

var q2 = `INSERT INTO "adsbone" (data) VALUES ($1)`

func createAdsboneRecord(data []byte) {
	_, err := db.Exec(q2, data)
	if err != nil {
		log.Println("cannot create adsbone record:", err)
		log.Println(string(data))
	}
}
