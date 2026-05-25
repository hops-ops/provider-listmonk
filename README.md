# Provider Upjet OpenPanel

`provider-listmonk` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
[OpenPanel](https://listmonk.com/) API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/hops-ops/provider-listmonk):
```
up ctp provider install hops-ops/provider-listmonk:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-listmonk
spec:
  package: xpkg.upbound.io/hops-ops/provider-listmonk:v0.1.0
EOF
```

You can see the API reference [here](https://doc.crds.dev/github.com/hops-ops/provider-listmonk).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/hops-ops/provider-listmonk/issues).
