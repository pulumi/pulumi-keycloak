// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public static class GetGroup
    {
        /// <summary>
        /// This data source can be used to fetch properties of a Keycloak group for
        /// usage with other resources, such as `keycloak.GroupRoles`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Keycloak = Pulumi.Keycloak;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
        ///         {
        ///             Realm = "my-realm",
        ///             Enabled = true,
        ///         });
        ///         var offlineAccess = realm.Id.Apply(id =&gt; Keycloak.GetRole.InvokeAsync(new Keycloak.GetRoleArgs
        ///         {
        ///             RealmId = id,
        ///             Name = "offline_access",
        ///         }));
        ///         var @group = realm.Id.Apply(id =&gt; Keycloak.GetGroup.InvokeAsync(new Keycloak.GetGroupArgs
        ///         {
        ///             RealmId = id,
        ///             Name = "group",
        ///         }));
        ///         var groupRoles = new Keycloak.GroupRoles("groupRoles", new Keycloak.GroupRolesArgs
        ///         {
        ///             RealmId = realm.Id,
        ///             GroupId = @group.Apply(@group =&gt; @group.Id),
        ///             RoleIds = 
        ///             {
        ///                 offlineAccess.Apply(offlineAccess =&gt; offlineAccess.Id),
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("keycloak:index/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithVersion());
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the group. If there are multiple groups match `name`, the first result will be returned.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The realm this group exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        public readonly ImmutableDictionary<string, object> Attributes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ParentId;
        public readonly string Path;
        public readonly string RealmId;

        [OutputConstructor]
        private GetGroupResult(
            ImmutableDictionary<string, object> attributes,

            string id,

            string name,

            string parentId,

            string path,

            string realmId)
        {
            Attributes = attributes;
            Id = id;
            Name = name;
            ParentId = parentId;
            Path = path;
            RealmId = realmId;
        }
    }
}
