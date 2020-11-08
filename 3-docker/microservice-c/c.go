package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

type Result struct {
	Status string
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

	result := Result{Status: isValid}

	jsonData, err := json.Marshal(result)

	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}
