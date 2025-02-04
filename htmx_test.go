package htmx_test

import (
	"fmt"

	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
)

func ExampleHx() {
	node := nodx.Div(
		htmx.Hx("get", "/api/data"),
	)
	fmt.Println(node)
	// Output: <div hx-get="/api/data"></div>
}
