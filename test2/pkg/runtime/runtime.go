package runtime

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// WaitSignal method is a runtime utility function
// that blocks the runtime until
// a signal is received
func WaitSignal() os.Signal {
	return <-getSignalChan()
}

func getSignalChan() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	return sig
}

// RunUntilSignal starts run function and continues
// until run function returns error or until signal
// is received from system, when signal is received
// stop function is triggered with timeout.
func RunUntilSignal(
	run func() error,
	stop func(context.Context) error,
	timeout time.Duration,
) error {
	sigChan := getSignalChan()
	errSig := make(chan error)

	go func() {
		errSig <- run()
	}()

	select {
	case err := <-errSig:
		return err
	case sig := <-sigChan:
		ctx, cancel := context.WithTimeout(
			context.Background(),
			timeout,
		)
		defer cancel()

		log.Printf("received signal: %s\n", sig.String())
		if stop != nil {
			err := stop(ctx)
			if err != nil {
				log.Printf("could not stop server: %v\n", err)
			}
		}
	}

	return nil
}
