import * as pulumi from "@pulumi/pulumi";
import * as timescale from "@itoam/pulumi-timescale";

const config = new pulumi.Config();
const accessKey = config.require("accessKey");
const secretKey = config.require("secretKey");
const projectId = config.require("projectId");

const provider = new timescale.Provider("provider", {
  accessKey: "XXX",
  secretKey: "XXX",
  projectId: "XXX",
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

const testReplica = new timescale.Service(
  "test-replica",
  {
    name: "test-replica",
    readReplicaSource: testService.id,
  },
  { provider }
);

export const testHost = pulumi.interpolate`${testService.hostname}:${testService.port}`;
export const testReplicaHost = pulumi.interpolate`${testReplica.hostname}:${testReplica.port}`;
