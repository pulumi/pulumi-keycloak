// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * ### Exhaustive Groups)
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const group = new keycloak.Group("group", {realmId: realm.id});
 * const user = new keycloak.User("user", {
 *     realmId: realm.id,
 *     username: "my-user",
 * });
 * const userGroups = new keycloak.UserGroups("userGroups", {
 *     realmId: realm.id,
 *     userId: user.id,
 *     groupIds: [group.id],
 * });
 * ```
 * ### Non Exhaustive Groups)
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const groupFoo = new keycloak.Group("groupFoo", {realmId: realm.id});
 * const groupBar = new keycloak.Group("groupBar", {realmId: realm.id});
 * const user = new keycloak.User("user", {
 *     realmId: realm.id,
 *     username: "my-user",
 * });
 * const userGroupsAssociation1UserGroups = new keycloak.UserGroups("userGroupsAssociation1UserGroups", {
 *     realmId: realm.id,
 *     userId: user.id,
 *     exhaustive: false,
 *     groupIds: [groupFoo.id],
 * });
 * const userGroupsAssociation1Index_userGroupsUserGroups = new keycloak.UserGroups("userGroupsAssociation1Index/userGroupsUserGroups", {
 *     realmId: realm.id,
 *     userId: user.id,
 *     exhaustive: false,
 *     groupIds: [groupBar.id],
 * });
 * ```
 *
 * ## Import
 *
 * This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server.
 */
export class UserGroups extends pulumi.CustomResource {
    /**
     * Get an existing UserGroups resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserGroupsState, opts?: pulumi.CustomResourceOptions): UserGroups {
        return new UserGroups(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/userGroups:UserGroups';

    /**
     * Returns true if the given object is an instance of UserGroups.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserGroups {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserGroups.__pulumiType;
    }

    /**
     * Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
     */
    public readonly exhaustive!: pulumi.Output<boolean | undefined>;
    /**
     * A list of group IDs that the user is member of.
     */
    public readonly groupIds!: pulumi.Output<string[]>;
    /**
     * The realm this group exists in.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * The ID of the user this resource should manage groups for.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a UserGroups resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserGroupsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserGroupsArgs | UserGroupsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserGroupsState | undefined;
            resourceInputs["exhaustive"] = state ? state.exhaustive : undefined;
            resourceInputs["groupIds"] = state ? state.groupIds : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserGroupsArgs | undefined;
            if ((!args || args.groupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupIds'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["exhaustive"] = args ? args.exhaustive : undefined;
            resourceInputs["groupIds"] = args ? args.groupIds : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserGroups.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserGroups resources.
 */
export interface UserGroupsState {
    /**
     * Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
     */
    exhaustive?: pulumi.Input<boolean>;
    /**
     * A list of group IDs that the user is member of.
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The realm this group exists in.
     */
    realmId?: pulumi.Input<string>;
    /**
     * The ID of the user this resource should manage groups for.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserGroups resource.
 */
export interface UserGroupsArgs {
    /**
     * Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
     */
    exhaustive?: pulumi.Input<boolean>;
    /**
     * A list of group IDs that the user is member of.
     */
    groupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The realm this group exists in.
     */
    realmId: pulumi.Input<string>;
    /**
     * The ID of the user this resource should manage groups for.
     */
    userId: pulumi.Input<string>;
}
