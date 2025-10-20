package dev

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/wailsapp/wails/v2/cmd/wails/internal/logutils"
	"golang.org/x/mod/semver"
)

// stdoutScanner acts as a stdout target that will scan the incoming
// data to find out the vite server url
type stdoutScanner struct {
	ViteServerURLChan  chan string
	ViteServerVersionC chan string
	versionDetected    bool
	urlDetected        bool
}

// NewStdoutScanner creates a new stdoutScanner
func NewStdoutScanner() *stdoutScanner {
	return &stdoutScanner{
		ViteServerURLChan:  make(chan string, 2),
		ViteServerVersionC: make(chan string, 2),
	}
}

// Write bytes to the scanner. Will copy the bytes to stdout
func (s *stdoutScanner) Write(data []byte) (n int, err error) {
	input := stripansi.Strip(string(data))
	if !s.versionDetected {
		v, err := detectViteVersion(input)
		if v != "" || err != nil {
			if err != nil {
				logutils.LogRed("ViteStdoutScanner: %s", err)
				v = "v0.0.0"
			}
			s.ViteServerVersionC <- v
			s.versionDetected = true
		}
	}

	match := strings.Index(input, "Local:")
	if match != -1 {
		sc := bufio.NewScanner(strings.NewReader(input))
		for sc.Scan() {
			line := sc.Text()
			index := strings.Index(line, "Local:")
			if index == -1 || len(line) < 7 {
				continue
			}
			viteServerURL := strings.TrimSpace(line[index+6:])
			logutils.LogGreen("Vite Server URL: %s", viteServerURL)
			_, err := url.Parse(viteServerURL)
			if err != nil {
				logutils.LogRed("%s", err.Error())
			} else {
				if !s.urlDetected {
					s.urlDetected = true
					s.ViteServerURLChan <- viteServerURL
				}
			}
		}
	}

	// Fallback: detect generic http(s) URLs (e.g., Flutter web-server output)
	if !s.urlDetected {
		// Look for first http(s) URL in the output
		re := regexp.MustCompile(`https?://[^\s]+`)
		urls := re.FindAllString(input, -1)
		for _, u := range urls {
			// Prefer localhost/127.0.0.1/0.0.0.0
			parsed, perr := url.Parse(u)
			if perr != nil {
				continue
			}
			host := parsed.Hostname()
			if host == "localhost" || host == "127.0.0.1" || host == "0.0.0.0" {
				if !s.urlDetected {
					s.urlDetected = true
					logutils.LogGreen("Detected DevServer URL: %s", u)
					s.ViteServerURLChan <- u
				}
				break
			}
		}
	}
	return os.Stdout.Write(data)
}

func detectViteVersion(line string) (string, error) {
	s := strings.Split(strings.TrimSpace(line), " ")
	if strings.ToLower(s[0]) != "vite" {
		return "", nil
	}

	if len(line) < 2 {
		return "", fmt.Errorf("unable to parse vite version")
	}

	v := s[1]
	if !semver.IsValid(v) {
		return "", fmt.Errorf("%s is not a valid vite version string", v)
	}

	return v, nil
}
