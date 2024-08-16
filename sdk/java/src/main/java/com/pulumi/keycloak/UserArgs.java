// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.inputs.UserFederatedIdentityArgs;
import com.pulumi.keycloak.inputs.UserInitialPasswordArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserArgs Empty = new UserArgs();

    @Import(name="attributes")
    private @Nullable Output<Map<String,String>> attributes;

    public Optional<Output<Map<String,String>>> attributes() {
        return Optional.ofNullable(this.attributes);
    }

    @Import(name="email")
    private @Nullable Output<String> email;

    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    @Import(name="emailVerified")
    private @Nullable Output<Boolean> emailVerified;

    public Optional<Output<Boolean>> emailVerified() {
        return Optional.ofNullable(this.emailVerified);
    }

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="federatedIdentities")
    private @Nullable Output<List<UserFederatedIdentityArgs>> federatedIdentities;

    public Optional<Output<List<UserFederatedIdentityArgs>>> federatedIdentities() {
        return Optional.ofNullable(this.federatedIdentities);
    }

    @Import(name="firstName")
    private @Nullable Output<String> firstName;

    public Optional<Output<String>> firstName() {
        return Optional.ofNullable(this.firstName);
    }

    @Import(name="initialPassword")
    private @Nullable Output<UserInitialPasswordArgs> initialPassword;

    public Optional<Output<UserInitialPasswordArgs>> initialPassword() {
        return Optional.ofNullable(this.initialPassword);
    }

    @Import(name="lastName")
    private @Nullable Output<String> lastName;

    public Optional<Output<String>> lastName() {
        return Optional.ofNullable(this.lastName);
    }

    @Import(name="realmId", required=true)
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }

    @Import(name="requiredActions")
    private @Nullable Output<List<String>> requiredActions;

    public Optional<Output<List<String>>> requiredActions() {
        return Optional.ofNullable(this.requiredActions);
    }

    @Import(name="username", required=true)
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    private UserArgs() {}

    private UserArgs(UserArgs $) {
        this.attributes = $.attributes;
        this.email = $.email;
        this.emailVerified = $.emailVerified;
        this.enabled = $.enabled;
        this.federatedIdentities = $.federatedIdentities;
        this.firstName = $.firstName;
        this.initialPassword = $.initialPassword;
        this.lastName = $.lastName;
        this.realmId = $.realmId;
        this.requiredActions = $.requiredActions;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserArgs $;

        public Builder() {
            $ = new UserArgs();
        }

        public Builder(UserArgs defaults) {
            $ = new UserArgs(Objects.requireNonNull(defaults));
        }

        public Builder attributes(@Nullable Output<Map<String,String>> attributes) {
            $.attributes = attributes;
            return this;
        }

        public Builder attributes(Map<String,String> attributes) {
            return attributes(Output.of(attributes));
        }

        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        public Builder email(String email) {
            return email(Output.of(email));
        }

        public Builder emailVerified(@Nullable Output<Boolean> emailVerified) {
            $.emailVerified = emailVerified;
            return this;
        }

        public Builder emailVerified(Boolean emailVerified) {
            return emailVerified(Output.of(emailVerified));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder federatedIdentities(@Nullable Output<List<UserFederatedIdentityArgs>> federatedIdentities) {
            $.federatedIdentities = federatedIdentities;
            return this;
        }

        public Builder federatedIdentities(List<UserFederatedIdentityArgs> federatedIdentities) {
            return federatedIdentities(Output.of(federatedIdentities));
        }

        public Builder federatedIdentities(UserFederatedIdentityArgs... federatedIdentities) {
            return federatedIdentities(List.of(federatedIdentities));
        }

        public Builder firstName(@Nullable Output<String> firstName) {
            $.firstName = firstName;
            return this;
        }

        public Builder firstName(String firstName) {
            return firstName(Output.of(firstName));
        }

        public Builder initialPassword(@Nullable Output<UserInitialPasswordArgs> initialPassword) {
            $.initialPassword = initialPassword;
            return this;
        }

        public Builder initialPassword(UserInitialPasswordArgs initialPassword) {
            return initialPassword(Output.of(initialPassword));
        }

        public Builder lastName(@Nullable Output<String> lastName) {
            $.lastName = lastName;
            return this;
        }

        public Builder lastName(String lastName) {
            return lastName(Output.of(lastName));
        }

        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder requiredActions(@Nullable Output<List<String>> requiredActions) {
            $.requiredActions = requiredActions;
            return this;
        }

        public Builder requiredActions(List<String> requiredActions) {
            return requiredActions(Output.of(requiredActions));
        }

        public Builder requiredActions(String... requiredActions) {
            return requiredActions(List.of(requiredActions));
        }

        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public UserArgs build() {
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("UserArgs", "realmId");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("UserArgs", "username");
            }
            return $;
        }
    }

}
