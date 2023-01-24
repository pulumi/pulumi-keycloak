// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HardcodedAttributeMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final HardcodedAttributeMapperArgs Empty = new HardcodedAttributeMapperArgs();

    /**
     * The name of the LDAP attribute to set.
     * 
     */
    @Import(name="attributeName", required=true)
    private Output<String> attributeName;

    /**
     * @return The name of the LDAP attribute to set.
     * 
     */
    public Output<String> attributeName() {
        return this.attributeName;
    }

    /**
     * The value to set to the LDAP attribute. You can hardcode any value like &#39;foo&#39;.
     * 
     */
    @Import(name="attributeValue", required=true)
    private Output<String> attributeValue;

    /**
     * @return The value to set to the LDAP attribute. You can hardcode any value like &#39;foo&#39;.
     * 
     */
    public Output<String> attributeValue() {
        return this.attributeValue;
    }

    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    @Import(name="ldapUserFederationId", required=true)
    private Output<String> ldapUserFederationId;

    /**
     * @return The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    public Output<String> ldapUserFederationId() {
        return this.ldapUserFederationId;
    }

    /**
     * Display name of this mapper when displayed in the console.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Display name of this mapper when displayed in the console.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The realm that this LDAP mapper will exist in.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm that this LDAP mapper will exist in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private HardcodedAttributeMapperArgs() {}

    private HardcodedAttributeMapperArgs(HardcodedAttributeMapperArgs $) {
        this.attributeName = $.attributeName;
        this.attributeValue = $.attributeValue;
        this.ldapUserFederationId = $.ldapUserFederationId;
        this.name = $.name;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HardcodedAttributeMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HardcodedAttributeMapperArgs $;

        public Builder() {
            $ = new HardcodedAttributeMapperArgs();
        }

        public Builder(HardcodedAttributeMapperArgs defaults) {
            $ = new HardcodedAttributeMapperArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attributeName The name of the LDAP attribute to set.
         * 
         * @return builder
         * 
         */
        public Builder attributeName(Output<String> attributeName) {
            $.attributeName = attributeName;
            return this;
        }

        /**
         * @param attributeName The name of the LDAP attribute to set.
         * 
         * @return builder
         * 
         */
        public Builder attributeName(String attributeName) {
            return attributeName(Output.of(attributeName));
        }

        /**
         * @param attributeValue The value to set to the LDAP attribute. You can hardcode any value like &#39;foo&#39;.
         * 
         * @return builder
         * 
         */
        public Builder attributeValue(Output<String> attributeValue) {
            $.attributeValue = attributeValue;
            return this;
        }

        /**
         * @param attributeValue The value to set to the LDAP attribute. You can hardcode any value like &#39;foo&#39;.
         * 
         * @return builder
         * 
         */
        public Builder attributeValue(String attributeValue) {
            return attributeValue(Output.of(attributeValue));
        }

        /**
         * @param ldapUserFederationId The ID of the LDAP user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(Output<String> ldapUserFederationId) {
            $.ldapUserFederationId = ldapUserFederationId;
            return this;
        }

        /**
         * @param ldapUserFederationId The ID of the LDAP user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(String ldapUserFederationId) {
            return ldapUserFederationId(Output.of(ldapUserFederationId));
        }

        /**
         * @param name Display name of this mapper when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Display name of this mapper when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm that this LDAP mapper will exist in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm that this LDAP mapper will exist in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public HardcodedAttributeMapperArgs build() {
            $.attributeName = Objects.requireNonNull($.attributeName, "expected parameter 'attributeName' to be non-null");
            $.attributeValue = Objects.requireNonNull($.attributeValue, "expected parameter 'attributeValue' to be non-null");
            $.ldapUserFederationId = Objects.requireNonNull($.ldapUserFederationId, "expected parameter 'ldapUserFederationId' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}
