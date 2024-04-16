// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * !> **WARNING:** This resource is deprecated and will be removed in the next major version. Please use `keycloak.GenericRoleMapper` instead.
 *
 * Allow for creating and managing a client's scope mappings within Keycloak.
 *
 * By default, all the user role mappings of the user are added as claims within the token (OIDC) or assertion (SAML). When
 * `fullScopeAllowed` is set to `false` for a client, role scope mapping allows you to limit the roles that get declared
 * inside an access token for a client.
 *
 * ## Example Usage
 *
 * ### Realm Role To Client)
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
 * const client = new keycloak.openid.Client("client", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     name: "client",
 *     enabled: true,
 *     accessType: "BEARER-ONLY",
 * });
 * const realmRole = new keycloak.Role("realm_role", {
 *     realmId: realm.id,
 *     name: "my-realm-role",
 *     description: "My Realm Role",
 * });
 * const clientRoleMapper = new keycloak.GenericClientRoleMapper("client_role_mapper", {
 *     realmId: realm.id,
 *     clientId: client.id,
 *     roleId: realmRole.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Client Role To Client)
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
 * const clientA = new keycloak.openid.Client("client_a", {
 *     realmId: realm.id,
 *     clientId: "client-a",
 *     name: "client-a",
 *     enabled: true,
 *     accessType: "BEARER-ONLY",
 *     fullScopeAllowed: false,
 * });
 * const clientRoleA = new keycloak.Role("client_role_a", {
 *     realmId: realm.id,
 *     clientId: clientA.id,
 *     name: "my-client-role",
 *     description: "My Client Role",
 * });
 * const clientB = new keycloak.openid.Client("client_b", {
 *     realmId: realm.id,
 *     clientId: "client-b",
 *     name: "client-b",
 *     enabled: true,
 *     accessType: "BEARER-ONLY",
 * });
 * const clientRoleB = new keycloak.Role("client_role_b", {
 *     realmId: realm.id,
 *     clientId: clientB.id,
 *     name: "my-client-role",
 *     description: "My Client Role",
 * });
 * const clientBRoleMapper = new keycloak.GenericClientRoleMapper("client_b_role_mapper", {
 *     realmId: realm.id,
 *     clientId: clientB.id,
 *     roleId: clientRoleA.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Realm Role To Client Scope)
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
 * const clientScope = new keycloak.openid.ClientScope("client_scope", {
 *     realmId: realm.id,
 *     name: "my-client-scope",
 * });
 * const realmRole = new keycloak.Role("realm_role", {
 *     realmId: realm.id,
 *     name: "my-realm-role",
 *     description: "My Realm Role",
 * });
 * const clientRoleMapper = new keycloak.GenericClientRoleMapper("client_role_mapper", {
 *     realmId: realm.id,
 *     clientScopeId: clientScope.id,
 *     roleId: realmRole.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Client Role To Client Scope)
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
 * const client = new keycloak.openid.Client("client", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     name: "client",
 *     enabled: true,
 *     accessType: "BEARER-ONLY",
 * });
 * const clientRole = new keycloak.Role("client_role", {
 *     realmId: realm.id,
 *     clientId: client.id,
 *     name: "my-client-role",
 *     description: "My Client Role",
 * });
 * const clientScope = new keycloak.openid.ClientScope("client_scope", {
 *     realmId: realm.id,
 *     name: "my-client-scope",
 * });
 * const clientBRoleMapper = new keycloak.GenericClientRoleMapper("client_b_role_mapper", {
 *     realmId: realm.id,
 *     clientScopeId: clientScope.id,
 *     roleId: clientRole.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Generic client role mappers can be imported using one of the following two formats:
 *
 * - When mapping a role to a client, use the format `{{realmId}}/client/{{clientId}}/scope-mappings/{{roleClientId}}/{{roleId}}`
 *
 * - When mapping a role to a client scope, use the format `{{realmId}}/client-scope/{{clientScopeId}}/scope-mappings/{{roleClientId}}/{{roleId}}`
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:index/genericClientRoleMapper:GenericClientRoleMapper client_role_mapper my-realm/client/23888550-5dcd-41f6-85ba-554233021e9c/scope-mappings/ce51f004-bdfb-4dd5-a963-c4487d2dec5b/ff3aa49f-bc07-4030-8783-41918c3614a3
 * ```
 */
export class GenericClientRoleMapper extends pulumi.CustomResource {
    /**
     * Get an existing GenericClientRoleMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GenericClientRoleMapperState, opts?: pulumi.CustomResourceOptions): GenericClientRoleMapper {
        return new GenericClientRoleMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/genericClientRoleMapper:GenericClientRoleMapper';

    /**
     * Returns true if the given object is an instance of GenericClientRoleMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GenericClientRoleMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GenericClientRoleMapper.__pulumiType;
    }

    /**
     * The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * The realm this role mapper exists within.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * The ID of the role to be added to this role mapper.
     */
    public readonly roleId!: pulumi.Output<string>;

    /**
     * Create a GenericClientRoleMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GenericClientRoleMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GenericClientRoleMapperArgs | GenericClientRoleMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GenericClientRoleMapperState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as GenericClientRoleMapperArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GenericClientRoleMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GenericClientRoleMapper resources.
 */
export interface GenericClientRoleMapperState {
    /**
     * The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * The realm this role mapper exists within.
     */
    realmId?: pulumi.Input<string>;
    /**
     * The ID of the role to be added to this role mapper.
     */
    roleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GenericClientRoleMapper resource.
 */
export interface GenericClientRoleMapperArgs {
    /**
     * The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * The realm this role mapper exists within.
     */
    realmId: pulumi.Input<string>;
    /**
     * The ID of the role to be added to this role mapper.
     */
    roleId: pulumi.Input<string>;
}
