/*

This package is used to retrieve about information from a document
in the Fauna database.

*/

package fauna

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
	"os"
)

// The following structs are used to store the response
// from the Fauna database
type AboutData struct {
	Title   string   `fauna:"title"`
	Details []string `fauna:"details"`
}

type About struct {
	Data AboutData `fauna:"data"`
	Ts   float64   `fauna:"ts"`
}

// Calls the Fauna API and retrieves the information about
// the about section
func GetAboutData() AboutData {
	client := f.NewFaunaClient(os.Getenv("FAUNA_SECRET"))

	res, err := client.Query(f.Get(f.Ref(f.Collection(os.Getenv("FAUNA_COLLECTION_NAME")), os.Getenv("FAUNA_ITEM_NUM"))))
	if err != nil {
		panic(err)
	}

	var about About

	if err := res.Get(&about); err != nil {
		panic(err)
	}
	return about.Data
}
