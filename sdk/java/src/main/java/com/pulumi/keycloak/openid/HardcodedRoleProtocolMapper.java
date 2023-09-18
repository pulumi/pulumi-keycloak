// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.HardcodedRoleProtocolMapperArgs;
import com.pulumi.keycloak.openid.inputs.HardcodedRoleProtocolMapperState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing hardcoded role protocol mappers within Keycloak.
 * 
 * Hardcoded role protocol mappers allow you to specify a single role to always map to an access token for a client.
 * 
 * Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
 * multiple different clients.
 * 
 * ## Example Usage
 * ### Client)
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
 * import com.pulumi.keycloak.openid.HardcodedRoleProtocolMapper;
 * import com.pulumi.keycloak.openid.HardcodedRoleProtocolMapperArgs;
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
 *         var role = new Role(&#34;role&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
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
 *         var hardcodedRoleMapper = new HardcodedRoleProtocolMapper(&#34;hardcodedRoleMapper&#34;, HardcodedRoleProtocolMapperArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(openidClient.id())
 *             .roleId(role.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Client Scope)
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
 * import com.pulumi.keycloak.openid.ClientScope;
 * import com.pulumi.keycloak.openid.ClientScopeArgs;
 * import com.pulumi.keycloak.openid.HardcodedRoleProtocolMapper;
 * import com.pulumi.keycloak.openid.HardcodedRoleProtocolMapperArgs;
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
 *         var role = new Role(&#34;role&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .build());
 * 
 *         var clientScope = new ClientScope(&#34;clientScope&#34;, ClientScopeArgs.builder()        
 *             .realmId(realm.id())
 *             .build());
 * 
 *         var hardcodedRoleMapper = new HardcodedRoleProtocolMapper(&#34;hardcodedRoleMapper&#34;, HardcodedRoleProtocolMapperArgs.builder()        
 *             .realmId(realm.id())
 *             .clientScopeId(clientScope.id())
 *             .roleId(role.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper hardcoded_role_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 * 
 * ```sh
 *  $ pulumi import keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper hardcoded_role_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 * 
 */
@ResourceType(type="keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper")
public class HardcodedRoleProtocolMapper extends com.pulumi.resources.CustomResource {
    /**
     * The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    @Export(name="clientScopeId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientScopeId;

    /**
     * @return The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    public Output<Optional<String>> clientScopeId() {
        return Codegen.optional(this.clientScopeId);
    }
    /**
     * The display name of this protocol mapper in the GUI.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The display name of this protocol mapper in the GUI.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The realm this protocol mapper exists within.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm this protocol mapper exists within.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }
    /**
     * The ID of the role to map to an access token.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The ID of the role to map to an access token.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HardcodedRoleProtocolMapper(String name) {
        this(name, HardcodedRoleProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HardcodedRoleProtocolMapper(String name, HardcodedRoleProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HardcodedRoleProtocolMapper(String name, HardcodedRoleProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, args == null ? HardcodedRoleProtocolMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HardcodedRoleProtocolMapper(String name, Output<String> id, @Nullable HardcodedRoleProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, state, makeResourceOptions(options, id));
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
    public static HardcodedRoleProtocolMapper get(String name, Output<String> id, @Nullable HardcodedRoleProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HardcodedRoleProtocolMapper(name, id, state, options);
    }
}
