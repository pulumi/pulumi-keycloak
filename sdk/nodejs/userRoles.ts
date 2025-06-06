// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows you to manage roles assigned to a Keycloak user.
 *
 * If `exhaustive` is true, this resource attempts to be an **authoritative** source over user roles: roles that are manually added to the user will be removed, and roles that are manually removed from the
 * user will be added upon the next run of `pulumi up`.
 * If `exhaustive` is false, this resource is a partial assignation of roles to a user. As a result, you can use multiple `keycloak.UserRoles` for the same `userId`.
 *
 * Note that when assigning composite roles to a user, you may see a non-empty plan following a `pulumi up` if you assign
 * a role and a composite that includes that role to the same user.
 *
 * ## Example Usage
 *
 * ### Exhaustive Roles)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const realmRole = new keycloak.Role("realm_role", {
 *     realmId: realm.id,
 *     name: "my-realm-role",
 *     description: "My Realm Role",
 * });
 * const client = new keycloak.openid.Client("client", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     name: "client",
 *     enabled: true,
 *     accessType: "BEARER-ONLY",
 * });
 * const clientRole = new keycloak.Role("client_role", {
 *     realmId: realm.id,
 *     clientId: clientKeycloakClient.id,
 *     name: "my-client-role",
 *     description: "My Client Role",
 * });
 * const user = new keycloak.User("user", {
 *     realmId: realm.id,
 *     username: "bob",
 *     enabled: true,
 *     email: "bob@domain.com",
 *     firstName: "Bob",
 *     lastName: "Bobson",
 * });
 * const userRoles = new keycloak.UserRoles("user_roles", {
 *     realmId: realm.id,
 *     userId: user.id,
 *     roleIds: [
 *         realmRole.id,
 *         clientRole.id,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak
 *
 * assigns to the user upon creation. This value can be found in the GUI when editing the user, and is typically a GUID.
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:index/userRoles:UserRoles user_roles my-realm/b0ae6924-1bd5-4655-9e38-dae7c5e42924
 * ```
 */
export class UserRoles extends pulumi.CustomResource {
    /**
     * Get an existing UserRoles resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserRolesState, opts?: pulumi.CustomResourceOptions): UserRoles {
        return new UserRoles(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/userRoles:UserRoles';

    /**
     * Returns true if the given object is an instance of UserRoles.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserRoles {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserRoles.__pulumiType;
    }

    /**
     * Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
     */
    public readonly exhaustive!: pulumi.Output<boolean | undefined>;
    /**
     * The realm this user exists in.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * A list of role IDs to map to the user
     */
    public readonly roleIds!: pulumi.Output<string[]>;
    /**
     * The ID of the user this resource should manage roles for.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a UserRoles resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserRolesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserRolesArgs | UserRolesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserRolesState | undefined;
            resourceInputs["exhaustive"] = state ? state.exhaustive : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["roleIds"] = state ? state.roleIds : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserRolesArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.roleIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleIds'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["exhaustive"] = args ? args.exhaustive : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["roleIds"] = args ? args.roleIds : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserRoles.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserRoles resources.
 */
export interface UserRolesState {
    /**
     * Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
     */
    exhaustive?: pulumi.Input<boolean>;
    /**
     * The realm this user exists in.
     */
    realmId?: pulumi.Input<string>;
    /**
     * A list of role IDs to map to the user
     */
    roleIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the user this resource should manage roles for.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserRoles resource.
 */
export interface UserRolesArgs {
    /**
     * Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
     */
    exhaustive?: pulumi.Input<boolean>;
    /**
     * The realm this user exists in.
     */
    realmId: pulumi.Input<string>;
    /**
     * A list of role IDs to map to the user
     */
    roleIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the user this resource should manage roles for.
     */
    userId: pulumi.Input<string>;
}
