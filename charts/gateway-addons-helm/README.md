# gateway-addons-helm

![Version: v0.0.0-latest](https://img.shields.io/badge/Version-v0.0.0--latest-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: latest](https://img.shields.io/badge/AppVersion-latest-informational?style=flat-square)

An Add-ons Helm chart for Envoy Gateway

**Homepage:** <https://gateway.envoyproxy.io/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| envoy-gateway-steering-committee |  | <https://github.com/envoyproxy/gateway/blob/main/GOVERNANCE.md> |
| envoy-gateway-maintainers |  | <https://github.com/envoyproxy/gateway/blob/main/CODEOWNERS> |

## Source Code

* <https://github.com/envoyproxy/gateway>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://grafana.github.io/helm-charts | grafana | 8.0.0 |
| https://prometheus-community.github.io/helm-charts | prometheus | 25.21.0 |

## Usage

[Helm](https://helm.sh) must be installed to use the charts.
Please refer to Helm's [documentation](https://helm.sh/docs) to get started.

The Envoy Gateway must be installed before installing this chart.

### Install from DockerHub

Once Helm has been set up correctly, install the chart from dockerhub:

``` shell
    helm install eg-addons oci://docker.io/envoyproxy/gateway-addons-helm --version v0.0.0-latest -n monitoring --create-namespace
```

You can find all helm chart release in [Dockerhub](https://hub.docker.com/r/envoyproxy/gateway-addons-helm/tags)

To uninstall the chart:

``` shell
    helm uninstall eg-addons -n monitoring
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| grafana.adminPassword | string | `"admin"` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".apiVersion | int | `1` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].disableDeletion | bool | `false` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].editable | bool | `true` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].folder | string | `"envoy-gateway"` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].name | string | `"envoy-gateway"` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].options.path | string | `"/var/lib/grafana/dashboards/envoy-gateway"` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].orgId | int | `1` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].type | string | `"file"` |  |
| grafana.dashboardsConfigMaps.envoy-gateway | string | `"grafana-dashboards"` |  |
| grafana.datasources."datasources.yaml".apiVersion | int | `1` |  |
| grafana.datasources."datasources.yaml".datasources[0].name | string | `"Prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[0].type | string | `"prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[0].url | string | `"http://prometheus-server"` |  |
| grafana.fullnameOverride | string | `"grafana"` |  |
| grafana.service.type | string | `"LoadBalancer"` |  |
| prometheus.alertmanager.enabled | bool | `false` |  |
| prometheus.kube-state-metrics.enabled | bool | `false` |  |
| prometheus.prometheus-node-exporter.enabled | bool | `false` |  |
| prometheus.prometheus-pushgateway.enabled | bool | `false` |  |
| prometheus.server.fullnameOverride | string | `"prometheus-server"` |  |
| prometheus.server.global.scrape_interval | string | `"15s"` |  |
| prometheus.server.image.repository | string | `"prom/prometheus"` |  |
| prometheus.server.persistentVolume.enabled | bool | `false` |  |
| prometheus.server.readinessProbeInitialDelay | int | `0` |  |
| prometheus.server.securityContext | object | `{}` |  |
| prometheus.server.service.type | string | `"LoadBalancer"` |  |

