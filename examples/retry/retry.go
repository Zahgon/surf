package main

import (
	"fmt"
)

func main() {
	// Example 1: Simple GET retry
	fmt.Println("=== Example 1: GET retry ===")
	simpleGetRetry()

	// Example 2: POST retry with body preservation
	fmt.Println("\n=== Example 2: POST retry with body ===")
	postRetryWithBody()

	// Example 3: Multipart retry with body preservation
	fmt.Println("\n=== Example 3: Multipart retry ===")
	multipartRetry()

	// Example 4: Retry honouring Retry-After header
	fmt.Println("\n=== Example 4: Retry-After ===")
	retryAfterScenario()
}

func simpleGetRetry() { _ = "STUB: not implemented"; return }

func postRetryWithBody() {
	_ = "STUB: not implemented"
	// Start local test server that fails first 2 requests
	return
}

// Wait for server to start

func multipartRetry() {
	_ = "STUB: not implemented"
	// Start local test server that fails first request
	return
}

// Wait for server to start

func retryAfterScenario() {
	_ = "STUB: not implemented"
	// Local test server: first attempt returns 429 with Retry-After: 2,
	// every subsequent attempt returns 200 OK. retryWait is 100ms, so the
	// Retry-After header is what actually extends the pause to ~2s.
	return
}

// Wait for server to start

// Upper bound for the entire retry cycle (including a long Retry-After)
// is the request context deadline. Builder.Timeout would only bound a
// single cli.Do — it does not interrupt the inter-attempt sleep.
