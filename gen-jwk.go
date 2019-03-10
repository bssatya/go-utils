package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte("testSecret")
	str := base64.RawURLEncoding.EncodeToString(data)
	fmt.Println("base64:", str)
	
	jwk := map[string]interface{}{
		// "k": data,
		"k": str,
		"alg":"HS256",
		"kid":"default",
		"kty":"oct",
	}
	
	json,_ := json.MarshalIndent(jwk, "", "\t")
	
	fmt.Println(string(json))
}
