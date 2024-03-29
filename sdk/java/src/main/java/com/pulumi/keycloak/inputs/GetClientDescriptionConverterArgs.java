// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetClientDescriptionConverterArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetClientDescriptionConverterArgs Empty = new GetClientDescriptionConverterArgs();

    /**
     * The body of the request to convert.
     * 
     */
    @Import(name="body", required=true)
    private Output<String> body;

    /**
     * @return The body of the request to convert.
     * 
     */
    public Output<String> body() {
        return this.body;
    }

    /**
     * The realm to use for the client description converter API call.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm to use for the client description converter API call.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private GetClientDescriptionConverterArgs() {}

    private GetClientDescriptionConverterArgs(GetClientDescriptionConverterArgs $) {
        this.body = $.body;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetClientDescriptionConverterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetClientDescriptionConverterArgs $;

        public Builder() {
            $ = new GetClientDescriptionConverterArgs();
        }

        public Builder(GetClientDescriptionConverterArgs defaults) {
            $ = new GetClientDescriptionConverterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param body The body of the request to convert.
         * 
         * @return builder
         * 
         */
        public Builder body(Output<String> body) {
            $.body = body;
            return this;
        }

        /**
         * @param body The body of the request to convert.
         * 
         * @return builder
         * 
         */
        public Builder body(String body) {
            return body(Output.of(body));
        }

        /**
         * @param realmId The realm to use for the client description converter API call.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm to use for the client description converter API call.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public GetClientDescriptionConverterArgs build() {
            if ($.body == null) {
                throw new MissingRequiredPropertyException("GetClientDescriptionConverterArgs", "body");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("GetClientDescriptionConverterArgs", "realmId");
            }
            return $;
        }
    }

}
