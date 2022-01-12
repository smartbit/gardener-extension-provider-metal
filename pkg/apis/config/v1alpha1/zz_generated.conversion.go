//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	healthcheckconfig "github.com/gardener/gardener/extensions/pkg/controller/healthcheck/config"
	configv1alpha1 "github.com/gardener/gardener/extensions/pkg/controller/healthcheck/config/v1alpha1"
	config "github.com/metal-stack/gardener-extension-provider-metal/pkg/apis/config"
	resource "k8s.io/apimachinery/pkg/api/resource"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AccountingExporterClientConfiguration)(nil), (*config.AccountingExporterClientConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(a.(*AccountingExporterClientConfiguration), b.(*config.AccountingExporterClientConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AccountingExporterClientConfiguration)(nil), (*AccountingExporterClientConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(a.(*config.AccountingExporterClientConfiguration), b.(*AccountingExporterClientConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountingExporterConfiguration)(nil), (*config.AccountingExporterConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(a.(*AccountingExporterConfiguration), b.(*config.AccountingExporterConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AccountingExporterConfiguration)(nil), (*AccountingExporterConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(a.(*config.AccountingExporterConfiguration), b.(*AccountingExporterConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountingExporterNetworkTrafficConfiguration)(nil), (*config.AccountingExporterNetworkTrafficConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(a.(*AccountingExporterNetworkTrafficConfiguration), b.(*config.AccountingExporterNetworkTrafficConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AccountingExporterNetworkTrafficConfiguration)(nil), (*AccountingExporterNetworkTrafficConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(a.(*config.AccountingExporterNetworkTrafficConfiguration), b.(*AccountingExporterNetworkTrafficConfiguration), scope)
	}); err != nil {
		return err
	}
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
	if err := s.AddGeneratedConversionFunc((*Auth)(nil), (*config.Auth)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Auth_To_config_Auth(a.(*Auth), b.(*config.Auth), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Auth)(nil), (*Auth)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Auth_To_v1alpha1_Auth(a.(*config.Auth), b.(*Auth), scope)
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

func autoConvert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(in *AccountingExporterClientConfiguration, out *config.AccountingExporterClientConfiguration, s conversion.Scope) error {
	out.Hostname = in.Hostname
	out.Port = in.Port
	out.CA = in.CA
	out.Cert = in.Cert
	out.CertKey = in.CertKey
	return nil
}

// Convert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(in *AccountingExporterClientConfiguration, out *config.AccountingExporterClientConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(in, out, s)
}

func autoConvert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(in *config.AccountingExporterClientConfiguration, out *AccountingExporterClientConfiguration, s conversion.Scope) error {
	out.Hostname = in.Hostname
	out.Port = in.Port
	out.CA = in.CA
	out.Cert = in.Cert
	out.CertKey = in.CertKey
	return nil
}

// Convert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration is an autogenerated conversion function.
func Convert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(in *config.AccountingExporterClientConfiguration, out *AccountingExporterClientConfiguration, s conversion.Scope) error {
	return autoConvert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(in, out, s)
}

func autoConvert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(in *AccountingExporterConfiguration, out *config.AccountingExporterConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	if err := Convert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(&in.NetworkTraffic, &out.NetworkTraffic, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(&in.Client, &out.Client, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(in *AccountingExporterConfiguration, out *config.AccountingExporterConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(in, out, s)
}

func autoConvert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(in *config.AccountingExporterConfiguration, out *AccountingExporterConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	if err := Convert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(&in.NetworkTraffic, &out.NetworkTraffic, s); err != nil {
		return err
	}
	if err := Convert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(&in.Client, &out.Client, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration is an autogenerated conversion function.
func Convert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(in *config.AccountingExporterConfiguration, out *AccountingExporterConfiguration, s conversion.Scope) error {
	return autoConvert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(in, out, s)
}

func autoConvert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(in *AccountingExporterNetworkTrafficConfiguration, out *config.AccountingExporterNetworkTrafficConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.InternalNetworks = *(*[]string)(unsafe.Pointer(&in.InternalNetworks))
	return nil
}

// Convert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(in *AccountingExporterNetworkTrafficConfiguration, out *config.AccountingExporterNetworkTrafficConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(in, out, s)
}

func autoConvert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(in *config.AccountingExporterNetworkTrafficConfiguration, out *AccountingExporterNetworkTrafficConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.InternalNetworks = *(*[]string)(unsafe.Pointer(&in.InternalNetworks))
	return nil
}

// Convert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration is an autogenerated conversion function.
func Convert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(in *config.AccountingExporterNetworkTrafficConfiguration, out *AccountingExporterNetworkTrafficConfiguration, s conversion.Scope) error {
	return autoConvert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(in, out, s)
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

func autoConvert_v1alpha1_Auth_To_config_Auth(in *Auth, out *config.Auth, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.ProviderTenant = in.ProviderTenant
	return nil
}

// Convert_v1alpha1_Auth_To_config_Auth is an autogenerated conversion function.
func Convert_v1alpha1_Auth_To_config_Auth(in *Auth, out *config.Auth, s conversion.Scope) error {
	return autoConvert_v1alpha1_Auth_To_config_Auth(in, out, s)
}

func autoConvert_config_Auth_To_v1alpha1_Auth(in *config.Auth, out *Auth, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.ProviderTenant = in.ProviderTenant
	return nil
}

// Convert_config_Auth_To_v1alpha1_Auth is an autogenerated conversion function.
func Convert_config_Auth_To_v1alpha1_Auth(in *config.Auth, out *Auth, s conversion.Scope) error {
	return autoConvert_config_Auth_To_v1alpha1_Auth(in, out, s)
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
	out.MachineImages = *(*[]config.MachineImage)(unsafe.Pointer(&in.MachineImages))
	if err := Convert_v1alpha1_ETCD_To_config_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ClusterAudit_To_config_ClusterAudit(&in.ClusterAudit, &out.ClusterAudit, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AuditToSplunk_To_config_AuditToSplunk(&in.AuditToSplunk, &out.AuditToSplunk, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_Auth_To_config_Auth(&in.Auth, &out.Auth, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(&in.AccountingExporter, &out.AccountingExporter, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*healthcheckconfig.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	if err := Convert_v1alpha1_StorageConfiguration_To_config_StorageConfiguration(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	out.ImagePullSecret = (*config.ImagePullSecret)(unsafe.Pointer(in.ImagePullSecret))
	return nil
}

// Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in, out, s)
}

func autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	out.MachineImages = *(*[]MachineImage)(unsafe.Pointer(&in.MachineImages))
	if err := Convert_config_ETCD_To_v1alpha1_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	if err := Convert_config_ClusterAudit_To_v1alpha1_ClusterAudit(&in.ClusterAudit, &out.ClusterAudit, s); err != nil {
		return err
	}
	if err := Convert_config_AuditToSplunk_To_v1alpha1_AuditToSplunk(&in.AuditToSplunk, &out.AuditToSplunk, s); err != nil {
		return err
	}
	if err := Convert_config_Auth_To_v1alpha1_Auth(&in.Auth, &out.Auth, s); err != nil {
		return err
	}
	if err := Convert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(&in.AccountingExporter, &out.AccountingExporter, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*configv1alpha1.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	if err := Convert_config_StorageConfiguration_To_v1alpha1_StorageConfiguration(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	out.ImagePullSecret = (*ImagePullSecret)(unsafe.Pointer(in.ImagePullSecret))
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
	out.APIEndpoint = (*string)(unsafe.Pointer(in.APIEndpoint))
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
	out.APIEndpoint = (*string)(unsafe.Pointer(in.APIEndpoint))
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
