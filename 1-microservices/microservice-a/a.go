package main

import (
	"html/template"
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"encoding/json"
	
	"github.com/hashicorp/go-retryablehttp"
)

type Result struct {
	Status string
	Discount string
}

func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/process", Process)
	http.ListenAndServe(":9090", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, Result{})
}

func Process(w http.ResponseWriter, r *http.Request) {

	result := MakeHttpCall("http://localhost:9091", r.FormValue("coupon"), r.FormValue("cc-number"))

	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, result)
}

func MakeHttpCall(urlMicroservice string, coupon string, ccNumber string) Result {

	values := url.Values{}
	values.Add("coupon", coupon);
	values.Add("ccNumber", ccNumber);

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	res, err := retryClient.PostForm(urlMicroservice, values)

	if err != nil {
		result := Result{Status: "Servidor fora do ar!", Discount: ""}

		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result
}