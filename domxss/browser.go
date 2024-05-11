package main

import (
	"context"
	// "fmt"
	"log"

	"github.com/chromedp/chromedp"
)
 

func browse(url string)(string){
	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.ProxyServer("http://127.0.0.1:8081"), chromedp.Headless,
	)

	ctx, cancel = chromedp.NewContext(ctx)
	
	defer cancel()
	
	var html string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &html),
	})
	if err != nil {
		log.Fatal(err)
	}
	
	return html
	
}