package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"istio.io/istio/pkg/test/echo/common/scheme"

	"github.com/solo-io/gloo-mesh/pkg/test/apps/context"
	echo2 "github.com/solo-io/gloo-mesh/pkg/test/apps/echo"
	gloo_mesh "github.com/solo-io/gloo-mesh/pkg/test/apps/gloo-mesh"
	"github.com/solo-io/gloo-mesh/pkg/test/common"
	"istio.io/istio/pkg/test/framework/components/echo"
	"istio.io/istio/pkg/test/framework/components/environment/kube"
	"istio.io/istio/pkg/test/framework/config"
	"istio.io/istio/pkg/test/framework/resource"

	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/istio"
)

var (
	i             istio.Instance
	deploymentCtx context.DeploymentContext
)

func TestMain(m *testing.M) {
	if os.Getenv("RUN_INTEGRATION") == "" {
		fmt.Println("skipping E2E Integration tests")
		return
	}
	licenseKey := os.Getenv("GLOO_MESH_LICENSE_KEY")
	// get kube settings from command line
	config.Parse()
	kubeSettings, _ := kube.NewSettingsFromCommandLine()
	clusterKubeConfigs := make(map[string]string)

	// this is a hack to match the kube configs with the cluster names so we can match them when deploy happens
	for i, k := range kubeSettings.KubeConfig {
		clusterKubeConfigs[fmt.Sprintf("cluster-%d", i)] = k
	}

	framework.
		NewSuite(m).
		RequireMinClusters(2).
		Setup(istio.Setup(&i, common.IstioSetupFunc("gloo-mesh-istio.yaml"))).
		Setup(gloo_mesh.Deploy(&deploymentCtx, &gloo_mesh.Config{
			ClusterKubeConfigs:                  clusterKubeConfigs,
			DeployControlPlaneToManagementPlane: true,
		},
			licenseKey)).
		Setup(echo2.DeployEchos(&deploymentCtx)).
		Run()
}
func TestInMesh(t *testing.T) {
	framework.
		NewTest(t).
		Run(func(ctx framework.TestContext) {

			tgs := []common.TestGroup{
				{
					Name: "virtual-destination",
					Cases: []common.TestCase{
						{
							Name:        "same-cluster-http",
							Description: "Testing http routing in same cluster",
							Test:        testGlobalVirtualDestinationHTTP,
							Namespace:   deploymentCtx.EchoContext.AppNamespace.Name(),
							FileName:    "virtual-destination-http.yaml",
							Folder:      "gloo-mesh/in-mesh",
						},
						{
							Name:        "same-cluster-https",
							Description: "Testing https routing in same cluster",
							Test:        testGlobalVirtualDestinationHTTPS,
							Namespace:   deploymentCtx.EchoContext.AppNamespace.Name(),
							FileName:    "virtual-destination-https.yaml",
							Folder:      "gloo-mesh/in-mesh",
						},
						{
							Name:        "same-cluster-tcp",
							Description: "Testing tcp routing in same cluster",
							Test:        testGlobalVirtualDestinationTCP,
							Namespace:   deploymentCtx.EchoContext.AppNamespace.Name(),
							FileName:    "virtual-destination-tcp.yaml",
							Folder:      "gloo-mesh/in-mesh",
						},
					},
				},
			}
			for _, tg := range tgs {
				tg.Run(ctx, t, &deploymentCtx)
			}
		})

}

