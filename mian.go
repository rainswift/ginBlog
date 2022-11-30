package main

import (
	"ginBlog/router"
)

func main() {
	//s, err := api.GenToken("ghz", "123")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("s: %v\n", s)
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdoeiIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNjY5ODAwOTk0LCJpc3MiOiJsYW9ndW8ifQ.HmgzxMK33jKcvaiIIyzCzwD7_5jDxJpEjEUKWwfdEgw"
	//mc, err := api.ParseToken(token)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("mc.Password: %v\n", mc.Password)
	//fmt.Printf("mc.Username: %v\n", mc.Username)

	router.Start()
}
