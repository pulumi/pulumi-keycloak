// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.RoleArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.RoleState;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing roles within Keycloak.
 * 
 * Roles allow you define privileges within Keycloak and map them to users and groups.
 * 
 * ## Example Usage
 * ### Realm Role)
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
 *             .description(&#34;My Realm Role&#34;)
 *             .attributes(Map.ofEntries(
 *                 Map.entry(&#34;key&#34;, &#34;value&#34;),
 *                 Map.entry(&#34;multivalue&#34;, &#34;value1##value2&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Client Role)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
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
 *         var openidClient = new Client(&#34;openidClient&#34;, ClientArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(&#34;client&#34;)
 *             .enabled(true)
 *             .accessType(&#34;CONFIDENTIAL&#34;)
 *             .validRedirectUris(&#34;http://localhost:8080/openid-callback&#34;)
 *             .build());
 * 
 *         var clientRole = new Role(&#34;clientRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(keycloak_client.openid_client().id())
 *             .description(&#34;My Client Role&#34;)
 *             .attributes(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Composite Role)
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
 *         var createRole = new Role(&#34;createRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .attributes(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .build());
 * 
 *         var readRole = new Role(&#34;readRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .attributes(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .build());
 * 
 *         var updateRole = new Role(&#34;updateRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .attributes(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .build());
 * 
 *         var deleteRole = new Role(&#34;deleteRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .attributes(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .build());
 * 
 *         var openidClient = new Client(&#34;openidClient&#34;, ClientArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(&#34;client&#34;)
 *             .enabled(true)
 *             .accessType(&#34;CONFIDENTIAL&#34;)
 *             .validRedirectUris(&#34;http://localhost:8080/openid-callback&#34;)
 *             .build());
 * 
 *         var clientRole = new Role(&#34;clientRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(keycloak_client.openid_client().id())
 *             .description(&#34;My Client Role&#34;)
 *             .attributes(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .build());
 * 
 *         var adminRole = new Role(&#34;adminRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .compositeRoles(            
 *                 createRole.id(),
 *                 readRole.id(),
 *                 updateRole.id(),
 *                 deleteRole.id(),
 *                 clientRole.id())
 *             .attributes(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Roles can be imported using the format `{{realm_id}}/{{role_id}}`, where `role_id` is the unique ID that Keycloak assigns to the role. The ID is not easy to find in the GUI, but it appears in the URL when editing the role. Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:index/role:Role role my-realm/7e8cf32a-8acb-4d34-89c4-04fb1d10ccad
 * ```
 * 
 */
@ResourceType(type="keycloak:index/role:Role")
public class Role extends com.pulumi.resources.CustomResource {
    /**
     * A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     * 
     */
    @Export(name="attributes", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> attributes;

    /**
     * @return A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     * 
     */
    public Output<Optional<Map<String,Object>>> attributes() {
        return Codegen.optional(this.attributes);
    }
    /**
     * When specified, this role will be created as a client role attached to the client with the provided ID
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return When specified, this role will be created as a client role attached to the client with the provided ID
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
     * 
     */
    @Export(name="compositeRoles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> compositeRoles;

    /**
     * @return When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
     * 
     */
    public Output<Optional<List<String>>> compositeRoles() {
        return Codegen.optional(this.compositeRoles);
    }
    /**
     * The description of the role
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the role
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of the role
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the role
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The realm this role exists within.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm this role exists within.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Role(String name) {
        this(name, RoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Role(String name, RoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Role(String name, RoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/role:Role", name, args == null ? RoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Role(String name, Output<String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/role:Role", name, state, makeResourceOptions(options, id));
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
    public static Role get(String name, Output<String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Role(name, id, state, options);
    }
}
