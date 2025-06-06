// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class UserFederatedIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the identity provider
        /// </summary>
        [Input("identityProvider", required: true)]
        public Input<string> IdentityProvider { get; set; } = null!;

        /// <summary>
        /// The ID of the user defined in the identity provider
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        /// <summary>
        /// The username of the user defined in the identity provider
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public UserFederatedIdentityArgs()
        {
        }
        public static new UserFederatedIdentityArgs Empty => new UserFederatedIdentityArgs();
    }
}
