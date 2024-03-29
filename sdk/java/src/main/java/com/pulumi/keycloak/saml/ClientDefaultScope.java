// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.saml.ClientDefaultScopeArgs;
import com.pulumi.keycloak.saml.inputs.ClientDefaultScopeState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.keycloak.saml.Client;
 * import com.pulumi.keycloak.saml.ClientArgs;
 * import com.pulumi.keycloak.saml.ClientScope;
 * import com.pulumi.keycloak.saml.ClientScopeArgs;
 * import com.pulumi.keycloak.saml.ClientDefaultScope;
 * import com.pulumi.keycloak.saml.ClientDefaultScopeArgs;
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
 *         var samlClient = new Client(&#34;samlClient&#34;, ClientArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(&#34;saml-client&#34;)
 *             .signDocuments(false)
 *             .signAssertions(true)
 *             .includeAuthnStatement(true)
 *             .signingCertificate(Files.readString(Paths.get(&#34;saml-cert.pem&#34;)))
 *             .signingPrivateKey(Files.readString(Paths.get(&#34;saml-key.pem&#34;)))
 *             .build());
 * 
 *         var clientScope = new ClientScope(&#34;clientScope&#34;, ClientScopeArgs.builder()        
 *             .realmId(realm.id())
 *             .build());
 * 
 *         var clientDefaultScopes = new ClientDefaultScope(&#34;clientDefaultScopes&#34;, ClientDefaultScopeArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(keycloak_saml_client.client().id())
 *             .defaultScopes(            
 *                 &#34;role_list&#34;,
 *                 clientScope.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist
 * 
 * on the server.
 * 
 */
@ResourceType(type="keycloak:saml/clientDefaultScope:ClientDefaultScope")
public class ClientDefaultScope extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
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
    @Export(name="defaultScopes", refs={List.class,String.class}, tree="[0,1]")
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
    @Export(name="realmId", refs={String.class}, tree="[0]")
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
    public ClientDefaultScope(String name) {
        this(name, ClientDefaultScopeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClientDefaultScope(String name, ClientDefaultScopeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClientDefaultScope(String name, ClientDefaultScopeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/clientDefaultScope:ClientDefaultScope", name, args == null ? ClientDefaultScopeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClientDefaultScope(String name, Output<String> id, @Nullable ClientDefaultScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/clientDefaultScope:ClientDefaultScope", name, state, makeResourceOptions(options, id));
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
    public static ClientDefaultScope get(String name, Output<String> id, @Nullable ClientDefaultScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClientDefaultScope(name, id, state, options);
    }
}
