// Package managerfactory re-exports the interfaces from apiserver/pkg/managerfactory.
package managerfactory

import "github.com/devsy-org/apiserver/pkg/managerfactory"

// SharedManagerFactory is the interface for retrieving managers.
type SharedManagerFactory = managerfactory.SharedManagerFactory

// ClusterClientAccess holds the functions for cluster access.
type ClusterClientAccess = managerfactory.ClusterClientAccess

// ManagementClientAccess holds the functions for management access.
type ManagementClientAccess = managerfactory.ManagementClientAccess
