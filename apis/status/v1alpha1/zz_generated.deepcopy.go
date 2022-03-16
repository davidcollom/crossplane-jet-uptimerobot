//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Page) DeepCopyInto(out *Page) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Page.
func (in *Page) DeepCopy() *Page {
	if in == nil {
		return nil
	}
	out := new(Page)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Page) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PageList) DeepCopyInto(out *PageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Page, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PageList.
func (in *PageList) DeepCopy() *PageList {
	if in == nil {
		return nil
	}
	out := new(PageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PageObservation) DeepCopyInto(out *PageObservation) {
	*out = *in
	if in.CustomURL != nil {
		in, out := &in.CustomURL, &out.CustomURL
		*out = new(string)
		**out = **in
	}
	if in.DNSAddress != nil {
		in, out := &in.DNSAddress, &out.DNSAddress
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.StandardURL != nil {
		in, out := &in.StandardURL, &out.StandardURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PageObservation.
func (in *PageObservation) DeepCopy() *PageObservation {
	if in == nil {
		return nil
	}
	out := new(PageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PageParameters) DeepCopyInto(out *PageParameters) {
	*out = *in
	if in.CustomDomain != nil {
		in, out := &in.CustomDomain, &out.CustomDomain
		*out = new(string)
		**out = **in
	}
	if in.FriendlyName != nil {
		in, out := &in.FriendlyName, &out.FriendlyName
		*out = new(string)
		**out = **in
	}
	if in.Monitors != nil {
		in, out := &in.Monitors, &out.Monitors
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Sort != nil {
		in, out := &in.Sort, &out.Sort
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PageParameters.
func (in *PageParameters) DeepCopy() *PageParameters {
	if in == nil {
		return nil
	}
	out := new(PageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PageSpec) DeepCopyInto(out *PageSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PageSpec.
func (in *PageSpec) DeepCopy() *PageSpec {
	if in == nil {
		return nil
	}
	out := new(PageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PageStatus) DeepCopyInto(out *PageStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PageStatus.
func (in *PageStatus) DeepCopy() *PageStatus {
	if in == nil {
		return nil
	}
	out := new(PageStatus)
	in.DeepCopyInto(out)
	return out
}
