// Code generated by skv2. DO NOT EDIT.

package v1beta1

import (
	networking_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for WasmDeploymentClient from Clientset
func WasmDeploymentClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeploymentClient {
	return clients.WasmDeployments()
}

// Provider for WasmDeployment Client from Client
func WasmDeploymentClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeploymentClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewWasmDeploymentClient(client)
}

type WasmDeploymentClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeploymentClient

func WasmDeploymentClientFactoryProvider() WasmDeploymentClientFactory {
	return WasmDeploymentClientProvider
}

type WasmDeploymentClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeploymentClient, error)

func WasmDeploymentClientFromConfigFactoryProvider() WasmDeploymentClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeploymentClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.WasmDeployments(), nil
	}
}

// Provider for RateLimitClientConfigClient from Clientset
func RateLimitClientConfigClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitClientConfigClient {
	return clients.RateLimitClientConfigs()
}

// Provider for RateLimitClientConfig Client from Client
func RateLimitClientConfigClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitClientConfigClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewRateLimitClientConfigClient(client)
}

type RateLimitClientConfigClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitClientConfigClient

func RateLimitClientConfigClientFactoryProvider() RateLimitClientConfigClientFactory {
	return RateLimitClientConfigClientProvider
}

type RateLimitClientConfigClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitClientConfigClient, error)

func RateLimitClientConfigClientFromConfigFactoryProvider() RateLimitClientConfigClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitClientConfigClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.RateLimitClientConfigs(), nil
	}
}

// Provider for RateLimitServerConfigClient from Clientset
func RateLimitServerConfigClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitServerConfigClient {
	return clients.RateLimitServerConfigs()
}

// Provider for RateLimitServerConfig Client from Client
func RateLimitServerConfigClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitServerConfigClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewRateLimitServerConfigClient(client)
}

type RateLimitServerConfigClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitServerConfigClient

func RateLimitServerConfigClientFactoryProvider() RateLimitServerConfigClientFactory {
	return RateLimitServerConfigClientProvider
}

type RateLimitServerConfigClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitServerConfigClient, error)

func RateLimitServerConfigClientFromConfigFactoryProvider() RateLimitServerConfigClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.RateLimitServerConfigClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.RateLimitServerConfigs(), nil
	}
}

// Provider for VirtualDestinationClient from Clientset
func VirtualDestinationClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestinationClient {
	return clients.VirtualDestinations()
}

// Provider for VirtualDestination Client from Client
func VirtualDestinationClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestinationClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewVirtualDestinationClient(client)
}

type VirtualDestinationClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestinationClient

func VirtualDestinationClientFactoryProvider() VirtualDestinationClientFactory {
	return VirtualDestinationClientProvider
}

type VirtualDestinationClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestinationClient, error)

func VirtualDestinationClientFromConfigFactoryProvider() VirtualDestinationClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestinationClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualDestinations(), nil
	}
}

// Provider for VirtualGatewayClient from Clientset
func VirtualGatewayClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGatewayClient {
	return clients.VirtualGateways()
}

// Provider for VirtualGateway Client from Client
func VirtualGatewayClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGatewayClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewVirtualGatewayClient(client)
}

type VirtualGatewayClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGatewayClient

func VirtualGatewayClientFactoryProvider() VirtualGatewayClientFactory {
	return VirtualGatewayClientProvider
}

type VirtualGatewayClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGatewayClient, error)

func VirtualGatewayClientFromConfigFactoryProvider() VirtualGatewayClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGatewayClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualGateways(), nil
	}
}

// Provider for VirtualHostClient from Clientset
func VirtualHostClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHostClient {
	return clients.VirtualHosts()
}

// Provider for VirtualHost Client from Client
func VirtualHostClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHostClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewVirtualHostClient(client)
}

type VirtualHostClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHostClient

func VirtualHostClientFactoryProvider() VirtualHostClientFactory {
	return VirtualHostClientProvider
}

type VirtualHostClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHostClient, error)

func VirtualHostClientFromConfigFactoryProvider() VirtualHostClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHostClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualHosts(), nil
	}
}

// Provider for RouteTableClient from Clientset
func RouteTableClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTableClient {
	return clients.RouteTables()
}

// Provider for RouteTable Client from Client
func RouteTableClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTableClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewRouteTableClient(client)
}

type RouteTableClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTableClient

func RouteTableClientFactoryProvider() RouteTableClientFactory {
	return RouteTableClientProvider
}

type RouteTableClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTableClient, error)

func RouteTableClientFromConfigFactoryProvider() RouteTableClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTableClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.RouteTables(), nil
	}
}

// Provider for ServiceDependencyClient from Clientset
func ServiceDependencyClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependencyClient {
	return clients.ServiceDependencies()
}

// Provider for ServiceDependency Client from Client
func ServiceDependencyClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependencyClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewServiceDependencyClient(client)
}

type ServiceDependencyClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependencyClient

func ServiceDependencyClientFactoryProvider() ServiceDependencyClientFactory {
	return ServiceDependencyClientProvider
}

type ServiceDependencyClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependencyClient, error)

func ServiceDependencyClientFromConfigFactoryProvider() ServiceDependencyClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependencyClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ServiceDependencies(), nil
	}
}

// Provider for CertificateVerificationClient from Clientset
func CertificateVerificationClientFromClientsetProvider(clients networking_enterprise_mesh_gloo_solo_io_v1beta1.Clientset) networking_enterprise_mesh_gloo_solo_io_v1beta1.CertificateVerificationClient {
	return clients.CertificateVerifications()
}

// Provider for CertificateVerification Client from Client
func CertificateVerificationClientProvider(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.CertificateVerificationClient {
	return networking_enterprise_mesh_gloo_solo_io_v1beta1.NewCertificateVerificationClient(client)
}

type CertificateVerificationClientFactory func(client client.Client) networking_enterprise_mesh_gloo_solo_io_v1beta1.CertificateVerificationClient

func CertificateVerificationClientFactoryProvider() CertificateVerificationClientFactory {
	return CertificateVerificationClientProvider
}

type CertificateVerificationClientFromConfigFactory func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.CertificateVerificationClient, error)

func CertificateVerificationClientFromConfigFactoryProvider() CertificateVerificationClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_enterprise_mesh_gloo_solo_io_v1beta1.CertificateVerificationClient, error) {
		clients, err := networking_enterprise_mesh_gloo_solo_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.CertificateVerifications(), nil
	}
}
