//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
2024 Copyright metal-stack Authors.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	apisconfig "github.com/gardener/gardener/extensions/pkg/apis/config"
	apisconfigv1alpha1 "github.com/gardener/gardener/extensions/pkg/apis/config/v1alpha1"
	config "github.com/metal-stack/gardener-extension-provider-metal/pkg/apis/config"
	resource "k8s.io/apimachinery/pkg/api/resource"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfig "k8s.io/component-base/config"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AuditToSplunk)(nil), (*config.AuditToSplunk)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AuditToSplunk_To_config_AuditToSplunk(a.(*AuditToSplunk), b.(*config.AuditToSplunk), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AuditToSplunk)(nil), (*AuditToSplunk)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AuditToSplunk_To_v1alpha1_AuditToSplunk(a.(*config.AuditToSplunk), b.(*AuditToSplunk), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterAudit)(nil), (*config.ClusterAudit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClusterAudit_To_config_ClusterAudit(a.(*ClusterAudit), b.(*config.ClusterAudit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ClusterAudit)(nil), (*ClusterAudit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ClusterAudit_To_v1alpha1_ClusterAudit(a.(*config.ClusterAudit), b.(*ClusterAudit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerConfiguration)(nil), (*config.ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(a.(*ControllerConfiguration), b.(*config.ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControllerConfiguration)(nil), (*ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(a.(*config.ControllerConfiguration), b.(*ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DurosConfiguration)(nil), (*config.DurosConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DurosConfiguration_To_config_DurosConfiguration(a.(*DurosConfiguration), b.(*config.DurosConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DurosConfiguration)(nil), (*DurosConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DurosConfiguration_To_v1alpha1_DurosConfiguration(a.(*config.DurosConfiguration), b.(*DurosConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DurosPartitionConfiguration)(nil), (*config.DurosPartitionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DurosPartitionConfiguration_To_config_DurosPartitionConfiguration(a.(*DurosPartitionConfiguration), b.(*config.DurosPartitionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DurosPartitionConfiguration)(nil), (*DurosPartitionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DurosPartitionConfiguration_To_v1alpha1_DurosPartitionConfiguration(a.(*config.DurosPartitionConfiguration), b.(*DurosPartitionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DurosSeedStorageClass)(nil), (*config.DurosSeedStorageClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DurosSeedStorageClass_To_config_DurosSeedStorageClass(a.(*DurosSeedStorageClass), b.(*config.DurosSeedStorageClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DurosSeedStorageClass)(nil), (*DurosSeedStorageClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DurosSeedStorageClass_To_v1alpha1_DurosSeedStorageClass(a.(*config.DurosSeedStorageClass), b.(*DurosSeedStorageClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCD)(nil), (*config.ETCD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCD_To_config_ETCD(a.(*ETCD), b.(*config.ETCD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCD)(nil), (*ETCD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCD_To_v1alpha1_ETCD(a.(*config.ETCD), b.(*ETCD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCDBackup)(nil), (*config.ETCDBackup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(a.(*ETCDBackup), b.(*config.ETCDBackup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCDBackup)(nil), (*ETCDBackup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(a.(*config.ETCDBackup), b.(*ETCDBackup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCDStorage)(nil), (*config.ETCDStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(a.(*ETCDStorage), b.(*config.ETCDStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCDStorage)(nil), (*ETCDStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(a.(*config.ETCDStorage), b.(*ETCDStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EgressDest)(nil), (*config.EgressDest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EgressDest_To_config_EgressDest(a.(*EgressDest), b.(*config.EgressDest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.EgressDest)(nil), (*EgressDest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_EgressDest_To_v1alpha1_EgressDest(a.(*config.EgressDest), b.(*EgressDest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ImagePullSecret)(nil), (*config.ImagePullSecret)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ImagePullSecret_To_config_ImagePullSecret(a.(*ImagePullSecret), b.(*config.ImagePullSecret), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ImagePullSecret)(nil), (*ImagePullSecret)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ImagePullSecret_To_v1alpha1_ImagePullSecret(a.(*config.ImagePullSecret), b.(*ImagePullSecret), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImage)(nil), (*config.MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImage_To_config_MachineImage(a.(*MachineImage), b.(*config.MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MachineImage)(nil), (*MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MachineImage_To_v1alpha1_MachineImage(a.(*config.MachineImage), b.(*MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*StorageConfiguration)(nil), (*config.StorageConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_StorageConfiguration_To_config_StorageConfiguration(a.(*StorageConfiguration), b.(*config.StorageConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.StorageConfiguration)(nil), (*StorageConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_StorageConfiguration_To_v1alpha1_StorageConfiguration(a.(*config.StorageConfiguration), b.(*StorageConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AuditToSplunk_To_config_AuditToSplunk(in *AuditToSplunk, out *config.AuditToSplunk, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.HECToken = in.HECToken
	out.Index = in.Index
	out.HECHost = in.HECHost
	out.HECPort = in.HECPort
	out.TLSEnabled = in.TLSEnabled
	out.HECCAFile = in.HECCAFile
	return nil
}

// Convert_v1alpha1_AuditToSplunk_To_config_AuditToSplunk is an autogenerated conversion function.
func Convert_v1alpha1_AuditToSplunk_To_config_AuditToSplunk(in *AuditToSplunk, out *config.AuditToSplunk, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditToSplunk_To_config_AuditToSplunk(in, out, s)
}

func autoConvert_config_AuditToSplunk_To_v1alpha1_AuditToSplunk(in *config.AuditToSplunk, out *AuditToSplunk, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.HECToken = in.HECToken
	out.Index = in.Index
	out.HECHost = in.HECHost
	out.HECPort = in.HECPort
	out.TLSEnabled = in.TLSEnabled
	out.HECCAFile = in.HECCAFile
	return nil
}

// Convert_config_AuditToSplunk_To_v1alpha1_AuditToSplunk is an autogenerated conversion function.
func Convert_config_AuditToSplunk_To_v1alpha1_AuditToSplunk(in *config.AuditToSplunk, out *AuditToSplunk, s conversion.Scope) error {
	return autoConvert_config_AuditToSplunk_To_v1alpha1_AuditToSplunk(in, out, s)
}

func autoConvert_v1alpha1_ClusterAudit_To_config_ClusterAudit(in *ClusterAudit, out *config.ClusterAudit, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_ClusterAudit_To_config_ClusterAudit is an autogenerated conversion function.
func Convert_v1alpha1_ClusterAudit_To_config_ClusterAudit(in *ClusterAudit, out *config.ClusterAudit, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterAudit_To_config_ClusterAudit(in, out, s)
}

func autoConvert_config_ClusterAudit_To_v1alpha1_ClusterAudit(in *config.ClusterAudit, out *ClusterAudit, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_config_ClusterAudit_To_v1alpha1_ClusterAudit is an autogenerated conversion function.
func Convert_config_ClusterAudit_To_v1alpha1_ClusterAudit(in *config.ClusterAudit, out *ClusterAudit, s conversion.Scope) error {
	return autoConvert_config_ClusterAudit_To_v1alpha1_ClusterAudit(in, out, s)
}

func autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	out.ClientConnection = (*componentbaseconfig.ClientConnectionConfiguration)(unsafe.Pointer(in.ClientConnection))
	out.MachineImages = *(*[]config.MachineImage)(unsafe.Pointer(&in.MachineImages))
	out.FirewallInternalPrefixes = *(*[]string)(unsafe.Pointer(&in.FirewallInternalPrefixes))
	if err := Convert_v1alpha1_ETCD_To_config_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ClusterAudit_To_config_ClusterAudit(&in.ClusterAudit, &out.ClusterAudit, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AuditToSplunk_To_config_AuditToSplunk(&in.AuditToSplunk, &out.AuditToSplunk, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*apisconfig.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	if err := Convert_v1alpha1_StorageConfiguration_To_config_StorageConfiguration(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	out.ImagePullPolicy = in.ImagePullPolicy
	out.ImagePullSecret = (*config.ImagePullSecret)(unsafe.Pointer(in.ImagePullSecret))
	out.EgressDestinations = *(*[]config.EgressDest)(unsafe.Pointer(&in.EgressDestinations))
	return nil
}

// Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in, out, s)
}

func autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	out.ClientConnection = (*configv1alpha1.ClientConnectionConfiguration)(unsafe.Pointer(in.ClientConnection))
	out.MachineImages = *(*[]MachineImage)(unsafe.Pointer(&in.MachineImages))
	out.FirewallInternalPrefixes = *(*[]string)(unsafe.Pointer(&in.FirewallInternalPrefixes))
	if err := Convert_config_ETCD_To_v1alpha1_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	if err := Convert_config_ClusterAudit_To_v1alpha1_ClusterAudit(&in.ClusterAudit, &out.ClusterAudit, s); err != nil {
		return err
	}
	if err := Convert_config_AuditToSplunk_To_v1alpha1_AuditToSplunk(&in.AuditToSplunk, &out.AuditToSplunk, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*apisconfigv1alpha1.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	if err := Convert_config_StorageConfiguration_To_v1alpha1_StorageConfiguration(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	out.ImagePullPolicy = in.ImagePullPolicy
	out.ImagePullSecret = (*ImagePullSecret)(unsafe.Pointer(in.ImagePullSecret))
	out.EgressDestinations = *(*[]EgressDest)(unsafe.Pointer(&in.EgressDestinations))
	return nil
}

// Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration is an autogenerated conversion function.
func Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_DurosConfiguration_To_config_DurosConfiguration(in *DurosConfiguration, out *config.DurosConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.PartitionConfig = *(*map[string]config.DurosPartitionConfiguration)(unsafe.Pointer(&in.PartitionConfig))
	return nil
}

// Convert_v1alpha1_DurosConfiguration_To_config_DurosConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_DurosConfiguration_To_config_DurosConfiguration(in *DurosConfiguration, out *config.DurosConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_DurosConfiguration_To_config_DurosConfiguration(in, out, s)
}

func autoConvert_config_DurosConfiguration_To_v1alpha1_DurosConfiguration(in *config.DurosConfiguration, out *DurosConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.PartitionConfig = *(*map[string]DurosPartitionConfiguration)(unsafe.Pointer(&in.PartitionConfig))
	return nil
}

// Convert_config_DurosConfiguration_To_v1alpha1_DurosConfiguration is an autogenerated conversion function.
func Convert_config_DurosConfiguration_To_v1alpha1_DurosConfiguration(in *config.DurosConfiguration, out *DurosConfiguration, s conversion.Scope) error {
	return autoConvert_config_DurosConfiguration_To_v1alpha1_DurosConfiguration(in, out, s)
}

func autoConvert_v1alpha1_DurosPartitionConfiguration_To_config_DurosPartitionConfiguration(in *DurosPartitionConfiguration, out *config.DurosPartitionConfiguration, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.AdminKey = in.AdminKey
	out.AdminToken = in.AdminToken
	out.StorageClasses = *(*[]config.DurosSeedStorageClass)(unsafe.Pointer(&in.StorageClasses))
	out.APIEndpoint = in.APIEndpoint
	out.APICA = in.APICA
	out.APICert = in.APICert
	out.APIKey = in.APIKey
	return nil
}

// Convert_v1alpha1_DurosPartitionConfiguration_To_config_DurosPartitionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_DurosPartitionConfiguration_To_config_DurosPartitionConfiguration(in *DurosPartitionConfiguration, out *config.DurosPartitionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_DurosPartitionConfiguration_To_config_DurosPartitionConfiguration(in, out, s)
}

func autoConvert_config_DurosPartitionConfiguration_To_v1alpha1_DurosPartitionConfiguration(in *config.DurosPartitionConfiguration, out *DurosPartitionConfiguration, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.AdminKey = in.AdminKey
	out.AdminToken = in.AdminToken
	out.StorageClasses = *(*[]DurosSeedStorageClass)(unsafe.Pointer(&in.StorageClasses))
	out.APIEndpoint = in.APIEndpoint
	out.APICA = in.APICA
	out.APICert = in.APICert
	out.APIKey = in.APIKey
	return nil
}

// Convert_config_DurosPartitionConfiguration_To_v1alpha1_DurosPartitionConfiguration is an autogenerated conversion function.
func Convert_config_DurosPartitionConfiguration_To_v1alpha1_DurosPartitionConfiguration(in *config.DurosPartitionConfiguration, out *DurosPartitionConfiguration, s conversion.Scope) error {
	return autoConvert_config_DurosPartitionConfiguration_To_v1alpha1_DurosPartitionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_DurosSeedStorageClass_To_config_DurosSeedStorageClass(in *DurosSeedStorageClass, out *config.DurosSeedStorageClass, s conversion.Scope) error {
	out.Name = in.Name
	out.ReplicaCount = in.ReplicaCount
	out.Compression = in.Compression
	out.Encryption = in.Encryption
	return nil
}

// Convert_v1alpha1_DurosSeedStorageClass_To_config_DurosSeedStorageClass is an autogenerated conversion function.
func Convert_v1alpha1_DurosSeedStorageClass_To_config_DurosSeedStorageClass(in *DurosSeedStorageClass, out *config.DurosSeedStorageClass, s conversion.Scope) error {
	return autoConvert_v1alpha1_DurosSeedStorageClass_To_config_DurosSeedStorageClass(in, out, s)
}

func autoConvert_config_DurosSeedStorageClass_To_v1alpha1_DurosSeedStorageClass(in *config.DurosSeedStorageClass, out *DurosSeedStorageClass, s conversion.Scope) error {
	out.Name = in.Name
	out.ReplicaCount = in.ReplicaCount
	out.Compression = in.Compression
	out.Encryption = in.Encryption
	return nil
}

// Convert_config_DurosSeedStorageClass_To_v1alpha1_DurosSeedStorageClass is an autogenerated conversion function.
func Convert_config_DurosSeedStorageClass_To_v1alpha1_DurosSeedStorageClass(in *config.DurosSeedStorageClass, out *DurosSeedStorageClass, s conversion.Scope) error {
	return autoConvert_config_DurosSeedStorageClass_To_v1alpha1_DurosSeedStorageClass(in, out, s)
}

func autoConvert_v1alpha1_ETCD_To_config_ETCD(in *ETCD, out *config.ETCD, s conversion.Scope) error {
	if err := Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(&in.Backup, &out.Backup, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ETCD_To_config_ETCD is an autogenerated conversion function.
func Convert_v1alpha1_ETCD_To_config_ETCD(in *ETCD, out *config.ETCD, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCD_To_config_ETCD(in, out, s)
}

func autoConvert_config_ETCD_To_v1alpha1_ETCD(in *config.ETCD, out *ETCD, s conversion.Scope) error {
	if err := Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	if err := Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(&in.Backup, &out.Backup, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ETCD_To_v1alpha1_ETCD is an autogenerated conversion function.
func Convert_config_ETCD_To_v1alpha1_ETCD(in *config.ETCD, out *ETCD, s conversion.Scope) error {
	return autoConvert_config_ETCD_To_v1alpha1_ETCD(in, out, s)
}

func autoConvert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in *ETCDBackup, out *config.ETCDBackup, s conversion.Scope) error {
	out.Schedule = (*string)(unsafe.Pointer(in.Schedule))
	out.DeltaSnapshotPeriod = (*string)(unsafe.Pointer(in.DeltaSnapshotPeriod))
	return nil
}

// Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup is an autogenerated conversion function.
func Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in *ETCDBackup, out *config.ETCDBackup, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in, out, s)
}

func autoConvert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in *config.ETCDBackup, out *ETCDBackup, s conversion.Scope) error {
	out.Schedule = (*string)(unsafe.Pointer(in.Schedule))
	out.DeltaSnapshotPeriod = (*string)(unsafe.Pointer(in.DeltaSnapshotPeriod))
	return nil
}

// Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup is an autogenerated conversion function.
func Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in *config.ETCDBackup, out *ETCDBackup, s conversion.Scope) error {
	return autoConvert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in, out, s)
}

func autoConvert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in *ETCDStorage, out *config.ETCDStorage, s conversion.Scope) error {
	out.ClassName = (*string)(unsafe.Pointer(in.ClassName))
	out.Capacity = (*resource.Quantity)(unsafe.Pointer(in.Capacity))
	return nil
}

// Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage is an autogenerated conversion function.
func Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in *ETCDStorage, out *config.ETCDStorage, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in, out, s)
}

func autoConvert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in *config.ETCDStorage, out *ETCDStorage, s conversion.Scope) error {
	out.ClassName = (*string)(unsafe.Pointer(in.ClassName))
	out.Capacity = (*resource.Quantity)(unsafe.Pointer(in.Capacity))
	return nil
}

// Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage is an autogenerated conversion function.
func Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in *config.ETCDStorage, out *ETCDStorage, s conversion.Scope) error {
	return autoConvert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in, out, s)
}

func autoConvert_v1alpha1_EgressDest_To_config_EgressDest(in *EgressDest, out *config.EgressDest, s conversion.Scope) error {
	out.Description = in.Description
	out.MatchPattern = in.MatchPattern
	out.MatchName = in.MatchName
	out.Protocol = in.Protocol
	out.Port = in.Port
	return nil
}

// Convert_v1alpha1_EgressDest_To_config_EgressDest is an autogenerated conversion function.
func Convert_v1alpha1_EgressDest_To_config_EgressDest(in *EgressDest, out *config.EgressDest, s conversion.Scope) error {
	return autoConvert_v1alpha1_EgressDest_To_config_EgressDest(in, out, s)
}

func autoConvert_config_EgressDest_To_v1alpha1_EgressDest(in *config.EgressDest, out *EgressDest, s conversion.Scope) error {
	out.Description = in.Description
	out.MatchPattern = in.MatchPattern
	out.MatchName = in.MatchName
	out.Protocol = in.Protocol
	out.Port = in.Port
	return nil
}

// Convert_config_EgressDest_To_v1alpha1_EgressDest is an autogenerated conversion function.
func Convert_config_EgressDest_To_v1alpha1_EgressDest(in *config.EgressDest, out *EgressDest, s conversion.Scope) error {
	return autoConvert_config_EgressDest_To_v1alpha1_EgressDest(in, out, s)
}

func autoConvert_v1alpha1_ImagePullSecret_To_config_ImagePullSecret(in *ImagePullSecret, out *config.ImagePullSecret, s conversion.Scope) error {
	out.DockerConfigJSON = in.DockerConfigJSON
	return nil
}

// Convert_v1alpha1_ImagePullSecret_To_config_ImagePullSecret is an autogenerated conversion function.
func Convert_v1alpha1_ImagePullSecret_To_config_ImagePullSecret(in *ImagePullSecret, out *config.ImagePullSecret, s conversion.Scope) error {
	return autoConvert_v1alpha1_ImagePullSecret_To_config_ImagePullSecret(in, out, s)
}

func autoConvert_config_ImagePullSecret_To_v1alpha1_ImagePullSecret(in *config.ImagePullSecret, out *ImagePullSecret, s conversion.Scope) error {
	out.DockerConfigJSON = in.DockerConfigJSON
	return nil
}

// Convert_config_ImagePullSecret_To_v1alpha1_ImagePullSecret is an autogenerated conversion function.
func Convert_config_ImagePullSecret_To_v1alpha1_ImagePullSecret(in *config.ImagePullSecret, out *ImagePullSecret, s conversion.Scope) error {
	return autoConvert_config_ImagePullSecret_To_v1alpha1_ImagePullSecret(in, out, s)
}

func autoConvert_v1alpha1_MachineImage_To_config_MachineImage(in *MachineImage, out *config.MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Image = in.Image
	return nil
}

// Convert_v1alpha1_MachineImage_To_config_MachineImage is an autogenerated conversion function.
func Convert_v1alpha1_MachineImage_To_config_MachineImage(in *MachineImage, out *config.MachineImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImage_To_config_MachineImage(in, out, s)
}

func autoConvert_config_MachineImage_To_v1alpha1_MachineImage(in *config.MachineImage, out *MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Image = in.Image
	return nil
}

// Convert_config_MachineImage_To_v1alpha1_MachineImage is an autogenerated conversion function.
func Convert_config_MachineImage_To_v1alpha1_MachineImage(in *config.MachineImage, out *MachineImage, s conversion.Scope) error {
	return autoConvert_config_MachineImage_To_v1alpha1_MachineImage(in, out, s)
}

func autoConvert_v1alpha1_StorageConfiguration_To_config_StorageConfiguration(in *StorageConfiguration, out *config.StorageConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_DurosConfiguration_To_config_DurosConfiguration(&in.Duros, &out.Duros, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_StorageConfiguration_To_config_StorageConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_StorageConfiguration_To_config_StorageConfiguration(in *StorageConfiguration, out *config.StorageConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_StorageConfiguration_To_config_StorageConfiguration(in, out, s)
}

func autoConvert_config_StorageConfiguration_To_v1alpha1_StorageConfiguration(in *config.StorageConfiguration, out *StorageConfiguration, s conversion.Scope) error {
	if err := Convert_config_DurosConfiguration_To_v1alpha1_DurosConfiguration(&in.Duros, &out.Duros, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_StorageConfiguration_To_v1alpha1_StorageConfiguration is an autogenerated conversion function.
func Convert_config_StorageConfiguration_To_v1alpha1_StorageConfiguration(in *config.StorageConfiguration, out *StorageConfiguration, s conversion.Scope) error {
	return autoConvert_config_StorageConfiguration_To_v1alpha1_StorageConfiguration(in, out, s)
}
