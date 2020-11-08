package main

import (
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
	
	"github.com/hashicorp/go-retryablehttp"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

type Result struct {
	Status string
	Discount string
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return "valid"
		}
	}
	return "invalid"
}

var coupons Coupons 

func main()  {
	coupons.Coupon = append(
		coupons.Coupon,  
		Coupon{ Code: "GO10" },  
		Coupon{ Code: "GO15" },
		Coupon{ Code: "GO20" },
	)

	http.HandleFunc("/", Home)
	http.ListenAndServe(":9092", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	isValid := coupons.Check(coupon)

	resultCouponDiscount := MakeHttpCall("http://localhost:9093", coupon, isValid)
	
	result := Result{Status: resultCouponDiscount.Status, Discount: resultCouponDiscount.Discount}

	jsonData, err := json.Marshal(result)

	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))

}

func MakeHttpCall(urlMicroservice string, coupon string, isValid string) Result {

	values := url.Values{}
	values.Add("coupon", coupon);
	values.Add("isValid", isValid);

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	res, err := retryClient.PostForm(urlMicroservice, values)

	if err != nil {
		result := Result{Status: "Servidor fora do ar!",Discount: ""}

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