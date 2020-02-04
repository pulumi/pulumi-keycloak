// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class GroupMapper extends pulumi.CustomResource {
    /**
     * Get an existing GroupMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMapperState, opts?: pulumi.CustomResourceOptions): GroupMapper {
        return new GroupMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:Ldap/groupMapper:GroupMapper';

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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupMapperState | undefined;
            inputs["dropNonExistingGroupsDuringSync"] = state ? state.dropNonExistingGroupsDuringSync : undefined;
            inputs["groupNameLdapAttribute"] = state ? state.groupNameLdapAttribute : undefined;
            inputs["groupObjectClasses"] = state ? state.groupObjectClasses : undefined;
            inputs["groupsLdapFilter"] = state ? state.groupsLdapFilter : undefined;
            inputs["ignoreMissingGroups"] = state ? state.ignoreMissingGroups : undefined;
            inputs["ldapGroupsDn"] = state ? state.ldapGroupsDn : undefined;
            inputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            inputs["mappedGroupAttributes"] = state ? state.mappedGroupAttributes : undefined;
            inputs["memberofLdapAttribute"] = state ? state.memberofLdapAttribute : undefined;
            inputs["membershipAttributeType"] = state ? state.membershipAttributeType : undefined;
            inputs["membershipLdapAttribute"] = state ? state.membershipLdapAttribute : undefined;
            inputs["membershipUserLdapAttribute"] = state ? state.membershipUserLdapAttribute : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["preserveGroupInheritance"] = state ? state.preserveGroupInheritance : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["userRolesRetrieveStrategy"] = state ? state.userRolesRetrieveStrategy : undefined;
        } else {
            const args = argsOrState as GroupMapperArgs | undefined;
            if (!args || args.groupNameLdapAttribute === undefined) {
                throw new Error("Missing required property 'groupNameLdapAttribute'");
            }
            if (!args || args.groupObjectClasses === undefined) {
                throw new Error("Missing required property 'groupObjectClasses'");
            }
            if (!args || args.ldapGroupsDn === undefined) {
                throw new Error("Missing required property 'ldapGroupsDn'");
            }
            if (!args || args.ldapUserFederationId === undefined) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if (!args || args.membershipLdapAttribute === undefined) {
                throw new Error("Missing required property 'membershipLdapAttribute'");
            }
            if (!args || args.membershipUserLdapAttribute === undefined) {
                throw new Error("Missing required property 'membershipUserLdapAttribute'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["dropNonExistingGroupsDuringSync"] = args ? args.dropNonExistingGroupsDuringSync : undefined;
            inputs["groupNameLdapAttribute"] = args ? args.groupNameLdapAttribute : undefined;
            inputs["groupObjectClasses"] = args ? args.groupObjectClasses : undefined;
            inputs["groupsLdapFilter"] = args ? args.groupsLdapFilter : undefined;
            inputs["ignoreMissingGroups"] = args ? args.ignoreMissingGroups : undefined;
            inputs["ldapGroupsDn"] = args ? args.ldapGroupsDn : undefined;
            inputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            inputs["mappedGroupAttributes"] = args ? args.mappedGroupAttributes : undefined;
            inputs["memberofLdapAttribute"] = args ? args.memberofLdapAttribute : undefined;
            inputs["membershipAttributeType"] = args ? args.membershipAttributeType : undefined;
            inputs["membershipLdapAttribute"] = args ? args.membershipLdapAttribute : undefined;
            inputs["membershipUserLdapAttribute"] = args ? args.membershipUserLdapAttribute : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["preserveGroupInheritance"] = args ? args.preserveGroupInheritance : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["userRolesRetrieveStrategy"] = args ? args.userRolesRetrieveStrategy : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMapper resources.
 */
export interface GroupMapperState {
    readonly dropNonExistingGroupsDuringSync?: pulumi.Input<boolean>;
    readonly groupNameLdapAttribute?: pulumi.Input<string>;
    readonly groupObjectClasses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly groupsLdapFilter?: pulumi.Input<string>;
    readonly ignoreMissingGroups?: pulumi.Input<boolean>;
    readonly ldapGroupsDn?: pulumi.Input<string>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    readonly ldapUserFederationId?: pulumi.Input<string>;
    readonly mappedGroupAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly memberofLdapAttribute?: pulumi.Input<string>;
    readonly membershipAttributeType?: pulumi.Input<string>;
    readonly membershipLdapAttribute?: pulumi.Input<string>;
    readonly membershipUserLdapAttribute?: pulumi.Input<string>;
    readonly mode?: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    readonly name?: pulumi.Input<string>;
    readonly preserveGroupInheritance?: pulumi.Input<boolean>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    readonly realmId?: pulumi.Input<string>;
    readonly userRolesRetrieveStrategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMapper resource.
 */
export interface GroupMapperArgs {
    readonly dropNonExistingGroupsDuringSync?: pulumi.Input<boolean>;
    readonly groupNameLdapAttribute: pulumi.Input<string>;
    readonly groupObjectClasses: pulumi.Input<pulumi.Input<string>[]>;
    readonly groupsLdapFilter?: pulumi.Input<string>;
    readonly ignoreMissingGroups?: pulumi.Input<boolean>;
    readonly ldapGroupsDn: pulumi.Input<string>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    readonly ldapUserFederationId: pulumi.Input<string>;
    readonly mappedGroupAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly memberofLdapAttribute?: pulumi.Input<string>;
    readonly membershipAttributeType?: pulumi.Input<string>;
    readonly membershipLdapAttribute: pulumi.Input<string>;
    readonly membershipUserLdapAttribute: pulumi.Input<string>;
    readonly mode?: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    readonly name?: pulumi.Input<string>;
    readonly preserveGroupInheritance?: pulumi.Input<boolean>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    readonly realmId: pulumi.Input<string>;
    readonly userRolesRetrieveStrategy?: pulumi.Input<string>;
}
