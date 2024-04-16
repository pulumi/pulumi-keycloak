// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.GroupRolesArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.GroupRolesState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.GroupRoles
 * 
 * Allows you to manage roles assigned to a Keycloak group.
 * 
 * Note that this resource attempts to be an **authoritative** source over
 * group roles. When this resource takes control over a group&#39;s roles,
 * roles that are manually added to the group will be removed, and roles
 * that are manually removed from the group will be added upon the next run
 * of `pulumi up`.
 * 
 * Note that when assigning composite roles to a group, you may see a
 * non-empty plan following a `pulumi up` if you assign a role and a
 * composite that includes that role to the same group.
 * 
 * ### Example Usage
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
 * import com.pulumi.keycloak.Group;
 * import com.pulumi.keycloak.GroupArgs;
 * import com.pulumi.keycloak.GroupRoles;
 * import com.pulumi.keycloak.GroupRolesArgs;
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
 *         var group = new Group(&#34;group&#34;, GroupArgs.builder()        
 *             .realmId(realm.id())
 *             .name(&#34;my-group&#34;)
 *             .build());
 * 
 *         var groupRoles = new GroupRoles(&#34;groupRoles&#34;, GroupRolesArgs.builder()        
 *             .realmId(realm.id())
 *             .groupId(group.id())
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
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realm_id` - (Required) The realm this group exists in.
 * - `group_id` - (Required) The ID of the group this resource should
 *   manage roles for.
 * - `role_ids` - (Required) A list of role IDs to map to the group
 * 
 * ### Import
 * 
 * This resource can be imported using the format
 * `{{realm_id}}/{{group_id}}`, where `group_id` is the unique ID that
 * Keycloak assigns to the group upon creation. This value can be found in
 * the URI when editing this group in the GUI, and is typically a GUID.
 * 
 * Example:
 * 
 */
@ResourceType(type="keycloak:index/groupRoles:GroupRoles")
public class GroupRoles extends com.pulumi.resources.CustomResource {
    @Export(name="exhaustive", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exhaustive;

    public Output<Optional<Boolean>> exhaustive() {
        return Codegen.optional(this.exhaustive);
    }
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    public Output<String> groupId() {
        return this.groupId;
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="roleIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> roleIds;

    public Output<List<String>> roleIds() {
        return this.roleIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupRoles(String name) {
        this(name, GroupRolesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupRoles(String name, GroupRolesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupRoles(String name, GroupRolesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/groupRoles:GroupRoles", name, args == null ? GroupRolesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupRoles(String name, Output<String> id, @Nullable GroupRolesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/groupRoles:GroupRoles", name, state, makeResourceOptions(options, id));
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
    public static GroupRoles get(String name, Output<String> id, @Nullable GroupRolesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupRoles(name, id, state, options);
    }
}
