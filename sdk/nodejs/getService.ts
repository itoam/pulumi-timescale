// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Service data source
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("timescale:index/getService:getService", {
        "environmentTag": args.environmentTag,
        "id": args.id,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceArgs {
    /**
     * Environment tag for this service.
     */
    environmentTag?: string;
    /**
     * Service ID is the unique identifier for this service
     */
    id: string;
    /**
     * VPC ID this service is linked to.
     */
    vpcId?: number;
}

/**
 * A collection of values returned by getService.
 */
export interface GetServiceResult {
    /**
     * Created is the time this service was created.
     */
    readonly created: string;
    /**
     * Environment tag for this service.
     */
    readonly environmentTag: string;
    /**
     * Service ID is the unique identifier for this service
     */
    readonly id: string;
    /**
     * Service Name is the configurable name assigned to this resource. If none is provided, a default will be generated by the provider.
     */
    readonly name: string;
    /**
     * Region Code is the physical data center where this service is located.
     */
    readonly regionCode: string;
    readonly resources: outputs.GetServiceResource[];
    readonly spec: outputs.GetServiceSpec;
    /**
     * VPC ID this service is linked to.
     */
    readonly vpcId: number;
}
/**
 * Service data source
 */
export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("timescale:index/getService:getService", {
        "environmentTag": args.environmentTag,
        "id": args.id,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceOutputArgs {
    /**
     * Environment tag for this service.
     */
    environmentTag?: pulumi.Input<string>;
    /**
     * Service ID is the unique identifier for this service
     */
    id: pulumi.Input<string>;
    /**
     * VPC ID this service is linked to.
     */
    vpcId?: pulumi.Input<number>;
}
