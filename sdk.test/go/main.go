package main

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
  err := sentry.Init(sentry.ClientOptions{
    Dsn: "http://b10af59299974ef1659c10b54f54ce62@localhost:9000/2",
    // Set TracesSampleRate to 1.0 to capture 100%
    // of transactions for performance monitoring.
    // We recommend adjusting this value in production,
    TracesSampleRate: 1.0,
  })
  if err != nil {
    log.Fatalf("sentry.Init: %s", err)
  }

  // Flush buffered events before the program terminates.
  defer sentry.Flush(2 * time.Second)

  sentry.CaptureMessage("It works!")
}
