// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.saml.IdentityProviderArgs;
import com.pulumi.keycloak.saml.inputs.IdentityProviderState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing SAML Identity Providers within Keycloak.
 * 
 * SAML (Security Assertion Markup Language) identity providers allows users to authenticate through a third-party system using the SAML protocol.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.saml.IdentityProvider;
 * import com.pulumi.keycloak.saml.IdentityProviderArgs;
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
 *         var realm = new Realm("realm", RealmArgs.builder()
 *             .realm("my-realm")
 *             .enabled(true)
 *             .build());
 * 
 *         var realmSamlIdentityProvider = new IdentityProvider("realmSamlIdentityProvider", IdentityProviderArgs.builder()
 *             .realm(realm.id())
 *             .alias("my-saml-idp")
 *             .entityId("https://domain.com/entity_id")
 *             .singleSignOnServiceUrl("https://domain.com/adfs/ls/")
 *             .singleLogoutServiceUrl("https://domain.com/adfs/ls/?wa=wsignout1.0")
 *             .backchannelSupported(true)
 *             .postBindingResponse(true)
 *             .postBindingLogout(true)
 *             .postBindingAuthnRequest(true)
 *             .storeToken(false)
 *             .trustEmail(true)
 *             .forceAuthn(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Identity providers can be imported using the format `{{realm_id}}/{{idp_alias}}`, where `idp_alias` is the identity provider alias.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:saml/identityProvider:IdentityProvider realm_saml_identity_provider my-realm/my-saml-idp
 * ```
 * 
 */
@ResourceType(type="keycloak:saml/identityProvider:IdentityProvider")
public class IdentityProvider extends com.pulumi.resources.CustomResource {
    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     * 
     */
    @Export(name="addReadTokenRoleOnCreate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addReadTokenRoleOnCreate;

    /**
     * @return When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> addReadTokenRoleOnCreate() {
        return Codegen.optional(this.addReadTokenRoleOnCreate);
    }
    /**
     * The unique name of identity provider.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return The unique name of identity provider.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * Authenticate users by default. Defaults to `false`.
     * 
     */
    @Export(name="authenticateByDefault", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> authenticateByDefault;

    /**
     * @return Authenticate users by default. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> authenticateByDefault() {
        return Codegen.optional(this.authenticateByDefault);
    }
    /**
     * Ordered list of requested AuthnContext ClassRefs.
     * 
     */
    @Export(name="authnContextClassRefs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authnContextClassRefs;

    /**
     * @return Ordered list of requested AuthnContext ClassRefs.
     * 
     */
    public Output<Optional<List<String>>> authnContextClassRefs() {
        return Codegen.optional(this.authnContextClassRefs);
    }
    /**
     * Specifies the comparison method used to evaluate the requested context classes or statements.
     * 
     */
    @Export(name="authnContextComparisonType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authnContextComparisonType;

    /**
     * @return Specifies the comparison method used to evaluate the requested context classes or statements.
     * 
     */
    public Output<Optional<String>> authnContextComparisonType() {
        return Codegen.optional(this.authnContextComparisonType);
    }
    /**
     * Ordered list of requested AuthnContext DeclRefs.
     * 
     */
    @Export(name="authnContextDeclRefs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authnContextDeclRefs;

    /**
     * @return Ordered list of requested AuthnContext DeclRefs.
     * 
     */
    public Output<Optional<List<String>>> authnContextDeclRefs() {
        return Codegen.optional(this.authnContextDeclRefs);
    }
    /**
     * Does the external IDP support backchannel logout?. Defaults to `false`.
     * 
     */
    @Export(name="backchannelSupported", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> backchannelSupported;

    /**
     * @return Does the external IDP support backchannel logout?. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> backchannelSupported() {
        return Codegen.optional(this.backchannelSupported);
    }
    /**
     * The display name for the realm that is shown when logging in to the admin console.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The display name for the realm that is shown when logging in to the admin console.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * When `false`, users and clients will not be able to access this realm. Defaults to `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When `false`, users and clients will not be able to access this realm. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The Entity ID that will be used to uniquely identify this SAML Service Provider.
     * 
     */
    @Export(name="entityId", refs={String.class}, tree="[0]")
    private Output<String> entityId;

    /**
     * @return The Entity ID that will be used to uniquely identify this SAML Service Provider.
     * 
     */
    public Output<String> entityId() {
        return this.entityId;
    }
    @Export(name="extraConfig", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> extraConfig;

    public Output<Optional<Map<String,String>>> extraConfig() {
        return Codegen.optional(this.extraConfig);
    }
    /**
     * Alias of authentication flow, which is triggered after first login with this identity provider. Term &#39;First Login&#39; means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
     * 
     */
    @Export(name="firstBrokerLoginFlowAlias", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> firstBrokerLoginFlowAlias;

    /**
     * @return Alias of authentication flow, which is triggered after first login with this identity provider. Term &#39;First Login&#39; means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
     * 
     */
    public Output<Optional<String>> firstBrokerLoginFlowAlias() {
        return Codegen.optional(this.firstBrokerLoginFlowAlias);
    }
    /**
     * Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
     * 
     */
    @Export(name="forceAuthn", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceAuthn;

    /**
     * @return Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
     * 
     */
    public Output<Optional<Boolean>> forceAuthn() {
        return Codegen.optional(this.forceAuthn);
    }
    /**
     * A number defining the order of this identity provider in the GUI.
     * 
     */
    @Export(name="guiOrder", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> guiOrder;

    /**
     * @return A number defining the order of this identity provider in the GUI.
     * 
     */
    public Output<Optional<String>> guiOrder() {
        return Codegen.optional(this.guiOrder);
    }
    /**
     * If hidden, then login with this provider is possible only if requested explicitly, e.g. using the &#39;kc_idp_hint&#39; parameter.
     * 
     */
    @Export(name="hideOnLoginPage", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hideOnLoginPage;

    /**
     * @return If hidden, then login with this provider is possible only if requested explicitly, e.g. using the &#39;kc_idp_hint&#39; parameter.
     * 
     */
    public Output<Optional<Boolean>> hideOnLoginPage() {
        return Codegen.optional(this.hideOnLoginPage);
    }
    /**
     * Internal Identity Provider Id
     * 
     */
    @Export(name="internalId", refs={String.class}, tree="[0]")
    private Output<String> internalId;

    /**
     * @return Internal Identity Provider Id
     * 
     */
    public Output<String> internalId() {
        return this.internalId;
    }
    /**
     * When `true`, users cannot log in using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     * 
     */
    @Export(name="linkOnly", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> linkOnly;

    /**
     * @return When `true`, users cannot log in using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> linkOnly() {
        return Codegen.optional(this.linkOnly);
    }
    /**
     * Login Hint.
     * 
     */
    @Export(name="loginHint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loginHint;

    /**
     * @return Login Hint.
     * 
     */
    public Output<Optional<String>> loginHint() {
        return Codegen.optional(this.loginHint);
    }
    /**
     * Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
     * 
     */
    @Export(name="nameIdPolicyFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nameIdPolicyFormat;

    /**
     * @return Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
     * 
     */
    public Output<Optional<String>> nameIdPolicyFormat() {
        return Codegen.optional(this.nameIdPolicyFormat);
    }
    /**
     * The organization domain to associate this identity provider with. It is used to map users to an organization based on their email domain and to authenticate them accordingly in the scope of the organization.
     * 
     */
    @Export(name="orgDomain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> orgDomain;

    /**
     * @return The organization domain to associate this identity provider with. It is used to map users to an organization based on their email domain and to authenticate them accordingly in the scope of the organization.
     * 
     */
    public Output<Optional<String>> orgDomain() {
        return Codegen.optional(this.orgDomain);
    }
    /**
     * Indicates whether to automatically redirect users to this identity provider when email domain matches domain.
     * 
     */
    @Export(name="orgRedirectModeEmailMatches", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> orgRedirectModeEmailMatches;

    /**
     * @return Indicates whether to automatically redirect users to this identity provider when email domain matches domain.
     * 
     */
    public Output<Optional<Boolean>> orgRedirectModeEmailMatches() {
        return Codegen.optional(this.orgRedirectModeEmailMatches);
    }
    /**
     * The ID of the organization to link this identity provider to.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> organizationId;

    /**
     * @return The ID of the organization to link this identity provider to.
     * 
     */
    public Output<Optional<String>> organizationId() {
        return Codegen.optional(this.organizationId);
    }
    /**
     * Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
     * 
     */
    @Export(name="postBindingAuthnRequest", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> postBindingAuthnRequest;

    /**
     * @return Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
     * 
     */
    public Output<Optional<Boolean>> postBindingAuthnRequest() {
        return Codegen.optional(this.postBindingAuthnRequest);
    }
    /**
     * Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
     * 
     */
    @Export(name="postBindingLogout", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> postBindingLogout;

    /**
     * @return Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
     * 
     */
    public Output<Optional<Boolean>> postBindingLogout() {
        return Codegen.optional(this.postBindingLogout);
    }
    /**
     * Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
     * 
     */
    @Export(name="postBindingResponse", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> postBindingResponse;

    /**
     * @return Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
     * 
     */
    public Output<Optional<Boolean>> postBindingResponse() {
        return Codegen.optional(this.postBindingResponse);
    }
    /**
     * Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don&#39;t want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
     * 
     */
    @Export(name="postBrokerLoginFlowAlias", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> postBrokerLoginFlowAlias;

    /**
     * @return Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don&#39;t want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
     * 
     */
    public Output<Optional<String>> postBrokerLoginFlowAlias() {
        return Codegen.optional(this.postBrokerLoginFlowAlias);
    }
    /**
     * The principal attribute.
     * 
     */
    @Export(name="principalAttribute", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> principalAttribute;

    /**
     * @return The principal attribute.
     * 
     */
    public Output<Optional<String>> principalAttribute() {
        return Codegen.optional(this.principalAttribute);
    }
    /**
     * The principal type. Can be one of `SUBJECT`, `ATTRIBUTE` or `FRIENDLY_ATTRIBUTE`.
     * 
     */
    @Export(name="principalType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> principalType;

    /**
     * @return The principal type. Can be one of `SUBJECT`, `ATTRIBUTE` or `FRIENDLY_ATTRIBUTE`.
     * 
     */
    public Output<Optional<String>> principalType() {
        return Codegen.optional(this.principalType);
    }
    /**
     * The ID of the identity provider to use. Defaults to `saml`, which should be used unless you have extended Keycloak and provided your own implementation.
     * 
     */
    @Export(name="providerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> providerId;

    /**
     * @return The ID of the identity provider to use. Defaults to `saml`, which should be used unless you have extended Keycloak and provided your own implementation.
     * 
     */
    public Output<Optional<String>> providerId() {
        return Codegen.optional(this.providerId);
    }
    /**
     * The name of the realm. This is unique across Keycloak.
     * 
     */
    @Export(name="realm", refs={String.class}, tree="[0]")
    private Output<String> realm;

    /**
     * @return The name of the realm. This is unique across Keycloak.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }
    /**
     * Signing Algorithm. Defaults to empty.
     * 
     */
    @Export(name="signatureAlgorithm", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> signatureAlgorithm;

    /**
     * @return Signing Algorithm. Defaults to empty.
     * 
     */
    public Output<Optional<String>> signatureAlgorithm() {
        return Codegen.optional(this.signatureAlgorithm);
    }
    /**
     * Signing Certificate.
     * 
     */
    @Export(name="signingCertificate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> signingCertificate;

    /**
     * @return Signing Certificate.
     * 
     */
    public Output<Optional<String>> signingCertificate() {
        return Codegen.optional(this.signingCertificate);
    }
    /**
     * The Url that must be used to send logout requests.
     * 
     */
    @Export(name="singleLogoutServiceUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> singleLogoutServiceUrl;

    /**
     * @return The Url that must be used to send logout requests.
     * 
     */
    public Output<Optional<String>> singleLogoutServiceUrl() {
        return Codegen.optional(this.singleLogoutServiceUrl);
    }
    /**
     * The Url that must be used to send authentication requests (SAML AuthnRequest).
     * 
     */
    @Export(name="singleSignOnServiceUrl", refs={String.class}, tree="[0]")
    private Output<String> singleSignOnServiceUrl;

    /**
     * @return The Url that must be used to send authentication requests (SAML AuthnRequest).
     * 
     */
    public Output<String> singleSignOnServiceUrl() {
        return this.singleSignOnServiceUrl;
    }
    /**
     * When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     * 
     */
    @Export(name="storeToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> storeToken;

    /**
     * @return When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> storeToken() {
        return Codegen.optional(this.storeToken);
    }
    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be one of `IMPORT`, `FORCE`, or `LEGACY`.
     * 
     */
    @Export(name="syncMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> syncMode;

    /**
     * @return The default sync mode to use for all mappers attached to this identity provider. Can be one of `IMPORT`, `FORCE`, or `LEGACY`.
     * 
     */
    public Output<Optional<String>> syncMode() {
        return Codegen.optional(this.syncMode);
    }
    /**
     * When `true`, email addresses for users in this provider will automatically be verified regardless of the realm&#39;s email verification policy. Defaults to `false`.
     * 
     */
    @Export(name="trustEmail", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> trustEmail;

    /**
     * @return When `true`, email addresses for users in this provider will automatically be verified regardless of the realm&#39;s email verification policy. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> trustEmail() {
        return Codegen.optional(this.trustEmail);
    }
    /**
     * Enable/disable signature validation of SAML responses.
     * 
     */
    @Export(name="validateSignature", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> validateSignature;

    /**
     * @return Enable/disable signature validation of SAML responses.
     * 
     */
    public Output<Optional<Boolean>> validateSignature() {
        return Codegen.optional(this.validateSignature);
    }
    /**
     * Indicates whether this service provider expects an encrypted Assertion.
     * 
     */
    @Export(name="wantAssertionsEncrypted", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> wantAssertionsEncrypted;

    /**
     * @return Indicates whether this service provider expects an encrypted Assertion.
     * 
     */
    public Output<Optional<Boolean>> wantAssertionsEncrypted() {
        return Codegen.optional(this.wantAssertionsEncrypted);
    }
    /**
     * Indicates whether this service provider expects a signed Assertion.
     * 
     */
    @Export(name="wantAssertionsSigned", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> wantAssertionsSigned;

    /**
     * @return Indicates whether this service provider expects a signed Assertion.
     * 
     */
    public Output<Optional<Boolean>> wantAssertionsSigned() {
        return Codegen.optional(this.wantAssertionsSigned);
    }
    /**
     * The SAML signature key name. Can be one of `NONE`, `KEY_ID`, or `CERT_SUBJECT`.
     * 
     */
    @Export(name="xmlSignKeyInfoKeyNameTransformer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> xmlSignKeyInfoKeyNameTransformer;

    /**
     * @return The SAML signature key name. Can be one of `NONE`, `KEY_ID`, or `CERT_SUBJECT`.
     * 
     */
    public Output<Optional<String>> xmlSignKeyInfoKeyNameTransformer() {
        return Codegen.optional(this.xmlSignKeyInfoKeyNameTransformer);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IdentityProvider(java.lang.String name) {
        this(name, IdentityProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IdentityProvider(java.lang.String name, IdentityProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IdentityProvider(java.lang.String name, IdentityProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/identityProvider:IdentityProvider", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IdentityProvider(java.lang.String name, Output<java.lang.String> id, @Nullable IdentityProviderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/identityProvider:IdentityProvider", name, state, makeResourceOptions(options, id), false);
    }

    private static IdentityProviderArgs makeArgs(IdentityProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IdentityProviderArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static IdentityProvider get(java.lang.String name, Output<java.lang.String> id, @Nullable IdentityProviderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IdentityProvider(name, id, state, options);
    }
}
