// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Ldap.Inputs
{

    public sealed class UserFederationKerberosArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the kerberos realm, e.g. FOO.LOCAL.
        /// </summary>
        [Input("kerberosRealm", required: true)]
        public Input<string> KerberosRealm { get; set; } = null!;

        /// <summary>
        /// Path to the kerberos keytab file on the server with credentials of the service principal.
        /// </summary>
        [Input("keyTab", required: true)]
        public Input<string> KeyTab { get; set; } = null!;

        /// <summary>
        /// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
        /// </summary>
        [Input("serverPrincipal", required: true)]
        public Input<string> ServerPrincipal { get; set; } = null!;

        /// <summary>
        /// Use kerberos login module instead of ldap service api. Defaults to `false`.
        /// </summary>
        [Input("useKerberosForPasswordAuthentication")]
        public Input<bool>? UseKerberosForPasswordAuthentication { get; set; }

        public UserFederationKerberosArgs()
        {
        }
        public static new UserFederationKerberosArgs Empty => new UserFederationKerberosArgs();
    }
}
