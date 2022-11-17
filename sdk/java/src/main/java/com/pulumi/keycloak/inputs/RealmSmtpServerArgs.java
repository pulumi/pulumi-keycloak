// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.keycloak.inputs.RealmSmtpServerAuthArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmSmtpServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final RealmSmtpServerArgs Empty = new RealmSmtpServerArgs();

    /**
     * Enables authentication to the SMTP server.  This block supports the following arguments:
     * 
     */
    @Import(name="auth")
    private @Nullable Output<RealmSmtpServerAuthArgs> auth;

    /**
     * @return Enables authentication to the SMTP server.  This block supports the following arguments:
     * 
     */
    public Optional<Output<RealmSmtpServerAuthArgs>> auth() {
        return Optional.ofNullable(this.auth);
    }

    /**
     * The email address uses for bounces.
     * 
     */
    @Import(name="envelopeFrom")
    private @Nullable Output<String> envelopeFrom;

    /**
     * @return The email address uses for bounces.
     * 
     */
    public Optional<Output<String>> envelopeFrom() {
        return Optional.ofNullable(this.envelopeFrom);
    }

    /**
     * The email address for the sender.
     * 
     */
    @Import(name="from", required=true)
    private Output<String> from;

    /**
     * @return The email address for the sender.
     * 
     */
    public Output<String> from() {
        return this.from;
    }

    /**
     * The display name of the sender email address.
     * 
     */
    @Import(name="fromDisplayName")
    private @Nullable Output<String> fromDisplayName;

    /**
     * @return The display name of the sender email address.
     * 
     */
    public Optional<Output<String>> fromDisplayName() {
        return Optional.ofNullable(this.fromDisplayName);
    }

    /**
     * The host of the SMTP server.
     * 
     */
    @Import(name="host", required=true)
    private Output<String> host;

    /**
     * @return The host of the SMTP server.
     * 
     */
    public Output<String> host() {
        return this.host;
    }

    /**
     * The port of the SMTP server (defaults to 25).
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return The port of the SMTP server (defaults to 25).
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The &#34;reply to&#34; email address.
     * 
     */
    @Import(name="replyTo")
    private @Nullable Output<String> replyTo;

    /**
     * @return The &#34;reply to&#34; email address.
     * 
     */
    public Optional<Output<String>> replyTo() {
        return Optional.ofNullable(this.replyTo);
    }

    /**
     * The display name of the &#34;reply to&#34; email address.
     * 
     */
    @Import(name="replyToDisplayName")
    private @Nullable Output<String> replyToDisplayName;

    /**
     * @return The display name of the &#34;reply to&#34; email address.
     * 
     */
    public Optional<Output<String>> replyToDisplayName() {
        return Optional.ofNullable(this.replyToDisplayName);
    }

    /**
     * When `true`, enables SSL. Defaults to `false`.
     * 
     */
    @Import(name="ssl")
    private @Nullable Output<Boolean> ssl;

    /**
     * @return When `true`, enables SSL. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> ssl() {
        return Optional.ofNullable(this.ssl);
    }

    /**
     * When `true`, enables StartTLS. Defaults to `false`.
     * 
     */
    @Import(name="starttls")
    private @Nullable Output<Boolean> starttls;

    /**
     * @return When `true`, enables StartTLS. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> starttls() {
        return Optional.ofNullable(this.starttls);
    }

    private RealmSmtpServerArgs() {}

    private RealmSmtpServerArgs(RealmSmtpServerArgs $) {
        this.auth = $.auth;
        this.envelopeFrom = $.envelopeFrom;
        this.from = $.from;
        this.fromDisplayName = $.fromDisplayName;
        this.host = $.host;
        this.port = $.port;
        this.replyTo = $.replyTo;
        this.replyToDisplayName = $.replyToDisplayName;
        this.ssl = $.ssl;
        this.starttls = $.starttls;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmSmtpServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmSmtpServerArgs $;

        public Builder() {
            $ = new RealmSmtpServerArgs();
        }

        public Builder(RealmSmtpServerArgs defaults) {
            $ = new RealmSmtpServerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param auth Enables authentication to the SMTP server.  This block supports the following arguments:
         * 
         * @return builder
         * 
         */
        public Builder auth(@Nullable Output<RealmSmtpServerAuthArgs> auth) {
            $.auth = auth;
            return this;
        }

        /**
         * @param auth Enables authentication to the SMTP server.  This block supports the following arguments:
         * 
         * @return builder
         * 
         */
        public Builder auth(RealmSmtpServerAuthArgs auth) {
            return auth(Output.of(auth));
        }

        /**
         * @param envelopeFrom The email address uses for bounces.
         * 
         * @return builder
         * 
         */
        public Builder envelopeFrom(@Nullable Output<String> envelopeFrom) {
            $.envelopeFrom = envelopeFrom;
            return this;
        }

        /**
         * @param envelopeFrom The email address uses for bounces.
         * 
         * @return builder
         * 
         */
        public Builder envelopeFrom(String envelopeFrom) {
            return envelopeFrom(Output.of(envelopeFrom));
        }

        /**
         * @param from The email address for the sender.
         * 
         * @return builder
         * 
         */
        public Builder from(Output<String> from) {
            $.from = from;
            return this;
        }

        /**
         * @param from The email address for the sender.
         * 
         * @return builder
         * 
         */
        public Builder from(String from) {
            return from(Output.of(from));
        }

        /**
         * @param fromDisplayName The display name of the sender email address.
         * 
         * @return builder
         * 
         */
        public Builder fromDisplayName(@Nullable Output<String> fromDisplayName) {
            $.fromDisplayName = fromDisplayName;
            return this;
        }

        /**
         * @param fromDisplayName The display name of the sender email address.
         * 
         * @return builder
         * 
         */
        public Builder fromDisplayName(String fromDisplayName) {
            return fromDisplayName(Output.of(fromDisplayName));
        }

        /**
         * @param host The host of the SMTP server.
         * 
         * @return builder
         * 
         */
        public Builder host(Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The host of the SMTP server.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param port The port of the SMTP server (defaults to 25).
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port of the SMTP server (defaults to 25).
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param replyTo The &#34;reply to&#34; email address.
         * 
         * @return builder
         * 
         */
        public Builder replyTo(@Nullable Output<String> replyTo) {
            $.replyTo = replyTo;
            return this;
        }

        /**
         * @param replyTo The &#34;reply to&#34; email address.
         * 
         * @return builder
         * 
         */
        public Builder replyTo(String replyTo) {
            return replyTo(Output.of(replyTo));
        }

        /**
         * @param replyToDisplayName The display name of the &#34;reply to&#34; email address.
         * 
         * @return builder
         * 
         */
        public Builder replyToDisplayName(@Nullable Output<String> replyToDisplayName) {
            $.replyToDisplayName = replyToDisplayName;
            return this;
        }

        /**
         * @param replyToDisplayName The display name of the &#34;reply to&#34; email address.
         * 
         * @return builder
         * 
         */
        public Builder replyToDisplayName(String replyToDisplayName) {
            return replyToDisplayName(Output.of(replyToDisplayName));
        }

        /**
         * @param ssl When `true`, enables SSL. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ssl(@Nullable Output<Boolean> ssl) {
            $.ssl = ssl;
            return this;
        }

        /**
         * @param ssl When `true`, enables SSL. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ssl(Boolean ssl) {
            return ssl(Output.of(ssl));
        }

        /**
         * @param starttls When `true`, enables StartTLS. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder starttls(@Nullable Output<Boolean> starttls) {
            $.starttls = starttls;
            return this;
        }

        /**
         * @param starttls When `true`, enables StartTLS. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder starttls(Boolean starttls) {
            return starttls(Output.of(starttls));
        }

        public RealmSmtpServerArgs build() {
            $.from = Objects.requireNonNull($.from, "expected parameter 'from' to be non-null");
            $.host = Objects.requireNonNull($.host, "expected parameter 'host' to be non-null");
            return $;
        }
    }

}