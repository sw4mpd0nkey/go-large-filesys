package p2p

const (
	IncomingMessage = 0x1
	IncomingStream  = 0x2
)

// RPC holds any arb data being sent over wire
// between nodes
type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
