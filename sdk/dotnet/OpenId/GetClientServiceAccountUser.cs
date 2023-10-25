// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public static class GetClientServiceAccountUser
    {
        /// <summary>
        /// This data source can be used to fetch information about the service account user that is associated with an OpenID client
        /// that has service accounts enabled.
        /// </summary>
        public static Task<GetClientServiceAccountUserResult> InvokeAsync(GetClientServiceAccountUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientServiceAccountUserResult>("keycloak:openid/getClientServiceAccountUser:getClientServiceAccountUser", args ?? new GetClientServiceAccountUserArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch information about the service account user that is associated with an OpenID client
        /// that has service accounts enabled.
        /// </summary>
        public static Output<GetClientServiceAccountUserResult> Invoke(GetClientServiceAccountUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientServiceAccountUserResult>("keycloak:openid/getClientServiceAccountUser:getClientServiceAccountUser", args ?? new GetClientServiceAccountUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClientServiceAccountUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the OpenID client with service accounts enabled.
        /// </summary>
        [Input("clientId", required: true)]
        public string ClientId { get; set; } = null!;

        /// <summary>
        /// The realm that the OpenID client exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetClientServiceAccountUserArgs()
        {
        }
        public static new GetClientServiceAccountUserArgs Empty => new GetClientServiceAccountUserArgs();
    }

    public sealed class GetClientServiceAccountUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the OpenID client with service accounts enabled.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The realm that the OpenID client exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GetClientServiceAccountUserInvokeArgs()
        {
        }
        public static new GetClientServiceAccountUserInvokeArgs Empty => new GetClientServiceAccountUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetClientServiceAccountUserResult
    {
        public readonly ImmutableDictionary<string, object> Attributes;
        public readonly string ClientId;
        public readonly string Email;
        public readonly bool EmailVerified;
        public readonly bool Enabled;
        public readonly ImmutableArray<Outputs.GetClientServiceAccountUserFederatedIdentityResult> FederatedIdentities;
        public readonly string FirstName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LastName;
        public readonly string RealmId;
        public readonly string Username;

        [OutputConstructor]
        private GetClientServiceAccountUserResult(
            ImmutableDictionary<string, object> attributes,

            string clientId,

            string email,

            bool emailVerified,

            bool enabled,

            ImmutableArray<Outputs.GetClientServiceAccountUserFederatedIdentityResult> federatedIdentities,

            string firstName,

            string id,

            string lastName,

            string realmId,

            string username)
        {
            Attributes = attributes;
            ClientId = clientId;
            Email = email;
            EmailVerified = emailVerified;
            Enabled = enabled;
            FederatedIdentities = federatedIdentities;
            FirstName = firstName;
            Id = id;
            LastName = lastName;
            RealmId = realmId;
            Username = username;
        }
    }
}
