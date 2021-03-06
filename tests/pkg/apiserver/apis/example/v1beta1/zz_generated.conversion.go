// +build !ignore_autogenerated

// Copyright 2019. PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	example "github.com/pingcap/tidb-operator/tests/pkg/apiserver/apis/example"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ContainerSpec)(nil), (*example.ContainerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ContainerSpec_To_example_ContainerSpec(a.(*ContainerSpec), b.(*example.ContainerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ContainerSpec)(nil), (*ContainerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ContainerSpec_To_v1beta1_ContainerSpec(a.(*example.ContainerSpec), b.(*ContainerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Pod)(nil), (*example.Pod)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Pod_To_example_Pod(a.(*Pod), b.(*example.Pod), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.Pod)(nil), (*Pod)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_Pod_To_v1beta1_Pod(a.(*example.Pod), b.(*Pod), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodList)(nil), (*example.PodList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodList_To_example_PodList(a.(*PodList), b.(*example.PodList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.PodList)(nil), (*PodList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_PodList_To_v1beta1_PodList(a.(*example.PodList), b.(*PodList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodSpec)(nil), (*example.PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodSpec_To_example_PodSpec(a.(*PodSpec), b.(*example.PodSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.PodSpec)(nil), (*PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_PodSpec_To_v1beta1_PodSpec(a.(*example.PodSpec), b.(*PodSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodStatus)(nil), (*example.PodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodStatus_To_example_PodStatus(a.(*PodStatus), b.(*example.PodStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.PodStatus)(nil), (*PodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_PodStatus_To_v1beta1_PodStatus(a.(*example.PodStatus), b.(*PodStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*example.PodSpec)(nil), (*PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_PodSpec_To_v1beta1_PodSpec(a.(*example.PodSpec), b.(*PodSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*PodSpec)(nil), (*example.PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodSpec_To_example_PodSpec(a.(*PodSpec), b.(*example.PodSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_ContainerSpec_To_example_ContainerSpec(in *ContainerSpec, out *example.ContainerSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.ImagePullPolicy = in.ImagePullPolicy
	return nil
}

// Convert_v1beta1_ContainerSpec_To_example_ContainerSpec is an autogenerated conversion function.
func Convert_v1beta1_ContainerSpec_To_example_ContainerSpec(in *ContainerSpec, out *example.ContainerSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ContainerSpec_To_example_ContainerSpec(in, out, s)
}

func autoConvert_example_ContainerSpec_To_v1beta1_ContainerSpec(in *example.ContainerSpec, out *ContainerSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.ImagePullPolicy = in.ImagePullPolicy
	return nil
}

// Convert_example_ContainerSpec_To_v1beta1_ContainerSpec is an autogenerated conversion function.
func Convert_example_ContainerSpec_To_v1beta1_ContainerSpec(in *example.ContainerSpec, out *ContainerSpec, s conversion.Scope) error {
	return autoConvert_example_ContainerSpec_To_v1beta1_ContainerSpec(in, out, s)
}

func autoConvert_v1beta1_Pod_To_example_Pod(in *Pod, out *example.Pod, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_PodSpec_To_example_PodSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_PodStatus_To_example_PodStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Pod_To_example_Pod is an autogenerated conversion function.
func Convert_v1beta1_Pod_To_example_Pod(in *Pod, out *example.Pod, s conversion.Scope) error {
	return autoConvert_v1beta1_Pod_To_example_Pod(in, out, s)
}

func autoConvert_example_Pod_To_v1beta1_Pod(in *example.Pod, out *Pod, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_example_PodSpec_To_v1beta1_PodSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_example_PodStatus_To_v1beta1_PodStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_example_Pod_To_v1beta1_Pod is an autogenerated conversion function.
func Convert_example_Pod_To_v1beta1_Pod(in *example.Pod, out *Pod, s conversion.Scope) error {
	return autoConvert_example_Pod_To_v1beta1_Pod(in, out, s)
}

func autoConvert_v1beta1_PodList_To_example_PodList(in *PodList, out *example.PodList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]example.Pod, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_Pod_To_example_Pod(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_PodList_To_example_PodList is an autogenerated conversion function.
func Convert_v1beta1_PodList_To_example_PodList(in *PodList, out *example.PodList, s conversion.Scope) error {
	return autoConvert_v1beta1_PodList_To_example_PodList(in, out, s)
}

func autoConvert_example_PodList_To_v1beta1_PodList(in *example.PodList, out *PodList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pod, len(*in))
		for i := range *in {
			if err := Convert_example_Pod_To_v1beta1_Pod(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_example_PodList_To_v1beta1_PodList is an autogenerated conversion function.
func Convert_example_PodList_To_v1beta1_PodList(in *example.PodList, out *PodList, s conversion.Scope) error {
	return autoConvert_example_PodList_To_v1beta1_PodList(in, out, s)
}

func autoConvert_v1beta1_PodSpec_To_example_PodSpec(in *PodSpec, out *example.PodSpec, s conversion.Scope) error {
	out.Containers = *(*[]example.ContainerSpec)(unsafe.Pointer(&in.Containers))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Tolerations = *(*[]string)(unsafe.Pointer(&in.Tolerations))
	out.HostName = in.HostName
	return nil
}

func autoConvert_example_PodSpec_To_v1beta1_PodSpec(in *example.PodSpec, out *PodSpec, s conversion.Scope) error {
	out.Containers = *(*[]ContainerSpec)(unsafe.Pointer(&in.Containers))
	// WARNING: in.Container requires manual conversion: does not exist in peer-type
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Tolerations = *(*[]string)(unsafe.Pointer(&in.Tolerations))
	out.HostName = in.HostName
	// WARNING: in.HasTolerations requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1beta1_PodStatus_To_example_PodStatus(in *PodStatus, out *example.PodStatus, s conversion.Scope) error {
	out.Phase = in.Phase
	return nil
}

// Convert_v1beta1_PodStatus_To_example_PodStatus is an autogenerated conversion function.
func Convert_v1beta1_PodStatus_To_example_PodStatus(in *PodStatus, out *example.PodStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_PodStatus_To_example_PodStatus(in, out, s)
}

func autoConvert_example_PodStatus_To_v1beta1_PodStatus(in *example.PodStatus, out *PodStatus, s conversion.Scope) error {
	out.Phase = in.Phase
	return nil
}

// Convert_example_PodStatus_To_v1beta1_PodStatus is an autogenerated conversion function.
func Convert_example_PodStatus_To_v1beta1_PodStatus(in *example.PodStatus, out *PodStatus, s conversion.Scope) error {
	return autoConvert_example_PodStatus_To_v1beta1_PodStatus(in, out, s)
}
