// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can be used to fetch information about the service account user that is associated with an OpenID client
 * that has service accounts enabled.
 *
 * ## Example Usage
 *
 * In this example, we'll create an OpenID client with service accounts enabled. This causes Keycloak to create a special user
 * that represents the service account. We'll use this data source to grab this user's ID in order to assign some roles to this
 * user, using the `keycloak.UserRoles` resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const client = new keycloak.openid.Client("client", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     accessType: "CONFIDENTIAL",
 *     serviceAccountsEnabled: true,
 * });
 * const serviceAccountUser = keycloak.openid.getClientServiceAccountUserOutput({
 *     realmId: realm.id,
 *     clientId: client.id,
 * });
 * const offlineAccess = keycloak.getRoleOutput({
 *     realmId: realm.id,
 *     name: "offline_access",
 * });
 * const serviceAccountUserRoles = new keycloak.UserRoles("serviceAccountUserRoles", {
 *     realmId: realm.id,
 *     userId: serviceAccountUser.apply(serviceAccountUser => serviceAccountUser.id),
 *     roleIds: [offlineAccess.apply(offlineAccess => offlineAccess.id)],
 * });
 * ```
 */
export function getClientServiceAccountUser(args: GetClientServiceAccountUserArgs, opts?: pulumi.InvokeOptions): Promise<GetClientServiceAccountUserResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("keycloak:openid/getClientServiceAccountUser:getClientServiceAccountUser", {
        "clientId": args.clientId,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getClientServiceAccountUser.
 */
export interface GetClientServiceAccountUserArgs {
    /**
     * The ID of the OpenID client with service accounts enabled.
     */
    clientId: string;
    /**
     * The realm that the OpenID client exists within.
     */
    realmId: string;
}

/**
 * A collection of values returned by getClientServiceAccountUser.
 */
export interface GetClientServiceAccountUserResult {
    readonly attributes: {[key: string]: any};
    readonly clientId: string;
    readonly email: string;
    readonly emailVerified: boolean;
    readonly enabled: boolean;
    readonly federatedIdentities: outputs.openid.GetClientServiceAccountUserFederatedIdentity[];
    readonly firstName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lastName: string;
    readonly realmId: string;
    readonly requiredActions: string[];
    readonly username: string;
}
/**
 * This data source can be used to fetch information about the service account user that is associated with an OpenID client
 * that has service accounts enabled.
 *
 * ## Example Usage
 *
 * In this example, we'll create an OpenID client with service accounts enabled. This causes Keycloak to create a special user
 * that represents the service account. We'll use this data source to grab this user's ID in order to assign some roles to this
 * user, using the `keycloak.UserRoles` resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const client = new keycloak.openid.Client("client", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     accessType: "CONFIDENTIAL",
 *     serviceAccountsEnabled: true,
 * });
 * const serviceAccountUser = keycloak.openid.getClientServiceAccountUserOutput({
 *     realmId: realm.id,
 *     clientId: client.id,
 * });
 * const offlineAccess = keycloak.getRoleOutput({
 *     realmId: realm.id,
 *     name: "offline_access",
 * });
 * const serviceAccountUserRoles = new keycloak.UserRoles("serviceAccountUserRoles", {
 *     realmId: realm.id,
 *     userId: serviceAccountUser.apply(serviceAccountUser => serviceAccountUser.id),
 *     roleIds: [offlineAccess.apply(offlineAccess => offlineAccess.id)],
 * });
 * ```
 */
export function getClientServiceAccountUserOutput(args: GetClientServiceAccountUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClientServiceAccountUserResult> {
    return pulumi.output(args).apply((a: any) => getClientServiceAccountUser(a, opts))
}

/**
 * A collection of arguments for invoking getClientServiceAccountUser.
 */
export interface GetClientServiceAccountUserOutputArgs {
    /**
     * The ID of the OpenID client with service accounts enabled.
     */
    clientId: pulumi.Input<string>;
    /**
     * The realm that the OpenID client exists within.
     */
    realmId: pulumi.Input<string>;
}
