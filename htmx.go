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
// Issues a GET to the specified URL.
//
// https://htmx.org/attributes/hx-get/
func HxGet(value string) nodx.Node {
	return Hx("get", value)
}

// HxPost renders an hx-post="[value]" attribute.
//
// Issues a POST to the specified URL.
//
// https://htmx.org/attributes/hx-post/
func HxPost(value string) nodx.Node {
	return Hx("post", value)
}

// HxPut renders an hx-put="[value]" attribute.
//
// Issues a PUT to the specified URL.
//
// https://htmx.org/attributes/hx-put/
func HxPut(value string) nodx.Node {
	return Hx("put", value)
}

// HxPatch renders an hx-patch="[value]" attribute.
//
// Issues a PATCH to the specified URL.
//
// https://htmx.org/attributes/hx-patch/
func HxPatch(value string) nodx.Node {
	return Hx("patch", value)
}

// HxDelete renders an hx-delete="[value]" attribute.
//
// Issues a DELETE to the specified URL.
//
// https://htmx.org/attributes/hx-delete/
func HxDelete(value string) nodx.Node {
	return Hx("delete", value)
}

// HxOn renders an hx-on:[eventName]="[value]" attribute.
//
// Handle events with inline scripts on elements.
//
// https://htmx.org/attributes/hx-on/
func HxOn(eventName string, value string) nodx.Node {
	return Hx("on:"+eventName, value)
}

// HxPushURL renders an hx-push-url="[value]" attribute.
//
// Push a URL into the browser location bar to create history.
//
// https://htmx.org/attributes/hx-push-url/
func HxPushURL(value string) nodx.Node {
	return Hx("push-url", value)
}

// HxSelect renders an hx-select="[value]" attribute.
//
// Select content to swap in from a response.
//
// https://htmx.org/attributes/hx-select/
func HxSelect(value string) nodx.Node {
	return Hx("select", value)
}

// HxSelectOOB renders an hx-select-oob="[value]" attribute.
//
// Select content to swap in from a response, somewhere other than the target (out of band).
//
// https://htmx.org/attributes/hx-select-oob/
func HxSelectOOB(value string) nodx.Node {
	return Hx("select-oob", value)
}

// HxSwap renders an hx-swap="[value]" attribute.
//
// Controls how content will swap in (outerHTML, beforeend, afterend, …).
//
// https://htmx.org/attributes/hx-swap/
func HxSwap(value string) nodx.Node {
	return Hx("swap", value)
}

// HxSwapOOB renders an hx-swap-oob="[value]" attribute.
//
// Mark element to swap in from a response (out of band).
//
// https://htmx.org/attributes/hx-swap-oob/
func HxSwapOOB(value string) nodx.Node {
	return Hx("swap-oob", value)
}

// HxTarget renders an hx-target="[value]" attribute.
//
// Specifies the target element to be swapped.
//
// https://htmx.org/attributes/hx-target/
func HxTarget(value string) nodx.Node {
	return Hx("target", value)
}

// HxTrigger renders an hx-trigger="[value]" attribute.
//
// Specifies the event that triggers the request.
//
// https://htmx.org/attributes/hx-trigger/
func HxTrigger(value string) nodx.Node {
	return Hx("trigger", value)
}

// HxVals renders an hx-vals="[value]" attribute.
//
// Add values to submit with the request (JSON format).
//
// https://htmx.org/attributes/hx-vals/
func HxVals(value string) nodx.Node {
	return Hx("vals", value)
}

// HxBoost renders an hx-boost="[value]" attribute.
//
// Add progressive enhancement for links and forms.
//
// https://htmx.org/attributes/hx-boost/
func HxBoost(value string) nodx.Node {
	return Hx("boost", value)
}

// HxConfirm renders an hx-confirm="[value]" attribute.
//
// Shows a confirm() dialog before issuing a request.
//
// https://htmx.org/attributes/hx-confirm/
func HxConfirm(value string) nodx.Node {
	return Hx("confirm", value)
}

// HxDisable renders an hx-disable="[value]" attribute.
//
// Disables htmx processing for the given node and any children nodes.
//
// https://htmx.org/attributes/hx-disable/
func HxDisable(value string) nodx.Node {
	return Hx("disable", value)
}

// HxDisabledELT renders an hx-disabled-elt="[value]" attribute.
//
// Adds the disabled attribute to the specified elements while a request is in flight.
//
// https://htmx.org/attributes/hx-disabled-elt/
func HxDisabledELT(value string) nodx.Node {
	return Hx("disabled-elt", value)
}

