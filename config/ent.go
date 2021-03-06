package config

import (
	"log"

	"github.com/dreddsa5dies/httprestapient/ent"

	_ "github.com/lib/pq"
)

var (
	client *ent.Client
)

// GetClient - getter for client
func GetClient() *ent.Client {
	return client
}

// SetClient - setter for client
func SetClient(newClient *ent.Client) {
	client = newClient
}

// NewEntClient - connect to postgresql
func NewEntClient() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=muser dbname=matrix_attack password=12345678 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}
