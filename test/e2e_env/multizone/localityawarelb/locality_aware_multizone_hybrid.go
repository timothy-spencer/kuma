package localityawarelb

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gruntwork-io/terratest/modules/k8s"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	mesh_http_route_api "github.com/kumahq/kuma/pkg/plugins/policies/meshhttproute/api/v1alpha1"
	mesh_loadbalancing_api "github.com/kumahq/kuma/pkg/plugins/policies/meshloadbalancingstrategy/api/v1alpha1"
	. "github.com/kumahq/kuma/test/framework"
	"github.com/kumahq/kuma/test/framework/client"
	"github.com/kumahq/kuma/test/framework/deployments/democlient"
	"github.com/kumahq/kuma/test/framework/deployments/testserver"
	"github.com/kumahq/kuma/test/framework/envs/multizone"
)

func LocalityAwareLB() {
	const mesh = "locality-aware-lb"
	const namespace = "locality-aware-lb"

	meshLoadBalancingStrategyDemoClient := YamlUniversal(fmt.Sprintf(`
type: MeshLoadBalancingStrategy
name: mlbs-1
mesh: %s
spec:
  targetRef:
    kind: MeshService
    name: demo-client_locality-aware-lb_svc
  to:
    - targetRef:
        kind: MeshService
        name: test-server_locality-aware-lb_svc_80
      default:
        localityAwareness:
          localZone:
            affinityTags:
            - key: k8s.io/node
            - key: k8s.io/az
          crossZone:
            failover:
              - from:
                  zones: ["kuma-4", "kuma-5"]
                to:
                  type: Only
                  zones: ["kuma-4", "kuma-5"]
              - from:
                  zones: ["kuma-1-zone", "kuma-2-zone"]
                to:
                  type: Only
                  zones: ["kuma-1-zone", "kuma-2-zone"]
              - from:
                  zones: ["kuma-4"]
                to:
                  type: Only
                  zones: ["kuma-1-zone"]`, mesh))
	meshLoadBalancingStrategyDemoClientMeshRoute := YamlUniversal(fmt.Sprintf(`
type: MeshLoadBalancingStrategy
name: mlbs-2
mesh: %s
spec:
  targetRef:
    kind: MeshService
    name: demo-client-mesh-route_locality-aware-lb_svc
  to:
    - targetRef:
        kind: MeshService
        name: test-server-mesh-route_locality-aware-lb_svc_80
      default:
        localityAwareness:
          crossZone:
            failover:
              - to:
                  type: Only
                  zones: ["kuma-1-zone"]`, mesh))

	BeforeAll(func() {
		// Global
		Expect(NewClusterSetup().
			Install(MTLSMeshUniversal(mesh)).
			Setup(multizone.Global)).To(Succeed())
		Expect(WaitForMesh(mesh, multizone.Zones())).To(Succeed())

		// Universal Zone 4
		Expect(NewClusterSetup().
			Install(DemoClientUniversal(
				"demo-client_locality-aware-lb_svc",
				mesh,
				WithTransparentProxy(true),
				WithAdditionalTags(
					map[string]string{
						"k8s.io/node": "node-1",
						"k8s.io/az":   "az-1",
					}),
			)).
			Install(TestServerUniversal("test-server-node-1", mesh, WithAdditionalTags(
				map[string]string{
					"k8s.io/node": "node-1",
					"k8s.io/az":   "az-1",
				}),
				WithServiceName("test-server_locality-aware-lb_svc_80"),
				WithArgs([]string{"echo", "--instance", "test-server-node-1-zone-4"}),
			)).
			Install(TestServerUniversal("test-server-az-1", mesh, WithAdditionalTags(
				map[string]string{
					"k8s.io/az": "az-1",
				}),
				WithServiceName("test-server_locality-aware-lb_svc_80"),
				WithArgs([]string{"echo", "--instance", "test-server-az-1-zone-4"}),
			)).
			Install(TestServerUniversal("test-server-node-2", mesh, WithAdditionalTags(
				map[string]string{
					"k8s.io/node": "node-2",
				}),
				WithServiceName("test-server_locality-aware-lb_svc_80"),
				WithArgs([]string{"echo", "--instance", "test-server-node-2-zone-4"}),
			)).
			Install(TestServerUniversal("test-server-no-tags", mesh,
				WithArgs([]string{"echo", "--instance", "test-server-no-tags-zone-4"}),
				WithServiceName("test-server_locality-aware-lb_svc_80"),
			)).
			Setup(multizone.UniZone1),
		).To(Succeed())

		// Universal Zone 5
		Expect(NewClusterSetup().
			Install(DemoClientUniversal(
				"demo-client_locality-aware-lb_svc",
				mesh,
				WithTransparentProxy(true),
			)).
			Install(TestServerUniversal("test-server", mesh,
				WithServiceName("test-server_locality-aware-lb_svc_80"),
				WithArgs([]string{"echo", "--instance", "test-server-zone-5"}),
			)).
			Install(TestServerUniversal("test-server-mesh-route", mesh,
				WithServiceName("test-server-mesh-route_locality-aware-lb_svc_80"),
				WithArgs([]string{"echo", "--instance", "test-server-mesh-route-zone-5"}),
			)).
			Setup(multizone.UniZone2),
		).To(Succeed())

		// Kubernetes Zone 1
		Expect(NewClusterSetup().
			Install(NamespaceWithSidecarInjection(namespace)).
			Install(democlient.Install(democlient.WithMesh(mesh), democlient.WithNamespace(namespace))).
			Install(testserver.Install(
				testserver.WithName("test-server"),
				testserver.WithMesh(mesh),
				testserver.WithNamespace(namespace),
				testserver.WithEchoArgs("echo", "--instance", "test-server-zone-1"),
			)).
			Install(testserver.Install(
				testserver.WithName("test-server-mesh-route"),
				testserver.WithMesh(mesh),
				testserver.WithNamespace(namespace),
				testserver.WithEchoArgs("echo", "--instance", "test-server-mesh-route-zone-1"),
			)).
			Setup(multizone.KubeZone1)).ToNot(HaveOccurred())

		// Kubernetes Zone 2
		Expect(NewClusterSetup().
			Install(NamespaceWithSidecarInjection(namespace)).
			Install(democlient.Install(democlient.WithMesh(mesh), democlient.WithNamespace(namespace))).
			Install(democlient.Install(
				democlient.WithMesh(mesh),
				democlient.WithNamespace(namespace),
				democlient.WithName("demo-client-mesh-route"),
			)).
			Install(testserver.Install(
				testserver.WithName("test-server"),
				testserver.WithMesh(mesh),
				testserver.WithNamespace(namespace),
				testserver.WithEchoArgs("echo", "--instance", "test-server-zone-2"),
			)).
			Setup(multizone.KubeZone2)).ToNot(HaveOccurred())
	})

	E2EAfterAll(func() {
		Expect(multizone.UniZone1.DeleteMeshApps(mesh)).To(Succeed())
		Expect(multizone.UniZone2.DeleteMeshApps(mesh)).To(Succeed())
		Expect(multizone.KubeZone1.TriggerDeleteNamespace(namespace)).To(Succeed())
		Expect(multizone.KubeZone2.TriggerDeleteNamespace(namespace)).To(Succeed())
		Expect(multizone.Global.DeleteMesh(mesh)).To(Succeed())
	})

	E2EAfterEach(func() {
		Expect(DeleteMeshResources(
			multizone.Global,
			mesh,
			mesh_loadbalancing_api.MeshLoadBalancingStrategyResourceTypeDescriptor,
		)).To(Succeed())
		Expect(DeleteMeshResources(
			multizone.Global,
			mesh,
			mesh_http_route_api.MeshHTTPRouteResourceTypeDescriptor,
		)).To(Succeed())
	})

	successRequestCount := func(cluster Cluster, name string) (int, error) {
		return collectMetric(cluster, name, namespace, "test-server-mesh-route_locality-aware-lb_svc_80-f3615f466e37e855.upstream_rq_2xx")
	}

	failRequestCount := func(cluster Cluster, name string) (int, error) {
		return collectMetric(cluster, name, namespace, "test-server-mesh-route_locality-aware-lb_svc_80-ce3d32a0f959e460.upstream_rq_5xx")
	}

	It("should route based on defined strategy", func() {
		// should load balance traffic equally when no policy
		Eventually(func() (map[string]int, error) {
			return client.CollectResponsesByInstance(multizone.UniZone1, "demo-client_locality-aware-lb_svc", "test-server_locality-aware-lb_svc_80.mesh", client.WithNumberOfRequests(100))
		}, "30s", "5s").Should(
			And(
				HaveKeyWithValue(Equal(`test-server-node-1-zone-4`), BeNumerically("~", 20, 3)),
				HaveKeyWithValue(Equal(`test-server-az-1-zone-4`), BeNumerically("~", 20, 3)),
				HaveKeyWithValue(Equal(`test-server-node-2-zone-4`), BeNumerically("~", 20, 3)),
				HaveKeyWithValue(Equal(`test-server-no-tags-zone-4`), BeNumerically("~", 20, 3)),
				HaveKeyWithValue(Equal(`test-server-zone-5`), BeNumerically("~", 6, 2)),
				HaveKeyWithValue(Equal(`test-server-zone-1`), BeNumerically("~", 6, 2)),
				HaveKeyWithValue(Equal(`test-server-zone-2`), BeNumerically("~", 6, 2)),
			),
		)

		// when load balancing policy created
		Expect(NewClusterSetup().Install(meshLoadBalancingStrategyDemoClient).Setup(multizone.Global)).To(Succeed())

		// then traffic should be routed based on the policy
		Eventually(func() (map[string]int, error) {
			return client.CollectResponsesByInstance(multizone.UniZone1, "demo-client_locality-aware-lb_svc", "test-server_locality-aware-lb_svc_80.mesh", client.WithNumberOfRequests(100))
		}, "30s", "5s").Should(
			And(
				HaveKeyWithValue(Equal(`test-server-node-1-zone-4`), BeNumerically("~", 93, 5)),
				HaveKeyWithValue(Equal(`test-server-az-1-zone-4`), BeNumerically("~", 7, 5)),
				Or(
					HaveKeyWithValue(Equal(`test-server-node-2-zone-4`), BeNumerically("~", 2, 2)),
					HaveKeyWithValue(Equal(`test-server-no-tags-zone-4`), BeNumerically("~", 2, 2)),
					Not(HaveKeyWithValue(Equal(`test-server-node-2-zone-4`), BeNumerically("~", 2, 2))),
					Not(HaveKeyWithValue(Equal(`test-server-no-tags-zone-4`), BeNumerically("~", 2, 2))),
				),
			),
		)

		// when app with the highest weight is down
		Expect(multizone.UniZone1.DeleteApp("test-server-node-1")).To(Succeed())

		// then traffic goes to the next highest
		Eventually(func() (map[string]int, error) {
			return client.CollectResponsesByInstance(multizone.UniZone1, "demo-client_locality-aware-lb_svc", "test-server_locality-aware-lb_svc_80.mesh", client.WithNumberOfRequests(100))
		}, "30s", "5s").Should(
			And(
				HaveKeyWithValue(Equal(`test-server-az-1-zone-4`), BeNumerically("~", 90, 5)),
				Or(
					HaveKeyWithValue(Equal(`test-server-node-2-zone-4`), BeNumerically("~", 5, 3)),
					HaveKeyWithValue(Equal(`test-server-no-tags-zone-4`), BeNumerically("~", 5, 3)),
				),
			),
		)

		// when the next app with the highest weight is down
		Expect(multizone.UniZone1.DeleteApp("test-server-az-1")).To(Succeed())

		// then traffic goes to the next highest
		Eventually(func() (map[string]int, error) {
			return client.CollectResponsesByInstance(multizone.UniZone1, "demo-client_locality-aware-lb_svc", "test-server_locality-aware-lb_svc_80.mesh", client.WithNumberOfRequests(100))
		}, "30s", "5s").Should(
			And(
				HaveKeyWithValue(Equal(`test-server-node-2-zone-4`), BeNumerically("~", 50, 5)),
				HaveKeyWithValue(Equal(`test-server-no-tags-zone-4`), BeNumerically("~", 50, 5)),
			),
		)

		// when all apps in local zone down
		Expect(multizone.UniZone1.DeleteApp("test-server-node-2")).To(Succeed())
		Expect(multizone.UniZone1.DeleteApp("test-server-no-tags")).To(Succeed())

		// then traffic goes to the zone with the next priority, kuma-5
		Eventually(func() (map[string]int, error) {
			return client.CollectResponsesByInstance(multizone.UniZone1, "demo-client_locality-aware-lb_svc", "test-server_locality-aware-lb_svc_80.mesh", client.WithNumberOfRequests(100))
		}, "30s", "5s").Should(HaveKeyWithValue(Equal(`test-server-zone-5`), BeNumerically("==", 100)))

		// when zone kuma-5 is disabled
		Expect(multizone.UniZone2.DeleteApp("test-server")).To(Succeed())

		// then traffic should goes to k8s
		Eventually(func() (map[string]int, error) {
			return client.CollectResponsesByInstance(multizone.UniZone1, "demo-client_locality-aware-lb_svc", "test-server_locality-aware-lb_svc_80.mesh", client.WithNumberOfRequests(100))
		}, "30s", "5s").Should(HaveKeyWithValue(Equal(`test-server-zone-1`), BeNumerically("==", 100)))

		// when zone kuma-1-zone is disabled
		Expect(DeleteK8sApp(multizone.KubeZone1, "test-server", namespace)).To(Succeed())

		// then traffic should goes to k8s
		Eventually(func(g Gomega) {
			_, err := client.CollectEchoResponse(
				multizone.UniZone1, "demo-client_locality-aware-lb_svc", "test-server_locality-aware-lb_svc_80.mesh",
			)
			g.Expect(err).To(HaveOccurred())
		}).Should(Succeed())
	})

	It("should route based on the strategy when split defined", func() {
		// given
		Expect(YamlUniversal(fmt.Sprintf(`
type: MeshHTTPRoute
name: route-1
mesh: %s
spec:
  targetRef:
    kind: MeshService
    name: demo-client-mesh-route_locality-aware-lb_svc
  to:
    - targetRef:
        kind: MeshService
        name: test-server-mesh-route_locality-aware-lb_svc_80
      rules:
        - matches:
          - path:
              value: /
              type: PathPrefix
          default:
            backendRefs:
              - kind: MeshServiceSubset
                name: test-server-mesh-route_locality-aware-lb_svc_80
                weight: 50
                tags:
                  kuma.io/zone: kuma-5
              - kind: MeshServiceSubset
                name: test-server-mesh-route_locality-aware-lb_svc_80
                weight: 50
                tags:
                  kuma.io/zone: kuma-1-zone
`, mesh))(multizone.Global)).To(Succeed())

		// then traffic should be routed equally
		Eventually(func() (map[string]int, error) {
			return client.CollectResponsesByInstance(
				multizone.KubeZone2, "demo-client-mesh-route", "test-server-mesh-route_locality-aware-lb_svc_80.mesh",
				client.WithNumberOfRequests(200),
				client.FromKubernetesPod(namespace, "demo-client-mesh-route"),
			)
		}, "30s", "10s").Should(
			And(
				HaveKeyWithValue(Equal(`test-server-mesh-route-zone-1`), BeNumerically("~", 100, 20)),
				HaveKeyWithValue(Equal(`test-server-mesh-route-zone-5`), BeNumerically("~", 100, 20)),
			),
		)

		// clean stats
		Expect(resetCounter(multizone.KubeZone2, "demo-client-mesh-route", namespace)).To(Succeed())

		// when load balancing policy created
		Expect(NewClusterSetup().Install(meshLoadBalancingStrategyDemoClientMeshRoute).Setup(multizone.Global)).To(Succeed())

		// and generate some traffic
		_, err := client.CollectResponsesAndFailures(
			multizone.KubeZone2, "demo-client-mesh-route", "test-server-mesh-route_locality-aware-lb_svc_80.mesh",
			client.WithNumberOfRequests(200),
			client.NoFail(),
			client.WithoutRetries(),
			client.FromKubernetesPod(namespace, "demo-client-mesh-route"),
		)
		Expect(err).ToNot(HaveOccurred())

		// then
		failedRequests, err := failRequestCount(multizone.KubeZone2, "demo-client-mesh-route")
		Expect(err).ToNot(HaveOccurred())
		Expect(failedRequests).To(BeNumerically("~", 100, 25))
		successRequests, err := successRequestCount(multizone.KubeZone2, "demo-client-mesh-route")
		Expect(err).ToNot(HaveOccurred())
		Expect(successRequests).To(BeNumerically("~", 100, 25))
	})
}

