package server

import (
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core"
	core_discovery "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/discovery"
	mesh_core "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/apis/mesh"
	core_model "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/model"
	model "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/xds"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/xds/generator"

	envoy "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	envoy_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_cache "github.com/envoyproxy/go-control-plane/pkg/cache"
)

var (
	reconcileLog = core.Log.WithName("xds-server").WithName("reconcile")
)

type reconciler struct {
	generator snapshotGenerator
	cacher    snapshotCacher
}

// Make sure that reconciler implements all relevant interfaces
var _ core_discovery.DataplaneDiscoveryConsumer = &reconciler{}

func (r *reconciler) OnDataplaneUpdate(dataplane *mesh_core.DataplaneResource) error {
	proxyId := model.ProxyId{Name: dataplane.Meta.GetName(), Namespace: dataplane.Meta.GetNamespace(), Mesh: dataplane.Meta.GetMesh()}
	return r.reconcile(
		&envoy_core.Node{Id: proxyId.String()},
		&model.Proxy{
			Id:        proxyId,
			Dataplane: dataplane,
		})
}
func (r *reconciler) OnDataplaneDelete(key core_model.ResourceKey) error {
	proxyId := model.ProxyId{Name: key.Name, Namespace: key.Namespace, Mesh: key.Mesh}
	// clearing cache does not push a new (empty) configuration to Envoy
	// that's why instead of clearing we explicitly set new configuration to empty value
	// cache will be actually cleared on Envoy disconnect
	return r.cacher.Cache(&envoy_core.Node{Id: proxyId.String()}, envoy_cache.Snapshot{})
}

func (r *reconciler) reconcile(node *envoy_core.Node, proxy *model.Proxy) error {
	snapshot, err := r.generator.GenerateSnapshot(proxy)
	if err != nil {
		reconcileLog.Error(err, "failed to generate a snapshot", "node", node, "proxy", proxy)
		return err
	}
	if err := snapshot.Consistent(); err != nil {
		reconcileLog.Error(err, "inconsistent snapshot", "snapshot", snapshot, "proxy", proxy)
	}
	// to avoid assigning a new version every time,
	// compare with the previous snaphsot and reuse its version whenever possible,
	// fallback to UUID otherwise
	previous, err := r.cacher.Get(node)
	if err != nil {
		previous = envoy_cache.Snapshot{}
	}
	snapshot = r.autoVersion(previous, snapshot)
	if err := r.cacher.Cache(node, snapshot); err != nil {
		reconcileLog.Error(err, "failed to store snapshot", "snapshot", snapshot, "proxy", proxy)
	}
	return nil
}

func (r *reconciler) autoVersion(old envoy_cache.Snapshot, new envoy_cache.Snapshot) envoy_cache.Snapshot {
	if new.Listeners.Version == "" {
		new.Listeners.Version = old.Listeners.Version
		if !EqualSnapshots(old.Listeners.Items, new.Listeners.Items) {
			new.Listeners.Version = core.NewUUID()
		}
	}
	if new.Routes.Version == "" {
		new.Routes.Version = old.Routes.Version
		if !EqualSnapshots(old.Routes.Items, new.Routes.Items) {
			new.Routes.Version = core.NewUUID()
		}
	}
	if new.Clusters.Version == "" {
		new.Clusters.Version = old.Clusters.Version
		if !EqualSnapshots(old.Clusters.Items, new.Clusters.Items) {
			new.Clusters.Version = core.NewUUID()
		}
	}
	if new.Endpoints.Version == "" {
		new.Endpoints.Version = old.Endpoints.Version
		if !EqualSnapshots(old.Endpoints.Items, new.Endpoints.Items) {
			new.Endpoints.Version = core.NewUUID()
		}
	}
	if new.Secrets.Version == "" {
		new.Secrets.Version = old.Secrets.Version
		if !EqualSnapshots(old.Secrets.Items, new.Secrets.Items) {
			new.Secrets.Version = core.NewUUID()
		}
	}
	return new
}

func EqualSnapshots(old, new map[string]envoy_cache.Resource) bool {
	if len(new) != len(old) {
		return false
	}
	for key, newValue := range new {
		if oldValue, hasOldValue := old[key]; !hasOldValue || !newValue.Equal(oldValue) {
			return false
		}
	}
	return true
}

type snapshotGenerator interface {
	GenerateSnapshot(proxy *model.Proxy) (envoy_cache.Snapshot, error)
}

type templateSnapshotGenerator struct {
	ProxyTemplateResolver proxyTemplateResolver
}

func (s *templateSnapshotGenerator) GenerateSnapshot(proxy *model.Proxy) (envoy_cache.Snapshot, error) {
	template := s.ProxyTemplateResolver.GetTemplate(proxy)

	gen := generator.TemplateProxyGenerator{ProxyTemplate: template}

	rs, err := gen.Generate(proxy)
	if err != nil {
		reconcileLog.Error(err, "failed to generate a snapshot", "proxy", proxy, "template", template)
		return envoy_cache.Snapshot{}, err
	}

	listeners := []envoy_cache.Resource{}
	routes := []envoy_cache.Resource{}
	clusters := []envoy_cache.Resource{}
	endpoints := []envoy_cache.Resource{}
	secrets := []envoy_cache.Resource{}

	for _, r := range rs {
		switch r.Resource.(type) {
		case *envoy.Listener:
			listeners = append(listeners, r.Resource)
		case *envoy.RouteConfiguration:
			routes = append(routes, r.Resource)
		case *envoy.Cluster:
			clusters = append(clusters, r.Resource)
		case *envoy.ClusterLoadAssignment:
			endpoints = append(endpoints, r.Resource)
		case *envoy_auth.Secret:
			secrets = append(secrets, r.Resource)
		default:
		}
	}

	version := "" // empty value is a sign to other components to generate the version automatically
	out := envoy_cache.Snapshot{
		Endpoints: envoy_cache.NewResources(version, endpoints),
		Clusters:  envoy_cache.NewResources(version, clusters),
		Routes:    envoy_cache.NewResources(version, routes),
		Listeners: envoy_cache.NewResources(version, listeners),
		Secrets:   envoy_cache.NewResources(version, secrets),
	}

	return out, nil
}

type snapshotCacher interface {
	Get(*envoy_core.Node) (envoy_cache.Snapshot, error)
	Cache(*envoy_core.Node, envoy_cache.Snapshot) error
	Clear(*envoy_core.Node)
}

type simpleSnapshotCacher struct {
	hasher envoy_cache.NodeHash
	store  envoy_cache.SnapshotCache
}

func (s *simpleSnapshotCacher) Get(node *envoy_core.Node) (envoy_cache.Snapshot, error) {
	return s.store.GetSnapshot(s.hasher.ID(node))
}

func (s *simpleSnapshotCacher) Cache(node *envoy_core.Node, snapshot envoy_cache.Snapshot) error {
	return s.store.SetSnapshot(s.hasher.ID(node), snapshot)
}

func (s *simpleSnapshotCacher) Clear(node *envoy_core.Node) {
	s.store.ClearSnapshot(s.hasher.ID(node))
}
