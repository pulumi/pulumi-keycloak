// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupMapperArgs Empty = new GroupMapperArgs();

    /**
     * When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
     * 
     */
    @Import(name="dropNonExistingGroupsDuringSync")
    private @Nullable Output<Boolean> dropNonExistingGroupsDuringSync;

    /**
     * @return When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> dropNonExistingGroupsDuringSync() {
        return Optional.ofNullable(this.dropNonExistingGroupsDuringSync);
    }

    /**
     * The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
     * 
     */
    @Import(name="groupNameLdapAttribute", required=true)
    private Output<String> groupNameLdapAttribute;

    /**
     * @return The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
     * 
     */
    public Output<String> groupNameLdapAttribute() {
        return this.groupNameLdapAttribute;
    }

    /**
     * List of strings representing the object classes for the group. Must contain at least one.
     * 
     */
    @Import(name="groupObjectClasses", required=true)
    private Output<List<String>> groupObjectClasses;

    /**
     * @return List of strings representing the object classes for the group. Must contain at least one.
     * 
     */
    public Output<List<String>> groupObjectClasses() {
        return this.groupObjectClasses;
    }

    /**
     * When specified, adds a custom filter to be used when querying for groups. Must start with `(` and end with `)`.
     * 
     */
    @Import(name="groupsLdapFilter")
    private @Nullable Output<String> groupsLdapFilter;

    /**
     * @return When specified, adds a custom filter to be used when querying for groups. Must start with `(` and end with `)`.
     * 
     */
    public Optional<Output<String>> groupsLdapFilter() {
        return Optional.ofNullable(this.groupsLdapFilter);
    }

    /**
     * Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
     * 
     */
    @Import(name="groupsPath")
    private @Nullable Output<String> groupsPath;

    /**
     * @return Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
     * 
     */
    public Optional<Output<String>> groupsPath() {
        return Optional.ofNullable(this.groupsPath);
    }

    /**
     * When `true`, missing groups in the hierarchy will be ignored.
     * 
     */
    @Import(name="ignoreMissingGroups")
    private @Nullable Output<Boolean> ignoreMissingGroups;

    /**
     * @return When `true`, missing groups in the hierarchy will be ignored.
     * 
     */
    public Optional<Output<Boolean>> ignoreMissingGroups() {
        return Optional.ofNullable(this.ignoreMissingGroups);
    }

    /**
     * The LDAP DN where groups can be found.
     * 
     */
    @Import(name="ldapGroupsDn", required=true)
    private Output<String> ldapGroupsDn;

    /**
     * @return The LDAP DN where groups can be found.
     * 
     */
    public Output<String> ldapGroupsDn() {
        return this.ldapGroupsDn;
    }

    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    @Import(name="ldapUserFederationId", required=true)
    private Output<String> ldapUserFederationId;

    /**
     * @return The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    public Output<String> ldapUserFederationId() {
        return this.ldapUserFederationId;
    }

    /**
     * Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
     * 
     */
    @Import(name="mappedGroupAttributes")
    private @Nullable Output<List<String>> mappedGroupAttributes;

    /**
     * @return Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
     * 
     */
    public Optional<Output<List<String>>> mappedGroupAttributes() {
        return Optional.ofNullable(this.mappedGroupAttributes);
    }

    /**
     * Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
     * 
     */
    @Import(name="memberofLdapAttribute")
    private @Nullable Output<String> memberofLdapAttribute;

    /**
     * @return Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
     * 
     */
    public Optional<Output<String>> memberofLdapAttribute() {
        return Optional.ofNullable(this.memberofLdapAttribute);
    }

    /**
     * Can be one of `DN` or `UID`. Defaults to `DN`.
     * 
     */
    @Import(name="membershipAttributeType")
    private @Nullable Output<String> membershipAttributeType;

    /**
     * @return Can be one of `DN` or `UID`. Defaults to `DN`.
     * 
     */
    public Optional<Output<String>> membershipAttributeType() {
        return Optional.ofNullable(this.membershipAttributeType);
    }

    /**
     * The name of the LDAP attribute that is used for membership mappings.
     * 
     */
    @Import(name="membershipLdapAttribute", required=true)
    private Output<String> membershipLdapAttribute;

    /**
     * @return The name of the LDAP attribute that is used for membership mappings.
     * 
     */
    public Output<String> membershipLdapAttribute() {
        return this.membershipLdapAttribute;
    }

    /**
     * The name of the LDAP attribute on a user that is used for membership mappings.
     * 
     */
    @Import(name="membershipUserLdapAttribute", required=true)
    private Output<String> membershipUserLdapAttribute;

    /**
     * @return The name of the LDAP attribute on a user that is used for membership mappings.
     * 
     */
    public Output<String> membershipUserLdapAttribute() {
        return this.membershipUserLdapAttribute;
    }

    /**
     * Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * Display name of this mapper when displayed in the console.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Display name of this mapper when displayed in the console.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
     * 
     */
    @Import(name="preserveGroupInheritance")
    private @Nullable Output<Boolean> preserveGroupInheritance;

    /**
     * @return When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
     * 
     */
    public Optional<Output<Boolean>> preserveGroupInheritance() {
        return Optional.ofNullable(this.preserveGroupInheritance);
    }

    /**
     * The realm that this LDAP mapper will exist in.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm that this LDAP mapper will exist in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
     * 
     */
    @Import(name="userRolesRetrieveStrategy")
    private @Nullable Output<String> userRolesRetrieveStrategy;

    /**
     * @return Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
     * 
     */
    public Optional<Output<String>> userRolesRetrieveStrategy() {
        return Optional.ofNullable(this.userRolesRetrieveStrategy);
    }

    private GroupMapperArgs() {}

    private GroupMapperArgs(GroupMapperArgs $) {
        this.dropNonExistingGroupsDuringSync = $.dropNonExistingGroupsDuringSync;
        this.groupNameLdapAttribute = $.groupNameLdapAttribute;
        this.groupObjectClasses = $.groupObjectClasses;
        this.groupsLdapFilter = $.groupsLdapFilter;
        this.groupsPath = $.groupsPath;
        this.ignoreMissingGroups = $.ignoreMissingGroups;
        this.ldapGroupsDn = $.ldapGroupsDn;
        this.ldapUserFederationId = $.ldapUserFederationId;
        this.mappedGroupAttributes = $.mappedGroupAttributes;
        this.memberofLdapAttribute = $.memberofLdapAttribute;
        this.membershipAttributeType = $.membershipAttributeType;
        this.membershipLdapAttribute = $.membershipLdapAttribute;
        this.membershipUserLdapAttribute = $.membershipUserLdapAttribute;
        this.mode = $.mode;
        this.name = $.name;
        this.preserveGroupInheritance = $.preserveGroupInheritance;
        this.realmId = $.realmId;
        this.userRolesRetrieveStrategy = $.userRolesRetrieveStrategy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMapperArgs $;

        public Builder() {
            $ = new GroupMapperArgs();
        }

        public Builder(GroupMapperArgs defaults) {
            $ = new GroupMapperArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dropNonExistingGroupsDuringSync When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder dropNonExistingGroupsDuringSync(@Nullable Output<Boolean> dropNonExistingGroupsDuringSync) {
            $.dropNonExistingGroupsDuringSync = dropNonExistingGroupsDuringSync;
            return this;
        }

        /**
         * @param dropNonExistingGroupsDuringSync When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder dropNonExistingGroupsDuringSync(Boolean dropNonExistingGroupsDuringSync) {
            return dropNonExistingGroupsDuringSync(Output.of(dropNonExistingGroupsDuringSync));
        }

        /**
         * @param groupNameLdapAttribute The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
         * 
         * @return builder
         * 
         */
        public Builder groupNameLdapAttribute(Output<String> groupNameLdapAttribute) {
            $.groupNameLdapAttribute = groupNameLdapAttribute;
            return this;
        }

        /**
         * @param groupNameLdapAttribute The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
         * 
         * @return builder
         * 
         */
        public Builder groupNameLdapAttribute(String groupNameLdapAttribute) {
            return groupNameLdapAttribute(Output.of(groupNameLdapAttribute));
        }

        /**
         * @param groupObjectClasses List of strings representing the object classes for the group. Must contain at least one.
         * 
         * @return builder
         * 
         */
        public Builder groupObjectClasses(Output<List<String>> groupObjectClasses) {
            $.groupObjectClasses = groupObjectClasses;
            return this;
        }

        /**
         * @param groupObjectClasses List of strings representing the object classes for the group. Must contain at least one.
         * 
         * @return builder
         * 
         */
        public Builder groupObjectClasses(List<String> groupObjectClasses) {
            return groupObjectClasses(Output.of(groupObjectClasses));
        }

        /**
         * @param groupObjectClasses List of strings representing the object classes for the group. Must contain at least one.
         * 
         * @return builder
         * 
         */
        public Builder groupObjectClasses(String... groupObjectClasses) {
            return groupObjectClasses(List.of(groupObjectClasses));
        }

        /**
         * @param groupsLdapFilter When specified, adds a custom filter to be used when querying for groups. Must start with `(` and end with `)`.
         * 
         * @return builder
         * 
         */
        public Builder groupsLdapFilter(@Nullable Output<String> groupsLdapFilter) {
            $.groupsLdapFilter = groupsLdapFilter;
            return this;
        }

        /**
         * @param groupsLdapFilter When specified, adds a custom filter to be used when querying for groups. Must start with `(` and end with `)`.
         * 
         * @return builder
         * 
         */
        public Builder groupsLdapFilter(String groupsLdapFilter) {
            return groupsLdapFilter(Output.of(groupsLdapFilter));
        }

        /**
         * @param groupsPath Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
         * 
         * @return builder
         * 
         */
        public Builder groupsPath(@Nullable Output<String> groupsPath) {
            $.groupsPath = groupsPath;
            return this;
        }

        /**
         * @param groupsPath Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
         * 
         * @return builder
         * 
         */
        public Builder groupsPath(String groupsPath) {
            return groupsPath(Output.of(groupsPath));
        }

        /**
         * @param ignoreMissingGroups When `true`, missing groups in the hierarchy will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder ignoreMissingGroups(@Nullable Output<Boolean> ignoreMissingGroups) {
            $.ignoreMissingGroups = ignoreMissingGroups;
            return this;
        }

        /**
         * @param ignoreMissingGroups When `true`, missing groups in the hierarchy will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder ignoreMissingGroups(Boolean ignoreMissingGroups) {
            return ignoreMissingGroups(Output.of(ignoreMissingGroups));
        }

        /**
         * @param ldapGroupsDn The LDAP DN where groups can be found.
         * 
         * @return builder
         * 
         */
        public Builder ldapGroupsDn(Output<String> ldapGroupsDn) {
            $.ldapGroupsDn = ldapGroupsDn;
            return this;
        }

        /**
         * @param ldapGroupsDn The LDAP DN where groups can be found.
         * 
         * @return builder
         * 
         */
        public Builder ldapGroupsDn(String ldapGroupsDn) {
            return ldapGroupsDn(Output.of(ldapGroupsDn));
        }

        /**
         * @param ldapUserFederationId The ID of the LDAP user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(Output<String> ldapUserFederationId) {
            $.ldapUserFederationId = ldapUserFederationId;
            return this;
        }

        /**
         * @param ldapUserFederationId The ID of the LDAP user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(String ldapUserFederationId) {
            return ldapUserFederationId(Output.of(ldapUserFederationId));
        }

        /**
         * @param mappedGroupAttributes Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
         * 
         * @return builder
         * 
         */
        public Builder mappedGroupAttributes(@Nullable Output<List<String>> mappedGroupAttributes) {
            $.mappedGroupAttributes = mappedGroupAttributes;
            return this;
        }

        /**
         * @param mappedGroupAttributes Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
         * 
         * @return builder
         * 
         */
        public Builder mappedGroupAttributes(List<String> mappedGroupAttributes) {
            return mappedGroupAttributes(Output.of(mappedGroupAttributes));
        }

        /**
         * @param mappedGroupAttributes Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
         * 
         * @return builder
         * 
         */
        public Builder mappedGroupAttributes(String... mappedGroupAttributes) {
            return mappedGroupAttributes(List.of(mappedGroupAttributes));
        }

        /**
         * @param memberofLdapAttribute Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
         * 
         * @return builder
         * 
         */
        public Builder memberofLdapAttribute(@Nullable Output<String> memberofLdapAttribute) {
            $.memberofLdapAttribute = memberofLdapAttribute;
            return this;
        }

        /**
         * @param memberofLdapAttribute Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
         * 
         * @return builder
         * 
         */
        public Builder memberofLdapAttribute(String memberofLdapAttribute) {
            return memberofLdapAttribute(Output.of(memberofLdapAttribute));
        }

        /**
         * @param membershipAttributeType Can be one of `DN` or `UID`. Defaults to `DN`.
         * 
         * @return builder
         * 
         */
        public Builder membershipAttributeType(@Nullable Output<String> membershipAttributeType) {
            $.membershipAttributeType = membershipAttributeType;
            return this;
        }

        /**
         * @param membershipAttributeType Can be one of `DN` or `UID`. Defaults to `DN`.
         * 
         * @return builder
         * 
         */
        public Builder membershipAttributeType(String membershipAttributeType) {
            return membershipAttributeType(Output.of(membershipAttributeType));
        }

        /**
         * @param membershipLdapAttribute The name of the LDAP attribute that is used for membership mappings.
         * 
         * @return builder
         * 
         */
        public Builder membershipLdapAttribute(Output<String> membershipLdapAttribute) {
            $.membershipLdapAttribute = membershipLdapAttribute;
            return this;
        }

        /**
         * @param membershipLdapAttribute The name of the LDAP attribute that is used for membership mappings.
         * 
         * @return builder
         * 
         */
        public Builder membershipLdapAttribute(String membershipLdapAttribute) {
            return membershipLdapAttribute(Output.of(membershipLdapAttribute));
        }

        /**
         * @param membershipUserLdapAttribute The name of the LDAP attribute on a user that is used for membership mappings.
         * 
         * @return builder
         * 
         */
        public Builder membershipUserLdapAttribute(Output<String> membershipUserLdapAttribute) {
            $.membershipUserLdapAttribute = membershipUserLdapAttribute;
            return this;
        }

        /**
         * @param membershipUserLdapAttribute The name of the LDAP attribute on a user that is used for membership mappings.
         * 
         * @return builder
         * 
         */
        public Builder membershipUserLdapAttribute(String membershipUserLdapAttribute) {
            return membershipUserLdapAttribute(Output.of(membershipUserLdapAttribute));
        }

        /**
         * @param mode Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name Display name of this mapper when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Display name of this mapper when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param preserveGroupInheritance When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
         * 
         * @return builder
         * 
         */
        public Builder preserveGroupInheritance(@Nullable Output<Boolean> preserveGroupInheritance) {
            $.preserveGroupInheritance = preserveGroupInheritance;
            return this;
        }

        /**
         * @param preserveGroupInheritance When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
         * 
         * @return builder
         * 
         */
        public Builder preserveGroupInheritance(Boolean preserveGroupInheritance) {
            return preserveGroupInheritance(Output.of(preserveGroupInheritance));
        }

        /**
         * @param realmId The realm that this LDAP mapper will exist in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm that this LDAP mapper will exist in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param userRolesRetrieveStrategy Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
         * 
         * @return builder
         * 
         */
        public Builder userRolesRetrieveStrategy(@Nullable Output<String> userRolesRetrieveStrategy) {
            $.userRolesRetrieveStrategy = userRolesRetrieveStrategy;
            return this;
        }

        /**
         * @param userRolesRetrieveStrategy Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
         * 
         * @return builder
         * 
         */
        public Builder userRolesRetrieveStrategy(String userRolesRetrieveStrategy) {
            return userRolesRetrieveStrategy(Output.of(userRolesRetrieveStrategy));
        }

        public GroupMapperArgs build() {
            if ($.groupNameLdapAttribute == null) {
                throw new MissingRequiredPropertyException("GroupMapperArgs", "groupNameLdapAttribute");
            }
            if ($.groupObjectClasses == null) {
                throw new MissingRequiredPropertyException("GroupMapperArgs", "groupObjectClasses");
            }
            if ($.ldapGroupsDn == null) {
                throw new MissingRequiredPropertyException("GroupMapperArgs", "ldapGroupsDn");
            }
            if ($.ldapUserFederationId == null) {
                throw new MissingRequiredPropertyException("GroupMapperArgs", "ldapUserFederationId");
            }
            if ($.membershipLdapAttribute == null) {
                throw new MissingRequiredPropertyException("GroupMapperArgs", "membershipLdapAttribute");
            }
            if ($.membershipUserLdapAttribute == null) {
                throw new MissingRequiredPropertyException("GroupMapperArgs", "membershipUserLdapAttribute");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("GroupMapperArgs", "realmId");
            }
            return $;
        }
    }

}
