// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing OIDC Identity Providers within Keycloak.
 *
 * OIDC (OpenID Connect) identity providers allows users to authenticate through a third party system using the OIDC standard.
 *
 * > **NOTICE:** This resource now supports write-only arguments
 * for client secret via the new arguments `clientSecretWo` and `clientSecretWoVersion`. Using write-only arguments
 * prevents sensitive values from being stored in plan and state files. You cannot use `clientSecretWo` and
 * `clientSecretWoVersion` alongside `clientSecret` as this will result in a validation error due to conflicts.
 * > 
 * > For backward compatibility, the behavior of the original `clientSecret` argument remains unchanged.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const realmIdentityProvider = new keycloak.oidc.IdentityProvider("realm_identity_provider", {
 *     realm: realm.id,
 *     alias: "my-idp",
 *     authorizationUrl: "https://authorizationurl.com",
 *     clientId: "clientID",
 *     clientSecret: "clientSecret",
 *     tokenUrl: "https://tokenurl.com",
 *     extraConfig: {
 *         clientAuthMethod: "client_secret_post",
 *     },
 * });
 * ```
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
 * $ pulumi import keycloak:oidc/identityProvider:IdentityProvider realm_identity_provider my-realm/my-idp
 * ```
 */
export class IdentityProvider extends pulumi.CustomResource {
    /**
     * Get an existing IdentityProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityProviderState, opts?: pulumi.CustomResourceOptions): IdentityProvider {
        return new IdentityProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:oidc/identityProvider:IdentityProvider';

    /**
     * Returns true if the given object is an instance of IdentityProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityProvider.__pulumiType;
    }

    /**
     * When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
     */
    public readonly acceptsPromptNoneForwardFromClient!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     */
    public readonly addReadTokenRoleOnCreate!: pulumi.Output<boolean | undefined>;
    /**
     * The alias uniquely identifies an identity provider, and it is also used to build the redirect uri.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * Enable/disable authenticate users by default.
     */
    public readonly authenticateByDefault!: pulumi.Output<boolean | undefined>;
    /**
     * The Authorization Url.
     */
    public readonly authorizationUrl!: pulumi.Output<string>;
    /**
     * Does the external IDP support backchannel logout? Defaults to `true`.
     */
    public readonly backchannelSupported!: pulumi.Output<boolean | undefined>;
    /**
     * The client or client identifier registered within the identity provider.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format. Required without `clientSecretWo` and `clientSecretWoVersion`.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Version of the Client secret write-only argument
     */
    public readonly clientSecretWoVersion!: pulumi.Output<number | undefined>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
     */
    public readonly defaultScopes!: pulumi.Output<string | undefined>;
    /**
     * When `true`, disables the check for the `typ` claim of tokens received from the identity provider. Defaults to `false`.
     */
    public readonly disableTypeClaimCheck!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     */
    public readonly disableUserInfo!: pulumi.Output<boolean | undefined>;
    /**
     * Display name for the identity provider in the GUI.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly extraConfig!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     */
    public readonly firstBrokerLoginFlowAlias!: pulumi.Output<string | undefined>;
    /**
     * A number defining the order of this identity provider in the GUI.
     */
    public readonly guiOrder!: pulumi.Output<string | undefined>;
    /**
     * When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
     */
    public readonly hideOnLoginPage!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
     */
    public /*out*/ readonly internalId!: pulumi.Output<string>;
    /**
     * The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
     */
    public readonly issuer!: pulumi.Output<string | undefined>;
    /**
     * JSON Web Key Set URL.
     */
    public readonly jwksUrl!: pulumi.Output<string | undefined>;
    /**
     * When `true`, users cannot sign-in using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     */
    public readonly linkOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Pass login hint to identity provider.
     */
    public readonly loginHint!: pulumi.Output<string | undefined>;
    /**
     * The Logout URL is the end session endpoint to use to sign-out the user from external identity provider.
     */
    public readonly logoutUrl!: pulumi.Output<string | undefined>;
    /**
     * The organization domain to associate this identity provider with. it is used to map users to an organization based on their email domain and to authenticate them accordingly in the scope of the organization.
     */
    public readonly orgDomain!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether to automatically redirect user to this identity provider when email domain matches domain.
     */
    public readonly orgRedirectModeEmailMatches!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the organization to link this identity provider to.
     */
    public readonly organizationId!: pulumi.Output<string | undefined>;
    /**
     * The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     */
    public readonly postBrokerLoginFlowAlias!: pulumi.Output<string | undefined>;
    /**
     * The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
     */
    public readonly providerId!: pulumi.Output<string | undefined>;
    /**
     * The name of the realm. This is unique across Keycloak.
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     */
    public readonly storeToken!: pulumi.Output<boolean | undefined>;
    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     */
    public readonly syncMode!: pulumi.Output<string | undefined>;
    /**
     * The Token URL.
     */
    public readonly tokenUrl!: pulumi.Output<string>;
    /**
     * When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
     */
    public readonly trustEmail!: pulumi.Output<boolean | undefined>;
    /**
     * Pass current locale to identity provider. Defaults to `false`.
     */
    public readonly uiLocales!: pulumi.Output<boolean | undefined>;
    /**
     * User Info URL.
     */
    public readonly userInfoUrl!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable signature validation of external IDP signatures. Defaults to `false`.
     */
    public readonly validateSignature!: pulumi.Output<boolean | undefined>;

