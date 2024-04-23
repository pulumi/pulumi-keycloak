// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing required actions within Keycloak.
 *
 * [Required actions](https://www.keycloak.org/docs/latest/server_admin/#con-required-actions_server_administration_guide) specify actions required before the first login of all new users.
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
 * const requiredAction = new keycloak.RequiredAction("required_action", {
 *     realmId: realm.realm,
 *     alias: "webauthn-register",
 *     enabled: true,
 *     name: "Webauthn Register",
 * });
 * ```
 *
 * ## Import
 *
 * Authentication executions can be imported using the formats: `{{realm}}/{{alias}}`.
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:index/requiredAction:RequiredAction required_action my-realm/my-default-action-alias
 * ```
 */
export class RequiredAction extends pulumi.CustomResource {
    /**
     * Get an existing RequiredAction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RequiredActionState, opts?: pulumi.CustomResourceOptions): RequiredAction {
        return new RequiredAction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/requiredAction:RequiredAction';

    /**
     * Returns true if the given object is an instance of RequiredAction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RequiredAction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RequiredAction.__pulumiType;
    }

    /**
     * The alias of the action to attach as a required action.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * When `true`, the required action is set as the default action for new users. Defaults to `false`.
     */
    public readonly defaultAction!: pulumi.Output<boolean | undefined>;
    /**
     * When `false`, the required action is not enabled for new users. Defaults to `false`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the required action.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The priority of the required action.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The realm the required action exists in.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a RequiredAction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RequiredActionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RequiredActionArgs | RequiredActionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RequiredActionState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["defaultAction"] = state ? state.defaultAction : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as RequiredActionArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["defaultAction"] = args ? args.defaultAction : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RequiredAction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RequiredAction resources.
 */
export interface RequiredActionState {
    /**
     * The alias of the action to attach as a required action.
     */
    alias?: pulumi.Input<string>;
    /**
     * When `true`, the required action is set as the default action for new users. Defaults to `false`.
     */
    defaultAction?: pulumi.Input<boolean>;
    /**
     * When `false`, the required action is not enabled for new users. Defaults to `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the required action.
     */
    name?: pulumi.Input<string>;
    /**
     * The priority of the required action.
     */
    priority?: pulumi.Input<number>;
    /**
     * The realm the required action exists in.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RequiredAction resource.
 */
export interface RequiredActionArgs {
    /**
     * The alias of the action to attach as a required action.
     */
    alias: pulumi.Input<string>;
    /**
     * When `true`, the required action is set as the default action for new users. Defaults to `false`.
     */
    defaultAction?: pulumi.Input<boolean>;
    /**
     * When `false`, the required action is not enabled for new users. Defaults to `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the required action.
     */
    name?: pulumi.Input<string>;
    /**
     * The priority of the required action.
     */
    priority?: pulumi.Input<number>;
    /**
     * The realm the required action exists in.
     */
    realmId: pulumi.Input<string>;
}
