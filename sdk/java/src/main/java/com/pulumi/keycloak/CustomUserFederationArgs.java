// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CustomUserFederationArgs extends com.pulumi.resources.ResourceArgs {

    public static final CustomUserFederationArgs Empty = new CustomUserFederationArgs();

    /**
     * Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     * 
     */
    @Import(name="cachePolicy")
    private @Nullable Output<String> cachePolicy;

    /**
     * @return Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     * 
     */
    public Optional<Output<String>> cachePolicy() {
        return Optional.ofNullable(this.cachePolicy);
    }

    /**
     * How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
     * 
     */
    @Import(name="changedSyncPeriod")
    private @Nullable Output<Integer> changedSyncPeriod;

    /**
     * @return How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
     * 
     */
    public Optional<Output<Integer>> changedSyncPeriod() {
        return Optional.ofNullable(this.changedSyncPeriod);
    }

    /**
     * The provider configuration handed over to your custom user federation provider.
     * 
     */
    @Import(name="config")
    private @Nullable Output<Map<String,Object>> config;

    /**
     * @return The provider configuration handed over to your custom user federation provider.
     * 
     */
    public Optional<Output<Map<String,Object>>> config() {
        return Optional.ofNullable(this.config);
    }

    /**
     * When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
     * 
     */
    @Import(name="fullSyncPeriod")
    private @Nullable Output<Integer> fullSyncPeriod;

    /**
     * @return How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
     * 
     */
    public Optional<Output<Integer>> fullSyncPeriod() {
        return Optional.ofNullable(this.fullSyncPeriod);
    }

    /**
     * Display name of the provider when displayed in the console.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Display name of the provider when displayed in the console.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Must be set to the realms&#39; `internal_id`  when it differs from the realm. This can happen when existing resources are imported into the state.
     * 
     */
    @Import(name="parentId")
    private @Nullable Output<String> parentId;

    /**
     * @return Must be set to the realms&#39; `internal_id`  when it differs from the realm. This can happen when existing resources are imported into the state.
     * 
     */
    public Optional<Output<String>> parentId() {
        return Optional.ofNullable(this.parentId);
    }

    /**
     * Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
     * 
     */
    @Import(name="providerId", required=true)
    private Output<String> providerId;

    /**
     * @return The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
     * 
     */
    public Output<String> providerId() {
        return this.providerId;
    }

    /**
     * The realm that this provider will provide user federation for.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm that this provider will provide user federation for.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private CustomUserFederationArgs() {}

    private CustomUserFederationArgs(CustomUserFederationArgs $) {
        this.cachePolicy = $.cachePolicy;
        this.changedSyncPeriod = $.changedSyncPeriod;
        this.config = $.config;
        this.enabled = $.enabled;
        this.fullSyncPeriod = $.fullSyncPeriod;
        this.name = $.name;
        this.parentId = $.parentId;
        this.priority = $.priority;
        this.providerId = $.providerId;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CustomUserFederationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CustomUserFederationArgs $;

        public Builder() {
            $ = new CustomUserFederationArgs();
        }

        public Builder(CustomUserFederationArgs defaults) {
            $ = new CustomUserFederationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cachePolicy Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
         * 
         * @return builder
         * 
         */
        public Builder cachePolicy(@Nullable Output<String> cachePolicy) {
            $.cachePolicy = cachePolicy;
            return this;
        }

        /**
         * @param cachePolicy Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
         * 
         * @return builder
         * 
         */
        public Builder cachePolicy(String cachePolicy) {
            return cachePolicy(Output.of(cachePolicy));
        }

        /**
         * @param changedSyncPeriod How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
         * 
         * @return builder
         * 
         */
        public Builder changedSyncPeriod(@Nullable Output<Integer> changedSyncPeriod) {
            $.changedSyncPeriod = changedSyncPeriod;
            return this;
        }

        /**
         * @param changedSyncPeriod How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
         * 
         * @return builder
         * 
         */
        public Builder changedSyncPeriod(Integer changedSyncPeriod) {
            return changedSyncPeriod(Output.of(changedSyncPeriod));
        }

        /**
         * @param config The provider configuration handed over to your custom user federation provider.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<Map<String,Object>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config The provider configuration handed over to your custom user federation provider.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,Object> config) {
            return config(Output.of(config));
        }

        /**
         * @param enabled When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param fullSyncPeriod How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
         * 
         * @return builder
         * 
         */
        public Builder fullSyncPeriod(@Nullable Output<Integer> fullSyncPeriod) {
            $.fullSyncPeriod = fullSyncPeriod;
            return this;
        }

        /**
         * @param fullSyncPeriod How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
         * 
         * @return builder
         * 
         */
        public Builder fullSyncPeriod(Integer fullSyncPeriod) {
            return fullSyncPeriod(Output.of(fullSyncPeriod));
        }

        /**
         * @param name Display name of the provider when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Display name of the provider when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentId Must be set to the realms&#39; `internal_id`  when it differs from the realm. This can happen when existing resources are imported into the state.
         * 
         * @return builder
         * 
         */
        public Builder parentId(@Nullable Output<String> parentId) {
            $.parentId = parentId;
            return this;
        }

        /**
         * @param parentId Must be set to the realms&#39; `internal_id`  when it differs from the realm. This can happen when existing resources are imported into the state.
         * 
         * @return builder
         * 
         */
        public Builder parentId(String parentId) {
            return parentId(Output.of(parentId));
        }

        /**
         * @param priority Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param providerId The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
         * 
         * @return builder
         * 
         */
        public Builder providerId(Output<String> providerId) {
            $.providerId = providerId;
            return this;
        }

        /**
         * @param providerId The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
         * 
         * @return builder
         * 
         */
        public Builder providerId(String providerId) {
            return providerId(Output.of(providerId));
        }

        /**
         * @param realmId The realm that this provider will provide user federation for.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm that this provider will provide user federation for.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public CustomUserFederationArgs build() {
            $.providerId = Objects.requireNonNull($.providerId, "expected parameter 'providerId' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}
