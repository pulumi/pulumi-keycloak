// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing group mappers for Keycloak users federated via LDAP.
 *
 * The LDAP group mapper can be used to map an LDAP user's groups from some DN to Keycloak groups. This group mapper will also
 * create the groups within Keycloak if they do not already exist.
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
 * const ldapGroupMapper = new keycloak.ldap.GroupMapper("ldapGroupMapper", {
 *     realmId: realm.id,
 *     ldapUserFederationId: ldapUserFederation.id,
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
 * ## Import
 *
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:ldap/groupMapper:GroupMapper ldap_group_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
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

    /**
     * When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
     */
    public readonly dropNonExistingGroupsDuringSync!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
     */
    public readonly groupNameLdapAttribute!: pulumi.Output<string>;
    /**
     * List of strings representing the object classes for the group. Must contain at least one.
     */
    public readonly groupObjectClasses!: pulumi.Output<string[]>;
    /**
     * When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
     */
    public readonly groupsLdapFilter!: pulumi.Output<string | undefined>;
    /**
     * Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
     */
    public readonly groupsPath!: pulumi.Output<string>;
    /**
     * When `true`, missing groups in the hierarchy will be ignored.
     */
    public readonly ignoreMissingGroups!: pulumi.Output<boolean | undefined>;
    /**
     * The LDAP DN where groups can be found.
     */
    public readonly ldapGroupsDn!: pulumi.Output<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    public readonly ldapUserFederationId!: pulumi.Output<string>;
    /**
     * Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
     */
    public readonly mappedGroupAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
     */
    public readonly memberofLdapAttribute!: pulumi.Output<string | undefined>;
    /**
     * Can be one of `DN` or `UID`. Defaults to `DN`.
     */
    public readonly membershipAttributeType!: pulumi.Output<string | undefined>;
    /**
     * The name of the LDAP attribute that is used for membership mappings.
     */
    public readonly membershipLdapAttribute!: pulumi.Output<string>;
    /**
     * The name of the LDAP attribute on a user that is used for membership mappings.
     */
    public readonly membershipUserLdapAttribute!: pulumi.Output<string>;
    /**
     * Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
     */
    public readonly preserveGroupInheritance!: pulumi.Output<boolean | undefined>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
     */
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
    /**
     * When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
     */
    dropNonExistingGroupsDuringSync?: pulumi.Input<boolean>;
    /**
     * The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
     */
    groupNameLdapAttribute?: pulumi.Input<string>;
    /**
     * List of strings representing the object classes for the group. Must contain at least one.
     */
    groupObjectClasses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
     */
    groupsLdapFilter?: pulumi.Input<string>;
    /**
     * Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
     */
    groupsPath?: pulumi.Input<string>;
    /**
     * When `true`, missing groups in the hierarchy will be ignored.
     */
    ignoreMissingGroups?: pulumi.Input<boolean>;
    /**
     * The LDAP DN where groups can be found.
     */
    ldapGroupsDn?: pulumi.Input<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    ldapUserFederationId?: pulumi.Input<string>;
    /**
     * Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
     */
    mappedGroupAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
     */
    memberofLdapAttribute?: pulumi.Input<string>;
    /**
     * Can be one of `DN` or `UID`. Defaults to `DN`.
     */
    membershipAttributeType?: pulumi.Input<string>;
    /**
     * The name of the LDAP attribute that is used for membership mappings.
     */
    membershipLdapAttribute?: pulumi.Input<string>;
    /**
     * The name of the LDAP attribute on a user that is used for membership mappings.
     */
    membershipUserLdapAttribute?: pulumi.Input<string>;
    /**
     * Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
     */
    preserveGroupInheritance?: pulumi.Input<boolean>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    realmId?: pulumi.Input<string>;
    /**
     * Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
     */
    userRolesRetrieveStrategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMapper resource.
 */
export interface GroupMapperArgs {
    /**
     * When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
     */
    dropNonExistingGroupsDuringSync?: pulumi.Input<boolean>;
    /**
     * The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
     */
    groupNameLdapAttribute: pulumi.Input<string>;
    /**
     * List of strings representing the object classes for the group. Must contain at least one.
     */
    groupObjectClasses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
     */
    groupsLdapFilter?: pulumi.Input<string>;
    /**
     * Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
     */
    groupsPath?: pulumi.Input<string>;
    /**
     * When `true`, missing groups in the hierarchy will be ignored.
     */
    ignoreMissingGroups?: pulumi.Input<boolean>;
    /**
     * The LDAP DN where groups can be found.
     */
    ldapGroupsDn: pulumi.Input<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    ldapUserFederationId: pulumi.Input<string>;
    /**
     * Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
     */
    mappedGroupAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
     */
    memberofLdapAttribute?: pulumi.Input<string>;
    /**
     * Can be one of `DN` or `UID`. Defaults to `DN`.
     */
    membershipAttributeType?: pulumi.Input<string>;
    /**
     * The name of the LDAP attribute that is used for membership mappings.
     */
    membershipLdapAttribute: pulumi.Input<string>;
    /**
     * The name of the LDAP attribute on a user that is used for membership mappings.
     */
    membershipUserLdapAttribute: pulumi.Input<string>;
    /**
     * Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
     */
    preserveGroupInheritance?: pulumi.Input<boolean>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    realmId: pulumi.Input<string>;
    /**
     * Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
     */
    userRolesRetrieveStrategy?: pulumi.Input<string>;
}
