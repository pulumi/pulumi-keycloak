// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FullNameMapperState extends com.pulumi.resources.ResourceArgs {

    public static final FullNameMapperState Empty = new FullNameMapperState();

    @Import(name="ldapFullNameAttribute")
    private @Nullable Output<String> ldapFullNameAttribute;

    public Optional<Output<String>> ldapFullNameAttribute() {
        return Optional.ofNullable(this.ldapFullNameAttribute);
    }

    /**
     * The ldap user federation provider to attach this mapper to.
     * 
     */
    @Import(name="ldapUserFederationId")
    private @Nullable Output<String> ldapUserFederationId;

    /**
     * @return The ldap user federation provider to attach this mapper to.
     * 
     */
    public Optional<Output<String>> ldapUserFederationId() {
        return Optional.ofNullable(this.ldapUserFederationId);
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

    @Import(name="readOnly")
    private @Nullable Output<Boolean> readOnly;

    public Optional<Output<Boolean>> readOnly() {
        return Optional.ofNullable(this.readOnly);
    }

    /**
     * The realm in which the ldap user federation provider exists.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm in which the ldap user federation provider exists.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    @Import(name="writeOnly")
    private @Nullable Output<Boolean> writeOnly;

    public Optional<Output<Boolean>> writeOnly() {
        return Optional.ofNullable(this.writeOnly);
    }

    private FullNameMapperState() {}

    private FullNameMapperState(FullNameMapperState $) {
        this.ldapFullNameAttribute = $.ldapFullNameAttribute;
        this.ldapUserFederationId = $.ldapUserFederationId;
        this.name = $.name;
        this.readOnly = $.readOnly;
        this.realmId = $.realmId;
        this.writeOnly = $.writeOnly;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FullNameMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FullNameMapperState $;

        public Builder() {
            $ = new FullNameMapperState();
        }

        public Builder(FullNameMapperState defaults) {
            $ = new FullNameMapperState(Objects.requireNonNull(defaults));
        }

        public Builder ldapFullNameAttribute(@Nullable Output<String> ldapFullNameAttribute) {
            $.ldapFullNameAttribute = ldapFullNameAttribute;
            return this;
        }

        public Builder ldapFullNameAttribute(String ldapFullNameAttribute) {
            return ldapFullNameAttribute(Output.of(ldapFullNameAttribute));
        }

        /**
         * @param ldapUserFederationId The ldap user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(@Nullable Output<String> ldapUserFederationId) {
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

        public Builder readOnly(@Nullable Output<Boolean> readOnly) {
            $.readOnly = readOnly;
            return this;
        }

        public Builder readOnly(Boolean readOnly) {
            return readOnly(Output.of(readOnly));
        }

        /**
         * @param realmId The realm in which the ldap user federation provider exists.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
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

        public Builder writeOnly(@Nullable Output<Boolean> writeOnly) {
            $.writeOnly = writeOnly;
            return this;
        }

        public Builder writeOnly(Boolean writeOnly) {
            return writeOnly(Output.of(writeOnly));
        }

        public FullNameMapperState build() {
            return $;
        }
    }

}
