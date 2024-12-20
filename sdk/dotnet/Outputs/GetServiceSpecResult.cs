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
    public sealed class GetServiceSpecResult
    {
        /// <summary>
        /// Hostname is the hostname of this service.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// Hostname of the pooler of this service.
        /// </summary>
        public readonly string PoolerHostname;
        /// <summary>
        /// Port of the pooler of this service.
        /// </summary>
        public readonly int PoolerPort;
        /// <summary>
        /// Port is the port assigned to this service.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Hostname of the HA-Replica of this service.
        /// </summary>
        public readonly string ReplicaHostname;
        /// <summary>
        /// Port of the HA-Replica of this service.
        /// </summary>
        public readonly int ReplicaPort;
        /// <summary>
        /// Username is the Postgres username.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetServiceSpecResult(
            string hostname,

            string poolerHostname,

            int poolerPort,

            int port,

            string replicaHostname,

            int replicaPort,

            string username)
        {
            Hostname = hostname;
            PoolerHostname = poolerHostname;
            PoolerPort = poolerPort;
            Port = port;
            ReplicaHostname = replicaHostname;
            ReplicaPort = replicaPort;
            Username = username;
        }
    }
}
