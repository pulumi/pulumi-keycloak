// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CustomMapperState extends com.pulumi.resources.ResourceArgs {

    public static final CustomMapperState Empty = new CustomMapperState();

    /**
     * A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
     * 
     */
    @Import(name="config")
    private @Nullable Output<Map<String,Object>> config;

    /**
     * @return A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
     * 
     */
    public Optional<Output<Map<String,Object>>> config() {
        return Optional.ofNullable(this.config);
    }

    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    @Import(name="ldapUserFederationId")
    private @Nullable Output<String> ldapUserFederationId;

    /**
     * @return The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    public Optional<Output<String>> ldapUserFederationId() {
        return Optional.ofNullable(this.ldapUserFederationId);
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
     * The id of the LDAP mapper implemented in MapperFactory.
     * 
     */
    @Import(name="providerId")
    private @Nullable Output<String> providerId;

    /**
     * @return The id of the LDAP mapper implemented in MapperFactory.
     * 
     */
    public Optional<Output<String>> providerId() {
        return Optional.ofNullable(this.providerId);
    }

    /**
     * The fully-qualified Java class name of the custom LDAP mapper.
     * 
     */
    @Import(name="providerType")
    private @Nullable Output<String> providerType;

    /**
     * @return The fully-qualified Java class name of the custom LDAP mapper.
     * 
     */
    public Optional<Output<String>> providerType() {
        return Optional.ofNullable(this.providerType);
    }

    /**
     * The realm that this LDAP mapper will exist in.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm that this LDAP mapper will exist in.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private CustomMapperState() {}

    private CustomMapperState(CustomMapperState $) {
        this.config = $.config;
        this.ldapUserFederationId = $.ldapUserFederationId;
        this.name = $.name;
        this.providerId = $.providerId;
        this.providerType = $.providerType;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CustomMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CustomMapperState $;

        public Builder() {
            $ = new CustomMapperState();
        }

        public Builder(CustomMapperState defaults) {
            $ = new CustomMapperState(Objects.requireNonNull(defaults));
        }

        /**
         * @param config A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<Map<String,Object>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,Object> config) {
            return config(Output.of(config));
        }

        /**
         * @param ldapUserFederationId The ID of the LDAP user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(@Nullable Output<String> ldapUserFederationId) {
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
         * @param providerId The id of the LDAP mapper implemented in MapperFactory.
         * 
         * @return builder
         * 
         */
        public Builder providerId(@Nullable Output<String> providerId) {
            $.providerId = providerId;
            return this;
        }

        /**
         * @param providerId The id of the LDAP mapper implemented in MapperFactory.
         * 
         * @return builder
         * 
         */
        public Builder providerId(String providerId) {
            return providerId(Output.of(providerId));
        }

        /**
         * @param providerType The fully-qualified Java class name of the custom LDAP mapper.
         * 
         * @return builder
         * 
         */
        public Builder providerType(@Nullable Output<String> providerType) {
            $.providerType = providerType;
            return this;
        }

        /**
         * @param providerType The fully-qualified Java class name of the custom LDAP mapper.
         * 
         * @return builder
         * 
         */
        public Builder providerType(String providerType) {
            return providerType(Output.of(providerType));
        }

        /**
         * @param realmId The realm that this LDAP mapper will exist in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
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

        public CustomMapperState build() {
            return $;
        }
    }

}
