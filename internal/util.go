package internal

import "context"

// IsDone returns true if context is canceled.
func IsDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
