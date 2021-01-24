package generators

import (
	"math/rand"
	"time"

	"github.com/PatrikOlin/fejkApi/util"
)
		  
func getBankgiro() string {
	return formatBankgiro(generateBankgiro())
}		  

func generateBankgiro() string {
	rand.Seed(time.Now().UTC().UnixNano())
	c := rand.Intn(10) % 2 == 0
	var num string

	if c {
		num = util.GenerateLuhns(7)
	} else {
		num = util.GenerateLuhns(8)
	}


	return num
}

func formatBankgiro(raw string) string {
	index := len(raw) - 4
	f := raw[:index] + "-" + raw[index:]
		  
	return f
}
