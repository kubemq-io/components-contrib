/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Pod Security Policy Spec defines the policy enforced.
type ExtensionsV1beta1PodSecurityPolicySpec struct {

	// AllowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`

	// AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities.
	AllowedCapabilities []string `json:"allowedCapabilities,omitempty"`

	// AllowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the \"Volumes\" field.
	AllowedFlexVolumes []ExtensionsV1beta1AllowedFlexVolume `json:"allowedFlexVolumes,omitempty"`

	// is a white list of allowed host paths. Empty indicates that all host paths may be used.
	AllowedHostPaths []ExtensionsV1beta1AllowedHostPath `json:"allowedHostPaths,omitempty"`

	// DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both DefaultAddCapabilities and RequiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the AllowedCapabilities list.
	DefaultAddCapabilities []string `json:"defaultAddCapabilities,omitempty"`

	// DefaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
	DefaultAllowPrivilegeEscalation bool `json:"defaultAllowPrivilegeEscalation,omitempty"`

	// FSGroup is the strategy that will dictate what fs group is used by the SecurityContext.
	FsGroup *ExtensionsV1beta1FsGroupStrategyOptions `json:"fsGroup"`

	// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	HostIPC bool `json:"hostIPC,omitempty"`

	// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	HostNetwork bool `json:"hostNetwork,omitempty"`

	// hostPID determines if the policy allows the use of HostPID in the pod spec.
	HostPID bool `json:"hostPID,omitempty"`

	// hostPorts determines which host port ranges are allowed to be exposed.
	HostPorts []ExtensionsV1beta1HostPortRange `json:"hostPorts,omitempty"`

	// privileged determines if a pod can request to be run as privileged.
	Privileged bool `json:"privileged,omitempty"`

	// ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
	ReadOnlyRootFilesystem bool `json:"readOnlyRootFilesystem,omitempty"`

	// RequiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.
	RequiredDropCapabilities []string `json:"requiredDropCapabilities,omitempty"`

	// runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
	RunAsUser *ExtensionsV1beta1RunAsUserStrategyOptions `json:"runAsUser"`

	// seLinux is the strategy that will dictate the allowable labels that may be set.
	SeLinux *ExtensionsV1beta1SeLinuxStrategyOptions `json:"seLinux"`

	// SupplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
	SupplementalGroups *ExtensionsV1beta1SupplementalGroupsStrategyOptions `json:"supplementalGroups"`

	// volumes is a white list of allowed volume plugins.  Empty indicates that all plugins may be used.
	Volumes []string `json:"volumes,omitempty"`
}