package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PatrikOlin/fejkApi/generators"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	amount, err := getAmountParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "invalid datatype for parameter"}`))
		return
	}

	data := generators.GetArticles(amount)

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	amount, err := getAmountParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "invalid datatype for parameter"}`))
		return
	}

	data := generators.GetPeople(amount)

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	amount, err := getAmountParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "invalid datatype for parameter"}`))
		return
	}

	data := generators.GetCompanies(amount)

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getAmountParam(r *http.Request) (int, error) {
	amount := 1
	queryParams := r.URL.Query()
	a := queryParams.Get("amount")
	if a != "" {
		val, err := strconv.Atoi(a)
		if err != nil {
			return amount, err
		}
		amount = val
	}
	if amount > 100 {
		amount = 100
	}
	return amount, nil
}
