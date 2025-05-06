// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

//go:build e2e

package tests

import (
	"context"
	"strings"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/gateway-api/conformance/utils/http"
	"sigs.k8s.io/gateway-api/conformance/utils/kubernetes"
	"sigs.k8s.io/gateway-api/conformance/utils/suite"
)

func init() {
	GatewayNamespaceModeTests = append(GatewayNamespaceModeTests, GatewayNamespaceModeTest)
}

var GatewayNamespaceModeTest = suite.ConformanceTest{
	ShortName:   "BasicGatewayNamespaceMode",
	Description: "Basic test for Gateway Namespace Mode feature",
	Manifests:   []string{"testdata/basic-gateway-namespace-mode.yaml"},
	Test: func(t *testing.T, suite *suite.ConformanceTestSuite) {
		ns1 := "ns-1"
		routeNN := types.NamespacedName{Name: "route-ns-1", Namespace: ns1}
		gwNN := types.NamespacedName{Name: "gateway-ns-1", Namespace: ns1}
		gwAddr := kubernetes.GatewayAndHTTPRoutesMustBeAccepted(t, suite.Client, suite.TimeoutConfig, suite.ControllerName, kubernetes.NewGatewayRef(gwNN), routeNN)
		OkResp := http.ExpectedResponse{
			Request: http.Request{
				Path: "/",
			},
			Response: http.Response{
				StatusCode: 200,
			},
			Namespace: ns1,
		}

		// Send a request to an valid path and expect a successful response
		http.MakeRequestAndExpectEventuallyConsistentResponse(t, suite.RoundTripper, suite.TimeoutConfig, gwAddr, OkResp)
		// Check if the Envoy pod exists in the second namespace
		checkEnvoyPodExistsInInfraNs(t, suite.Client, ns1)

		ns2 := "ns-2"
		routeNN2 := types.NamespacedName{Name: "route-ns-2", Namespace: ns2}
		gwNN2 := types.NamespacedName{Name: "gateway-ns-2", Namespace: ns2}
		gwAddr2 := kubernetes.GatewayAndHTTPRoutesMustBeAccepted(t, suite.Client, suite.TimeoutConfig, suite.ControllerName, kubernetes.NewGatewayRef(gwNN2), routeNN2)
		OkResp = http.ExpectedResponse{
			Request: http.Request{
				Path: "/",
			},
			Response: http.Response{
				StatusCode: 200,
			},
			Namespace: ns2,
		}

		// Send a request to an valid path and expect a successful response
		http.MakeRequestAndExpectEventuallyConsistentResponse(t, suite.RoundTripper, suite.TimeoutConfig, gwAddr2, OkResp)
		// Check if the Envoy pod exists in the second namespace
		checkEnvoyPodExistsInInfraNs(t, suite.Client, ns2)
	},
}

func checkEnvoyPodExistsInInfraNs(t *testing.T, k8sClient client.Client, namespace string) {
	t.Helper()
	podList := &corev1.PodList{}
	err := k8sClient.List(context.TODO(), podList, client.InNamespace(namespace))
	if err != nil {
		t.Fatalf("failed to list pods in namespace %s: %v", namespace, err)
	}

	for _, pod := range podList.Items {
		if strings.Contains(pod.Name, "envoy") {
			t.Logf("Envoy pod found in namespace %s: %s", namespace, pod.Name)
			return
		}
	}

	t.Fatalf("no Envoy pod found in namespace %s", namespace)
}
