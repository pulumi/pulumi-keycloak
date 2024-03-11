// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # keycloak.ldap.HardcodedRoleMapper
 *
 * This mapper will grant a specified Keycloak role to each Keycloak user linked with LDAP.
 *
 * ### Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "test",
 *     enabled: true,
 * });
 * const ldapUserFederation = new keycloak.ldap.UserFederation("ldapUserFederation", {
 *     realmId: realm.id,
 *     usernameLdapAttribute: "cn",
 *     rdnLdapAttribute: "cn",
 *     uuidLdapAttribute: "entryDN",
 *     userObjectClasses: [
 *         "simpleSecurityObject",
 *         "organizationalRole",
 *     ],
 *     connectionUrl: "ldap://openldap",
 *     usersDn: "dc=example,dc=org",
 *     bindDn: "cn=admin,dc=example,dc=org",
 *     bindCredential: "admin",
 * });
 * const assignAdminRoleToAllUsers = new keycloak.ldap.HardcodedRoleMapper("assignAdminRoleToAllUsers", {
 *     realmId: realm.id,
 *     ldapUserFederationId: ldapUserFederation.id,
 *     role: "admin",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm that this LDAP mapper will exist in.
 * - `ldapUserFederationId` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
 * - `name` - (Required) Display name of this mapper when displayed in the console.
 * - `role` - (Required) The role which should be assigned to the users.
 *
 * ### Import
 *
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
 * The ID of the LDAP user federation provider and the mapper can be found within
 * the Keycloak GUI, and they are typically GUIDs:
 */
export class HardcodedRoleMapper extends pulumi.CustomResource {
    /**
     * Get an existing HardcodedRoleMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HardcodedRoleMapperState, opts?: pulumi.CustomResourceOptions): HardcodedRoleMapper {
        return new HardcodedRoleMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper';

    /**
     * Returns true if the given object is an instance of HardcodedRoleMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HardcodedRoleMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HardcodedRoleMapper.__pulumiType;
    }

    /**
     * The ldap user federation provider to attach this mapper to.
     */
    public readonly ldapUserFederationId!: pulumi.Output<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * Role to grant to user.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a HardcodedRoleMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HardcodedRoleMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HardcodedRoleMapperArgs | HardcodedRoleMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HardcodedRoleMapperState | undefined;
            resourceInputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as HardcodedRoleMapperArgs | undefined;
            if ((!args || args.ldapUserFederationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HardcodedRoleMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HardcodedRoleMapper resources.
 */
export interface HardcodedRoleMapperState {
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    ldapUserFederationId?: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    realmId?: pulumi.Input<string>;
    /**
     * Role to grant to user.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HardcodedRoleMapper resource.
 */
export interface HardcodedRoleMapperArgs {
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    ldapUserFederationId: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    realmId: pulumi.Input<string>;
    /**
     * Role to grant to user.
     */
    role: pulumi.Input<string>;
}
