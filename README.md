# nodxgo-htmx

<a href="https://pkg.go.dev/github.com/nodxdev/nodxgo-htmx">
  <img src="https://pkg.go.dev/badge/github.com/nodxdev/nodxgo-htmx" alt="Go Reference"/>
</a>
<a href="https://goreportcard.com/report/nodxdev/nodxgo-htmx">
  <img src="https://goreportcard.com/badge/nodxdev/nodxgo-htmx" alt="Go Report Card"/>
</a>
<a href="https://github.com/nodxdev/nodxgo-htmx/releases/latest">
  <img src="https://img.shields.io/github/release/nodxdev/nodxgo-htmx.svg" alt="Release Version"/>
</a>
<a href="LICENSE">
  <img src="https://img.shields.io/github/license/nodxdev/nodxgo-htmx.svg" alt="License"/>
</a>
<a href="https://github.com/nodxdev/nodxgo-htmx">
  <img src="https://img.shields.io/github/stars/nodxdev/nodxgo-htmx?style=flat&label=github+stars"/>
</a>

## Overview

nodxgo-htmx provides Go-based integration for [HTMX](https://htmx.org) within
the [NodX Go](https://github.com/nodxdev/nodxgo) template engine. This package
enables the generation of dynamic HTML attributes and server-side helpers for
handling HTMX requests efficiently.

## Installation

To install the package, run:

```sh
# Go 1.22 or later is required
go get github.com/nodxdev/nodxgo-htmx
```

## Usage

You can see all the included functions in the
[Go reference](https://pkg.go.dev/github.com/nodxdev/nodxgo-htmx) or with your
editor's auto-completion.

### Example: Using HTMX Attributes

#### `hx-get` Attribute

```go
node := nodx.Div(
  htmx.HxGet("/api/data"),
)
fmt.Println(node)
```

**Output:**

```html
<div hx-get="/api/data"></div>
```

#### `hx-trigger` Attribute

```go
node := nodx.Button(
  htmx.HxTrigger("click"),
  nodx.Text("Click Me"),
)
fmt.Println(node)
```

**Output:**

```html
<button hx-trigger="click">Click Me</button>
```

### Example: Using Server Helpers

#### Detecting an HTMX Request

```go
func handler(w http.ResponseWriter, r *http.Request) {
  if htmx.ServerGetIsHtmxRequest(r.Header) {
    fmt.Println("HTMX request detected")
  }
}
```

#### Setting a Redirect via Response Header

```go
func handler(w http.ResponseWriter, r *http.Request) {
  htmx.ServerSetRedirect(w.Header(), "/new-url")
  w.WriteHeader(http.StatusOK)
}
```

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to
improve the project.

---

For more details, check out the [HTMX documentation](https://htmx.org) and
[NodX Go](https://github.com/nodxdev/nodxgo).
