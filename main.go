package main

import (
	"context"
	"os/signal"
	"syscall"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	connectDb()
	migrateDb()
	defer closeDb()

	newAdsblol()
	adsblolInstance.start()
	defer adsblolInstance.stop()

	newAdsbone()
	adsboneInstance.start()
	defer adsboneInstance.stop()

	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
}
