package htmx

import (
	"net/http"
	"testing"
)

/*******************/
/* Request Headers */
/*******************/

func TestServerGetIsBoosted(t *testing.T) {
	headers := http.Header{}
	if ServerGetIsBoosted(headers) {
		t.Error("Expected false when HX-Boosted header is not set")
	}
	headers.Set("HX-Boosted", "true")
	if !ServerGetIsBoosted(headers) {
		t.Error("Expected true when HX-Boosted header is set")
	}
}

func TestServerGetCurrentURL(t *testing.T) {
	headers := http.Header{}
	expected := "http://example.com"
	headers.Set("HX-Current-URL", expected)
	if got := ServerGetCurrentURL(headers); got != expected {
		t.Errorf("ServerGetCurrentURL: expected %q, got %q", expected, got)
	}
}

func TestServerGetIsHistoryRestoreRequest(t *testing.T) {
	headers := http.Header{}
	headers.Set("HX-History-Restore-Request", "true")
	if !ServerGetIsHistoryRestoreRequest(headers) {
		t.Error("Expected true when HX-History-Restore-Request is 'true'")
	}
	headers.Set("HX-History-Restore-Request", "false")
	if ServerGetIsHistoryRestoreRequest(headers) {
		t.Error("Expected false when HX-History-Restore-Request is not 'true'")
	}
}

func TestServerGetPrompt(t *testing.T) {
	headers := http.Header{}
	expected := "Are you sure?"
	headers.Set("HX-Prompt", expected)
	if got := ServerGetPrompt(headers); got != expected {
		t.Errorf("ServerGetPrompt: expected %q, got %q", expected, got)
	}
}

func TestServerGetIsHtmxRequest(t *testing.T) {
	headers := http.Header{}
	headers.Set("HX-Request", "true")
	if !ServerGetIsHtmxRequest(headers) {
		t.Error("Expected true when HX-Request is 'true'")
	}
	headers.Set("HX-Request", "false")
	if ServerGetIsHtmxRequest(headers) {
		t.Error("Expected false when HX-Request is not 'true'")
	}
}

func TestServerGetTarget(t *testing.T) {
	headers := http.Header{}
	expected := "#content"
	headers.Set("HX-Target", expected)
	if got := ServerGetTarget(headers); got != expected {
		t.Errorf("ServerGetTarget: expected %q, got %q", expected, got)
	}
}

func TestServerGetTriggerName(t *testing.T) {
	headers := http.Header{}
	expected := "btnSubmit"
	headers.Set("HX-Trigger-Name", expected)
	if got := ServerGetTriggerName(headers); got != expected {
		t.Errorf("ServerGetTriggerName: expected %q, got %q", expected, got)
	}
}

func TestServerGetTrigger(t *testing.T) {
	headers := http.Header{}
	expected := "btnSubmit"
	headers.Set("HX-Trigger", expected)
	if got := ServerGetTrigger(headers); got != expected {
		t.Errorf("ServerGetTrigger: expected %q, got %q", expected, got)
	}
}

/********************/
/* Response Headers */
/********************/

func TestServerSetLocation(t *testing.T) {
	headers := http.Header{}
	expected := "/new-location"
	ServerSetLocation(headers, expected)
	if got := headers.Get("HX-Location"); got != expected {
		t.Errorf("ServerSetLocation: expected %q, got %q", expected, got)
	}
}

func TestServerSetPushURL(t *testing.T) {
	headers := http.Header{}
	expected := "/pushed-url"
	ServerSetPushURL(headers, expected)
	if got := headers.Get("HX-Push-Url"); got != expected {
		t.Errorf("ServerSetPushURL: expected %q, got %q", expected, got)
	}
}

func TestServerSetRedirect(t *testing.T) {
	headers := http.Header{}
	expected := "/redirect-url"
	ServerSetRedirect(headers, expected)
	if got := headers.Get("HX-Redirect"); got != expected {
		t.Errorf("ServerSetRedirect: expected %q, got %q", expected, got)
	}
}

func TestServerSetRefresh(t *testing.T) {
	headers := http.Header{}
	expected := "true"
	ServerSetRefresh(headers, expected)
	if got := headers.Get("HX-Refresh"); got != expected {
		t.Errorf("ServerSetRefresh: expected %q, got %q", expected, got)
	}
}

func TestServerSetReplaceURL(t *testing.T) {
	headers := http.Header{}
	expected := "/replace-url"
	ServerSetReplaceURL(headers, expected)
	if got := headers.Get("HX-Replace-Url"); got != expected {
		t.Errorf("ServerSetReplaceURL: expected %q, got %q", expected, got)
	}
}

func TestServerSetReswap(t *testing.T) {
	headers := http.Header{}
	expected := "outerHTML"
	ServerSetReswap(headers, expected)
	if got := headers.Get("HX-Reswap"); got != expected {
		t.Errorf("ServerSetReswap: expected %q, got %q", expected, got)
	}
}

func TestServerSetRetarget(t *testing.T) {
	headers := http.Header{}
	expected := "#target"
	ServerSetRetarget(headers, expected)
	if got := headers.Get("HX-Retarget"); got != expected {
		t.Errorf("ServerSetRetarget: expected %q, got %q", expected, got)
	}
}

func TestServerSetReselect(t *testing.T) {
	headers := http.Header{}
	expected := ".selector"
	ServerSetReselect(headers, expected)
	if got := headers.Get("HX-Reselect"); got != expected {
		t.Errorf("ServerSetReselect: expected %q, got %q", expected, got)
	}
}

func TestServerSetTrigger(t *testing.T) {
	headers := http.Header{}
	expected := "eventTrigger"
	ServerSetTrigger(headers, expected)
	if got := headers.Get("HX-Trigger"); got != expected {
		t.Errorf("ServerSetTrigger: expected %q, got %q", expected, got)
	}
}

func TestServerSetTriggerAfterSettle(t *testing.T) {
	headers := http.Header{}
	expected := "afterSettle"
	ServerSetTriggerAfterSettle(headers, expected)
	if got := headers.Get("HX-Trigger-After-Settle"); got != expected {
		t.Errorf("ServerSetTriggerAfterSettle: expected %q, got %q", expected, got)
	}
}

func TestServerSetTriggerAfterSwap(t *testing.T) {
	headers := http.Header{}
	expected := "afterSwap"
	ServerSetTriggerAfterSwap(headers, expected)
	if got := headers.Get("HX-Trigger-After-Swap"); got != expected {
		t.Errorf("ServerSetTriggerAfterSwap: expected %q, got %q", expected, got)
	}
}
