package main_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/chromedp/chromedp"
)

func writeHTML(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, strings.TrimSpace(content))
	})
}

func ExampleTitle() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`
<head>
	<title>fancy website title</title>
</head>
<body>
	<div id="content"></div>
</body>
	`))
	defer ts.Close()

	var title string
	// fmt.Println("URL: ", ts.URL)
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Title(&title),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(title)

	// Output:
	// fancy website title
}
