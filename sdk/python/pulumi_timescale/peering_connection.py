# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['PeeringConnectionArgs', 'PeeringConnection']

@pulumi.input_type
class PeeringConnectionArgs:
    def __init__(__self__, *,
                 peer_account_id: pulumi.Input[str],
                 peer_region_code: pulumi.Input[str],
                 peer_vpc_id: pulumi.Input[str],
                 timescale_vpc_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a PeeringConnection resource.
        :param pulumi.Input[str] peer_account_id: AWS account ID where the VPC to be paired
        :param pulumi.Input[str] peer_region_code: Region code for the VPC to be paired
        :param pulumi.Input[str] peer_vpc_id: AWS ID for the VPC to be paired
        :param pulumi.Input[int] timescale_vpc_id: Timescale internal ID for a vpc
        """
        pulumi.set(__self__, "peer_account_id", peer_account_id)
        pulumi.set(__self__, "peer_region_code", peer_region_code)
        pulumi.set(__self__, "peer_vpc_id", peer_vpc_id)
        pulumi.set(__self__, "timescale_vpc_id", timescale_vpc_id)

    @property
    @pulumi.getter(name="peerAccountId")
    def peer_account_id(self) -> pulumi.Input[str]:
        """
        AWS account ID where the VPC to be paired
        """
        return pulumi.get(self, "peer_account_id")

    @peer_account_id.setter
    def peer_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_account_id", value)

    @property
    @pulumi.getter(name="peerRegionCode")
    def peer_region_code(self) -> pulumi.Input[str]:
        """
        Region code for the VPC to be paired
        """
        return pulumi.get(self, "peer_region_code")

    @peer_region_code.setter
    def peer_region_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_region_code", value)

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> pulumi.Input[str]:
        """
        AWS ID for the VPC to be paired
        """
        return pulumi.get(self, "peer_vpc_id")

    @peer_vpc_id.setter
    def peer_vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_vpc_id", value)

    @property
    @pulumi.getter(name="timescaleVpcId")
    def timescale_vpc_id(self) -> pulumi.Input[int]:
        """
        Timescale internal ID for a vpc
        """
        return pulumi.get(self, "timescale_vpc_id")

    @timescale_vpc_id.setter
    def timescale_vpc_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "timescale_vpc_id", value)


@pulumi.input_type
class _PeeringConnectionState:
    def __init__(__self__, *,
                 error_message: Optional[pulumi.Input[str]] = None,
                 peer_account_id: Optional[pulumi.Input[str]] = None,
                 peer_cidr: Optional[pulumi.Input[str]] = None,
                 peer_region_code: Optional[pulumi.Input[str]] = None,
                 peer_vpc_id: Optional[pulumi.Input[str]] = None,
                 provisioned_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 timescale_vpc_id: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PeeringConnection resources.
        :param pulumi.Input[str] peer_account_id: AWS account ID where the VPC to be paired
        :param pulumi.Input[str] peer_cidr: CIDR for the VPC to be paired
        :param pulumi.Input[str] peer_region_code: Region code for the VPC to be paired
        :param pulumi.Input[str] peer_vpc_id: AWS ID for the VPC to be paired
        :param pulumi.Input[str] provisioned_id: AWS ID of the peering connection (starts with pcx-...)
        :param pulumi.Input[str] status: Peering connection status
        :param pulumi.Input[int] timescale_vpc_id: Timescale internal ID for a vpc
        :param pulumi.Input[str] vpc_id: AWS VPC ID of the timescale instance VPC
        """
        if error_message is not None:
            pulumi.set(__self__, "error_message", error_message)
        if peer_account_id is not None:
            pulumi.set(__self__, "peer_account_id", peer_account_id)
        if peer_cidr is not None:
            pulumi.set(__self__, "peer_cidr", peer_cidr)
        if peer_region_code is not None:
            pulumi.set(__self__, "peer_region_code", peer_region_code)
        if peer_vpc_id is not None:
            pulumi.set(__self__, "peer_vpc_id", peer_vpc_id)
        if provisioned_id is not None:
            pulumi.set(__self__, "provisioned_id", provisioned_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if timescale_vpc_id is not None:
            pulumi.set(__self__, "timescale_vpc_id", timescale_vpc_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "error_message")

    @error_message.setter
    def error_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "error_message", value)

    @property
    @pulumi.getter(name="peerAccountId")
    def peer_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID where the VPC to be paired
        """
        return pulumi.get(self, "peer_account_id")

    @peer_account_id.setter
    def peer_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_account_id", value)

    @property
    @pulumi.getter(name="peerCidr")
    def peer_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        CIDR for the VPC to be paired
        """
        return pulumi.get(self, "peer_cidr")

    @peer_cidr.setter
    def peer_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_cidr", value)

    @property
    @pulumi.getter(name="peerRegionCode")
    def peer_region_code(self) -> Optional[pulumi.Input[str]]:
        """
        Region code for the VPC to be paired
        """
        return pulumi.get(self, "peer_region_code")

    @peer_region_code.setter
    def peer_region_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_region_code", value)

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS ID for the VPC to be paired
        """
        return pulumi.get(self, "peer_vpc_id")

    @peer_vpc_id.setter
    def peer_vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_vpc_id", value)

    @property
    @pulumi.getter(name="provisionedId")
    def provisioned_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS ID of the peering connection (starts with pcx-...)
        """
        return pulumi.get(self, "provisioned_id")

    @provisioned_id.setter
    def provisioned_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provisioned_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Peering connection status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="timescaleVpcId")
    def timescale_vpc_id(self) -> Optional[pulumi.Input[int]]:
        """
        Timescale internal ID for a vpc
        """
        return pulumi.get(self, "timescale_vpc_id")

    @timescale_vpc_id.setter
    def timescale_vpc_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timescale_vpc_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS VPC ID of the timescale instance VPC
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class PeeringConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peer_account_id: Optional[pulumi.Input[str]] = None,
                 peer_region_code: Optional[pulumi.Input[str]] = None,
                 peer_vpc_id: Optional[pulumi.Input[str]] = None,
                 timescale_vpc_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Schema for a peering connection. Import can be done with timescale_vpc_id,peer_account_id,peer_region_code,peer_vpc_id format

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peer_account_id: AWS account ID where the VPC to be paired
        :param pulumi.Input[str] peer_region_code: Region code for the VPC to be paired
        :param pulumi.Input[str] peer_vpc_id: AWS ID for the VPC to be paired
        :param pulumi.Input[int] timescale_vpc_id: Timescale internal ID for a vpc
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PeeringConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Schema for a peering connection. Import can be done with timescale_vpc_id,peer_account_id,peer_region_code,peer_vpc_id format

        :param str resource_name: The name of the resource.
        :param PeeringConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PeeringConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peer_account_id: Optional[pulumi.Input[str]] = None,
                 peer_region_code: Optional[pulumi.Input[str]] = None,
                 peer_vpc_id: Optional[pulumi.Input[str]] = None,
                 timescale_vpc_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PeeringConnectionArgs.__new__(PeeringConnectionArgs)

            if peer_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'peer_account_id'")
            __props__.__dict__["peer_account_id"] = peer_account_id
            if peer_region_code is None and not opts.urn:
                raise TypeError("Missing required property 'peer_region_code'")
            __props__.__dict__["peer_region_code"] = peer_region_code
            if peer_vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'peer_vpc_id'")
            __props__.__dict__["peer_vpc_id"] = peer_vpc_id
            if timescale_vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'timescale_vpc_id'")
            __props__.__dict__["timescale_vpc_id"] = timescale_vpc_id
            __props__.__dict__["error_message"] = None
            __props__.__dict__["peer_cidr"] = None
            __props__.__dict__["provisioned_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["vpc_id"] = None
        super(PeeringConnection, __self__).__init__(
            'timescale:index/peeringConnection:PeeringConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            error_message: Optional[pulumi.Input[str]] = None,
            peer_account_id: Optional[pulumi.Input[str]] = None,
            peer_cidr: Optional[pulumi.Input[str]] = None,
            peer_region_code: Optional[pulumi.Input[str]] = None,
            peer_vpc_id: Optional[pulumi.Input[str]] = None,
            provisioned_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            timescale_vpc_id: Optional[pulumi.Input[int]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'PeeringConnection':
        """
        Get an existing PeeringConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peer_account_id: AWS account ID where the VPC to be paired
        :param pulumi.Input[str] peer_cidr: CIDR for the VPC to be paired
        :param pulumi.Input[str] peer_region_code: Region code for the VPC to be paired
        :param pulumi.Input[str] peer_vpc_id: AWS ID for the VPC to be paired
        :param pulumi.Input[str] provisioned_id: AWS ID of the peering connection (starts with pcx-...)
        :param pulumi.Input[str] status: Peering connection status
        :param pulumi.Input[int] timescale_vpc_id: Timescale internal ID for a vpc
        :param pulumi.Input[str] vpc_id: AWS VPC ID of the timescale instance VPC
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PeeringConnectionState.__new__(_PeeringConnectionState)

        __props__.__dict__["error_message"] = error_message
        __props__.__dict__["peer_account_id"] = peer_account_id
        __props__.__dict__["peer_cidr"] = peer_cidr
        __props__.__dict__["peer_region_code"] = peer_region_code
        __props__.__dict__["peer_vpc_id"] = peer_vpc_id
        __props__.__dict__["provisioned_id"] = provisioned_id
        __props__.__dict__["status"] = status
        __props__.__dict__["timescale_vpc_id"] = timescale_vpc_id
        __props__.__dict__["vpc_id"] = vpc_id
        return PeeringConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> pulumi.Output[str]:
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter(name="peerAccountId")
    def peer_account_id(self) -> pulumi.Output[str]:
        """
        AWS account ID where the VPC to be paired
        """
        return pulumi.get(self, "peer_account_id")

    @property
    @pulumi.getter(name="peerCidr")
    def peer_cidr(self) -> pulumi.Output[str]:
        """
        CIDR for the VPC to be paired
        """
        return pulumi.get(self, "peer_cidr")

    @property
    @pulumi.getter(name="peerRegionCode")
    def peer_region_code(self) -> pulumi.Output[str]:
        """
        Region code for the VPC to be paired
        """
        return pulumi.get(self, "peer_region_code")

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> pulumi.Output[str]:
        """
        AWS ID for the VPC to be paired
        """
        return pulumi.get(self, "peer_vpc_id")

    @property
    @pulumi.getter(name="provisionedId")
    def provisioned_id(self) -> pulumi.Output[str]:
        """
        AWS ID of the peering connection (starts with pcx-...)
        """
        return pulumi.get(self, "provisioned_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Peering connection status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="timescaleVpcId")
    def timescale_vpc_id(self) -> pulumi.Output[int]:
        """
        Timescale internal ID for a vpc
        """
        return pulumi.get(self, "timescale_vpc_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        AWS VPC ID of the timescale instance VPC
        """
        return pulumi.get(self, "vpc_id")
