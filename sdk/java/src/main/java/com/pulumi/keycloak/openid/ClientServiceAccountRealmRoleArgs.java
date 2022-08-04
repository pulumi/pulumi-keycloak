// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ClientServiceAccountRealmRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClientServiceAccountRealmRoleArgs Empty = new ClientServiceAccountRealmRoleArgs();

    /**
     * The realm that the client and role belong to.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm that the client and role belong to.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * The name of the role that is assigned.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The name of the role that is assigned.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     * The id of the service account that is assigned the role (the service account of the client that &#34;consumes&#34; the role).
     * 
     */
    @Import(name="serviceAccountUserId", required=true)
    private Output<String> serviceAccountUserId;

    /**
     * @return The id of the service account that is assigned the role (the service account of the client that &#34;consumes&#34; the role).
     * 
     */
    public Output<String> serviceAccountUserId() {
        return this.serviceAccountUserId;
    }

    private ClientServiceAccountRealmRoleArgs() {}

    private ClientServiceAccountRealmRoleArgs(ClientServiceAccountRealmRoleArgs $) {
        this.realmId = $.realmId;
        this.role = $.role;
        this.serviceAccountUserId = $.serviceAccountUserId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientServiceAccountRealmRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientServiceAccountRealmRoleArgs $;

        public Builder() {
            $ = new ClientServiceAccountRealmRoleArgs();
        }

        public Builder(ClientServiceAccountRealmRoleArgs defaults) {
            $ = new ClientServiceAccountRealmRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param realmId The realm that the client and role belong to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm that the client and role belong to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param role The name of the role that is assigned.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name of the role that is assigned.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param serviceAccountUserId The id of the service account that is assigned the role (the service account of the client that &#34;consumes&#34; the role).
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountUserId(Output<String> serviceAccountUserId) {
            $.serviceAccountUserId = serviceAccountUserId;
            return this;
        }

        /**
         * @param serviceAccountUserId The id of the service account that is assigned the role (the service account of the client that &#34;consumes&#34; the role).
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountUserId(String serviceAccountUserId) {
            return serviceAccountUserId(Output.of(serviceAccountUserId));
        }

        public ClientServiceAccountRealmRoleArgs build() {
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            $.serviceAccountUserId = Objects.requireNonNull($.serviceAccountUserId, "expected parameter 'serviceAccountUserId' to be non-null");
            return $;
        }
    }

}
