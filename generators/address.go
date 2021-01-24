package generators

import (
	"github.com/PatrikOlin/fejkApi/db"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getFullAddress() string {
	rand.Seed(time.Now().UnixNano())
	var fullAddress strings.Builder
	fullAddress.WriteString(getStreet())
	fullAddress.WriteString(",\n")
	fullAddress.WriteString(db.GetPostalAddress())

	return fullAddress.String()
}

func getStreetPrefix() string {
	stPrefix := db.GetRandomLine("streetprefixes")
	return stPrefix
}

func getStreetSuffix() string {
	stPrefix := db.GetRandomLine("streetsuffixes")
	return stPrefix
}

func getStreet() string {
	var street strings.Builder
	street.WriteString(getStreetPrefix())
	street.WriteString(getStreetSuffix())
	street.WriteString(" ")
	street.WriteString(strconv.Itoa(rand.Intn(198) + 1))
	if rand.Intn(8) < 2 {
		street.WriteString(randomLetter())
	}

	return street.String()
}

func randomLetter() string {
	var letters = []rune("ABCDEFGHIJK")
	letter := make([]rune, 1)
	for i := range letter {
		letter[i] = letters[rand.Intn(len(letters))]
	}

	return string(letter)
}

// func getPostalAddress() string {
// 	var id int
// 	var postalTown string
// 	var postalCode string

// 	db, _ := sql.Open("sqlite3", "./datastore/store.db")
// 	row := db.QueryRow("SELECT * FROM postaladdresses ORDER BY RANDOM() LIMIT 1")
// 	err := row.Scan(&id, &postalCode, &postalTown)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No rows found")
// 		} else {
// 			panic(err)
// 		}
// 	}

// 	db.Close()

// 	return postalCode + ", " + postalTown
// }
