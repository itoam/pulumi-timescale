// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Timescale
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("timescale");

        private static readonly __Value<string?> _accessKey = new __Value<string?>(() => __config.Get("accessKey"));
        /// <summary>
        /// Access Key
        /// </summary>
        public static string? AccessKey
        {
            get => _accessKey.Get();
            set => _accessKey.Set(value);
        }

        private static readonly __Value<string?> _accessToken = new __Value<string?>(() => __config.Get("accessToken"));
        /// <summary>
        /// Access Token
        /// </summary>
        public static string? AccessToken
        {
            get => _accessToken.Get();
            set => _accessToken.Set(value);
        }

        private static readonly __Value<string?> _projectId = new __Value<string?>(() => __config.Get("projectId"));
        /// <summary>
        /// Project ID
        /// </summary>
        public static string? ProjectId
        {
            get => _projectId.Get();
            set => _projectId.Set(value);
        }

        private static readonly __Value<string?> _secretKey = new __Value<string?>(() => __config.Get("secretKey"));
        /// <summary>
        /// Secret Key
        /// </summary>
        public static string? SecretKey
        {
            get => _secretKey.Get();
            set => _secretKey.Set(value);
        }

    }
}
