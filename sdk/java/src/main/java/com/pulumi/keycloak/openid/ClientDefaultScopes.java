// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.ClientDefaultScopesArgs;
import com.pulumi.keycloak.openid.inputs.ClientDefaultScopesState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ## Example Usage
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
 * import com.pulumi.keycloak.openid.ClientScope;
 * import com.pulumi.keycloak.openid.ClientScopeArgs;
 * import com.pulumi.keycloak.openid.ClientDefaultScopes;
 * import com.pulumi.keycloak.openid.ClientDefaultScopesArgs;
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
 *         var client = new Client(&#34;client&#34;, ClientArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(&#34;test-client&#34;)
 *             .accessType(&#34;CONFIDENTIAL&#34;)
 *             .build());
 * 
 *         var clientScope = new ClientScope(&#34;clientScope&#34;, ClientScopeArgs.builder()        
 *             .realmId(realm.id())
 *             .build());
 * 
 *         var clientDefaultScopes = new ClientDefaultScopes(&#34;clientDefaultScopes&#34;, ClientDefaultScopesArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(client.id())
 *             .defaultScopes(            
 *                 &#34;profile&#34;,
 *                 &#34;email&#34;,
 *                 &#34;roles&#34;,
 *                 &#34;web-origins&#34;,
 *                 clientScope.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server.
 * 
 */
@ResourceType(type="keycloak:openid/clientDefaultScopes:ClientDefaultScopes")
public class ClientDefaultScopes extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
     * 
     */
    @Export(name="clientId", type=String.class, parameters={})
    private Output<String> clientId;

    /**
     * @return The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * An array of client scope names to attach to this client.
     * 
     */
    @Export(name="defaultScopes", type=List.class, parameters={String.class})
    private Output<List<String>> defaultScopes;

    /**
     * @return An array of client scope names to attach to this client.
     * 
     */
    public Output<List<String>> defaultScopes() {
        return this.defaultScopes;
    }
    /**
     * The realm this client and scopes exists in.
     * 
     */
    @Export(name="realmId", type=String.class, parameters={})
    private Output<String> realmId;

    /**
     * @return The realm this client and scopes exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClientDefaultScopes(String name) {
        this(name, ClientDefaultScopesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClientDefaultScopes(String name, ClientDefaultScopesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClientDefaultScopes(String name, ClientDefaultScopesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientDefaultScopes:ClientDefaultScopes", name, args == null ? ClientDefaultScopesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClientDefaultScopes(String name, Output<String> id, @Nullable ClientDefaultScopesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientDefaultScopes:ClientDefaultScopes", name, state, makeResourceOptions(options, id));
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
    public static ClientDefaultScopes get(String name, Output<String> id, @Nullable ClientDefaultScopesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClientDefaultScopes(name, id, state, options);
    }
}