//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1alpha3

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	output "k8s.io/kubernetes/cmd/kubeadm/app/apis/output"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*BootstrapToken)(nil), (*output.BootstrapToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_BootstrapToken_To_output_BootstrapToken(a.(*BootstrapToken), b.(*output.BootstrapToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*output.BootstrapToken)(nil), (*BootstrapToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_output_BootstrapToken_To_v1alpha3_BootstrapToken(a.(*output.BootstrapToken), b.(*BootstrapToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Certificate)(nil), (*output.Certificate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Certificate_To_output_Certificate(a.(*Certificate), b.(*output.Certificate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*output.Certificate)(nil), (*Certificate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_output_Certificate_To_v1alpha3_Certificate(a.(*output.Certificate), b.(*Certificate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CertificateExpirationInfo)(nil), (*output.CertificateExpirationInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_CertificateExpirationInfo_To_output_CertificateExpirationInfo(a.(*CertificateExpirationInfo), b.(*output.CertificateExpirationInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*output.CertificateExpirationInfo)(nil), (*CertificateExpirationInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_output_CertificateExpirationInfo_To_v1alpha3_CertificateExpirationInfo(a.(*output.CertificateExpirationInfo), b.(*CertificateExpirationInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ComponentConfigVersionState)(nil), (*output.ComponentConfigVersionState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ComponentConfigVersionState_To_output_ComponentConfigVersionState(a.(*ComponentConfigVersionState), b.(*output.ComponentConfigVersionState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*output.ComponentConfigVersionState)(nil), (*ComponentConfigVersionState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_output_ComponentConfigVersionState_To_v1alpha3_ComponentConfigVersionState(a.(*output.ComponentConfigVersionState), b.(*ComponentConfigVersionState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ComponentUpgradePlan)(nil), (*output.ComponentUpgradePlan)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ComponentUpgradePlan_To_output_ComponentUpgradePlan(a.(*ComponentUpgradePlan), b.(*output.ComponentUpgradePlan), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*output.ComponentUpgradePlan)(nil), (*ComponentUpgradePlan)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_output_ComponentUpgradePlan_To_v1alpha3_ComponentUpgradePlan(a.(*output.ComponentUpgradePlan), b.(*ComponentUpgradePlan), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Images)(nil), (*output.Images)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Images_To_output_Images(a.(*Images), b.(*output.Images), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*output.Images)(nil), (*Images)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_output_Images_To_v1alpha3_Images(a.(*output.Images), b.(*Images), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UpgradePlan)(nil), (*output.UpgradePlan)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_UpgradePlan_To_output_UpgradePlan(a.(*UpgradePlan), b.(*output.UpgradePlan), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*output.UpgradePlan)(nil), (*UpgradePlan)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_output_UpgradePlan_To_v1alpha3_UpgradePlan(a.(*output.UpgradePlan), b.(*UpgradePlan), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_BootstrapToken_To_output_BootstrapToken(in *BootstrapToken, out *output.BootstrapToken, s conversion.Scope) error {
	out.BootstrapToken = in.BootstrapToken
	return nil
}

// Convert_v1alpha3_BootstrapToken_To_output_BootstrapToken is an autogenerated conversion function.
func Convert_v1alpha3_BootstrapToken_To_output_BootstrapToken(in *BootstrapToken, out *output.BootstrapToken, s conversion.Scope) error {
	return autoConvert_v1alpha3_BootstrapToken_To_output_BootstrapToken(in, out, s)
}

func autoConvert_output_BootstrapToken_To_v1alpha3_BootstrapToken(in *output.BootstrapToken, out *BootstrapToken, s conversion.Scope) error {
	out.BootstrapToken = in.BootstrapToken
	return nil
}

// Convert_output_BootstrapToken_To_v1alpha3_BootstrapToken is an autogenerated conversion function.
func Convert_output_BootstrapToken_To_v1alpha3_BootstrapToken(in *output.BootstrapToken, out *BootstrapToken, s conversion.Scope) error {
	return autoConvert_output_BootstrapToken_To_v1alpha3_BootstrapToken(in, out, s)
}

func autoConvert_v1alpha3_Certificate_To_output_Certificate(in *Certificate, out *output.Certificate, s conversion.Scope) error {
	out.Name = in.Name
	out.ExpirationDate = in.ExpirationDate
	out.ResidualTime = in.ResidualTime
	out.ExternallyManaged = in.ExternallyManaged
	out.CAName = in.CAName
	out.Missing = in.Missing
	return nil
}

// Convert_v1alpha3_Certificate_To_output_Certificate is an autogenerated conversion function.
func Convert_v1alpha3_Certificate_To_output_Certificate(in *Certificate, out *output.Certificate, s conversion.Scope) error {
	return autoConvert_v1alpha3_Certificate_To_output_Certificate(in, out, s)
}

func autoConvert_output_Certificate_To_v1alpha3_Certificate(in *output.Certificate, out *Certificate, s conversion.Scope) error {
	out.Name = in.Name
	out.ExpirationDate = in.ExpirationDate
	out.ResidualTime = in.ResidualTime
	out.ExternallyManaged = in.ExternallyManaged
	out.CAName = in.CAName
	out.Missing = in.Missing
	return nil
}

// Convert_output_Certificate_To_v1alpha3_Certificate is an autogenerated conversion function.
func Convert_output_Certificate_To_v1alpha3_Certificate(in *output.Certificate, out *Certificate, s conversion.Scope) error {
	return autoConvert_output_Certificate_To_v1alpha3_Certificate(in, out, s)
}

func autoConvert_v1alpha3_CertificateExpirationInfo_To_output_CertificateExpirationInfo(in *CertificateExpirationInfo, out *output.CertificateExpirationInfo, s conversion.Scope) error {
	out.Certificates = *(*[]output.Certificate)(unsafe.Pointer(&in.Certificates))
	out.CertificateAuthorities = *(*[]output.Certificate)(unsafe.Pointer(&in.CertificateAuthorities))
	return nil
}

// Convert_v1alpha3_CertificateExpirationInfo_To_output_CertificateExpirationInfo is an autogenerated conversion function.
func Convert_v1alpha3_CertificateExpirationInfo_To_output_CertificateExpirationInfo(in *CertificateExpirationInfo, out *output.CertificateExpirationInfo, s conversion.Scope) error {
	return autoConvert_v1alpha3_CertificateExpirationInfo_To_output_CertificateExpirationInfo(in, out, s)
}

func autoConvert_output_CertificateExpirationInfo_To_v1alpha3_CertificateExpirationInfo(in *output.CertificateExpirationInfo, out *CertificateExpirationInfo, s conversion.Scope) error {
	out.Certificates = *(*[]Certificate)(unsafe.Pointer(&in.Certificates))
	out.CertificateAuthorities = *(*[]Certificate)(unsafe.Pointer(&in.CertificateAuthorities))
	return nil
}

// Convert_output_CertificateExpirationInfo_To_v1alpha3_CertificateExpirationInfo is an autogenerated conversion function.
func Convert_output_CertificateExpirationInfo_To_v1alpha3_CertificateExpirationInfo(in *output.CertificateExpirationInfo, out *CertificateExpirationInfo, s conversion.Scope) error {
	return autoConvert_output_CertificateExpirationInfo_To_v1alpha3_CertificateExpirationInfo(in, out, s)
}

func autoConvert_v1alpha3_ComponentConfigVersionState_To_output_ComponentConfigVersionState(in *ComponentConfigVersionState, out *output.ComponentConfigVersionState, s conversion.Scope) error {
	out.Group = in.Group
	out.CurrentVersion = in.CurrentVersion
	out.PreferredVersion = in.PreferredVersion
	out.ManualUpgradeRequired = in.ManualUpgradeRequired
	return nil
}

// Convert_v1alpha3_ComponentConfigVersionState_To_output_ComponentConfigVersionState is an autogenerated conversion function.
func Convert_v1alpha3_ComponentConfigVersionState_To_output_ComponentConfigVersionState(in *ComponentConfigVersionState, out *output.ComponentConfigVersionState, s conversion.Scope) error {
	return autoConvert_v1alpha3_ComponentConfigVersionState_To_output_ComponentConfigVersionState(in, out, s)
}

func autoConvert_output_ComponentConfigVersionState_To_v1alpha3_ComponentConfigVersionState(in *output.ComponentConfigVersionState, out *ComponentConfigVersionState, s conversion.Scope) error {
	out.Group = in.Group
	out.CurrentVersion = in.CurrentVersion
	out.PreferredVersion = in.PreferredVersion
	out.ManualUpgradeRequired = in.ManualUpgradeRequired
	return nil
}

// Convert_output_ComponentConfigVersionState_To_v1alpha3_ComponentConfigVersionState is an autogenerated conversion function.
func Convert_output_ComponentConfigVersionState_To_v1alpha3_ComponentConfigVersionState(in *output.ComponentConfigVersionState, out *ComponentConfigVersionState, s conversion.Scope) error {
	return autoConvert_output_ComponentConfigVersionState_To_v1alpha3_ComponentConfigVersionState(in, out, s)
}

func autoConvert_v1alpha3_ComponentUpgradePlan_To_output_ComponentUpgradePlan(in *ComponentUpgradePlan, out *output.ComponentUpgradePlan, s conversion.Scope) error {
	out.Name = in.Name
	out.CurrentVersion = in.CurrentVersion
	out.NewVersion = in.NewVersion
	return nil
}

// Convert_v1alpha3_ComponentUpgradePlan_To_output_ComponentUpgradePlan is an autogenerated conversion function.
func Convert_v1alpha3_ComponentUpgradePlan_To_output_ComponentUpgradePlan(in *ComponentUpgradePlan, out *output.ComponentUpgradePlan, s conversion.Scope) error {
	return autoConvert_v1alpha3_ComponentUpgradePlan_To_output_ComponentUpgradePlan(in, out, s)
}

func autoConvert_output_ComponentUpgradePlan_To_v1alpha3_ComponentUpgradePlan(in *output.ComponentUpgradePlan, out *ComponentUpgradePlan, s conversion.Scope) error {
	out.Name = in.Name
	out.CurrentVersion = in.CurrentVersion
	out.NewVersion = in.NewVersion
	return nil
}

// Convert_output_ComponentUpgradePlan_To_v1alpha3_ComponentUpgradePlan is an autogenerated conversion function.
func Convert_output_ComponentUpgradePlan_To_v1alpha3_ComponentUpgradePlan(in *output.ComponentUpgradePlan, out *ComponentUpgradePlan, s conversion.Scope) error {
	return autoConvert_output_ComponentUpgradePlan_To_v1alpha3_ComponentUpgradePlan(in, out, s)
}

func autoConvert_v1alpha3_Images_To_output_Images(in *Images, out *output.Images, s conversion.Scope) error {
	out.Images = *(*[]string)(unsafe.Pointer(&in.Images))
	return nil
}

// Convert_v1alpha3_Images_To_output_Images is an autogenerated conversion function.
func Convert_v1alpha3_Images_To_output_Images(in *Images, out *output.Images, s conversion.Scope) error {
	return autoConvert_v1alpha3_Images_To_output_Images(in, out, s)
}

func autoConvert_output_Images_To_v1alpha3_Images(in *output.Images, out *Images, s conversion.Scope) error {
	out.Images = *(*[]string)(unsafe.Pointer(&in.Images))
	return nil
}

// Convert_output_Images_To_v1alpha3_Images is an autogenerated conversion function.
func Convert_output_Images_To_v1alpha3_Images(in *output.Images, out *Images, s conversion.Scope) error {
	return autoConvert_output_Images_To_v1alpha3_Images(in, out, s)
}

func autoConvert_v1alpha3_UpgradePlan_To_output_UpgradePlan(in *UpgradePlan, out *output.UpgradePlan, s conversion.Scope) error {
	out.Components = *(*[]output.ComponentUpgradePlan)(unsafe.Pointer(&in.Components))
	out.ConfigVersions = *(*[]output.ComponentConfigVersionState)(unsafe.Pointer(&in.ConfigVersions))
	return nil
}

// Convert_v1alpha3_UpgradePlan_To_output_UpgradePlan is an autogenerated conversion function.
func Convert_v1alpha3_UpgradePlan_To_output_UpgradePlan(in *UpgradePlan, out *output.UpgradePlan, s conversion.Scope) error {
	return autoConvert_v1alpha3_UpgradePlan_To_output_UpgradePlan(in, out, s)
}

func autoConvert_output_UpgradePlan_To_v1alpha3_UpgradePlan(in *output.UpgradePlan, out *UpgradePlan, s conversion.Scope) error {
	out.Components = *(*[]ComponentUpgradePlan)(unsafe.Pointer(&in.Components))
	out.ConfigVersions = *(*[]ComponentConfigVersionState)(unsafe.Pointer(&in.ConfigVersions))
	return nil
}

// Convert_output_UpgradePlan_To_v1alpha3_UpgradePlan is an autogenerated conversion function.
func Convert_output_UpgradePlan_To_v1alpha3_UpgradePlan(in *output.UpgradePlan, out *UpgradePlan, s conversion.Scope) error {
	return autoConvert_output_UpgradePlan_To_v1alpha3_UpgradePlan(in, out, s)
}
