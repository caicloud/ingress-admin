// +build !ignore_autogenerated

/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunProvider) DeepCopyInto(out *AliyunProvider) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunProvider.
func (in *AliyunProvider) DeepCopy() *AliyunProvider {
	if in == nil {
		return nil
	}
	out := new(AliyunProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunProviderStatus) DeepCopyInto(out *AliyunProviderStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunProviderStatus.
func (in *AliyunProviderStatus) DeepCopy() *AliyunProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AliyunProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureIPAddressProperties) DeepCopyInto(out *AzureIPAddressProperties) {
	*out = *in
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = new(AzurePrivateIPAddressProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(AzurePublicIPAddressProperties)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureIPAddressProperties.
func (in *AzureIPAddressProperties) DeepCopy() *AzureIPAddressProperties {
	if in == nil {
		return nil
	}
	out := new(AzureIPAddressProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzurePrivateIPAddressProperties) DeepCopyInto(out *AzurePrivateIPAddressProperties) {
	*out = *in
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzurePrivateIPAddressProperties.
func (in *AzurePrivateIPAddressProperties) DeepCopy() *AzurePrivateIPAddressProperties {
	if in == nil {
		return nil
	}
	out := new(AzurePrivateIPAddressProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureProvider) DeepCopyInto(out *AzureProvider) {
	*out = *in
	if in.ReserveAzure != nil {
		in, out := &in.ReserveAzure, &out.ReserveAzure
		*out = new(bool)
		**out = **in
	}
	in.IPAddressProperties.DeepCopyInto(&out.IPAddressProperties)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureProvider.
func (in *AzureProvider) DeepCopy() *AzureProvider {
	if in == nil {
		return nil
	}
	out := new(AzureProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureProviderStatus) DeepCopyInto(out *AzureProviderStatus) {
	*out = *in
	if in.PublicIPAddress != nil {
		in, out := &in.PublicIPAddress, &out.PublicIPAddress
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureProviderStatus.
func (in *AzureProviderStatus) DeepCopy() *AzureProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AzureProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzurePublicIPAddressProperties) DeepCopyInto(out *AzurePublicIPAddressProperties) {
	*out = *in
	if in.PublicIPAddressID != nil {
		in, out := &in.PublicIPAddressID, &out.PublicIPAddressID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzurePublicIPAddressProperties.
func (in *AzurePublicIPAddressProperties) DeepCopy() *AzurePublicIPAddressProperties {
	if in == nil {
		return nil
	}
	out := new(AzurePublicIPAddressProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpternalProviderStatus) DeepCopyInto(out *ExpternalProviderStatus) {
	*out = *in
	if in.VIPs != nil {
		in, out := &in.VIPs, &out.VIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpternalProviderStatus.
func (in *ExpternalProviderStatus) DeepCopy() *ExpternalProviderStatus {
	if in == nil {
		return nil
	}
	out := new(ExpternalProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalProvider) DeepCopyInto(out *ExternalProvider) {
	*out = *in
	if in.VIPs != nil {
		in, out := &in.VIPs, &out.VIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalProvider.
func (in *ExternalProvider) DeepCopy() *ExternalProvider {
	if in == nil {
		return nil
	}
	out := new(ExternalProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpvsdrProvider) DeepCopyInto(out *IpvsdrProvider) {
	*out = *in
	if in.VIPs != nil {
		in, out := &in.VIPs, &out.VIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpvsdrProvider.
func (in *IpvsdrProvider) DeepCopy() *IpvsdrProvider {
	if in == nil {
		return nil
	}
	out := new(IpvsdrProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpvsdrProviderStatus) DeepCopyInto(out *IpvsdrProviderStatus) {
	*out = *in
	in.PodStatuses.DeepCopyInto(&out.PodStatuses)
	if in.VIPs != nil {
		in, out := &in.VIPs, &out.VIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Vrid != nil {
		in, out := &in.Vrid, &out.Vrid
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpvsdrProviderStatus.
func (in *IpvsdrProviderStatus) DeepCopy() *IpvsdrProviderStatus {
	if in == nil {
		return nil
	}
	out := new(IpvsdrProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerList) DeepCopyInto(out *LoadBalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerList.
func (in *LoadBalancerList) DeepCopy() *LoadBalancerList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	in.Nodes.DeepCopyInto(&out.Nodes)
	in.Proxy.DeepCopyInto(&out.Proxy)
	in.Providers.DeepCopyInto(&out.Providers)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	if in.AccessIPs != nil {
		in, out := &in.AccessIPs, &out.AccessIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeIPs != nil {
		in, out := &in.NodeIPs, &out.NodeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ProxyStatus.DeepCopyInto(&out.ProxyStatus)
	in.ProvidersStatuses.DeepCopyInto(&out.ProvidersStatuses)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesSpec) DeepCopyInto(out *NodesSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Effect != nil {
		in, out := &in.Effect, &out.Effect
		*out = new(v1.TaintEffect)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesSpec.
func (in *NodesSpec) DeepCopy() *NodesSpec {
	if in == nil {
		return nil
	}
	out := new(NodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatuses) DeepCopyInto(out *PodStatuses) {
	*out = *in
	if in.Statuses != nil {
		in, out := &in.Statuses, &out.Statuses
		*out = make([]PodStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatuses.
func (in *PodStatuses) DeepCopy() *PodStatuses {
	if in == nil {
		return nil
	}
	out := new(PodStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRange) DeepCopyInto(out *PortRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRange.
func (in *PortRange) DeepCopy() *PortRange {
	if in == nil {
		return nil
	}
	out := new(PortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvidersSpec) DeepCopyInto(out *ProvidersSpec) {
	*out = *in
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Ipvsdr != nil {
		in, out := &in.Ipvsdr, &out.Ipvsdr
		*out = new(IpvsdrProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Aliyun != nil {
		in, out := &in.Aliyun, &out.Aliyun
		*out = new(AliyunProvider)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureProvider)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvidersSpec.
func (in *ProvidersSpec) DeepCopy() *ProvidersSpec {
	if in == nil {
		return nil
	}
	out := new(ProvidersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvidersStatuses) DeepCopyInto(out *ProvidersStatuses) {
	*out = *in
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExpternalProviderStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Ipvsdr != nil {
		in, out := &in.Ipvsdr, &out.Ipvsdr
		*out = new(IpvsdrProviderStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Aliyun != nil {
		in, out := &in.Aliyun, &out.Aliyun
		*out = new(AliyunProviderStatus)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureProviderStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvidersStatuses.
func (in *ProvidersStatuses) DeepCopy() *ProvidersStatuses {
	if in == nil {
		return nil
	}
	out := new(ProvidersStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySpec) DeepCopyInto(out *ProxySpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.PortRanges != nil {
		in, out := &in.PortRanges, &out.PortRanges
		*out = make([]PortRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySpec.
func (in *ProxySpec) DeepCopy() *ProxySpec {
	if in == nil {
		return nil
	}
	out := new(ProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyStatus) DeepCopyInto(out *ProxyStatus) {
	*out = *in
	in.PodStatuses.DeepCopyInto(&out.PodStatuses)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyStatus.
func (in *ProxyStatus) DeepCopy() *ProxyStatus {
	if in == nil {
		return nil
	}
	out := new(ProxyStatus)
	in.DeepCopyInto(out)
	return out
}
