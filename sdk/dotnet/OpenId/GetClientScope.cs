// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public static class GetClientScope
    {
        /// <summary>
        /// This data source can be used to fetch properties of a Keycloak OpenID client scope for usage with other resources.
        /// </summary>
        public static Task<GetClientScopeResult> InvokeAsync(GetClientScopeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientScopeResult>("keycloak:openid/getClientScope:getClientScope", args ?? new GetClientScopeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch properties of a Keycloak OpenID client scope for usage with other resources.
        /// </summary>
        public static Output<GetClientScopeResult> Invoke(GetClientScopeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientScopeResult>("keycloak:openid/getClientScope:getClientScope", args ?? new GetClientScopeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClientScopeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the client scope.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The realm id.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetClientScopeArgs()
        {
        }
        public static new GetClientScopeArgs Empty => new GetClientScopeArgs();
    }

    public sealed class GetClientScopeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the client scope.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The realm id.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GetClientScopeInvokeArgs()
        {
        }
        public static new GetClientScopeInvokeArgs Empty => new GetClientScopeInvokeArgs();
    }


    [OutputType]
    public sealed class GetClientScopeResult
    {
        public readonly string ConsentScreenText;
        public readonly string Description;
        public readonly int GuiOrder;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IncludeInTokenScope;
        public readonly string Name;
        public readonly string RealmId;

        [OutputConstructor]
        private GetClientScopeResult(
            string consentScreenText,

            string description,

            int guiOrder,

            string id,

            bool includeInTokenScope,

            string name,

            string realmId)
        {
            ConsentScreenText = consentScreenText;
            Description = description;
            GuiOrder = guiOrder;
            Id = id;
            IncludeInTokenScope = includeInTokenScope;
            Name = name;
            RealmId = realmId;
        }
    }
}
