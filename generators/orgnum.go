package generators

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/PatrikOlin/fejkApi/util"
	"github.com/ShiraazMoollatjie/goluhn"
)

func getFormattedOrgNum(timeSeed int64) string {
	orgNumLuhn := getOrgNum(timeSeed)

	for i := 6; i < len(orgNumLuhn); i += 7 {
		orgNumLuhn = orgNumLuhn[:i] + "-" + orgNumLuhn[i:]
	}

	return orgNumLuhn
}

func getOrgNum(timeSeed int64) string {
	rand.Seed(timeSeed)
	orgNum := orgNumPartial()

	_, orgNumLuhn, err := goluhn.Calculate(orgNum)
	util.Check(err)

	return orgNumLuhn
}

func orgNumPartial() string {
	group := strconv.Itoa(util.RangeIn(2, 9))
	second := strconv.Itoa(util.RangeIn(1, 9))
	leading := strconv.Itoa(util.RangeIn(20, 99))
	numbers := strconv.Itoa(util.RangeIn(20, 99))
	ending := strconv.Itoa(util.RangeIn(100, 999))

	stringArray := []string{group, second, leading, numbers, ending}

	return strings.Join(stringArray, "")

}

func toOrgNumString(org []int) string {
	string := make([]string, len(org))

	for v := range org {
		string = append(string, strconv.Itoa(v))
	}
	return strings.Join(string, "")
}
