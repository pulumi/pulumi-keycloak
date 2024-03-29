// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetUserPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserPlainArgs Empty = new GetUserPlainArgs();

    /**
     * The realm this user belongs to.
     * 
     */
    @Import(name="realmId", required=true)
    private String realmId;

    /**
     * @return The realm this user belongs to.
     * 
     */
    public String realmId() {
        return this.realmId;
    }

    /**
     * The unique username of this user.
     * 
     */
    @Import(name="username", required=true)
    private String username;

    /**
     * @return The unique username of this user.
     * 
     */
    public String username() {
        return this.username;
    }

    private GetUserPlainArgs() {}

    private GetUserPlainArgs(GetUserPlainArgs $) {
        this.realmId = $.realmId;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserPlainArgs $;

        public Builder() {
            $ = new GetUserPlainArgs();
        }

        public Builder(GetUserPlainArgs defaults) {
            $ = new GetUserPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param realmId The realm this user belongs to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param username The unique username of this user.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            $.username = username;
            return this;
        }

        public GetUserPlainArgs build() {
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("GetUserPlainArgs", "realmId");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("GetUserPlainArgs", "username");
            }
            return $;
        }
    }

}
