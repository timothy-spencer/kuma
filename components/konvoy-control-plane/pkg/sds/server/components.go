package server

import (
	"context"

	"github.com/pkg/errors"

	envoy "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"

	konvoy_cp "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/config/app/konvoy-cp"
	core_mesh "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/apis/mesh"
	core_manager "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/manager"
	core_store "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/store"
	core_runtime "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/runtime"
	core_xds "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/xds"
	sds_auth "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/sds/auth"
	k8s_sds_auth "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/sds/auth/k8s"
	universal_sds_auth "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/sds/auth/universal"
	sds_provider "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/sds/provider"
	ca_sds_provider "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/sds/provider/ca"
	identity_sds_provider "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/sds/provider/identity"
)

const (
	MeshCaResource       = "mesh_ca"
	IdentityCertResource = "identity_cert"
)

func DefaultAuthenticator(rt core_runtime.Runtime) (sds_auth.Authenticator, error) {
	switch env := rt.Config().Environment; env {
	case konvoy_cp.KubernetesEnvironment:
		return k8s_sds_auth.New(), nil
	case konvoy_cp.UniversalEnvironment:
		return universal_sds_auth.New(), nil
	default:
		return nil, errors.Errorf("unable to choose SDS authenticator for environment type %q", env)
	}
}

func DefaultDataplaneResolver(resourceManager core_manager.ResourceManager) func(core_xds.ProxyId) (*core_mesh.DataplaneResource, error) {
	return func(proxyId core_xds.ProxyId) (*core_mesh.DataplaneResource, error) {
		ctx := context.Background()
		dataplane := &core_mesh.DataplaneResource{}
		if err := resourceManager.Get(ctx, dataplane, core_store.GetBy(proxyId.ToResourceKey())); err != nil {
			return nil, err
		}
		return dataplane, nil
	}
}

func DefaultMeshCaProvider() sds_provider.SecretProvider {
	return ca_sds_provider.New()
}

func DefaultIdentityCertProvider() sds_provider.SecretProvider {
	return identity_sds_provider.New()
}

func GetSecretProvider(resource string) (sds_provider.SecretProvider, error) {
	switch resource {
	case MeshCaResource:
		return DefaultMeshCaProvider(), nil
	case IdentityCertResource:
		return DefaultIdentityCertProvider(), nil
	default:
		return nil, errors.Errorf("SDS request for %q resource is not supported", resource)
	}
}

func DefaultSecretDiscoveryHandler(rt core_runtime.Runtime) (SecretDiscoveryHandler, error) {
	authenticator, err := DefaultAuthenticator(rt)
	if err != nil {
		return nil, err
	}
	return SecretDiscoveryHandlerFunc(func(ctx context.Context, req envoy.DiscoveryRequest) (*envoy_auth.Secret, error) {
		resource := req.ResourceNames[0]
		provider, err := GetSecretProvider(resource)
		if err != nil {
			return nil, err
		}
		proxyId, err := core_xds.ParseProxyId(req.Node)
		if err != nil {
			return nil, errors.Wrap(err, "SDS request must have a valid Proxy Id")
		}
		requestor := sds_auth.Identity{Mesh: proxyId.Mesh}
		if provider.RequiresIdentity() {
			credential, err := sds_auth.ExtractCredential(ctx)
			if err != nil {
				return nil, err
			}
			requestor, err = authenticator.Authenticate(ctx, *proxyId, credential)
			if err != nil {
				return nil, err
			}
		}
		secret, err := provider.Get(ctx, resource, requestor)
		if err != nil {
			return nil, err
		}
		return secret.ToResource(resource), nil
	}), nil
}

type SecretDiscoveryHandlerFunc func(ctx context.Context, req envoy.DiscoveryRequest) (*envoy_auth.Secret, error)

func (f SecretDiscoveryHandlerFunc) Handle(ctx context.Context, req envoy.DiscoveryRequest) (*envoy_auth.Secret, error) {
	return f(ctx, req)
}
