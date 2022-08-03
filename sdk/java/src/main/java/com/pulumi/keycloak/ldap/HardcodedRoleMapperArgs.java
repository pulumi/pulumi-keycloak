// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HardcodedRoleMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final HardcodedRoleMapperArgs Empty = new HardcodedRoleMapperArgs();

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
     * The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private HardcodedRoleMapperArgs() {}

    private HardcodedRoleMapperArgs(HardcodedRoleMapperArgs $) {
        this.ldapUserFederationId = $.ldapUserFederationId;
        this.name = $.name;
        this.realmId = $.realmId;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HardcodedRoleMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HardcodedRoleMapperArgs $;

        public Builder() {
            $ = new HardcodedRoleMapperArgs();
        }

        public Builder(HardcodedRoleMapperArgs defaults) {
            $ = new HardcodedRoleMapperArgs(Objects.requireNonNull(defaults));
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
         * @param role The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public HardcodedRoleMapperArgs build() {
            $.ldapUserFederationId = Objects.requireNonNull($.ldapUserFederationId, "expected parameter 'ldapUserFederationId' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
