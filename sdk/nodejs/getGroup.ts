// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can be used to fetch properties of a Keycloak group for
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
 * const group = keycloak.getGroupOutput({
 *     realmId: realm.id,
 *     name: "group",
 * });
 * const groupRoles = new keycloak.GroupRoles("groupRoles", {
 *     realmId: realm.id,
 *     groupId: group.apply(group => group.id),
 *     roleIds: [offlineAccess.apply(offlineAccess => offlineAccess.id)],
 * });
 * ```
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("keycloak:index/getGroup:getGroup", {
        "name": args.name,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The name of the group. If there are multiple groups match `name`, the first result will be returned.
     */
    name: string;
    /**
     * The realm this group exists within.
     */
    realmId: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    readonly attributes: {[key: string]: any};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly parentId: string;
    readonly path: string;
    readonly realmId: string;
}
/**
 * This data source can be used to fetch properties of a Keycloak group for
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
 * const group = keycloak.getGroupOutput({
 *     realmId: realm.id,
 *     name: "group",
 * });
 * const groupRoles = new keycloak.GroupRoles("groupRoles", {
 *     realmId: realm.id,
 *     groupId: group.apply(group => group.id),
 *     roleIds: [offlineAccess.apply(offlineAccess => offlineAccess.id)],
 * });
 * ```
 */
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply((a: any) => getGroup(a, opts))
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupOutputArgs {
    /**
     * The name of the group. If there are multiple groups match `name`, the first result will be returned.
     */
    name: pulumi.Input<string>;
    /**
     * The realm this group exists within.
     */
    realmId: pulumi.Input<string>;
}
