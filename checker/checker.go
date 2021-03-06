package checker

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net/http"
	"os"
)

var wd = _GetWebDriver()

func _GetWebDriver() selenium.WebDriver {
	const (
		seleniumPath    = "/home/n30/go/src/github.com/tebeka/selenium/vendor/selenium-server.jar"
		geckoDriverPath = "/home/n30/go/src/github.com/tebeka/selenium/vendor/geckodriver"
		port            = 4000
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	_, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	return wd
}

func PayloadWasExecuted(r *http.Request, anchor string) bool {
	host := r.Host
	path := r.URL.Path
	params := r.URL.RawQuery
	if params == "" {
		params = "bot=1"
	} else {
		params += "&bot=1"
	}

	wd.DeleteAllCookies()
	wd.Get(fmt.Sprintf("http://%s%s?%s%s", host, path, params, anchor))
	_, err := wd.AlertText()

	if err != nil {
		wd.Refresh()
		_, err = wd.AlertText()
	}
	if err != nil {
		return false
	}
	return true
}
