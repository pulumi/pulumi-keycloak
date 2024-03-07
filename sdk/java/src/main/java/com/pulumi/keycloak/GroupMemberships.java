// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.GroupMembershipsArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.GroupMembershipsState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for managing a Keycloak group&#39;s members.
 * 
 * Note that this resource attempts to be an **authoritative** source over group members. When this resource takes control
 * over a group&#39;s members, users that are manually added to the group will be removed, and users that are manually removed
 * from the group will be added upon the next run of `pulumi up`.
 * 
 * Also note that you should not use `keycloak.GroupMemberships` with a group has been assigned as a default group via
 * `keycloak.DefaultGroups`.
 * 
 * This resource **should not** be used to control membership of a group that has its members federated from an external
 * source via group mapping.
 * 
 * To non-exclusively manage the group&#39;s of a user, see the [`keycloak.UserGroups` resource][1]
 * 
 * This resource paginates its data loading on refresh by 50 items.
 * 
 * ## Example Usage
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
 * import com.pulumi.keycloak.Group;
 * import com.pulumi.keycloak.GroupArgs;
 * import com.pulumi.keycloak.User;
 * import com.pulumi.keycloak.UserArgs;
 * import com.pulumi.keycloak.GroupMemberships;
 * import com.pulumi.keycloak.GroupMembershipsArgs;
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
 *             .build());
 * 
 *         var user = new User(&#34;user&#34;, UserArgs.builder()        
 *             .realmId(realm.id())
 *             .username(&#34;my-user&#34;)
 *             .build());
 * 
 *         var groupMembers = new GroupMemberships(&#34;groupMembers&#34;, GroupMembershipsArgs.builder()        
 *             .realmId(realm.id())
 *             .groupId(group.id())
 *             .members(user.username())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource does not support import. Instead of importing, feel free to create this resource
 * 
 * as if it did not already exist on the server.
 * 
 * [1]: providers/mrparkers/keycloak/latest/docs/resources/group_memberships
 * 
 */
@ResourceType(type="keycloak:index/groupMemberships:GroupMemberships")
public class GroupMemberships extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the group this resource should manage memberships for.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupId;

    /**
     * @return The ID of the group this resource should manage memberships for.
     * 
     */
    public Output<Optional<String>> groupId() {
        return Codegen.optional(this.groupId);
    }
    /**
     * A list of usernames that belong to this group.
     * 
     */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> members;

    /**
     * @return A list of usernames that belong to this group.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupMemberships(String name) {
        this(name, GroupMembershipsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupMemberships(String name, GroupMembershipsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupMemberships(String name, GroupMembershipsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/groupMemberships:GroupMemberships", name, args == null ? GroupMembershipsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupMemberships(String name, Output<String> id, @Nullable GroupMembershipsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/groupMemberships:GroupMemberships", name, state, makeResourceOptions(options, id));
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
    public static GroupMemberships get(String name, Output<String> id, @Nullable GroupMembershipsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupMemberships(name, id, state, options);
    }
}
