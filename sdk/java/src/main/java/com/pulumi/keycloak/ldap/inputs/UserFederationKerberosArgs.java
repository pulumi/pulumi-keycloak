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


public final class UserFederationKerberosArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserFederationKerberosArgs Empty = new UserFederationKerberosArgs();

    /**
     * The name of the kerberos realm, e.g. FOO.LOCAL.
     * 
     */
    @Import(name="kerberosRealm", required=true)
    private Output<String> kerberosRealm;

    /**
     * @return The name of the kerberos realm, e.g. FOO.LOCAL.
     * 
     */
    public Output<String> kerberosRealm() {
        return this.kerberosRealm;
    }

    /**
     * Path to the kerberos keytab file on the server with credentials of the service principal.
     * 
     */
    @Import(name="keyTab", required=true)
    private Output<String> keyTab;

    /**
     * @return Path to the kerberos keytab file on the server with credentials of the service principal.
     * 
     */
    public Output<String> keyTab() {
        return this.keyTab;
    }

    /**
     * The kerberos server principal, e.g. &#39;HTTP/host.foo.com@FOO.LOCAL&#39;.
     * 
     */
    @Import(name="serverPrincipal", required=true)
    private Output<String> serverPrincipal;

    /**
     * @return The kerberos server principal, e.g. &#39;HTTP/host.foo.com@FOO.LOCAL&#39;.
     * 
     */
    public Output<String> serverPrincipal() {
        return this.serverPrincipal;
    }

    /**
     * Use kerberos login module instead of ldap service api. Defaults to `false`.
     * 
     */
    @Import(name="useKerberosForPasswordAuthentication")
    private @Nullable Output<Boolean> useKerberosForPasswordAuthentication;

    /**
     * @return Use kerberos login module instead of ldap service api. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> useKerberosForPasswordAuthentication() {
        return Optional.ofNullable(this.useKerberosForPasswordAuthentication);
    }

    private UserFederationKerberosArgs() {}

    private UserFederationKerberosArgs(UserFederationKerberosArgs $) {
        this.kerberosRealm = $.kerberosRealm;
        this.keyTab = $.keyTab;
        this.serverPrincipal = $.serverPrincipal;
        this.useKerberosForPasswordAuthentication = $.useKerberosForPasswordAuthentication;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserFederationKerberosArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserFederationKerberosArgs $;

        public Builder() {
            $ = new UserFederationKerberosArgs();
        }

        public Builder(UserFederationKerberosArgs defaults) {
            $ = new UserFederationKerberosArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param kerberosRealm The name of the kerberos realm, e.g. FOO.LOCAL.
         * 
         * @return builder
         * 
         */
        public Builder kerberosRealm(Output<String> kerberosRealm) {
            $.kerberosRealm = kerberosRealm;
            return this;
        }

        /**
         * @param kerberosRealm The name of the kerberos realm, e.g. FOO.LOCAL.
         * 
         * @return builder
         * 
         */
        public Builder kerberosRealm(String kerberosRealm) {
            return kerberosRealm(Output.of(kerberosRealm));
        }

        /**
         * @param keyTab Path to the kerberos keytab file on the server with credentials of the service principal.
         * 
         * @return builder
         * 
         */
        public Builder keyTab(Output<String> keyTab) {
            $.keyTab = keyTab;
            return this;
        }

        /**
         * @param keyTab Path to the kerberos keytab file on the server with credentials of the service principal.
         * 
         * @return builder
         * 
         */
        public Builder keyTab(String keyTab) {
            return keyTab(Output.of(keyTab));
        }

        /**
         * @param serverPrincipal The kerberos server principal, e.g. &#39;HTTP/host.foo.com@FOO.LOCAL&#39;.
         * 
         * @return builder
         * 
         */
        public Builder serverPrincipal(Output<String> serverPrincipal) {
            $.serverPrincipal = serverPrincipal;
            return this;
        }

        /**
         * @param serverPrincipal The kerberos server principal, e.g. &#39;HTTP/host.foo.com@FOO.LOCAL&#39;.
         * 
         * @return builder
         * 
         */
        public Builder serverPrincipal(String serverPrincipal) {
            return serverPrincipal(Output.of(serverPrincipal));
        }

        /**
         * @param useKerberosForPasswordAuthentication Use kerberos login module instead of ldap service api. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder useKerberosForPasswordAuthentication(@Nullable Output<Boolean> useKerberosForPasswordAuthentication) {
            $.useKerberosForPasswordAuthentication = useKerberosForPasswordAuthentication;
            return this;
        }

        /**
         * @param useKerberosForPasswordAuthentication Use kerberos login module instead of ldap service api. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder useKerberosForPasswordAuthentication(Boolean useKerberosForPasswordAuthentication) {
            return useKerberosForPasswordAuthentication(Output.of(useKerberosForPasswordAuthentication));
        }

        public UserFederationKerberosArgs build() {
            $.kerberosRealm = Objects.requireNonNull($.kerberosRealm, "expected parameter 'kerberosRealm' to be non-null");
            $.keyTab = Objects.requireNonNull($.keyTab, "expected parameter 'keyTab' to be non-null");
            $.serverPrincipal = Objects.requireNonNull($.serverPrincipal, "expected parameter 'serverPrincipal' to be non-null");
            return $;
        }
    }

}
