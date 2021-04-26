package main

import (
	"time"
	"wuhang-du/raptorq/consumer"
	"wuhang-du/raptorq/microserver"
	"wuhang-du/raptorq/raqserver"
	"wuhang-du/raptorq/tracker"
)

func main() {
	t := tracker.NewTracker()
	_ = raqserver.NewRaqServer(t)
	_ = microserver.NewMicroServer(t)
	c := consumer.NewConsumer(t)
	c.Start(0)

	time.Sleep(10 * time.Second)
}
