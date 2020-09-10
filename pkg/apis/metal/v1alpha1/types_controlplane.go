package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControlPlaneConfig contains configuration settings for the control plane.
type ControlPlaneConfig struct {
	metav1.TypeMeta `json:",inline"`

	// CloudControllerManager contains configuration settings for the cloud-controller-manager.
	// +optional
	CloudControllerManager *CloudControllerManagerConfig `json:"cloudControllerManager,omitempty"`

	// IAMConfig contains the config for all AuthN/AuthZ related components
	IAMConfig *IAMConfig `json:"iamconfig" optional:"false"`
}

// CloudControllerManagerConfig contains configuration settings for the cloud-controller-manager.
type CloudControllerManagerConfig struct {
	// FeatureGates contains information about enabled feature gates.
	// +optional
	FeatureGates map[string]bool `json:"featureGates,omitempty"`
	// DefaultExternalNetwork explicitly defines the network from which the CCM allocates IPs for services of type load balancer
	// If not defined, it will use the first network with the default external network tag from the infrastructure firewall networks
	// +optional
	DefaultExternalNetwork *string `json:"defaultExternalNetwork" optional:"true"`
}
