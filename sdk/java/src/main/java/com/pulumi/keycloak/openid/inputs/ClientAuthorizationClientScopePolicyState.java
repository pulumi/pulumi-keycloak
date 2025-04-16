// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.keycloak.openid.inputs.ClientAuthorizationClientScopePolicyScopeArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientAuthorizationClientScopePolicyState extends com.pulumi.resources.ResourceArgs {

    public static final ClientAuthorizationClientScopePolicyState Empty = new ClientAuthorizationClientScopePolicyState();

    @Import(name="decisionStrategy")
    private @Nullable Output<String> decisionStrategy;

    public Optional<Output<String>> decisionStrategy() {
        return Optional.ofNullable(this.decisionStrategy);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="logic")
    private @Nullable Output<String> logic;

    public Optional<Output<String>> logic() {
        return Optional.ofNullable(this.logic);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    @Import(name="resourceServerId")
    private @Nullable Output<String> resourceServerId;

    public Optional<Output<String>> resourceServerId() {
        return Optional.ofNullable(this.resourceServerId);
    }

    @Import(name="scopes")
    private @Nullable Output<List<ClientAuthorizationClientScopePolicyScopeArgs>> scopes;

    public Optional<Output<List<ClientAuthorizationClientScopePolicyScopeArgs>>> scopes() {
        return Optional.ofNullable(this.scopes);
    }

    private ClientAuthorizationClientScopePolicyState() {}

    private ClientAuthorizationClientScopePolicyState(ClientAuthorizationClientScopePolicyState $) {
        this.decisionStrategy = $.decisionStrategy;
        this.description = $.description;
        this.logic = $.logic;
        this.name = $.name;
        this.realmId = $.realmId;
        this.resourceServerId = $.resourceServerId;
        this.scopes = $.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientAuthorizationClientScopePolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientAuthorizationClientScopePolicyState $;

        public Builder() {
            $ = new ClientAuthorizationClientScopePolicyState();
        }

        public Builder(ClientAuthorizationClientScopePolicyState defaults) {
            $ = new ClientAuthorizationClientScopePolicyState(Objects.requireNonNull(defaults));
        }

        public Builder decisionStrategy(@Nullable Output<String> decisionStrategy) {
            $.decisionStrategy = decisionStrategy;
            return this;
        }

        public Builder decisionStrategy(String decisionStrategy) {
            return decisionStrategy(Output.of(decisionStrategy));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder logic(@Nullable Output<String> logic) {
            $.logic = logic;
            return this;
        }

        public Builder logic(String logic) {
            return logic(Output.of(logic));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder resourceServerId(@Nullable Output<String> resourceServerId) {
            $.resourceServerId = resourceServerId;
            return this;
        }

        public Builder resourceServerId(String resourceServerId) {
            return resourceServerId(Output.of(resourceServerId));
        }

        public Builder scopes(@Nullable Output<List<ClientAuthorizationClientScopePolicyScopeArgs>> scopes) {
            $.scopes = scopes;
            return this;
        }

        public Builder scopes(List<ClientAuthorizationClientScopePolicyScopeArgs> scopes) {
            return scopes(Output.of(scopes));
        }

        public Builder scopes(ClientAuthorizationClientScopePolicyScopeArgs... scopes) {
            return scopes(List.of(scopes));
        }

        public ClientAuthorizationClientScopePolicyState build() {
            return $;
        }
    }

}
