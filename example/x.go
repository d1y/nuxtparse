package main

import (
	"fmt"

	"github.com/d1y/nuxtparse"
)

func main() {
	var localTestServer = "https://jikipedia.com/definition/514476769"
	data, _ := nuxtparse.EasyParseURL(localTestServer)
	var value = data.Get("layout").String()
	fmt.Println("value", value)
}
