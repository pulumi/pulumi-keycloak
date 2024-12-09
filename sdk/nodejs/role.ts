// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing roles within Keycloak.
 *
 * Roles allow you define privileges within Keycloak and map them to users and groups.
 *
 * ## Example Usage
 *
 * ### Realm Role)
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
 *     attributes: {
 *         key: "value",
 *         multivalue: "value1##value2",
 *     },
 * });
 * ```
 *
 * ### Client Role)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const openidClient = new keycloak.openid.Client("openid_client", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     name: "client",
 *     enabled: true,
 *     accessType: "CONFIDENTIAL",
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const clientRole = new keycloak.Role("client_role", {
 *     realmId: realm.id,
 *     clientId: openidClientKeycloakClient.id,
 *     name: "my-client-role",
 *     description: "My Client Role",
 *     attributes: {
 *         key: "value",
 *     },
 * });
 * ```
 *
 * ### Composite Role)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * // realm roles
 * const createRole = new keycloak.Role("create_role", {
 *     realmId: realm.id,
 *     name: "create",
 *     attributes: {
 *         key: "value",
 *     },
 * });
 * const readRole = new keycloak.Role("read_role", {
 *     realmId: realm.id,
 *     name: "read",
 *     attributes: {
 *         key: "value",
 *     },
 * });
 * const updateRole = new keycloak.Role("update_role", {
 *     realmId: realm.id,
 *     name: "update",
 *     attributes: {
 *         key: "value",
 *     },
 * });
 * const deleteRole = new keycloak.Role("delete_role", {
 *     realmId: realm.id,
 *     name: "delete",
 *     attributes: {
 *         key: "value",
 *     },
 * });
 * // client role
 * const openidClient = new keycloak.openid.Client("openid_client", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     name: "client",
 *     enabled: true,
 *     accessType: "CONFIDENTIAL",
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const clientRole = new keycloak.Role("client_role", {
 *     realmId: realm.id,
 *     clientId: openidClientKeycloakClient.id,
 *     name: "my-client-role",
 *     description: "My Client Role",
 *     attributes: {
 *         key: "value",
 *     },
 * });
 * const adminRole = new keycloak.Role("admin_role", {
 *     realmId: realm.id,
 *     name: "admin",
 *     compositeRoles: [
 *         createRole.id,
 *         readRole.id,
 *         updateRole.id,
 *         deleteRole.id,
 *         clientRole.id,
 *     ],
 *     attributes: {
 *         key: "value",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Roles can be imported using the format `{{realm_id}}/{{role_id}}`, where `role_id` is the unique ID that Keycloak assigns
 *
 * to the role. The ID is not easy to find in the GUI, but it appears in the URL when editing the role.
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:index/role:Role role my-realm/7e8cf32a-8acb-4d34-89c4-04fb1d10ccad
 * ```
 */
export class Role extends pulumi.CustomResource {
    /**
     * Get an existing Role resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleState, opts?: pulumi.CustomResourceOptions): Role {
        return new Role(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/role:Role';

    /**
     * Returns true if the given object is an instance of Role.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Role {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Role.__pulumiType;
    }

    /**
     * A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     */
    public readonly attributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * When specified, this role will be created as a client role attached to the client with the provided ID
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
     */
    public readonly compositeRoles!: pulumi.Output<string[] | undefined>;
    /**
     * The description of the role
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the role
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm this role exists within.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a Role resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleArgs | RoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoleState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["compositeRoles"] = state ? state.compositeRoles : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as RoleArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["compositeRoles"] = args ? args.compositeRoles : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Role.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Role resources.
 */
export interface RoleState {
    /**
     * A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     */
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * When specified, this role will be created as a client role attached to the client with the provided ID
     */
    clientId?: pulumi.Input<string>;
    /**
     * When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
     */
    compositeRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the role
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the role
     */
    name?: pulumi.Input<string>;
    /**
     * The realm this role exists within.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Role resource.
 */
export interface RoleArgs {
    /**
     * A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     */
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * When specified, this role will be created as a client role attached to the client with the provided ID
     */
    clientId?: pulumi.Input<string>;
    /**
     * When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
     */
    compositeRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the role
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the role
     */
    name?: pulumi.Input<string>;
    /**
     * The realm this role exists within.
     */
    realmId: pulumi.Input<string>;
}
