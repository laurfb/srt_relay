package auth

import "github.com/laurfb/srt_relay/stream"

type Authenticator interface {
	Authenticate(stream.StreamID) bool
}
