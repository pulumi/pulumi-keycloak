// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserAttributeMapperState extends com.pulumi.resources.ResourceArgs {

    public static final UserAttributeMapperState Empty = new UserAttributeMapperState();

    /**
     * When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
     * 
     */
    @Import(name="alwaysReadValueFromLdap")
    private @Nullable Output<Boolean> alwaysReadValueFromLdap;

    /**
     * @return When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> alwaysReadValueFromLdap() {
        return Optional.ofNullable(this.alwaysReadValueFromLdap);
    }

    /**
     * Default value to set in LDAP if `is_mandatory_in_ldap` is true and the value is empty.
     * 
     */
    @Import(name="attributeDefaultValue")
    private @Nullable Output<String> attributeDefaultValue;

    /**
     * @return Default value to set in LDAP if `is_mandatory_in_ldap` is true and the value is empty.
     * 
     */
    public Optional<Output<String>> attributeDefaultValue() {
        return Optional.ofNullable(this.attributeDefaultValue);
    }

    /**
     * Should be true for binary LDAP attributes.
     * 
     */
    @Import(name="isBinaryAttribute")
    private @Nullable Output<Boolean> isBinaryAttribute;

    /**
     * @return Should be true for binary LDAP attributes.
     * 
     */
    public Optional<Output<Boolean>> isBinaryAttribute() {
        return Optional.ofNullable(this.isBinaryAttribute);
    }

    /**
     * When `true`, this attribute must exist in LDAP. Defaults to `false`.
     * 
     */
    @Import(name="isMandatoryInLdap")
    private @Nullable Output<Boolean> isMandatoryInLdap;

    /**
     * @return When `true`, this attribute must exist in LDAP. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> isMandatoryInLdap() {
        return Optional.ofNullable(this.isMandatoryInLdap);
    }

    /**
     * Name of the mapped attribute on the LDAP object.
     * 
     */
    @Import(name="ldapAttribute")
    private @Nullable Output<String> ldapAttribute;

    /**
     * @return Name of the mapped attribute on the LDAP object.
     * 
     */
    public Optional<Output<String>> ldapAttribute() {
        return Optional.ofNullable(this.ldapAttribute);
    }

    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    @Import(name="ldapUserFederationId")
    private @Nullable Output<String> ldapUserFederationId;

    /**
     * @return The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    public Optional<Output<String>> ldapUserFederationId() {
        return Optional.ofNullable(this.ldapUserFederationId);
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
     * When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
     * 
     */
    @Import(name="readOnly")
    private @Nullable Output<Boolean> readOnly;

    /**
     * @return When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> readOnly() {
        return Optional.ofNullable(this.readOnly);
    }

    /**
     * The realm that this LDAP mapper will exist in.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm that this LDAP mapper will exist in.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    /**
     * Name of the user property or attribute you want to map the LDAP attribute into.
     * 
     */
    @Import(name="userModelAttribute")
    private @Nullable Output<String> userModelAttribute;

    /**
     * @return Name of the user property or attribute you want to map the LDAP attribute into.
     * 
     */
    public Optional<Output<String>> userModelAttribute() {
        return Optional.ofNullable(this.userModelAttribute);
    }

    private UserAttributeMapperState() {}

    private UserAttributeMapperState(UserAttributeMapperState $) {
        this.alwaysReadValueFromLdap = $.alwaysReadValueFromLdap;
        this.attributeDefaultValue = $.attributeDefaultValue;
        this.isBinaryAttribute = $.isBinaryAttribute;
        this.isMandatoryInLdap = $.isMandatoryInLdap;
        this.ldapAttribute = $.ldapAttribute;
        this.ldapUserFederationId = $.ldapUserFederationId;
        this.name = $.name;
        this.readOnly = $.readOnly;
        this.realmId = $.realmId;
        this.userModelAttribute = $.userModelAttribute;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserAttributeMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserAttributeMapperState $;

        public Builder() {
            $ = new UserAttributeMapperState();
        }

        public Builder(UserAttributeMapperState defaults) {
            $ = new UserAttributeMapperState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alwaysReadValueFromLdap When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder alwaysReadValueFromLdap(@Nullable Output<Boolean> alwaysReadValueFromLdap) {
            $.alwaysReadValueFromLdap = alwaysReadValueFromLdap;
            return this;
        }

        /**
         * @param alwaysReadValueFromLdap When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder alwaysReadValueFromLdap(Boolean alwaysReadValueFromLdap) {
            return alwaysReadValueFromLdap(Output.of(alwaysReadValueFromLdap));
        }

        /**
         * @param attributeDefaultValue Default value to set in LDAP if `is_mandatory_in_ldap` is true and the value is empty.
         * 
         * @return builder
         * 
         */
        public Builder attributeDefaultValue(@Nullable Output<String> attributeDefaultValue) {
            $.attributeDefaultValue = attributeDefaultValue;
            return this;
        }

        /**
         * @param attributeDefaultValue Default value to set in LDAP if `is_mandatory_in_ldap` is true and the value is empty.
         * 
         * @return builder
         * 
         */
        public Builder attributeDefaultValue(String attributeDefaultValue) {
            return attributeDefaultValue(Output.of(attributeDefaultValue));
        }

        /**
         * @param isBinaryAttribute Should be true for binary LDAP attributes.
         * 
         * @return builder
         * 
         */
        public Builder isBinaryAttribute(@Nullable Output<Boolean> isBinaryAttribute) {
            $.isBinaryAttribute = isBinaryAttribute;
            return this;
        }

        /**
         * @param isBinaryAttribute Should be true for binary LDAP attributes.
         * 
         * @return builder
         * 
         */
        public Builder isBinaryAttribute(Boolean isBinaryAttribute) {
            return isBinaryAttribute(Output.of(isBinaryAttribute));
        }

        /**
         * @param isMandatoryInLdap When `true`, this attribute must exist in LDAP. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder isMandatoryInLdap(@Nullable Output<Boolean> isMandatoryInLdap) {
            $.isMandatoryInLdap = isMandatoryInLdap;
            return this;
        }

        /**
         * @param isMandatoryInLdap When `true`, this attribute must exist in LDAP. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder isMandatoryInLdap(Boolean isMandatoryInLdap) {
            return isMandatoryInLdap(Output.of(isMandatoryInLdap));
        }

        /**
         * @param ldapAttribute Name of the mapped attribute on the LDAP object.
         * 
         * @return builder
         * 
         */
        public Builder ldapAttribute(@Nullable Output<String> ldapAttribute) {
            $.ldapAttribute = ldapAttribute;
            return this;
        }

        /**
         * @param ldapAttribute Name of the mapped attribute on the LDAP object.
         * 
         * @return builder
         * 
         */
        public Builder ldapAttribute(String ldapAttribute) {
            return ldapAttribute(Output.of(ldapAttribute));
        }

        /**
         * @param ldapUserFederationId The ID of the LDAP user federation provider to attach this mapper to.
         * 
         * @return builder
         * 
         */
        public Builder ldapUserFederationId(@Nullable Output<String> ldapUserFederationId) {
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
         * @param readOnly When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder readOnly(@Nullable Output<Boolean> readOnly) {
            $.readOnly = readOnly;
            return this;
        }

        /**
         * @param readOnly When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder readOnly(Boolean readOnly) {
            return readOnly(Output.of(readOnly));
        }

        /**
         * @param realmId The realm that this LDAP mapper will exist in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
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

        /**
         * @param userModelAttribute Name of the user property or attribute you want to map the LDAP attribute into.
         * 
         * @return builder
         * 
         */
        public Builder userModelAttribute(@Nullable Output<String> userModelAttribute) {
            $.userModelAttribute = userModelAttribute;
            return this;
        }

        /**
         * @param userModelAttribute Name of the user property or attribute you want to map the LDAP attribute into.
         * 
         * @return builder
         * 
         */
        public Builder userModelAttribute(String userModelAttribute) {
            return userModelAttribute(Output.of(userModelAttribute));
        }

        public UserAttributeMapperState build() {
            return $;
        }
    }

}
