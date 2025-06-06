// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can be used to fetch properties of a Keycloak role for
 * usage with other resources, such as `keycloak.GroupRoles`.
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
 * const offlineAccess = keycloak.getRoleOutput({
 *     realmId: realm.id,
 *     name: "offline_access",
 * });
 * // use the data source
 * const group = new keycloak.Group("group", {
 *     realmId: realm.id,
 *     name: "group",
 * });
 * const groupRoles = new keycloak.GroupRoles("group_roles", {
 *     realmId: realm.id,
 *     groupId: group.id,
 *     roleIds: [offlineAccess.apply(offlineAccess => offlineAccess.id)],
 * });
 * ```
 */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("keycloak:index/getRole:getRole", {
        "clientId": args.clientId,
        "name": args.name,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleArgs {
    /**
     * When specified, this role is assumed to be a client role belonging to the client with the provided ID. The `id` attribute of a `keycloakClient` resource should be used here.
     */
    clientId?: string;
    /**
     * The name of the role.
     */
    name: string;
    /**
     * The realm this role exists within.
     */
    realmId: string;
}

/**
 * A collection of values returned by getRole.
 */
export interface GetRoleResult {
    readonly attributes: {[key: string]: string};
    readonly clientId?: string;
    readonly compositeRoles: string[];
    /**
     * (Computed) The description of the role.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly realmId: string;
}
/**
 * This data source can be used to fetch properties of a Keycloak role for
 * usage with other resources, such as `keycloak.GroupRoles`.
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
 * const offlineAccess = keycloak.getRoleOutput({
 *     realmId: realm.id,
 *     name: "offline_access",
 * });
 * // use the data source
 * const group = new keycloak.Group("group", {
 *     realmId: realm.id,
 *     name: "group",
 * });
 * const groupRoles = new keycloak.GroupRoles("group_roles", {
 *     realmId: realm.id,
 *     groupId: group.id,
 *     roleIds: [offlineAccess.apply(offlineAccess => offlineAccess.id)],
 * });
 * ```
 */
export function getRoleOutput(args: GetRoleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("keycloak:index/getRole:getRole", {
        "clientId": args.clientId,
        "name": args.name,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleOutputArgs {
    /**
     * When specified, this role is assumed to be a client role belonging to the client with the provided ID. The `id` attribute of a `keycloakClient` resource should be used here.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    name: pulumi.Input<string>;
    /**
     * The realm this role exists within.
     */
    realmId: pulumi.Input<string>;
}
