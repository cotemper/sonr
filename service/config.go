package service

import (
	"github.com/sonr-io/core/common"
	v1 "github.com/sonr-io/core/service/v1"
)

// Option is a function that can be used to configure a ServiceConfig.
type Option func(*v1.ServiceConfig)

// WithDescription is a Service Option that sets the service description.
func WithDescription(d string) Option {
	return func(c *v1.ServiceConfig) {
		c.Description = d
	}
}

// WithOwner is a Service Option that sets the service owner with a Did
// and a public key.
func WithOwner(did *common.Did) Option {
	return func(c *v1.ServiceConfig) {
		c.Owner = did
	}
}

// WithId is a Service Option that sets the service tags.
func WithId(did *common.Did) Option {
	return func(c *v1.ServiceConfig) {
		c.Id = did
	}
}

// WithChannels is a Service Option that sets the service channels.
func WithChannels(channels ...*common.Did) Option {
	return func(c *v1.ServiceConfig) {
		c.Channels = channels
	}
}

// WithBuckets is a Service Option that sets the service buckets.
func WithBuckets(buckets ...*common.Did) Option {
	return func(c *v1.ServiceConfig) {
		c.Buckets = buckets
	}
}

// WithObjects is a Service Option that sets the service objects.
func WithObjects(objects ...*common.ObjectDoc) Option {
	// Create Objects map from objects
	objectsMap := make(map[string]*common.ObjectDoc)
	for _, o := range objects {
		objectsMap[o.GetDid()] = o
	}

	// Return option
	return func(c *v1.ServiceConfig) {
		c.Objects = objectsMap
	}
}

// WithEndpoints is a Service Option that sets the service endpoints.
func WithEndpoints(endpoints ...string) Option {
	return func(c *v1.ServiceConfig) {
		c.Endpoints = endpoints
	}
}

// WithMetadata is a Service Option that sets the service metadata.
func WithMetadata(metadata map[string]string) Option {
	return func(c *v1.ServiceConfig) {
		c.Metadata = metadata
	}
}

// WithVersion is a Service Option that sets the service version.
func WithVersion(version string) Option {
	return func(c *v1.ServiceConfig) {
		c.Version = version
	}
}

func defaultConfig() *v1.ServiceConfig {
	return &v1.ServiceConfig{
		Name:        "",
		Description: "",
		Owner:       nil,
		Id:          nil,
		Channels:    []*common.Did{},
		Buckets:     []*common.Did{},
		Objects:     make(map[string]*common.ObjectDoc),
		Endpoints:   []string{},
		Metadata:    make(map[string]string),
		Version:     "0.0.1",
	}
}
