// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.authentication;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FlowArgs extends com.pulumi.resources.ResourceArgs {

    public static final FlowArgs Empty = new FlowArgs();

    /**
     * The alias for this authentication flow.
     * 
     */
    @Import(name="alias", required=true)
    private Output<String> alias;

    /**
     * @return The alias for this authentication flow.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }

    /**
     * A description for the authentication flow.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the authentication flow.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
     * 
     */
    @Import(name="providerId")
    private @Nullable Output<String> providerId;

    /**
     * @return The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
     * 
     */
    public Optional<Output<String>> providerId() {
        return Optional.ofNullable(this.providerId);
    }

    /**
     * The realm that the authentication flow exists in.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm that the authentication flow exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private FlowArgs() {}

    private FlowArgs(FlowArgs $) {
        this.alias = $.alias;
        this.description = $.description;
        this.providerId = $.providerId;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FlowArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FlowArgs $;

        public Builder() {
            $ = new FlowArgs();
        }

        public Builder(FlowArgs defaults) {
            $ = new FlowArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias The alias for this authentication flow.
         * 
         * @return builder
         * 
         */
        public Builder alias(Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias The alias for this authentication flow.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param description A description for the authentication flow.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the authentication flow.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param providerId The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
         * 
         * @return builder
         * 
         */
        public Builder providerId(@Nullable Output<String> providerId) {
            $.providerId = providerId;
            return this;
        }

        /**
         * @param providerId The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
         * 
         * @return builder
         * 
         */
        public Builder providerId(String providerId) {
            return providerId(Output.of(providerId));
        }

        /**
         * @param realmId The realm that the authentication flow exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm that the authentication flow exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public FlowArgs build() {
            if ($.alias == null) {
                throw new MissingRequiredPropertyException("FlowArgs", "alias");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("FlowArgs", "realmId");
            }
            return $;
        }
    }

}
