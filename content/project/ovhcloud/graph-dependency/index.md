---
title: "Graph dependency"
date: 2021-04-21T10:40:53+02:00
---

## Goal

Obviously the goal of a patent is to protect innovation.

## My role

I am the **inventor**, I designed the solution and communicated how it works to the team.

## Context

Designed during development of the [Managed Public Registry project](../managed-private-registry).
A first implementation of the innovation was open-sourced before submiting the patent to the European commission, so the patent could only be send to the US one.

Note: [OVHcloud](https://www.ovhcloud.com) helped me to write and submit the final patent and we were able to works only with EU and US.

## Technical side

The innovation is strongly linked to the CLoud Native world: being able to keep a set of resources consistent and in the specified state.

Using Kubernetes features, it means ensuring that any resources depending on others are updated (and restarted in case of pods, deployments, daemonsets...)

{{< mxgraph title="design" src="sample.xml" />}}
