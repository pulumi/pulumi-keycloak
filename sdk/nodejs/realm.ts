// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing Realms within Keycloak.
 *
 * A realm manages a logical collection of users, credentials, roles, and groups. Users log in to realms and can be federated
 * from multiple sources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     accessCodeLifespan: "1h",
 *     attributes: {
 *         mycustomAttribute: "myCustomValue",
 *     },
 *     displayName: "my realm",
 *     displayNameHtml: "<b>my realm</b>",
 *     enabled: true,
 *     internationalization: {
 *         defaultLocale: "en",
 *         supportedLocales: [
 *             "en",
 *             "de",
 *             "es",
 *         ],
 *     },
 *     loginTheme: "base",
 *     passwordPolicy: "upperCase(1) and length(8) and forceExpiredPasswordChange(365) and notUsername",
 *     realm: "my-realm",
 *     securityDefenses: {
 *         bruteForceDetection: {
 *             failureResetTimeSeconds: 43200,
 *             maxFailureWaitSeconds: 900,
 *             maxLoginFailures: 30,
 *             minimumQuickLoginWaitSeconds: 60,
 *             permanentLockout: false,
 *             quickLoginCheckMilliSeconds: 1000,
 *             waitIncrementSeconds: 60,
 *         },
 *         headers: {
 *             contentSecurityPolicy: "frame-src 'self'; frame-ancestors 'self'; object-src 'none';",
 *             contentSecurityPolicyReportOnly: "",
 *             strictTransportSecurity: "max-age=31536000; includeSubDomains",
 *             xContentTypeOptions: "nosniff",
 *             xFrameOptions: "DENY",
 *             xRobotsTag: "none",
 *             xXssProtection: "1; mode=block",
 *         },
 *     },
 *     smtpServer: {
 *         auth: {
 *             password: "password",
 *             username: "tom",
 *         },
 *         from: "example@example.com",
 *         host: "smtp.example.com",
 *     },
 *     sslRequired: "external",
 *     webAuthnPolicy: {
 *         relyingPartyEntityName: "Example",
 *         relyingPartyId: "keycloak.example.com",
 *         signatureAlgorithms: [
 *             "ES256",
 *             "RS256",
 *         ],
 *     },
 * });
 * ```
 * ## Default Client Scopes
 *
 * - `defaultDefaultClientScopes` - (Optional) A list of default default client scopes to be used for client definitions. Defaults to `[]` or keycloak's built-in default default client-scopes.
 * - `defaultOptionalClientScopes` - (Optional) A list of default optional client scopes to be used for client definitions. Defaults to `[]` or keycloak's built-in default optional client-scopes.
 *
 * ## Import
 *
 * Realms can be imported using their name. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:index/realm:Realm realm my-realm
 * ```
 */
