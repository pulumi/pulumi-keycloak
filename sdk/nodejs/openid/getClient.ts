// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## # keycloak.openid.Client data source
 *
 * This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
 *
 * ### Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realmManagement = keycloak.openid.getClient({
 *     realmId: "my-realm",
 *     clientId: "realm-management",
 * });
 * // use the data source
 * const admin = realmManagement.then(realmManagement => keycloak.getRole({
 *     realmId: "my-realm",
 *     clientId: realmManagement.id,
 *     name: "realm-admin",
 * }));
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm id.
 * - `clientId` - (Required) The client id.
 *
 * ### Attributes Reference
 *
 * See the docs for the `keycloak.openid.Client` resource for details on the exported attributes.
 */
export function getClient(args: GetClientArgs, opts?: pulumi.InvokeOptions): Promise<GetClientResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("keycloak:openid/getClient:getClient", {
        "clientId": args.clientId,
        "consentScreenText": args.consentScreenText,
        "displayOnConsentScreen": args.displayOnConsentScreen,
        "extraConfig": args.extraConfig,
        "oauth2DeviceAuthorizationGrantEnabled": args.oauth2DeviceAuthorizationGrantEnabled,
        "oauth2DeviceCodeLifespan": args.oauth2DeviceCodeLifespan,
        "oauth2DevicePollingInterval": args.oauth2DevicePollingInterval,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getClient.
 */
export interface GetClientArgs {
    clientId: string;
    consentScreenText?: string;
    displayOnConsentScreen?: boolean;
    extraConfig?: {[key: string]: string};
    oauth2DeviceAuthorizationGrantEnabled?: boolean;
    oauth2DeviceCodeLifespan?: string;
    oauth2DevicePollingInterval?: string;
    realmId: string;
}

/**
 * A collection of values returned by getClient.
 */
export interface GetClientResult {
    readonly accessTokenLifespan: string;
    readonly accessType: string;
    readonly adminUrl: string;
    readonly authenticationFlowBindingOverrides: outputs.openid.GetClientAuthenticationFlowBindingOverride[];
    readonly authorizations: outputs.openid.GetClientAuthorization[];
    readonly backchannelLogoutRevokeOfflineSessions: boolean;
    readonly backchannelLogoutSessionRequired: boolean;
    readonly backchannelLogoutUrl: string;
    readonly baseUrl: string;
    readonly clientAuthenticatorType: string;
    readonly clientId: string;
    readonly clientOfflineSessionIdleTimeout: string;
    readonly clientOfflineSessionMaxLifespan: string;
    readonly clientSecret: string;
    readonly clientSessionIdleTimeout: string;
    readonly clientSessionMaxLifespan: string;
    readonly consentRequired: boolean;
    readonly consentScreenText?: string;
    readonly description: string;
    readonly directAccessGrantsEnabled: boolean;
    readonly displayOnConsentScreen?: boolean;
    readonly enabled: boolean;
    readonly excludeSessionStateFromAuthResponse: boolean;
    readonly extraConfig: {[key: string]: string};
    readonly frontchannelLogoutEnabled: boolean;
    readonly frontchannelLogoutUrl: string;
    readonly fullScopeAllowed: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly implicitFlowEnabled: boolean;
    readonly loginTheme: string;
    readonly name: string;
    readonly oauth2DeviceAuthorizationGrantEnabled?: boolean;
    readonly oauth2DeviceCodeLifespan?: string;
    readonly oauth2DevicePollingInterval?: string;
    readonly pkceCodeChallengeMethod: string;
    readonly realmId: string;
    readonly resourceServerId: string;
    readonly rootUrl: string;
    readonly serviceAccountUserId: string;
    readonly serviceAccountsEnabled: boolean;
    readonly standardFlowEnabled: boolean;
    readonly useRefreshTokens: boolean;
    readonly useRefreshTokensClientCredentials: boolean;
    readonly validPostLogoutRedirectUris: string[];
    readonly validRedirectUris: string[];
    readonly webOrigins: string[];
}
/**
 * ## # keycloak.openid.Client data source
 *
 * This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
 *
 * ### Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realmManagement = keycloak.openid.getClient({
 *     realmId: "my-realm",
 *     clientId: "realm-management",
 * });
 * // use the data source
 * const admin = realmManagement.then(realmManagement => keycloak.getRole({
 *     realmId: "my-realm",
 *     clientId: realmManagement.id,
 *     name: "realm-admin",
 * }));
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm id.
 * - `clientId` - (Required) The client id.
 *
 * ### Attributes Reference
 *
 * See the docs for the `keycloak.openid.Client` resource for details on the exported attributes.
 */
export function getClientOutput(args: GetClientOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClientResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("keycloak:openid/getClient:getClient", {
        "clientId": args.clientId,
        "consentScreenText": args.consentScreenText,
        "displayOnConsentScreen": args.displayOnConsentScreen,
        "extraConfig": args.extraConfig,
        "oauth2DeviceAuthorizationGrantEnabled": args.oauth2DeviceAuthorizationGrantEnabled,
        "oauth2DeviceCodeLifespan": args.oauth2DeviceCodeLifespan,
        "oauth2DevicePollingInterval": args.oauth2DevicePollingInterval,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getClient.
 */
export interface GetClientOutputArgs {
    clientId: pulumi.Input<string>;
    consentScreenText?: pulumi.Input<string>;
    displayOnConsentScreen?: pulumi.Input<boolean>;
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    oauth2DeviceAuthorizationGrantEnabled?: pulumi.Input<boolean>;
    oauth2DeviceCodeLifespan?: pulumi.Input<string>;
    oauth2DevicePollingInterval?: pulumi.Input<string>;
    realmId: pulumi.Input<string>;
}
