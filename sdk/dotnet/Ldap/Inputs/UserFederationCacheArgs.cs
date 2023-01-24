// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Ldap.Inputs
{

    public sealed class UserFederationCacheArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Day of the week the entry will become invalid on
        /// </summary>
        [Input("evictionDay")]
        public Input<int>? EvictionDay { get; set; }

        /// <summary>
        /// Hour of day the entry will become invalid on.
        /// </summary>
        [Input("evictionHour")]
        public Input<int>? EvictionHour { get; set; }

        /// <summary>
        /// Minute of day the entry will become invalid on.
        /// </summary>
        [Input("evictionMinute")]
        public Input<int>? EvictionMinute { get; set; }

        /// <summary>
        /// Max lifespan of cache entry (duration string).
        /// </summary>
        [Input("maxLifespan")]
        public Input<string>? MaxLifespan { get; set; }

        /// <summary>
        /// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public UserFederationCacheArgs()
        {
        }
        public static new UserFederationCacheArgs Empty => new UserFederationCacheArgs();
    }
}
