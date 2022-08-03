// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserPropertyProtocolMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserPropertyProtocolMapperArgs Empty = new UserPropertyProtocolMapperArgs();

    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     * 
     */
    @Import(name="addToAccessToken")
    private @Nullable Output<Boolean> addToAccessToken;

    /**
     * @return Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> addToAccessToken() {
        return Optional.ofNullable(this.addToAccessToken);
    }

    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     * 
     */
    @Import(name="addToIdToken")
    private @Nullable Output<Boolean> addToIdToken;

    /**
     * @return Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> addToIdToken() {
        return Optional.ofNullable(this.addToIdToken);
    }

    /**
     * Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     * 
     */
    @Import(name="addToUserinfo")
    private @Nullable Output<Boolean> addToUserinfo;

    /**
     * @return Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> addToUserinfo() {
        return Optional.ofNullable(this.addToUserinfo);
    }

    /**
     * The name of the claim to insert into a token.
     * 
     */
    @Import(name="claimName", required=true)
    private Output<String> claimName;

    /**
     * @return The name of the claim to insert into a token.
     * 
     */
    public Output<String> claimName() {
        return this.claimName;
    }

    /**
     * The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     * 
     */
    @Import(name="claimValueType")
    private @Nullable Output<String> claimValueType;

    /**
     * @return The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     * 
     */
    public Optional<Output<String>> claimValueType() {
        return Optional.ofNullable(this.claimValueType);
    }

    /**
     * The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified. `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
     * 
     */
    @Import(name="clientScopeId")
    private @Nullable Output<String> clientScopeId;

    /**
     * @return The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified. `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
     * 
     */
    public Optional<Output<String>> clientScopeId() {
        return Optional.ofNullable(this.clientScopeId);
    }

    /**
     * The display name of this protocol mapper in the GUI.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The display name of this protocol mapper in the GUI.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The realm this protocol mapper exists within.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this protocol mapper exists within.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * The built in user property (such as email) to map a claim for.
     * 
     */
    @Import(name="userProperty", required=true)
    private Output<String> userProperty;

    /**
     * @return The built in user property (such as email) to map a claim for.
     * 
     */
    public Output<String> userProperty() {
        return this.userProperty;
    }

    private UserPropertyProtocolMapperArgs() {}

    private UserPropertyProtocolMapperArgs(UserPropertyProtocolMapperArgs $) {
        this.addToAccessToken = $.addToAccessToken;
        this.addToIdToken = $.addToIdToken;
        this.addToUserinfo = $.addToUserinfo;
        this.claimName = $.claimName;
        this.claimValueType = $.claimValueType;
        this.clientId = $.clientId;
        this.clientScopeId = $.clientScopeId;
        this.name = $.name;
        this.realmId = $.realmId;
        this.userProperty = $.userProperty;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserPropertyProtocolMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserPropertyProtocolMapperArgs $;

        public Builder() {
            $ = new UserPropertyProtocolMapperArgs();
        }

        public Builder(UserPropertyProtocolMapperArgs defaults) {
            $ = new UserPropertyProtocolMapperArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addToAccessToken Indicates if the property should be added as a claim to the access token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(@Nullable Output<Boolean> addToAccessToken) {
            $.addToAccessToken = addToAccessToken;
            return this;
        }

        /**
         * @param addToAccessToken Indicates if the property should be added as a claim to the access token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(Boolean addToAccessToken) {
            return addToAccessToken(Output.of(addToAccessToken));
        }

        /**
         * @param addToIdToken Indicates if the property should be added as a claim to the id token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(@Nullable Output<Boolean> addToIdToken) {
            $.addToIdToken = addToIdToken;
            return this;
        }

        /**
         * @param addToIdToken Indicates if the property should be added as a claim to the id token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(Boolean addToIdToken) {
            return addToIdToken(Output.of(addToIdToken));
        }

        /**
         * @param addToUserinfo Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToUserinfo(@Nullable Output<Boolean> addToUserinfo) {
            $.addToUserinfo = addToUserinfo;
            return this;
        }

        /**
         * @param addToUserinfo Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToUserinfo(Boolean addToUserinfo) {
            return addToUserinfo(Output.of(addToUserinfo));
        }

        /**
         * @param claimName The name of the claim to insert into a token.
         * 
         * @return builder
         * 
         */
        public Builder claimName(Output<String> claimName) {
            $.claimName = claimName;
            return this;
        }

        /**
         * @param claimName The name of the claim to insert into a token.
         * 
         * @return builder
         * 
         */
        public Builder claimName(String claimName) {
            return claimName(Output.of(claimName));
        }

        /**
         * @param claimValueType The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
         * 
         * @return builder
         * 
         */
        public Builder claimValueType(@Nullable Output<String> claimValueType) {
            $.claimValueType = claimValueType;
            return this;
        }

        /**
         * @param claimValueType The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
         * 
         * @return builder
         * 
         */
        public Builder claimValueType(String claimValueType) {
            return claimValueType(Output.of(claimValueType));
        }

        /**
         * @param clientId The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param clientScopeId The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified. `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(@Nullable Output<String> clientScopeId) {
            $.clientScopeId = clientScopeId;
            return this;
        }

        /**
         * @param clientScopeId The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified. `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(String clientScopeId) {
            return clientScopeId(Output.of(clientScopeId));
        }

        /**
         * @param name The display name of this protocol mapper in the GUI.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The display name of this protocol mapper in the GUI.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm this protocol mapper exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this protocol mapper exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param userProperty The built in user property (such as email) to map a claim for.
         * 
         * @return builder
         * 
         */
        public Builder userProperty(Output<String> userProperty) {
            $.userProperty = userProperty;
            return this;
        }

        /**
         * @param userProperty The built in user property (such as email) to map a claim for.
         * 
         * @return builder
         * 
         */
        public Builder userProperty(String userProperty) {
            return userProperty(Output.of(userProperty));
        }

        public UserPropertyProtocolMapperArgs build() {
            $.claimName = Objects.requireNonNull($.claimName, "expected parameter 'claimName' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            $.userProperty = Objects.requireNonNull($.userProperty, "expected parameter 'userProperty' to be non-null");
            return $;
        }
    }

}
