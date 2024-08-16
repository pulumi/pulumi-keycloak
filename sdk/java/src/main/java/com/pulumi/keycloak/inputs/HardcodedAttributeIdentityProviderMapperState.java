// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HardcodedAttributeIdentityProviderMapperState extends com.pulumi.resources.ResourceArgs {

    public static final HardcodedAttributeIdentityProviderMapperState Empty = new HardcodedAttributeIdentityProviderMapperState();

    /**
     * The name of the IDP attribute to set.
     * 
     */
    @Import(name="attributeName")
    private @Nullable Output<String> attributeName;

    /**
     * @return The name of the IDP attribute to set.
     * 
     */
    public Optional<Output<String>> attributeName() {
        return Optional.ofNullable(this.attributeName);
    }

    /**
     * The value to set to the attribute. You can hardcode any value like &#39;foo&#39;.
     * 
     */
    @Import(name="attributeValue")
    private @Nullable Output<String> attributeValue;

    /**
     * @return The value to set to the attribute. You can hardcode any value like &#39;foo&#39;.
     * 
     */
    public Optional<Output<String>> attributeValue() {
        return Optional.ofNullable(this.attributeValue);
    }

    @Import(name="extraConfig")
    private @Nullable Output<Map<String,String>> extraConfig;

    public Optional<Output<Map<String,String>>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    /**
     * The IDP alias of the attribute to set.
     * 
     */
    @Import(name="identityProviderAlias")
    private @Nullable Output<String> identityProviderAlias;

    /**
     * @return The IDP alias of the attribute to set.
     * 
     */
    public Optional<Output<String>> identityProviderAlias() {
        return Optional.ofNullable(this.identityProviderAlias);
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
     * The realm ID that this mapper will exist in.
     * 
     */
    @Import(name="realm")
    private @Nullable Output<String> realm;

    /**
     * @return The realm ID that this mapper will exist in.
     * 
     */
    public Optional<Output<String>> realm() {
        return Optional.ofNullable(this.realm);
    }

    /**
     * Is Attribute related to a User Session.
     * 
     */
    @Import(name="userSession")
    private @Nullable Output<Boolean> userSession;

    /**
     * @return Is Attribute related to a User Session.
     * 
     */
    public Optional<Output<Boolean>> userSession() {
        return Optional.ofNullable(this.userSession);
    }

    private HardcodedAttributeIdentityProviderMapperState() {}

    private HardcodedAttributeIdentityProviderMapperState(HardcodedAttributeIdentityProviderMapperState $) {
        this.attributeName = $.attributeName;
        this.attributeValue = $.attributeValue;
        this.extraConfig = $.extraConfig;
        this.identityProviderAlias = $.identityProviderAlias;
        this.name = $.name;
        this.realm = $.realm;
        this.userSession = $.userSession;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HardcodedAttributeIdentityProviderMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HardcodedAttributeIdentityProviderMapperState $;

        public Builder() {
            $ = new HardcodedAttributeIdentityProviderMapperState();
        }

        public Builder(HardcodedAttributeIdentityProviderMapperState defaults) {
            $ = new HardcodedAttributeIdentityProviderMapperState(Objects.requireNonNull(defaults));
        }

        /**
         * @param attributeName The name of the IDP attribute to set.
         * 
         * @return builder
         * 
         */
        public Builder attributeName(@Nullable Output<String> attributeName) {
            $.attributeName = attributeName;
            return this;
        }

        /**
         * @param attributeName The name of the IDP attribute to set.
         * 
         * @return builder
         * 
         */
        public Builder attributeName(String attributeName) {
            return attributeName(Output.of(attributeName));
        }

        /**
         * @param attributeValue The value to set to the attribute. You can hardcode any value like &#39;foo&#39;.
         * 
         * @return builder
         * 
         */
        public Builder attributeValue(@Nullable Output<String> attributeValue) {
            $.attributeValue = attributeValue;
            return this;
        }

        /**
         * @param attributeValue The value to set to the attribute. You can hardcode any value like &#39;foo&#39;.
         * 
         * @return builder
         * 
         */
        public Builder attributeValue(String attributeValue) {
            return attributeValue(Output.of(attributeValue));
        }

        public Builder extraConfig(@Nullable Output<Map<String,String>> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder extraConfig(Map<String,String> extraConfig) {
            return extraConfig(Output.of(extraConfig));
        }

        /**
         * @param identityProviderAlias The IDP alias of the attribute to set.
         * 
         * @return builder
         * 
         */
        public Builder identityProviderAlias(@Nullable Output<String> identityProviderAlias) {
            $.identityProviderAlias = identityProviderAlias;
            return this;
        }

        /**
         * @param identityProviderAlias The IDP alias of the attribute to set.
         * 
         * @return builder
         * 
         */
        public Builder identityProviderAlias(String identityProviderAlias) {
            return identityProviderAlias(Output.of(identityProviderAlias));
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
         * @param realm The realm ID that this mapper will exist in.
         * 
         * @return builder
         * 
         */
        public Builder realm(@Nullable Output<String> realm) {
            $.realm = realm;
            return this;
        }

        /**
         * @param realm The realm ID that this mapper will exist in.
         * 
         * @return builder
         * 
         */
        public Builder realm(String realm) {
            return realm(Output.of(realm));
        }

        /**
         * @param userSession Is Attribute related to a User Session.
         * 
         * @return builder
         * 
         */
        public Builder userSession(@Nullable Output<Boolean> userSession) {
            $.userSession = userSession;
            return this;
        }

        /**
         * @param userSession Is Attribute related to a User Session.
         * 
         * @return builder
         * 
         */
        public Builder userSession(Boolean userSession) {
            return userSession(Output.of(userSession));
        }

        public HardcodedAttributeIdentityProviderMapperState build() {
            return $;
        }
    }

}
