// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    @Import(name="additionalHeaders", json=true)
    private @Nullable Output<Map<String,String>> additionalHeaders;

    public Optional<Output<Map<String,String>>> additionalHeaders() {
        return Optional.ofNullable(this.additionalHeaders);
    }

    @Import(name="basePath")
    private @Nullable Output<String> basePath;

    public Optional<Output<String>> basePath() {
        return Optional.ofNullable(this.basePath);
    }

    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    @Import(name="clientSecret")
    private @Nullable Output<String> clientSecret;

    public Optional<Output<String>> clientSecret() {
        return Optional.ofNullable(this.clientSecret);
    }

    /**
     * Timeout (in seconds) of the Keycloak client
     * 
     */
    @Import(name="clientTimeout", json=true)
    private @Nullable Output<Integer> clientTimeout;

    /**
     * @return Timeout (in seconds) of the Keycloak client
     * 
     */
    public Optional<Output<Integer>> clientTimeout() {
        return Optional.ofNullable(this.clientTimeout);
    }

    /**
     * Whether or not to login to Keycloak instance on provider initialization
     * 
     */
    @Import(name="initialLogin", json=true)
    private @Nullable Output<Boolean> initialLogin;

    /**
     * @return Whether or not to login to Keycloak instance on provider initialization
     * 
     */
    public Optional<Output<Boolean>> initialLogin() {
        return Optional.ofNullable(this.initialLogin);
    }

    @Import(name="password")
    private @Nullable Output<String> password;

    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="realm")
    private @Nullable Output<String> realm;

    public Optional<Output<String>> realm() {
        return Optional.ofNullable(this.realm);
    }

    /**
     * When true, the provider will treat the Keycloak instance as a Red Hat SSO server, specifically when parsing the version
     * returned from the /serverinfo API endpoint.
     * 
     */
    @Import(name="redHatSso", json=true)
    private @Nullable Output<Boolean> redHatSso;

    /**
     * @return When true, the provider will treat the Keycloak instance as a Red Hat SSO server, specifically when parsing the version
     * returned from the /serverinfo API endpoint.
     * 
     */
    public Optional<Output<Boolean>> redHatSso() {
        return Optional.ofNullable(this.redHatSso);
    }

    /**
     * Allows x509 calls using an unknown CA certificate (for development purposes)
     * 
     */
    @Import(name="rootCaCertificate")
    private @Nullable Output<String> rootCaCertificate;

    /**
     * @return Allows x509 calls using an unknown CA certificate (for development purposes)
     * 
     */
    public Optional<Output<String>> rootCaCertificate() {
        return Optional.ofNullable(this.rootCaCertificate);
    }

    /**
     * Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
     * should be avoided.
     * 
     */
    @Import(name="tlsInsecureSkipVerify", json=true)
    private @Nullable Output<Boolean> tlsInsecureSkipVerify;

    /**
     * @return Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
     * should be avoided.
     * 
     */
    public Optional<Output<Boolean>> tlsInsecureSkipVerify() {
        return Optional.ofNullable(this.tlsInsecureSkipVerify);
    }

    /**
     * The base URL of the Keycloak instance, before `/auth`
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The base URL of the Keycloak instance, before `/auth`
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    @Import(name="username")
    private @Nullable Output<String> username;

    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.additionalHeaders = $.additionalHeaders;
        this.basePath = $.basePath;
        this.clientId = $.clientId;
        this.clientSecret = $.clientSecret;
        this.clientTimeout = $.clientTimeout;
        this.initialLogin = $.initialLogin;
        this.password = $.password;
        this.realm = $.realm;
        this.redHatSso = $.redHatSso;
        this.rootCaCertificate = $.rootCaCertificate;
        this.tlsInsecureSkipVerify = $.tlsInsecureSkipVerify;
        this.url = $.url;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        public Builder additionalHeaders(@Nullable Output<Map<String,String>> additionalHeaders) {
            $.additionalHeaders = additionalHeaders;
            return this;
        }

        public Builder additionalHeaders(Map<String,String> additionalHeaders) {
            return additionalHeaders(Output.of(additionalHeaders));
        }

        public Builder basePath(@Nullable Output<String> basePath) {
            $.basePath = basePath;
            return this;
        }

        public Builder basePath(String basePath) {
            return basePath(Output.of(basePath));
        }

        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        public Builder clientSecret(@Nullable Output<String> clientSecret) {
            $.clientSecret = clientSecret;
            return this;
        }

        public Builder clientSecret(String clientSecret) {
            return clientSecret(Output.of(clientSecret));
        }

        /**
         * @param clientTimeout Timeout (in seconds) of the Keycloak client
         * 
         * @return builder
         * 
         */
        public Builder clientTimeout(@Nullable Output<Integer> clientTimeout) {
            $.clientTimeout = clientTimeout;
            return this;
        }

        /**
         * @param clientTimeout Timeout (in seconds) of the Keycloak client
         * 
         * @return builder
         * 
         */
        public Builder clientTimeout(Integer clientTimeout) {
            return clientTimeout(Output.of(clientTimeout));
        }

        /**
         * @param initialLogin Whether or not to login to Keycloak instance on provider initialization
         * 
         * @return builder
         * 
         */
        public Builder initialLogin(@Nullable Output<Boolean> initialLogin) {
            $.initialLogin = initialLogin;
            return this;
        }

        /**
         * @param initialLogin Whether or not to login to Keycloak instance on provider initialization
         * 
         * @return builder
         * 
         */
        public Builder initialLogin(Boolean initialLogin) {
            return initialLogin(Output.of(initialLogin));
        }

        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder realm(@Nullable Output<String> realm) {
            $.realm = realm;
            return this;
        }

        public Builder realm(String realm) {
            return realm(Output.of(realm));
        }

        /**
         * @param redHatSso When true, the provider will treat the Keycloak instance as a Red Hat SSO server, specifically when parsing the version
         * returned from the /serverinfo API endpoint.
         * 
         * @return builder
         * 
         */
        public Builder redHatSso(@Nullable Output<Boolean> redHatSso) {
            $.redHatSso = redHatSso;
            return this;
        }

        /**
         * @param redHatSso When true, the provider will treat the Keycloak instance as a Red Hat SSO server, specifically when parsing the version
         * returned from the /serverinfo API endpoint.
         * 
         * @return builder
         * 
         */
        public Builder redHatSso(Boolean redHatSso) {
            return redHatSso(Output.of(redHatSso));
        }

        /**
         * @param rootCaCertificate Allows x509 calls using an unknown CA certificate (for development purposes)
         * 
         * @return builder
         * 
         */
        public Builder rootCaCertificate(@Nullable Output<String> rootCaCertificate) {
            $.rootCaCertificate = rootCaCertificate;
            return this;
        }

        /**
         * @param rootCaCertificate Allows x509 calls using an unknown CA certificate (for development purposes)
         * 
         * @return builder
         * 
         */
        public Builder rootCaCertificate(String rootCaCertificate) {
            return rootCaCertificate(Output.of(rootCaCertificate));
        }

        /**
         * @param tlsInsecureSkipVerify Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
         * should be avoided.
         * 
         * @return builder
         * 
         */
        public Builder tlsInsecureSkipVerify(@Nullable Output<Boolean> tlsInsecureSkipVerify) {
            $.tlsInsecureSkipVerify = tlsInsecureSkipVerify;
            return this;
        }

        /**
         * @param tlsInsecureSkipVerify Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
         * should be avoided.
         * 
         * @return builder
         * 
         */
        public Builder tlsInsecureSkipVerify(Boolean tlsInsecureSkipVerify) {
            return tlsInsecureSkipVerify(Output.of(tlsInsecureSkipVerify));
        }

        /**
         * @param url The base URL of the Keycloak instance, before `/auth`
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The base URL of the Keycloak instance, before `/auth`
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProviderArgs build() {
            $.clientTimeout = Codegen.integerProp("clientTimeout").output().arg($.clientTimeout).env("KEYCLOAK_CLIENT_TIMEOUT").def(5).getNullable();
            return $;
        }
    }

}
