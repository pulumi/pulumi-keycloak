// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.authentication.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExecutionState extends com.pulumi.resources.ResourceArgs {

    public static final ExecutionState Empty = new ExecutionState();

    /**
     * The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser&#39;s development tools.
     * 
     */
    @Import(name="authenticator")
    private @Nullable Output<String> authenticator;

    /**
     * @return The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser&#39;s development tools.
     * 
     */
    public Optional<Output<String>> authenticator() {
        return Optional.ofNullable(this.authenticator);
    }

    /**
     * The alias of the flow this execution is attached to.
     * 
     */
    @Import(name="parentFlowAlias")
    private @Nullable Output<String> parentFlowAlias;

    /**
     * @return The alias of the flow this execution is attached to.
     * 
     */
    public Optional<Output<String>> parentFlowAlias() {
        return Optional.ofNullable(this.parentFlowAlias);
    }

    /**
     * The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak &gt;= 25).
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak &gt;= 25).
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
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

    /**
     * The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
     * 
     */
    @Import(name="requirement")
    private @Nullable Output<String> requirement;

    /**
     * @return The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
     * 
     */
    public Optional<Output<String>> requirement() {
        return Optional.ofNullable(this.requirement);
    }

    private ExecutionState() {}

    private ExecutionState(ExecutionState $) {
        this.authenticator = $.authenticator;
        this.parentFlowAlias = $.parentFlowAlias;
        this.priority = $.priority;
        this.realmId = $.realmId;
        this.requirement = $.requirement;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExecutionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExecutionState $;

        public Builder() {
            $ = new ExecutionState();
        }

        public Builder(ExecutionState defaults) {
            $ = new ExecutionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authenticator The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser&#39;s development tools.
         * 
         * @return builder
         * 
         */
        public Builder authenticator(@Nullable Output<String> authenticator) {
            $.authenticator = authenticator;
            return this;
        }

        /**
         * @param authenticator The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser&#39;s development tools.
         * 
         * @return builder
         * 
         */
        public Builder authenticator(String authenticator) {
            return authenticator(Output.of(authenticator));
        }

        /**
         * @param parentFlowAlias The alias of the flow this execution is attached to.
         * 
         * @return builder
         * 
         */
        public Builder parentFlowAlias(@Nullable Output<String> parentFlowAlias) {
            $.parentFlowAlias = parentFlowAlias;
            return this;
        }

        /**
         * @param parentFlowAlias The alias of the flow this execution is attached to.
         * 
         * @return builder
         * 
         */
        public Builder parentFlowAlias(String parentFlowAlias) {
            return parentFlowAlias(Output.of(parentFlowAlias));
        }

        /**
         * @param priority The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak &gt;= 25).
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak &gt;= 25).
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
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

        /**
         * @param requirement The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
         * 
         * @return builder
         * 
         */
        public Builder requirement(@Nullable Output<String> requirement) {
            $.requirement = requirement;
            return this;
        }

        /**
         * @param requirement The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
         * 
         * @return builder
         * 
         */
        public Builder requirement(String requirement) {
            return requirement(Output.of(requirement));
        }

        public ExecutionState build() {
            return $;
        }
    }

}
