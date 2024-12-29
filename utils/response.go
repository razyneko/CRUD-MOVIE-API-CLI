// Helper functions for responses (e.g., setting headers)
// Helper for setting response headers.
package utils

import "net/http"

func SetJSONHeader(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
}
