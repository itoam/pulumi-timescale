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
from . import outputs

__all__ = [
    'GetProductsResult',
    'AwaitableGetProductsResult',
    'get_products',
    'get_products_output',
]

@pulumi.output_type
class GetProductsResult:
    """
    A collection of values returned by getProducts.
    """
    def __init__(__self__, id=None, products=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if products and not isinstance(products, list):
            raise TypeError("Expected argument 'products' to be a list")
        pulumi.set(__self__, "products", products)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def products(self) -> Sequence['outputs.GetProductsProductResult']:
        return pulumi.get(self, "products")


class AwaitableGetProductsResult(GetProductsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProductsResult(
            id=self.id,
            products=self.products)


def get_products(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProductsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('timescale:index/getProducts:getProducts', __args__, opts=opts, typ=GetProductsResult).value

    return AwaitableGetProductsResult(
        id=pulumi.get(__ret__, 'id'),
        products=pulumi.get(__ret__, 'products'))
def get_products_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProductsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('timescale:index/getProducts:getProducts', __args__, opts=opts, typ=GetProductsResult)
    return __ret__.apply(lambda __response__: GetProductsResult(
        id=pulumi.get(__response__, 'id'),
        products=pulumi.get(__response__, 'products')))
