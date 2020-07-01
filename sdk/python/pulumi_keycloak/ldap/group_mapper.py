# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GroupMapper(pulumi.CustomResource):
    drop_non_existing_groups_during_sync: pulumi.Output[bool]
    group_name_ldap_attribute: pulumi.Output[str]
    group_object_classes: pulumi.Output[list]
    groups_ldap_filter: pulumi.Output[str]
    ignore_missing_groups: pulumi.Output[bool]
    ldap_groups_dn: pulumi.Output[str]
    ldap_user_federation_id: pulumi.Output[str]
    """
    The ldap user federation provider to attach this mapper to.
    """
    mapped_group_attributes: pulumi.Output[list]
    memberof_ldap_attribute: pulumi.Output[str]
    membership_attribute_type: pulumi.Output[str]
    membership_ldap_attribute: pulumi.Output[str]
    membership_user_ldap_attribute: pulumi.Output[str]
    mode: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    Display name of the mapper when displayed in the console.
    """
    preserve_group_inheritance: pulumi.Output[bool]
    realm_id: pulumi.Output[str]
    """
    The realm in which the ldap user federation provider exists.
    """
    user_roles_retrieve_strategy: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, drop_non_existing_groups_during_sync=None, group_name_ldap_attribute=None, group_object_classes=None, groups_ldap_filter=None, ignore_missing_groups=None, ldap_groups_dn=None, ldap_user_federation_id=None, mapped_group_attributes=None, memberof_ldap_attribute=None, membership_attribute_type=None, membership_ldap_attribute=None, membership_user_ldap_attribute=None, mode=None, name=None, preserve_group_inheritance=None, realm_id=None, user_roles_retrieve_strategy=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # ldap.GroupMapper

        Allows for creating and managing group mappers for Keycloak users federated
        via LDAP.

        The LDAP group mapper can be used to map an LDAP user's groups from some DN
        to Keycloak groups. This group mapper will also create the groups within Keycloak
        if they do not already exist.

        ### Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            enabled=True,
            realm="test")
        ldap_user_federation = keycloak.ldap.UserFederation("ldapUserFederation",
            bind_credential="admin",
            bind_dn="cn=admin,dc=example,dc=org",
            connection_url="ldap://openldap",
            rdn_ldap_attribute="cn",
            realm_id=realm.id,
            user_object_classes=[
                "simpleSecurityObject",
                "organizationalRole",
            ],
            username_ldap_attribute="cn",
            users_dn="dc=example,dc=org",
            uuid_ldap_attribute="entryDN")
        ldap_group_mapper = keycloak.ldap.GroupMapper("ldapGroupMapper",
            group_name_ldap_attribute="cn",
            group_object_classes=["groupOfNames"],
            ldap_groups_dn="dc=example,dc=org",
            ldap_user_federation_id=ldap_user_federation.id,
            memberof_ldap_attribute="memberOf",
            membership_attribute_type="DN",
            membership_ldap_attribute="member",
            membership_user_ldap_attribute="cn",
            realm_id=realm.id)
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm that this LDAP mapper will exist in.
        - `ldap_user_federation_id` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
        - `name` - (Required) Display name of this mapper when displayed in the console.
        - `ldap_groups_dn` - (Required) The LDAP DN where groups can be found.
        - `group_name_ldap_attribute` - (Required) The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
        - `group_object_classes` - (Required) Array of strings representing the object classes for the group. Must contain at least one.
        - `preserve_group_inheritance` - (Optional) When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
        - `ignore_missing_groups` - (Optional) When `true`, missing groups in the hierarchy will be ignored.
        - `membership_ldap_attribute` - (Required) The name of the LDAP attribute that is used for membership mappings.
        - `membership_attribute_type` - (Optional) Can be one of `DN` or `UID`. Defaults to `DN`.
        - `membership_user_ldap_attribute` - (Required) The name of the LDAP attribute on a user that is used for membership mappings.
        - `groups_ldap_filter` - (Optional) When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
        - `mode` - (Optional) Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
        - `user_roles_retrieve_strategy` - (Optional) Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
        - `memberof_ldap_attribute` - (Optional) Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
        - `mapped_group_attributes` - (Optional) Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
        - `drop_non_existing_groups_during_sync` - (Optional) When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ldap user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of the mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm in which the ldap user federation provider exists.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['drop_non_existing_groups_during_sync'] = drop_non_existing_groups_during_sync
            if group_name_ldap_attribute is None:
                raise TypeError("Missing required property 'group_name_ldap_attribute'")
            __props__['group_name_ldap_attribute'] = group_name_ldap_attribute
            if group_object_classes is None:
                raise TypeError("Missing required property 'group_object_classes'")
            __props__['group_object_classes'] = group_object_classes
            __props__['groups_ldap_filter'] = groups_ldap_filter
            __props__['ignore_missing_groups'] = ignore_missing_groups
            if ldap_groups_dn is None:
                raise TypeError("Missing required property 'ldap_groups_dn'")
            __props__['ldap_groups_dn'] = ldap_groups_dn
            if ldap_user_federation_id is None:
                raise TypeError("Missing required property 'ldap_user_federation_id'")
            __props__['ldap_user_federation_id'] = ldap_user_federation_id
            __props__['mapped_group_attributes'] = mapped_group_attributes
            __props__['memberof_ldap_attribute'] = memberof_ldap_attribute
            __props__['membership_attribute_type'] = membership_attribute_type
            if membership_ldap_attribute is None:
                raise TypeError("Missing required property 'membership_ldap_attribute'")
            __props__['membership_ldap_attribute'] = membership_ldap_attribute
            if membership_user_ldap_attribute is None:
                raise TypeError("Missing required property 'membership_user_ldap_attribute'")
            __props__['membership_user_ldap_attribute'] = membership_user_ldap_attribute
            __props__['mode'] = mode
            __props__['name'] = name
            __props__['preserve_group_inheritance'] = preserve_group_inheritance
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            __props__['user_roles_retrieve_strategy'] = user_roles_retrieve_strategy
        super(GroupMapper, __self__).__init__(
            'keycloak:ldap/groupMapper:GroupMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, drop_non_existing_groups_during_sync=None, group_name_ldap_attribute=None, group_object_classes=None, groups_ldap_filter=None, ignore_missing_groups=None, ldap_groups_dn=None, ldap_user_federation_id=None, mapped_group_attributes=None, memberof_ldap_attribute=None, membership_attribute_type=None, membership_ldap_attribute=None, membership_user_ldap_attribute=None, mode=None, name=None, preserve_group_inheritance=None, realm_id=None, user_roles_retrieve_strategy=None):
        """
        Get an existing GroupMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ldap user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of the mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm in which the ldap user federation provider exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["drop_non_existing_groups_during_sync"] = drop_non_existing_groups_during_sync
        __props__["group_name_ldap_attribute"] = group_name_ldap_attribute
        __props__["group_object_classes"] = group_object_classes
        __props__["groups_ldap_filter"] = groups_ldap_filter
        __props__["ignore_missing_groups"] = ignore_missing_groups
        __props__["ldap_groups_dn"] = ldap_groups_dn
        __props__["ldap_user_federation_id"] = ldap_user_federation_id
        __props__["mapped_group_attributes"] = mapped_group_attributes
        __props__["memberof_ldap_attribute"] = memberof_ldap_attribute
        __props__["membership_attribute_type"] = membership_attribute_type
        __props__["membership_ldap_attribute"] = membership_ldap_attribute
        __props__["membership_user_ldap_attribute"] = membership_user_ldap_attribute
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["preserve_group_inheritance"] = preserve_group_inheritance
        __props__["realm_id"] = realm_id
        __props__["user_roles_retrieve_strategy"] = user_roles_retrieve_strategy
        return GroupMapper(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
