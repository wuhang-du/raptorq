package main

import (
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
	c.Start()
}
