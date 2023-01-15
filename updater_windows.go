package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/abemedia/go-winsparkle"
)

func init() {
	winsparkle.SetAppcastURL("https://raw.githubusercontent.com/abemedia/appcast-example/dist/sparkle/sparkle.xml")
	winsparkle.SetAppDetails("AppCast", "Example", version)
	winsparkle.SetAutomaticCheckForUpdates(true)

	if err := winsparkle.SetDSAPubPEM(dsaPublicKey); err != nil {
		panic(err)
	}

	winsparkle.Init()
	defer winsparkle.Cleanup()

	wait := make(chan struct{})
	winsparkle.SetUpdateCancelledCallback(func() {
		fmt.Println("Update cancelled.")
		close(wait)
	})
	winsparkle.SetDidNotFindUpdateCallback(func() {
		fmt.Println("No updates found.")
		close(wait)
	})
	winsparkle.CheckUpdateWithUIAndInstall()

	select {
	case <-wait:
	case <-time.After(time.Minute):
	}
}

//go:embed keys/dsa_public_key
var dsaPublicKey string
