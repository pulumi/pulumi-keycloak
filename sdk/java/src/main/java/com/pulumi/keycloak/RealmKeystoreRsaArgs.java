// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmKeystoreRsaArgs extends com.pulumi.resources.ResourceArgs {

    public static final RealmKeystoreRsaArgs Empty = new RealmKeystoreRsaArgs();

    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     * 
     */
    @Import(name="active")
    private @Nullable Output<Boolean> active;

    /**
     * @return When `false`, key in not used for signing. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    /**
     * Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * X509 Certificate encoded in PEM format.
     * 
     */
    @Import(name="certificate", required=true)
    private Output<String> certificate;

    /**
     * @return X509 Certificate encoded in PEM format.
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }

    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return When `false`, key is not accessible in this realm. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Display name of provider when linked in admin console.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Display name of provider when linked in admin console.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Priority for the provider. Defaults to `0`
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return Priority for the provider. Defaults to `0`
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * Private RSA Key encoded in PEM format.
     * 
     */
    @Import(name="privateKey", required=true)
    private Output<String> privateKey;

    /**
     * @return Private RSA Key encoded in PEM format.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }

    /**
     * Use `rsa` for signing keys, `rsa-enc` for encryption keys
     * 
     */
    @Import(name="providerId")
    private @Nullable Output<String> providerId;

    /**
     * @return Use `rsa` for signing keys, `rsa-enc` for encryption keys
     * 
     */
    public Optional<Output<String>> providerId() {
        return Optional.ofNullable(this.providerId);
    }

    /**
     * The realm this keystore exists in.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this keystore exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private RealmKeystoreRsaArgs() {}

    private RealmKeystoreRsaArgs(RealmKeystoreRsaArgs $) {
        this.active = $.active;
        this.algorithm = $.algorithm;
        this.certificate = $.certificate;
        this.enabled = $.enabled;
        this.name = $.name;
        this.priority = $.priority;
        this.privateKey = $.privateKey;
        this.providerId = $.providerId;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmKeystoreRsaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmKeystoreRsaArgs $;

        public Builder() {
            $ = new RealmKeystoreRsaArgs();
        }

        public Builder(RealmKeystoreRsaArgs defaults) {
            $ = new RealmKeystoreRsaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param active When `false`, key in not used for signing. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        /**
         * @param active When `false`, key in not used for signing. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        /**
         * @param algorithm Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param certificate X509 Certificate encoded in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder certificate(Output<String> certificate) {
            $.certificate = certificate;
            return this;
        }

        /**
         * @param certificate X509 Certificate encoded in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder certificate(String certificate) {
            return certificate(Output.of(certificate));
        }

        /**
         * @param enabled When `false`, key is not accessible in this realm. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled When `false`, key is not accessible in this realm. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param name Display name of provider when linked in admin console.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Display name of provider when linked in admin console.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param priority Priority for the provider. Defaults to `0`
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority Priority for the provider. Defaults to `0`
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param privateKey Private RSA Key encoded in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey Private RSA Key encoded in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param providerId Use `rsa` for signing keys, `rsa-enc` for encryption keys
         * 
         * @return builder
         * 
         */
        public Builder providerId(@Nullable Output<String> providerId) {
            $.providerId = providerId;
            return this;
        }

        /**
         * @param providerId Use `rsa` for signing keys, `rsa-enc` for encryption keys
         * 
         * @return builder
         * 
         */
        public Builder providerId(String providerId) {
            return providerId(Output.of(providerId));
        }

        /**
         * @param realmId The realm this keystore exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this keystore exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public RealmKeystoreRsaArgs build() {
            if ($.certificate == null) {
                throw new MissingRequiredPropertyException("RealmKeystoreRsaArgs", "certificate");
            }
            if ($.privateKey == null) {
                throw new MissingRequiredPropertyException("RealmKeystoreRsaArgs", "privateKey");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("RealmKeystoreRsaArgs", "realmId");
            }
            return $;
        }
    }

}
