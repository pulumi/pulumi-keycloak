// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.UsersPermissionsArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.UsersPermissionsState;
import com.pulumi.keycloak.outputs.UsersPermissionsImpersonateScope;
import com.pulumi.keycloak.outputs.UsersPermissionsManageGroupMembershipScope;
import com.pulumi.keycloak.outputs.UsersPermissionsManageScope;
import com.pulumi.keycloak.outputs.UsersPermissionsMapRolesScope;
import com.pulumi.keycloak.outputs.UsersPermissionsUserImpersonatedScope;
import com.pulumi.keycloak.outputs.UsersPermissionsViewScope;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows you to manage fine-grained permissions for all users in a realm: https://www.keycloak.org/docs/latest/server_admin/#_users-permissions
 * 
 * This is part of a preview Keycloak feature: `admin_fine_grained_authz` (see https://www.keycloak.org/docs/latest/server_admin/#_fine_grain_permissions).
 * This feature can be enabled with the Keycloak option `-Dkeycloak.profile.feature.admin_fine_grained_authz=enabled`. See the
 * example `docker-compose.yml` file for an example.
 * 
 * When enabling fine-grained permissions for users, Keycloak does several things automatically:
 * 1. Enable Authorization on built-in `realm-management` client (if not already enabled).
 * 2. Create a resource representing the users permissions.
 * 3. Create scopes `view`, `manage`, `map-roles`, `manage-group-membership`, `impersonate`, and `user-impersonated`.
 * 4. Create all scope based permission for the scopes and users resources.
 * 
 * &gt; This resource should only be created once per realm.
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.openid.OpenidFunctions;
 * import com.pulumi.keycloak.openid.inputs.GetClientArgs;
 * import com.pulumi.keycloak.openid.ClientPermissions;
 * import com.pulumi.keycloak.openid.ClientPermissionsArgs;
 * import com.pulumi.keycloak.User;
 * import com.pulumi.keycloak.UserArgs;
 * import com.pulumi.keycloak.openid.ClientUserPolicy;
 * import com.pulumi.keycloak.openid.ClientUserPolicyArgs;
 * import com.pulumi.keycloak.UsersPermissions;
 * import com.pulumi.keycloak.UsersPermissionsArgs;
 * import com.pulumi.keycloak.inputs.UsersPermissionsViewScopeArgs;
 * import com.pulumi.keycloak.inputs.UsersPermissionsManageScopeArgs;
 * import com.pulumi.keycloak.inputs.UsersPermissionsMapRolesScopeArgs;
 * import com.pulumi.keycloak.inputs.UsersPermissionsManageGroupMembershipScopeArgs;
 * import com.pulumi.keycloak.inputs.UsersPermissionsImpersonateScopeArgs;
 * import com.pulumi.keycloak.inputs.UsersPermissionsUserImpersonatedScopeArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var realm = new Realm(&#34;realm&#34;, RealmArgs.builder()        
 *             .realm(&#34;my-realm&#34;)
 *             .build());
 * 
 *         final var realmManagement = OpenidFunctions.getClient(GetClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(&#34;realm-management&#34;)
 *             .build());
 * 
 *         var realmManagementPermission = new ClientPermissions(&#34;realmManagementPermission&#34;, ClientPermissionsArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(realmManagement.applyValue(getClientResult -&gt; getClientResult).applyValue(realmManagement -&gt; realmManagement.applyValue(getClientResult -&gt; getClientResult.id())))
 *             .enabled(true)
 *             .build());
 * 
 *         var testUser = new User(&#34;testUser&#34;, UserArgs.builder()        
 *             .realmId(realm.id())
 *             .username(&#34;test-user&#34;)
 *             .email(&#34;test-user@fakedomain.com&#34;)
 *             .firstName(&#34;Testy&#34;)
 *             .lastName(&#34;Tester&#34;)
 *             .build());
 * 
 *         var testClientUserPolicy = new ClientUserPolicy(&#34;testClientUserPolicy&#34;, ClientUserPolicyArgs.builder()        
 *             .realmId(realm.id())
 *             .resourceServerId(realmManagement.applyValue(getClientResult -&gt; getClientResult).applyValue(realmManagement -&gt; realmManagement.applyValue(getClientResult -&gt; getClientResult.id())))
 *             .users(testUser.id())
 *             .logic(&#34;POSITIVE&#34;)
 *             .decisionStrategy(&#34;UNANIMOUS&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(realmManagementPermission)
 *                 .build());
 * 
 *         var usersPermissions = new UsersPermissions(&#34;usersPermissions&#34;, UsersPermissionsArgs.builder()        
 *             .realmId(realm.id())
 *             .viewScope(UsersPermissionsViewScopeArgs.builder()
 *                 .policies(testClientUserPolicy.id())
 *                 .description(&#34;description&#34;)
 *                 .decisionStrategy(&#34;UNANIMOUS&#34;)
 *                 .build())
 *             .manageScope(UsersPermissionsManageScopeArgs.builder()
 *                 .policies(testClientUserPolicy.id())
 *                 .description(&#34;description&#34;)
 *                 .decisionStrategy(&#34;UNANIMOUS&#34;)
 *                 .build())
 *             .mapRolesScope(UsersPermissionsMapRolesScopeArgs.builder()
 *                 .policies(testClientUserPolicy.id())
 *                 .description(&#34;description&#34;)
 *                 .decisionStrategy(&#34;UNANIMOUS&#34;)
 *                 .build())
 *             .manageGroupMembershipScope(UsersPermissionsManageGroupMembershipScopeArgs.builder()
 *                 .policies(testClientUserPolicy.id())
 *                 .description(&#34;description&#34;)
 *                 .decisionStrategy(&#34;UNANIMOUS&#34;)
 *                 .build())
 *             .impersonateScope(UsersPermissionsImpersonateScopeArgs.builder()
 *                 .policies(testClientUserPolicy.id())
 *                 .description(&#34;description&#34;)
 *                 .decisionStrategy(&#34;UNANIMOUS&#34;)
 *                 .build())
 *             .userImpersonatedScope(UsersPermissionsUserImpersonatedScopeArgs.builder()
 *                 .policies(testClientUserPolicy.id())
 *                 .description(&#34;description&#34;)
 *                 .decisionStrategy(&#34;UNANIMOUS&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realm_id` - (Required) The realm in which to manage fine-grained user permissions.
 * 
 * Each of the scopes that can be managed are defined below:
 * 
 * - `view_scope` - (Optional) When specified, set the scope based view permission.
 * - `manage_scope` - (Optional) When specified, set the scope based manage permission.
 * - `map_roles_scope` - (Optional) When specified, set the scope based map_roles permission.
 * - `manage_group_membership_scope` - (Optional) When specified, set the scope based manage_group_membership permission.
 * - `impersonate_scope` - (Optional) When specified, set the scope based impersonate permission.
 * - `user_impersonated_scope` - (Optional) When specified, set the scope based user_impersonated permission.
 * 
 * The configuration block for each of these scopes supports the following arguments:
 * 
 * - `policies` - (Optional) Assigned policies to the permission. Each element within this list should be a policy ID.
 * - `description` - (Optional) Description of the permission.
 * - `decision_strategy` - (Optional) Decision strategy of the permission.
 * 
 * ### Attributes Reference
 * 
 * In addition to the arguments listed above, the following computed attributes are exported:
 * 
 * - `enabled` - When true, this indicates that fine-grained user permissions are enabled. This will always be `true`.
 * - `authorization_resource_server_id` - Resource server id representing the realm management client on which these permissions are managed.
 * 
 */
@ResourceType(type="keycloak:index/usersPermissions:UsersPermissions")
public class UsersPermissions extends com.pulumi.resources.CustomResource {
    /**
     * Resource server id representing the realm management client on which this permission is managed
     * 
     */
    @Export(name="authorizationResourceServerId", refs={String.class}, tree="[0]")
    private Output<String> authorizationResourceServerId;

    /**
     * @return Resource server id representing the realm management client on which this permission is managed
     * 
     */
    public Output<String> authorizationResourceServerId() {
        return this.authorizationResourceServerId;
    }
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    public Output<Boolean> enabled() {
        return this.enabled;
    }
    @Export(name="impersonateScope", refs={UsersPermissionsImpersonateScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsImpersonateScope> impersonateScope;

    public Output<Optional<UsersPermissionsImpersonateScope>> impersonateScope() {
        return Codegen.optional(this.impersonateScope);
    }
    @Export(name="manageGroupMembershipScope", refs={UsersPermissionsManageGroupMembershipScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsManageGroupMembershipScope> manageGroupMembershipScope;

    public Output<Optional<UsersPermissionsManageGroupMembershipScope>> manageGroupMembershipScope() {
        return Codegen.optional(this.manageGroupMembershipScope);
    }
    @Export(name="manageScope", refs={UsersPermissionsManageScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsManageScope> manageScope;

    public Output<Optional<UsersPermissionsManageScope>> manageScope() {
        return Codegen.optional(this.manageScope);
    }
    @Export(name="mapRolesScope", refs={UsersPermissionsMapRolesScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsMapRolesScope> mapRolesScope;

    public Output<Optional<UsersPermissionsMapRolesScope>> mapRolesScope() {
        return Codegen.optional(this.mapRolesScope);
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="userImpersonatedScope", refs={UsersPermissionsUserImpersonatedScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsUserImpersonatedScope> userImpersonatedScope;

    public Output<Optional<UsersPermissionsUserImpersonatedScope>> userImpersonatedScope() {
        return Codegen.optional(this.userImpersonatedScope);
    }
    @Export(name="viewScope", refs={UsersPermissionsViewScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsViewScope> viewScope;

    public Output<Optional<UsersPermissionsViewScope>> viewScope() {
        return Codegen.optional(this.viewScope);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UsersPermissions(String name) {
        this(name, UsersPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UsersPermissions(String name, UsersPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UsersPermissions(String name, UsersPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/usersPermissions:UsersPermissions", name, args == null ? UsersPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UsersPermissions(String name, Output<String> id, @Nullable UsersPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/usersPermissions:UsersPermissions", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static UsersPermissions get(String name, Output<String> id, @Nullable UsersPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UsersPermissions(name, id, state, options);
    }
}
