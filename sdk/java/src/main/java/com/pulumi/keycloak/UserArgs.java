// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.inputs.UserFederatedIdentityArgs;
import com.pulumi.keycloak.inputs.UserInitialPasswordArgs;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserArgs Empty = new UserArgs();

    /**
     * A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     * 
     */
    @Import(name="attributes")
    private @Nullable Output<Map<String,Object>> attributes;

    /**
     * @return A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     * 
     */
    public Optional<Output<Map<String,Object>>> attributes() {
        return Optional.ofNullable(this.attributes);
    }

    /**
     * The user&#39;s email.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return The user&#39;s email.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * Whether the email address was validated or not. Default to `false`.
     * 
     */
    @Import(name="emailVerified")
    private @Nullable Output<Boolean> emailVerified;

    /**
     * @return Whether the email address was validated or not. Default to `false`.
     * 
     */
    public Optional<Output<Boolean>> emailVerified() {
        return Optional.ofNullable(this.emailVerified);
    }

    /**
     * When false, this user cannot log in. Defaults to `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return When false, this user cannot log in. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="federatedIdentities")
    private @Nullable Output<List<UserFederatedIdentityArgs>> federatedIdentities;

    public Optional<Output<List<UserFederatedIdentityArgs>>> federatedIdentities() {
        return Optional.ofNullable(this.federatedIdentities);
    }

    /**
     * The user&#39;s first name.
     * 
     */
    @Import(name="firstName")
    private @Nullable Output<String> firstName;

    /**
     * @return The user&#39;s first name.
     * 
     */
    public Optional<Output<String>> firstName() {
        return Optional.ofNullable(this.firstName);
    }

    /**
     * When given, the user&#39;s initial password will be set. This attribute is only respected during initial user creation.
     * 
     */
    @Import(name="initialPassword")
    private @Nullable Output<UserInitialPasswordArgs> initialPassword;

    /**
     * @return When given, the user&#39;s initial password will be set. This attribute is only respected during initial user creation.
     * 
     */
    public Optional<Output<UserInitialPasswordArgs>> initialPassword() {
        return Optional.ofNullable(this.initialPassword);
    }

    /**
     * The user&#39;s last name.
     * 
     */
    @Import(name="lastName")
    private @Nullable Output<String> lastName;

    /**
     * @return The user&#39;s last name.
     * 
     */
    public Optional<Output<String>> lastName() {
        return Optional.ofNullable(this.lastName);
    }

    /**
     * The realm this user belongs to.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this user belongs to.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * A list of required user actions.
     * 
     */
    @Import(name="requiredActions")
    private @Nullable Output<List<String>> requiredActions;

    /**
     * @return A list of required user actions.
     * 
     */
    public Optional<Output<List<String>>> requiredActions() {
        return Optional.ofNullable(this.requiredActions);
    }

    /**
     * The unique username of this user.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The unique username of this user.
     * 
     */
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

        /**
         * @param attributes A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
         * 
         * @return builder
         * 
         */
        public Builder attributes(@Nullable Output<Map<String,Object>> attributes) {
            $.attributes = attributes;
            return this;
        }

        /**
         * @param attributes A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
         * 
         * @return builder
         * 
         */
        public Builder attributes(Map<String,Object> attributes) {
            return attributes(Output.of(attributes));
        }

        /**
         * @param email The user&#39;s email.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The user&#39;s email.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param emailVerified Whether the email address was validated or not. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder emailVerified(@Nullable Output<Boolean> emailVerified) {
            $.emailVerified = emailVerified;
            return this;
        }

        /**
         * @param emailVerified Whether the email address was validated or not. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder emailVerified(Boolean emailVerified) {
            return emailVerified(Output.of(emailVerified));
        }

        /**
         * @param enabled When false, this user cannot log in. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled When false, this user cannot log in. Defaults to `true`.
         * 
         * @return builder
         * 
         */
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

        /**
         * @param firstName The user&#39;s first name.
         * 
         * @return builder
         * 
         */
        public Builder firstName(@Nullable Output<String> firstName) {
            $.firstName = firstName;
            return this;
        }

        /**
         * @param firstName The user&#39;s first name.
         * 
         * @return builder
         * 
         */
        public Builder firstName(String firstName) {
            return firstName(Output.of(firstName));
        }

        /**
         * @param initialPassword When given, the user&#39;s initial password will be set. This attribute is only respected during initial user creation.
         * 
         * @return builder
         * 
         */
        public Builder initialPassword(@Nullable Output<UserInitialPasswordArgs> initialPassword) {
            $.initialPassword = initialPassword;
            return this;
        }

        /**
         * @param initialPassword When given, the user&#39;s initial password will be set. This attribute is only respected during initial user creation.
         * 
         * @return builder
         * 
         */
        public Builder initialPassword(UserInitialPasswordArgs initialPassword) {
            return initialPassword(Output.of(initialPassword));
        }

        /**
         * @param lastName The user&#39;s last name.
         * 
         * @return builder
         * 
         */
        public Builder lastName(@Nullable Output<String> lastName) {
            $.lastName = lastName;
            return this;
        }

        /**
         * @param lastName The user&#39;s last name.
         * 
         * @return builder
         * 
         */
        public Builder lastName(String lastName) {
            return lastName(Output.of(lastName));
        }

        /**
         * @param realmId The realm this user belongs to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this user belongs to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param requiredActions A list of required user actions.
         * 
         * @return builder
         * 
         */
        public Builder requiredActions(@Nullable Output<List<String>> requiredActions) {
            $.requiredActions = requiredActions;
            return this;
        }

        /**
         * @param requiredActions A list of required user actions.
         * 
         * @return builder
         * 
         */
        public Builder requiredActions(List<String> requiredActions) {
            return requiredActions(Output.of(requiredActions));
        }

        /**
         * @param requiredActions A list of required user actions.
         * 
         * @return builder
         * 
         */
        public Builder requiredActions(String... requiredActions) {
            return requiredActions(List.of(requiredActions));
        }

        /**
         * @param username The unique username of this user.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The unique username of this user.
         * 
         * @return builder
         * 
         */
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
