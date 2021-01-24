package generators

import (
	"github.com/PatrikOlin/fejkApi/db"
	"github.com/PatrikOlin/fejkApi/util"
	"math/rand"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Company struct {
	CompanyName string `json:"companyName"`
	OrgNum      string `json:"orgNum"`
	VatCode     string `json:"vatCode"`
	Address     string `json:"address"`
	Bankgiro    string `json:"bankgiro"`
	CEO         Person `json:"CEO"`
}

func GetCompanies(amount int) []Company {
	var wg sync.WaitGroup
	wg.Add(amount)
	var companies []Company

	for i := 1; i <= amount; i++ {
		go func(i int) {
			defer wg.Done()
			companies = append(companies, generateCompany())
		}(i)
	}

	wg.Wait()
	return companies
}

func generateCompany() Company {
	seed := time.Now().UnixNano()

	company := Company{
		CompanyName: getCompanyname(),
		OrgNum:      getFormattedOrgNum(seed),
		VatCode:     getVatNumForOrgNum(getOrgNum(seed)),
		Address:     getFullAddress(),
		Bankgiro:    getBankgiro(),
		CEO:         generatePerson(),
	}

	company.CEO.Email = strings.Split(company.CEO.Email, "@")[0] + getCompanyDomain(company.CompanyName)
	return company
}

func getCompanyDomain(name string) string {
	tlds := []string{".se", ".dk", ".com", ".fi", ".tk", ".org", ".xyz", ".nu", ".gov"}
	r := rand.Intn(len(tlds))
	if len(name) > 25 {
		n := strings.Split(name, " ")
		n = n[:len(n)-1]
		name = strings.Join(n, " ")
	}	
	name = strings.ReplaceAll(name, " ", "")

	e := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(e, name)
	
	return strings.ToLower("@" + result + tlds[r])
}

func getCompanyname() string {
	rand.Seed(time.Now().UnixNano())
	numOfWords := util.RangeIn(2, 4)
	numType := util.RangeIn(0, 100)
	var companyname strings.Builder
	if numType < 5 {
		pt := db.GetRandomLine("companynametypes")
		companyname.WriteString(pt)
		companyname.WriteString(" ")
	}

	for i := 0; i < numOfWords; i++ {
		s := db.GetRandomLine("companynameparts")
		if s == "&" && (i == 0 || i == numOfWords-1) {
			continue
		}
		companyname.WriteString(s)
		companyname.WriteString(" ")
	}
		
	if numType <= 25 && numType >= 5 {
		t := db.GetRandomLine("companynametypes")
		companyname.WriteString(t)
	}

	finishedCompanyname := strings.TrimSpace(companyname.String())

	return finishedCompanyname
}

func getCompanyType() string {
	return db.GetRandomLine("companynametypes")
}

func getVatNumForOrgNum(orgNum string) string {
	prefix := "SE"
	suffix := "01"

	parts := []string{prefix, orgNum, suffix}

	vatcode := strings.Join(parts, "")

	return vatcode
}

