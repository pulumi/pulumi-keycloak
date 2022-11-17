// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmKeystoreRsaGeneratedState extends com.pulumi.resources.ResourceArgs {

    public static final RealmKeystoreRsaGeneratedState Empty = new RealmKeystoreRsaGeneratedState();

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
     * Intended algorithm for the key. Defaults to `RS256`
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return Intended algorithm for the key. Defaults to `RS256`
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
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
     * Size for the generated keys
     * 
     */
    @Import(name="keySize")
    private @Nullable Output<Integer> keySize;

    /**
     * @return Size for the generated keys
     * 
     */
    public Optional<Output<Integer>> keySize() {
        return Optional.ofNullable(this.keySize);
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
     * The realm this keystore exists in.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm this keystore exists in.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private RealmKeystoreRsaGeneratedState() {}

    private RealmKeystoreRsaGeneratedState(RealmKeystoreRsaGeneratedState $) {
        this.active = $.active;
        this.algorithm = $.algorithm;
        this.enabled = $.enabled;
        this.keySize = $.keySize;
        this.name = $.name;
        this.priority = $.priority;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmKeystoreRsaGeneratedState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmKeystoreRsaGeneratedState $;

        public Builder() {
            $ = new RealmKeystoreRsaGeneratedState();
        }

        public Builder(RealmKeystoreRsaGeneratedState defaults) {
            $ = new RealmKeystoreRsaGeneratedState(Objects.requireNonNull(defaults));
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
         * @param algorithm Intended algorithm for the key. Defaults to `RS256`
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Intended algorithm for the key. Defaults to `RS256`
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
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
         * @param keySize Size for the generated keys
         * 
         * @return builder
         * 
         */
        public Builder keySize(@Nullable Output<Integer> keySize) {
            $.keySize = keySize;
            return this;
        }

        /**
         * @param keySize Size for the generated keys
         * 
         * @return builder
         * 
         */
        public Builder keySize(Integer keySize) {
            return keySize(Output.of(keySize));
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
         * @param realmId The realm this keystore exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
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

        public RealmKeystoreRsaGeneratedState build() {
            return $;
        }
    }

}