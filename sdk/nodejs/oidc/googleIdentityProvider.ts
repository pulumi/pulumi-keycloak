// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing OIDC Identity Providers within Keycloak.
 *
 * OIDC (OpenID Connect) identity providers allows users to authenticate through a third party system using the OIDC standard.
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
 * const google = new keycloak.oidc.GoogleIdentityProvider("google", {
 *     realm: realm.id,
 *     clientId: googleIdentityProviderClientId,
 *     clientSecret: googleIdentityProviderClientSecret,
 *     trustEmail: true,
 *     hostedDomain: "example.com",
 *     syncMode: "IMPORT",
 *     extraConfig: {
 *         myCustomConfigKey: "myValue",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Google Identity providers can be imported using the format {{realm_id}}/{{idp_alias}}, where idp_alias is the identity provider alias.
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider google_identity_provider my-realm/my-google-idp
 * ```
 */
export class GoogleIdentityProvider extends pulumi.CustomResource {
    /**
     * Get an existing GoogleIdentityProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GoogleIdentityProviderState, opts?: pulumi.CustomResourceOptions): GoogleIdentityProvider {
        return new GoogleIdentityProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider';

    /**
     * Returns true if the given object is an instance of GoogleIdentityProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GoogleIdentityProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GoogleIdentityProvider.__pulumiType;
    }

    /**
     * When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
     */
    public readonly acceptsPromptNoneForwardFromClient!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     */
    public readonly addReadTokenRoleOnCreate!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) The alias for the Google identity provider.
     */
    public /*out*/ readonly alias!: pulumi.Output<string>;
    /**
     * Enable/disable authenticate users by default.
     */
    public readonly authenticateByDefault!: pulumi.Output<boolean | undefined>;
    /**
     * The client or client identifier registered within the identity provider.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
     */
    public readonly defaultScopes!: pulumi.Output<string | undefined>;
    /**
     * When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     */
    public readonly disableUserInfo!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) Display name for the Google identity provider in the GUI.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly extraConfig!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     */
    public readonly firstBrokerLoginFlowAlias!: pulumi.Output<string | undefined>;
    /**
     * A number defining the order of this identity provider in the GUI.
     */
    public readonly guiOrder!: pulumi.Output<string | undefined>;
    /**
     * When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
     */
    public readonly hideOnLoginPage!: pulumi.Output<boolean | undefined>;
    /**
     * Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
     */
    public readonly hostedDomain!: pulumi.Output<string | undefined>;
    /**
     * (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
     */
    public /*out*/ readonly internalId!: pulumi.Output<string>;
    /**
     * When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     */
    public readonly linkOnly!: pulumi.Output<boolean | undefined>;
    /**
     * The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     */
    public readonly postBrokerLoginFlowAlias!: pulumi.Output<string | undefined>;
    /**
     * The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
     */
    public readonly providerId!: pulumi.Output<string | undefined>;
    /**
     * The name of the realm. This is unique across Keycloak.
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * Sets the "accessType" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
     */
    public readonly requestRefreshToken!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     */
    public readonly storeToken!: pulumi.Output<boolean | undefined>;
    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     */
    public readonly syncMode!: pulumi.Output<string | undefined>;
    /**
     * When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
     */
    public readonly trustEmail!: pulumi.Output<boolean | undefined>;
    /**
     * Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
     */
    public readonly useUserIpParam!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GoogleIdentityProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GoogleIdentityProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GoogleIdentityProviderArgs | GoogleIdentityProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GoogleIdentityProviderState | undefined;
            resourceInputs["acceptsPromptNoneForwardFromClient"] = state ? state.acceptsPromptNoneForwardFromClient : undefined;
            resourceInputs["addReadTokenRoleOnCreate"] = state ? state.addReadTokenRoleOnCreate : undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["authenticateByDefault"] = state ? state.authenticateByDefault : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["defaultScopes"] = state ? state.defaultScopes : undefined;
            resourceInputs["disableUserInfo"] = state ? state.disableUserInfo : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["extraConfig"] = state ? state.extraConfig : undefined;
            resourceInputs["firstBrokerLoginFlowAlias"] = state ? state.firstBrokerLoginFlowAlias : undefined;
            resourceInputs["guiOrder"] = state ? state.guiOrder : undefined;
            resourceInputs["hideOnLoginPage"] = state ? state.hideOnLoginPage : undefined;
            resourceInputs["hostedDomain"] = state ? state.hostedDomain : undefined;
            resourceInputs["internalId"] = state ? state.internalId : undefined;
            resourceInputs["linkOnly"] = state ? state.linkOnly : undefined;
            resourceInputs["postBrokerLoginFlowAlias"] = state ? state.postBrokerLoginFlowAlias : undefined;
            resourceInputs["providerId"] = state ? state.providerId : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["requestRefreshToken"] = state ? state.requestRefreshToken : undefined;
            resourceInputs["storeToken"] = state ? state.storeToken : undefined;
            resourceInputs["syncMode"] = state ? state.syncMode : undefined;
            resourceInputs["trustEmail"] = state ? state.trustEmail : undefined;
            resourceInputs["useUserIpParam"] = state ? state.useUserIpParam : undefined;
        } else {
            const args = argsOrState as GoogleIdentityProviderArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if ((!args || args.realm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realm'");
            }
            resourceInputs["acceptsPromptNoneForwardFromClient"] = args ? args.acceptsPromptNoneForwardFromClient : undefined;
            resourceInputs["addReadTokenRoleOnCreate"] = args ? args.addReadTokenRoleOnCreate : undefined;
            resourceInputs["authenticateByDefault"] = args ? args.authenticateByDefault : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["defaultScopes"] = args ? args.defaultScopes : undefined;
            resourceInputs["disableUserInfo"] = args ? args.disableUserInfo : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["extraConfig"] = args ? args.extraConfig : undefined;
            resourceInputs["firstBrokerLoginFlowAlias"] = args ? args.firstBrokerLoginFlowAlias : undefined;
            resourceInputs["guiOrder"] = args ? args.guiOrder : undefined;
            resourceInputs["hideOnLoginPage"] = args ? args.hideOnLoginPage : undefined;
            resourceInputs["hostedDomain"] = args ? args.hostedDomain : undefined;
            resourceInputs["linkOnly"] = args ? args.linkOnly : undefined;
            resourceInputs["postBrokerLoginFlowAlias"] = args ? args.postBrokerLoginFlowAlias : undefined;
            resourceInputs["providerId"] = args ? args.providerId : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["requestRefreshToken"] = args ? args.requestRefreshToken : undefined;
            resourceInputs["storeToken"] = args ? args.storeToken : undefined;
            resourceInputs["syncMode"] = args ? args.syncMode : undefined;
            resourceInputs["trustEmail"] = args ? args.trustEmail : undefined;
            resourceInputs["useUserIpParam"] = args ? args.useUserIpParam : undefined;
            resourceInputs["alias"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["internalId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(GoogleIdentityProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GoogleIdentityProvider resources.
 */
export interface GoogleIdentityProviderState {
    /**
     * When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
     */
    acceptsPromptNoneForwardFromClient?: pulumi.Input<boolean>;
    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     */
    addReadTokenRoleOnCreate?: pulumi.Input<boolean>;
    /**
     * (Computed) The alias for the Google identity provider.
     */
    alias?: pulumi.Input<string>;
    /**
     * Enable/disable authenticate users by default.
     */
    authenticateByDefault?: pulumi.Input<boolean>;
    /**
     * The client or client identifier registered within the identity provider.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
     */
    defaultScopes?: pulumi.Input<string>;
    /**
     * When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     */
    disableUserInfo?: pulumi.Input<boolean>;
    /**
     * (Computed) Display name for the Google identity provider in the GUI.
     */
    displayName?: pulumi.Input<string>;
    /**
     * When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     */
    firstBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * A number defining the order of this identity provider in the GUI.
     */
    guiOrder?: pulumi.Input<string>;
    /**
     * When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
     */
    hideOnLoginPage?: pulumi.Input<boolean>;
    /**
     * Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
     */
    hostedDomain?: pulumi.Input<string>;
    /**
     * (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
     */
    internalId?: pulumi.Input<string>;
    /**
     * When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     */
    linkOnly?: pulumi.Input<boolean>;
    /**
     * The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     */
    postBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
     */
    providerId?: pulumi.Input<string>;
    /**
     * The name of the realm. This is unique across Keycloak.
     */
    realm?: pulumi.Input<string>;
    /**
     * Sets the "accessType" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
     */
    requestRefreshToken?: pulumi.Input<boolean>;
    /**
     * When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     */
    storeToken?: pulumi.Input<boolean>;
    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     */
    syncMode?: pulumi.Input<string>;
    /**
     * When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
     */
    trustEmail?: pulumi.Input<boolean>;
    /**
     * Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
     */
    useUserIpParam?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GoogleIdentityProvider resource.
 */
export interface GoogleIdentityProviderArgs {
    /**
     * When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
     */
    acceptsPromptNoneForwardFromClient?: pulumi.Input<boolean>;
    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     */
    addReadTokenRoleOnCreate?: pulumi.Input<boolean>;
    /**
     * Enable/disable authenticate users by default.
     */
    authenticateByDefault?: pulumi.Input<boolean>;
    /**
     * The client or client identifier registered within the identity provider.
     */
    clientId: pulumi.Input<string>;
    /**
     * The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
     */
    clientSecret: pulumi.Input<string>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
     */
    defaultScopes?: pulumi.Input<string>;
    /**
     * When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     */
    disableUserInfo?: pulumi.Input<boolean>;
    /**
     * When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     */
    firstBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * A number defining the order of this identity provider in the GUI.
     */
    guiOrder?: pulumi.Input<string>;
    /**
     * When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
     */
    hideOnLoginPage?: pulumi.Input<boolean>;
    /**
     * Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
     */
    hostedDomain?: pulumi.Input<string>;
    /**
     * When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     */
    linkOnly?: pulumi.Input<boolean>;
    /**
     * The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     */
    postBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
     */
    providerId?: pulumi.Input<string>;
    /**
     * The name of the realm. This is unique across Keycloak.
     */
    realm: pulumi.Input<string>;
    /**
     * Sets the "accessType" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
     */
    requestRefreshToken?: pulumi.Input<boolean>;
    /**
     * When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     */
    storeToken?: pulumi.Input<boolean>;
    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     */
    syncMode?: pulumi.Input<string>;
    /**
     * When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
     */
    trustEmail?: pulumi.Input<boolean>;
    /**
     * Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
     */
    useUserIpParam?: pulumi.Input<boolean>;
}
