// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # keycloak.openid.HardcodedRoleProtocolMapper
 *
 * Allows for creating and managing hardcoded role protocol mappers within
 * Keycloak.
 *
 * Hardcoded role protocol mappers allow you to specify a single role to
 * always map to an access token for a client. Protocol mappers can be
 * defined for a single client, or they can be defined for a client scope
 * which can be shared between multiple different clients.
 *
 * ### Example Usage (Client)
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const role = new keycloak.Role("role", {
 *     realmId: realm.id,
 *     name: "my-role",
 * });
 * const openidClient = new keycloak.openid.Client("openid_client", {
 *     realmId: realm.id,
 *     clientId: "test-client",
 *     name: "test client",
 *     enabled: true,
 *     accessType: "CONFIDENTIAL",
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const hardcodedRoleMapper = new keycloak.openid.HardcodedRoleProtocolMapper("hardcoded_role_mapper", {
 *     realmId: realm.id,
 *     clientId: openidClient.id,
 *     name: "hardcoded-role-mapper",
 *     roleId: role.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Example Usage (Client Scope)
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const role = new keycloak.Role("role", {
 *     realmId: realm.id,
 *     name: "my-role",
 * });
 * const clientScope = new keycloak.openid.ClientScope("client_scope", {
 *     realmId: realm.id,
 *     name: "test-client-scope",
 * });
 * const hardcodedRoleMapper = new keycloak.openid.HardcodedRoleProtocolMapper("hardcoded_role_mapper", {
 *     realmId: realm.id,
 *     clientScopeId: clientScope.id,
 *     name: "hardcoded-role-mapper",
 *     roleId: role.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm this protocol mapper exists within.
 * - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
 * - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the
 *   GUI.
 * - `roleId` - (Required) The ID of the role to map to an access token.
 *
 * ### Import
 *
 * Protocol mappers can be imported using one of the following formats:
 * - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
 * - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
 *
 * Example:
 */
export class HardcodedRoleProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing HardcodedRoleProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HardcodedRoleProtocolMapperState, opts?: pulumi.CustomResourceOptions): HardcodedRoleProtocolMapper {
        return new HardcodedRoleProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper';

    /**
     * Returns true if the given object is an instance of HardcodedRoleProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HardcodedRoleProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HardcodedRoleProtocolMapper.__pulumiType;
    }

    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    public readonly realmId!: pulumi.Output<string>;
    public readonly roleId!: pulumi.Output<string>;

    /**
     * Create a HardcodedRoleProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HardcodedRoleProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HardcodedRoleProtocolMapperArgs | HardcodedRoleProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HardcodedRoleProtocolMapperState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as HardcodedRoleProtocolMapperArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HardcodedRoleProtocolMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HardcodedRoleProtocolMapper resources.
 */
export interface HardcodedRoleProtocolMapperState {
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    realmId?: pulumi.Input<string>;
    roleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HardcodedRoleProtocolMapper resource.
 */
export interface HardcodedRoleProtocolMapperArgs {
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    realmId: pulumi.Input<string>;
    roleId: pulumi.Input<string>;
}
