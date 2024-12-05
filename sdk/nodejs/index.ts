// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { GetProductsResult } from "./getProducts";
export const getProducts: typeof import("./getProducts").getProducts = null as any;
export const getProductsOutput: typeof import("./getProducts").getProductsOutput = null as any;
utilities.lazyLoad(exports, ["getProducts","getProductsOutput"], () => require("./getProducts"));

export { GetServiceArgs, GetServiceResult, GetServiceOutputArgs } from "./getService";
export const getService: typeof import("./getService").getService = null as any;
export const getServiceOutput: typeof import("./getService").getServiceOutput = null as any;
utilities.lazyLoad(exports, ["getService","getServiceOutput"], () => require("./getService"));

export { GetVpcsResult } from "./getVpcs";
export const getVpcs: typeof import("./getVpcs").getVpcs = null as any;
export const getVpcsOutput: typeof import("./getVpcs").getVpcsOutput = null as any;
utilities.lazyLoad(exports, ["getVpcs","getVpcsOutput"], () => require("./getVpcs"));

export { PeeringConnectionArgs, PeeringConnectionState } from "./peeringConnection";
export type PeeringConnection = import("./peeringConnection").PeeringConnection;
export const PeeringConnection: typeof import("./peeringConnection").PeeringConnection = null as any;
utilities.lazyLoad(exports, ["PeeringConnection"], () => require("./peeringConnection"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { ServiceArgs, ServiceState } from "./service";
export type Service = import("./service").Service;
export const Service: typeof import("./service").Service = null as any;
utilities.lazyLoad(exports, ["Service"], () => require("./service"));

export { VpcsArgs, VpcsState } from "./vpcs";
export type Vpcs = import("./vpcs").Vpcs;
export const Vpcs: typeof import("./vpcs").Vpcs = null as any;
utilities.lazyLoad(exports, ["Vpcs"], () => require("./vpcs"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "timescale:index/peeringConnection:PeeringConnection":
                return new PeeringConnection(name, <any>undefined, { urn })
            case "timescale:index/service:Service":
                return new Service(name, <any>undefined, { urn })
            case "timescale:index/vpcs:Vpcs":
                return new Vpcs(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("timescale", "index/peeringConnection", _module)
pulumi.runtime.registerResourceModule("timescale", "index/service", _module)
pulumi.runtime.registerResourceModule("timescale", "index/vpcs", _module)
pulumi.runtime.registerResourcePackage("timescale", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:timescale") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});