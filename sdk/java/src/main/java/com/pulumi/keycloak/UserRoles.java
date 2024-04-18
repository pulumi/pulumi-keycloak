// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.UserRolesArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.UserRolesState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows you to manage roles assigned to a Keycloak user.
 * 
 * If `exhaustive` is true, this resource attempts to be an **authoritative** source over user roles: roles that are manually added to the user will be removed, and roles that are manually removed from the
 * user will be added upon the next run of `pulumi up`.
 * If `exhaustive` is false, this resource is a partial assignation of roles to a user. As a result, you can use multiple `keycloak.UserRoles` for the same `user_id`.
 * 
 * Note that when assigning composite roles to a user, you may see a non-empty plan following a `pulumi up` if you assign
 * a role and a composite that includes that role to the same user.
 * 
 * ## Example Usage
 * 
 * ### Exhaustive Roles)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
 * import com.pulumi.keycloak.User;
 * import com.pulumi.keycloak.UserArgs;
 * import com.pulumi.keycloak.UserRoles;
 * import com.pulumi.keycloak.UserRolesArgs;
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
 *             .enabled(true)
 *             .build());
 * 
 *         var realmRole = new Role(&#34;realmRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .name(&#34;my-realm-role&#34;)
 *             .description(&#34;My Realm Role&#34;)
 *             .build());
 * 
 *         var client = new Client(&#34;client&#34;, ClientArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(&#34;client&#34;)
 *             .name(&#34;client&#34;)
 *             .enabled(true)
 *             .accessType(&#34;BEARER-ONLY&#34;)
 *             .build());
 * 
 *         var clientRole = new Role(&#34;clientRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(clientKeycloakClient.id())
 *             .name(&#34;my-client-role&#34;)
 *             .description(&#34;My Client Role&#34;)
 *             .build());
 * 
 *         var user = new User(&#34;user&#34;, UserArgs.builder()        
 *             .realmId(realm.id())
 *             .username(&#34;bob&#34;)
 *             .enabled(true)
 *             .email(&#34;bob@domain.com&#34;)
 *             .firstName(&#34;Bob&#34;)
 *             .lastName(&#34;Bobson&#34;)
 *             .build());
 * 
 *         var userRoles = new UserRoles(&#34;userRoles&#34;, UserRolesArgs.builder()        
 *             .realmId(realm.id())
 *             .userId(user.id())
 *             .roleIds(            
 *                 realmRole.id(),
 *                 clientRole.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak
 * 
 * assigns to the user upon creation. This value can be found in the GUI when editing the user, and is typically a GUID.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:index/userRoles:UserRoles user_roles my-realm/b0ae6924-1bd5-4655-9e38-dae7c5e42924
 * ```
 * 
 */
@ResourceType(type="keycloak:index/userRoles:UserRoles")
public class UserRoles extends com.pulumi.resources.CustomResource {
    /**
     * Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
     * 
     */
    @Export(name="exhaustive", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exhaustive;

    /**
     * @return Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> exhaustive() {
        return Codegen.optional(this.exhaustive);
    }
    /**
     * The realm this user exists in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm this user exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }
    /**
     * A list of role IDs to map to the user
     * 
     */
    @Export(name="roleIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> roleIds;

    /**
     * @return A list of role IDs to map to the user
     * 
     */
    public Output<List<String>> roleIds() {
        return this.roleIds;
    }
    /**
     * The ID of the user this resource should manage roles for.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return The ID of the user this resource should manage roles for.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserRoles(String name) {
        this(name, UserRolesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserRoles(String name, UserRolesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserRoles(String name, UserRolesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/userRoles:UserRoles", name, args == null ? UserRolesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserRoles(String name, Output<String> id, @Nullable UserRolesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/userRoles:UserRoles", name, state, makeResourceOptions(options, id));
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
    public static UserRoles get(String name, Output<String> id, @Nullable UserRolesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserRoles(name, id, state, options);
    }
}
