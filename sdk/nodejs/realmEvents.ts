// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for managing Realm Events settings within Keycloak.
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
 * const realmEvents = new keycloak.RealmEvents("realm_events", {
 *     realmId: realm.id,
 *     eventsEnabled: true,
 *     eventsExpiration: 3600,
 *     adminEventsEnabled: true,
 *     adminEventsDetailsEnabled: true,
 *     enabledEventTypes: [
 *         "LOGIN",
 *         "LOGOUT",
 *     ],
 *     eventsListeners: ["jboss-logging"],
 * });
 * ```
 *
 * ## Import
 *
 * This resource currently does not support importing.
 */
export class RealmEvents extends pulumi.CustomResource {
    /**
     * Get an existing RealmEvents resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RealmEventsState, opts?: pulumi.CustomResourceOptions): RealmEvents {
        return new RealmEvents(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/realmEvents:RealmEvents';

    /**
     * Returns true if the given object is an instance of RealmEvents.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RealmEvents {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RealmEvents.__pulumiType;
    }

    /**
     * When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
     */
    public readonly adminEventsDetailsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
     */
    public readonly adminEventsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
     */
    public readonly enabledEventTypes!: pulumi.Output<string[] | undefined>;
    /**
     * When `true`, events from `enabledEventTypes` are saved to the database, making them available through the admin console. Defaults to `false`.
     */
    public readonly eventsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
     */
    public readonly eventsExpiration!: pulumi.Output<number | undefined>;
    /**
     * The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
     */
    public readonly eventsListeners!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the realm the event settings apply to.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a RealmEvents resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealmEventsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RealmEventsArgs | RealmEventsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RealmEventsState | undefined;
            resourceInputs["adminEventsDetailsEnabled"] = state ? state.adminEventsDetailsEnabled : undefined;
            resourceInputs["adminEventsEnabled"] = state ? state.adminEventsEnabled : undefined;
            resourceInputs["enabledEventTypes"] = state ? state.enabledEventTypes : undefined;
            resourceInputs["eventsEnabled"] = state ? state.eventsEnabled : undefined;
            resourceInputs["eventsExpiration"] = state ? state.eventsExpiration : undefined;
            resourceInputs["eventsListeners"] = state ? state.eventsListeners : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as RealmEventsArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["adminEventsDetailsEnabled"] = args ? args.adminEventsDetailsEnabled : undefined;
            resourceInputs["adminEventsEnabled"] = args ? args.adminEventsEnabled : undefined;
            resourceInputs["enabledEventTypes"] = args ? args.enabledEventTypes : undefined;
            resourceInputs["eventsEnabled"] = args ? args.eventsEnabled : undefined;
            resourceInputs["eventsExpiration"] = args ? args.eventsExpiration : undefined;
            resourceInputs["eventsListeners"] = args ? args.eventsListeners : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RealmEvents.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RealmEvents resources.
 */
export interface RealmEventsState {
    /**
     * When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
     */
    adminEventsDetailsEnabled?: pulumi.Input<boolean>;
    /**
     * When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
     */
    adminEventsEnabled?: pulumi.Input<boolean>;
    /**
     * The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
     */
    enabledEventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When `true`, events from `enabledEventTypes` are saved to the database, making them available through the admin console. Defaults to `false`.
     */
    eventsEnabled?: pulumi.Input<boolean>;
    /**
     * The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
     */
    eventsExpiration?: pulumi.Input<number>;
    /**
     * The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
     */
    eventsListeners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the realm the event settings apply to.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RealmEvents resource.
 */
export interface RealmEventsArgs {
    /**
     * When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
     */
    adminEventsDetailsEnabled?: pulumi.Input<boolean>;
    /**
     * When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
     */
    adminEventsEnabled?: pulumi.Input<boolean>;
    /**
     * The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
     */
    enabledEventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When `true`, events from `enabledEventTypes` are saved to the database, making them available through the admin console. Defaults to `false`.
     */
    eventsEnabled?: pulumi.Input<boolean>;
    /**
     * The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
     */
    eventsExpiration?: pulumi.Input<number>;
    /**
     * The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
     */
    eventsListeners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the realm the event settings apply to.
     */
    realmId: pulumi.Input<string>;
}
