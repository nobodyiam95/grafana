package apiregistry

import (
	"context"

	playlistsv1 "github.com/grafana/grafana/pkg/apis/playlist/v1"
	"github.com/grafana/grafana/pkg/registry"
)

var (
	_ registry.BackgroundService = (*Service)(nil)
)

type Service struct{}

func ProvideService(
	_ *playlistsv1.PlaylistAPIBuilder,
) *Service {
	return &Service{}
}

func (s *Service) Run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}
