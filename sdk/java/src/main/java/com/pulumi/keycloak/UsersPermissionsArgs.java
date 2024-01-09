// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.inputs.UsersPermissionsImpersonateScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsManageGroupMembershipScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsManageScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsMapRolesScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsUserImpersonatedScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsViewScopeArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UsersPermissionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final UsersPermissionsArgs Empty = new UsersPermissionsArgs();

    @Import(name="impersonateScope")
    private @Nullable Output<UsersPermissionsImpersonateScopeArgs> impersonateScope;

    public Optional<Output<UsersPermissionsImpersonateScopeArgs>> impersonateScope() {
        return Optional.ofNullable(this.impersonateScope);
    }

    @Import(name="manageGroupMembershipScope")
    private @Nullable Output<UsersPermissionsManageGroupMembershipScopeArgs> manageGroupMembershipScope;

    public Optional<Output<UsersPermissionsManageGroupMembershipScopeArgs>> manageGroupMembershipScope() {
        return Optional.ofNullable(this.manageGroupMembershipScope);
    }

    @Import(name="manageScope")
    private @Nullable Output<UsersPermissionsManageScopeArgs> manageScope;

    public Optional<Output<UsersPermissionsManageScopeArgs>> manageScope() {
        return Optional.ofNullable(this.manageScope);
    }

    @Import(name="mapRolesScope")
    private @Nullable Output<UsersPermissionsMapRolesScopeArgs> mapRolesScope;

    public Optional<Output<UsersPermissionsMapRolesScopeArgs>> mapRolesScope() {
        return Optional.ofNullable(this.mapRolesScope);
    }

    @Import(name="realmId", required=true)
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }

    @Import(name="userImpersonatedScope")
    private @Nullable Output<UsersPermissionsUserImpersonatedScopeArgs> userImpersonatedScope;

    public Optional<Output<UsersPermissionsUserImpersonatedScopeArgs>> userImpersonatedScope() {
        return Optional.ofNullable(this.userImpersonatedScope);
    }

    @Import(name="viewScope")
    private @Nullable Output<UsersPermissionsViewScopeArgs> viewScope;

    public Optional<Output<UsersPermissionsViewScopeArgs>> viewScope() {
        return Optional.ofNullable(this.viewScope);
    }

    private UsersPermissionsArgs() {}

    private UsersPermissionsArgs(UsersPermissionsArgs $) {
        this.impersonateScope = $.impersonateScope;
        this.manageGroupMembershipScope = $.manageGroupMembershipScope;
        this.manageScope = $.manageScope;
        this.mapRolesScope = $.mapRolesScope;
        this.realmId = $.realmId;
        this.userImpersonatedScope = $.userImpersonatedScope;
        this.viewScope = $.viewScope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UsersPermissionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UsersPermissionsArgs $;

        public Builder() {
            $ = new UsersPermissionsArgs();
        }

        public Builder(UsersPermissionsArgs defaults) {
            $ = new UsersPermissionsArgs(Objects.requireNonNull(defaults));
        }

        public Builder impersonateScope(@Nullable Output<UsersPermissionsImpersonateScopeArgs> impersonateScope) {
            $.impersonateScope = impersonateScope;
            return this;
        }

        public Builder impersonateScope(UsersPermissionsImpersonateScopeArgs impersonateScope) {
            return impersonateScope(Output.of(impersonateScope));
        }

        public Builder manageGroupMembershipScope(@Nullable Output<UsersPermissionsManageGroupMembershipScopeArgs> manageGroupMembershipScope) {
            $.manageGroupMembershipScope = manageGroupMembershipScope;
            return this;
        }

        public Builder manageGroupMembershipScope(UsersPermissionsManageGroupMembershipScopeArgs manageGroupMembershipScope) {
            return manageGroupMembershipScope(Output.of(manageGroupMembershipScope));
        }

        public Builder manageScope(@Nullable Output<UsersPermissionsManageScopeArgs> manageScope) {
            $.manageScope = manageScope;
            return this;
        }

        public Builder manageScope(UsersPermissionsManageScopeArgs manageScope) {
            return manageScope(Output.of(manageScope));
        }

        public Builder mapRolesScope(@Nullable Output<UsersPermissionsMapRolesScopeArgs> mapRolesScope) {
            $.mapRolesScope = mapRolesScope;
            return this;
        }

        public Builder mapRolesScope(UsersPermissionsMapRolesScopeArgs mapRolesScope) {
            return mapRolesScope(Output.of(mapRolesScope));
        }

        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder userImpersonatedScope(@Nullable Output<UsersPermissionsUserImpersonatedScopeArgs> userImpersonatedScope) {
            $.userImpersonatedScope = userImpersonatedScope;
            return this;
        }

        public Builder userImpersonatedScope(UsersPermissionsUserImpersonatedScopeArgs userImpersonatedScope) {
            return userImpersonatedScope(Output.of(userImpersonatedScope));
        }

        public Builder viewScope(@Nullable Output<UsersPermissionsViewScopeArgs> viewScope) {
            $.viewScope = viewScope;
            return this;
        }

        public Builder viewScope(UsersPermissionsViewScopeArgs viewScope) {
            return viewScope(Output.of(viewScope));
        }

        public UsersPermissionsArgs build() {
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("UsersPermissionsArgs", "realmId");
            }
            return $;
        }
    }

}
