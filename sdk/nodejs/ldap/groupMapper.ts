// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # keycloak.ldap.GroupMapper
 *
 * Allows for creating and managing group mappers for Keycloak users federated
 * via LDAP.
 *
 * The LDAP group mapper can be used to map an LDAP user's groups from some DN
 * to Keycloak groups. This group mapper will also create the groups within Keycloak
 * if they do not already exist.
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
 * const ldapUserFederation = new keycloak.ldap.UserFederation("ldap_user_federation", {
 *     name: "openldap",
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
 * const ldapGroupMapper = new keycloak.ldap.GroupMapper("ldap_group_mapper", {
 *     realmId: realm.id,
 *     ldapUserFederationId: ldapUserFederation.id,
 *     name: "group-mapper",
 *     ldapGroupsDn: "dc=example,dc=org",
 *     groupNameLdapAttribute: "cn",
 *     groupObjectClasses: ["groupOfNames"],
 *     membershipAttributeType: "DN",
 *     membershipLdapAttribute: "member",
 *     membershipUserLdapAttribute: "cn",
 *     memberofLdapAttribute: "memberOf",
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
 * - `ldapGroupsDn` - (Required) The LDAP DN where groups can be found.
 * - `groupNameLdapAttribute` - (Required) The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
 * - `groupObjectClasses` - (Required) Array of strings representing the object classes for the group. Must contain at least one.
 * - `preserveGroupInheritance` - (Optional) When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
 * - `ignoreMissingGroups` - (Optional) When `true`, missing groups in the hierarchy will be ignored.
 * - `membershipLdapAttribute` - (Required) The name of the LDAP attribute that is used for membership mappings.
 * - `membershipAttributeType` - (Optional) Can be one of `DN` or `UID`. Defaults to `DN`.
 * - `membershipUserLdapAttribute` - (Required) The name of the LDAP attribute on a user that is used for membership mappings.
 * - `groupsLdapFilter` - (Optional) When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
 * - `mode` - (Optional) Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
 * - `userRolesRetrieveStrategy` - (Optional) Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
 * - `memberofLdapAttribute` - (Optional) Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
 * - `mappedGroupAttributes` - (Optional) Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
 * - `dropNonExistingGroupsDuringSync` - (Optional) When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
 *
 * ### Import
 *
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
 * The ID of the LDAP user federation provider and the mapper can be found within
 * the Keycloak GUI, and they are typically GUIDs:
 */
export class GroupMapper extends pulumi.CustomResource {
    /**
     * Get an existing GroupMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMapperState, opts?: pulumi.CustomResourceOptions): GroupMapper {
        return new GroupMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:ldap/groupMapper:GroupMapper';

    /**
     * Returns true if the given object is an instance of GroupMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMapper.__pulumiType;
    }

    public readonly dropNonExistingGroupsDuringSync!: pulumi.Output<boolean | undefined>;
    public readonly groupNameLdapAttribute!: pulumi.Output<string>;
    public readonly groupObjectClasses!: pulumi.Output<string[]>;
    public readonly groupsLdapFilter!: pulumi.Output<string | undefined>;
    public readonly groupsPath!: pulumi.Output<string>;
    public readonly ignoreMissingGroups!: pulumi.Output<boolean | undefined>;
    public readonly ldapGroupsDn!: pulumi.Output<string>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    public readonly ldapUserFederationId!: pulumi.Output<string>;
    public readonly mappedGroupAttributes!: pulumi.Output<string[] | undefined>;
    public readonly memberofLdapAttribute!: pulumi.Output<string | undefined>;
    public readonly membershipAttributeType!: pulumi.Output<string | undefined>;
    public readonly membershipLdapAttribute!: pulumi.Output<string>;
    public readonly membershipUserLdapAttribute!: pulumi.Output<string>;
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly preserveGroupInheritance!: pulumi.Output<boolean | undefined>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    public readonly realmId!: pulumi.Output<string>;
    public readonly userRolesRetrieveStrategy!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMapperArgs | GroupMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMapperState | undefined;
            resourceInputs["dropNonExistingGroupsDuringSync"] = state ? state.dropNonExistingGroupsDuringSync : undefined;
            resourceInputs["groupNameLdapAttribute"] = state ? state.groupNameLdapAttribute : undefined;
            resourceInputs["groupObjectClasses"] = state ? state.groupObjectClasses : undefined;
            resourceInputs["groupsLdapFilter"] = state ? state.groupsLdapFilter : undefined;
            resourceInputs["groupsPath"] = state ? state.groupsPath : undefined;
            resourceInputs["ignoreMissingGroups"] = state ? state.ignoreMissingGroups : undefined;
            resourceInputs["ldapGroupsDn"] = state ? state.ldapGroupsDn : undefined;
            resourceInputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            resourceInputs["mappedGroupAttributes"] = state ? state.mappedGroupAttributes : undefined;
            resourceInputs["memberofLdapAttribute"] = state ? state.memberofLdapAttribute : undefined;
            resourceInputs["membershipAttributeType"] = state ? state.membershipAttributeType : undefined;
            resourceInputs["membershipLdapAttribute"] = state ? state.membershipLdapAttribute : undefined;
            resourceInputs["membershipUserLdapAttribute"] = state ? state.membershipUserLdapAttribute : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["preserveGroupInheritance"] = state ? state.preserveGroupInheritance : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["userRolesRetrieveStrategy"] = state ? state.userRolesRetrieveStrategy : undefined;
        } else {
            const args = argsOrState as GroupMapperArgs | undefined;
            if ((!args || args.groupNameLdapAttribute === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupNameLdapAttribute'");
            }
            if ((!args || args.groupObjectClasses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupObjectClasses'");
            }
            if ((!args || args.ldapGroupsDn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapGroupsDn'");
            }
            if ((!args || args.ldapUserFederationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if ((!args || args.membershipLdapAttribute === undefined) && !opts.urn) {
                throw new Error("Missing required property 'membershipLdapAttribute'");
            }
            if ((!args || args.membershipUserLdapAttribute === undefined) && !opts.urn) {
                throw new Error("Missing required property 'membershipUserLdapAttribute'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["dropNonExistingGroupsDuringSync"] = args ? args.dropNonExistingGroupsDuringSync : undefined;
            resourceInputs["groupNameLdapAttribute"] = args ? args.groupNameLdapAttribute : undefined;
            resourceInputs["groupObjectClasses"] = args ? args.groupObjectClasses : undefined;
            resourceInputs["groupsLdapFilter"] = args ? args.groupsLdapFilter : undefined;
            resourceInputs["groupsPath"] = args ? args.groupsPath : undefined;
            resourceInputs["ignoreMissingGroups"] = args ? args.ignoreMissingGroups : undefined;
            resourceInputs["ldapGroupsDn"] = args ? args.ldapGroupsDn : undefined;
            resourceInputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            resourceInputs["mappedGroupAttributes"] = args ? args.mappedGroupAttributes : undefined;
            resourceInputs["memberofLdapAttribute"] = args ? args.memberofLdapAttribute : undefined;
            resourceInputs["membershipAttributeType"] = args ? args.membershipAttributeType : undefined;
            resourceInputs["membershipLdapAttribute"] = args ? args.membershipLdapAttribute : undefined;
            resourceInputs["membershipUserLdapAttribute"] = args ? args.membershipUserLdapAttribute : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["preserveGroupInheritance"] = args ? args.preserveGroupInheritance : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["userRolesRetrieveStrategy"] = args ? args.userRolesRetrieveStrategy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMapper resources.
 */
export interface GroupMapperState {
    dropNonExistingGroupsDuringSync?: pulumi.Input<boolean>;
    groupNameLdapAttribute?: pulumi.Input<string>;
    groupObjectClasses?: pulumi.Input<pulumi.Input<string>[]>;
    groupsLdapFilter?: pulumi.Input<string>;
    groupsPath?: pulumi.Input<string>;
    ignoreMissingGroups?: pulumi.Input<boolean>;
    ldapGroupsDn?: pulumi.Input<string>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    ldapUserFederationId?: pulumi.Input<string>;
    mappedGroupAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    memberofLdapAttribute?: pulumi.Input<string>;
    membershipAttributeType?: pulumi.Input<string>;
    membershipLdapAttribute?: pulumi.Input<string>;
    membershipUserLdapAttribute?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    preserveGroupInheritance?: pulumi.Input<boolean>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    realmId?: pulumi.Input<string>;
    userRolesRetrieveStrategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMapper resource.
 */
export interface GroupMapperArgs {
    dropNonExistingGroupsDuringSync?: pulumi.Input<boolean>;
    groupNameLdapAttribute: pulumi.Input<string>;
    groupObjectClasses: pulumi.Input<pulumi.Input<string>[]>;
    groupsLdapFilter?: pulumi.Input<string>;
    groupsPath?: pulumi.Input<string>;
    ignoreMissingGroups?: pulumi.Input<boolean>;
    ldapGroupsDn: pulumi.Input<string>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    ldapUserFederationId: pulumi.Input<string>;
    mappedGroupAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    memberofLdapAttribute?: pulumi.Input<string>;
    membershipAttributeType?: pulumi.Input<string>;
    membershipLdapAttribute: pulumi.Input<string>;
    membershipUserLdapAttribute: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    preserveGroupInheritance?: pulumi.Input<boolean>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    realmId: pulumi.Input<string>;
    userRolesRetrieveStrategy?: pulumi.Input<string>;
}
