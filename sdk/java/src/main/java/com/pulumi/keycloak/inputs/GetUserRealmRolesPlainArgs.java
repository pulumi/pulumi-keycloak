// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetUserRealmRolesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserRealmRolesPlainArgs Empty = new GetUserRealmRolesPlainArgs();

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
     * The ID of the user to query realm roles for.
     * 
     */
    @Import(name="userId", required=true)
    private String userId;

    /**
     * @return The ID of the user to query realm roles for.
     * 
     */
    public String userId() {
        return this.userId;
    }

    private GetUserRealmRolesPlainArgs() {}

    private GetUserRealmRolesPlainArgs(GetUserRealmRolesPlainArgs $) {
        this.realmId = $.realmId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserRealmRolesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserRealmRolesPlainArgs $;

        public Builder() {
            $ = new GetUserRealmRolesPlainArgs();
        }

        public Builder(GetUserRealmRolesPlainArgs defaults) {
            $ = new GetUserRealmRolesPlainArgs(Objects.requireNonNull(defaults));
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
         * @param userId The ID of the user to query realm roles for.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            $.userId = userId;
            return this;
        }

        public GetUserRealmRolesPlainArgs build() {
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            $.userId = Objects.requireNonNull($.userId, "expected parameter 'userId' to be non-null");
            return $;
        }
    }

}
