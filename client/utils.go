package client

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

// openBrowser opens the specified URL in the default web browser.
// openBrowser opens the specified URL in the default web browser of the user's operating system.
// It supports Windows, macOS (Darwin), and Linux platforms. If the platform is unsupported,
// a warning is logged.
//
// Parameters:
//   - url: The URL to be opened in the web browser.
//
// Example usage:
//
//	openBrowser("https://example.com")
func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	case "linux":
		cmd = "xdg-open"
	default:
		logger.Warn("Unsupported platform", "platform", runtime.GOOS)
		return
	}

	args = append(args, url)
	if err := exec.Command(cmd, args...).Start(); err != nil {
		logger.Error("Error opening browser", "error", err)
	}
}

// splitApplianceURL splits the appliance URL into protocol, hostname, domain name, and port.
// Returns an error if the URL format is invalid.
// splitApplianceURL splits the appliance URL into its components: protocol, hostname, domain name, and port.
// It returns the protocol (e.g., "http", "https"), the hostname, the domain name, and the port.
// If the URL is invalid or the protocol is unknown, it returns an error.
//
// Returns:
//   - protocol: The URL scheme (e.g., "http", "https").
//   - hostname: The hostname part of the URL.
//   - domainName: The domain name part of the URL.
//   - port: The port number. If no port is specified in the URL, it returns the default port for the protocol.
//   - err: An error if the URL is invalid or the protocol is unknown.
func splitApplianceURL(applianceUrl string) (protocol, hostname, domainName, port string, err error) {
	parsedURL, err := url.Parse(applianceUrl)
	if err != nil {
		return "", "", "", "", fmt.Errorf("invalid URL format: %v", err)
	}

	protocol = parsedURL.Scheme
	host := parsedURL.Hostname()
	port = parsedURL.Port()

	// Split the host into parts to extract the domain name
	hostParts := strings.Split(host, ".")
	if len(hostParts) < 2 {
		hostname = hostParts[0]
		domainName = ""
	} else {
		hostname = hostParts[0]
		domainName = strings.Join(hostParts[1:], ".")
	}

	// If no port is specified, use the default port for the protocol
	if port == "" {
		switch protocol {
		case "http":
			port = "80"
		case "https":
			port = "443"
		default:
			return "", "", "", "", fmt.Errorf("unknown protocol: %s", protocol)
		}
	}

	return protocol, hostname, domainName, port, nil
}
