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

func ExampleHxGet() {
	node := nodx.Div(
		htmx.HxGet("/api/data"),
	)
	fmt.Println(node)
	// Output: <div hx-get="/api/data"></div>
}

func ExampleHxPost() {
	node := nodx.Div(
		htmx.HxPost("/api/data"),
	)
	fmt.Println(node)
	// Output: <div hx-post="/api/data"></div>
}

func ExampleHxPut() {
	node := nodx.Div(
		htmx.HxPut("/api/data"),
	)
	fmt.Println(node)
	// Output: <div hx-put="/api/data"></div>
}

func ExampleHxPatch() {
	node := nodx.Div(
		htmx.HxPatch("/api/data"),
	)
	fmt.Println(node)
	// Output: <div hx-patch="/api/data"></div>
}

func ExampleHxDelete() {
	node := nodx.Div(
		htmx.HxDelete("/api/data"),
	)
	fmt.Println(node)
	// Output: <div hx-delete="/api/data"></div>
}

func ExampleHxOn() {
	node := nodx.Div(
		htmx.HxOn("click", "alert(true)"),
	)
	fmt.Println(node)
	// Output: <div hx-on:click="alert(true)"></div>
}

func ExampleHxPushURL() {
	node := nodx.Div(
		htmx.HxPushURL("/new-url"),
	)
	fmt.Println(node)
	// Output: <div hx-push-url="/new-url"></div>
}

func ExampleHxSelect() {
	node := nodx.Div(
		htmx.HxSelect("#content"),
	)
	fmt.Println(node)
	// Output: <div hx-select="#content"></div>
}

func ExampleHxSelectOOB() {
	node := nodx.Div(
		htmx.HxSelectOOB("#content"),
	)
	fmt.Println(node)
	// Output: <div hx-select-oob="#content"></div>
}

func ExampleHxSwap() {
	node := nodx.Div(
		htmx.HxSwap("outerHTML"),
	)
	fmt.Println(node)
	// Output: <div hx-swap="outerHTML"></div>
}

func ExampleHxSwapOOB() {
	node := nodx.Div(
		htmx.HxSwapOOB("innerHTML"),
	)
	fmt.Println(node)
	// Output: <div hx-swap-oob="innerHTML"></div>
}

func ExampleHxTarget() {
	node := nodx.Div(
		htmx.HxTarget("#target"),
	)
	fmt.Println(node)
	// Output: <div hx-target="#target"></div>
}

func ExampleHxTrigger() {
	node := nodx.Div(
		htmx.HxTrigger("click"),
	)
	fmt.Println(node)
	// Output: <div hx-trigger="click"></div>
}

func ExampleHxVals() {
	node := nodx.Div(
		htmx.HxVals("{}"),
	)
	fmt.Println(node)
	// Output: <div hx-vals="{}"></div>
}

func ExampleHxBoost() {
	node := nodx.Div(
		htmx.HxBoost("true"),
	)
	fmt.Println(node)
	// Output: <div hx-boost="true"></div>
}

func ExampleHxConfirm() {
	node := nodx.Div(
		htmx.HxConfirm("Are you sure?"),
	)
	fmt.Println(node)
	// Output: <div hx-confirm="Are you sure?"></div>
}

func ExampleHxDisable() {
	node := nodx.Div(
		htmx.HxDisable("true"),
	)
	fmt.Println(node)
	// Output: <div hx-disable="true"></div>
}

func ExampleHxDisabledELT() {
	node := nodx.Div(
		htmx.HxDisabledELT("#button"),
	)
	fmt.Println(node)
	// Output: <div hx-disabled-elt="#button"></div>
}

func ExampleHxDisinherit() {
	node := nodx.Div(
		htmx.HxDisinherit("true"),
	)
	fmt.Println(node)
	// Output: <div hx-disinherit="true"></div>
}

func ExampleHxEncoding() {
	node := nodx.Div(
		htmx.HxEncoding("UTF-8"),
	)
	fmt.Println(node)
	// Output: <div hx-encoding="UTF-8"></div>
}

func ExampleHxExt() {
	node := nodx.Div(
		htmx.HxExt("ext-value"),
	)
	fmt.Println(node)
	// Output: <div hx-ext="ext-value"></div>
}

func ExampleHxHeaders() {
	node := nodx.Div(
		htmx.HxHeaders("{}"),
	)
	fmt.Println(node)
	// Output: <div hx-headers="{}"></div>
}

func ExampleHxHistory() {
	node := nodx.Div(
		htmx.HxHistory("true"),
	)
	fmt.Println(node)
	// Output: <div hx-history="true"></div>
}

func ExampleHxHistoryElt() {
	node := nodx.Div(
		htmx.HxHistoryElt("#history"),
	)
	fmt.Println(node)
	// Output: <div hx-history-elt="#history"></div>
}

func ExampleHxInclude() {
	node := nodx.Div(
		htmx.HxInclude("#include"),
	)
	fmt.Println(node)
	// Output: <div hx-include="#include"></div>
}

func ExampleHxIndicator() {
	node := nodx.Div(
		htmx.HxIndicator("#spinner"),
	)
	fmt.Println(node)
	// Output: <div hx-indicator="#spinner"></div>
}

func ExampleHxInherit() {
	node := nodx.Div(
		htmx.HxInherit("true"),
	)
	fmt.Println(node)
	// Output: <div hx-inherit="true"></div>
}

func ExampleHxParams() {
	node := nodx.Div(
		htmx.HxParams("a,b,c"),
	)
	fmt.Println(node)
	// Output: <div hx-params="a,b,c"></div>
}

func ExampleHxPreserve() {
	node := nodx.Div(
		htmx.HxPreserve("true"),
	)
	fmt.Println(node)
	// Output: <div hx-preserve="true"></div>
}

func ExampleHxPrompt() {
	node := nodx.Div(
		htmx.HxPrompt("Enter value:"),
	)
	fmt.Println(node)
	// Output: <div hx-prompt="Enter value:"></div>
}

func ExampleHxReplaceURL() {
	node := nodx.Div(
		htmx.HxReplaceURL("/replace"),
	)
	fmt.Println(node)
	// Output: <div hx-replace-url="/replace"></div>
}

func ExampleHxRequest() {
	node := nodx.Div(
		htmx.HxRequest("GET"),
	)
	fmt.Println(node)
	// Output: <div hx-request="GET"></div>
}

func ExampleHxSync() {
	node := nodx.Div(
		htmx.HxSync("true"),
	)
	fmt.Println(node)
	// Output: <div hx-sync="true"></div>
}

func ExampleHxValidate() {
	node := nodx.Div(
		htmx.HxValidate("true"),
	)
	fmt.Println(node)
	// Output: <div hx-validate="true"></div>
}

func ExampleHxVars() {
	node := nodx.Div(
		htmx.HxVars("{}"),
	)
	fmt.Println(node)
	// Output: <div hx-vars="{}"></div>
}
