//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	configv1alpha1 "github.com/gardener/gardener/extensions/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACME) DeepCopyInto(out *ACME) {
	*out = *in
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	if in.PropagationTimeout != nil {
		in, out := &in.PropagationTimeout, &out.PropagationTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.PrecheckNameservers != nil {
		in, out := &in.PrecheckNameservers, &out.PrecheckNameservers
		*out = new(string)
		**out = **in
	}
	if in.CACertificates != nil {
		in, out := &in.CACertificates, &out.CACertificates
		*out = new(string)
		**out = **in
	}
	if in.DeactivateAuthorizations != nil {
		in, out := &in.DeactivateAuthorizations, &out.DeactivateAuthorizations
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACME.
func (in *ACME) DeepCopy() *ACME {
	if in == nil {
		return nil
	}
	out := new(ACME)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CA) DeepCopyInto(out *CA) {
	*out = *in
	if in.CACertificates != nil {
		in, out := &in.CACertificates, &out.CACertificates
		*out = new(string)
		**out = **in
	}
	if in.CAPrivateKey != nil {
		in, out := &in.CAPrivateKey, &out.CAPrivateKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CA.
func (in *CA) DeepCopy() *CA {
	if in == nil {
		return nil
	}
	out := new(CA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.RestrictIssuer != nil {
		in, out := &in.RestrictIssuer, &out.RestrictIssuer
		*out = new(bool)
		**out = **in
	}
	if in.DefaultRequestsPerDayQuota != nil {
		in, out := &in.DefaultRequestsPerDayQuota, &out.DefaultRequestsPerDayQuota
		*out = new(int32)
		**out = **in
	}
	if in.ShootIssuers != nil {
		in, out := &in.ShootIssuers, &out.ShootIssuers
		*out = new(ShootIssuers)
		**out = **in
	}
	in.ACME.DeepCopyInto(&out.ACME)
	in.CA.DeepCopyInto(&out.CA)
	if in.HealthCheckConfig != nil {
		in, out := &in.HealthCheckConfig, &out.HealthCheckConfig
		*out = new(configv1alpha1.HealthCheckConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateKeyDefaults != nil {
		in, out := &in.PrivateKeyDefaults, &out.PrivateKeyDefaults
		*out = new(PrivateKeyDefaults)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateKeyDefaults) DeepCopyInto(out *PrivateKeyDefaults) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.SizeRSA != nil {
		in, out := &in.SizeRSA, &out.SizeRSA
		*out = new(int)
		**out = **in
	}
	if in.SizeECDSA != nil {
		in, out := &in.SizeECDSA, &out.SizeECDSA
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateKeyDefaults.
func (in *PrivateKeyDefaults) DeepCopy() *PrivateKeyDefaults {
	if in == nil {
		return nil
	}
	out := new(PrivateKeyDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootIssuers) DeepCopyInto(out *ShootIssuers) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootIssuers.
func (in *ShootIssuers) DeepCopy() *ShootIssuers {
	if in == nil {
		return nil
	}
	out := new(ShootIssuers)
	in.DeepCopyInto(out)
	return out
}
