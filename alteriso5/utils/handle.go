package utils

import (
	"os"
	"os/signal"
)

func SignalChan(sigs ...os.Signal) *chan os.Signal {
	sig := make(chan os.Signal, 1)
	if len(sigs) == 0 {
		sigs = []os.Signal{os.Interrupt}
	}
	signal.Notify(sig, sigs...)
	return &sig
}

func OnSignal(on func(s os.Signal), sigs ...os.Signal) {
	go func() {
		sig := SignalChan(sigs...)
		s := <-*sig
		on(s)
	}()
}
