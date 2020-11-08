package main

import (
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Result struct {
	Status string
	Discount string
}

func main()  {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":9091", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	ccNumber := r.PostFormValue("ccNumber")

	resultCoupon := MakeHttpCall("http://localhost:9092", coupon)

	lengthCCNumber := len([]rune(ccNumber))

	result := Result{Status: "Declined", Discount: ""}

	if lengthCCNumber == 16 && resultCoupon.Status == "valid" {
		result.Status = "Approved"
		result.Discount = resultCoupon.Discount
	}

	if resultCoupon.Status == "invalid" {
		result.Status = "Invalid coupon"
		result.Discount = resultCoupon.Discount
	}

	jsonData, err := json.Marshal(result)

	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}

func MakeHttpCall(urlMicroservice string, coupon string) Result {

	values := url.Values{}
	values.Add("coupon", coupon);

	res, err := http.PostForm(urlMicroservice, values)

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