// testGlobalVirtualDestinationHTTP making http requests for a virtual destination
// because of locality priority routing, we should see routing to local cluster first always
func testGlobalVirtualDestinationHTTP(ctx resource.Context, t *testing.T, deploymentCtx *context.DeploymentContext) {
	cluster := ctx.Clusters()[0]
	// frontend calling backend in mesh using virtual destination in same cluster
	src := deploymentCtx.EchoContext.Deployments.GetOrFail(t, echo.Service("frontend").And(echo.InCluster(cluster)))
	backendHost := "http-backend.solo.io"

	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "http",
			ServicePort: 8090,
		},
		Scheme:    scheme.HTTP,
		Address:   backendHost,
		Method:    http.MethodGet,
		Path:      "",
		Count:     5,
		Validator: echo.And(echo.ExpectOK(), echo.ExpectCluster(cluster.Name())),
	})
	// cluster 2 test
	cluster = ctx.Clusters()[1]
	// frontend calling backend in mesh using virtual destination in same cluster
	src = deploymentCtx.EchoContext.Deployments.GetOrFail(t, echo.Service("frontend").And(echo.InCluster(cluster)))
	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "http",
			ServicePort: 8090,
		},
		Scheme:    scheme.HTTP,
		Address:   backendHost,
		Method:    http.MethodGet,
		Path:      "",
		Count:     5,
		Validator: echo.And(echo.ExpectOK(), echo.ExpectCluster(cluster.Name())),
	})
}

// testGlobalVirtualDestinationHTTPS making https requests for a virtual destination
// because of locality priority routing, we should see routing to local cluster first always
func testGlobalVirtualDestinationHTTPS(ctx resource.Context, t *testing.T, deploymentCtx *context.DeploymentContext) {
	cluster := ctx.Clusters()[0]
	// frontend calling backend in mesh using virtual destination in same cluster
	src := deploymentCtx.EchoContext.Deployments.GetOrFail(t, echo.Service("frontend").And(echo.InCluster(cluster)))
	backendHost := "https-backend.solo.io"

	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "https",
			ServicePort: 9443,
			TLS:         true,
		},
		Scheme:    scheme.HTTPS,
		Address:   backendHost,
		Method:    http.MethodGet,
		Path:      "",
		Count:     5,
		Validator: echo.And(echo.ExpectOK(), echo.ExpectCluster(cluster.Name())),
		CaCert:    echo2.GetEchoCACert(),
	})

	cluster = ctx.Clusters()[1]
	// frontend calling backend in mesh using virtual destination in same cluster
	src = deploymentCtx.EchoContext.Deployments.GetOrFail(t, echo.Service("frontend").And(echo.InCluster(cluster)))

	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "https",
			ServicePort: 9443,
			TLS:         true,
		},
		Scheme:    scheme.HTTPS,
		Address:   backendHost,
		Method:    http.MethodGet,
		Path:      "",
		Count:     5,
		Validator: echo.And(echo.ExpectOK(), echo.ExpectCluster(cluster.Name())),
		CaCert:    echo2.GetEchoCACert(),
	})
}

// testGlobalVirtualDestinationTCP making tcp requests for a virtual destination
// because of locality priority routing, we should see routing to local cluster first always
func testGlobalVirtualDestinationTCP(ctx resource.Context, t *testing.T, deploymentCtx *context.DeploymentContext) {
	cluster := ctx.Clusters()[0]
	// frontend calling backend in mesh using virtual destination in same cluster
	src := deploymentCtx.EchoContext.Deployments.GetOrFail(t, echo.Service("frontend").And(echo.InCluster(cluster)))
	backendHost := "tcp-backend.solo.io"

	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "tcp",
			ServicePort: 9000,
		},
		Scheme:    scheme.TCP,
		Address:   backendHost,
		Count:     5,
		Validator: echo.And(echo.ExpectOK(), echo.ExpectCluster(cluster.Name())),
	})

	cluster = ctx.Clusters()[1]
	// frontend calling backend in mesh using virtual destination in same cluster
	src = deploymentCtx.EchoContext.Deployments.GetOrFail(t, echo.Service("frontend").And(echo.InCluster(cluster)))
	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "tcp",
			ServicePort: 9000,
		},
		Scheme:    scheme.TCP,
		Address:   backendHost,
		Count:     5,
		Validator: echo.And(echo.ExpectOK(), echo.ExpectCluster(cluster.Name())),
	})
}