package function

import (
	"context"
	"fmt"
	"net/http"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */

	fmt.Fprintf(res, " _______________________________________\n/ What a RUSH! We just deploy a Knative \\n\\ Function live on a Keynote!           /\n---------------------------------------\n\t\t\\   ^__^\n\t\t \\  (oo)\\_______\n\t\t\t (__)\\       )\\/\\n\t\t\t\t||----w |\n\t\t\t\t||     ||\n")
}
