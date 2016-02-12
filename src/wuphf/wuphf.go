package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strings"
)

const (
	accountSid = "afasdffdgfsdfsfs"
	authToken = "asdfafdafdfsdfsf"
	urlStr =  "https://api.twilio.com/2010-04-01/Accounts/" + 
				accountSid + "/Messages.json"
)

func main(){
	v := url.Values{}
	v.Set("To" , "123145124134")
	v.Set("From" ,"124124144141" )
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
