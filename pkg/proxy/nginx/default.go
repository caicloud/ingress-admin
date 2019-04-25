/*
Copyright 2017 Caicloud authors. All rights reserved.

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

package nginx

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	log "k8s.io/klog"
)

const (
	defaultHTTPBackendName      = "default-http-backend"
	defaultHTTPBackendNamespace = metav1.NamespaceSystem
)

var (
	defaultHTTPBackendLabels = map[string]string{
		"app": defaultHTTPBackendName,
	}
)

func (f *nginx) ensureDefaultHTTPBackend() error {
	dp := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: defaultHTTPBackendNamespace,
			Name:      defaultHTTPBackendName,
			Labels:    defaultHTTPBackendLabels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: defaultHTTPBackendLabels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: defaultHTTPBackendLabels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            defaultHTTPBackendName,
							Image:           f.defaultHTTPbackend,
							ImagePullPolicy: v1.PullAlways,
							Resources: v1.ResourceRequirements{
								Limits: v1.ResourceList{
									v1.ResourceCPU:    resource.MustParse("100m"),
									v1.ResourceMemory: resource.MustParse("100Mi"),
								},
							},
							Ports: []v1.ContainerPort{
								{
									ContainerPort: int32(80),
									Protocol:      v1.ProtocolTCP,
								},
							},
						},
					},
				},
			},
		},
	}

	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: defaultHTTPBackendNamespace,
			Name:      defaultHTTPBackendName,
			Labels:    defaultHTTPBackendLabels,
		},
		Spec: v1.ServiceSpec{
			Type:            v1.ServiceTypeClusterIP,
			SessionAffinity: v1.ServiceAffinityNone,
			Selector:        defaultHTTPBackendLabels,
			Ports: []v1.ServicePort{
				{
					Port:       int32(80),
					TargetPort: intstr.FromInt(80),
					Protocol:   v1.ProtocolTCP,
				},
			},
		},
	}

	if _, err := f.client.AppsV1().Deployments(defaultHTTPBackendNamespace).Create(dp); err != nil && !errors.IsAlreadyExists(err) {
		log.Errorf("Cannot create Deployment %v/%v: %v", defaultHTTPBackendNamespace, defaultHTTPBackendName, err)
		return err
	}

	if _, err := f.client.CoreV1().Services(defaultHTTPBackendNamespace).Create(svc); err != nil && !errors.IsAlreadyExists(err) {
		log.Errorf("Cannot create Service %v/%v: %v", defaultHTTPBackendNamespace, defaultHTTPBackendName, err)
		return err
	}

	log.Info("Ensure default-http-backend service for ingress controller success")

	return nil
}
