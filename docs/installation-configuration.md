---
title: Timescale Installation & Configuration
meta_desc: Information on how to install the Timescale provider.
layout: package
---

## Installation

The Timescale provider is available as a package in some Pulumi languages:

- JavaScript/TypeScript: [`@itoam/pulumi-timescale`](https://www.npmjs.com/package/@itoam/pulumi-timescale)
- Python: [`pulumi_timescale`](https://pypi.org/project/pulumi-timescale/)
- Go: [`github.com/itoam/pulumi-timescale/sdk/go/timescale`](https://pkg.go.dev/github.com/itoam/pulumi-timescale/sdk)
- .NET: Not supported

### Authorization

When you log in to your [Timescale Account](https://console.cloud.timescale.com/), navigate to the `Project settings` page.
From here, you can create client credentials for programmatic usage. Click the `Create credentials` button to generate a new public/secret key pair.

Find more information on creating Client Credentials in the [Timescale docs](https://docs.timescale.com/use-timescale/latest/security/client-credentials/#creating-client-credentials).

## Configuration Options

Use `pulumi config set timescale:<option> (--secret)`.

| Option      | Environment Variable | Required/Optional | Description          |
| ----------- | -------------------- | ----------------- | -------------------- |
| `projectId` |                      | Required          | Timescale project ID |
| `accessKey` |                      | Required          | Timescale access key |
| `secretKey` |                      | Required          | Timescale secret key |
