// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml
{
    public static class GetClientInstallationProvider
    {
        /// <summary>
        /// This data source can be used to retrieve Installation Provider of a SAML Client.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// In the example below, we extract the SAML metadata IDPSSODescriptor to pass it to the AWS IAM SAML Provider.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.IO;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
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
        ///     var samlClient = new Keycloak.Saml.Client("samlClient", new()
        ///     {
        ///         RealmId = realm.Id,
        ///         ClientId = "test-saml-client",
        ///         SignDocuments = false,
        ///         SignAssertions = true,
        ///         IncludeAuthnStatement = true,
        ///         SigningCertificate = File.ReadAllText("saml-cert.pem"),
        ///         SigningPrivateKey = File.ReadAllText("saml-key.pem"),
        ///     });
        /// 
        ///     var samlIdpDescriptor = Keycloak.Saml.GetClientInstallationProvider.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         ClientId = samlClient.Id,
        ///         ProviderId = "saml-idp-descriptor",
        ///     });
        /// 
        ///     var @default = new Aws.Iam.SamlProvider("default", new()
        ///     {
        ///         SamlMetadataDocument = samlIdpDescriptor.Apply(getClientInstallationProviderResult =&gt; getClientInstallationProviderResult.Value),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClientInstallationProviderResult> InvokeAsync(GetClientInstallationProviderArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientInstallationProviderResult>("keycloak:saml/getClientInstallationProvider:getClientInstallationProvider", args ?? new GetClientInstallationProviderArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to retrieve Installation Provider of a SAML Client.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// In the example below, we extract the SAML metadata IDPSSODescriptor to pass it to the AWS IAM SAML Provider.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.IO;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
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
        ///     var samlClient = new Keycloak.Saml.Client("samlClient", new()
        ///     {
        ///         RealmId = realm.Id,
        ///         ClientId = "test-saml-client",
        ///         SignDocuments = false,
        ///         SignAssertions = true,
        ///         IncludeAuthnStatement = true,
        ///         SigningCertificate = File.ReadAllText("saml-cert.pem"),
        ///         SigningPrivateKey = File.ReadAllText("saml-key.pem"),
        ///     });
        /// 
        ///     var samlIdpDescriptor = Keycloak.Saml.GetClientInstallationProvider.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         ClientId = samlClient.Id,
        ///         ProviderId = "saml-idp-descriptor",
        ///     });
        /// 
        ///     var @default = new Aws.Iam.SamlProvider("default", new()
        ///     {
        ///         SamlMetadataDocument = samlIdpDescriptor.Apply(getClientInstallationProviderResult =&gt; getClientInstallationProviderResult.Value),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClientInstallationProviderResult> Invoke(GetClientInstallationProviderInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientInstallationProviderResult>("keycloak:saml/getClientInstallationProvider:getClientInstallationProvider", args ?? new GetClientInstallationProviderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClientInstallationProviderArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the SAML client. The `id` attribute of a `keycloak_client` resource should be used here.
        /// </summary>
        [Input("clientId", required: true)]
        public string ClientId { get; set; } = null!;

        /// <summary>
        /// The ID of the SAML installation provider. Could be one of `saml-idp-descriptor`, `keycloak-saml`, `saml-sp-descriptor`, `keycloak-saml-subsystem`, `mod-auth-mellon`, etc.
        /// </summary>
        [Input("providerId", required: true)]
        public string ProviderId { get; set; } = null!;

        /// <summary>
        /// The realm that the SAML client exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetClientInstallationProviderArgs()
        {
        }
        public static new GetClientInstallationProviderArgs Empty => new GetClientInstallationProviderArgs();
    }

    public sealed class GetClientInstallationProviderInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the SAML client. The `id` attribute of a `keycloak_client` resource should be used here.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The ID of the SAML installation provider. Could be one of `saml-idp-descriptor`, `keycloak-saml`, `saml-sp-descriptor`, `keycloak-saml-subsystem`, `mod-auth-mellon`, etc.
        /// </summary>
        [Input("providerId", required: true)]
        public Input<string> ProviderId { get; set; } = null!;

        /// <summary>
        /// The realm that the SAML client exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GetClientInstallationProviderInvokeArgs()
        {
        }
        public static new GetClientInstallationProviderInvokeArgs Empty => new GetClientInstallationProviderInvokeArgs();
    }


    [OutputType]
    public sealed class GetClientInstallationProviderResult
    {
        public readonly string ClientId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProviderId;
        public readonly string RealmId;
        /// <summary>
        /// (Computed) The returned document needed for SAML installation.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetClientInstallationProviderResult(
            string clientId,

            string id,

            string providerId,

            string realmId,

            string value)
        {
            ClientId = clientId;
            Id = id;
            ProviderId = providerId;
            RealmId = realmId;
            Value = value;
        }
    }
}
