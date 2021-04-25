---
title: "Harbor Operator"
date: 2021-04-21T17:48:23+02:00
---

A [Kubernetes operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) to manage Harbor application.

I developped the Harbor Operator for [OVHcloud](https://ovhcloud.com) needs [when I was employee]({{< ref "/project/ovhcloud" >}}).

<https://github.com/goharbor/harbor-operator/>

## Goal

Manage hundreds of [Harbor](https://goharbor.io) at scale.

## Technical side

A Kubernetes operator is an application that control resource life cycle.

### Architecture

### Requests flow

{{< mxgraph title="design" src="design-flow.xml" />}}

### Kubernetes Resources

Harbor Operator manages multiple resources which represents running application:

1. Docker registry
2. Chartmuseum
3. Notary
4. Trivy
5. Harbor components (Core, JobService, Portal, RegistryController)

The operator can also manage some extra resources:

1. Harbor. A full Harbor deployment: all previous software configured to communicate with each others.
2. HarborCluster. A full Harbor deployment with some external resources to deploy databases and redis.

The operator uses external resources to do the job:

1. Certificate and Issuer from CertManager.
2. Postgresql from PostgresqlOperator.
3. Redis from RedisOperator.

Here is a sample of an Kubernetes resource that describe a Harbor.

{{< highlight yaml >}}
apiVersion: goharbor.io/v1alpha3
kind: HarborCluster
metadata:
  name: my-harbor
spec:
  ...
{{< /highlight >}}
