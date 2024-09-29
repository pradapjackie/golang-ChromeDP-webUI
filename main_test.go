
package main

import (
    "context"
    "strings"
    "testing"
    "time"

    "github.com/chromedp/chromedp"
    "github.com/chromedp/chromedp/kb"
)

// TestGoogleSearch_Golang checks that the search results contain "Golang"
func TestGoogleSearch_Golang(t *testing.T) {
    opts := chromedp.DefaultExecAllocatorOptions[:]
    opts = append(opts,
        chromedp.Flag("headless", false),
        chromedp.Flag("disable-gpu", false),
        chromedp.Flag("start-maximized", true),
    )
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    var result string

    err := chromedp.Run(ctx,
        chromedp.Navigate("https://www.google.com"),
        chromedp.WaitVisible(`//textarea[@name="q"]`),
        chromedp.SendKeys(`//textarea[@name="q"]`, "Golang"),
        chromedp.SendKeys(`//textarea[@name="q"]`, kb.Enter),
        chromedp.WaitVisible(`#search`),
        chromedp.Text(`#search`, &result),
    )

    if err != nil {
        t.Fatalf("Test Failed: %v", err)
    }

    if strings.Contains(result, "Golang") {
        t.Log("Test Passed: 'Golang' found in search results")
    } else {
        t.Errorf("Test Failed: 'Golang' not found in search results")
    }
}

// TestGoogleSearch_SearchBarVisible checks that the search bar is visible
func TestGoogleSearch_SearchBarVisible(t *testing.T) {
    opts := chromedp.DefaultExecAllocatorOptions[:]
    opts = append(opts,
        chromedp.Flag("headless", false),
        chromedp.Flag("disable-gpu", false),
        chromedp.Flag("start-maximized", true),
    )
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    err := chromedp.Run(ctx,
        chromedp.Navigate("https://www.google.com"),
        chromedp.WaitVisible(`//textarea[@name="q"]`),
    )

    if err != nil {
        t.Errorf("Test Failed: Search bar not visible - %v", err)
    } else {
        t.Log("Test Passed: Search bar is visible")
    }
}

// TestGoogleSearch_EmptyQuery checks if there are results when the search query is empty
func TestGoogleSearch_EmptyQuery(t *testing.T) {
    opts := chromedp.DefaultExecAllocatorOptions[:]
    opts = append(opts,
        chromedp.Flag("headless", false),
        chromedp.Flag("disable-gpu", false),
        chromedp.Flag("start-maximized", true),
    )
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    var result string

    err := chromedp.Run(ctx,
        chromedp.Navigate("https://www.google.com"),
        chromedp.WaitVisible(`//textarea[@name="q"]`),
        chromedp.SendKeys(`//textarea[@name="q"]`, ""),  // Empty query
        chromedp.SendKeys(`//textarea[@name="q"]`, kb.Enter),
        chromedp.WaitVisible(`#search`, chromedp.ByID),
        chromedp.Text(`#search`, &result),
    )

    if err != nil {
        t.Fatalf("Test Failed: %v", err)
    }

    if result == "" {
        t.Log("Test Passed: No search results for empty query")
    } else {
        t.Errorf("Test Failed: Unexpected results for empty query")
    }
}
