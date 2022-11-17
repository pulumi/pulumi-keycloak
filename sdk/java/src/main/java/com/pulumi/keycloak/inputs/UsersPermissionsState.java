// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.keycloak.inputs.UsersPermissionsImpersonateScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsManageGroupMembershipScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsManageScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsMapRolesScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsUserImpersonatedScopeArgs;
import com.pulumi.keycloak.inputs.UsersPermissionsViewScopeArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UsersPermissionsState extends com.pulumi.resources.ResourceArgs {

    public static final UsersPermissionsState Empty = new UsersPermissionsState();

    /**
     * Resource server id representing the realm management client on which this permission is managed
     * 
     */
    @Import(name="authorizationResourceServerId")
    private @Nullable Output<String> authorizationResourceServerId;

    /**
     * @return Resource server id representing the realm management client on which this permission is managed
     * 
     */
    public Optional<Output<String>> authorizationResourceServerId() {
        return Optional.ofNullable(this.authorizationResourceServerId);
    }

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

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

    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
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

    private UsersPermissionsState() {}

    private UsersPermissionsState(UsersPermissionsState $) {
        this.authorizationResourceServerId = $.authorizationResourceServerId;
        this.enabled = $.enabled;
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
    public static Builder builder(UsersPermissionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UsersPermissionsState $;

        public Builder() {
            $ = new UsersPermissionsState();
        }

        public Builder(UsersPermissionsState defaults) {
            $ = new UsersPermissionsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorizationResourceServerId Resource server id representing the realm management client on which this permission is managed
         * 
         * @return builder
         * 
         */
        public Builder authorizationResourceServerId(@Nullable Output<String> authorizationResourceServerId) {
            $.authorizationResourceServerId = authorizationResourceServerId;
            return this;
        }

        /**
         * @param authorizationResourceServerId Resource server id representing the realm management client on which this permission is managed
         * 
         * @return builder
         * 
         */
        public Builder authorizationResourceServerId(String authorizationResourceServerId) {
            return authorizationResourceServerId(Output.of(authorizationResourceServerId));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
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

        public Builder realmId(@Nullable Output<String> realmId) {
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

        public UsersPermissionsState build() {
            return $;
        }
    }

}