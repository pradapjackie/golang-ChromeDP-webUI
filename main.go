package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func main() {
	// Check environment variable for headless mode (default: true for headless)
	headless := os.Getenv("HEADLESS") != "false"

	// Chrome options
	opts := chromedp.DefaultExecAllocatorOptions[:]
	opts = append(opts,
		chromedp.Flag("headless", headless), // Toggle headless mode based on env
		chromedp.Flag("disable-gpu", headless), // Disable GPU in headless mode
		chromedp.Flag("no-sandbox", true), // Disable sandbox
	)

	// Create Chrome allocator and context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Set a timeout for the test
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Variables to hold test result
	var result string

	// Perform the UI automation tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.google.com"),
		chromedp.WaitVisible(`//textarea[@name="q"]`),
		chromedp.SendKeys(`//textarea[@name="q"]`, "Golang"),
		chromedp.SendKeys(`//textarea[@name="q"]`, kb.Enter),
		chromedp.WaitVisible(`#search`),
		chromedp.Text(`#search`, &result),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the result contains "Golang"
	if strings.Contains(result, "Golang") {
		fmt.Println("Test Passed: 'Golang' found in search results")
	} else {
		fmt.Println("Test Failed: 'Golang' not found in search results")
	}
}
