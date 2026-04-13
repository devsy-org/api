// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/skevetter/api/pkg/apis/virtualcluster
// +k8s:defaulter-gen=TypeMeta
// +groupName=virtualcluster.devsy.sh
package v1 // import "github.com/skevetter/api/pkg/apis/virtualcluster/v1"
