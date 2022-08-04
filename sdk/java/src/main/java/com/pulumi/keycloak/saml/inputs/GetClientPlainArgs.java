// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetClientPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetClientPlainArgs Empty = new GetClientPlainArgs();

    /**
     * The client id (not its unique ID).
     * 
     */
    @Import(name="clientId", required=true)
    private String clientId;

    /**
     * @return The client id (not its unique ID).
     * 
     */
    public String clientId() {
        return this.clientId;
    }

    /**
     * The realm id.
     * 
     */
    @Import(name="realmId", required=true)
    private String realmId;

    /**
     * @return The realm id.
     * 
     */
    public String realmId() {
        return this.realmId;
    }

    private GetClientPlainArgs() {}

    private GetClientPlainArgs(GetClientPlainArgs $) {
        this.clientId = $.clientId;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetClientPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetClientPlainArgs $;

        public Builder() {
            $ = new GetClientPlainArgs();
        }

        public Builder(GetClientPlainArgs defaults) {
            $ = new GetClientPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId The client id (not its unique ID).
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param realmId The realm id.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            $.realmId = realmId;
            return this;
        }

        public GetClientPlainArgs build() {
            $.clientId = Objects.requireNonNull($.clientId, "expected parameter 'clientId' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}
