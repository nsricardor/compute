/*
Copyright 2024 the Unikorn Authors.

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

package cluster

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/pflag"

	unikornv1 "github.com/unikorn-cloud/baremetal/pkg/apis/unikorn/v1alpha1"
	"github.com/unikorn-cloud/baremetal/pkg/constants"
	baremetalprovisioners "github.com/unikorn-cloud/baremetal/pkg/provisioners"
	unikornv1core "github.com/unikorn-cloud/core/pkg/apis/unikorn/v1alpha1"
	coreclient "github.com/unikorn-cloud/core/pkg/client"
	coreconstants "github.com/unikorn-cloud/core/pkg/constants"
	coreerrors "github.com/unikorn-cloud/core/pkg/errors"
	"github.com/unikorn-cloud/core/pkg/manager"
	coreapi "github.com/unikorn-cloud/core/pkg/openapi"
	"github.com/unikorn-cloud/core/pkg/provisioners"
	identityclient "github.com/unikorn-cloud/identity/pkg/client"
	regionclient "github.com/unikorn-cloud/region/pkg/client"
	regionapi "github.com/unikorn-cloud/region/pkg/openapi"

	"sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	ErrAnnotation = errors.New("required annotation missing")

	ErrResourceDependency = errors.New("resource deplendedncy error")
)

// Options allows access to CLI options in the provisioner.
type Options struct {
	// identityOptions allow the identity host and CA to be set.
	identityOptions *identityclient.Options
	// regionOptions allows the region host and CA to be set.
	regionOptions *regionclient.Options
	// clientOptions give access to client certificate information as
	// we need to talk to identity to get a token, and then to region
	// to ensure cloud identities and networks are provisioned, as well
	// as deptovisioning them.
	clientOptions coreclient.HTTPClientOptions
}

func (o *Options) AddFlags(f *pflag.FlagSet) {
	if o.identityOptions == nil {
		o.identityOptions = identityclient.NewOptions()
	}

	if o.regionOptions == nil {
		o.regionOptions = regionclient.NewOptions()
	}

	o.identityOptions.AddFlags(f)
	o.regionOptions.AddFlags(f)
	o.clientOptions.AddFlags(f)
}

// Provisioner encapsulates control plane provisioning.
type Provisioner struct {
	provisioners.Metadata

	// cluster is the baremetal cluster we're provisioning.
	cluster unikornv1.BaremetalCluster

	// options are documented for the type.
	options *Options
}

// New returns a new initialized provisioner object.
func New(options manager.ControllerOptions) provisioners.ManagerProvisioner {
	o, _ := options.(*Options)

	return &Provisioner{
		options: o,
	}
}

// Ensure the ManagerProvisioner interface is implemented.
var _ provisioners.ManagerProvisioner = &Provisioner{}

func (p *Provisioner) Object() unikornv1core.ManagableResourceInterface {
	return &p.cluster
}

// getRegionClient returns an authenticated context with a client credentials access token
// and a client.  The context must be used by subseqent API calls in order to extract
// the access token.
func (p *Provisioner) getRegionClient(ctx context.Context, traceName string) (context.Context, regionapi.ClientWithResponsesInterface, error) {
	cli, err := coreclient.ProvisionerClientFromContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	tokenIssuer := identityclient.NewTokenIssuer(cli, p.options.identityOptions, &p.options.clientOptions, constants.Application, constants.Version)

	ctx, err = tokenIssuer.Context(ctx, traceName)
	if err != nil {
		return nil, nil, err
	}

	getter := regionclient.New(cli, p.options.regionOptions, &p.options.clientOptions)

	client, err := getter.Client(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ctx, client, nil
}

func (p *Provisioner) getIdentity(ctx context.Context, client regionapi.ClientWithResponsesInterface) (*regionapi.IdentityRead, error) {
	log := log.FromContext(ctx)

	response, err := client.GetApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDWithResponse(ctx, p.cluster.Labels[coreconstants.OrganizationLabel], p.cluster.Labels[coreconstants.ProjectLabel], p.cluster.Annotations[coreconstants.IdentityAnnotation])
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("%w: identity GET expected 200 got %d", coreerrors.ErrAPIStatus, response.StatusCode())
	}

	resource := response.JSON200

	//nolint:exhaustive
	switch resource.Metadata.ProvisioningStatus {
	case coreapi.ResourceProvisioningStatusProvisioned:
		return resource, nil
	case coreapi.ResourceProvisioningStatusUnknown, coreapi.ResourceProvisioningStatusProvisioning:
		log.Info("waiting for identity to become ready")

		return nil, provisioners.ErrYield
	}

	return nil, fmt.Errorf("%w: unhandled status %s", ErrResourceDependency, resource.Metadata.ProvisioningStatus)
}

func (p *Provisioner) deleteIdentity(ctx context.Context, client regionapi.ClientWithResponsesInterface) error {
	response, err := client.DeleteApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDWithResponse(ctx, p.cluster.Labels[coreconstants.OrganizationLabel], p.cluster.Labels[coreconstants.ProjectLabel], p.cluster.Annotations[coreconstants.IdentityAnnotation])
	if err != nil {
		return err
	}

	statusCode := response.StatusCode()

	// An accepted status means the API has recoded the deletion event and
	// we can delete the cluster, a not found means it's been deleted already
	// and again can proceed.  The goal here is not to leak resources.
	if statusCode != http.StatusAccepted && statusCode != http.StatusNotFound {
		return fmt.Errorf("%w: identity DELETE expected 202,404 got %d", ErrResourceDependency, statusCode)
	}

	return nil
}

func (p *Provisioner) getPhysicalNetwork(ctx context.Context, client regionapi.ClientWithResponsesInterface) (*regionapi.PhysicalNetworkRead, error) {
	log := log.FromContext(ctx)

	physicalNetworkID, ok := p.cluster.Annotations[coreconstants.PhysicalNetworkAnnotation]
	if !ok {
		//nolint: nilnil
		return nil, nil
	}

	response, err := client.GetApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDPhysicalnetworksPhysicalNetworkIDWithResponse(ctx, p.cluster.Labels[coreconstants.OrganizationLabel], p.cluster.Labels[coreconstants.ProjectLabel], p.cluster.Annotations[coreconstants.IdentityAnnotation], physicalNetworkID)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("%w: physical network GET expected 200 got %d", coreerrors.ErrAPIStatus, response.StatusCode())
	}

	resource := response.JSON200

	//nolint:exhaustive
	switch resource.Metadata.ProvisioningStatus {
	case coreapi.ResourceProvisioningStatusProvisioned:
		return resource, nil
	case coreapi.ResourceProvisioningStatusUnknown, coreapi.ResourceProvisioningStatusProvisioning:
		log.Info("waiting for physical network to become ready")

		return nil, provisioners.ErrYield
	}

	return nil, fmt.Errorf("%w: unhandled status %s", ErrResourceDependency, resource.Metadata.ProvisioningStatus)
}

func (p *Provisioner) identityOptions(ctx context.Context, client regionapi.ClientWithResponsesInterface) (*baremetalprovisioners.ClusterOpenstackOptions, error) {
	identity, err := p.getIdentity(ctx, client)
	if err != nil {
		return nil, err
	}

	physicalNetwork, err := p.getPhysicalNetwork(ctx, client)
	if err != nil {
		return nil, err
	}

	options := &baremetalprovisioners.ClusterOpenstackOptions{
		CloudConfig:   *identity.Spec.Openstack.CloudConfig,
		Cloud:         *identity.Spec.Openstack.Cloud,
		ServerGroupID: identity.Spec.Openstack.ServerGroupId,
		ProviderNetwork: &baremetalprovisioners.ClusterOpenstackProviderOptions{
			NetworkID: physicalNetwork.Spec.Openstack.NetworkId,
			SubnetID:  physicalNetwork.Spec.Openstack.SubnetId,
		},
	}

	return options, nil
}

// Provision implements the Provision interface.
func (p *Provisioner) Provision(ctx context.Context) error {
	// Likewise identity creation is provisioned asynchronously as it too takes a
	// long time, epspectially if a physical network is being provisioned and that
	// needs to go out and talk to swiches.
	clientContext, client, err := p.getRegionClient(ctx, "provision")
	if err != nil {
		return err
	}

	_, err = p.identityOptions(clientContext, client)
	if err != nil {
		return err
	}

	// TODO: do something!
	return nil
}

// Deprovision implements the Provision interface.
func (p *Provisioner) Deprovision(ctx context.Context) error {
	// TODO: do something!
	// Clean up the identity when everything has cleanly deprovisioned.
	// An accepted status means the API has recoded the deletion event and
	// we can delete the cluster, a not found means it's been deleted already
	// and again can proceed.  The goal here is not to leak resources.
	clientContext, client, err := p.getRegionClient(ctx, "deprovision")
	if err != nil {
		return err
	}

	if err := p.deleteIdentity(clientContext, client); err != nil {
		return err
	}

	return nil
}