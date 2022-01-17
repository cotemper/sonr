package motor

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/core/node"
	v1 "github.com/sonr-io/sonr/motor/v1"
	"google.golang.org/protobuf/proto"
)

// Start starts the host, node, and rpc service.
func Start(reqBuf []byte) {
	// Unmarshal request
	req := &v1.InitializeRequest{}
	if err := proto.Unmarshal(reqBuf, req); err != nil {
		golog.Warn("%s - Failed to unmarshal InitializeRequest. Using defaults...", err)
	}

	// Start the app
	node.NewHighway(context.Background(), node.WithLogLevel(node.DebugLevel))
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	node.Pause()
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	node.Resume()
}

// Stop closes the host, node, and rpc service.
func Stop() {
	node.Exit(0)
}