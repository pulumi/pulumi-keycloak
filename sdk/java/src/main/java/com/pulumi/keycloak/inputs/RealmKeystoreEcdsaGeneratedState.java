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


public final class RealmKeystoreEcdsaGeneratedState extends com.pulumi.resources.ResourceArgs {

    public static final RealmKeystoreEcdsaGeneratedState Empty = new RealmKeystoreEcdsaGeneratedState();

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
     * Elliptic Curve used in ECDSA. Defaults to `P-256`.
     * 
     */
    @Import(name="ellipticCurveKey")
    private @Nullable Output<String> ellipticCurveKey;

    /**
     * @return Elliptic Curve used in ECDSA. Defaults to `P-256`.
     * 
     */
    public Optional<Output<String>> ellipticCurveKey() {
        return Optional.ofNullable(this.ellipticCurveKey);
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

    private RealmKeystoreEcdsaGeneratedState() {}

    private RealmKeystoreEcdsaGeneratedState(RealmKeystoreEcdsaGeneratedState $) {
        this.active = $.active;
        this.ellipticCurveKey = $.ellipticCurveKey;
        this.enabled = $.enabled;
        this.name = $.name;
        this.priority = $.priority;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmKeystoreEcdsaGeneratedState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmKeystoreEcdsaGeneratedState $;

        public Builder() {
            $ = new RealmKeystoreEcdsaGeneratedState();
        }

        public Builder(RealmKeystoreEcdsaGeneratedState defaults) {
            $ = new RealmKeystoreEcdsaGeneratedState(Objects.requireNonNull(defaults));
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
         * @param ellipticCurveKey Elliptic Curve used in ECDSA. Defaults to `P-256`.
         * 
         * @return builder
         * 
         */
        public Builder ellipticCurveKey(@Nullable Output<String> ellipticCurveKey) {
            $.ellipticCurveKey = ellipticCurveKey;
            return this;
        }

        /**
         * @param ellipticCurveKey Elliptic Curve used in ECDSA. Defaults to `P-256`.
         * 
         * @return builder
         * 
         */
        public Builder ellipticCurveKey(String ellipticCurveKey) {
            return ellipticCurveKey(Output.of(ellipticCurveKey));
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

        public RealmKeystoreEcdsaGeneratedState build() {
            return $;
        }
    }

}