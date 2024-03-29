// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserAttributeProtocolMapperState extends com.pulumi.resources.ResourceArgs {

    public static final UserAttributeProtocolMapperState Empty = new UserAttributeProtocolMapperState();

    /**
     * Indicates if the attribute should be a claim in the access token.
     * 
     */
    @Import(name="addToAccessToken")
    private @Nullable Output<Boolean> addToAccessToken;

    /**
     * @return Indicates if the attribute should be a claim in the access token.
     * 
     */
    public Optional<Output<Boolean>> addToAccessToken() {
        return Optional.ofNullable(this.addToAccessToken);
    }

    /**
     * Indicates if the attribute should be a claim in the id token.
     * 
     */
    @Import(name="addToIdToken")
    private @Nullable Output<Boolean> addToIdToken;

    /**
     * @return Indicates if the attribute should be a claim in the id token.
     * 
     */
    public Optional<Output<Boolean>> addToIdToken() {
        return Optional.ofNullable(this.addToIdToken);
    }

    /**
     * Indicates if the attribute should appear in the userinfo response body.
     * 
     */
    @Import(name="addToUserinfo")
    private @Nullable Output<Boolean> addToUserinfo;

    /**
     * @return Indicates if the attribute should appear in the userinfo response body.
     * 
     */
    public Optional<Output<Boolean>> addToUserinfo() {
        return Optional.ofNullable(this.addToUserinfo);
    }

    /**
     * Indicates if attribute values should be aggregated within the group attributes
     * 
     */
    @Import(name="aggregateAttributes")
    private @Nullable Output<Boolean> aggregateAttributes;

    /**
     * @return Indicates if attribute values should be aggregated within the group attributes
     * 
     */
    public Optional<Output<Boolean>> aggregateAttributes() {
        return Optional.ofNullable(this.aggregateAttributes);
    }

    @Import(name="claimName")
    private @Nullable Output<String> claimName;

    public Optional<Output<String>> claimName() {
        return Optional.ofNullable(this.claimName);
    }

    /**
     * Claim type used when serializing tokens.
     * 
     */
    @Import(name="claimValueType")
    private @Nullable Output<String> claimValueType;

    /**
     * @return Claim type used when serializing tokens.
     * 
     */
    public Optional<Output<String>> claimValueType() {
        return Optional.ofNullable(this.claimValueType);
    }

    /**
     * The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
     * 
     */
    @Import(name="clientScopeId")
    private @Nullable Output<String> clientScopeId;

    /**
     * @return The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
     * 
     */
    public Optional<Output<String>> clientScopeId() {
        return Optional.ofNullable(this.clientScopeId);
    }

    /**
     * Indicates whether this attribute is a single value or an array of values.
     * 
     */
    @Import(name="multivalued")
    private @Nullable Output<Boolean> multivalued;

    /**
     * @return Indicates whether this attribute is a single value or an array of values.
     * 
     */
    public Optional<Output<Boolean>> multivalued() {
        return Optional.ofNullable(this.multivalued);
    }

    /**
     * A human-friendly name that will appear in the Keycloak console.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A human-friendly name that will appear in the Keycloak console.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The realm id where the associated client or client scope exists.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm id where the associated client or client scope exists.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    @Import(name="userAttribute")
    private @Nullable Output<String> userAttribute;

    public Optional<Output<String>> userAttribute() {
        return Optional.ofNullable(this.userAttribute);
    }

    private UserAttributeProtocolMapperState() {}

    private UserAttributeProtocolMapperState(UserAttributeProtocolMapperState $) {
        this.addToAccessToken = $.addToAccessToken;
        this.addToIdToken = $.addToIdToken;
        this.addToUserinfo = $.addToUserinfo;
        this.aggregateAttributes = $.aggregateAttributes;
        this.claimName = $.claimName;
        this.claimValueType = $.claimValueType;
        this.clientId = $.clientId;
        this.clientScopeId = $.clientScopeId;
        this.multivalued = $.multivalued;
        this.name = $.name;
        this.realmId = $.realmId;
        this.userAttribute = $.userAttribute;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserAttributeProtocolMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserAttributeProtocolMapperState $;

        public Builder() {
            $ = new UserAttributeProtocolMapperState();
        }

        public Builder(UserAttributeProtocolMapperState defaults) {
            $ = new UserAttributeProtocolMapperState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addToAccessToken Indicates if the attribute should be a claim in the access token.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(@Nullable Output<Boolean> addToAccessToken) {
            $.addToAccessToken = addToAccessToken;
            return this;
        }

        /**
         * @param addToAccessToken Indicates if the attribute should be a claim in the access token.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(Boolean addToAccessToken) {
            return addToAccessToken(Output.of(addToAccessToken));
        }

        /**
         * @param addToIdToken Indicates if the attribute should be a claim in the id token.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(@Nullable Output<Boolean> addToIdToken) {
            $.addToIdToken = addToIdToken;
            return this;
        }

        /**
         * @param addToIdToken Indicates if the attribute should be a claim in the id token.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(Boolean addToIdToken) {
            return addToIdToken(Output.of(addToIdToken));
        }

        /**
         * @param addToUserinfo Indicates if the attribute should appear in the userinfo response body.
         * 
         * @return builder
         * 
         */
        public Builder addToUserinfo(@Nullable Output<Boolean> addToUserinfo) {
            $.addToUserinfo = addToUserinfo;
            return this;
        }

        /**
         * @param addToUserinfo Indicates if the attribute should appear in the userinfo response body.
         * 
         * @return builder
         * 
         */
        public Builder addToUserinfo(Boolean addToUserinfo) {
            return addToUserinfo(Output.of(addToUserinfo));
        }

        /**
         * @param aggregateAttributes Indicates if attribute values should be aggregated within the group attributes
         * 
         * @return builder
         * 
         */
        public Builder aggregateAttributes(@Nullable Output<Boolean> aggregateAttributes) {
            $.aggregateAttributes = aggregateAttributes;
            return this;
        }

        /**
         * @param aggregateAttributes Indicates if attribute values should be aggregated within the group attributes
         * 
         * @return builder
         * 
         */
        public Builder aggregateAttributes(Boolean aggregateAttributes) {
            return aggregateAttributes(Output.of(aggregateAttributes));
        }

        public Builder claimName(@Nullable Output<String> claimName) {
            $.claimName = claimName;
            return this;
        }

        public Builder claimName(String claimName) {
            return claimName(Output.of(claimName));
        }

        /**
         * @param claimValueType Claim type used when serializing tokens.
         * 
         * @return builder
         * 
         */
        public Builder claimValueType(@Nullable Output<String> claimValueType) {
            $.claimValueType = claimValueType;
            return this;
        }

        /**
         * @param claimValueType Claim type used when serializing tokens.
         * 
         * @return builder
         * 
         */
        public Builder claimValueType(String claimValueType) {
            return claimValueType(Output.of(claimValueType));
        }

        /**
         * @param clientId The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param clientScopeId The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(@Nullable Output<String> clientScopeId) {
            $.clientScopeId = clientScopeId;
            return this;
        }

        /**
         * @param clientScopeId The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(String clientScopeId) {
            return clientScopeId(Output.of(clientScopeId));
        }

        /**
         * @param multivalued Indicates whether this attribute is a single value or an array of values.
         * 
         * @return builder
         * 
         */
        public Builder multivalued(@Nullable Output<Boolean> multivalued) {
            $.multivalued = multivalued;
            return this;
        }

        /**
         * @param multivalued Indicates whether this attribute is a single value or an array of values.
         * 
         * @return builder
         * 
         */
        public Builder multivalued(Boolean multivalued) {
            return multivalued(Output.of(multivalued));
        }

        /**
         * @param name A human-friendly name that will appear in the Keycloak console.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A human-friendly name that will appear in the Keycloak console.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm id where the associated client or client scope exists.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm id where the associated client or client scope exists.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder userAttribute(@Nullable Output<String> userAttribute) {
            $.userAttribute = userAttribute;
            return this;
        }

        public Builder userAttribute(String userAttribute) {
            return userAttribute(Output.of(userAttribute));
        }

        public UserAttributeProtocolMapperState build() {
            return $;
        }
    }

}
