// Code generated by solo-kit. DO NOT EDIT.

package enterprise

import (
	"context"
	"errors"
	"fmt"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/client"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/server"
)

// Type Definitions:

const RateLimitConfigType = cache.TypePrefix + "/glooe.solo.io.RateLimitConfig"

/* Defined a resource - to be used by snapshot */
type RateLimitConfigXdsResourceWrapper struct {
	// TODO(yuval-k): This is public for mitchellh hashstructure to work properly. consider better alternatives.
	Resource *RateLimitConfig
}

// Make sure the Resource interface is implemented
var _ cache.Resource = &RateLimitConfigXdsResourceWrapper{}

func NewRateLimitConfigXdsResourceWrapper(resourceProto *RateLimitConfig) *RateLimitConfigXdsResourceWrapper {
	return &RateLimitConfigXdsResourceWrapper{
		Resource: resourceProto,
	}
}

func (e *RateLimitConfigXdsResourceWrapper) Self() cache.XdsResourceReference {
	return cache.XdsResourceReference{Name: e.Resource.Domain, Type: RateLimitConfigType}
}

func (e *RateLimitConfigXdsResourceWrapper) ResourceProto() cache.ResourceProto {
	return e.Resource
}
func (e *RateLimitConfigXdsResourceWrapper) References() []cache.XdsResourceReference {
	return nil
}

// Define a type record. This is used by the generic client library.
var RateLimitConfigTypeRecord = client.NewTypeRecord(
	RateLimitConfigType,

	// Return an empty message, that can be used to deserialize bytes into it.
	func() cache.ResourceProto { return &RateLimitConfig{} },

	// Covert the message to a resource suitable for use for protobuf's Any.
	func(r cache.ResourceProto) cache.Resource {
		return &RateLimitConfigXdsResourceWrapper{Resource: r.(*RateLimitConfig)}
	},
)

// Server Implementation:

// Wrap the generic server and implement the type sepcific methods:
type rateLimitDiscoveryServiceServer struct {
	server.Server
}

func NewRateLimitDiscoveryServiceServer(genericServer server.Server) RateLimitDiscoveryServiceServer {
	return &rateLimitDiscoveryServiceServer{Server: genericServer}
}

func (s *rateLimitDiscoveryServiceServer) StreamRateLimitConfig(stream RateLimitDiscoveryService_StreamRateLimitConfigServer) error {
	return s.Server.Stream(stream, RateLimitConfigType)
}

func (s *rateLimitDiscoveryServiceServer) FetchRateLimitConfig(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = RateLimitConfigType
	return s.Server.Fetch(ctx, req)
}

func (s *rateLimitDiscoveryServiceServer) DeltaRateLimitConfig(_ RateLimitDiscoveryService_DeltaRateLimitConfigServer) error {
	return errors.New("not implemented")
}

// Client Implementation: Generate a strongly typed client over the generic client

// The apply functions receives resources and returns an error if they were applied correctly.
// In theory the configuration can become valid in the future (i.e. eventually consistent), but I don't think we need to worry about that now
// As our current use cases only have one configuration resource, so no interactions are expected.
type ApplyRateLimitConfig func(version string, resources []*RateLimitConfig) error

// Convert the strongly typed apply to a generic apply.
func applyRateLimitConfig(typedApply ApplyRateLimitConfig) func(cache.Resources) error {
	return func(resources cache.Resources) error {

		var configs []*RateLimitConfig
		for _, r := range resources.Items {
			if proto, ok := r.ResourceProto().(*RateLimitConfig); !ok {
				return fmt.Errorf("resource %s of type %s incorrect", r.Self().Name, r.Self().Type)
			} else {
				configs = append(configs, proto)
			}
		}

		return typedApply(resources.Version, configs)
	}
}

func NewRateLimitConfigClient(nodeinfo *core.Node, typedApply ApplyRateLimitConfig) client.Client {
	return client.NewClient(nodeinfo, RateLimitConfigTypeRecord, applyRateLimitConfig(typedApply))
}
