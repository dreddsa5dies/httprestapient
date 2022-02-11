package main

import (
	"log"

	"github.com/dreddsa5dies/httprestapient/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5555 user=muser dbname=matrixdb password=12345678")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
