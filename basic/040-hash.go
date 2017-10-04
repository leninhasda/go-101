package main

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {

	s := "secret"

	h := sha1.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	p(s)
	pf("%x\n", bs)

	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	p(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	p(string(sDec))
	p("")

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))
}
