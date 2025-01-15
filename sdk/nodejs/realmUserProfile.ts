// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Allows for managing Realm User Profiles within Keycloak.
 *
 * A user profile defines a schema for representing user attributes and how they are managed within a realm.
 *
 * Information for Keycloak versions < 24:
 * The realm linked to the `keycloak.RealmUserProfile` resource must have the user profile feature enabled.
 * It can be done via the administration UI, or by setting the `userProfileEnabled` realm attribute to `true`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {realm: "my-realm"});
 * const userprofile = new keycloak.RealmUserProfile("userprofile", {
 *     realmId: myRealm.id,
 *     unmanagedAttributePolicy: "ENABLED",
 *     attributes: [
 *         {
 *             name: "field1",
 *             displayName: "Field 1",
 *             group: "group1",
 *             enabledWhenScopes: ["offline_access"],
 *             requiredForRoles: ["user"],
 *             requiredForScopes: ["offline_access"],
 *             permissions: {
 *                 views: [
 *                     "admin",
 *                     "user",
 *                 ],
 *                 edits: [
 *                     "admin",
 *                     "user",
 *                 ],
 *             },
 *             validators: [
 *                 {
 *                     name: "person-name-prohibited-characters",
 *                 },
 *                 {
 *                     name: "pattern",
 *                     config: {
 *                         pattern: "^[a-z]+$",
 *                         "error-message": "Nope",
 *                     },
 *                 },
 *             ],
 *             annotations: {
 *                 foo: "bar",
 *             },
 *         },
 *         {
 *             name: "field2",
 *             validators: [{
 *                 name: "options",
 *                 config: {
 *                     options: JSON.stringify(["opt1"]),
 *                 },
 *             }],
 *             annotations: {
 *                 foo: JSON.stringify({
 *                     key: "val",
 *                 }),
 *             },
 *         },
 *     ],
 *     groups: [
 *         {
 *             name: "group1",
 *             displayHeader: "Group 1",
 *             displayDescription: "A first group",
 *             annotations: {
 *                 foo: "bar",
 *                 foo2: JSON.stringify({
 *                     key: "val",
 *                 }),
 *             },
 *         },
 *         {
 *             name: "group2",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * This resource currently does not support importing.
 */
export class RealmUserProfile extends pulumi.CustomResource {
    /**
     * Get an existing RealmUserProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RealmUserProfileState, opts?: pulumi.CustomResourceOptions): RealmUserProfile {
        return new RealmUserProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/realmUserProfile:RealmUserProfile';

    /**
     * Returns true if the given object is an instance of RealmUserProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RealmUserProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RealmUserProfile.__pulumiType;
    }

    /**
     * An ordered list of attributes.
     */
    public readonly attributes!: pulumi.Output<outputs.RealmUserProfileAttribute[] | undefined>;
    /**
     * A list of groups.
     */
    public readonly groups!: pulumi.Output<outputs.RealmUserProfileGroup[] | undefined>;
    /**
     * The ID of the realm the user profile applies to.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
     */
    public readonly unmanagedAttributePolicy!: pulumi.Output<string | undefined>;

    /**
     * Create a RealmUserProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealmUserProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RealmUserProfileArgs | RealmUserProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RealmUserProfileState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["unmanagedAttributePolicy"] = state ? state.unmanagedAttributePolicy : undefined;
        } else {
            const args = argsOrState as RealmUserProfileArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["unmanagedAttributePolicy"] = args ? args.unmanagedAttributePolicy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RealmUserProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RealmUserProfile resources.
 */
export interface RealmUserProfileState {
    /**
     * An ordered list of attributes.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.RealmUserProfileAttribute>[]>;
    /**
     * A list of groups.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.RealmUserProfileGroup>[]>;
    /**
     * The ID of the realm the user profile applies to.
     */
    realmId?: pulumi.Input<string>;
    /**
     * Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
     */
    unmanagedAttributePolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RealmUserProfile resource.
 */
export interface RealmUserProfileArgs {
    /**
     * An ordered list of attributes.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.RealmUserProfileAttribute>[]>;
    /**
     * A list of groups.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.RealmUserProfileGroup>[]>;
    /**
     * The ID of the realm the user profile applies to.
     */
    realmId: pulumi.Input<string>;
    /**
     * Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
     */
    unmanagedAttributePolicy?: pulumi.Input<string>;
}
