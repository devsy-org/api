# devsy-org/api

API type definitions and generated clients for the Devsy platform.

## Code Generation

All generated code (`zz_generated.*`, `pkg/clientset/`, `pkg/listers/`, `pkg/informers/`, `pkg/openapi/`) is produced by `hack/generate.go`.

### Prerequisites

- Go 1.25+
- [Task](https://taskfile.dev) (`go install github.com/go-task/task/v3/cmd/task@latest`)

### Run all generators

```sh
task generate
```

### Run individual generators

```sh
task generate:register    # API registration (zz_generated.api.register.go)
task generate:deepcopy    # DeepCopy methods
task generate:defaults    # Defaulter functions
task generate:conversion  # Conversion functions
task generate:openapi     # OpenAPI definitions
task generate:clients     # Clientset, listers, informers
```

### Install tools only

```sh
task generate:install
```

### Verify generation is up to date

```sh
task verify
```

## Generator Architecture

`hack/generate.go` is a Go program that orchestrates two categories of generators:

1. **API Register Generator** — Uses `github.com/devsy-org/apiserver/pkg/generate` to scan all `+resource`-annotated types under `pkg/apis/...` and produces `zz_generated.api.register.go` files containing internal (hub) types, scheme registration, storage wiring, and registry interfaces.

2. **k8s.io/code-generator tools** — Standard Kubernetes code generators invoked as subprocesses:
   - `deepcopy-gen` — `DeepCopyInto` / `DeepCopyObject` methods
   - `defaulter-gen` — `SetDefaults_*` functions
   - `conversion-gen` — `Convert_*` functions between versioned and internal types
   - `client-gen` — Typed clientset
   - `lister-gen` — Listers for informer caches
   - `informer-gen` — SharedInformerFactory and informers
   - `openapi-gen` — OpenAPI v2 schema definitions (from `k8s.io/kube-openapi`)
