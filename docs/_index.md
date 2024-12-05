---
title: Timescale
meta_desc: Provides an overview of the Timescale Provider for Pulumi.
layout: package
---

The Timescale Provider lets you manage [Timescale Cloud](https://www.timescale.com/cloud) resources.

## Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as timescale from "@itoam/pulumi-timescale";

const provider = new timescale.Provider("provider", {
  accessKey: "XXXX",
  secretKey: "XXXX",
  projectId: "XXXX",
});

const testService = new timescale.Service(
  "test",
  {
    name: "test",
    milliCpu: 500,
    memoryGb: 2,
  },
  { provider }
);
```
