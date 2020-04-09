/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/custominformers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AlertingRules returns a AlertingRuleInformer.
	AlertingRules() AlertingRuleInformer
	// AlertingSubRules returns a AlertingSubRuleInformer.
	AlertingSubRules() AlertingSubRuleInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AlertingRules returns a AlertingRuleInformer.
func (v *version) AlertingRules() AlertingRuleInformer {
	return &alertingRuleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// AlertingSubRules returns a AlertingSubRuleInformer.
func (v *version) AlertingSubRules() AlertingSubRuleInformer {
	return &alertingSubRuleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
