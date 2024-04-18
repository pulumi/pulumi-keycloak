// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.UserGroupsArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.UserGroupsState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for managing a Keycloak user&#39;s groups.
 * 
 * If `exhaustive` is true, this resource attempts to be an **authoritative** source over user groups: groups that are manually added to the user will be removed, and groups that are manually removed from the user group will be added upon the next run of `pulumi up`.
 * If `exhaustive` is false, this resource is a partial assignation of groups to a user. As a result, you can get multiple `keycloak.UserGroups` for the same `user_id`.
 * 
 * ## Example Usage
 * 
 * ### Exhaustive Groups)
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.Group;
 * import com.pulumi.keycloak.GroupArgs;
 * import com.pulumi.keycloak.User;
 * import com.pulumi.keycloak.UserArgs;
 * import com.pulumi.keycloak.UserGroups;
 * import com.pulumi.keycloak.UserGroupsArgs;
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
 *         var group = new Group(&#34;group&#34;, GroupArgs.builder()        
 *             .realmId(realm.id())
 *             .name(&#34;foo&#34;)
 *             .build());
 * 
 *         var user = new User(&#34;user&#34;, UserArgs.builder()        
 *             .realmId(realm.id())
 *             .username(&#34;my-user&#34;)
 *             .build());
 * 
 *         var userGroups = new UserGroups(&#34;userGroups&#34;, UserGroupsArgs.builder()        
 *             .realmId(realm.id())
 *             .userId(user.id())
 *             .groupIds(group.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="keycloak:index/userGroups:UserGroups")
public class UserGroups extends com.pulumi.resources.CustomResource {
    /**
     * Indicates if the list of the user&#39;s groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
     * 
     */
    @Export(name="exhaustive", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exhaustive;

    /**
     * @return Indicates if the list of the user&#39;s groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> exhaustive() {
        return Codegen.optional(this.exhaustive);
    }
    /**
     * A list of group IDs that the user is member of.
     * 
     */
    @Export(name="groupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> groupIds;

    /**
     * @return A list of group IDs that the user is member of.
     * 
     */
    public Output<List<String>> groupIds() {
        return this.groupIds;
    }
    /**
     * The realm this group exists in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm this group exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }
    /**
     * The ID of the user this resource should manage groups for.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return The ID of the user this resource should manage groups for.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserGroups(String name) {
        this(name, UserGroupsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserGroups(String name, UserGroupsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserGroups(String name, UserGroupsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/userGroups:UserGroups", name, args == null ? UserGroupsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserGroups(String name, Output<String> id, @Nullable UserGroupsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/userGroups:UserGroups", name, state, makeResourceOptions(options, id));
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
    public static UserGroups get(String name, Output<String> id, @Nullable UserGroupsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserGroups(name, id, state, options);
    }
}
