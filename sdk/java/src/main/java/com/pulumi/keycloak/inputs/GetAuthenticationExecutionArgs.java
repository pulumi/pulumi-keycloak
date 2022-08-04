// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetAuthenticationExecutionArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthenticationExecutionArgs Empty = new GetAuthenticationExecutionArgs();

    /**
     * The alias of the flow this execution is attached to.
     * 
     */
    @Import(name="parentFlowAlias", required=true)
    private Output<String> parentFlowAlias;

    /**
     * @return The alias of the flow this execution is attached to.
     * 
     */
    public Output<String> parentFlowAlias() {
        return this.parentFlowAlias;
    }

    /**
     * The name of the provider. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser&#39;s development tools. This was previously known as the &#34;authenticator&#34;.
     * 
     */
    @Import(name="providerId", required=true)
    private Output<String> providerId;

    /**
     * @return The name of the provider. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser&#39;s development tools. This was previously known as the &#34;authenticator&#34;.
     * 
     */
    public Output<String> providerId() {
        return this.providerId;
    }

    /**
     * The realm the authentication execution exists in.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm the authentication execution exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private GetAuthenticationExecutionArgs() {}

    private GetAuthenticationExecutionArgs(GetAuthenticationExecutionArgs $) {
        this.parentFlowAlias = $.parentFlowAlias;
        this.providerId = $.providerId;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthenticationExecutionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthenticationExecutionArgs $;

        public Builder() {
            $ = new GetAuthenticationExecutionArgs();
        }

        public Builder(GetAuthenticationExecutionArgs defaults) {
            $ = new GetAuthenticationExecutionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param parentFlowAlias The alias of the flow this execution is attached to.
         * 
         * @return builder
         * 
         */
        public Builder parentFlowAlias(Output<String> parentFlowAlias) {
            $.parentFlowAlias = parentFlowAlias;
            return this;
        }

        /**
         * @param parentFlowAlias The alias of the flow this execution is attached to.
         * 
         * @return builder
         * 
         */
        public Builder parentFlowAlias(String parentFlowAlias) {
            return parentFlowAlias(Output.of(parentFlowAlias));
        }

        /**
         * @param providerId The name of the provider. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser&#39;s development tools. This was previously known as the &#34;authenticator&#34;.
         * 
         * @return builder
         * 
         */
        public Builder providerId(Output<String> providerId) {
            $.providerId = providerId;
            return this;
        }

        /**
         * @param providerId The name of the provider. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser&#39;s development tools. This was previously known as the &#34;authenticator&#34;.
         * 
         * @return builder
         * 
         */
        public Builder providerId(String providerId) {
            return providerId(Output.of(providerId));
        }

        /**
         * @param realmId The realm the authentication execution exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm the authentication execution exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public GetAuthenticationExecutionArgs build() {
            $.parentFlowAlias = Objects.requireNonNull($.parentFlowAlias, "expected parameter 'parentFlowAlias' to be non-null");
            $.providerId = Objects.requireNonNull($.providerId, "expected parameter 'providerId' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}
