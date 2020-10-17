// Author: d1y<chenhonzhou@gmail.com>
// nuxt parse

package nuxtparse

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
)

// EasyParseURL easy parse URL
func EasyParseURL(URL string) (*goja.Object, error) {
	dom, err := goquery.NewDocument(URL)
	if err != nil {
		return nil, err
	}
	var sele = dom.Find("body script")
	var text = sele.Text()
	var matchText = "window."
	if strings.Index(text, matchText) == 0 {
		text = text[len(matchText):]
	}
	vm := goja.New()
	v, err := vm.RunString(text)
	if err != nil {
		return nil, err
	}
	var objectWord = "__NUXT__"
	v.ToObject(vm)
	var jsonData = vm.Get(objectWord)
	var data = jsonData.ToObject(vm)
	return data, nil
}
