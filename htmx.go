// Package htmx provides HTMX integration for NodX Go and stdlib server utilities.
//
//   - https://htmx.org/reference
//   - https://github.com/nodxdev/nodxgo
package htmx

import nodx "github.com/nodxdev/nodxgo"

// Hx renders an hx-[key]="[value]" attribute.
//
// https://htmx.org/reference
func Hx(key string, value string) nodx.Node {
	return nodx.Attr("hx-"+key, value)
}

// HxGet renders an hx-get="[value]" attribute.
//
// https://htmx.org/attributes/hx-get/
func HxGet(value string) nodx.Node {
	return Hx("get", value)
}

// HxPost renders an hx-post="[value]" attribute.
//
// https://htmx.org/attributes/hx-post/
func HxPost(value string) nodx.Node {
	return Hx("post", value)
}

// HxPut renders an hx-put="[value]" attribute.
//
// https://htmx.org/attributes/hx-put/
func HxPut(value string) nodx.Node {
	return Hx("put", value)
}

// HxPatch renders an hx-patch="[value]" attribute.
//
// https://htmx.org/attributes/hx-patch/
func HxPatch(value string) nodx.Node {
	return Hx("patch", value)
}

// HxDelete renders an hx-delete="[value]" attribute.
//
// https://htmx.org/attributes/hx-delete/
func HxDelete(value string) nodx.Node {
	return Hx("delete", value)
}

// HxOn renders an hx-on:[eventName]="[value]" attribute.
//
// https://htmx.org/attributes/hx-on/
func HxOn(eventName string, value string) nodx.Node {
	return Hx("on:"+eventName, value)
}

// HxPushURL renders an hx-push-url="[value]" attribute.
//
// https://htmx.org/attributes/hx-push-url/
func HxPushURL(value string) nodx.Node {
	return Hx("push-url", value)
}

// HxSelect renders an hx-select="[value]" attribute.
//
// https://htmx.org/attributes/hx-select/
func HxSelect(value string) nodx.Node {
	return Hx("select", value)
}

// HxSelectOOB renders an hx-select-oob="[value]" attribute.
//
// https://htmx.org/attributes/hx-select-oob/
func HxSelectOOB(value string) nodx.Node {
	return Hx("select-oob", value)
}

// HxSwap renders an hx-swap="[value]" attribute.
//
// https://htmx.org/attributes/hx-swap/
func HxSwap(value string) nodx.Node {
	return Hx("swap", value)
}

// HxSwapOOB renders an hx-swap-oob="[value]" attribute.
//
// https://htmx.org/attributes/hx-swap-oob/
func HxSwapOOB(value string) nodx.Node {
	return Hx("swap-oob", value)
}

// HxTarget renders an hx-target="[value]" attribute.
//
// https://htmx.org/attributes/hx-target/
func HxTarget(value string) nodx.Node {
	return Hx("target", value)
}

// HxTrigger renders an hx-trigger="[value]" attribute.
//
// https://htmx.org/attributes/hx-trigger/
func HxTrigger(value string) nodx.Node {
	return Hx("trigger", value)
}

// HxVals renders an hx-vals="[value]" attribute.
//
// https://htmx.org/attributes/hx-vals/
func HxVals(value string) nodx.Node {
	return Hx("vals", value)
}

// HxBoost renders an hx-boost="[value]" attribute.
//
// https://htmx.org/attributes/hx-boost/
func HxBoost(value string) nodx.Node {
	return Hx("boost", value)
}

// HxConfirm renders an hx-confirm="[value]" attribute.
//
// https://htmx.org/attributes/hx-confirm/
func HxConfirm(value string) nodx.Node {
	return Hx("confirm", value)
}

// HxDisable renders an hx-disable="[value]" attribute.
//
// https://htmx.org/attributes/hx-disable/
func HxDisable(value string) nodx.Node {
	return Hx("disable", value)
}

// HxDisabledELT renders an hx-disabled-elt="[value]" attribute.
//
// https://htmx.org/attributes/hx-disabled-elt/
func HxDisabledELT(value string) nodx.Node {
	return Hx("disabled-elt", value)
}

// HxDisinherit renders an hx-disinherit="[value]" attribute.
//
// https://htmx.org/attributes/hx-disinherit/
func HxDisinherit(value string) nodx.Node {
	return Hx("disinherit", value)
}

// HxEncoding renders an hx-encoding="[value]" attribute.
//
// https://htmx.org/attributes/hx-encoding/
func HxEncoding(value string) nodx.Node {
	return Hx("encoding", value)
}

// HxExt renders an hx-ext="[value]" attribute.
//
// https://htmx.org/attributes/hx-ext/
func HxExt(value string) nodx.Node {
	return Hx("ext", value)
}

// HxHeaders renders an hx-headers="[value]" attribute.
//
// https://htmx.org/attributes/hx-headers/
func HxHeaders(value string) nodx.Node {
	return Hx("headers", value)
}

// HxHistory renders an hx-history="[value]" attribute.
//
// https://htmx.org/attributes/hx-history/
func HxHistory(value string) nodx.Node {
	return Hx("history", value)
}

// HxHistoryElt renders an hx-history-elt="[value]" attribute.
//
// https://htmx.org/attributes/hx-history-elt/
func HxHistoryElt(value string) nodx.Node {
	return Hx("history-elt", value)
}

// HxInclude renders an hx-include="[value]" attribute.
//
// https://htmx.org/attributes/hx-include/
func HxInclude(value string) nodx.Node {
	return Hx("include", value)
}

// HxIndicator renders an hx-indicator="[value]" attribute.
//
// https://htmx.org/attributes/hx-indicator/
func HxIndicator(value string) nodx.Node {
	return Hx("indicator", value)
}

// HxInherit renders an hx-inherit="[value]" attribute.
//
// https://htmx.org/attributes/hx-inherit/
func HxInherit(value string) nodx.Node {
	return Hx("inherit", value)
}

// HxParams renders an hx-params="[value]" attribute.
//
// https://htmx.org/attributes/hx-params/
func HxParams(value string) nodx.Node {
	return Hx("params", value)
}

// HxPreserve renders an hx-preserve="[value]" attribute.
//
// https://htmx.org/attributes/hx-preserve/
func HxPreserve(value string) nodx.Node {
	return Hx("preserve", value)
}

// HxPrompt renders an hx-prompt="[value]" attribute.
//
// https://htmx.org/attributes/hx-prompt/
func HxPrompt(value string) nodx.Node {
	return Hx("prompt", value)
}

// HxReplaceURL renders an hx-replace-url="[value]" attribute.
//
// https://htmx.org/attributes/hx-replace-url/
func HxReplaceURL(value string) nodx.Node {
	return Hx("replace-url", value)
}

// HxRequest renders an hx-request="[value]" attribute.
//
// https://htmx.org/attributes/hx-request/
func HxRequest(value string) nodx.Node {
	return Hx("request", value)
}

// HxSync renders an hx-sync="[value]" attribute.
//
// https://htmx.org/attributes/hx-sync/
func HxSync(value string) nodx.Node {
	return Hx("sync", value)
}

// HxValidate renders an hx-validate="[value]" attribute.
//
// https://htmx.org/attributes/hx-validate/
func HxValidate(value string) nodx.Node {
	return Hx("validate", value)
}

// HxVars renders an hx-vars="[value]" attribute.
//
// https://htmx.org/attributes/hx-vars/
func HxVars(value string) nodx.Node {
	return Hx("vars", value)
}
