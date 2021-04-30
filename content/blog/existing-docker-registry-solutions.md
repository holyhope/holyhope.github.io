---
title: "Existing Solutions"
date: 2019-12-10T10:00:00+02:00
draft: false
name: existing-solutions
---

About the [Private Managed Registry]({{< ref "/project/ovhcloud/managed-private-registry" >}}) project, here is a landscape of the different solutions on the market on December 2019.

Attention following information were collected long time ago and are probably obsolete.

## Official Docker registry

<https://docs.docker.com/registry/>

Strengths | Weaknesses | Opportunities | Threats
----------|------------|---------------|--------
Official docker registry.{{< br >}} Compliant [APIv2](https://docs.docker.com/registry/spec/api/).{{< br >}} Cheap infrastructure.{{< br >}} Prod ready.{{< br >}} Integration with Swift storage. | Few features. | Can add custom frontend and other features manually.{{< br >}} Authorization are delegated. | Part of the [Docker distribution](https://github.com/docker/distribution).

## Quay enterprise

<https://quay.io>

Strengths | Weaknesses | Opportunities | Threats
----------|------------|---------------|--------
Supported by CoreOS - RedHat.{{< br >}} Pull via p2p.{{< br >}} Security scanning.{{< br >}} APIv2 supported.{{< br >}} HA: multi-region.{{< br >}} LDAP authentication. | Close sources: licences required.{{< br >}} Recommended with Tectonic. |  | Close sources.{{< br >}} May be long to POC.{{< br >}} May be hard to integrate with OVHcloud ecosystem.{{< br >}} Specific cli.

## JFrog

<https://jfrog.com/>

Strengths | Weaknesses | Opportunities | Threats
----------|------------|---------------|--------
Prod ready (used by other competitors).{{< br >}} Can hosts helm charts. | Close sources: licences required | Can hosts many kind of artifact puppet, maven, ... | Close sources.{{< br >}} May be long to POC.{{< br >}} May be hard to integrate with OVHcloud ecosystem.

## Portus

<http://port.us.org>

Strengths | Weaknesses | Opportunities | Threats
----------|------------|---------------|--------
Security: scans container images for vulnerabilities.{{< br >}} Cloud native.{{< br >}} Multi-tenant.{{< br >}} UI provided: Oauth and OpenID authentication.{{< br >}} Built on top of Docker registry.{{< br >}} Good documentation. |  | Ready to deploy on Kubernetes: Helm Chart available.{{< br >}} content-signing and validation. | May be hard to provide all features to clients nicely: multi tenant ui?

## Harbor

<https://goharbor/io>

Strengths | Weaknesses | Opportunities | Threats
----------|------------|---------------|--------
Cloud native.{{< br >}} Multi-tenant.{{< br >}} content-signing and validation.{{< br >}} Extensible API : To dig.{{< br >}} UI provided.{{< br >}} Builded on top of Docker registry. |  | Ready to deploy on Kubernetes: Helm Chart available.{{< br >}} Supported by [CNCF](https://www.cncf.io).{{< br >}} Security: scans container images for vulnerabilities. | May be hard to provide all features to clients nicely: multi tenant ui?

## Popularity

{{< chartjs/horizontal-bar title="Projects popularity" height=200 >}}
Name,Star,Fork,Watch
Portus,2148,396,115
"Official docker registry",4301,1424,281
Harbor,6138,1695,394
{{< /chartjs/horizontal-bar >}}
It was very important to take care of the popularity of each projects.
Popularities were computed easily but roughly, thanks to GitHub interactions.

Closed source projects popularity could not easily be compared to others.
That is why [Quay enterprise](#quay-enterprise) and [JFrog](#jfrog) are not in the chart.
