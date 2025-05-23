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


public final class HardcodedGroupIdentityProviderMapperState extends com.pulumi.resources.ResourceArgs {

    public static final HardcodedGroupIdentityProviderMapperState Empty = new HardcodedGroupIdentityProviderMapperState();

    @Import(name="extraConfig")
    private @Nullable Output<Map<String,String>> extraConfig;

    public Optional<Output<Map<String,String>>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    /**
     * The name of the group which should be assigned to the users.
     * 
     */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return The name of the group which should be assigned to the users.
     * 
     */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
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

    private HardcodedGroupIdentityProviderMapperState() {}

    private HardcodedGroupIdentityProviderMapperState(HardcodedGroupIdentityProviderMapperState $) {
        this.extraConfig = $.extraConfig;
        this.group = $.group;
        this.identityProviderAlias = $.identityProviderAlias;
        this.name = $.name;
        this.realm = $.realm;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HardcodedGroupIdentityProviderMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HardcodedGroupIdentityProviderMapperState $;

        public Builder() {
            $ = new HardcodedGroupIdentityProviderMapperState();
        }

        public Builder(HardcodedGroupIdentityProviderMapperState defaults) {
            $ = new HardcodedGroupIdentityProviderMapperState(Objects.requireNonNull(defaults));
        }

        public Builder extraConfig(@Nullable Output<Map<String,String>> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder extraConfig(Map<String,String> extraConfig) {
            return extraConfig(Output.of(extraConfig));
        }

        /**
         * @param group The name of the group which should be assigned to the users.
         * 
         * @return builder
         * 
         */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The name of the group which should be assigned to the users.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
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

        public HardcodedGroupIdentityProviderMapperState build() {
            return $;
        }
    }

}
