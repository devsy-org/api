#!/usr/bin/env bash
set -euo pipefail

# Code generation script for devsy-org/api.
# Regenerates all generated code from source types.
#
# Usage:
#   ./hack/generate.sh          # Run all generators
#   ./hack/generate.sh register # Run only apiregister-gen
#   ./hack/generate.sh clients  # Run only client/lister/informer-gen

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${REPO_ROOT}"

# Ensure GOBIN is on PATH
export PATH="${GOBIN:-$(go env GOPATH)/bin}:${PATH}"

MODULE="github.com/devsy-org/api"
BOILERPLATE="${REPO_ROOT}/hack/boilerplate.go.txt"

# API packages (input-dirs style)
API_PACKAGES=(
  "${MODULE}/pkg/apis/audit/v1"
  "${MODULE}/pkg/apis/management"
  "${MODULE}/pkg/apis/management/v1"
  "${MODULE}/pkg/apis/storage/v1"
  "${MODULE}/pkg/apis/ui"
  "${MODULE}/pkg/apis/ui/v1"
  "${MODULE}/pkg/apis/virtualcluster"
  "${MODULE}/pkg/apis/virtualcluster/v1"
)
# Versioned API packages for client-gen
CLIENT_INPUT_DIRS="management/v1,storage/v1,virtualcluster/v1"

# Conversion packages
CONVERSION_PACKAGES=(
  "${MODULE}/pkg/apis/management/v1"
  "${MODULE}/pkg/apis/virtualcluster/v1"
)

# OpenAPI extra input packages
OPENAPI_EXTRA=(
  "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/apimachinery/pkg/api/resource"
  "k8s.io/apimachinery/pkg/version"
  "k8s.io/apimachinery/pkg/runtime"
  "k8s.io/apimachinery/pkg/util/intstr"
  "k8s.io/api/core/v1"
  "k8s.io/api/rbac/v1"
  "k8s.io/api/apps/v1"
  "k8s.io/api/networking/v1"
  "k8s.io/api/storage/v1"
  "k8s.io/api/batch/v1"
)

CODEGEN_VERSION="v0.35.3"

install_tools() {
  echo "Installing code-generator tools..."
  go install "k8s.io/code-generator/cmd/deepcopy-gen@${CODEGEN_VERSION}"
  go install "k8s.io/code-generator/cmd/defaulter-gen@${CODEGEN_VERSION}"
  go install "k8s.io/code-generator/cmd/conversion-gen@${CODEGEN_VERSION}"
  go install "k8s.io/code-generator/cmd/client-gen@${CODEGEN_VERSION}"
  go install "k8s.io/code-generator/cmd/lister-gen@${CODEGEN_VERSION}"
  go install "k8s.io/code-generator/cmd/informer-gen@${CODEGEN_VERSION}"
  go install k8s.io/kube-openapi/cmd/openapi-gen@latest
}

generate_register() {
  echo "==> Generating API register (apiserver-gen)..."
  go run ./hack/gen/main.go
}

generate_deepcopy() {
  echo "==> Generating deepcopy..."
  deepcopy-gen \
    --go-header-file "${BOILERPLATE}" \
    --output-file zz_generated.deepcopy.go \
    "${API_PACKAGES[@]}"
}

generate_defaults() {
  echo "==> Generating defaults..."
  defaulter-gen \
    --go-header-file "${BOILERPLATE}" \
    --output-file zz_generated.defaults.go \
    "${API_PACKAGES[@]}"
}

generate_conversion() {
  echo "==> Generating conversion..."
  conversion-gen \
    --go-header-file "${BOILERPLATE}" \
    --output-file zz_generated.conversion.go \
    "${CONVERSION_PACKAGES[@]}"
}

generate_openapi() {
  echo "==> Generating openapi..."
  openapi-gen \
    --go-header-file "${BOILERPLATE}" \
    --input-dirs "${API_PACKAGES[*]},${OPENAPI_EXTRA[*]}" \
    --output-package "${MODULE}/pkg/openapi" \
    --output-file zz_generated.openapi.go \
    --report-filename /dev/null
}

generate_clients() {
  echo "==> Generating clientset..."
  client-gen \
    --go-header-file "${BOILERPLATE}" \
    --input-base "${MODULE}/pkg/apis" \
    --input "${CLIENT_INPUT_DIRS}" \
    --output-package "${MODULE}/pkg/clientset" \
    --clientset-name versioned \
    --output-dir pkg/clientset

  echo "==> Generating listers..."
  lister-gen \
    --go-header-file "${BOILERPLATE}" \
    --output-package "${MODULE}/pkg/listers" \
    --output-dir pkg/listers \
    "${API_PACKAGES[@]}"

  echo "==> Generating informers..."
  informer-gen \
    --go-header-file "${BOILERPLATE}" \
    --output-package "${MODULE}/pkg/informers" \
    --output-dir pkg/informers \
    --versioned-clientset-package "${MODULE}/pkg/clientset/versioned" \
    --listers-package "${MODULE}/pkg/listers" \
    "${API_PACKAGES[@]}"
}

case "${1:-all}" in
  register)
    generate_register
    ;;
  deepcopy)
    generate_deepcopy
    ;;
  defaults)
    generate_defaults
    ;;
  conversion)
    generate_conversion
    ;;
  openapi)
    generate_openapi
    ;;
  clients)
    generate_clients
    ;;
  install)
    install_tools
    ;;
  all)
    install_tools
    generate_register
    generate_deepcopy
    generate_defaults
    generate_conversion
    generate_openapi
    generate_clients
    ;;
  *)
    echo "Usage: $0 {all|install|register|deepcopy|defaults|conversion|openapi|clients}"
    exit 1
    ;;
esac

echo "==> Done."
