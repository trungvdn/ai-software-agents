package oauth

import "os/exec"

type WindowsBrowser struct{}

func (WindowsBrowser) Open(url string) error {
	return exec.Command(
		"rundll32",
		"url.dll,FileProtocolHandler",
		url,
	).Start()
}
