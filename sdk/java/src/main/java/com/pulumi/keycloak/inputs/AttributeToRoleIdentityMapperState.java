// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AttributeToRoleIdentityMapperState extends com.pulumi.resources.ResourceArgs {

    public static final AttributeToRoleIdentityMapperState Empty = new AttributeToRoleIdentityMapperState();

    /**
     * Attribute Friendly Name. Conflicts with `attribute_name`.
     * 
     */
    @Import(name="attributeFriendlyName")
    private @Nullable Output<String> attributeFriendlyName;

    /**
     * @return Attribute Friendly Name. Conflicts with `attribute_name`.
     * 
     */
    public Optional<Output<String>> attributeFriendlyName() {
        return Optional.ofNullable(this.attributeFriendlyName);
    }

    /**
     * Attribute Name.
     * 
     */
    @Import(name="attributeName")
    private @Nullable Output<String> attributeName;

    /**
     * @return Attribute Name.
     * 
     */
    public Optional<Output<String>> attributeName() {
        return Optional.ofNullable(this.attributeName);
    }

    /**
     * Attribute Value.
     * 
     */
    @Import(name="attributeValue")
    private @Nullable Output<String> attributeValue;

    /**
     * @return Attribute Value.
     * 
     */
    public Optional<Output<String>> attributeValue() {
        return Optional.ofNullable(this.attributeValue);
    }

    /**
     * OIDC Claim Name
     * 
     */
    @Import(name="claimName")
    private @Nullable Output<String> claimName;

    /**
     * @return OIDC Claim Name
     * 
     */
    public Optional<Output<String>> claimName() {
        return Optional.ofNullable(this.claimName);
    }

    /**
     * OIDC Claim Value
     * 
     */
    @Import(name="claimValue")
    private @Nullable Output<String> claimValue;

    /**
     * @return OIDC Claim Value
     * 
     */
    public Optional<Output<String>> claimValue() {
        return Optional.ofNullable(this.claimValue);
    }

    /**
     * Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     * 
     */
    @Import(name="extraConfig")
    private @Nullable Output<Map<String,String>> extraConfig;

    /**
     * @return Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     * 
     */
    public Optional<Output<Map<String,String>>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    /**
     * The alias of the associated identity provider.
     * 
     */
    @Import(name="identityProviderAlias")
    private @Nullable Output<String> identityProviderAlias;

    /**
     * @return The alias of the associated identity provider.
     * 
     */
    public Optional<Output<String>> identityProviderAlias() {
        return Optional.ofNullable(this.identityProviderAlias);
    }

    /**
     * The name of the mapper.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the mapper.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the realm.
     * 
     */
    @Import(name="realm")
    private @Nullable Output<String> realm;

    /**
     * @return The name of the realm.
     * 
     */
    public Optional<Output<String>> realm() {
        return Optional.ofNullable(this.realm);
    }

    /**
     * Role Name.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return Role Name.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    private AttributeToRoleIdentityMapperState() {}

    private AttributeToRoleIdentityMapperState(AttributeToRoleIdentityMapperState $) {
        this.attributeFriendlyName = $.attributeFriendlyName;
        this.attributeName = $.attributeName;
        this.attributeValue = $.attributeValue;
        this.claimName = $.claimName;
        this.claimValue = $.claimValue;
        this.extraConfig = $.extraConfig;
        this.identityProviderAlias = $.identityProviderAlias;
        this.name = $.name;
        this.realm = $.realm;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AttributeToRoleIdentityMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AttributeToRoleIdentityMapperState $;

        public Builder() {
            $ = new AttributeToRoleIdentityMapperState();
        }

        public Builder(AttributeToRoleIdentityMapperState defaults) {
            $ = new AttributeToRoleIdentityMapperState(Objects.requireNonNull(defaults));
        }

        /**
         * @param attributeFriendlyName Attribute Friendly Name. Conflicts with `attribute_name`.
         * 
         * @return builder
         * 
         */
        public Builder attributeFriendlyName(@Nullable Output<String> attributeFriendlyName) {
            $.attributeFriendlyName = attributeFriendlyName;
            return this;
        }

        /**
         * @param attributeFriendlyName Attribute Friendly Name. Conflicts with `attribute_name`.
         * 
         * @return builder
         * 
         */
        public Builder attributeFriendlyName(String attributeFriendlyName) {
            return attributeFriendlyName(Output.of(attributeFriendlyName));
        }

        /**
         * @param attributeName Attribute Name.
         * 
         * @return builder
         * 
         */
        public Builder attributeName(@Nullable Output<String> attributeName) {
            $.attributeName = attributeName;
            return this;
        }

        /**
         * @param attributeName Attribute Name.
         * 
         * @return builder
         * 
         */
        public Builder attributeName(String attributeName) {
            return attributeName(Output.of(attributeName));
        }

        /**
         * @param attributeValue Attribute Value.
         * 
         * @return builder
         * 
         */
        public Builder attributeValue(@Nullable Output<String> attributeValue) {
            $.attributeValue = attributeValue;
            return this;
        }

        /**
         * @param attributeValue Attribute Value.
         * 
         * @return builder
         * 
         */
        public Builder attributeValue(String attributeValue) {
            return attributeValue(Output.of(attributeValue));
        }

        /**
         * @param claimName OIDC Claim Name
         * 
         * @return builder
         * 
         */
        public Builder claimName(@Nullable Output<String> claimName) {
            $.claimName = claimName;
            return this;
        }

        /**
         * @param claimName OIDC Claim Name
         * 
         * @return builder
         * 
         */
        public Builder claimName(String claimName) {
            return claimName(Output.of(claimName));
        }

        /**
         * @param claimValue OIDC Claim Value
         * 
         * @return builder
         * 
         */
        public Builder claimValue(@Nullable Output<String> claimValue) {
            $.claimValue = claimValue;
            return this;
        }

        /**
         * @param claimValue OIDC Claim Value
         * 
         * @return builder
         * 
         */
        public Builder claimValue(String claimValue) {
            return claimValue(Output.of(claimValue));
        }

        /**
         * @param extraConfig Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
         * 
         * @return builder
         * 
         */
        public Builder extraConfig(@Nullable Output<Map<String,String>> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        /**
         * @param extraConfig Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
         * 
         * @return builder
         * 
         */
        public Builder extraConfig(Map<String,String> extraConfig) {
            return extraConfig(Output.of(extraConfig));
        }

        /**
         * @param identityProviderAlias The alias of the associated identity provider.
         * 
         * @return builder
         * 
         */
        public Builder identityProviderAlias(@Nullable Output<String> identityProviderAlias) {
            $.identityProviderAlias = identityProviderAlias;
            return this;
        }

        /**
         * @param identityProviderAlias The alias of the associated identity provider.
         * 
         * @return builder
         * 
         */
        public Builder identityProviderAlias(String identityProviderAlias) {
            return identityProviderAlias(Output.of(identityProviderAlias));
        }

        /**
         * @param name The name of the mapper.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the mapper.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realm The name of the realm.
         * 
         * @return builder
         * 
         */
        public Builder realm(@Nullable Output<String> realm) {
            $.realm = realm;
            return this;
        }

        /**
         * @param realm The name of the realm.
         * 
         * @return builder
         * 
         */
        public Builder realm(String realm) {
            return realm(Output.of(realm));
        }

        /**
         * @param role Role Name.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Role Name.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public AttributeToRoleIdentityMapperState build() {
            return $;
        }
    }

}