// HxDisinherit renders an hx-disinherit="[value]" attribute.
//
// Control and disable automatic attribute inheritance for child nodes.
//
// https://htmx.org/attributes/hx-disinherit/
func HxDisinherit(value string) nodx.Node {
	return Hx("disinherit", value)
}

// HxEncoding renders an hx-encoding="[value]" attribute.
//
// Changes the request encoding type.
//
// https://htmx.org/attributes/hx-encoding/
func HxEncoding(value string) nodx.Node {
	return Hx("encoding", value)
}

// HxExt renders an hx-ext="[value]" attribute.
//
// Extensions to use for this element.
//
// https://htmx.org/attributes/hx-ext/
func HxExt(value string) nodx.Node {
	return Hx("ext", value)
}

// HxHeaders renders an hx-headers="[value]" attribute.
//
// Adds to the headers that will be submitted with the request.
//
// https://htmx.org/attributes/hx-headers/
func HxHeaders(value string) nodx.Node {
	return Hx("headers", value)
}

// HxHistory renders an hx-history="[value]" attribute.
//
// Prevent sensitive data being saved to the history cache.
//
// https://htmx.org/attributes/hx-history/
func HxHistory(value string) nodx.Node {
	return Hx("history", value)
}

// HxHistoryElt renders an hx-history-elt="[value]" attribute.
//
// The element to snapshot and restore during history navigation.
//
// https://htmx.org/attributes/hx-history-elt/
func HxHistoryElt(value string) nodx.Node {
	return Hx("history-elt", value)
}

// HxInclude renders an hx-include="[value]" attribute.
//
// Include additional data in requests.
//
// https://htmx.org/attributes/hx-include/
func HxInclude(value string) nodx.Node {
	return Hx("include", value)
}

// HxIndicator renders an hx-indicator="[value]" attribute.
//
// The element to put the htmx-request class on during the request.
//
// https://htmx.org/attributes/hx-indicator/
func HxIndicator(value string) nodx.Node {
	return Hx("indicator", value)
}

// HxInherit renders an hx-inherit="[value]" attribute.
//
// Control and enable automatic attribute inheritance for child nodes if it has been disabled by default.
//
// https://htmx.org/attributes/hx-inherit/
func HxInherit(value string) nodx.Node {
	return Hx("inherit", value)
}

// HxParams renders an hx-params="[value]" attribute.
//
// Filters the parameters that will be submitted with a request.
//
// https://htmx.org/attributes/hx-params/
func HxParams(value string) nodx.Node {
	return Hx("params", value)
}

// HxPreserve renders an hx-preserve="[value]" attribute.
//
// Specifies elements to keep unchanged between requests.
//
// https://htmx.org/attributes/hx-preserve/
func HxPreserve(value string) nodx.Node {
	return Hx("preserve", value)
}

// HxPrompt renders an hx-prompt="[value]" attribute.
//
// Shows a prompt() before submitting a request.
//
// https://htmx.org/attributes/hx-prompt/
func HxPrompt(value string) nodx.Node {
	return Hx("prompt", value)
}

// HxReplaceURL renders an hx-replace-url="[value]" attribute.
//
// Replace the URL in the browser location bar.
//
// https://htmx.org/attributes/hx-replace-url/
func HxReplaceURL(value string) nodx.Node {
	return Hx("replace-url", value)
}

// HxRequest renders an hx-request="[value]" attribute.
//
// Configures various aspects of the request.
//
// https://htmx.org/attributes/hx-request/
func HxRequest(value string) nodx.Node {
	return Hx("request", value)
}

// HxSync renders an hx-sync="[value]" attribute.
//
// Control how requests made by different elements are synchronized.
//
// https://htmx.org/attributes/hx-sync/
func HxSync(value string) nodx.Node {
	return Hx("sync", value)
}

// HxValidate renders an hx-validate="[value]" attribute.
//
// Force elements to validate themselves before a request.
//
// https://htmx.org/attributes/hx-validate/
func HxValidate(value string) nodx.Node {
	return Hx("validate", value)
}

// HxVars renders an hx-vars="[value]" attribute.
//
// Adds values dynamically to the parameters to submit with the request (deprecated, please use hx-vals).
//
// https://htmx.org/attributes/hx-vars/
func HxVars(value string) nodx.Node {
	return Hx("vars", value)
}
