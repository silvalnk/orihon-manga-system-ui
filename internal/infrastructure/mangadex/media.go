package mangadex

import (
	"net"
	"net/url"
	"strings"
)

func Allowed(raw string) bool {
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" || parsed.User != nil {
		return false
	}
	host := strings.ToLower(parsed.Hostname())
	if host == "localhost" || strings.HasSuffix(host, ".localhost") || net.ParseIP(host) != nil {
		return false
	}
	path := parsed.EscapedPath()
	if strings.HasSuffix(host, "mangadex.org") || strings.HasSuffix(host, "mangadex.network") {
		return strings.Contains(path, "/covers/") || strings.Contains(path, "/data/") || strings.Contains(path, "/data-saver/")
	}
	return strings.Contains(path, "/data-saver/") || strings.Contains(path, "/data/")
}
