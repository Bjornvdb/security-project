package session

import "github.com/gofiber/fiber/v2/middleware/session"

var (
	// Sessions here
	Sessions *session.Store
)

// MakeSession Create sessions
func MakeSession() {
	Sessions = session.New()
}
