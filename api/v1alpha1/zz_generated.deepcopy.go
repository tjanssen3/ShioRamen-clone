// +build !ignore_autogenerated

/*
Copyright 2021 The RamenDR authors.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPlacementControl) DeepCopyInto(out *DRPlacementControl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPlacementControl.
func (in *DRPlacementControl) DeepCopy() *DRPlacementControl {
	if in == nil {
		return nil
	}
	out := new(DRPlacementControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRPlacementControl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPlacementControlList) DeepCopyInto(out *DRPlacementControlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DRPlacementControl, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPlacementControlList.
func (in *DRPlacementControlList) DeepCopy() *DRPlacementControlList {
	if in == nil {
		return nil
	}
	out := new(DRPlacementControlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRPlacementControlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPlacementControlSpec) DeepCopyInto(out *DRPlacementControlSpec) {
	*out = *in
	out.PlacementRef = in.PlacementRef
	out.DRPolicyRef = in.DRPolicyRef
	in.PVCSelector.DeepCopyInto(&out.PVCSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPlacementControlSpec.
func (in *DRPlacementControlSpec) DeepCopy() *DRPlacementControlSpec {
	if in == nil {
		return nil
	}
	out := new(DRPlacementControlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPlacementControlStatus) DeepCopyInto(out *DRPlacementControlStatus) {
	*out = *in
	out.PreferredDecision = in.PreferredDecision
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ResourceConditions.DeepCopyInto(&out.ResourceConditions)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPlacementControlStatus.
func (in *DRPlacementControlStatus) DeepCopy() *DRPlacementControlStatus {
	if in == nil {
		return nil
	}
	out := new(DRPlacementControlStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPolicy) DeepCopyInto(out *DRPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(DRPolicyStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPolicy.
func (in *DRPolicy) DeepCopy() *DRPolicy {
	if in == nil {
		return nil
	}
	out := new(DRPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPolicyList) DeepCopyInto(out *DRPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DRPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPolicyList.
func (in *DRPolicyList) DeepCopy() *DRPolicyList {
	if in == nil {
		return nil
	}
	out := new(DRPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DRPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPolicySpec) DeepCopyInto(out *DRPolicySpec) {
	*out = *in
	in.ReplicationClassSelector.DeepCopyInto(&out.ReplicationClassSelector)
	if in.DRClusterSet != nil {
		in, out := &in.DRClusterSet, &out.DRClusterSet
		*out = make([]ManagedCluster, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPolicySpec.
func (in *DRPolicySpec) DeepCopy() *DRPolicySpec {
	if in == nil {
		return nil
	}
	out := new(DRPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRPolicyStatus) DeepCopyInto(out *DRPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRPolicyStatus.
func (in *DRPolicyStatus) DeepCopy() *DRPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DRPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedCluster) DeepCopyInto(out *ManagedCluster) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedCluster.
func (in *ManagedCluster) DeepCopy() *ManagedCluster {
	if in == nil {
		return nil
	}
	out := new(ManagedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedPVC) DeepCopyInto(out *ProtectedPVC) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedPVC.
func (in *ProtectedPVC) DeepCopy() *ProtectedPVC {
	if in == nil {
		return nil
	}
	out := new(ProtectedPVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ProtectedPVCMap) DeepCopyInto(out *ProtectedPVCMap) {
	{
		in := &in
		*out = make(ProtectedPVCMap, len(*in))
		for key, val := range *in {
			var outVal *ProtectedPVC
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ProtectedPVC)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedPVCMap.
func (in ProtectedPVCMap) DeepCopy() ProtectedPVCMap {
	if in == nil {
		return nil
	}
	out := new(ProtectedPVCMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RamenConfig) DeepCopyInto(out *RamenConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	if in.S3StoreProfiles != nil {
		in, out := &in.S3StoreProfiles, &out.S3StoreProfiles
		*out = make([]S3StoreProfile, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RamenConfig.
func (in *RamenConfig) DeepCopy() *RamenConfig {
	if in == nil {
		return nil
	}
	out := new(RamenConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RamenConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3StoreProfile) DeepCopyInto(out *S3StoreProfile) {
	*out = *in
	out.S3SecretRef = in.S3SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3StoreProfile.
func (in *S3StoreProfile) DeepCopy() *S3StoreProfile {
	if in == nil {
		return nil
	}
	out := new(S3StoreProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRGConditions) DeepCopyInto(out *VRGConditions) {
	*out = *in
	out.ResourceMeta = in.ResourceMeta
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRGConditions.
func (in *VRGConditions) DeepCopy() *VRGConditions {
	if in == nil {
		return nil
	}
	out := new(VRGConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRGResourceMeta) DeepCopyInto(out *VRGResourceMeta) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRGResourceMeta.
func (in *VRGResourceMeta) DeepCopy() *VRGResourceMeta {
	if in == nil {
		return nil
	}
	out := new(VRGResourceMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplicationGroup) DeepCopyInto(out *VolumeReplicationGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplicationGroup.
func (in *VolumeReplicationGroup) DeepCopy() *VolumeReplicationGroup {
	if in == nil {
		return nil
	}
	out := new(VolumeReplicationGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeReplicationGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplicationGroupList) DeepCopyInto(out *VolumeReplicationGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeReplicationGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplicationGroupList.
func (in *VolumeReplicationGroupList) DeepCopy() *VolumeReplicationGroupList {
	if in == nil {
		return nil
	}
	out := new(VolumeReplicationGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeReplicationGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplicationGroupSpec) DeepCopyInto(out *VolumeReplicationGroupSpec) {
	*out = *in
	in.PVCSelector.DeepCopyInto(&out.PVCSelector)
	in.ReplicationClassSelector.DeepCopyInto(&out.ReplicationClassSelector)
	if in.S3ProfileList != nil {
		in, out := &in.S3ProfileList, &out.S3ProfileList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplicationGroupSpec.
func (in *VolumeReplicationGroupSpec) DeepCopy() *VolumeReplicationGroupSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeReplicationGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplicationGroupStatus) DeepCopyInto(out *VolumeReplicationGroupStatus) {
	*out = *in
	if in.ProtectedPVCs != nil {
		in, out := &in.ProtectedPVCs, &out.ProtectedPVCs
		*out = make(ProtectedPVCMap, len(*in))
		for key, val := range *in {
			var outVal *ProtectedPVC
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ProtectedPVC)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplicationGroupStatus.
func (in *VolumeReplicationGroupStatus) DeepCopy() *VolumeReplicationGroupStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeReplicationGroupStatus)
	in.DeepCopyInto(out)
	return out
}
