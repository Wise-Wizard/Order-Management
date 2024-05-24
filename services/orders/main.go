package main

import "log"

func main() {
	gRPCServer := NewGRPCServcer(":9000")
	err := gRPCServer.Run()
	if err != nil {
		log.Fatalf("Failed to run gRPC server: %v", err)
	}
}
