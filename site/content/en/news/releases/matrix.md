---
title: Compatibility Matrix
description: This section includes Compatibility Matrix of Envoy Gateway.
---

Envoy Gateway relies on the Envoy Proxy and the Gateway API, and runs within a Kubernetes cluster. Not all versions of each of these products can function together for Envoy Gateway. Supported version combinations are listed below; **bold** type indicates the versions of the Envoy Proxy and the Gateway API actually compiled into each Envoy Gateway release.

| Envoy Gateway version | Envoy Proxy version         | Rate Limit version | Gateway API version | Kubernetes version         |  End of Life |
|-----------------------|-----------------------------|--------------------|---------------------|----------------------------|--------------|
| latest                | **dev-latest**              | **master**         | **v1.3.0**          | v1.30, v1.31, v1.32, v1.33 |  n/a         |
| v1.4                  | **distroless-v1.34.1**      | **3e085e5b**       | **v1.3.0**          | v1.30, v1.31, v1.32, v1.33 |  2025/11/13  |
| v1.3                  | **distroless-v1.33.0**      | **60d8e81b**       | **v1.2.1**          | v1.29, v1.30, v1.31, v1.32 |  2025/07/30  |
| v1.2                  | **distroless-v1.32.1**      | **28b1629a**       | **v1.2.0**          | v1.28, v1.29, v1.30, v1.31 |  2025/05/06  |
| v1.1                  | **distroless-v1.31.0**      | **91484c59**       | **v1.1.0**          | v1.27, v1.28, v1.29, v1.30 |  2025/01/22  |
| v1.0                  | **distroless-v1.29.2**      | **19f2079f**       | **v1.0.0**          | v1.26, v1.27, v1.28, v1.29 |  2024/09/13  |
| v0.6                  | **distroless-v1.28-latest** | **b9796237**       | **v1.0.0**          | v1.26, v1.27, v1.28        |  2024/05/02  |
| v0.5                  | **v1.27-latest**            | **e059638d**       | **v0.7.1**          | v1.25, v1.26, v1.27        |  2024/01/02  |
| v0.4                  | **v1.26-latest**            | **542a6047**       | **v0.6.2**          | v1.25, v1.26, v1.27        |  2023/10/24  |
| v0.3                  | **v1.25-latest**            | **f28024e3**       | **v0.6.1**          | v1.24, v1.25, v1.26        |  2023/08/09  |
| v0.2                  | **v1.23-latest**            |                    | **v0.5.1**          | v1.24                      |  2023/04/20  |
