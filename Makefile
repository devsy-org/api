.PHONY: generate generate-register generate-deepcopy generate-defaults generate-conversion generate-openapi generate-clients install-tools

# Run all code generators
generate: install-tools generate-register generate-deepcopy generate-defaults generate-conversion generate-openapi generate-clients

# Install k8s.io/code-generator tools
install-tools:
	./hack/generate.sh install

# Generate API registration code (custom apiserver generator)
generate-register:
	./hack/generate.sh register

# Generate DeepCopy methods
generate-deepcopy:
	./hack/generate.sh deepcopy

# Generate Defaulter functions
generate-defaults:
	./hack/generate.sh defaults

# Generate Conversion functions
generate-conversion:
	./hack/generate.sh conversion

# Generate OpenAPI definitions
generate-openapi:
	./hack/generate.sh openapi

# Generate clientset, listers, informers
generate-clients:
	./hack/generate.sh clients
