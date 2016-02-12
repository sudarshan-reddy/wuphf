package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strings"
)

const (
	accountSid = "ACff2b52113ad171f7aa2976c8a3f55f7c"
	authToken = "abb129ba34817882b26af8c7324508fb"
	urlStr =  "https://api.twilio.com/2010-04-01/Accounts/" + 
				accountSid + "/Messages.json"
)

func main(){
	v := url.Values{}
	v.Set("To" , "+919789099524")
	v.Set("From" ,"+17722176343" )
	v.Set("Body", "Wuphf!")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req , err := http.NewRequest("POST", urlStr, &rb)
	if err != nil {
		fmt.Println("ERROR: " , err)	
	}else{
		req.SetBasicAuth(accountSid, authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")	
		resp, err := client.Do(req)
		if err != nil{
			fmt.Println("ERROR: ", err)	
		}else{
			fmt.Println("Success: " , resp.Status)	
		}
	}


}
