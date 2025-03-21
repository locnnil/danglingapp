package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Container started. Waiting for termination signals...")

	// Create a channel to receive OS signals
	signalChan := make(chan os.Signal, 1)

	// Notify the channel on SIGINT and SIGTERM (but don't handle SIGKILL since it can't be caught)
	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGSEGV,
		// syscall.SIGKILL,
	)

	go func() {
		for sig := range signalChan {
			msg := fmt.Sprintf("Received signal: %s, but ignoring it.\n", sig)
			fmt.Printf("%s", msg)
			log.Printf("%s", msg)
			time.Sleep(30 * time.Second)
		}
	}()

	// Block forever (or until SIGKILL is received, which cannot be caught)
	select {}
}
