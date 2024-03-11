// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class UserFederatedIdentityArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserFederatedIdentityArgs Empty = new UserFederatedIdentityArgs();

    @Import(name="identityProvider", required=true)
    private Output<String> identityProvider;

    public Output<String> identityProvider() {
        return this.identityProvider;
    }

    @Import(name="userId", required=true)
    private Output<String> userId;

    public Output<String> userId() {
        return this.userId;
    }

    @Import(name="userName", required=true)
    private Output<String> userName;

    public Output<String> userName() {
        return this.userName;
    }

    private UserFederatedIdentityArgs() {}

    private UserFederatedIdentityArgs(UserFederatedIdentityArgs $) {
        this.identityProvider = $.identityProvider;
        this.userId = $.userId;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserFederatedIdentityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserFederatedIdentityArgs $;

        public Builder() {
            $ = new UserFederatedIdentityArgs();
        }

        public Builder(UserFederatedIdentityArgs defaults) {
            $ = new UserFederatedIdentityArgs(Objects.requireNonNull(defaults));
        }

        public Builder identityProvider(Output<String> identityProvider) {
            $.identityProvider = identityProvider;
            return this;
        }

        public Builder identityProvider(String identityProvider) {
            return identityProvider(Output.of(identityProvider));
        }

        public Builder userId(Output<String> userId) {
            $.userId = userId;
            return this;
        }

        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        public Builder userName(Output<String> userName) {
            $.userName = userName;
            return this;
        }

        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public UserFederatedIdentityArgs build() {
            if ($.identityProvider == null) {
                throw new MissingRequiredPropertyException("UserFederatedIdentityArgs", "identityProvider");
            }
            if ($.userId == null) {
                throw new MissingRequiredPropertyException("UserFederatedIdentityArgs", "userId");
            }
            if ($.userName == null) {
                throw new MissingRequiredPropertyException("UserFederatedIdentityArgs", "userName");
            }
            return $;
        }
    }

}
