package inspect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kumahq/kuma/api/openapi/types"
	. "github.com/kumahq/kuma/test/framework"
	"github.com/kumahq/kuma/test/framework/envs/multizone"
)

func Inspect() {
	const meshName = "inspect"

	BeforeAll(func() {
		Expect(multizone.Global.Install(MTLSMeshUniversal(meshName))).To(Succeed())
		Expect(WaitForMesh(meshName, multizone.Zones())).To(Succeed())

		err := multizone.UniZone1.Install(TestServerUniversal("test-server", meshName,
			WithArgs([]string{"echo", "--instance", "echo"}),
		))
		Expect(err).ToNot(HaveOccurred())
	})

	E2EAfterAll(func() {
		Expect(multizone.UniZone1.DeleteMeshApps(meshName)).To(Succeed())
		Expect(multizone.Global.DeleteMesh(meshName)).To(Succeed())
	})

	type testCase struct {
		cluster     func() Cluster
		args        []string
		expectedOut string
	}
	GlobalCluster := func() Cluster {
		return multizone.Global
	}
	UniZone1Cluster := func() Cluster {
		return multizone.UniZone1
	}

	Context("Dataplane", func() {
		DescribeTable("should execute envoy inspection",
			func(given testCase) {
				Eventually(func(g Gomega) {
					args := append([]string{"inspect"}, given.args...)
					out, err := given.cluster().GetKumactlOptions().RunKumactlAndGetOutput(args...)
					g.Expect(err).ToNot(HaveOccurred())
					g.Expect(out).Should(ContainSubstring(given.expectedOut))
				}, "30s", "1s").Should(Succeed())
			},
			Entry("of config dump for a dataplane using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"dataplane", "kuma-4.test-server", "--type", "config-dump", "--mesh", meshName},
				expectedOut: `"dataplane.proxyType": "dataplane"`,
			}),
			Entry("of stats for a dataplane using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"dataplane", "kuma-4.test-server", "--type", "stats", "--mesh", meshName},
				expectedOut: `server.live: 1`,
			}),
			Entry("of clusters for a dataplane using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"dataplane", "kuma-4.test-server", "--type", "clusters", "--mesh", meshName},
				expectedOut: `kuma:envoy:admin::`,
			}),
			Entry("of config dump for a dataplane using Zone CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"dataplane", "test-server", "--type", "config-dump", "--mesh", meshName},
				expectedOut: `"dataplane.proxyType": "dataplane"`,
			}),
			Entry("of stats for a dataplane using Zone CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"dataplane", "test-server", "--type", "stats", "--mesh", meshName},
				expectedOut: `server.live: 1`,
			}),
			Entry("of clusters for a dataplane using Zone CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"dataplane", "test-server", "--type", "clusters", "--mesh", meshName},
				expectedOut: `kuma:envoy:admin::`,
			}),
			Entry("of config dump for a zoneingress using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"zoneingress", "kuma-4.ingress", "--type", "config-dump"},
				expectedOut: `"dataplane.proxyType": "ingress"`,
			}),
			Entry("of stats for a zoneingress using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"zoneingress", "kuma-4.ingress", "--type", "stats"},
				expectedOut: `server.live: 1`,
			}),
			Entry("of clusters for a zoneingress using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"zoneingress", "kuma-4.ingress", "--type", "clusters"},
				expectedOut: `kuma:envoy:admin::`,
			}),
			Entry("of config dump for a zoneingress using Zone CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"zoneingress", "ingress", "--type", "config-dump"},
				expectedOut: `"dataplane.proxyType": "ingress"`,
			}),
			Entry("of stats for a zoneingress using Global CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"zoneingress", "ingress", "--type", "stats"},
				expectedOut: `server.live: 1`,
			}),
			Entry("of clusters for a zoneingress using Global CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"zoneingress", "ingress", "--type", "clusters"},
				expectedOut: `kuma:envoy:admin::`,
			}),
			Entry("of config dump for a zoneegress using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"zoneegress", "kuma-4.egress", "--type", "config-dump"},
				expectedOut: `"dataplane.proxyType": "egress"`,
			}),
			Entry("of stats for a zoneegress using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"zoneegress", "kuma-4.egress", "--type", "stats"},
				expectedOut: `server.live: 1`,
			}),
			Entry("of clusters for a zoneegress using Global CP", testCase{
				cluster:     GlobalCluster,
				args:        []string{"zoneegress", "kuma-4.egress", "--type", "clusters"},
				expectedOut: `kuma:envoy:admin::`,
			}),
			Entry("of config dump for a zoneegress using Zone CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"zoneegress", "egress", "--type", "config-dump"},
				expectedOut: `"dataplane.proxyType": "egress"`,
			}),
			Entry("of stats for a zoneegress using Global CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"zoneegress", "egress", "--type", "stats"},
				expectedOut: `server.live: 1`,
			}),
			Entry("of clusters for a zoneegress using Global CP", testCase{
				cluster:     UniZone1Cluster,
				args:        []string{"zoneegress", "egress", "--type", "clusters"},
				expectedOut: `kuma:envoy:admin::`,
			}),
		)

		It("match dataplanes of policy", func() {
			Eventually(func(g Gomega) {
				r, err := http.Get(multizone.Global.GetKuma().GetAPIServerAddress() + fmt.Sprintf("/meshes/%s/timeouts/timeout-all-%s/_resources/dataplanes", meshName, meshName))
				g.Expect(err).ToNot(HaveOccurred())
				defer r.Body.Close()
				g.Expect(r).To(HaveHTTPStatus(200))

				body, err := io.ReadAll(r.Body)
				g.Expect(err).ToNot(HaveOccurred())
				result := types.InspectDataplanesForPolicyResponse{}
				g.Expect(json.Unmarshal(body, &result)).To(Succeed())

				g.Expect(result.Items).To(HaveLen(1))
				g.Expect(result.Total).To(Equal(1))
				g.Expect(result.Items[0].Name).To(HavePrefix("kuma-4.test-server"))
			}, "30s", "1s").Should(Succeed())
		})

		It("should execute inspect rules of dataplane", func() {
			Expect(YamlUniversal(fmt.Sprintf(`
type: MeshTimeout
name: mt1
mesh: %s
spec:
  targetRef:
    kind: Mesh
  to:
    - targetRef:
        kind: Mesh
      default:
        idleTimeout: 20s
        http:
          requestTimeout: 2s
          maxStreamDuration: 20s`, meshName))(multizone.Global)).To(Succeed())
			Eventually(func(g Gomega) {
				clientDp := "kuma-4.test-server"
				r, err := http.Get(multizone.Global.GetKuma().GetAPIServerAddress() + fmt.Sprintf("/meshes/%s/dataplanes/%s/_rules", meshName, clientDp))
				g.Expect(err).ToNot(HaveOccurred())
				defer r.Body.Close()
				g.Expect(r).To(HaveHTTPStatus(200))

				body, err := io.ReadAll(r.Body)
				g.Expect(err).ToNot(HaveOccurred())
				result := types.InspectRulesForDataplaneResponse{}
				g.Expect(json.Unmarshal(body, &result)).To(Succeed())

				g.Expect(result.Resource.Name).To(Equal(clientDp))
				g.Expect(result.Rules).ToNot(BeEmpty())
				for _, rule := range result.Rules {
					if rule.Type == "MeshTimeout" {
						g.Expect(rule.ToRules).ToNot(BeNil())
						g.Expect(*rule.ToRules).ToNot(BeEmpty())
						g.Expect((*rule.ToRules)[0].Origin[0].Name).To(ContainSubstring("mt1"))
					}
				}
			}, "30s", "1s").Should(Succeed())
		})
	})
}
