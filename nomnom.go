package main

import (
	"fmt"
	"time"
)

func main() {
	memAlloc := 528
	// Allocate approximately 128MB of memory
	// 128MB = 128 * 1024 * 1024 bytes
	memSize := memAlloc * 1024 * 1024
	
	// Create a byte slice of the required size
	memory := make([]byte, memSize)
	
	// Write some data to the memory to ensure it's actually allocated
	// (Go might optimize away unused allocations)
	for i := 0; i < memSize; i += 1024 {
		memory[i] = 1
	}
	
	fmt.Printf("Allocated %d MB of memory\n", memAlloc)
	
	// Keep the program running until interrupted
	fmt.Println("Press Ctrl+C to exit")
	for {
		time.Sleep(1 * time.Second)
	}
}
