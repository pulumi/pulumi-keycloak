// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public static class GetUser
    {
        /// <summary>
        /// This data source can be used to fetch properties of a user within Keycloak.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Keycloak = Pulumi.Keycloak;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var masterRealm = Keycloak.GetRealm.Invoke(new()
        ///     {
        ///         Realm = "master",
        ///     });
        /// 
        ///     // use the keycloak_user data source to grab the admin user's ID
        ///     var defaultAdminUser = Keycloak.GetUser.Invoke(new()
        ///     {
        ///         RealmId = masterRealm.Apply(getRealmResult =&gt; getRealmResult.Id),
        ///         Username = "keycloak",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["keycloakUserId"] = defaultAdminUser.Apply(getUserResult =&gt; getUserResult.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("keycloak:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch properties of a user within Keycloak.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Keycloak = Pulumi.Keycloak;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var masterRealm = Keycloak.GetRealm.Invoke(new()
        ///     {
        ///         Realm = "master",
        ///     });
        /// 
        ///     // use the keycloak_user data source to grab the admin user's ID
        ///     var defaultAdminUser = Keycloak.GetUser.Invoke(new()
        ///     {
        ///         RealmId = masterRealm.Apply(getRealmResult =&gt; getRealmResult.Id),
        ///         Username = "keycloak",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["keycloakUserId"] = defaultAdminUser.Apply(getUserResult =&gt; getUserResult.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("keycloak:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch properties of a user within Keycloak.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Keycloak = Pulumi.Keycloak;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var masterRealm = Keycloak.GetRealm.Invoke(new()
        ///     {
        ///         Realm = "master",
        ///     });
        /// 
        ///     // use the keycloak_user data source to grab the admin user's ID
        ///     var defaultAdminUser = Keycloak.GetUser.Invoke(new()
        ///     {
        ///         RealmId = masterRealm.Apply(getRealmResult =&gt; getRealmResult.Id),
        ///         Username = "keycloak",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["keycloakUserId"] = defaultAdminUser.Apply(getUserResult =&gt; getUserResult.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("keycloak:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The realm this user belongs to.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        /// <summary>
        /// The unique username of this user.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The realm this user belongs to.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// The unique username of this user.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// (Computed) A map representing attributes for the user
        /// </summary>
        public readonly ImmutableDictionary<string, string> Attributes;
        /// <summary>
        /// (Computed) The user's email.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// (Computed) Whether the email address was validated or not. Default to `false`.
        /// </summary>
        public readonly bool EmailVerified;
        /// <summary>
        /// (Computed) When false, this user cannot log in. Defaults to `true`.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (Computed) The user's federated identities, if applicable. This block has the following schema:
        /// </summary>
        public readonly ImmutableArray<string> FederatedIdentities;
        /// <summary>
        /// (Computed) The user's first name.
        /// </summary>
        public readonly string FirstName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) The user's last name.
        /// </summary>
        public readonly string LastName;
        public readonly string RealmId;
        public readonly ImmutableArray<string> RequiredActions;
        public readonly string Username;

        [OutputConstructor]
        private GetUserResult(
            ImmutableDictionary<string, string> attributes,

            string email,

            bool emailVerified,

            bool enabled,

            ImmutableArray<string> federatedIdentities,

            string firstName,

            string id,

            string lastName,

            string realmId,

            ImmutableArray<string> requiredActions,

            string username)
        {
            Attributes = attributes;
            Email = email;
            EmailVerified = emailVerified;
            Enabled = enabled;
            FederatedIdentities = federatedIdentities;
            FirstName = firstName;
            Id = id;
            LastName = lastName;
            RealmId = realmId;
            RequiredActions = requiredActions;
            Username = username;
        }
    }
}
