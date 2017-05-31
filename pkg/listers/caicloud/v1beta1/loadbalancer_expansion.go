package v1beta1

import (
	"fmt"

	"github.com/caicloud/loadbalancer-controller/api"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/labels"
)

// LoadBalancerListerExpansion allows custom methods to be added to
// LoadBalancerLister.
type LoadBalancerListerExpansion interface {
	GetLoadBalancersForControllee(obj interface{}) ([]*api.LoadBalancer, error)
}

// LoadBalancerNamespaceListerExpansion allows custom methods to be added to
// LoadBalancerNamespaeLister.
type LoadBalancerNamespaceListerExpansion interface{}

// GetLoadBalancersForControllee
func (s *loadBalancerList) GetLoadBalancersForControllee(obj interface{}) ([]*api.LoadBalancer, error) {
	meta, err := apimeta.Accessor(obj)
	if err != nil {
		return nil, fmt.Errorf("object has no meta: %v", err)
	}

	tpy, err := apimeta.TypeAccessor(obj)
	if err != nil {
		return nil, fmt.Errorf("object has no type: %v", err)
	}

	if len(meta.GetLabels()) == 0 {
		return nil, fmt.Errorf("no loadbalancers found for daemonset %v because it has no labels", meta.GetName())
	}

	lbList, err := s.LoadBalancers(meta.GetNamespace()).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var lbs []*api.LoadBalancer
	for _, lb := range lbList {
		// use loadbalancer namespace and name construct unique key
		selector := labels.Set{
			api.LabelKeyCreateby: fmt.Sprintf(api.LabelValueFormatCreateby, lb.Namespace, lb.Name),
		}.AsSelector()

		if !selector.Matches(labels.Set(meta.GetLabels())) {
			// d is not creatby this lb
			continue
		}
		lbs = append(lbs, lb)
	}

	if len(lbs) == 0 {
		return nil, fmt.Errorf("could not find loadbalancer for %v %s in namespace %s with labels: %v", tpy.GetKind(), meta.GetName(), meta.GetNamespace(), meta.GetLabels())
	}

	return lbs, nil
}
