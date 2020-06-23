// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/pascomnet/nomad-driver-podman/client/containers"
	"github.com/pascomnet/nomad-driver-podman/client/containers_compat"
	"github.com/pascomnet/nomad-driver-podman/client/exec"
	"github.com/pascomnet/nomad-driver-podman/client/exec_compat"
	"github.com/pascomnet/nomad-driver-podman/client/images"
	"github.com/pascomnet/nomad-driver-podman/client/images_compat"
	"github.com/pascomnet/nomad-driver-podman/client/images_libpod"
	"github.com/pascomnet/nomad-driver-podman/client/manifests"
	"github.com/pascomnet/nomad-driver-podman/client/networks"
	"github.com/pascomnet/nomad-driver-podman/client/networks_compat"
	"github.com/pascomnet/nomad-driver-podman/client/pods"
	"github.com/pascomnet/nomad-driver-podman/client/system"
	"github.com/pascomnet/nomad-driver-podman/client/system_compat"
	"github.com/pascomnet/nomad-driver-podman/client/volumes"
)

// Default providesa container compatible interface experimental HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "podman.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new providesa container compatible interface experimental HTTP client.
func NewHTTPClient(formats strfmt.Registry) *ProvidesaContainerCompatibleInterfaceExperimental {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new providesa container compatible interface experimental HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *ProvidesaContainerCompatibleInterfaceExperimental {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new providesa container compatible interface experimental client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *ProvidesaContainerCompatibleInterfaceExperimental {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(ProvidesaContainerCompatibleInterfaceExperimental)
	cli.Transport = transport
	cli.Containers = containers.New(transport, formats)
	cli.ContainersCompat = containers_compat.New(transport, formats)
	cli.Exec = exec.New(transport, formats)
	cli.ExecCompat = exec_compat.New(transport, formats)
	cli.Images = images.New(transport, formats)
	cli.ImagesCompat = images_compat.New(transport, formats)
	cli.ImagesLibpod = images_libpod.New(transport, formats)
	cli.Manifests = manifests.New(transport, formats)
	cli.Networks = networks.New(transport, formats)
	cli.NetworksCompat = networks_compat.New(transport, formats)
	cli.Pods = pods.New(transport, formats)
	cli.System = system.New(transport, formats)
	cli.SystemCompat = system_compat.New(transport, formats)
	cli.Volumes = volumes.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// ProvidesaContainerCompatibleInterfaceExperimental is a client for providesa container compatible interface experimental
type ProvidesaContainerCompatibleInterfaceExperimental struct {
	Containers containers.ClientService

	ContainersCompat containers_compat.ClientService

	Exec exec.ClientService

	ExecCompat exec_compat.ClientService

	Images images.ClientService

	ImagesCompat images_compat.ClientService

	ImagesLibpod images_libpod.ClientService

	Manifests manifests.ClientService

	Networks networks.ClientService

	NetworksCompat networks_compat.ClientService

	Pods pods.ClientService

	System system.ClientService

	SystemCompat system_compat.ClientService

	Volumes volumes.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *ProvidesaContainerCompatibleInterfaceExperimental) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Containers.SetTransport(transport)
	c.ContainersCompat.SetTransport(transport)
	c.Exec.SetTransport(transport)
	c.ExecCompat.SetTransport(transport)
	c.Images.SetTransport(transport)
	c.ImagesCompat.SetTransport(transport)
	c.ImagesLibpod.SetTransport(transport)
	c.Manifests.SetTransport(transport)
	c.Networks.SetTransport(transport)
	c.NetworksCompat.SetTransport(transport)
	c.Pods.SetTransport(transport)
	c.System.SetTransport(transport)
	c.SystemCompat.SetTransport(transport)
	c.Volumes.SetTransport(transport)
}