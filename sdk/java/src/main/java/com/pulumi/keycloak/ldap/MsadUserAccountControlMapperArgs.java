// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MsadUserAccountControlMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final MsadUserAccountControlMapperArgs Empty = new MsadUserAccountControlMapperArgs();

    @Import(name="ldapPasswordPolicyHintsEnabled")
    private @Nullable Output<Boolean> ldapPasswordPolicyHintsEnabled;

    public Optional<Output<Boolean>> ldapPasswordPolicyHintsEnabled() {
        return Optional.ofNullable(this.ldapPasswordPolicyHintsEnabled);
    }

    /**
     * The ldap user federation provider to attach this mapper to.
     * 
     */
    @Import(name="ldapUserFederationId", required=true)
    private Output<String> ldapUserFederationId;

    /**
     * @return The ldap user federation provider to attach this mapper to.
     * 
     */
    public Output<String> ldapUserFederationId() {
        return this.ldapUserFederationId;
    }

    /**
     * Display name of the mapper when displayed in the console.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Display name of the mapper when displayed in the console.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The realm in which the ldap user federation provider exists.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm in which the ldap user federation provider exists.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private MsadUserAccountControlMapperArgs() {}

    private MsadUserAccountControlMapperArgs(MsadUserAccountControlMapperArgs $) {
        this.ldapPasswordPolicyHintsEnabled = $.ldapPasswordPolicyHintsEnabled;
        this.ldapUserFederationId = $.ldapUserFederationId;
        this.name = $.name;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MsadUserAccountControlMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MsadUserAccountControlMapperArgs $;

        public Builder() {
            $ = new MsadUserAccountControlMapperArgs();
        }

        public Builder(MsadUserAccountControlMapperArgs defaults) {
            $ = new MsadUserAccountControlMapperArgs(Objects.requireNonNull(defaults));
        }

        public Builder ldapPasswordPolicyHintsEnabled(@Nullable Output<Boolean> ldapPasswordPolicyHintsEnabled) {
            $.ldapPasswordPolicyHintsEnabled = ldapPasswordPolicyHintsEnabled;
            return this;
        }

        public Builder ldapPasswordPolicyHintsEnabled(Boolean ldapPasswordPolicyHintsEnabled) {
            return ldapPasswordPolicyHintsEnabled(Output.of(ldapPasswordPolicyHintsEnabled));
        }

        /**
         * @param ldapUserFederationId The ldap user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(Output<String> ldapUserFederationId) {
            $.ldapUserFederationId = ldapUserFederationId;
            return this;
        }

        /**
         * @param ldapUserFederationId The ldap user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(String ldapUserFederationId) {
            return ldapUserFederationId(Output.of(ldapUserFederationId));
        }

        /**
         * @param name Display name of the mapper when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Display name of the mapper when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm in which the ldap user federation provider exists.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm in which the ldap user federation provider exists.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public MsadUserAccountControlMapperArgs build() {
            if ($.ldapUserFederationId == null) {
                throw new MissingRequiredPropertyException("MsadUserAccountControlMapperArgs", "ldapUserFederationId");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("MsadUserAccountControlMapperArgs", "realmId");
            }
            return $;
        }
    }

}
