// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.authentication.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExecutionConfigState extends com.pulumi.resources.ResourceArgs {

    public static final ExecutionConfigState Empty = new ExecutionConfigState();

    /**
     * The name of the configuration.
     * 
     */
    @Import(name="alias")
    private @Nullable Output<String> alias;

    /**
     * @return The name of the configuration.
     * 
     */
    public Optional<Output<String>> alias() {
        return Optional.ofNullable(this.alias);
    }

    /**
     * The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
     * 
     */
    @Import(name="config")
    private @Nullable Output<Map<String,String>> config;

    /**
     * @return The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
     * 
     */
    public Optional<Output<Map<String,String>>> config() {
        return Optional.ofNullable(this.config);
    }

    /**
     * The authentication execution this configuration is attached to.
     * 
     */
    @Import(name="executionId")
    private @Nullable Output<String> executionId;

    /**
     * @return The authentication execution this configuration is attached to.
     * 
     */
    public Optional<Output<String>> executionId() {
        return Optional.ofNullable(this.executionId);
    }

    /**
     * The realm the authentication execution exists in.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm the authentication execution exists in.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private ExecutionConfigState() {}

    private ExecutionConfigState(ExecutionConfigState $) {
        this.alias = $.alias;
        this.config = $.config;
        this.executionId = $.executionId;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExecutionConfigState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExecutionConfigState $;

        public Builder() {
            $ = new ExecutionConfigState();
        }

        public Builder(ExecutionConfigState defaults) {
            $ = new ExecutionConfigState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias The name of the configuration.
         * 
         * @return builder
         * 
         */
        public Builder alias(@Nullable Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias The name of the configuration.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param config The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<Map<String,String>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,String> config) {
            return config(Output.of(config));
        }

        /**
         * @param executionId The authentication execution this configuration is attached to.
         * 
         * @return builder
         * 
         */
        public Builder executionId(@Nullable Output<String> executionId) {
            $.executionId = executionId;
            return this;
        }

        /**
         * @param executionId The authentication execution this configuration is attached to.
         * 
         * @return builder
         * 
         */
        public Builder executionId(String executionId) {
            return executionId(Output.of(executionId));
        }

        /**
         * @param realmId The realm the authentication execution exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm the authentication execution exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public ExecutionConfigState build() {
            return $;
        }
    }

}
