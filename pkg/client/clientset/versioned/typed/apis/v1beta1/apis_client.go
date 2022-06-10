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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	rest "k8s.io/client-go/rest"
	v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
	"sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/scheme"
)

type GatewayV1beta1Interface interface {
	RESTClient() rest.Interface
	GatewaysGetter
	GatewayClassesGetter
	HTTPRoutesGetter
}

// GatewayV1beta1Client is used to interact with features provided by the gateway.networking.k8s.io group.
type GatewayV1beta1Client struct {
	restClient rest.Interface
}

func (c *GatewayV1beta1Client) Gateways(namespace string) GatewayInterface {
	return newGateways(c, namespace)
}

func (c *GatewayV1beta1Client) GatewayClasses() GatewayClassInterface {
	return newGatewayClasses(c)
}

func (c *GatewayV1beta1Client) HTTPRoutes(namespace string) HTTPRouteInterface {
	return newHTTPRoutes(c, namespace)
}

// NewForConfig creates a new GatewayV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*GatewayV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &GatewayV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new GatewayV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GatewayV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GatewayV1beta1Client for the given RESTClient.
func New(c rest.Interface) *GatewayV1beta1Client {
	return &GatewayV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *GatewayV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
