// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.openid.inputs.ClientRolePolicyRoleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientRolePolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClientRolePolicyArgs Empty = new ClientRolePolicyArgs();

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

    @Import(name="roles", required=true)
    private Output<List<ClientRolePolicyRoleArgs>> roles;

    public Output<List<ClientRolePolicyRoleArgs>> roles() {
        return this.roles;
    }

    @Import(name="type", required=true)
    private Output<String> type;

    public Output<String> type() {
        return this.type;
    }

    private ClientRolePolicyArgs() {}

    private ClientRolePolicyArgs(ClientRolePolicyArgs $) {
        this.decisionStrategy = $.decisionStrategy;
        this.description = $.description;
        this.logic = $.logic;
        this.name = $.name;
        this.realmId = $.realmId;
        this.resourceServerId = $.resourceServerId;
        this.roles = $.roles;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientRolePolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientRolePolicyArgs $;

        public Builder() {
            $ = new ClientRolePolicyArgs();
        }

        public Builder(ClientRolePolicyArgs defaults) {
            $ = new ClientRolePolicyArgs(Objects.requireNonNull(defaults));
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

        public Builder roles(Output<List<ClientRolePolicyRoleArgs>> roles) {
            $.roles = roles;
            return this;
        }

        public Builder roles(List<ClientRolePolicyRoleArgs> roles) {
            return roles(Output.of(roles));
        }

        public Builder roles(ClientRolePolicyRoleArgs... roles) {
            return roles(List.of(roles));
        }

        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ClientRolePolicyArgs build() {
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("ClientRolePolicyArgs", "realmId");
            }
            if ($.resourceServerId == null) {
                throw new MissingRequiredPropertyException("ClientRolePolicyArgs", "resourceServerId");
            }
            if ($.roles == null) {
                throw new MissingRequiredPropertyException("ClientRolePolicyArgs", "roles");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("ClientRolePolicyArgs", "type");
            }
            return $;
        }
    }

}
