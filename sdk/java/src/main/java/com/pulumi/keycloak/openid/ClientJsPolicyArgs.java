// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientJsPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClientJsPolicyArgs Empty = new ClientJsPolicyArgs();

    @Import(name="code", required=true)
    private Output<String> code;

    public Output<String> code() {
        return this.code;
    }

    @Import(name="decisionStrategy", required=true)
    private Output<String> decisionStrategy;

    public Output<String> decisionStrategy() {
        return this.decisionStrategy;
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

    @Import(name="realmId", required=true)
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }

    @Import(name="resourceServerId", required=true)
    private Output<String> resourceServerId;

    public Output<String> resourceServerId() {
        return this.resourceServerId;
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ClientJsPolicyArgs() {}

    private ClientJsPolicyArgs(ClientJsPolicyArgs $) {
        this.code = $.code;
        this.decisionStrategy = $.decisionStrategy;
        this.description = $.description;
        this.logic = $.logic;
        this.name = $.name;
        this.realmId = $.realmId;
        this.resourceServerId = $.resourceServerId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientJsPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientJsPolicyArgs $;

        public Builder() {
            $ = new ClientJsPolicyArgs();
        }

        public Builder(ClientJsPolicyArgs defaults) {
            $ = new ClientJsPolicyArgs(Objects.requireNonNull(defaults));
        }

        public Builder code(Output<String> code) {
            $.code = code;
            return this;
        }

        public Builder code(String code) {
            return code(Output.of(code));
        }

        public Builder decisionStrategy(Output<String> decisionStrategy) {
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

        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder resourceServerId(Output<String> resourceServerId) {
            $.resourceServerId = resourceServerId;
            return this;
        }

        public Builder resourceServerId(String resourceServerId) {
            return resourceServerId(Output.of(resourceServerId));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ClientJsPolicyArgs build() {
            $.code = Objects.requireNonNull($.code, "expected parameter 'code' to be non-null");
            $.decisionStrategy = Objects.requireNonNull($.decisionStrategy, "expected parameter 'decisionStrategy' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            $.resourceServerId = Objects.requireNonNull($.resourceServerId, "expected parameter 'resourceServerId' to be non-null");
            return $;
        }
    }

}
