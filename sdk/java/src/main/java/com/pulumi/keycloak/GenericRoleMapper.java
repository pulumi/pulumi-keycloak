// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.GenericRoleMapperArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.GenericRoleMapperState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allow for creating and managing a client&#39;s or client scope&#39;s role mappings within Keycloak.
 * 
 * By default, all the user role mappings of the user are added as claims within the token (OIDC) or assertion (SAML). When
 * `full_scope_allowed` is set to `false` for a client, role scope mapping allows you to limit the roles that get declared
 * inside an access token for a client.
 * 
 * ## Example Usage
 * 
 * ### Realm Role To Client)
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
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
 * import com.pulumi.keycloak.GenericRoleMapper;
 * import com.pulumi.keycloak.GenericRoleMapperArgs;
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
 *         var client = new Client("client", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId("client")
 *             .name("client")
 *             .enabled(true)
 *             .accessType("BEARER-ONLY")
 *             .build());
 * 
 *         var realmRole = new Role("realmRole", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .name("my-realm-role")
 *             .description("My Realm Role")
 *             .build());
 * 
 *         var clientRoleMapper = new GenericRoleMapper("clientRoleMapper", GenericRoleMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(client.id())
 *             .roleId(realmRole.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Client Role To Client)
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
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
 * import com.pulumi.keycloak.GenericRoleMapper;
 * import com.pulumi.keycloak.GenericRoleMapperArgs;
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
 *         var clientA = new Client("clientA", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId("client-a")
 *             .name("client-a")
 *             .enabled(true)
 *             .accessType("BEARER-ONLY")
 *             .fullScopeAllowed(false)
 *             .build());
 * 
 *         var clientRoleA = new Role("clientRoleA", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(clientA.id())
 *             .name("my-client-role")
 *             .description("My Client Role")
 *             .build());
 * 
 *         var clientB = new Client("clientB", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId("client-b")
 *             .name("client-b")
 *             .enabled(true)
 *             .accessType("BEARER-ONLY")
 *             .build());
 * 
 *         var clientRoleB = new Role("clientRoleB", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(clientB.id())
 *             .name("my-client-role")
 *             .description("My Client Role")
 *             .build());
 * 
 *         var clientBRoleMapper = new GenericRoleMapper("clientBRoleMapper", GenericRoleMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(clientB.id())
 *             .roleId(clientRoleA.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Realm Role To Client Scope)
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
 * import com.pulumi.keycloak.openid.ClientScope;
 * import com.pulumi.keycloak.openid.ClientScopeArgs;
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
 * import com.pulumi.keycloak.GenericRoleMapper;
 * import com.pulumi.keycloak.GenericRoleMapperArgs;
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
 *         var clientScope = new ClientScope("clientScope", ClientScopeArgs.builder()
 *             .realmId(realm.id())
 *             .name("my-client-scope")
 *             .build());
 * 
 *         var realmRole = new Role("realmRole", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .name("my-realm-role")
 *             .description("My Realm Role")
 *             .build());
 * 
 *         var clientRoleMapper = new GenericRoleMapper("clientRoleMapper", GenericRoleMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientScopeId(clientScope.id())
 *             .roleId(realmRole.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Client Role To Client Scope)
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
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
 * import com.pulumi.keycloak.openid.ClientScope;
 * import com.pulumi.keycloak.openid.ClientScopeArgs;
 * import com.pulumi.keycloak.GenericRoleMapper;
 * import com.pulumi.keycloak.GenericRoleMapperArgs;
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
 *         var client = new Client("client", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId("client")
 *             .name("client")
 *             .enabled(true)
 *             .accessType("BEARER-ONLY")
 *             .build());
 * 
 *         var clientRole = new Role("clientRole", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(client.id())
 *             .name("my-client-role")
 *             .description("My Client Role")
 *             .build());
 * 
 *         var clientScope = new ClientScope("clientScope", ClientScopeArgs.builder()
 *             .realmId(realm.id())
 *             .name("my-client-scope")
 *             .build());
 * 
 *         var clientBRoleMapper = new GenericRoleMapper("clientBRoleMapper", GenericRoleMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientScopeId(clientScope.id())
 *             .roleId(clientRole.id())
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
 * Generic client role mappers can be imported using one of the following two formats:
 * 
 * - When mapping a role to a client, use the format `{{realmId}}/client/{{clientId}}/scope-mappings/{{roleClientId}}/{{roleId}}`
 * 
 * - When mapping a role to a client scope, use the format `{{realmId}}/client-scope/{{clientScopeId}}/scope-mappings/{{roleClientId}}/{{roleId}}`
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:index/genericRoleMapper:GenericRoleMapper client_role_mapper my-realm/client/23888550-5dcd-41f6-85ba-554233021e9c/scope-mappings/ce51f004-bdfb-4dd5-a963-c4487d2dec5b/ff3aa49f-bc07-4030-8783-41918c3614a3
 * ```
 * 
 */
@ResourceType(type="keycloak:index/genericRoleMapper:GenericRoleMapper")
public class GenericRoleMapper extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
     * 
     */
    @Export(name="clientScopeId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientScopeId;

    /**
     * @return The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
     * 
     */
    public Output<Optional<String>> clientScopeId() {
        return Codegen.optional(this.clientScopeId);
    }
    /**
     * The realm this role mapper exists within.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm this role mapper exists within.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }
    /**
     * The ID of the role to be added to this role mapper.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The ID of the role to be added to this role mapper.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GenericRoleMapper(String name) {
        this(name, GenericRoleMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GenericRoleMapper(String name, GenericRoleMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GenericRoleMapper(String name, GenericRoleMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/genericRoleMapper:GenericRoleMapper", name, args == null ? GenericRoleMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GenericRoleMapper(String name, Output<String> id, @Nullable GenericRoleMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/genericRoleMapper:GenericRoleMapper", name, state, makeResourceOptions(options, id));
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
    public static GenericRoleMapper get(String name, Output<String> id, @Nullable GenericRoleMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GenericRoleMapper(name, id, state, options);
    }
}
