# Timescale Resource Provider

The Timescale Resource Provider lets you manage [Timescale](https://www.timescale.com/cloud) resources with [Pulumi](https://www.pulumi.com/).

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install your preferred package manager:

```bash
pnpm add @pulumi/timescale
npm install @pulumi/timescale
yarn add @pulumi/timescale
```

### Python, .NET, and Go

We currently don't host these languages. Instead, you can install the provider from GitHub binaries using `pulumi plugin install`:

```bash
pulumi plugin install resource timescale <version> --server github://api.github.com/itoam/pulumi-timescale
```

## Configuration

The following configuration points are available for the `timescale` provider:

- `timescale:projectId` - the project ID
- `timescale:accessKey` - the access key
- `timescale:secretKey` - the secret key

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/timescale/api-docs/) or [docs](docs/_index.md).
