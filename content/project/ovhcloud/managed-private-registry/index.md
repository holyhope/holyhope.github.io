---
title: "Managed Private Registry"
date: 2021-04-21T10:40:53+02:00
---

Managed Public Registry is a product of [OVHcloud Public Cloud catalog](https://www.ovhcloud.com/en/public-cloud/).
Deploy on demand a dedicated [Docker Registry](https://docs.docker.com/registry/) in a specified [OpenStack project](https://www.ovhcloud.com/en/public-cloud/project-management/).

## Goal

Add a new service to the [OVHcloud Public Cloud catalog](https://www.ovhcloud.com/en/public-cloud/).

## My role

From project manager, to architect, I was a 1-man team at the begining.
Then I onboarded more people, developped an [operator](../../harbor-operator) and managed ~1000 registries on producion.

## Timeline

### Bootstrap

1. December 2018 — Start

   Propose a dedicated private registry.

2. February 2019 — Project approved

   Go for [Harbor](#harbor).

3. March 2019 — Proof Of Concept

   API which deploys/deletes a set of kubernetes resources.

### Beta

1. June 2019 — Public Beta opening.

   Endpoint on [APIv6](https://api.ovh.com) to deploy/delete a bunch of kubernetes resources.

### Production

1. January 2020 — Production ready

   Global Availability.

2. March 2021 — Harbor Update

   Update to 2.0.

## Technical side

[Harbor](https://goharbor.io) was chosen as product after an [analysis of existing solutions]({{< ref "/blog/existing-docker-registry-solutions" >}}).

### Kubernetes & Operator
