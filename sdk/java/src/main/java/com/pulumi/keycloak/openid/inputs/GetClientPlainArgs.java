// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetClientPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetClientPlainArgs Empty = new GetClientPlainArgs();

    /**
     * The client id (not its unique ID).
     * 
     */
    @Import(name="clientId", required=true)
    private String clientId;

    /**
     * @return The client id (not its unique ID).
     * 
     */
    public String clientId() {
        return this.clientId;
    }

    @Import(name="consentScreenText")
    private @Nullable String consentScreenText;

    public Optional<String> consentScreenText() {
        return Optional.ofNullable(this.consentScreenText);
    }

    @Import(name="displayOnConsentScreen")
    private @Nullable Boolean displayOnConsentScreen;

    public Optional<Boolean> displayOnConsentScreen() {
        return Optional.ofNullable(this.displayOnConsentScreen);
    }

    @Import(name="extraConfig")
    private @Nullable Map<String,Object> extraConfig;

    public Optional<Map<String,Object>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    @Import(name="oauth2DeviceAuthorizationGrantEnabled")
    private @Nullable Boolean oauth2DeviceAuthorizationGrantEnabled;

    public Optional<Boolean> oauth2DeviceAuthorizationGrantEnabled() {
        return Optional.ofNullable(this.oauth2DeviceAuthorizationGrantEnabled);
    }

    @Import(name="oauth2DeviceCodeLifespan")
    private @Nullable String oauth2DeviceCodeLifespan;

    public Optional<String> oauth2DeviceCodeLifespan() {
        return Optional.ofNullable(this.oauth2DeviceCodeLifespan);
    }

    @Import(name="oauth2DevicePollingInterval")
    private @Nullable String oauth2DevicePollingInterval;

    public Optional<String> oauth2DevicePollingInterval() {
        return Optional.ofNullable(this.oauth2DevicePollingInterval);
    }

    /**
     * The realm id.
     * 
     */
    @Import(name="realmId", required=true)
    private String realmId;

    /**
     * @return The realm id.
     * 
     */
    public String realmId() {
        return this.realmId;
    }

    private GetClientPlainArgs() {}

    private GetClientPlainArgs(GetClientPlainArgs $) {
        this.clientId = $.clientId;
        this.consentScreenText = $.consentScreenText;
        this.displayOnConsentScreen = $.displayOnConsentScreen;
        this.extraConfig = $.extraConfig;
        this.oauth2DeviceAuthorizationGrantEnabled = $.oauth2DeviceAuthorizationGrantEnabled;
        this.oauth2DeviceCodeLifespan = $.oauth2DeviceCodeLifespan;
        this.oauth2DevicePollingInterval = $.oauth2DevicePollingInterval;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetClientPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetClientPlainArgs $;

        public Builder() {
            $ = new GetClientPlainArgs();
        }

        public Builder(GetClientPlainArgs defaults) {
            $ = new GetClientPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId The client id (not its unique ID).
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            $.clientId = clientId;
            return this;
        }

        public Builder consentScreenText(@Nullable String consentScreenText) {
            $.consentScreenText = consentScreenText;
            return this;
        }

        public Builder displayOnConsentScreen(@Nullable Boolean displayOnConsentScreen) {
            $.displayOnConsentScreen = displayOnConsentScreen;
            return this;
        }

        public Builder extraConfig(@Nullable Map<String,Object> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder oauth2DeviceAuthorizationGrantEnabled(@Nullable Boolean oauth2DeviceAuthorizationGrantEnabled) {
            $.oauth2DeviceAuthorizationGrantEnabled = oauth2DeviceAuthorizationGrantEnabled;
            return this;
        }

        public Builder oauth2DeviceCodeLifespan(@Nullable String oauth2DeviceCodeLifespan) {
            $.oauth2DeviceCodeLifespan = oauth2DeviceCodeLifespan;
            return this;
        }

        public Builder oauth2DevicePollingInterval(@Nullable String oauth2DevicePollingInterval) {
            $.oauth2DevicePollingInterval = oauth2DevicePollingInterval;
            return this;
        }

        /**
         * @param realmId The realm id.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            $.realmId = realmId;
            return this;
        }

        public GetClientPlainArgs build() {
            $.clientId = Objects.requireNonNull($.clientId, "expected parameter 'clientId' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}
