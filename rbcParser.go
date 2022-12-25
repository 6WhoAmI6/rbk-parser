package main

import (
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func rbcParse() error {
	/*launcher.NewBrowser().MustGet()
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()*/

	l := launcher.New().
		Headless(false).
		Devtools(true)

	defer l.Cleanup() // remove launcher.FlagUserDataDir

	url := l.MustLaunch()

	// Trace shows verbose debug information for each action executed
	// SlowMotion is a debug related function that waits 2 seconds between
	// each action, making it easier to inspect what your code is doing.
	browser := rod.New().
		ControlURL(url).
		Trace(true).
		SlowMotion(2 * time.Second).
		MustConnect()

	// ServeMonitor plays screenshots of each tab. This feature is extremely
	// useful when debugging with headless mode.
	// You can also enable it with flag "-rod=monitor"
	launcher.Open(browser.ServeMonitor(""))

	defer browser.MustClose()

	//page := stealth.MustPage(browser)

	//page.MustNavigate("https://www.rbc.ru/").MustWaitLoad()

	page := browser.MustPage("https://www.rbc.ru/")

	page.MustScreenshotFullPage("screen.png")
	log.Println(page.MustHTML())

	return nil
}
