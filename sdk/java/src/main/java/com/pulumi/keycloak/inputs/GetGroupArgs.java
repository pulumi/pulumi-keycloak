// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetGroupArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupArgs Empty = new GetGroupArgs();

    /**
     * The name of the group. If there are multiple groups match `name`, the first result will be returned.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the group. If there are multiple groups match `name`, the first result will be returned.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The realm this group exists within.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this group exists within.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private GetGroupArgs() {}

    private GetGroupArgs(GetGroupArgs $) {
        this.name = $.name;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupArgs $;

        public Builder() {
            $ = new GetGroupArgs();
        }

        public Builder(GetGroupArgs defaults) {
            $ = new GetGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the group. If there are multiple groups match `name`, the first result will be returned.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the group. If there are multiple groups match `name`, the first result will be returned.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm this group exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this group exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public GetGroupArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetGroupArgs", "name");
            }
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("GetGroupArgs", "realmId");
            }
            return $;
        }
    }

}
