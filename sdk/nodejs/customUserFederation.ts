// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing custom user federation providers within Keycloak.
 *
 * A custom user federation provider is an implementation of Keycloak's [User Storage SPI](https://www.keycloak.org/docs/4.2/server_development/index.html#_user-storage-spi).
 * An example of this implementation can be found here.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "test",
 *     enabled: true,
 * });
 * const customUserFederation = new keycloak.CustomUserFederation("custom_user_federation", {
 *     name: "custom",
 *     realmId: realm.id,
 *     providerId: "custom",
 *     enabled: true,
 *     config: {
 *         dummyString: "foobar",
 *         dummyBool: "true",
 *         multivalue: "value1##value2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Custom user federation providers can be imported using the format `{{realm_id}}/{{custom_user_federation_id}}`.
 *
 * The ID of the custom user federation provider can be found within the Keycloak GUI and is typically a GUID:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:index/customUserFederation:CustomUserFederation custom_user_federation my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860
 * ```
 */
export class CustomUserFederation extends pulumi.CustomResource {
    /**
     * Get an existing CustomUserFederation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomUserFederationState, opts?: pulumi.CustomResourceOptions): CustomUserFederation {
        return new CustomUserFederation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/customUserFederation:CustomUserFederation';

    /**
     * Returns true if the given object is an instance of CustomUserFederation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomUserFederation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomUserFederation.__pulumiType;
    }

    /**
     * Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     */
    public readonly cachePolicy!: pulumi.Output<string | undefined>;
    /**
     * How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
     */
    public readonly changedSyncPeriod!: pulumi.Output<number | undefined>;
    /**
     * The provider configuration handed over to your custom user federation provider. In order to add multivalued settings, use `##` to separate the values.
     */
    public readonly config!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
     */
    public readonly fullSyncPeriod!: pulumi.Output<number | undefined>;
    /**
     * Display name of the provider when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
     */
    public readonly parentId!: pulumi.Output<string>;
    /**
     * Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
     */
    public readonly providerId!: pulumi.Output<string>;
    /**
     * The realm that this provider will provide user federation for.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a CustomUserFederation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomUserFederationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomUserFederationArgs | CustomUserFederationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomUserFederationState | undefined;
            resourceInputs["cachePolicy"] = state ? state.cachePolicy : undefined;
            resourceInputs["changedSyncPeriod"] = state ? state.changedSyncPeriod : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["fullSyncPeriod"] = state ? state.fullSyncPeriod : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["providerId"] = state ? state.providerId : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as CustomUserFederationArgs | undefined;
            if ((!args || args.providerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["cachePolicy"] = args ? args.cachePolicy : undefined;
            resourceInputs["changedSyncPeriod"] = args ? args.changedSyncPeriod : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["fullSyncPeriod"] = args ? args.fullSyncPeriod : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["providerId"] = args ? args.providerId : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomUserFederation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomUserFederation resources.
 */
export interface CustomUserFederationState {
    /**
     * Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     */
    cachePolicy?: pulumi.Input<string>;
    /**
     * How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
     */
    changedSyncPeriod?: pulumi.Input<number>;
    /**
     * The provider configuration handed over to your custom user federation provider. In order to add multivalued settings, use `##` to separate the values.
     */
    config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
     */
    fullSyncPeriod?: pulumi.Input<number>;
    /**
     * Display name of the provider when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
     */
    parentId?: pulumi.Input<string>;
    /**
     * Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
     */
    priority?: pulumi.Input<number>;
    /**
     * The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
     */
    providerId?: pulumi.Input<string>;
    /**
     * The realm that this provider will provide user federation for.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomUserFederation resource.
 */
export interface CustomUserFederationArgs {
    /**
     * Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     */
    cachePolicy?: pulumi.Input<string>;
    /**
     * How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
     */
    changedSyncPeriod?: pulumi.Input<number>;
    /**
     * The provider configuration handed over to your custom user federation provider. In order to add multivalued settings, use `##` to separate the values.
     */
    config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
     */
    fullSyncPeriod?: pulumi.Input<number>;
    /**
     * Display name of the provider when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
     */
    parentId?: pulumi.Input<string>;
    /**
     * Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
     */
    priority?: pulumi.Input<number>;
    /**
     * The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
     */
    providerId: pulumi.Input<string>;
    /**
     * The realm that this provider will provide user federation for.
     */
    realmId: pulumi.Input<string>;
}
