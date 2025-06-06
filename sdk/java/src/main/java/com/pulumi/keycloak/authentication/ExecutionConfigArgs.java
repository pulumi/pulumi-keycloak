// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.authentication;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExecutionConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ExecutionConfigArgs Empty = new ExecutionConfigArgs();

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
    @Import(name="config", required=true)
    private Output<Map<String,String>> config;

    /**
     * @return The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
     * 
     */
    public Output<Map<String,String>> config() {
        return this.config;
    }

    /**
     * The authentication execution this configuration is attached to.
     * 
     */
    @Import(name="executionId", required=true)
    private Output<String> executionId;

    /**
     * @return The authentication execution this configuration is attached to.
     * 
     */
    public Output<String> executionId() {
        return this.executionId;
    }

    /**
     * The realm the authentication execution exists in.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm the authentication execution exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private ExecutionConfigArgs() {}

    private ExecutionConfigArgs(ExecutionConfigArgs $) {
        this.alias = $.alias;
        this.config = $.config;
        this.executionId = $.executionId;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExecutionConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExecutionConfigArgs $;

        public Builder() {
            $ = new ExecutionConfigArgs();
        }

        public Builder(ExecutionConfigArgs defaults) {
            $ = new ExecutionConfigArgs(Objects.requireNonNull(defaults));
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
        public Builder config(Output<Map<String,String>> config) {
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
        public Builder executionId(Output<String> executionId) {
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
        public Builder realmId(Output<String> realmId) {
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

        public ExecutionConfigArgs build() {
            if ($.config == null) {
                throw new MissingRequiredPropertyException("ExecutionConfigArgs", "config");
            }
            if ($.executionId == null) {
                throw new MissingRequiredPropertyException("ExecutionConfigArgs", "executionId");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("ExecutionConfigArgs", "realmId");
            }
            return $;
        }
    }

}
