// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.inputs.GetRealmInternationalizationArgs;
import com.pulumi.keycloak.inputs.GetRealmOtpPolicyArgs;
import com.pulumi.keycloak.inputs.GetRealmSecurityDefenseArgs;
import com.pulumi.keycloak.inputs.GetRealmSmtpServerArgs;
import com.pulumi.keycloak.inputs.GetRealmWebAuthnPasswordlessPolicyArgs;
import com.pulumi.keycloak.inputs.GetRealmWebAuthnPolicyArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRealmArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRealmArgs Empty = new GetRealmArgs();

    @Import(name="attributes")
    private @Nullable Output<Map<String,String>> attributes;

    public Optional<Output<Map<String,String>>> attributes() {
        return Optional.ofNullable(this.attributes);
    }

    @Import(name="defaultDefaultClientScopes")
    private @Nullable Output<List<String>> defaultDefaultClientScopes;

    public Optional<Output<List<String>>> defaultDefaultClientScopes() {
        return Optional.ofNullable(this.defaultDefaultClientScopes);
    }

    @Import(name="defaultOptionalClientScopes")
    private @Nullable Output<List<String>> defaultOptionalClientScopes;

    public Optional<Output<List<String>>> defaultOptionalClientScopes() {
        return Optional.ofNullable(this.defaultOptionalClientScopes);
    }

    @Import(name="displayNameHtml")
    private @Nullable Output<String> displayNameHtml;

    public Optional<Output<String>> displayNameHtml() {
        return Optional.ofNullable(this.displayNameHtml);
    }

    @Import(name="internationalizations")
    private @Nullable Output<List<GetRealmInternationalizationArgs>> internationalizations;

    public Optional<Output<List<GetRealmInternationalizationArgs>>> internationalizations() {
        return Optional.ofNullable(this.internationalizations);
    }

    @Import(name="otpPolicy")
    private @Nullable Output<GetRealmOtpPolicyArgs> otpPolicy;

    public Optional<Output<GetRealmOtpPolicyArgs>> otpPolicy() {
        return Optional.ofNullable(this.otpPolicy);
    }

    /**
     * The realm name.
     * 
     */
    @Import(name="realm", required=true)
    private Output<String> realm;

    /**
     * @return The realm name.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }

    @Import(name="securityDefenses")
    private @Nullable Output<List<GetRealmSecurityDefenseArgs>> securityDefenses;

    public Optional<Output<List<GetRealmSecurityDefenseArgs>>> securityDefenses() {
        return Optional.ofNullable(this.securityDefenses);
    }

    @Import(name="smtpServers")
    private @Nullable Output<List<GetRealmSmtpServerArgs>> smtpServers;

    public Optional<Output<List<GetRealmSmtpServerArgs>>> smtpServers() {
        return Optional.ofNullable(this.smtpServers);
    }

    @Import(name="webAuthnPasswordlessPolicy")
    private @Nullable Output<GetRealmWebAuthnPasswordlessPolicyArgs> webAuthnPasswordlessPolicy;

    public Optional<Output<GetRealmWebAuthnPasswordlessPolicyArgs>> webAuthnPasswordlessPolicy() {
        return Optional.ofNullable(this.webAuthnPasswordlessPolicy);
    }

    @Import(name="webAuthnPolicy")
    private @Nullable Output<GetRealmWebAuthnPolicyArgs> webAuthnPolicy;

    public Optional<Output<GetRealmWebAuthnPolicyArgs>> webAuthnPolicy() {
        return Optional.ofNullable(this.webAuthnPolicy);
    }

    private GetRealmArgs() {}

    private GetRealmArgs(GetRealmArgs $) {
        this.attributes = $.attributes;
        this.defaultDefaultClientScopes = $.defaultDefaultClientScopes;
        this.defaultOptionalClientScopes = $.defaultOptionalClientScopes;
        this.displayNameHtml = $.displayNameHtml;
        this.internationalizations = $.internationalizations;
        this.otpPolicy = $.otpPolicy;
        this.realm = $.realm;
        this.securityDefenses = $.securityDefenses;
        this.smtpServers = $.smtpServers;
        this.webAuthnPasswordlessPolicy = $.webAuthnPasswordlessPolicy;
        this.webAuthnPolicy = $.webAuthnPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRealmArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRealmArgs $;

        public Builder() {
            $ = new GetRealmArgs();
        }

        public Builder(GetRealmArgs defaults) {
            $ = new GetRealmArgs(Objects.requireNonNull(defaults));
        }

        public Builder attributes(@Nullable Output<Map<String,String>> attributes) {
            $.attributes = attributes;
            return this;
        }

        public Builder attributes(Map<String,String> attributes) {
            return attributes(Output.of(attributes));
        }

        public Builder defaultDefaultClientScopes(@Nullable Output<List<String>> defaultDefaultClientScopes) {
            $.defaultDefaultClientScopes = defaultDefaultClientScopes;
            return this;
        }

        public Builder defaultDefaultClientScopes(List<String> defaultDefaultClientScopes) {
            return defaultDefaultClientScopes(Output.of(defaultDefaultClientScopes));
        }

        public Builder defaultDefaultClientScopes(String... defaultDefaultClientScopes) {
            return defaultDefaultClientScopes(List.of(defaultDefaultClientScopes));
        }

        public Builder defaultOptionalClientScopes(@Nullable Output<List<String>> defaultOptionalClientScopes) {
            $.defaultOptionalClientScopes = defaultOptionalClientScopes;
            return this;
        }

        public Builder defaultOptionalClientScopes(List<String> defaultOptionalClientScopes) {
            return defaultOptionalClientScopes(Output.of(defaultOptionalClientScopes));
        }

        public Builder defaultOptionalClientScopes(String... defaultOptionalClientScopes) {
            return defaultOptionalClientScopes(List.of(defaultOptionalClientScopes));
        }

        public Builder displayNameHtml(@Nullable Output<String> displayNameHtml) {
            $.displayNameHtml = displayNameHtml;
            return this;
        }

        public Builder displayNameHtml(String displayNameHtml) {
            return displayNameHtml(Output.of(displayNameHtml));
        }

        public Builder internationalizations(@Nullable Output<List<GetRealmInternationalizationArgs>> internationalizations) {
            $.internationalizations = internationalizations;
            return this;
        }

        public Builder internationalizations(List<GetRealmInternationalizationArgs> internationalizations) {
            return internationalizations(Output.of(internationalizations));
        }

        public Builder internationalizations(GetRealmInternationalizationArgs... internationalizations) {
            return internationalizations(List.of(internationalizations));
        }

        public Builder otpPolicy(@Nullable Output<GetRealmOtpPolicyArgs> otpPolicy) {
            $.otpPolicy = otpPolicy;
            return this;
        }

        public Builder otpPolicy(GetRealmOtpPolicyArgs otpPolicy) {
            return otpPolicy(Output.of(otpPolicy));
        }

        /**
         * @param realm The realm name.
         * 
         * @return builder
         * 
         */
        public Builder realm(Output<String> realm) {
            $.realm = realm;
            return this;
        }

        /**
         * @param realm The realm name.
         * 
         * @return builder
         * 
         */
        public Builder realm(String realm) {
            return realm(Output.of(realm));
        }

        public Builder securityDefenses(@Nullable Output<List<GetRealmSecurityDefenseArgs>> securityDefenses) {
            $.securityDefenses = securityDefenses;
            return this;
        }

        public Builder securityDefenses(List<GetRealmSecurityDefenseArgs> securityDefenses) {
            return securityDefenses(Output.of(securityDefenses));
        }

        public Builder securityDefenses(GetRealmSecurityDefenseArgs... securityDefenses) {
            return securityDefenses(List.of(securityDefenses));
        }

        public Builder smtpServers(@Nullable Output<List<GetRealmSmtpServerArgs>> smtpServers) {
            $.smtpServers = smtpServers;
            return this;
        }

        public Builder smtpServers(List<GetRealmSmtpServerArgs> smtpServers) {
            return smtpServers(Output.of(smtpServers));
        }

        public Builder smtpServers(GetRealmSmtpServerArgs... smtpServers) {
            return smtpServers(List.of(smtpServers));
        }

        public Builder webAuthnPasswordlessPolicy(@Nullable Output<GetRealmWebAuthnPasswordlessPolicyArgs> webAuthnPasswordlessPolicy) {
            $.webAuthnPasswordlessPolicy = webAuthnPasswordlessPolicy;
            return this;
        }

        public Builder webAuthnPasswordlessPolicy(GetRealmWebAuthnPasswordlessPolicyArgs webAuthnPasswordlessPolicy) {
            return webAuthnPasswordlessPolicy(Output.of(webAuthnPasswordlessPolicy));
        }

        public Builder webAuthnPolicy(@Nullable Output<GetRealmWebAuthnPolicyArgs> webAuthnPolicy) {
            $.webAuthnPolicy = webAuthnPolicy;
            return this;
        }

        public Builder webAuthnPolicy(GetRealmWebAuthnPolicyArgs webAuthnPolicy) {
            return webAuthnPolicy(Output.of(webAuthnPolicy));
        }

        public GetRealmArgs build() {
            if ($.realm == null) {
                throw new MissingRequiredPropertyException("GetRealmArgs", "realm");
            }
            return $;
        }
    }

}
