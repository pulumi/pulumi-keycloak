// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UsersPermissionsMapRolesScopeArgs extends com.pulumi.resources.ResourceArgs {

    public static final UsersPermissionsMapRolesScopeArgs Empty = new UsersPermissionsMapRolesScopeArgs();

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

    @Import(name="policies")
    private @Nullable Output<List<String>> policies;

    public Optional<Output<List<String>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    private UsersPermissionsMapRolesScopeArgs() {}

    private UsersPermissionsMapRolesScopeArgs(UsersPermissionsMapRolesScopeArgs $) {
        this.decisionStrategy = $.decisionStrategy;
        this.description = $.description;
        this.policies = $.policies;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UsersPermissionsMapRolesScopeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UsersPermissionsMapRolesScopeArgs $;

        public Builder() {
            $ = new UsersPermissionsMapRolesScopeArgs();
        }

        public Builder(UsersPermissionsMapRolesScopeArgs defaults) {
            $ = new UsersPermissionsMapRolesScopeArgs(Objects.requireNonNull(defaults));
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

        public Builder policies(@Nullable Output<List<String>> policies) {
            $.policies = policies;
            return this;
        }

        public Builder policies(List<String> policies) {
            return policies(Output.of(policies));
        }

        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }

        public UsersPermissionsMapRolesScopeArgs build() {
            return $;
        }
    }

}
