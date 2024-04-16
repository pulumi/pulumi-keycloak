// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml
{
    /// <summary>
    /// ## # keycloak.saml.Client
    /// 
    /// Allows for creating and managing Keycloak clients that use the SAML protocol.
    /// 
    /// Clients are entities that can use Keycloak for user authentication. Typically,
    /// clients are applications that redirect users to Keycloak for authentication
    /// in order to take advantage of Keycloak's user sessions for SSO.
    /// 
    /// ### Import
    /// 
    /// Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `client_keycloak_id` is the unique ID that Keycloak
    /// assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID.
    /// 
    /// Example:
    /// </summary>
    [KeycloakResourceType("keycloak:saml/client:Client")]
    public partial class Client : global::Pulumi.CustomResource
    {
        [Output("assertionConsumerPostUrl")]
        public Output<string?> AssertionConsumerPostUrl { get; private set; } = null!;

        [Output("assertionConsumerRedirectUrl")]
        public Output<string?> AssertionConsumerRedirectUrl { get; private set; } = null!;

        [Output("authenticationFlowBindingOverrides")]
        public Output<Outputs.ClientAuthenticationFlowBindingOverrides?> AuthenticationFlowBindingOverrides { get; private set; } = null!;

        [Output("baseUrl")]
        public Output<string?> BaseUrl { get; private set; } = null!;

        [Output("canonicalizationMethod")]
        public Output<string?> CanonicalizationMethod { get; private set; } = null!;

        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        [Output("clientSignatureRequired")]
        public Output<bool?> ClientSignatureRequired { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("encryptAssertions")]
        public Output<bool?> EncryptAssertions { get; private set; } = null!;

        [Output("encryptionCertificate")]
        public Output<string> EncryptionCertificate { get; private set; } = null!;

        [Output("encryptionCertificateSha1")]
        public Output<string> EncryptionCertificateSha1 { get; private set; } = null!;

        [Output("extraConfig")]
        public Output<ImmutableDictionary<string, object>?> ExtraConfig { get; private set; } = null!;

        [Output("forceNameIdFormat")]
        public Output<bool?> ForceNameIdFormat { get; private set; } = null!;

        [Output("forcePostBinding")]
        public Output<bool?> ForcePostBinding { get; private set; } = null!;

        [Output("frontChannelLogout")]
        public Output<bool?> FrontChannelLogout { get; private set; } = null!;

        [Output("fullScopeAllowed")]
        public Output<bool?> FullScopeAllowed { get; private set; } = null!;

        [Output("idpInitiatedSsoRelayState")]
        public Output<string?> IdpInitiatedSsoRelayState { get; private set; } = null!;

        [Output("idpInitiatedSsoUrlName")]
        public Output<string?> IdpInitiatedSsoUrlName { get; private set; } = null!;

        [Output("includeAuthnStatement")]
        public Output<bool?> IncludeAuthnStatement { get; private set; } = null!;

        [Output("loginTheme")]
        public Output<string?> LoginTheme { get; private set; } = null!;

        [Output("logoutServicePostBindingUrl")]
        public Output<string?> LogoutServicePostBindingUrl { get; private set; } = null!;

        [Output("logoutServiceRedirectBindingUrl")]
        public Output<string?> LogoutServiceRedirectBindingUrl { get; private set; } = null!;

        [Output("masterSamlProcessingUrl")]
        public Output<string?> MasterSamlProcessingUrl { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nameIdFormat")]
        public Output<string> NameIdFormat { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("rootUrl")]
        public Output<string?> RootUrl { get; private set; } = null!;

        [Output("signAssertions")]
        public Output<bool?> SignAssertions { get; private set; } = null!;

        [Output("signDocuments")]
        public Output<bool?> SignDocuments { get; private set; } = null!;

        [Output("signatureAlgorithm")]
        public Output<string?> SignatureAlgorithm { get; private set; } = null!;

        [Output("signatureKeyName")]
        public Output<string?> SignatureKeyName { get; private set; } = null!;

        [Output("signingCertificate")]
        public Output<string> SigningCertificate { get; private set; } = null!;

        [Output("signingCertificateSha1")]
        public Output<string> SigningCertificateSha1 { get; private set; } = null!;

        [Output("signingPrivateKey")]
        public Output<string> SigningPrivateKey { get; private set; } = null!;

        [Output("signingPrivateKeySha1")]
        public Output<string> SigningPrivateKeySha1 { get; private set; } = null!;

        [Output("validRedirectUris")]
        public Output<ImmutableArray<string>> ValidRedirectUris { get; private set; } = null!;


        /// <summary>
        /// Create a Client resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Client(string name, ClientArgs args, CustomResourceOptions? options = null)
            : base("keycloak:saml/client:Client", name, args ?? new ClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Client(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:saml/client:Client", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Client resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Client Get(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
        {
            return new Client(name, id, state, options);
        }
    }

    public sealed class ClientArgs : global::Pulumi.ResourceArgs
    {
        [Input("assertionConsumerPostUrl")]
        public Input<string>? AssertionConsumerPostUrl { get; set; }

        [Input("assertionConsumerRedirectUrl")]
        public Input<string>? AssertionConsumerRedirectUrl { get; set; }

        [Input("authenticationFlowBindingOverrides")]
        public Input<Inputs.ClientAuthenticationFlowBindingOverridesArgs>? AuthenticationFlowBindingOverrides { get; set; }

        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        [Input("canonicalizationMethod")]
        public Input<string>? CanonicalizationMethod { get; set; }

        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSignatureRequired")]
        public Input<bool>? ClientSignatureRequired { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("encryptAssertions")]
        public Input<bool>? EncryptAssertions { get; set; }

        [Input("encryptionCertificate")]
        public Input<string>? EncryptionCertificate { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        [Input("forceNameIdFormat")]
        public Input<bool>? ForceNameIdFormat { get; set; }

        [Input("forcePostBinding")]
        public Input<bool>? ForcePostBinding { get; set; }

        [Input("frontChannelLogout")]
        public Input<bool>? FrontChannelLogout { get; set; }

        [Input("fullScopeAllowed")]
        public Input<bool>? FullScopeAllowed { get; set; }

        [Input("idpInitiatedSsoRelayState")]
        public Input<string>? IdpInitiatedSsoRelayState { get; set; }

        [Input("idpInitiatedSsoUrlName")]
        public Input<string>? IdpInitiatedSsoUrlName { get; set; }

        [Input("includeAuthnStatement")]
        public Input<bool>? IncludeAuthnStatement { get; set; }

        [Input("loginTheme")]
        public Input<string>? LoginTheme { get; set; }

        [Input("logoutServicePostBindingUrl")]
        public Input<string>? LogoutServicePostBindingUrl { get; set; }

        [Input("logoutServiceRedirectBindingUrl")]
        public Input<string>? LogoutServiceRedirectBindingUrl { get; set; }

        [Input("masterSamlProcessingUrl")]
        public Input<string>? MasterSamlProcessingUrl { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameIdFormat")]
        public Input<string>? NameIdFormat { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        [Input("signAssertions")]
        public Input<bool>? SignAssertions { get; set; }

        [Input("signDocuments")]
        public Input<bool>? SignDocuments { get; set; }

        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        [Input("signatureKeyName")]
        public Input<string>? SignatureKeyName { get; set; }

        [Input("signingCertificate")]
        public Input<string>? SigningCertificate { get; set; }

        [Input("signingPrivateKey")]
        public Input<string>? SigningPrivateKey { get; set; }

        [Input("validRedirectUris")]
        private InputList<string>? _validRedirectUris;
        public InputList<string> ValidRedirectUris
        {
            get => _validRedirectUris ?? (_validRedirectUris = new InputList<string>());
            set => _validRedirectUris = value;
        }

        public ClientArgs()
        {
        }
        public static new ClientArgs Empty => new ClientArgs();
    }

    public sealed class ClientState : global::Pulumi.ResourceArgs
    {
        [Input("assertionConsumerPostUrl")]
        public Input<string>? AssertionConsumerPostUrl { get; set; }

        [Input("assertionConsumerRedirectUrl")]
        public Input<string>? AssertionConsumerRedirectUrl { get; set; }

        [Input("authenticationFlowBindingOverrides")]
        public Input<Inputs.ClientAuthenticationFlowBindingOverridesGetArgs>? AuthenticationFlowBindingOverrides { get; set; }

        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        [Input("canonicalizationMethod")]
        public Input<string>? CanonicalizationMethod { get; set; }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSignatureRequired")]
        public Input<bool>? ClientSignatureRequired { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("encryptAssertions")]
        public Input<bool>? EncryptAssertions { get; set; }

        [Input("encryptionCertificate")]
        public Input<string>? EncryptionCertificate { get; set; }

        [Input("encryptionCertificateSha1")]
        public Input<string>? EncryptionCertificateSha1 { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        [Input("forceNameIdFormat")]
        public Input<bool>? ForceNameIdFormat { get; set; }

        [Input("forcePostBinding")]
        public Input<bool>? ForcePostBinding { get; set; }

        [Input("frontChannelLogout")]
        public Input<bool>? FrontChannelLogout { get; set; }

        [Input("fullScopeAllowed")]
        public Input<bool>? FullScopeAllowed { get; set; }

        [Input("idpInitiatedSsoRelayState")]
        public Input<string>? IdpInitiatedSsoRelayState { get; set; }

        [Input("idpInitiatedSsoUrlName")]
        public Input<string>? IdpInitiatedSsoUrlName { get; set; }

        [Input("includeAuthnStatement")]
        public Input<bool>? IncludeAuthnStatement { get; set; }

        [Input("loginTheme")]
        public Input<string>? LoginTheme { get; set; }

        [Input("logoutServicePostBindingUrl")]
        public Input<string>? LogoutServicePostBindingUrl { get; set; }

        [Input("logoutServiceRedirectBindingUrl")]
        public Input<string>? LogoutServiceRedirectBindingUrl { get; set; }

        [Input("masterSamlProcessingUrl")]
        public Input<string>? MasterSamlProcessingUrl { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameIdFormat")]
        public Input<string>? NameIdFormat { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        [Input("signAssertions")]
        public Input<bool>? SignAssertions { get; set; }

        [Input("signDocuments")]
        public Input<bool>? SignDocuments { get; set; }

        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        [Input("signatureKeyName")]
        public Input<string>? SignatureKeyName { get; set; }

        [Input("signingCertificate")]
        public Input<string>? SigningCertificate { get; set; }

        [Input("signingCertificateSha1")]
        public Input<string>? SigningCertificateSha1 { get; set; }

        [Input("signingPrivateKey")]
        public Input<string>? SigningPrivateKey { get; set; }

        [Input("signingPrivateKeySha1")]
        public Input<string>? SigningPrivateKeySha1 { get; set; }

        [Input("validRedirectUris")]
        private InputList<string>? _validRedirectUris;
        public InputList<string> ValidRedirectUris
        {
            get => _validRedirectUris ?? (_validRedirectUris = new InputList<string>());
            set => _validRedirectUris = value;
        }

        public ClientState()
        {
        }
        public static new ClientState Empty => new ClientState();
    }
}
