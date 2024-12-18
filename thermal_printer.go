package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	printerPath := "/dev/usb/lp0" // Update the device path if necessary

	// Open printer device
	file, err := os.OpenFile(printerPath, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error: Unable to open printer device (%s): %v\n", printerPath, err)
		return
	}
	defer file.Close()

	// Commands
	init := []byte{0x1B, 0x40} // Initialize
	message := []byte(`Hello, Printer!
	Working with suddam,
	working with also haidar,
	also jewel,
	also kamrul
	thats all
	`) // Simple message
	cut := []byte{0x1D, 0x56, 0x42, 0x00} // Cut paper

	// Debug initialization
	fmt.Println("Sending Initialize Command:", init)
	_, err = file.Write(init)
	if err != nil {
		fmt.Println("Error writing Initialize Command:", err)
		return
	}
	time.Sleep(100 * time.Millisecond)

	// Debug message
	fmt.Println("Sending Message:", string(message))
	_, err = file.Write(message)
	if err != nil {
		fmt.Println("Error writing Message:", err)
		return
	}
	time.Sleep(100 * time.Millisecond)

	// Debug cut
	fmt.Println("Sending Cut Command:", cut)
	_, err = file.Write(cut)
	if err != nil {
		fmt.Println("Error writing Cut Command:", err)
		return
	}

	fmt.Println("Printing completed successfully!")
}
