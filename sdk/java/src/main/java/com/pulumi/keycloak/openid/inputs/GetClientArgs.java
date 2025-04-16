// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetClientArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetClientArgs Empty = new GetClientArgs();

    @Import(name="alwaysDisplayInConsole")
    private @Nullable Output<Boolean> alwaysDisplayInConsole;

    public Optional<Output<Boolean>> alwaysDisplayInConsole() {
        return Optional.ofNullable(this.alwaysDisplayInConsole);
    }

    /**
     * The client id (not its unique ID).
     * 
     */
    @Import(name="clientId", required=true)
    private Output<String> clientId;

    /**
     * @return The client id (not its unique ID).
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }

    @Import(name="consentScreenText")
    private @Nullable Output<String> consentScreenText;

    public Optional<Output<String>> consentScreenText() {
        return Optional.ofNullable(this.consentScreenText);
    }

    @Import(name="displayOnConsentScreen")
    private @Nullable Output<Boolean> displayOnConsentScreen;

    public Optional<Output<Boolean>> displayOnConsentScreen() {
        return Optional.ofNullable(this.displayOnConsentScreen);
    }

    @Import(name="extraConfig")
    private @Nullable Output<Map<String,String>> extraConfig;

    public Optional<Output<Map<String,String>>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    @Import(name="oauth2DeviceAuthorizationGrantEnabled")
    private @Nullable Output<Boolean> oauth2DeviceAuthorizationGrantEnabled;

    public Optional<Output<Boolean>> oauth2DeviceAuthorizationGrantEnabled() {
        return Optional.ofNullable(this.oauth2DeviceAuthorizationGrantEnabled);
    }

    @Import(name="oauth2DeviceCodeLifespan")
    private @Nullable Output<String> oauth2DeviceCodeLifespan;

    public Optional<Output<String>> oauth2DeviceCodeLifespan() {
        return Optional.ofNullable(this.oauth2DeviceCodeLifespan);
    }

    @Import(name="oauth2DevicePollingInterval")
    private @Nullable Output<String> oauth2DevicePollingInterval;

    public Optional<Output<String>> oauth2DevicePollingInterval() {
        return Optional.ofNullable(this.oauth2DevicePollingInterval);
    }

    /**
     * The realm id.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm id.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private GetClientArgs() {}

    private GetClientArgs(GetClientArgs $) {
        this.alwaysDisplayInConsole = $.alwaysDisplayInConsole;
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
    public static Builder builder(GetClientArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetClientArgs $;

        public Builder() {
            $ = new GetClientArgs();
        }

        public Builder(GetClientArgs defaults) {
            $ = new GetClientArgs(Objects.requireNonNull(defaults));
        }

        public Builder alwaysDisplayInConsole(@Nullable Output<Boolean> alwaysDisplayInConsole) {
            $.alwaysDisplayInConsole = alwaysDisplayInConsole;
            return this;
        }

        public Builder alwaysDisplayInConsole(Boolean alwaysDisplayInConsole) {
            return alwaysDisplayInConsole(Output.of(alwaysDisplayInConsole));
        }

        /**
         * @param clientId The client id (not its unique ID).
         * 
         * @return builder
         * 
         */
        public Builder clientId(Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The client id (not its unique ID).
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        public Builder consentScreenText(@Nullable Output<String> consentScreenText) {
            $.consentScreenText = consentScreenText;
            return this;
        }

        public Builder consentScreenText(String consentScreenText) {
            return consentScreenText(Output.of(consentScreenText));
        }

        public Builder displayOnConsentScreen(@Nullable Output<Boolean> displayOnConsentScreen) {
            $.displayOnConsentScreen = displayOnConsentScreen;
            return this;
        }

        public Builder displayOnConsentScreen(Boolean displayOnConsentScreen) {
            return displayOnConsentScreen(Output.of(displayOnConsentScreen));
        }

        public Builder extraConfig(@Nullable Output<Map<String,String>> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder extraConfig(Map<String,String> extraConfig) {
            return extraConfig(Output.of(extraConfig));
        }

        public Builder oauth2DeviceAuthorizationGrantEnabled(@Nullable Output<Boolean> oauth2DeviceAuthorizationGrantEnabled) {
            $.oauth2DeviceAuthorizationGrantEnabled = oauth2DeviceAuthorizationGrantEnabled;
            return this;
        }

        public Builder oauth2DeviceAuthorizationGrantEnabled(Boolean oauth2DeviceAuthorizationGrantEnabled) {
            return oauth2DeviceAuthorizationGrantEnabled(Output.of(oauth2DeviceAuthorizationGrantEnabled));
        }

        public Builder oauth2DeviceCodeLifespan(@Nullable Output<String> oauth2DeviceCodeLifespan) {
            $.oauth2DeviceCodeLifespan = oauth2DeviceCodeLifespan;
            return this;
        }

        public Builder oauth2DeviceCodeLifespan(String oauth2DeviceCodeLifespan) {
            return oauth2DeviceCodeLifespan(Output.of(oauth2DeviceCodeLifespan));
        }

        public Builder oauth2DevicePollingInterval(@Nullable Output<String> oauth2DevicePollingInterval) {
            $.oauth2DevicePollingInterval = oauth2DevicePollingInterval;
            return this;
        }

        public Builder oauth2DevicePollingInterval(String oauth2DevicePollingInterval) {
            return oauth2DevicePollingInterval(Output.of(oauth2DevicePollingInterval));
        }

        /**
         * @param realmId The realm id.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm id.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public GetClientArgs build() {
            if ($.clientId == null) {
                throw new MissingRequiredPropertyException("GetClientArgs", "clientId");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("GetClientArgs", "realmId");
            }
            return $;
        }
    }

}
