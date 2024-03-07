// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.ScriptProtocolMapperArgs;
import com.pulumi.keycloak.openid.inputs.ScriptProtocolMapperState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing script protocol mappers within Keycloak.
 * 
 * Script protocol mappers evaluate a JavaScript function to produce a token claim based on context information.
 * 
 * Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
 * multiple different clients.
 * 
 * &gt; Support for this protocol mapper was removed in Keycloak 18.
 * 
 * ## Example Usage
 * 
 * ### Client)
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
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
 * import com.pulumi.keycloak.openid.ScriptProtocolMapper;
 * import com.pulumi.keycloak.openid.ScriptProtocolMapperArgs;
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
 *         var scriptMapper = new ScriptProtocolMapper(&#34;scriptMapper&#34;, ScriptProtocolMapperArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(openidClient.id())
 *             .claimName(&#34;foo&#34;)
 *             .script(&#34;exports = &#39;foo&#39;;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Client Scope)
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
 * import com.pulumi.keycloak.openid.ClientScope;
 * import com.pulumi.keycloak.openid.ClientScopeArgs;
 * import com.pulumi.keycloak.openid.ScriptProtocolMapper;
 * import com.pulumi.keycloak.openid.ScriptProtocolMapperArgs;
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
 *         var clientScope = new ClientScope(&#34;clientScope&#34;, ClientScopeArgs.builder()        
 *             .realmId(realm.id())
 *             .build());
 * 
 *         var scriptMapper = new ScriptProtocolMapper(&#34;scriptMapper&#34;, ScriptProtocolMapperArgs.builder()        
 *             .realmId(realm.id())
 *             .clientScopeId(clientScope.id())
 *             .claimName(&#34;foo&#34;)
 *             .script(&#34;exports = &#39;foo&#39;;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Protocol mappers can be imported using one of the following formats:
 * 
 * - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
 * 
 * - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper script_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 * 
 * ```sh
 * $ pulumi import keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper script_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 * 
 */
@ResourceType(type="keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper")
public class ScriptProtocolMapper extends com.pulumi.resources.CustomResource {
    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     * 
     */
    @Export(name="addToAccessToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToAccessToken;

    /**
     * @return Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> addToAccessToken() {
        return Codegen.optional(this.addToAccessToken);
    }
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     * 
     */
    @Export(name="addToIdToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToIdToken;

    /**
     * @return Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> addToIdToken() {
        return Codegen.optional(this.addToIdToken);
    }
    /**
     * Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     * 
     */
    @Export(name="addToUserinfo", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToUserinfo;

    /**
     * @return Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> addToUserinfo() {
        return Codegen.optional(this.addToUserinfo);
    }
    /**
     * The name of the claim to insert into a token.
     * 
     */
    @Export(name="claimName", refs={String.class}, tree="[0]")
    private Output<String> claimName;

    /**
     * @return The name of the claim to insert into a token.
     * 
     */
    public Output<String> claimName() {
        return this.claimName;
    }
    /**
     * The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     * 
     */
    @Export(name="claimValueType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> claimValueType;

    /**
     * @return The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     * 
     */
    public Output<Optional<String>> claimValueType() {
        return Codegen.optional(this.claimValueType);
    }
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
     * Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
     * 
     */
    @Export(name="multivalued", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> multivalued;

    /**
     * @return Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> multivalued() {
        return Codegen.optional(this.multivalued);
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
     * JavaScript code to compute the claim value.
     * 
     */
    @Export(name="script", refs={String.class}, tree="[0]")
    private Output<String> script;

    /**
     * @return JavaScript code to compute the claim value.
     * 
     */
    public Output<String> script() {
        return this.script;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ScriptProtocolMapper(String name) {
        this(name, ScriptProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ScriptProtocolMapper(String name, ScriptProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ScriptProtocolMapper(String name, ScriptProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper", name, args == null ? ScriptProtocolMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ScriptProtocolMapper(String name, Output<String> id, @Nullable ScriptProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper", name, state, makeResourceOptions(options, id));
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
    public static ScriptProtocolMapper get(String name, Output<String> id, @Nullable ScriptProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ScriptProtocolMapper(name, id, state, options);
    }
}
