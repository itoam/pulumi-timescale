// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package timescale

import (
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	timescale "github.com/timescale/terraform-provider-timescale/shim" // Import the upstream provider

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

	"github.com/itoam/pulumi-timescale/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "timescale"
	// modules:
	mainMod = "index" // the timescale module
)

//go:embed cmd/pulumi-resource-timescale/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P: pfbridge.ShimProvider(timescale.NewProvider(version.Version)()),

		Name:    "timescale",
		Version: version.Version,
		DisplayName: "Timescale",
		Publisher: "ITOAM",
		// You may host a logo on a domain you control or add an PNG logo (100x100) for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://console.cloud.timescale.com/static/media/timescale-logo-mini.06f9d72ede30459fa1933fc0d3164144.svg",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g. https://github.com/org/pulumi-provider-name/releases/download/v${VERSION}/
		PluginDownloadURL: "https://github.com/itoam/pulumi-timescale/releases/download/v${VERSION}",
		Description:       "A Pulumi package for creating and managing timescale cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"timescale", "category/cloud", "category/database"},
		License:    "Apache-2.0",
		Homepage:   "https://www.timescale.com/cloud",
		Repository: "https://github.com/itoam/pulumi-timescale",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this should
		// match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "timescale",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"timescale_peering_connection": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {Type: "string"},
				},
			},
			"timescale_vpcs": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {Type: "string"},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Enable modern PyProject support in the generated Python SDK.
			PyProject: struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			// Set where the SDK is going to be published to.
			ImportBasePath: path.Join(
				"github.com/itoam/pulumi-timescale/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			// Opt in to all available code generation features.
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Use a wildcard import so NuGet will prefer the latest possible version.
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// MustComputeTokens maps all resources and datasources from the upstream provider into Pulumi.
	//
	// tokens.SingleModule puts every upstream item into your provider's main module.
	//
	// You shouldn't need to override anything, but if you do, use the [tfbridge.ProviderInfo.Resources]
	// and [tfbridge.ProviderInfo.DataSources].
	prov.MustComputeTokens(tokens.SingleModule("timescale_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
