// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newCustomDomainAssociationResource,
			Name:    "Custom Domain Association",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCredentials,
			TypeName: "aws_redshiftserverless_credentials",
			Name:     "Credentials",
		},
		{
			Factory:  dataSourceNamespace,
			TypeName: "aws_redshiftserverless_namespace",
			Name:     "Namespace",
		},
		{
			Factory:  dataSourceWorkgroup,
			TypeName: "aws_redshiftserverless_workgroup",
			Name:     "Workgroup",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceEndpointAccess,
			TypeName: "aws_redshiftserverless_endpoint_access",
			Name:     "Endpoint Access",
		},
		{
			Factory:  resourceNamespace,
			TypeName: "aws_redshiftserverless_namespace",
			Name:     "Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceResourcePolicy,
			TypeName: "aws_redshiftserverless_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  resourceSnapshot,
			TypeName: "aws_redshiftserverless_snapshot",
			Name:     "Snapshot",
		},
		{
			Factory:  resourceUsageLimit,
			TypeName: "aws_redshiftserverless_usage_limit",
			Name:     "Usage Limit",
		},
		{
			Factory:  resourceWorkgroup,
			TypeName: "aws_redshiftserverless_workgroup",
			Name:     "Workgroup",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RedshiftServerless
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*redshiftserverless.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return redshiftserverless.NewFromConfig(cfg,
		redshiftserverless.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
