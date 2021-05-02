package checker

import (
	"fmt"
	"github.com/tebeka/selenium"
	"os"
)

var wd = _GetWebDriver()

func _GetWebDriver() selenium.WebDriver {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
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
	//defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	//defer wd.Quit()

	return wd
}

func PayloadWasExecuted(host string, uri string) bool {
	if err := wd.Get(fmt.Sprintf("http://%s%s&bot=1", host, uri)); err != nil {
		panic(err)
	}

	_, err := wd.AlertText()
	if err != nil {
		return false
	} else {
		return true
	}
}
