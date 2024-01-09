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


public final class RoleMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final RoleMapperArgs Empty = new RoleMapperArgs();

    /**
     * When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `use_realm_roles_mapping` is `false`.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `use_realm_roles_mapping` is `false`.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The LDAP DN where roles can be found.
     * 
     */
    @Import(name="ldapRolesDn", required=true)
    private Output<String> ldapRolesDn;

    /**
     * @return The LDAP DN where roles can be found.
     * 
     */
    public Output<String> ldapRolesDn() {
        return this.ldapRolesDn;
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
     * Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
     * 
     */
    @Import(name="memberofLdapAttribute")
    private @Nullable Output<String> memberofLdapAttribute;

    /**
     * @return Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
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
     * The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
     * 
     */
    @Import(name="roleNameLdapAttribute", required=true)
    private Output<String> roleNameLdapAttribute;

    /**
     * @return The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
     * 
     */
    public Output<String> roleNameLdapAttribute() {
        return this.roleNameLdapAttribute;
    }

    /**
     * List of strings representing the object classes for the role. Must contain at least one.
     * 
     */
    @Import(name="roleObjectClasses", required=true)
    private Output<List<String>> roleObjectClasses;

    /**
     * @return List of strings representing the object classes for the role. Must contain at least one.
     * 
     */
    public Output<List<String>> roleObjectClasses() {
        return this.roleObjectClasses;
    }

    /**
     * When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
     * 
     */
    @Import(name="rolesLdapFilter")
    private @Nullable Output<String> rolesLdapFilter;

    /**
     * @return When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
     * 
     */
    public Optional<Output<String>> rolesLdapFilter() {
        return Optional.ofNullable(this.rolesLdapFilter);
    }

    /**
     * When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
     * 
     */
    @Import(name="useRealmRolesMapping")
    private @Nullable Output<Boolean> useRealmRolesMapping;

    /**
     * @return When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> useRealmRolesMapping() {
        return Optional.ofNullable(this.useRealmRolesMapping);
    }

    /**
     * Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
     * 
     */
    @Import(name="userRolesRetrieveStrategy")
    private @Nullable Output<String> userRolesRetrieveStrategy;

    /**
     * @return Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
     * 
     */
    public Optional<Output<String>> userRolesRetrieveStrategy() {
        return Optional.ofNullable(this.userRolesRetrieveStrategy);
    }

    private RoleMapperArgs() {}

    private RoleMapperArgs(RoleMapperArgs $) {
        this.clientId = $.clientId;
        this.ldapRolesDn = $.ldapRolesDn;
        this.ldapUserFederationId = $.ldapUserFederationId;
        this.memberofLdapAttribute = $.memberofLdapAttribute;
        this.membershipAttributeType = $.membershipAttributeType;
        this.membershipLdapAttribute = $.membershipLdapAttribute;
        this.membershipUserLdapAttribute = $.membershipUserLdapAttribute;
        this.mode = $.mode;
        this.name = $.name;
        this.realmId = $.realmId;
        this.roleNameLdapAttribute = $.roleNameLdapAttribute;
        this.roleObjectClasses = $.roleObjectClasses;
        this.rolesLdapFilter = $.rolesLdapFilter;
        this.useRealmRolesMapping = $.useRealmRolesMapping;
        this.userRolesRetrieveStrategy = $.userRolesRetrieveStrategy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RoleMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RoleMapperArgs $;

        public Builder() {
            $ = new RoleMapperArgs();
        }

        public Builder(RoleMapperArgs defaults) {
            $ = new RoleMapperArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `use_realm_roles_mapping` is `false`.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `use_realm_roles_mapping` is `false`.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param ldapRolesDn The LDAP DN where roles can be found.
         * 
         * @return builder
         * 
         */
        public Builder ldapRolesDn(Output<String> ldapRolesDn) {
            $.ldapRolesDn = ldapRolesDn;
            return this;
        }

        /**
         * @param ldapRolesDn The LDAP DN where roles can be found.
         * 
         * @return builder
         * 
         */
        public Builder ldapRolesDn(String ldapRolesDn) {
            return ldapRolesDn(Output.of(ldapRolesDn));
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
         * @param memberofLdapAttribute Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
         * 
         * @return builder
         * 
         */
        public Builder memberofLdapAttribute(@Nullable Output<String> memberofLdapAttribute) {
            $.memberofLdapAttribute = memberofLdapAttribute;
            return this;
        }

        /**
         * @param memberofLdapAttribute Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
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
         * @param roleNameLdapAttribute The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
         * 
         * @return builder
         * 
         */
        public Builder roleNameLdapAttribute(Output<String> roleNameLdapAttribute) {
            $.roleNameLdapAttribute = roleNameLdapAttribute;
            return this;
        }

        /**
         * @param roleNameLdapAttribute The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
         * 
         * @return builder
         * 
         */
        public Builder roleNameLdapAttribute(String roleNameLdapAttribute) {
            return roleNameLdapAttribute(Output.of(roleNameLdapAttribute));
        }

        /**
         * @param roleObjectClasses List of strings representing the object classes for the role. Must contain at least one.
         * 
         * @return builder
         * 
         */
        public Builder roleObjectClasses(Output<List<String>> roleObjectClasses) {
            $.roleObjectClasses = roleObjectClasses;
            return this;
        }

        /**
         * @param roleObjectClasses List of strings representing the object classes for the role. Must contain at least one.
         * 
         * @return builder
         * 
         */
        public Builder roleObjectClasses(List<String> roleObjectClasses) {
            return roleObjectClasses(Output.of(roleObjectClasses));
        }

        /**
         * @param roleObjectClasses List of strings representing the object classes for the role. Must contain at least one.
         * 
         * @return builder
         * 
         */
        public Builder roleObjectClasses(String... roleObjectClasses) {
            return roleObjectClasses(List.of(roleObjectClasses));
        }

        /**
         * @param rolesLdapFilter When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
         * 
         * @return builder
         * 
         */
        public Builder rolesLdapFilter(@Nullable Output<String> rolesLdapFilter) {
            $.rolesLdapFilter = rolesLdapFilter;
            return this;
        }

        /**
         * @param rolesLdapFilter When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
         * 
         * @return builder
         * 
         */
        public Builder rolesLdapFilter(String rolesLdapFilter) {
            return rolesLdapFilter(Output.of(rolesLdapFilter));
        }

        /**
         * @param useRealmRolesMapping When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder useRealmRolesMapping(@Nullable Output<Boolean> useRealmRolesMapping) {
            $.useRealmRolesMapping = useRealmRolesMapping;
            return this;
        }

        /**
         * @param useRealmRolesMapping When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder useRealmRolesMapping(Boolean useRealmRolesMapping) {
            return useRealmRolesMapping(Output.of(useRealmRolesMapping));
        }

        /**
         * @param userRolesRetrieveStrategy Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
         * 
         * @return builder
         * 
         */
        public Builder userRolesRetrieveStrategy(@Nullable Output<String> userRolesRetrieveStrategy) {
            $.userRolesRetrieveStrategy = userRolesRetrieveStrategy;
            return this;
        }

        /**
         * @param userRolesRetrieveStrategy Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
         * 
         * @return builder
         * 
         */
        public Builder userRolesRetrieveStrategy(String userRolesRetrieveStrategy) {
            return userRolesRetrieveStrategy(Output.of(userRolesRetrieveStrategy));
        }

        public RoleMapperArgs build() {
            if ($.ldapRolesDn == null) {
                throw new MissingRequiredPropertyException("RoleMapperArgs", "ldapRolesDn");
            }
            if ($.ldapUserFederationId == null) {
                throw new MissingRequiredPropertyException("RoleMapperArgs", "ldapUserFederationId");
            }
            if ($.membershipLdapAttribute == null) {
                throw new MissingRequiredPropertyException("RoleMapperArgs", "membershipLdapAttribute");
            }
            if ($.membershipUserLdapAttribute == null) {
                throw new MissingRequiredPropertyException("RoleMapperArgs", "membershipUserLdapAttribute");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("RoleMapperArgs", "realmId");
            }
            if ($.roleNameLdapAttribute == null) {
                throw new MissingRequiredPropertyException("RoleMapperArgs", "roleNameLdapAttribute");
            }
            if ($.roleObjectClasses == null) {
                throw new MissingRequiredPropertyException("RoleMapperArgs", "roleObjectClasses");
            }
            return $;
        }
    }

}
