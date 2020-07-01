// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # keycloak.Role data source
 *
 * This data source can be used to fetch properties of a Keycloak role for
 * usage with other resources, such as `keycloak.GroupRoles`.
 *
 * ### Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const offlineAccess = realm.id.apply(id => keycloak.getRole({
 *     name: "offline_access",
 *     realmId: id,
 * }, { async: true }));
 * const group = new keycloak.Group("group", {
 *     realmId: realm.id,
 * });
 * const groupRoles = new keycloak.GroupRoles("group_roles", {
 *     groupId: group.id,
 *     realmId: realm.id,
 *     roles: [offlineAccess.id],
 * });
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm this role exists within.
 * - `clientId` - (Optional) When specified, this role is assumed to be a
 *   client role belonging to the client with the provided ID
 * - `name` - (Required) The name of the role
 *
 * ### Attributes Reference
 *
 * In addition to the arguments listed above, the following computed attributes are exported:
 *
 * - `id` - The unique ID of the role, which can be used as an argument to
 *   other resources supported by this provider.
 * - `description` - The description of the role.
 */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
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
    readonly clientId?: string;
    readonly name: string;
    readonly realmId: string;
}

/**
 * A collection of values returned by getRole.
 */
export interface GetRoleResult {
    readonly clientId?: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly realmId: string;
}
