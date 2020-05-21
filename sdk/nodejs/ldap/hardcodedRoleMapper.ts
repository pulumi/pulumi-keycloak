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
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm that this LDAP mapper will exist in.
 * - `ldapUserFederationId` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
 * - `name` - (Required) Display name of this mapper when displayed in the console.
 * - `role` - (Required) The role which should be assigned to the users.
 */
export class HardcodedRoleMapper extends pulumi.CustomResource {
    /**
     * Get an existing HardcodedRoleMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HardcodedRoleMapperState | undefined;
            inputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as HardcodedRoleMapperArgs | undefined;
            if (!args || args.ldapUserFederationId === undefined) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["role"] = args ? args.role : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HardcodedRoleMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HardcodedRoleMapper resources.
 */
export interface HardcodedRoleMapperState {
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    readonly ldapUserFederationId?: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    readonly realmId?: pulumi.Input<string>;
    /**
     * Role to grant to user.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HardcodedRoleMapper resource.
 */
export interface HardcodedRoleMapperArgs {
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    readonly ldapUserFederationId: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    readonly realmId: pulumi.Input<string>;
    /**
     * Role to grant to user.
     */
    readonly role: pulumi.Input<string>;
}
