// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Outputs
{

    [OutputType]
    public sealed class UserFederatedIdentity
    {
        public readonly string IdentityProvider;
        public readonly string UserId;
        public readonly string UserName;

        [OutputConstructor]
        private UserFederatedIdentity(
            string identityProvider,

            string userId,

            string userName)
        {
            IdentityProvider = identityProvider;
            UserId = userId;
            UserName = userName;
        }
    }
}
