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


public final class ClientUserPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClientUserPolicyArgs Empty = new ClientUserPolicyArgs();

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

    @Import(name="users", required=true)
    private Output<List<String>> users;

    public Output<List<String>> users() {
        return this.users;
    }

    private ClientUserPolicyArgs() {}

    private ClientUserPolicyArgs(ClientUserPolicyArgs $) {
        this.decisionStrategy = $.decisionStrategy;
        this.description = $.description;
        this.logic = $.logic;
        this.name = $.name;
        this.realmId = $.realmId;
        this.resourceServerId = $.resourceServerId;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientUserPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientUserPolicyArgs $;

        public Builder() {
            $ = new ClientUserPolicyArgs();
        }

        public Builder(ClientUserPolicyArgs defaults) {
            $ = new ClientUserPolicyArgs(Objects.requireNonNull(defaults));
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

        public Builder users(Output<List<String>> users) {
            $.users = users;
            return this;
        }

        public Builder users(List<String> users) {
            return users(Output.of(users));
        }

        public Builder users(String... users) {
            return users(List.of(users));
        }

        public ClientUserPolicyArgs build() {
            if ($.decisionStrategy == null) {
                throw new MissingRequiredPropertyException("ClientUserPolicyArgs", "decisionStrategy");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("ClientUserPolicyArgs", "realmId");
            }
            if ($.resourceServerId == null) {
                throw new MissingRequiredPropertyException("ClientUserPolicyArgs", "resourceServerId");
            }
            if ($.users == null) {
                throw new MissingRequiredPropertyException("ClientUserPolicyArgs", "users");
            }
            return $;
        }
    }

}
