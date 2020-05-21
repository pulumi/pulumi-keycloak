// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # keycloak..Group
 *
 * Allows for creating and managing Groups within Keycloak.
 *
 * Groups provide a logical wrapping for users within Keycloak. Users within a
 * group can share attributes and roles, and group membership can be mapped
 * to a claim.
 *
 * Attributes can also be defined on Groups.
 *
 * Groups can also be federated from external data sources, such as LDAP or Active Directory.
 * This resource **should not** be used to manage groups that were created this way.
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
 * const parentGroup = new keycloak.Group("parentGroup", {
 *     realmId: realm.id,
 * });
 * const childGroup = new keycloak.Group("childGroup", {
 *     parentId: parentGroup.id,
 *     realmId: realm.id,
 * });
 * const childGroupWithOptionalAttributes = new keycloak.Group("childGroupWithOptionalAttributes", {
 *     attributes: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 *     parentId: parentGroup.id,
 *     realmId: realm.id,
 * });
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm this group exists in.
 * - `parentId` - (Optional) The ID of this group's parent. If omitted, this group will be defined at the root level.
 * - `name` - (Required) The name of the group.
 * - `attributes` - (Optional) A dict of key/value pairs to set as custom attributes for the group.
 *
 * ### Attributes Reference
 *
 * In addition to the arguments listed above, the following computed attributes are exported:
 *
 * - `path` - The complete path of the group. For example, the child group's path in the example configuration would be `/parent-group/child-group`.
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    public readonly attributes!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly parentId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly path!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupState | undefined;
            inputs["attributes"] = state ? state.attributes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parentId"] = state ? state.parentId : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parentId"] = args ? args.parentId : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["path"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Group.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    readonly name?: pulumi.Input<string>;
    readonly parentId?: pulumi.Input<string>;
    readonly path?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    readonly name?: pulumi.Input<string>;
    readonly parentId?: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
}
