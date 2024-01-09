// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetUserArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserArgs Empty = new GetUserArgs();

    /**
     * The realm this user belongs to.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this user belongs to.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * The unique username of this user.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The unique username of this user.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private GetUserArgs() {}

    private GetUserArgs(GetUserArgs $) {
        this.realmId = $.realmId;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserArgs $;

        public Builder() {
            $ = new GetUserArgs();
        }

        public Builder(GetUserArgs defaults) {
            $ = new GetUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param realmId The realm this user belongs to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this user belongs to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param username The unique username of this user.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The unique username of this user.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public GetUserArgs build() {
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("GetUserArgs", "realmId");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("GetUserArgs", "username");
            }
            return $;
        }
    }

}
