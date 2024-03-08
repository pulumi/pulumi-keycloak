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
        /// 
        /// ## Example Usage
        /// 
        /// In this example, we'll create an OpenID client with service accounts enabled. This causes Keycloak to create a special user
        /// that represents the service account. We'll use this data source to grab this user's ID in order to assign some roles to this
        /// user, using the `keycloak.UserRoles` resource.
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Keycloak = Pulumi.Keycloak;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var realm = new Keycloak.Realm("realm", new()
        ///     {
        ///         RealmName = "my-realm",
        ///         Enabled = true,
        ///     });
        /// 
        ///     var client = new Keycloak.OpenId.Client("client", new()
        ///     {
        ///         RealmId = realm.Id,
        ///         ClientId = "client",
        ///         AccessType = "CONFIDENTIAL",
        ///         ServiceAccountsEnabled = true,
        ///     });
        /// 
        ///     var serviceAccountUser = Keycloak.OpenId.GetClientServiceAccountUser.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         ClientId = client.Id,
        ///     });
        /// 
        ///     var offlineAccess = Keycloak.GetRole.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         Name = "offline_access",
        ///     });
        /// 
        ///     var serviceAccountUserRoles = new Keycloak.UserRoles("serviceAccountUserRoles", new()
        ///     {
        ///         RealmId = realm.Id,
        ///         UserId = serviceAccountUser.Apply(getClientServiceAccountUserResult =&gt; getClientServiceAccountUserResult.Id),
        ///         RoleIds = new[]
        ///         {
        ///             offlineAccess.Apply(getRoleResult =&gt; getRoleResult.Id),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetClientServiceAccountUserResult> InvokeAsync(GetClientServiceAccountUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientServiceAccountUserResult>("keycloak:openid/getClientServiceAccountUser:getClientServiceAccountUser", args ?? new GetClientServiceAccountUserArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch information about the service account user that is associated with an OpenID client
        /// that has service accounts enabled.
        /// 
        /// ## Example Usage
        /// 
        /// In this example, we'll create an OpenID client with service accounts enabled. This causes Keycloak to create a special user
        /// that represents the service account. We'll use this data source to grab this user's ID in order to assign some roles to this
        /// user, using the `keycloak.UserRoles` resource.
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Keycloak = Pulumi.Keycloak;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var realm = new Keycloak.Realm("realm", new()
        ///     {
        ///         RealmName = "my-realm",
        ///         Enabled = true,
        ///     });
        /// 
        ///     var client = new Keycloak.OpenId.Client("client", new()
        ///     {
        ///         RealmId = realm.Id,
        ///         ClientId = "client",
        ///         AccessType = "CONFIDENTIAL",
        ///         ServiceAccountsEnabled = true,
        ///     });
        /// 
        ///     var serviceAccountUser = Keycloak.OpenId.GetClientServiceAccountUser.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         ClientId = client.Id,
        ///     });
        /// 
        ///     var offlineAccess = Keycloak.GetRole.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         Name = "offline_access",
        ///     });
        /// 
        ///     var serviceAccountUserRoles = new Keycloak.UserRoles("serviceAccountUserRoles", new()
        ///     {
        ///         RealmId = realm.Id,
        ///         UserId = serviceAccountUser.Apply(getClientServiceAccountUserResult =&gt; getClientServiceAccountUserResult.Id),
        ///         RoleIds = new[]
        ///         {
        ///             offlineAccess.Apply(getRoleResult =&gt; getRoleResult.Id),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
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
        public readonly ImmutableArray<string> RequiredActions;
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

            ImmutableArray<string> requiredActions,

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
            RequiredActions = requiredActions;
            Username = username;
        }
    }
}
