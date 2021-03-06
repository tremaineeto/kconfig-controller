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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvConfig) DeepCopyInto(out *EnvConfig) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.RefName != nil {
		in, out := &in.RefName, &out.RefName
		*out = new(string)
		**out = **in
	}
	if in.RefKey != nil {
		in, out := &in.RefKey, &out.RefKey
		*out = new(string)
		**out = **in
	}
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.FieldRef != nil {
		in, out := &in.FieldRef, &out.FieldRef
		*out = new(v1.ObjectFieldSelector)
		**out = **in
	}
	if in.ResourceFieldRef != nil {
		in, out := &in.ResourceFieldRef, &out.ResourceFieldRef
		*out = new(v1.ResourceFieldSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvConfig.
func (in *EnvConfig) DeepCopy() *EnvConfig {
	if in == nil {
		return nil
	}
	out := new(EnvConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kconfig) DeepCopyInto(out *Kconfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kconfig.
func (in *Kconfig) DeepCopy() *Kconfig {
	if in == nil {
		return nil
	}
	out := new(Kconfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kconfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KconfigBinding) DeepCopyInto(out *KconfigBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KconfigBinding.
func (in *KconfigBinding) DeepCopy() *KconfigBinding {
	if in == nil {
		return nil
	}
	out := new(KconfigBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KconfigBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KconfigBindingList) DeepCopyInto(out *KconfigBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KconfigBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KconfigBindingList.
func (in *KconfigBindingList) DeepCopy() *KconfigBindingList {
	if in == nil {
		return nil
	}
	out := new(KconfigBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KconfigBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KconfigBindingSpec) DeepCopyInto(out *KconfigBindingSpec) {
	*out = *in
	if in.KconfigEnvsMap != nil {
		in, out := &in.KconfigEnvsMap, &out.KconfigEnvsMap
		*out = make(map[string]KconfigEnvs, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KconfigBindingSpec.
func (in *KconfigBindingSpec) DeepCopy() *KconfigBindingSpec {
	if in == nil {
		return nil
	}
	out := new(KconfigBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KconfigEnvs) DeepCopyInto(out *KconfigEnvs) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KconfigEnvs.
func (in *KconfigEnvs) DeepCopy() *KconfigEnvs {
	if in == nil {
		return nil
	}
	out := new(KconfigEnvs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KconfigList) DeepCopyInto(out *KconfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kconfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KconfigList.
func (in *KconfigList) DeepCopy() *KconfigList {
	if in == nil {
		return nil
	}
	out := new(KconfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KconfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KconfigSpec) DeepCopyInto(out *KconfigSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.EnvConfigs != nil {
		in, out := &in.EnvConfigs, &out.EnvConfigs
		*out = make([]EnvConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KconfigSpec.
func (in *KconfigSpec) DeepCopy() *KconfigSpec {
	if in == nil {
		return nil
	}
	out := new(KconfigSpec)
	in.DeepCopyInto(out)
	return out
}