export class Realm extends pulumi.CustomResource {
    /**
     * Get an existing Realm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RealmState, opts?: pulumi.CustomResourceOptions): Realm {
        return new Realm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/realm:Realm';

    /**
     * Returns true if the given object is an instance of Realm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Realm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Realm.__pulumiType;
    }

    /**
     * The maximum amount of time a client has to finish the authorization code flow.
     */
    public readonly accessCodeLifespan!: pulumi.Output<string>;
    /**
     * The maximum amount of time a user is permitted to stay on the login page before the authentication process must be restarted.
     */
    public readonly accessCodeLifespanLogin!: pulumi.Output<string>;
    /**
     * The maximum amount of time a user has to complete login related actions, such as updating a password.
     */
    public readonly accessCodeLifespanUserAction!: pulumi.Output<string>;
    /**
     * The amount of time an access token can be used before it expires.
     */
    public readonly accessTokenLifespan!: pulumi.Output<string>;
    /**
     * The amount of time an access token issued with the OpenID Connect Implicit Flow can be used before it expires.
     */
    public readonly accessTokenLifespanForImplicitFlow!: pulumi.Output<string>;
    /**
     * Used for account management pages.
     */
    public readonly accountTheme!: pulumi.Output<string | undefined>;
    /**
     * The maximum time a user has to use an admin-generated permit before it expires.
     */
    public readonly actionTokenGeneratedByAdminLifespan!: pulumi.Output<string>;
    /**
     * The maximum time a user has to use a user-generated permit before it expires.
     */
    public readonly actionTokenGeneratedByUserLifespan!: pulumi.Output<string>;
    /**
     * Used for the admin console.
     */
    public readonly adminTheme!: pulumi.Output<string | undefined>;
    /**
     * A map of custom attributes to add to the realm.
     */
    public readonly attributes!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The desired flow for browser authentication. Defaults to `browser`.
     */
    public readonly browserFlow!: pulumi.Output<string | undefined>;
    /**
     * The desired flow for client authentication. Defaults to `clients`.
     */
    public readonly clientAuthenticationFlow!: pulumi.Output<string | undefined>;
    public readonly defaultDefaultClientScopes!: pulumi.Output<string[] | undefined>;
    public readonly defaultOptionalClientScopes!: pulumi.Output<string[] | undefined>;
    /**
     * Default algorithm used to sign tokens for the realm.
     */
    public readonly defaultSignatureAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * The desired flow for direct access authentication. Defaults to `direct grant`.
     */
    public readonly directGrantFlow!: pulumi.Output<string | undefined>;
    /**
     * The display name for the realm that is shown when logging in to the admin console.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The display name for the realm that is rendered as HTML on the screen when logging in to the admin console.
     */
    public readonly displayNameHtml!: pulumi.Output<string | undefined>;
    /**
     * The desired flow for Docker authentication. Defaults to `docker auth`.
     */
    public readonly dockerAuthenticationFlow!: pulumi.Output<string | undefined>;
    /**
     * When true, multiple users will be allowed to have the same email address. This argument must be set to `false` if `loginWithEmailAllowed` is set to `true`.
     */
    public readonly duplicateEmailsAllowed!: pulumi.Output<boolean>;
    /**
     * When true, the username field is editable.
     */
    public readonly editUsernameAllowed!: pulumi.Output<boolean>;
    /**
     * Used for emails that are sent by Keycloak.
     */
    public readonly emailTheme!: pulumi.Output<string | undefined>;
    /**
     * When `false`, users and clients will not be able to access this realm. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly internalId!: pulumi.Output<string>;
    public readonly internationalization!: pulumi.Output<outputs.RealmInternationalization | undefined>;
    /**
     * Used for the login, forgot password, and registration pages.
     */
    public readonly loginTheme!: pulumi.Output<string | undefined>;
    /**
     * When true, users may log in with their email address.
     */
    public readonly loginWithEmailAllowed!: pulumi.Output<boolean>;
    /**
     * The amount of time an offline session can be idle before it expires.
     */
    public readonly offlineSessionIdleTimeout!: pulumi.Output<string>;
    /**
     * The maximum amount of time before an offline session expires regardless of activity.
     */
    public readonly offlineSessionMaxLifespan!: pulumi.Output<string>;
    /**
     * Enable `offlineSessionMaxLifespan`.
     */
    public readonly offlineSessionMaxLifespanEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The password policy for users within the realm.
     */
    public readonly passwordPolicy!: pulumi.Output<string | undefined>;
    /**
     * The name of the realm. This is unique across Keycloak. This will also be used as the realm's internal ID within Keycloak.
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * Maximum number of times a refresh token can be reused before they are revoked. If unspecified and 'revoke_refresh_token' is enabled the default value is 0 and refresh tokens can not be reused.
     */
    public readonly refreshTokenMaxReuse!: pulumi.Output<number | undefined>;
    /**
     * When true, user registration will be enabled, and a link for registration will be displayed on the login page.
     */
    public readonly registrationAllowed!: pulumi.Output<boolean>;
    /**
     * When true, the user's email will be used as their username during registration.
     */
    public readonly registrationEmailAsUsername!: pulumi.Output<boolean>;
    /**
     * The desired flow for user registration. Defaults to `registration`.
     */
    public readonly registrationFlow!: pulumi.Output<string | undefined>;
    /**
     * When true, a "remember me" checkbox will be displayed on the login page, and the user's session will not expire between browser restarts.
     */
    public readonly rememberMe!: pulumi.Output<boolean>;
    /**
     * The desired flow to use when a user attempts to reset their credentials. Defaults to `reset credentials`.
     */
    public readonly resetCredentialsFlow!: pulumi.Output<string | undefined>;
    /**
     * When true, a "forgot password" link will be displayed on the login page.
     */
    public readonly resetPasswordAllowed!: pulumi.Output<boolean>;
    /**
     * If enabled a refresh token can only be used number of times specified in 'refresh_token_max_reuse' before they are revoked. If unspecified, refresh tokens can be reused.
     */
    public readonly revokeRefreshToken!: pulumi.Output<boolean | undefined>;
    public readonly securityDefenses!: pulumi.Output<outputs.RealmSecurityDefenses | undefined>;
    public readonly smtpServer!: pulumi.Output<outputs.RealmSmtpServer | undefined>;
    /**
     * Can be one of following values: 'none, 'external' or 'all'
     */
    public readonly sslRequired!: pulumi.Output<string | undefined>;
    /**
     * The amount of time a session can be idle before it expires.
     */
    public readonly ssoSessionIdleTimeout!: pulumi.Output<string>;
    public readonly ssoSessionIdleTimeoutRememberMe!: pulumi.Output<string>;
    /**
     * The maximum amount of time before a session expires regardless of activity.
     */
    public readonly ssoSessionMaxLifespan!: pulumi.Output<string>;
    public readonly ssoSessionMaxLifespanRememberMe!: pulumi.Output<string>;
    /**
     * When `true`, users are allowed to manage their own resources. Defaults to `false`.
     */
    public readonly userManagedAccess!: pulumi.Output<boolean | undefined>;
    /**
     * When true, users are required to verify their email address after registration and after email address changes.
     */
    public readonly verifyEmail!: pulumi.Output<boolean>;
    /**
     * Configuration for WebAuthn Passwordless Policy authentication.
     */
    public readonly webAuthnPasswordlessPolicy!: pulumi.Output<outputs.RealmWebAuthnPasswordlessPolicy>;
    /**
     * Configuration for WebAuthn Policy authentication.
     */
    public readonly webAuthnPolicy!: pulumi.Output<outputs.RealmWebAuthnPolicy>;

