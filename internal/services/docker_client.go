package services

import (
	"context"
	"io"

	c "github.com/docker/docker/api/types/container"
	i "github.com/docker/docker/api/types/image"
	n "github.com/docker/docker/api/types/network"
	v "github.com/docker/docker/api/types/volume"
	o "github.com/opencontainers/image-spec/specs-go/v1"
)

type IDockerClient interface {
	ContainerStop(ctx context.Context, containerID string, options c.StopOptions) error
	ContainerRemove(ctx context.Context, containerID string, options c.RemoveOptions) error
	ContainerList(ctx context.Context, options c.ListOptions) ([]c.Summary, error)
	ContainerCreate(ctx context.Context, config *c.Config, hostConfig *c.HostConfig, networkingConfig *n.NetworkingConfig, platform *o.Platform, containerName string) (c.CreateResponse, error)
	ContainerStart(ctx context.Context, containerID string, options c.StartOptions) error
	VolumeCreate(ctx context.Context, options v.CreateOptions) (v.Volume, error)
	VolumeList(ctx context.Context, options v.ListOptions) (v.ListResponse, error)
	ImagePull(ctx context.Context, image string, options i.PullOptions) (io.ReadCloser, error)
}
