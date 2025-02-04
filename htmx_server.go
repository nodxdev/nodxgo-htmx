package htmx

import "net/http"

/*******************/
/* Request Headers */
/*******************/

// ServerGetIsBoosted returns true if the "HX-Boosted" request header is present.
//
// Indicates that the request is via an element using hx-boost.
func ServerGetIsBoosted(headers http.Header) bool {
	return headers.Get("HX-Boosted") != ""
}

// ServerGetCurrentURL returns the value of the "HX-Current-URL" request header.
//
// Contains the current URL of the browser.
func ServerGetCurrentURL(headers http.Header) string {
	return headers.Get("HX-Current-URL")
}

// ServerGetIsHistoryRestoreRequest returns true if the "HX-History-Restore-Request" header equals "true".
//
// Indicates that the request is for history restoration after a miss in the local history cache.
func ServerGetIsHistoryRestoreRequest(headers http.Header) bool {
	return headers.Get("HX-History-Restore-Request") == "true"
}

// ServerGetPrompt returns the value of the "HX-Prompt" request header.
//
// Contains the user response to an hx-prompt.
func ServerGetPrompt(headers http.Header) string {
	return headers.Get("HX-Prompt")
}

// ServerGetIsHtmxRequest returns true if the "HX-Request" header equals "true".
//
// Indicates that the request is made via HTMX.
func ServerGetIsHtmxRequest(headers http.Header) bool {
	return headers.Get("HX-Request") == "true"
}

// ServerGetTarget returns the value of the "HX-Target" request header.
//
// Contains the id of the target element, if it exists.
func ServerGetTarget(headers http.Header) string {
	return headers.Get("HX-Target")
}

// ServerGetTriggerName returns the value of the "HX-Trigger-Name" request header.
//
// Contains the name of the triggered element, if it exists.
func ServerGetTriggerName(headers http.Header) string {
	return headers.Get("HX-Trigger-Name")
}

// ServerGetTrigger returns the value of the "HX-Trigger" request header.
//
// Contains the id of the triggered element, if it exists.
func ServerGetTrigger(headers http.Header) string {
	return headers.Get("HX-Trigger")
}

/********************/
/* Response Headers */
/********************/

// ServerSetLocation sets the "HX-Location" response header to the given value.
//
// Allows a client-side redirect without a full page reload.
//
// https://htmx.org/headers/hx-location/
func ServerSetLocation(headers http.Header, value string) {
	headers.Set("HX-Location", value)
}

// ServerSetPushURL sets the "HX-Push-Url" response header to the given value.
//
// Pushes a new URL into the browser's history stack.
//
// https://htmx.org/headers/hx-push-url/
func ServerSetPushURL(headers http.Header, value string) {
	headers.Set("HX-Push-Url", value)
}

// ServerSetRedirect sets the "HX-Redirect" response header to the given value.
//
// Can be used to perform a client-side redirect to a new location.
//
// https://htmx.org/headers/hx-redirect/
func ServerSetRedirect(headers http.Header, value string) {
	headers.Set("HX-Redirect", value)
}

// ServerSetRefresh sets the "HX-Refresh" response header to the given value.
//
// If set to "true", the client will perform a full refresh of the page.
func ServerSetRefresh(headers http.Header, value string) {
	headers.Set("HX-Refresh", value)
}

// ServerSetReplaceURL sets the "HX-Replace-Url" response header to the given value.
//
// Replaces the current URL in the browser's location bar.
//
// https://htmx.org/headers/hx-replace-url/
func ServerSetReplaceURL(headers http.Header, value string) {
	headers.Set("HX-Replace-Url", value)
}

// ServerSetReswap sets the "HX-Reswap" response header to the given value.
//
// Specifies how the response will be swapped.
func ServerSetReswap(headers http.Header, value string) {
	headers.Set("HX-Reswap", value)
}

// ServerSetRetarget sets the "HX-Retarget" response header to the given value.
//
// Updates the target of the content update to a different element on the page.
func ServerSetRetarget(headers http.Header, value string) {
	headers.Set("HX-Retarget", value)
}

// ServerSetReselect sets the "HX-Reselect" response header to the given value.
//
// Specifies which part of the response is used for swapping, overriding any existing hx-select.
func ServerSetReselect(headers http.Header, value string) {
	headers.Set("HX-Reselect", value)
}

// ServerSetTrigger sets the "HX-Trigger" response header to the given value.
//
// Allows triggering client-side events.
//
// https://htmx.org/headers/hx-trigger/
func ServerSetTrigger(headers http.Header, value string) {
	headers.Set("HX-Trigger", value)
}

// ServerSetTriggerAfterSettle sets the "HX-Trigger-After-Settle" response header to the given value.
//
// Triggers client-side events after the settle step.
//
// https://htmx.org/headers/hx-trigger/
func ServerSetTriggerAfterSettle(headers http.Header, value string) {
	headers.Set("HX-Trigger-After-Settle", value)
}

// ServerSetTriggerAfterSwap sets the "HX-Trigger-After-Swap" response header to the given value.
//
// Triggers client-side events after the swap step.
//
// https://htmx.org/headers/hx-trigger/
func ServerSetTriggerAfterSwap(headers http.Header, value string) {
	headers.Set("HX-Trigger-After-Swap", value)
}
