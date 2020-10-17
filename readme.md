# nuxtparse

Vue nuxt framework use SSR

the first script

![Jietu20201017-133625.gif](https://i.loli.net/2020/10/17/jUYbxRF2ofkNAda.gif)

```js
window.__NUXT__ = {
  // ... => the more json data
}
```

![Jietu20201017-133913.gif](https://i.loli.net/2020/10/17/vn3k5EUi8tGM6KB.gif)

> the idea by: https://github.com/nuxt/nuxt.js/issues/2822

try write parse html codes to json

```
go get github.com/d1y/nuxtparse
```

example

```go
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

```

- github.com/PuerkitoBio/goquery
- github.com/dop251/goja