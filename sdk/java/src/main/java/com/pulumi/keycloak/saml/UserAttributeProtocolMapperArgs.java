// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserAttributeProtocolMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserAttributeProtocolMapperArgs Empty = new UserAttributeProtocolMapperArgs();

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
     * The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    @Import(name="clientScopeId")
    private @Nullable Output<String> clientScopeId;

    /**
     * @return The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    public Optional<Output<String>> clientScopeId() {
        return Optional.ofNullable(this.clientScopeId);
    }

    /**
     * An optional human-friendly name for this attribute.
     * 
     */
    @Import(name="friendlyName")
    private @Nullable Output<String> friendlyName;

    /**
     * @return An optional human-friendly name for this attribute.
     * 
     */
    public Optional<Output<String>> friendlyName() {
        return Optional.ofNullable(this.friendlyName);
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
     * The name of the SAML attribute.
     * 
     */
    @Import(name="samlAttributeName", required=true)
    private Output<String> samlAttributeName;

    /**
     * @return The name of the SAML attribute.
     * 
     */
    public Output<String> samlAttributeName() {
        return this.samlAttributeName;
    }

    /**
     * The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
     * 
     */
    @Import(name="samlAttributeNameFormat", required=true)
    private Output<String> samlAttributeNameFormat;

    /**
     * @return The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
     * 
     */
    public Output<String> samlAttributeNameFormat() {
        return this.samlAttributeNameFormat;
    }

    /**
     * The custom user attribute to map.
     * 
     */
    @Import(name="userAttribute", required=true)
    private Output<String> userAttribute;

    /**
     * @return The custom user attribute to map.
     * 
     */
    public Output<String> userAttribute() {
        return this.userAttribute;
    }

    private UserAttributeProtocolMapperArgs() {}

    private UserAttributeProtocolMapperArgs(UserAttributeProtocolMapperArgs $) {
        this.clientId = $.clientId;
        this.clientScopeId = $.clientScopeId;
        this.friendlyName = $.friendlyName;
        this.name = $.name;
        this.realmId = $.realmId;
        this.samlAttributeName = $.samlAttributeName;
        this.samlAttributeNameFormat = $.samlAttributeNameFormat;
        this.userAttribute = $.userAttribute;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserAttributeProtocolMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserAttributeProtocolMapperArgs $;

        public Builder() {
            $ = new UserAttributeProtocolMapperArgs();
        }

        public Builder(UserAttributeProtocolMapperArgs defaults) {
            $ = new UserAttributeProtocolMapperArgs(Objects.requireNonNull(defaults));
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
         * @param clientScopeId The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(@Nullable Output<String> clientScopeId) {
            $.clientScopeId = clientScopeId;
            return this;
        }

        /**
         * @param clientScopeId The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(String clientScopeId) {
            return clientScopeId(Output.of(clientScopeId));
        }

        /**
         * @param friendlyName An optional human-friendly name for this attribute.
         * 
         * @return builder
         * 
         */
        public Builder friendlyName(@Nullable Output<String> friendlyName) {
            $.friendlyName = friendlyName;
            return this;
        }

        /**
         * @param friendlyName An optional human-friendly name for this attribute.
         * 
         * @return builder
         * 
         */
        public Builder friendlyName(String friendlyName) {
            return friendlyName(Output.of(friendlyName));
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
         * @param samlAttributeName The name of the SAML attribute.
         * 
         * @return builder
         * 
         */
        public Builder samlAttributeName(Output<String> samlAttributeName) {
            $.samlAttributeName = samlAttributeName;
            return this;
        }

        /**
         * @param samlAttributeName The name of the SAML attribute.
         * 
         * @return builder
         * 
         */
        public Builder samlAttributeName(String samlAttributeName) {
            return samlAttributeName(Output.of(samlAttributeName));
        }

        /**
         * @param samlAttributeNameFormat The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
         * 
         * @return builder
         * 
         */
        public Builder samlAttributeNameFormat(Output<String> samlAttributeNameFormat) {
            $.samlAttributeNameFormat = samlAttributeNameFormat;
            return this;
        }

        /**
         * @param samlAttributeNameFormat The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
         * 
         * @return builder
         * 
         */
        public Builder samlAttributeNameFormat(String samlAttributeNameFormat) {
            return samlAttributeNameFormat(Output.of(samlAttributeNameFormat));
        }

        /**
         * @param userAttribute The custom user attribute to map.
         * 
         * @return builder
         * 
         */
        public Builder userAttribute(Output<String> userAttribute) {
            $.userAttribute = userAttribute;
            return this;
        }

        /**
         * @param userAttribute The custom user attribute to map.
         * 
         * @return builder
         * 
         */
        public Builder userAttribute(String userAttribute) {
            return userAttribute(Output.of(userAttribute));
        }

        public UserAttributeProtocolMapperArgs build() {
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("UserAttributeProtocolMapperArgs", "realmId");
            }
            if ($.samlAttributeName == null) {
                throw new MissingRequiredPropertyException("UserAttributeProtocolMapperArgs", "samlAttributeName");
            }
            if ($.samlAttributeNameFormat == null) {
                throw new MissingRequiredPropertyException("UserAttributeProtocolMapperArgs", "samlAttributeNameFormat");
            }
            if ($.userAttribute == null) {
                throw new MissingRequiredPropertyException("UserAttributeProtocolMapperArgs", "userAttribute");
            }
            return $;
        }
    }

}
