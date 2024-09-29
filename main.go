package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "time"

    "github.com/chromedp/chromedp"
    "github.com/chromedp/chromedp/kb"
)

func main() {
    // Set Chrome options to run in headed mode (with UI)
    opts := chromedp.DefaultExecAllocatorOptions[:]
    opts = append(opts,
        chromedp.Flag("headless", true),      // Disable headless mode
        chromedp.Flag("disable-gpu", false),   // Enable GPU (optional)
        chromedp.Flag("start-maximized", true),// Start with a maximized window
    )

    // Create a new browser allocator with the options
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()

    // Create a new context for Chrome using the allocator
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    // Set a timeout for the entire test
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    // Variables to hold the results
    var result string

    // List of tasks to perform
    err := chromedp.Run(ctx,
        // Step 1: Navigate to Google
        chromedp.Navigate("https://www.google.com"),

        // Step 2: Wait for the search bar to be visible
        chromedp.WaitVisible(`//textarea[@name="q"]`),

        // Step 3: Input "Golang" in the search bar
        chromedp.SendKeys(`//textarea[@name="q"]`, "Golang"),

        // Step 4: Submit the form (press Enter)
        chromedp.SendKeys(`//textarea[@name="q"]`, kb.Enter),

        // Step 5: Wait for search results page to load
        chromedp.WaitVisible(`#search`),

        // Step 6: Extract the text of the first search result
        chromedp.Text(`#search`, &result),
    )

    if err != nil {
        log.Fatal(err)
    }

    // Check if the word "Golang" is present in the search result
    if strings.Contains(result, "Golang") {
        fmt.Println("Test Passed: 'Golang' found in search results")
    } else {
        fmt.Println("Test Failed: 'Golang' not found in search results")
    }
}
