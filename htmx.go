// Package htmx provides HTMX integration for NodX Go and stdlib server utilities.
//
//   - https://htmx.org/reference
//   - https://github.com/nodxdev/nodxgo
package htmx

import nodx "github.com/nodxdev/nodxgo"

// Hx is an attribute that renders a hx-[key]="[value]" attribute.
func Hx(key string, value string) nodx.Node {
	return nodx.Attr("hx-"+key, value)
}

// HxGet is a NodX Node that renders hx-get="[value]" attribute.
//
// https://htmx.org/attributes/hx-get/
func HxGet(value string) nodx.Node {
	return Hx("get", value)
}

// HxPost is a NodX Node that renders hx-post="[value]" attribute.
//
// https://htmx.org/attributes/hx-post/
func HxPost(value string) nodx.Node {
	return Hx("post", value)
}

// HxPut is a NodX Node that renders hx-put="[value]" attribute.
//
// https://htmx.org/attributes/hx-put/
func HxPut(value string) nodx.Node {
	return Hx("put", value)
}

// HxPatch is a NodX Node that renders hx-patch="[value]" attribute.
//
// https://htmx.org/attributes/hx-patch/
func HxPatch(value string) nodx.Node {
	return Hx("patch", value)
}

// HxDelete is a NodX Node that renders hx-delete="[value]" attribute.
//
// https://htmx.org/attributes/hx-delete/
func HxDelete(value string) nodx.Node {
	return Hx("delete", value)
}

// HxOn is a Nodx Node that renders hx-on:[eventName]="[value]" attribute.
//
// https://htmx.org/attributes/hx-on/
func HxOn(eventName string, value string) nodx.Node {
	return Hx("on:"+eventName, value)
}

// HxPushURL is a NodX Node that renders hx-push-url="[value]" attribute.
//
// https://htmx.org/attributes/hx-push-url/
func HxPushURL(value string) nodx.Node {
	return Hx("push-url", value)
}

// HxSelect is a NodX Node that renders hx-select="[value]" attribute.
//
// https://htmx.org/attributes/hx-select/
func HxSelect(value string) nodx.Node {
	return Hx("select", value)
}

// HxSelectOOB is a NodX Node that renders hx-select-oob="[value]" attribute.
//
// https://htmx.org/attributes/hx-select-oob/
func HxSelectOOB(value string) nodx.Node {
	return Hx("select-oob", value)
}

// HxSwap is a NodX Node that renders hx-swap="[value]" attribute.
//
// https://htmx.org/attributes/hx-swap/
func HxSwap(value string) nodx.Node {
	return Hx("swap", value)
}

// HxSwapOOB is a NodX Node that renders hx-swap-oob="[value]" attribute.
//
// https://htmx.org/attributes/hx-swap-oob/
func HxSwapOOB(value string) nodx.Node {
	return Hx("swap-oob", value)
}

// HxTarget returns a NodX node with the hx-target
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-target/
func HxTarget(value string) nodx.Node {
	return Hx("target", value)
}

// HxTrigger returns a NodX node with the hx-trigger
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-trigger/
func HxTrigger(value string) nodx.Node {
	return Hx("trigger", value)
}

// HxVals returns a NodX node with the hx-vals
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-vals/
func HxVals(value string) nodx.Node {
	return Hx("vals", value)
}

// HxBoost returns a NodX node with the hx-boost
// attribute set to the given value.
//
// See https://htmx.org/attributes/hx-boost/
func HxBoost(value string) nodx.Node {
	return Hx("boost", value)
}

// HxConfirm returns a NodX node with the hx-confirm
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-confirm/
func HxConfirm(value string) nodx.Node {
	return Hx("confirm", value)
}

// HxDisable returns a NodX node with the hx-disable
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-disable/
func HxDisable(value string) nodx.Node {
	return Hx("disable", value)
}

// HxDisabledELT returns a NodX node with the hx-disabled-elt
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-disabled-elt/
func HxDisabledELT(value string) nodx.Node {
	return Hx("disabled-elt", value)
}

// HxDisinherit returns a NodX node with the hx-disinherit
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-disinherit/
func HxDisinherit(value string) nodx.Node {
	return Hx("disinherit", value)
}

// HxEncoding returns a NodX node with the hx-encoding
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-encoding/
func HxEncoding(value string) nodx.Node {
	return Hx("encoding", value)
}

// HxExt returns a NodX node with the hx-ext
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-ext/
func HxExt(value string) nodx.Node {
	return Hx("ext", value)
}

// HxHeaders returns a NodX node with the hx-headers
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-headers/
func HxHeaders(value string) nodx.Node {
	return Hx("headers", value)
}

// HxHistory returns a NodX node with the hx-history
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-history/
func HxHistory(value string) nodx.Node {
	return Hx("history", value)
}

// HxHistoryElt returns a NodX node with the hx-history-elt
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-history-elt/
func HxHistoryElt(value string) nodx.Node {
	return Hx("history-elt", value)
}

// HxInclude returns a NodX node with the hx-include
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-include/
func HxInclude(value string) nodx.Node {
	return Hx("include", value)
}

// HxIndicator returns a NodX node with the hx-indicator
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-indicator/
func HxIndicator(value string) nodx.Node {
	return Hx("indicator", value)
}

// HxInherit returns a NodX node with the hx-inherit
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-inherit/
func HxInherit(value string) nodx.Node {
	return Hx("inherit", value)
}

// HxParams returns a NodX node with the hx-params
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-params/
func HxParams(value string) nodx.Node {
	return Hx("params", value)
}

// HxPreserve returns a NodX node with the hx-preserve
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-preserve/
func HxPreserve(value string) nodx.Node {
	return Hx("preserve", value)
}

// HxPrompt returns a NodX node with the hx-prompt
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-prompt/
func HxPrompt(value string) nodx.Node {
	return Hx("prompt", value)
}

// HxReplaceURL returns a NodX node with the hx-replace-url
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-replace-url/
func HxReplaceURL(value string) nodx.Node {
	return Hx("replace-url", value)
}

// HxRequest returns a NodX node with the hx-request
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-request/
func HxRequest(value string) nodx.Node {
	return Hx("request", value)
}

// HxSync returns a NodX node with the hx-sync
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-sync/
func HxSync(value string) nodx.Node {
	return Hx("sync", value)
}

// HxValidate returns a NodX node with the hx-validate
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-validate/
func HxValidate(value string) nodx.Node {
	return Hx("validate", value)
}

// HxVars is a NodX Node that renders hx-vars="[value]" attribute.
//
// https://htmx.org/attributes/hx-vars/
func HxVars(value string) nodx.Node {
	return Hx("vars", value)
}
