//package mike //go run: cannot run non-main package
package main

import b64 "encoding/base64"
import "fmt"

var p = fmt.Println

func main() {
	data := "abc123!@#!#!@313'-=@$%^"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	p("sEnc:", sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	p("sDec:", string(sDec))
	p()

	urlEnc := b64.URLEncoding.EncodeToString([]byte(data))
	p("urlEnc:", urlEnc)
	urlDec, _ := b64.URLEncoding.DecodeString(urlEnc)
	p("urlDec:", string(urlDec))
}
