package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set a timeout
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Navigate to the website
	err := chromedp.Run(ctx, navigateToWebsite("https://www.dezlearn.com/nested-iframes-example/"))
	if err != nil {
		log.Fatal(err)
	}
}

func navigateToWebsite(url string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second), // Wait for the website to load (optional)

		// Accessing nested iframes
		//chromedp.Frame("iframe1", chromedp.Frame("iframe2", chromedp.Frame("iframe3", chromedp.Click("#buttonID")))),
		// You can replace chromedp.Click("#buttonID") with chromedp.SetValue("#inputID", "text") to input text

	}
}
