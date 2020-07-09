package middleware

import (
	"context"
	"net/http"
	"strconv"
)

// Type to retrieve our context objects
type ContextID int

// The only ID we've defined
const ID ContextID = 0

// Updates context with the id then increments it
func SetID(start int64) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ID,
				strconv.FormatInt(start, 10))
			start++
			r = r.WithContext(ctx)
			next(w, r)
		}
	}
}

// Grabs an ID from a context if set
// otherwise it returns an empty string
func GetID(ctx context.Context) string {
	if val, ok := ctx.Value(ID).(string); ok {
		return val
	}
	return ""
}