    /**
     * Create a IdentityProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityProviderArgs | IdentityProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdentityProviderState | undefined;
            resourceInputs["acceptsPromptNoneForwardFromClient"] = state ? state.acceptsPromptNoneForwardFromClient : undefined;
            resourceInputs["addReadTokenRoleOnCreate"] = state ? state.addReadTokenRoleOnCreate : undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["authenticateByDefault"] = state ? state.authenticateByDefault : undefined;
            resourceInputs["authorizationUrl"] = state ? state.authorizationUrl : undefined;
            resourceInputs["backchannelSupported"] = state ? state.backchannelSupported : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["clientSecretWoVersion"] = state ? state.clientSecretWoVersion : undefined;
            resourceInputs["defaultScopes"] = state ? state.defaultScopes : undefined;
            resourceInputs["disableTypeClaimCheck"] = state ? state.disableTypeClaimCheck : undefined;
            resourceInputs["disableUserInfo"] = state ? state.disableUserInfo : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["extraConfig"] = state ? state.extraConfig : undefined;
            resourceInputs["firstBrokerLoginFlowAlias"] = state ? state.firstBrokerLoginFlowAlias : undefined;
            resourceInputs["guiOrder"] = state ? state.guiOrder : undefined;
            resourceInputs["hideOnLoginPage"] = state ? state.hideOnLoginPage : undefined;
            resourceInputs["internalId"] = state ? state.internalId : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["jwksUrl"] = state ? state.jwksUrl : undefined;
            resourceInputs["linkOnly"] = state ? state.linkOnly : undefined;
            resourceInputs["loginHint"] = state ? state.loginHint : undefined;
            resourceInputs["logoutUrl"] = state ? state.logoutUrl : undefined;
            resourceInputs["orgDomain"] = state ? state.orgDomain : undefined;
            resourceInputs["orgRedirectModeEmailMatches"] = state ? state.orgRedirectModeEmailMatches : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["postBrokerLoginFlowAlias"] = state ? state.postBrokerLoginFlowAlias : undefined;
            resourceInputs["providerId"] = state ? state.providerId : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["storeToken"] = state ? state.storeToken : undefined;
            resourceInputs["syncMode"] = state ? state.syncMode : undefined;
            resourceInputs["tokenUrl"] = state ? state.tokenUrl : undefined;
            resourceInputs["trustEmail"] = state ? state.trustEmail : undefined;
            resourceInputs["uiLocales"] = state ? state.uiLocales : undefined;
            resourceInputs["userInfoUrl"] = state ? state.userInfoUrl : undefined;
            resourceInputs["validateSignature"] = state ? state.validateSignature : undefined;
        } else {
            const args = argsOrState as IdentityProviderArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
            }
            if ((!args || args.authorizationUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationUrl'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.realm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realm'");
            }
            if ((!args || args.tokenUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tokenUrl'");
            }
            resourceInputs["acceptsPromptNoneForwardFromClient"] = args ? args.acceptsPromptNoneForwardFromClient : undefined;
            resourceInputs["addReadTokenRoleOnCreate"] = args ? args.addReadTokenRoleOnCreate : undefined;
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["authenticateByDefault"] = args ? args.authenticateByDefault : undefined;
            resourceInputs["authorizationUrl"] = args ? args.authorizationUrl : undefined;
            resourceInputs["backchannelSupported"] = args ? args.backchannelSupported : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["clientSecretWoVersion"] = args ? args.clientSecretWoVersion : undefined;
            resourceInputs["defaultScopes"] = args ? args.defaultScopes : undefined;
            resourceInputs["disableTypeClaimCheck"] = args ? args.disableTypeClaimCheck : undefined;
            resourceInputs["disableUserInfo"] = args ? args.disableUserInfo : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["extraConfig"] = args ? args.extraConfig : undefined;
            resourceInputs["firstBrokerLoginFlowAlias"] = args ? args.firstBrokerLoginFlowAlias : undefined;
            resourceInputs["guiOrder"] = args ? args.guiOrder : undefined;
            resourceInputs["hideOnLoginPage"] = args ? args.hideOnLoginPage : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["jwksUrl"] = args ? args.jwksUrl : undefined;
            resourceInputs["linkOnly"] = args ? args.linkOnly : undefined;
            resourceInputs["loginHint"] = args ? args.loginHint : undefined;
            resourceInputs["logoutUrl"] = args ? args.logoutUrl : undefined;
            resourceInputs["orgDomain"] = args ? args.orgDomain : undefined;
            resourceInputs["orgRedirectModeEmailMatches"] = args ? args.orgRedirectModeEmailMatches : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["postBrokerLoginFlowAlias"] = args ? args.postBrokerLoginFlowAlias : undefined;
            resourceInputs["providerId"] = args ? args.providerId : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["storeToken"] = args ? args.storeToken : undefined;
            resourceInputs["syncMode"] = args ? args.syncMode : undefined;
            resourceInputs["tokenUrl"] = args ? args.tokenUrl : undefined;
            resourceInputs["trustEmail"] = args ? args.trustEmail : undefined;
            resourceInputs["uiLocales"] = args ? args.uiLocales : undefined;
            resourceInputs["userInfoUrl"] = args ? args.userInfoUrl : undefined;
            resourceInputs["validateSignature"] = args ? args.validateSignature : undefined;
            resourceInputs["internalId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IdentityProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityProvider resources.
 */
export interface IdentityProviderState {
    /**
     * When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
     */
    acceptsPromptNoneForwardFromClient?: pulumi.Input<boolean>;
    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     */
    addReadTokenRoleOnCreate?: pulumi.Input<boolean>;
    /**
     * The alias uniquely identifies an identity provider, and it is also used to build the redirect uri.
     */
    alias?: pulumi.Input<string>;
    /**
     * Enable/disable authenticate users by default.
     */
    authenticateByDefault?: pulumi.Input<boolean>;
    /**
     * The Authorization Url.
     */
    authorizationUrl?: pulumi.Input<string>;
    /**
     * Does the external IDP support backchannel logout? Defaults to `true`.
     */
    backchannelSupported?: pulumi.Input<boolean>;
    /**
     * The client or client identifier registered within the identity provider.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format. Required without `clientSecretWo` and `clientSecretWoVersion`.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Version of the Client secret write-only argument
     */
    clientSecretWoVersion?: pulumi.Input<number>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
     */
    defaultScopes?: pulumi.Input<string>;
    /**
     * When `true`, disables the check for the `typ` claim of tokens received from the identity provider. Defaults to `false`.
     */
    disableTypeClaimCheck?: pulumi.Input<boolean>;
    /**
     * When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     */
    disableUserInfo?: pulumi.Input<boolean>;
    /**
     * Display name for the identity provider in the GUI.
     */
    displayName?: pulumi.Input<string>;
    /**
     * When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     */
    firstBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * A number defining the order of this identity provider in the GUI.
     */
    guiOrder?: pulumi.Input<string>;
    /**
     * When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
     */
    hideOnLoginPage?: pulumi.Input<boolean>;
    /**
     * (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
     */
    internalId?: pulumi.Input<string>;
    /**
     * The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
     */
    issuer?: pulumi.Input<string>;
    /**
     * JSON Web Key Set URL.
     */
    jwksUrl?: pulumi.Input<string>;
    /**
     * When `true`, users cannot sign-in using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     */
    linkOnly?: pulumi.Input<boolean>;
    /**
     * Pass login hint to identity provider.
     */
    loginHint?: pulumi.Input<string>;
    /**
     * The Logout URL is the end session endpoint to use to sign-out the user from external identity provider.
     */
    logoutUrl?: pulumi.Input<string>;
    /**
     * The organization domain to associate this identity provider with. it is used to map users to an organization based on their email domain and to authenticate them accordingly in the scope of the organization.
     */
    orgDomain?: pulumi.Input<string>;
    /**
     * Indicates whether to automatically redirect user to this identity provider when email domain matches domain.
     */
    orgRedirectModeEmailMatches?: pulumi.Input<boolean>;
    /**
     * The ID of the organization to link this identity provider to.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     */
    postBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
     */
    providerId?: pulumi.Input<string>;
    /**
     * The name of the realm. This is unique across Keycloak.
     */
    realm?: pulumi.Input<string>;
    /**
     * When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     */
    storeToken?: pulumi.Input<boolean>;
    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     */
    syncMode?: pulumi.Input<string>;
    /**
     * The Token URL.
     */
    tokenUrl?: pulumi.Input<string>;
    /**
     * When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
     */
    trustEmail?: pulumi.Input<boolean>;
    /**
     * Pass current locale to identity provider. Defaults to `false`.
     */
    uiLocales?: pulumi.Input<boolean>;
    /**
     * User Info URL.
     */
    userInfoUrl?: pulumi.Input<string>;
    /**
     * Enable/disable signature validation of external IDP signatures. Defaults to `false`.
     */
    validateSignature?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a IdentityProvider resource.
 */
export interface IdentityProviderArgs {
    /**
     * When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
     */
    acceptsPromptNoneForwardFromClient?: pulumi.Input<boolean>;
    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     */
    addReadTokenRoleOnCreate?: pulumi.Input<boolean>;
    /**
     * The alias uniquely identifies an identity provider, and it is also used to build the redirect uri.
     */
    alias: pulumi.Input<string>;
    /**
     * Enable/disable authenticate users by default.
     */
    authenticateByDefault?: pulumi.Input<boolean>;
    /**
     * The Authorization Url.
     */
    authorizationUrl: pulumi.Input<string>;
    /**
     * Does the external IDP support backchannel logout? Defaults to `true`.
     */
    backchannelSupported?: pulumi.Input<boolean>;
    /**
     * The client or client identifier registered within the identity provider.
     */
    clientId: pulumi.Input<string>;
    /**
     * The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format. Required without `clientSecretWo` and `clientSecretWoVersion`.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Version of the Client secret write-only argument
     */
    clientSecretWoVersion?: pulumi.Input<number>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
     */
    defaultScopes?: pulumi.Input<string>;
    /**
     * When `true`, disables the check for the `typ` claim of tokens received from the identity provider. Defaults to `false`.
     */
    disableTypeClaimCheck?: pulumi.Input<boolean>;
    /**
     * When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     */
    disableUserInfo?: pulumi.Input<boolean>;
    /**
     * Display name for the identity provider in the GUI.
     */
    displayName?: pulumi.Input<string>;
    /**
     * When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     */
    firstBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * A number defining the order of this identity provider in the GUI.
     */
    guiOrder?: pulumi.Input<string>;
    /**
     * When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
     */
    hideOnLoginPage?: pulumi.Input<boolean>;
    /**
     * The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
     */
    issuer?: pulumi.Input<string>;
    /**
     * JSON Web Key Set URL.
     */
    jwksUrl?: pulumi.Input<string>;
    /**
     * When `true`, users cannot sign-in using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     */
    linkOnly?: pulumi.Input<boolean>;
    /**
     * Pass login hint to identity provider.
     */
    loginHint?: pulumi.Input<string>;
    /**
     * The Logout URL is the end session endpoint to use to sign-out the user from external identity provider.
     */
    logoutUrl?: pulumi.Input<string>;
    /**
     * The organization domain to associate this identity provider with. it is used to map users to an organization based on their email domain and to authenticate them accordingly in the scope of the organization.
     */
    orgDomain?: pulumi.Input<string>;
    /**
     * Indicates whether to automatically redirect user to this identity provider when email domain matches domain.
     */
    orgRedirectModeEmailMatches?: pulumi.Input<boolean>;
    /**
     * The ID of the organization to link this identity provider to.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     */
    postBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
     */
    providerId?: pulumi.Input<string>;
    /**
     * The name of the realm. This is unique across Keycloak.
     */
    realm: pulumi.Input<string>;
    /**
     * When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     */
    storeToken?: pulumi.Input<boolean>;
    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     */
    syncMode?: pulumi.Input<string>;
    /**
     * The Token URL.
     */
    tokenUrl: pulumi.Input<string>;
    /**
     * When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
     */
    trustEmail?: pulumi.Input<boolean>;
    /**
     * Pass current locale to identity provider. Defaults to `false`.
     */
    uiLocales?: pulumi.Input<boolean>;
    /**
     * User Info URL.
     */
    userInfoUrl?: pulumi.Input<string>;
    /**
     * Enable/disable signature validation of external IDP signatures. Defaults to `false`.
     */
    validateSignature?: pulumi.Input<boolean>;
}
