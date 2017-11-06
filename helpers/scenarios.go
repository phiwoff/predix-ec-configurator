package helpers

import "net/url"

func GetDebugString(value string) string {
	debugStr := ""
	if value == "yes" {
		debugStr = "-dbg"
	}
	return debugStr
}

func GetProxyString(proxyURL string) string {
	proxyString := ""
	if proxyURL != "" {
		_, err := url.Parse(proxyURL)
		if err == nil {
			proxyString = "-pxy " + proxyURL
		}
	}
	return proxyString
}
