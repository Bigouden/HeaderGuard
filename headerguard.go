// Package headerguard provides a middleware to allow access based on HTTP header values.
package headerguard

import (
	"context"
	"net/http"
	"strings"
)

// Config defines the plugin configuration.
type Config struct {
	Header string   `json:"header,omitempty"`
	Allow      []string `json:"allow,omitempty"`
	Separator  string   `json:"separator,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Header: "X-Auth-Request-Groups",
		Allow:      []string{},
		Separator:  "|",
	}
}

// HeaderGuard is a middleware that checks if a request's header matches allowed values.
type HeaderGuard struct {
	next       http.Handler
	header string
	allow      map[string]struct{}
	separator  string
	name       string
}

// New creates a new instance of the HeaderGuard middleware.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	allowedSet := make(map[string]struct{})
	for _, v := range config.Allow {
		allowedSet[strings.TrimSpace(v)] = struct{}{}
	}

	return &HeaderGuard{
		next:       next,
		header: config.Header,
		allow:      allowedSet,
		separator:  config.Separator,
		name:       name,
	}, nil
}

// ServeHTTP implements the middleware logic.
func (hg *HeaderGuard) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	headerValue := req.Header.Get(hg.header)

	if headerValue == "" {
		http.Error(rw, "Forbidden: missing header", http.StatusForbidden)
		return
	}

	values := strings.Split(headerValue, hg.separator)
	for _, v := range values {
		trimmed := strings.TrimSpace(v)
		if _, ok := hg.allow[trimmed]; ok {
			hg.next.ServeHTTP(rw, req)
			return
		}
	}

	http.Error(rw, "Forbidden", http.StatusForbidden)
}
