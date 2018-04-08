/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// This file was automatically generated by lister-gen

package v1beta1

import (
	v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterLister helps list Clusters.
type ClusterLister interface {
	// List lists all Clusters in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.Cluster, err error)
	// Get retrieves the Cluster from the index for a given name.
	Get(name string) (*v1beta1.Cluster, error)
	ClusterListerExpansion
}

// clusterLister implements the ClusterLister interface.
type clusterLister struct {
	indexer cache.Indexer
}

// NewClusterLister returns a new ClusterLister.
func NewClusterLister(indexer cache.Indexer) ClusterLister {
	return &clusterLister{indexer: indexer}
}

// List lists all Clusters in the indexer.
func (s *clusterLister) List(selector labels.Selector) (ret []*v1beta1.Cluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Cluster))
	})
	return ret, err
}

// Get retrieves the Cluster from the index for a given name.
func (s *clusterLister) Get(name string) (*v1beta1.Cluster, error) {
	key := &v1beta1.Cluster{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("cluster"), name)
	}
	return obj.(*v1beta1.Cluster), nil
}
