// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.DefaultGroupsArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.DefaultGroupsState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Allows for managing a realm&#39;s default groups.
 * 
 * &gt; You should not use `keycloak.DefaultGroups` with a group whose members are managed by `keycloak.GroupMemberships`.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.Group;
 * import com.pulumi.keycloak.GroupArgs;
 * import com.pulumi.keycloak.DefaultGroups;
 * import com.pulumi.keycloak.DefaultGroupsArgs;
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
 *         var realm = new Realm("realm", RealmArgs.builder()
 *             .realm("my-realm")
 *             .enabled(true)
 *             .build());
 * 
 *         var group = new Group("group", GroupArgs.builder()
 *             .realmId(realm.id())
 *             .name("my-group")
 *             .build());
 * 
 *         var default_ = new DefaultGroups("default", DefaultGroupsArgs.builder()
 *             .realmId(realm.id())
 *             .groupIds(group.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Default groups can be imported using the format `{{realm_id}}` where `realm_id` is the realm the group exists in.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:index/defaultGroups:DefaultGroups default my-realm
 * ```
 * 
 */
@ResourceType(type="keycloak:index/defaultGroups:DefaultGroups")
public class DefaultGroups extends com.pulumi.resources.CustomResource {
    /**
     * A set of group ids that should be default groups on the realm referenced by `realm_id`.
     * 
     */
    @Export(name="groupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> groupIds;

    /**
     * @return A set of group ids that should be default groups on the realm referenced by `realm_id`.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DefaultGroups(java.lang.String name) {
        this(name, DefaultGroupsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DefaultGroups(java.lang.String name, DefaultGroupsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DefaultGroups(java.lang.String name, DefaultGroupsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/defaultGroups:DefaultGroups", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DefaultGroups(java.lang.String name, Output<java.lang.String> id, @Nullable DefaultGroupsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/defaultGroups:DefaultGroups", name, state, makeResourceOptions(options, id), false);
    }

    private static DefaultGroupsArgs makeArgs(DefaultGroupsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DefaultGroupsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static DefaultGroups get(java.lang.String name, Output<java.lang.String> id, @Nullable DefaultGroupsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DefaultGroups(name, id, state, options);
    }
}