    /**
     * Create a Realm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RealmArgs | RealmState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RealmState | undefined;
            inputs["accessCodeLifespan"] = state ? state.accessCodeLifespan : undefined;
            inputs["accessCodeLifespanLogin"] = state ? state.accessCodeLifespanLogin : undefined;
            inputs["accessCodeLifespanUserAction"] = state ? state.accessCodeLifespanUserAction : undefined;
            inputs["accessTokenLifespan"] = state ? state.accessTokenLifespan : undefined;
            inputs["accessTokenLifespanForImplicitFlow"] = state ? state.accessTokenLifespanForImplicitFlow : undefined;
            inputs["accountTheme"] = state ? state.accountTheme : undefined;
            inputs["actionTokenGeneratedByAdminLifespan"] = state ? state.actionTokenGeneratedByAdminLifespan : undefined;
            inputs["actionTokenGeneratedByUserLifespan"] = state ? state.actionTokenGeneratedByUserLifespan : undefined;
            inputs["adminTheme"] = state ? state.adminTheme : undefined;
            inputs["attributes"] = state ? state.attributes : undefined;
            inputs["browserFlow"] = state ? state.browserFlow : undefined;
            inputs["clientAuthenticationFlow"] = state ? state.clientAuthenticationFlow : undefined;
            inputs["defaultDefaultClientScopes"] = state ? state.defaultDefaultClientScopes : undefined;
            inputs["defaultOptionalClientScopes"] = state ? state.defaultOptionalClientScopes : undefined;
            inputs["defaultSignatureAlgorithm"] = state ? state.defaultSignatureAlgorithm : undefined;
            inputs["directGrantFlow"] = state ? state.directGrantFlow : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["displayNameHtml"] = state ? state.displayNameHtml : undefined;
            inputs["dockerAuthenticationFlow"] = state ? state.dockerAuthenticationFlow : undefined;
            inputs["duplicateEmailsAllowed"] = state ? state.duplicateEmailsAllowed : undefined;
            inputs["editUsernameAllowed"] = state ? state.editUsernameAllowed : undefined;
            inputs["emailTheme"] = state ? state.emailTheme : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["internalId"] = state ? state.internalId : undefined;
            inputs["internationalization"] = state ? state.internationalization : undefined;
            inputs["loginTheme"] = state ? state.loginTheme : undefined;
            inputs["loginWithEmailAllowed"] = state ? state.loginWithEmailAllowed : undefined;
            inputs["offlineSessionIdleTimeout"] = state ? state.offlineSessionIdleTimeout : undefined;
            inputs["offlineSessionMaxLifespan"] = state ? state.offlineSessionMaxLifespan : undefined;
            inputs["offlineSessionMaxLifespanEnabled"] = state ? state.offlineSessionMaxLifespanEnabled : undefined;
            inputs["passwordPolicy"] = state ? state.passwordPolicy : undefined;
            inputs["realm"] = state ? state.realm : undefined;
            inputs["refreshTokenMaxReuse"] = state ? state.refreshTokenMaxReuse : undefined;
            inputs["registrationAllowed"] = state ? state.registrationAllowed : undefined;
            inputs["registrationEmailAsUsername"] = state ? state.registrationEmailAsUsername : undefined;
            inputs["registrationFlow"] = state ? state.registrationFlow : undefined;
            inputs["rememberMe"] = state ? state.rememberMe : undefined;
            inputs["resetCredentialsFlow"] = state ? state.resetCredentialsFlow : undefined;
            inputs["resetPasswordAllowed"] = state ? state.resetPasswordAllowed : undefined;
            inputs["revokeRefreshToken"] = state ? state.revokeRefreshToken : undefined;
            inputs["securityDefenses"] = state ? state.securityDefenses : undefined;
            inputs["smtpServer"] = state ? state.smtpServer : undefined;
            inputs["sslRequired"] = state ? state.sslRequired : undefined;
            inputs["ssoSessionIdleTimeout"] = state ? state.ssoSessionIdleTimeout : undefined;
            inputs["ssoSessionIdleTimeoutRememberMe"] = state ? state.ssoSessionIdleTimeoutRememberMe : undefined;
            inputs["ssoSessionMaxLifespan"] = state ? state.ssoSessionMaxLifespan : undefined;
            inputs["ssoSessionMaxLifespanRememberMe"] = state ? state.ssoSessionMaxLifespanRememberMe : undefined;
            inputs["userManagedAccess"] = state ? state.userManagedAccess : undefined;
            inputs["verifyEmail"] = state ? state.verifyEmail : undefined;
            inputs["webAuthnPasswordlessPolicy"] = state ? state.webAuthnPasswordlessPolicy : undefined;
            inputs["webAuthnPolicy"] = state ? state.webAuthnPolicy : undefined;
        } else {
            const args = argsOrState as RealmArgs | undefined;
            if ((!args || args.realm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realm'");
            }
            inputs["accessCodeLifespan"] = args ? args.accessCodeLifespan : undefined;
            inputs["accessCodeLifespanLogin"] = args ? args.accessCodeLifespanLogin : undefined;
            inputs["accessCodeLifespanUserAction"] = args ? args.accessCodeLifespanUserAction : undefined;
            inputs["accessTokenLifespan"] = args ? args.accessTokenLifespan : undefined;
            inputs["accessTokenLifespanForImplicitFlow"] = args ? args.accessTokenLifespanForImplicitFlow : undefined;
            inputs["accountTheme"] = args ? args.accountTheme : undefined;
            inputs["actionTokenGeneratedByAdminLifespan"] = args ? args.actionTokenGeneratedByAdminLifespan : undefined;
            inputs["actionTokenGeneratedByUserLifespan"] = args ? args.actionTokenGeneratedByUserLifespan : undefined;
            inputs["adminTheme"] = args ? args.adminTheme : undefined;
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["browserFlow"] = args ? args.browserFlow : undefined;
            inputs["clientAuthenticationFlow"] = args ? args.clientAuthenticationFlow : undefined;
            inputs["defaultDefaultClientScopes"] = args ? args.defaultDefaultClientScopes : undefined;
            inputs["defaultOptionalClientScopes"] = args ? args.defaultOptionalClientScopes : undefined;
            inputs["defaultSignatureAlgorithm"] = args ? args.defaultSignatureAlgorithm : undefined;
            inputs["directGrantFlow"] = args ? args.directGrantFlow : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["displayNameHtml"] = args ? args.displayNameHtml : undefined;
            inputs["dockerAuthenticationFlow"] = args ? args.dockerAuthenticationFlow : undefined;
            inputs["duplicateEmailsAllowed"] = args ? args.duplicateEmailsAllowed : undefined;
            inputs["editUsernameAllowed"] = args ? args.editUsernameAllowed : undefined;
            inputs["emailTheme"] = args ? args.emailTheme : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["internationalization"] = args ? args.internationalization : undefined;
            inputs["loginTheme"] = args ? args.loginTheme : undefined;
            inputs["loginWithEmailAllowed"] = args ? args.loginWithEmailAllowed : undefined;
            inputs["offlineSessionIdleTimeout"] = args ? args.offlineSessionIdleTimeout : undefined;
            inputs["offlineSessionMaxLifespan"] = args ? args.offlineSessionMaxLifespan : undefined;
            inputs["offlineSessionMaxLifespanEnabled"] = args ? args.offlineSessionMaxLifespanEnabled : undefined;
            inputs["passwordPolicy"] = args ? args.passwordPolicy : undefined;
            inputs["realm"] = args ? args.realm : undefined;
            inputs["refreshTokenMaxReuse"] = args ? args.refreshTokenMaxReuse : undefined;
            inputs["registrationAllowed"] = args ? args.registrationAllowed : undefined;
            inputs["registrationEmailAsUsername"] = args ? args.registrationEmailAsUsername : undefined;
            inputs["registrationFlow"] = args ? args.registrationFlow : undefined;
            inputs["rememberMe"] = args ? args.rememberMe : undefined;
            inputs["resetCredentialsFlow"] = args ? args.resetCredentialsFlow : undefined;
            inputs["resetPasswordAllowed"] = args ? args.resetPasswordAllowed : undefined;
            inputs["revokeRefreshToken"] = args ? args.revokeRefreshToken : undefined;
            inputs["securityDefenses"] = args ? args.securityDefenses : undefined;
            inputs["smtpServer"] = args ? args.smtpServer : undefined;
            inputs["sslRequired"] = args ? args.sslRequired : undefined;
            inputs["ssoSessionIdleTimeout"] = args ? args.ssoSessionIdleTimeout : undefined;
            inputs["ssoSessionIdleTimeoutRememberMe"] = args ? args.ssoSessionIdleTimeoutRememberMe : undefined;
            inputs["ssoSessionMaxLifespan"] = args ? args.ssoSessionMaxLifespan : undefined;
            inputs["ssoSessionMaxLifespanRememberMe"] = args ? args.ssoSessionMaxLifespanRememberMe : undefined;
            inputs["userManagedAccess"] = args ? args.userManagedAccess : undefined;
            inputs["verifyEmail"] = args ? args.verifyEmail : undefined;
            inputs["webAuthnPasswordlessPolicy"] = args ? args.webAuthnPasswordlessPolicy : undefined;
            inputs["webAuthnPolicy"] = args ? args.webAuthnPolicy : undefined;
            inputs["internalId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Realm.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Realm resources.
 */
export interface RealmState {
    /**
     * The maximum amount of time a client has to finish the authorization code flow.
     */
    readonly accessCodeLifespan?: pulumi.Input<string>;
    /**
     * The maximum amount of time a user is permitted to stay on the login page before the authentication process must be restarted.
     */
    readonly accessCodeLifespanLogin?: pulumi.Input<string>;
    /**
     * The maximum amount of time a user has to complete login related actions, such as updating a password.
     */
    readonly accessCodeLifespanUserAction?: pulumi.Input<string>;
    /**
     * The amount of time an access token can be used before it expires.
     */
    readonly accessTokenLifespan?: pulumi.Input<string>;
    /**
     * The amount of time an access token issued with the OpenID Connect Implicit Flow can be used before it expires.
     */
    readonly accessTokenLifespanForImplicitFlow?: pulumi.Input<string>;
    /**
     * Used for account management pages.
     */
    readonly accountTheme?: pulumi.Input<string>;
    /**
     * The maximum time a user has to use an admin-generated permit before it expires.
     */
    readonly actionTokenGeneratedByAdminLifespan?: pulumi.Input<string>;
    /**
     * The maximum time a user has to use a user-generated permit before it expires.
     */
    readonly actionTokenGeneratedByUserLifespan?: pulumi.Input<string>;
    /**
     * Used for the admin console.
     */
    readonly adminTheme?: pulumi.Input<string>;
    /**
     * A map of custom attributes to add to the realm.
     */
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * The desired flow for browser authentication. Defaults to `browser`.
     */
    readonly browserFlow?: pulumi.Input<string>;
    /**
     * The desired flow for client authentication. Defaults to `clients`.
     */
    readonly clientAuthenticationFlow?: pulumi.Input<string>;
    readonly defaultDefaultClientScopes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly defaultOptionalClientScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default algorithm used to sign tokens for the realm.
     */
    readonly defaultSignatureAlgorithm?: pulumi.Input<string>;
    /**
     * The desired flow for direct access authentication. Defaults to `direct grant`.
     */
    readonly directGrantFlow?: pulumi.Input<string>;
    /**
     * The display name for the realm that is shown when logging in to the admin console.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The display name for the realm that is rendered as HTML on the screen when logging in to the admin console.
     */
    readonly displayNameHtml?: pulumi.Input<string>;
    /**
     * The desired flow for Docker authentication. Defaults to `docker auth`.
     */
    readonly dockerAuthenticationFlow?: pulumi.Input<string>;
    /**
     * When true, multiple users will be allowed to have the same email address. This argument must be set to `false` if `loginWithEmailAllowed` is set to `true`.
     */
    readonly duplicateEmailsAllowed?: pulumi.Input<boolean>;
    /**
     * When true, the username field is editable.
     */
    readonly editUsernameAllowed?: pulumi.Input<boolean>;
    /**
     * Used for emails that are sent by Keycloak.
     */
    readonly emailTheme?: pulumi.Input<string>;
    /**
     * When `false`, users and clients will not be able to access this realm. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    readonly internalId?: pulumi.Input<string>;
    readonly internationalization?: pulumi.Input<inputs.RealmInternationalization>;
    /**
     * Used for the login, forgot password, and registration pages.
     */
    readonly loginTheme?: pulumi.Input<string>;
    /**
     * When true, users may log in with their email address.
     */
    readonly loginWithEmailAllowed?: pulumi.Input<boolean>;
    /**
     * The amount of time an offline session can be idle before it expires.
     */
    readonly offlineSessionIdleTimeout?: pulumi.Input<string>;
    /**
     * The maximum amount of time before an offline session expires regardless of activity.
     */
    readonly offlineSessionMaxLifespan?: pulumi.Input<string>;
    /**
     * Enable `offlineSessionMaxLifespan`.
     */
    readonly offlineSessionMaxLifespanEnabled?: pulumi.Input<boolean>;
    /**
     * The password policy for users within the realm.
     */
    readonly passwordPolicy?: pulumi.Input<string>;
    /**
     * The name of the realm. This is unique across Keycloak. This will also be used as the realm's internal ID within Keycloak.
     */
    readonly realm?: pulumi.Input<string>;
    /**
     * Maximum number of times a refresh token can be reused before they are revoked. If unspecified and 'revoke_refresh_token' is enabled the default value is 0 and refresh tokens can not be reused.
     */
    readonly refreshTokenMaxReuse?: pulumi.Input<number>;
    /**
     * When true, user registration will be enabled, and a link for registration will be displayed on the login page.
     */
    readonly registrationAllowed?: pulumi.Input<boolean>;
    /**
     * When true, the user's email will be used as their username during registration.
     */
    readonly registrationEmailAsUsername?: pulumi.Input<boolean>;
    /**
     * The desired flow for user registration. Defaults to `registration`.
     */
    readonly registrationFlow?: pulumi.Input<string>;
    /**
     * When true, a "remember me" checkbox will be displayed on the login page, and the user's session will not expire between browser restarts.
     */
    readonly rememberMe?: pulumi.Input<boolean>;
    /**
     * The desired flow to use when a user attempts to reset their credentials. Defaults to `reset credentials`.
     */
    readonly resetCredentialsFlow?: pulumi.Input<string>;
    /**
     * When true, a "forgot password" link will be displayed on the login page.
     */
    readonly resetPasswordAllowed?: pulumi.Input<boolean>;
    /**
     * If enabled a refresh token can only be used number of times specified in 'refresh_token_max_reuse' before they are revoked. If unspecified, refresh tokens can be reused.
     */
    readonly revokeRefreshToken?: pulumi.Input<boolean>;
    readonly securityDefenses?: pulumi.Input<inputs.RealmSecurityDefenses>;
    readonly smtpServer?: pulumi.Input<inputs.RealmSmtpServer>;
    /**
     * Can be one of following values: 'none, 'external' or 'all'
     */
    readonly sslRequired?: pulumi.Input<string>;
    /**
     * The amount of time a session can be idle before it expires.
     */
    readonly ssoSessionIdleTimeout?: pulumi.Input<string>;
    readonly ssoSessionIdleTimeoutRememberMe?: pulumi.Input<string>;
    /**
     * The maximum amount of time before a session expires regardless of activity.
     */
    readonly ssoSessionMaxLifespan?: pulumi.Input<string>;
    readonly ssoSessionMaxLifespanRememberMe?: pulumi.Input<string>;
    /**
     * When `true`, users are allowed to manage their own resources. Defaults to `false`.
     */
    readonly userManagedAccess?: pulumi.Input<boolean>;
    /**
     * When true, users are required to verify their email address after registration and after email address changes.
     */
    readonly verifyEmail?: pulumi.Input<boolean>;
    /**
     * Configuration for WebAuthn Passwordless Policy authentication.
     */
    readonly webAuthnPasswordlessPolicy?: pulumi.Input<inputs.RealmWebAuthnPasswordlessPolicy>;
    /**
     * Configuration for WebAuthn Policy authentication.
     */
    readonly webAuthnPolicy?: pulumi.Input<inputs.RealmWebAuthnPolicy>;
}

/**
 * The set of arguments for constructing a Realm resource.
 */
export interface RealmArgs {
    /**
     * The maximum amount of time a client has to finish the authorization code flow.
     */
    readonly accessCodeLifespan?: pulumi.Input<string>;
    /**
     * The maximum amount of time a user is permitted to stay on the login page before the authentication process must be restarted.
     */
    readonly accessCodeLifespanLogin?: pulumi.Input<string>;
    /**
     * The maximum amount of time a user has to complete login related actions, such as updating a password.
     */
    readonly accessCodeLifespanUserAction?: pulumi.Input<string>;
    /**
     * The amount of time an access token can be used before it expires.
     */
    readonly accessTokenLifespan?: pulumi.Input<string>;
    /**
     * The amount of time an access token issued with the OpenID Connect Implicit Flow can be used before it expires.
     */
    readonly accessTokenLifespanForImplicitFlow?: pulumi.Input<string>;
    /**
     * Used for account management pages.
     */
    readonly accountTheme?: pulumi.Input<string>;
    /**
     * The maximum time a user has to use an admin-generated permit before it expires.
     */
    readonly actionTokenGeneratedByAdminLifespan?: pulumi.Input<string>;
    /**
     * The maximum time a user has to use a user-generated permit before it expires.
     */
    readonly actionTokenGeneratedByUserLifespan?: pulumi.Input<string>;
    /**
     * Used for the admin console.
     */
    readonly adminTheme?: pulumi.Input<string>;
    /**
     * A map of custom attributes to add to the realm.
     */
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * The desired flow for browser authentication. Defaults to `browser`.
     */
    readonly browserFlow?: pulumi.Input<string>;
    /**
     * The desired flow for client authentication. Defaults to `clients`.
     */
    readonly clientAuthenticationFlow?: pulumi.Input<string>;
    readonly defaultDefaultClientScopes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly defaultOptionalClientScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default algorithm used to sign tokens for the realm.
     */
    readonly defaultSignatureAlgorithm?: pulumi.Input<string>;
    /**
     * The desired flow for direct access authentication. Defaults to `direct grant`.
     */
    readonly directGrantFlow?: pulumi.Input<string>;
    /**
     * The display name for the realm that is shown when logging in to the admin console.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The display name for the realm that is rendered as HTML on the screen when logging in to the admin console.
     */
    readonly displayNameHtml?: pulumi.Input<string>;
    /**
     * The desired flow for Docker authentication. Defaults to `docker auth`.
     */
    readonly dockerAuthenticationFlow?: pulumi.Input<string>;
    /**
     * When true, multiple users will be allowed to have the same email address. This argument must be set to `false` if `loginWithEmailAllowed` is set to `true`.
     */
    readonly duplicateEmailsAllowed?: pulumi.Input<boolean>;
    /**
     * When true, the username field is editable.
     */
    readonly editUsernameAllowed?: pulumi.Input<boolean>;
    /**
     * Used for emails that are sent by Keycloak.
     */
    readonly emailTheme?: pulumi.Input<string>;
    /**
     * When `false`, users and clients will not be able to access this realm. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    readonly internationalization?: pulumi.Input<inputs.RealmInternationalization>;
    /**
     * Used for the login, forgot password, and registration pages.
     */
    readonly loginTheme?: pulumi.Input<string>;
    /**
     * When true, users may log in with their email address.
     */
    readonly loginWithEmailAllowed?: pulumi.Input<boolean>;
    /**
     * The amount of time an offline session can be idle before it expires.
     */
    readonly offlineSessionIdleTimeout?: pulumi.Input<string>;
    /**
     * The maximum amount of time before an offline session expires regardless of activity.
     */
    readonly offlineSessionMaxLifespan?: pulumi.Input<string>;
    /**
     * Enable `offlineSessionMaxLifespan`.
     */
    readonly offlineSessionMaxLifespanEnabled?: pulumi.Input<boolean>;
    /**
     * The password policy for users within the realm.
     */
    readonly passwordPolicy?: pulumi.Input<string>;
    /**
     * The name of the realm. This is unique across Keycloak. This will also be used as the realm's internal ID within Keycloak.
     */
    readonly realm: pulumi.Input<string>;
    /**
     * Maximum number of times a refresh token can be reused before they are revoked. If unspecified and 'revoke_refresh_token' is enabled the default value is 0 and refresh tokens can not be reused.
     */
    readonly refreshTokenMaxReuse?: pulumi.Input<number>;
    /**
     * When true, user registration will be enabled, and a link for registration will be displayed on the login page.
     */
    readonly registrationAllowed?: pulumi.Input<boolean>;
    /**
     * When true, the user's email will be used as their username during registration.
     */
    readonly registrationEmailAsUsername?: pulumi.Input<boolean>;
    /**
     * The desired flow for user registration. Defaults to `registration`.
     */
    readonly registrationFlow?: pulumi.Input<string>;
    /**
     * When true, a "remember me" checkbox will be displayed on the login page, and the user's session will not expire between browser restarts.
     */
    readonly rememberMe?: pulumi.Input<boolean>;
    /**
     * The desired flow to use when a user attempts to reset their credentials. Defaults to `reset credentials`.
     */
    readonly resetCredentialsFlow?: pulumi.Input<string>;
    /**
     * When true, a "forgot password" link will be displayed on the login page.
     */
    readonly resetPasswordAllowed?: pulumi.Input<boolean>;
    /**
     * If enabled a refresh token can only be used number of times specified in 'refresh_token_max_reuse' before they are revoked. If unspecified, refresh tokens can be reused.
     */
    readonly revokeRefreshToken?: pulumi.Input<boolean>;
    readonly securityDefenses?: pulumi.Input<inputs.RealmSecurityDefenses>;
    readonly smtpServer?: pulumi.Input<inputs.RealmSmtpServer>;
    /**
     * Can be one of following values: 'none, 'external' or 'all'
     */
    readonly sslRequired?: pulumi.Input<string>;
    /**
     * The amount of time a session can be idle before it expires.
     */
    readonly ssoSessionIdleTimeout?: pulumi.Input<string>;
    readonly ssoSessionIdleTimeoutRememberMe?: pulumi.Input<string>;
    /**
     * The maximum amount of time before a session expires regardless of activity.
     */
    readonly ssoSessionMaxLifespan?: pulumi.Input<string>;
    readonly ssoSessionMaxLifespanRememberMe?: pulumi.Input<string>;
    /**
     * When `true`, users are allowed to manage their own resources. Defaults to `false`.
     */
    readonly userManagedAccess?: pulumi.Input<boolean>;
    /**
     * When true, users are required to verify their email address after registration and after email address changes.
     */
    readonly verifyEmail?: pulumi.Input<boolean>;
    /**
     * Configuration for WebAuthn Passwordless Policy authentication.
     */
    readonly webAuthnPasswordlessPolicy?: pulumi.Input<inputs.RealmWebAuthnPasswordlessPolicy>;
    /**
     * Configuration for WebAuthn Policy authentication.
     */
    readonly webAuthnPolicy?: pulumi.Input<inputs.RealmWebAuthnPolicy>;
}
