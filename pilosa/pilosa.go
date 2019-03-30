package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/pilosa/go-pilosa"
)

const restaurants = 65536

func fill(r *rand.Rand, field *pilosa.Field, probability float32) {
	for i := 0; i < restaurants; i++ {
		if r.Float32() < probability {
			// Send a Set query. If err is non-nil, response will be nil.
			_, err := client.Query(field.Set(0, i))
			if err != nil {
				log.Fatalf("set failed: %s", err)
			}
		}
	}
}

var (
	client              *pilosa.Client
	schema              *pilosa.Schema
	myindex             *pilosa.Index
	nearMetroField      *pilosa.Field
	privateParkingField *pilosa.Field
	terraceField        *pilosa.Field
	reservationsField   *pilosa.Field
	veganFriendlyField  *pilosa.Field
	expensiveField      *pilosa.Field
)

func main() {
	var err error

	// Create the default client
	client = pilosa.DefaultClient()

	// Retrieve the schema
	schema, err = client.Schema()
	if err != nil {
		log.Fatalf("creating schema failed: %s", err)
	}

	// Create an Index object
	myindex = schema.Index("myindex")

	nearMetroField = myindex.Field("near-metro")
	privateParkingField = myindex.Field("private-parking")
	terraceField = myindex.Field("terrace")
	reservationsField = myindex.Field("reservations")
	veganFriendlyField = myindex.Field("vegan-friendly")
	expensiveField = myindex.Field("expensive")

	// make sure the index and the field exists on the server
	err = client.SyncSchema(schema)
	if err != nil {
		log.Fatalf("error while syncing schema: %s", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	log.Print("filling the data...")

	fill(r, nearMetroField, 0.1)
	fill(r, privateParkingField, 0.01)
	fill(r, terraceField, 0.05)
	fill(r, reservationsField, 0.95)
	fill(r, veganFriendlyField, 0.2)
	fill(r, expensiveField, 0.1)

	log.Print("finished filling the data")

	var response *pilosa.QueryResponse

	response, err = client.Query(myindex.Intersect(
		terraceField.Row(0),
		myindex.Not(expensiveField.Row(0)),
		reservationsField.Row(0)))

	if err != nil {
		log.Fatalf("error while querying pilosa: %s", err)
	}

	// Get the result
	result := response.Result()

	// Act on the result
	if result != nil {
		columns := result.Row().Columns
		log.Printf("got %d columns", len(columns))
	}
}
