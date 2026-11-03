package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
)

// Recover turns a panic in any handler into a 500 UNKNOWN ERROR instead of a dropped connection.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rec := recover()
			if rec == nil {
				return
			}

			if rec == http.ErrAbortHandler {
				panic(rec)
			}

			slog.Error("[Recover] panic: "+fmt.Sprint(rec), "stack", string(debug.Stack()))
			responses.Error(w, http.StatusInternalServerError, responses.UNKNOWN_ERROR)
		}()

		next.ServeHTTP(w, r)
	})
}
