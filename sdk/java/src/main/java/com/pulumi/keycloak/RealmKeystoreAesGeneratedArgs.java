// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmKeystoreAesGeneratedArgs extends com.pulumi.resources.ResourceArgs {

    public static final RealmKeystoreAesGeneratedArgs Empty = new RealmKeystoreAesGeneratedArgs();

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
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this keystore exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
     * 
     */
    @Import(name="secretSize")
    private @Nullable Output<Integer> secretSize;

    /**
     * @return Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
     * 
     */
    public Optional<Output<Integer>> secretSize() {
        return Optional.ofNullable(this.secretSize);
    }

    private RealmKeystoreAesGeneratedArgs() {}

    private RealmKeystoreAesGeneratedArgs(RealmKeystoreAesGeneratedArgs $) {
        this.active = $.active;
        this.enabled = $.enabled;
        this.name = $.name;
        this.priority = $.priority;
        this.realmId = $.realmId;
        this.secretSize = $.secretSize;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmKeystoreAesGeneratedArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmKeystoreAesGeneratedArgs $;

        public Builder() {
            $ = new RealmKeystoreAesGeneratedArgs();
        }

        public Builder(RealmKeystoreAesGeneratedArgs defaults) {
            $ = new RealmKeystoreAesGeneratedArgs(Objects.requireNonNull(defaults));
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

        /**
         * @param secretSize Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
         * 
         * @return builder
         * 
         */
        public Builder secretSize(@Nullable Output<Integer> secretSize) {
            $.secretSize = secretSize;
            return this;
        }

        /**
         * @param secretSize Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
         * 
         * @return builder
         * 
         */
        public Builder secretSize(Integer secretSize) {
            return secretSize(Output.of(secretSize));
        }

        public RealmKeystoreAesGeneratedArgs build() {
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}