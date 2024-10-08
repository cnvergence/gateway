// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package infrastructure

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	clicfg "sigs.k8s.io/controller-runtime/pkg/client/config"

	egv1a1 "github.com/envoyproxy/gateway/api/v1alpha1"
	"github.com/envoyproxy/gateway/internal/envoygateway"
	"github.com/envoyproxy/gateway/internal/envoygateway/config"
	"github.com/envoyproxy/gateway/internal/infrastructure/kubernetes"
	"github.com/envoyproxy/gateway/internal/ir"
)

var _ Manager = (*kubernetes.Infra)(nil)

// Manager provides the scaffolding for managing infrastructure.
type Manager interface {
	// CreateOrUpdateProxyInfra creates or updates infra.
	CreateOrUpdateProxyInfra(ctx context.Context, infra *ir.Infra) error
	// DeleteProxyInfra deletes infra.
	DeleteProxyInfra(ctx context.Context, infra *ir.Infra) error
	// CreateOrUpdateRateLimitInfra creates or updates rate limit infra.
	CreateOrUpdateRateLimitInfra(ctx context.Context) error
	// DeleteRateLimitInfra deletes rate limit infra.
	DeleteRateLimitInfra(ctx context.Context) error
}

// NewManager returns a new infrastructure Manager.
func NewManager(cfg *config.Server) (Manager, error) {
	var mgr Manager

	switch cfg.EnvoyGateway.Provider.Type {
	case egv1a1.ProviderTypeKubernetes:
		cli, err := client.New(clicfg.GetConfigOrDie(), client.Options{Scheme: envoygateway.GetScheme()})
		if err != nil {
			return nil, err
		}
		mgr = kubernetes.NewInfra(cli, cfg)
	case egv1a1.ProviderTypeCustom:
		infra := cfg.EnvoyGateway.Provider.Custom.Infrastructure
		switch infra.Type {
		case egv1a1.InfrastructureProviderTypeHost:
			// TODO(sh2): implement host provider
			return nil, fmt.Errorf("host provider is not available yet")
		default:
			return nil, fmt.Errorf("unsupported provider type: %s", infra.Type)
		}
	}

	return mgr, nil
}
