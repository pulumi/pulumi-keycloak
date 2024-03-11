// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.saml.ClientArgs;
import com.pulumi.keycloak.saml.inputs.ClientState;
import com.pulumi.keycloak.saml.outputs.ClientAuthenticationFlowBindingOverrides;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.saml.Client
 * 
 * Allows for creating and managing Keycloak clients that use the SAML protocol.
 * 
 * Clients are entities that can use Keycloak for user authentication. Typically,
 * clients are applications that redirect users to Keycloak for authentication
 * in order to take advantage of Keycloak&#39;s user sessions for SSO.
 * 
 * ### Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.saml.Client;
 * import com.pulumi.keycloak.saml.ClientArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var realm = new Realm(&#34;realm&#34;, RealmArgs.builder()        
 *             .enabled(true)
 *             .realm(&#34;my-realm&#34;)
 *             .build());
 * 
 *         var samlClient = new Client(&#34;samlClient&#34;, ClientArgs.builder()        
 *             .clientId(&#34;test-saml-client&#34;)
 *             .includeAuthnStatement(true)
 *             .realmId(realm.id())
 *             .signAssertions(true)
 *             .signDocuments(false)
 *             .signingCertificate(Files.readString(Paths.get(&#34;saml-cert.pem&#34;)))
 *             .signingPrivateKey(Files.readString(Paths.get(&#34;saml-key.pem&#34;)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realm_id` - (Required) The realm this client is attached to.
 * - `client_id` - (Required) The unique ID of this client, referenced in the URI during authentication and in issued tokens.
 * - `name` - (Optional) The display name of this client in the GUI.
 * - `enabled` - (Optional) When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
 * - `description` - (Optional) The description of this client in the GUI.
 * - `include_authn_statement` - (Optional) When `true`, an `AuthnStatement` will be included in the SAML response.
 * - `sign_documents` - (Optional) When `true`, the SAML document will be signed by Keycloak using the realm&#39;s private key.
 * - `sign_assertions` - (Optional) When `true`, the SAML assertions will be signed by Keycloak using the realm&#39;s private key, and embedded within the SAML XML Auth response.
 * - `client_signature_required` - (Optional) When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signing_certificate` and `signing_private_key`.
 * - `force_post_binding` - (Optional) When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding.
 * - `front_channel_logout` - (Optional) When `true`, this client will require a browser redirect in order to perform a logout.
 * - `name_id_format` - (Optional) Sets the Name ID format for the subject.
 * - `root_url` - (Optional) When specified, this value is prepended to all relative URLs.
 * - `valid_redirect_uris` - (Optional) When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
 * - `base_url` - (Optional) When specified, this URL will be used whenever Keycloak needs to link to this client.
 * - `master_saml_processing_url` - (Optional) When specified, this URL will be used for all SAML requests.
 * - `signing_certificate` - (Optional) If documents or assertions from the client are signed, this certificate will be used to verify the signature.
 * - `signing_private_key` - (Optional) If documents or assertions from the client are signed, this private key will be used to verify the signature.
 * - `idp_initiated_sso_url_name` - (Optional) URL fragment name to reference client when you want to do IDP Initiated SSO.
 * - `idp_initiated_sso_relay_state` - (Optional) Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
 * - `assertion_consumer_post_url` - (Optional) SAML POST Binding URL for the client&#39;s assertion consumer service (login responses).
 * - `assertion_consumer_redirect_url` - (Optional) SAML Redirect Binding URL for the client&#39;s assertion consumer service (login responses).
 * - `logout_service_post_binding_url` - (Optional) SAML POST Binding URL for the client&#39;s single logout service.
 * - `logout_service_redirect_binding_url` - (Optional) SAML Redirect Binding URL for the client&#39;s single logout service.
 * - `full_scope_allowed` - (Optional) - Allow to include all roles mappings in the access token
 * 
 * ### Import
 * 
 * Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `client_keycloak_id` is the unique ID that Keycloak
 * assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID.
 * 
 * Example:
 * 
 */
@ResourceType(type="keycloak:saml/client:Client")
public class Client extends com.pulumi.resources.CustomResource {
    @Export(name="assertionConsumerPostUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> assertionConsumerPostUrl;

    public Output<Optional<String>> assertionConsumerPostUrl() {
        return Codegen.optional(this.assertionConsumerPostUrl);
    }
    @Export(name="assertionConsumerRedirectUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> assertionConsumerRedirectUrl;

    public Output<Optional<String>> assertionConsumerRedirectUrl() {
        return Codegen.optional(this.assertionConsumerRedirectUrl);
    }
    @Export(name="authenticationFlowBindingOverrides", refs={ClientAuthenticationFlowBindingOverrides.class}, tree="[0]")
    private Output</* @Nullable */ ClientAuthenticationFlowBindingOverrides> authenticationFlowBindingOverrides;

    public Output<Optional<ClientAuthenticationFlowBindingOverrides>> authenticationFlowBindingOverrides() {
        return Codegen.optional(this.authenticationFlowBindingOverrides);
    }
    @Export(name="baseUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> baseUrl;

    public Output<Optional<String>> baseUrl() {
        return Codegen.optional(this.baseUrl);
    }
    @Export(name="canonicalizationMethod", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> canonicalizationMethod;

    public Output<Optional<String>> canonicalizationMethod() {
        return Codegen.optional(this.canonicalizationMethod);
    }
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    public Output<String> clientId() {
        return this.clientId;
    }
    @Export(name="clientSignatureRequired", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> clientSignatureRequired;

    public Output<Optional<Boolean>> clientSignatureRequired() {
        return Codegen.optional(this.clientSignatureRequired);
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    @Export(name="encryptAssertions", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> encryptAssertions;

    public Output<Optional<Boolean>> encryptAssertions() {
        return Codegen.optional(this.encryptAssertions);
    }
    @Export(name="encryptionCertificate", refs={String.class}, tree="[0]")
    private Output<String> encryptionCertificate;

    public Output<String> encryptionCertificate() {
        return this.encryptionCertificate;
    }
    @Export(name="encryptionCertificateSha1", refs={String.class}, tree="[0]")
    private Output<String> encryptionCertificateSha1;

    public Output<String> encryptionCertificateSha1() {
        return this.encryptionCertificateSha1;
    }
    @Export(name="extraConfig", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> extraConfig;

    public Output<Optional<Map<String,Object>>> extraConfig() {
        return Codegen.optional(this.extraConfig);
    }
    @Export(name="forceNameIdFormat", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceNameIdFormat;

    public Output<Optional<Boolean>> forceNameIdFormat() {
        return Codegen.optional(this.forceNameIdFormat);
    }
    @Export(name="forcePostBinding", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forcePostBinding;

    public Output<Optional<Boolean>> forcePostBinding() {
        return Codegen.optional(this.forcePostBinding);
    }
    @Export(name="frontChannelLogout", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> frontChannelLogout;

    public Output<Optional<Boolean>> frontChannelLogout() {
        return Codegen.optional(this.frontChannelLogout);
    }
    @Export(name="fullScopeAllowed", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> fullScopeAllowed;

    public Output<Optional<Boolean>> fullScopeAllowed() {
        return Codegen.optional(this.fullScopeAllowed);
    }
    @Export(name="idpInitiatedSsoRelayState", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> idpInitiatedSsoRelayState;

    public Output<Optional<String>> idpInitiatedSsoRelayState() {
        return Codegen.optional(this.idpInitiatedSsoRelayState);
    }
    @Export(name="idpInitiatedSsoUrlName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> idpInitiatedSsoUrlName;

    public Output<Optional<String>> idpInitiatedSsoUrlName() {
        return Codegen.optional(this.idpInitiatedSsoUrlName);
    }
    @Export(name="includeAuthnStatement", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> includeAuthnStatement;

    public Output<Optional<Boolean>> includeAuthnStatement() {
        return Codegen.optional(this.includeAuthnStatement);
    }
    @Export(name="loginTheme", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loginTheme;

    public Output<Optional<String>> loginTheme() {
        return Codegen.optional(this.loginTheme);
    }
    @Export(name="logoutServicePostBindingUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logoutServicePostBindingUrl;

    public Output<Optional<String>> logoutServicePostBindingUrl() {
        return Codegen.optional(this.logoutServicePostBindingUrl);
    }
    @Export(name="logoutServiceRedirectBindingUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logoutServiceRedirectBindingUrl;

    public Output<Optional<String>> logoutServiceRedirectBindingUrl() {
        return Codegen.optional(this.logoutServiceRedirectBindingUrl);
    }
    @Export(name="masterSamlProcessingUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> masterSamlProcessingUrl;

    public Output<Optional<String>> masterSamlProcessingUrl() {
        return Codegen.optional(this.masterSamlProcessingUrl);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="nameIdFormat", refs={String.class}, tree="[0]")
    private Output<String> nameIdFormat;

    public Output<String> nameIdFormat() {
        return this.nameIdFormat;
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="rootUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> rootUrl;

    public Output<Optional<String>> rootUrl() {
        return Codegen.optional(this.rootUrl);
    }
    @Export(name="signAssertions", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> signAssertions;

    public Output<Optional<Boolean>> signAssertions() {
        return Codegen.optional(this.signAssertions);
    }
    @Export(name="signDocuments", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> signDocuments;

    public Output<Optional<Boolean>> signDocuments() {
        return Codegen.optional(this.signDocuments);
    }
    @Export(name="signatureAlgorithm", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> signatureAlgorithm;

    public Output<Optional<String>> signatureAlgorithm() {
        return Codegen.optional(this.signatureAlgorithm);
    }
    @Export(name="signatureKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> signatureKeyName;

    public Output<Optional<String>> signatureKeyName() {
        return Codegen.optional(this.signatureKeyName);
    }
    @Export(name="signingCertificate", refs={String.class}, tree="[0]")
    private Output<String> signingCertificate;

    public Output<String> signingCertificate() {
        return this.signingCertificate;
    }
    @Export(name="signingCertificateSha1", refs={String.class}, tree="[0]")
    private Output<String> signingCertificateSha1;

    public Output<String> signingCertificateSha1() {
        return this.signingCertificateSha1;
    }
    @Export(name="signingPrivateKey", refs={String.class}, tree="[0]")
    private Output<String> signingPrivateKey;

    public Output<String> signingPrivateKey() {
        return this.signingPrivateKey;
    }
    @Export(name="signingPrivateKeySha1", refs={String.class}, tree="[0]")
    private Output<String> signingPrivateKeySha1;

    public Output<String> signingPrivateKeySha1() {
        return this.signingPrivateKeySha1;
    }
    @Export(name="validRedirectUris", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> validRedirectUris;

    public Output<Optional<List<String>>> validRedirectUris() {
        return Codegen.optional(this.validRedirectUris);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Client(String name) {
        this(name, ClientArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Client(String name, ClientArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Client(String name, ClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/client:Client", name, args == null ? ClientArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Client(String name, Output<String> id, @Nullable ClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/client:Client", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Client get(String name, Output<String> id, @Nullable ClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Client(name, id, state, options);
    }
}