func collectMetric(cluster Cluster, name string, namespace string, metricName string) (int, error) {
	resp, _, err := client.CollectResponse(cluster, name, fmt.Sprintf("http://localhost:9901/stats?filter=%s", metricName), client.FromKubernetesPod(namespace, name))
	if err != nil {
		return -1, err
	}
	split := strings.Split(resp, ": ")
	if len(split) == 2 {
		i, err := strconv.Atoi(split[1])
		if err != nil {
			return -1, err
		}
		return i, nil
	} else {
		return -1, errors.New("no metric found")
	}
}

func resetCounter(cluster Cluster, name string, namespace string) error {
	_, _, err := client.CollectResponse(
		cluster, name, "http://localhost:9901/reset_counters",
		client.FromKubernetesPod(namespace, name),
		client.WithMethod("POST"),
	)
	return err
}

// TODO(lukidzi): use test-server implementation: https://github.com/kumahq/kuma/issues/8245
func DeleteK8sApp(c Cluster, name string, namespace string) error {
	if err := k8s.RunKubectlE(c.GetTesting(), c.GetKubectlOptions(namespace), "delete", "service", name); err != nil {
		return err
	}
	if err := k8s.RunKubectlE(c.GetTesting(), c.GetKubectlOptions(namespace), "delete", "deployment", name); err != nil {
		return err
	}
	return nil
}
