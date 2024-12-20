// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Timescale.Outputs
{

    [OutputType]
    public sealed class GetProductsProductPlanResult
    {
        public readonly string Id;
        public readonly int MemoryGb;
        public readonly int MilliCpu;
        public readonly double Price;
        public readonly string ProductId;
        public readonly string RegionCode;

        [OutputConstructor]
        private GetProductsProductPlanResult(
            string id,

            int memoryGb,

            int milliCpu,

            double price,

            string productId,

            string regionCode)
        {
            Id = id;
            MemoryGb = memoryGb;
            MilliCpu = milliCpu;
            Price = price;
            ProductId = productId;
            RegionCode = regionCode;
        }
    }
}
