// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmDefaultClientScopesState extends com.pulumi.resources.ResourceArgs {

    public static final RealmDefaultClientScopesState Empty = new RealmDefaultClientScopesState();

    /**
     * An array of default client scope names that should be used when creating new Keycloak clients.
     * 
     */
    @Import(name="defaultScopes")
    private @Nullable Output<List<String>> defaultScopes;

    /**
     * @return An array of default client scope names that should be used when creating new Keycloak clients.
     * 
     */
    public Optional<Output<List<String>>> defaultScopes() {
        return Optional.ofNullable(this.defaultScopes);
    }

    /**
     * The realm this client and scopes exists in.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm this client and scopes exists in.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private RealmDefaultClientScopesState() {}

    private RealmDefaultClientScopesState(RealmDefaultClientScopesState $) {
        this.defaultScopes = $.defaultScopes;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmDefaultClientScopesState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmDefaultClientScopesState $;

        public Builder() {
            $ = new RealmDefaultClientScopesState();
        }

        public Builder(RealmDefaultClientScopesState defaults) {
            $ = new RealmDefaultClientScopesState(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultScopes An array of default client scope names that should be used when creating new Keycloak clients.
         * 
         * @return builder
         * 
         */
        public Builder defaultScopes(@Nullable Output<List<String>> defaultScopes) {
            $.defaultScopes = defaultScopes;
            return this;
        }

        /**
         * @param defaultScopes An array of default client scope names that should be used when creating new Keycloak clients.
         * 
         * @return builder
         * 
         */
        public Builder defaultScopes(List<String> defaultScopes) {
            return defaultScopes(Output.of(defaultScopes));
        }

        /**
         * @param defaultScopes An array of default client scope names that should be used when creating new Keycloak clients.
         * 
         * @return builder
         * 
         */
        public Builder defaultScopes(String... defaultScopes) {
            return defaultScopes(List.of(defaultScopes));
        }

        /**
         * @param realmId The realm this client and scopes exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this client and scopes exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public RealmDefaultClientScopesState build() {
            return $;
        }
    }

}
