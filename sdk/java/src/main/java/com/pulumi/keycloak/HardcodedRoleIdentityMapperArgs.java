// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HardcodedRoleIdentityMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final HardcodedRoleIdentityMapperArgs Empty = new HardcodedRoleIdentityMapperArgs();

    @Import(name="extraConfig")
    private @Nullable Output<Map<String,Object>> extraConfig;

    public Optional<Output<Map<String,Object>>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    /**
     * The IDP alias of the attribute to set.
     * 
     */
    @Import(name="identityProviderAlias", required=true)
    private Output<String> identityProviderAlias;

    /**
     * @return The IDP alias of the attribute to set.
     * 
     */
    public Output<String> identityProviderAlias() {
        return this.identityProviderAlias;
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
    @Import(name="realm", required=true)
    private Output<String> realm;

    /**
     * @return The realm ID that this mapper will exist in.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }

    /**
     * The name of the role which should be assigned to the users.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The name of the role which should be assigned to the users.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    private HardcodedRoleIdentityMapperArgs() {}

    private HardcodedRoleIdentityMapperArgs(HardcodedRoleIdentityMapperArgs $) {
        this.extraConfig = $.extraConfig;
        this.identityProviderAlias = $.identityProviderAlias;
        this.name = $.name;
        this.realm = $.realm;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HardcodedRoleIdentityMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HardcodedRoleIdentityMapperArgs $;

        public Builder() {
            $ = new HardcodedRoleIdentityMapperArgs();
        }

        public Builder(HardcodedRoleIdentityMapperArgs defaults) {
            $ = new HardcodedRoleIdentityMapperArgs(Objects.requireNonNull(defaults));
        }

        public Builder extraConfig(@Nullable Output<Map<String,Object>> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder extraConfig(Map<String,Object> extraConfig) {
            return extraConfig(Output.of(extraConfig));
        }

        /**
         * @param identityProviderAlias The IDP alias of the attribute to set.
         * 
         * @return builder
         * 
         */
        public Builder identityProviderAlias(Output<String> identityProviderAlias) {
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
        public Builder realm(Output<String> realm) {
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
         * @param role The name of the role which should be assigned to the users.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name of the role which should be assigned to the users.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public HardcodedRoleIdentityMapperArgs build() {
            $.identityProviderAlias = Objects.requireNonNull($.identityProviderAlias, "expected parameter 'identityProviderAlias' to be non-null");
            $.realm = Objects.requireNonNull($.realm, "expected parameter 'realm' to be non-null");
            return $;
        }
    }

}
