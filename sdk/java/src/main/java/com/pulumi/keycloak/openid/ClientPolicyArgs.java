// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClientPolicyArgs Empty = new ClientPolicyArgs();

    /**
     * The clients allowed by this client policy.
     * 
     */
    @Import(name="clients", required=true)
    private Output<List<String>> clients;

    /**
     * @return The clients allowed by this client policy.
     * 
     */
    public Output<List<String>> clients() {
        return this.clients;
    }

    /**
     * (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
     * 
     */
    @Import(name="decisionStrategy")
    private @Nullable Output<String> decisionStrategy;

    /**
     * @return (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
     * 
     */
    public Optional<Output<String>> decisionStrategy() {
        return Optional.ofNullable(this.decisionStrategy);
    }

    /**
     * The description of this client policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of this client policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
     * 
     */
    @Import(name="logic")
    private @Nullable Output<String> logic;

    /**
     * @return (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
     * 
     */
    public Optional<Output<String>> logic() {
        return Optional.ofNullable(this.logic);
    }

    /**
     * The name of this client policy.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of this client policy.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The realm this client policy exists within.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this client policy exists within.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * The ID of the resource server this client policy is attached to.
     * 
     */
    @Import(name="resourceServerId", required=true)
    private Output<String> resourceServerId;

    /**
     * @return The ID of the resource server this client policy is attached to.
     * 
     */
    public Output<String> resourceServerId() {
        return this.resourceServerId;
    }

    private ClientPolicyArgs() {}

    private ClientPolicyArgs(ClientPolicyArgs $) {
        this.clients = $.clients;
        this.decisionStrategy = $.decisionStrategy;
        this.description = $.description;
        this.logic = $.logic;
        this.name = $.name;
        this.realmId = $.realmId;
        this.resourceServerId = $.resourceServerId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientPolicyArgs $;

        public Builder() {
            $ = new ClientPolicyArgs();
        }

        public Builder(ClientPolicyArgs defaults) {
            $ = new ClientPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clients The clients allowed by this client policy.
         * 
         * @return builder
         * 
         */
        public Builder clients(Output<List<String>> clients) {
            $.clients = clients;
            return this;
        }

        /**
         * @param clients The clients allowed by this client policy.
         * 
         * @return builder
         * 
         */
        public Builder clients(List<String> clients) {
            return clients(Output.of(clients));
        }

        /**
         * @param clients The clients allowed by this client policy.
         * 
         * @return builder
         * 
         */
        public Builder clients(String... clients) {
            return clients(List.of(clients));
        }

        /**
         * @param decisionStrategy (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
         * 
         * @return builder
         * 
         */
        public Builder decisionStrategy(@Nullable Output<String> decisionStrategy) {
            $.decisionStrategy = decisionStrategy;
            return this;
        }

        /**
         * @param decisionStrategy (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
         * 
         * @return builder
         * 
         */
        public Builder decisionStrategy(String decisionStrategy) {
            return decisionStrategy(Output.of(decisionStrategy));
        }

        /**
         * @param description The description of this client policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of this client policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param logic (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
         * 
         * @return builder
         * 
         */
        public Builder logic(@Nullable Output<String> logic) {
            $.logic = logic;
            return this;
        }

        /**
         * @param logic (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
         * 
         * @return builder
         * 
         */
        public Builder logic(String logic) {
            return logic(Output.of(logic));
        }

        /**
         * @param name The name of this client policy.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of this client policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm this client policy exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this client policy exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param resourceServerId The ID of the resource server this client policy is attached to.
         * 
         * @return builder
         * 
         */
        public Builder resourceServerId(Output<String> resourceServerId) {
            $.resourceServerId = resourceServerId;
            return this;
        }

        /**
         * @param resourceServerId The ID of the resource server this client policy is attached to.
         * 
         * @return builder
         * 
         */
        public Builder resourceServerId(String resourceServerId) {
            return resourceServerId(Output.of(resourceServerId));
        }

        public ClientPolicyArgs build() {
            if ($.clients == null) {
                throw new MissingRequiredPropertyException("ClientPolicyArgs", "clients");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("ClientPolicyArgs", "realmId");
            }
            if ($.resourceServerId == null) {
                throw new MissingRequiredPropertyException("ClientPolicyArgs", "resourceServerId");
            }
            return $;
        }
    }

}
