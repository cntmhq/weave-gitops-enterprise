// +build !ignore_autogenerated

/*
Copyright 2021.
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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPIResourceTemplate) DeepCopyInto(out *CAPIResourceTemplate) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPIResourceTemplate.
func (in *CAPIResourceTemplate) DeepCopy() *CAPIResourceTemplate {
	if in == nil {
		return nil
	}
	out := new(CAPIResourceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPITemplate) DeepCopyInto(out *CAPITemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPITemplate.
func (in *CAPITemplate) DeepCopy() *CAPITemplate {
	if in == nil {
		return nil
	}
	out := new(CAPITemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPITemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPITemplateList) DeepCopyInto(out *CAPITemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CAPITemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPITemplateList.
func (in *CAPITemplateList) DeepCopy() *CAPITemplateList {
	if in == nil {
		return nil
	}
	out := new(CAPITemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAPITemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPITemplateSpec) DeepCopyInto(out *CAPITemplateSpec) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]TemplateParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceTemplates != nil {
		in, out := &in.ResourceTemplates, &out.ResourceTemplates
		*out = make([]CAPIResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPITemplateSpec.
func (in *CAPITemplateSpec) DeepCopy() *CAPITemplateSpec {
	if in == nil {
		return nil
	}
	out := new(CAPITemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAPITemplateStatus) DeepCopyInto(out *CAPITemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAPITemplateStatus.
func (in *CAPITemplateStatus) DeepCopy() *CAPITemplateStatus {
	if in == nil {
		return nil
	}
	out := new(CAPITemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateParam) DeepCopyInto(out *TemplateParam) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateParam.
func (in *TemplateParam) DeepCopy() *TemplateParam {
	if in == nil {
		return nil
	}
	out := new(TemplateParam)
	in.DeepCopyInto(out)
	return out
}