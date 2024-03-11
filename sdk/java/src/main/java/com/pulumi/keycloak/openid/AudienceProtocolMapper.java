// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.AudienceProtocolMapperArgs;
import com.pulumi.keycloak.openid.inputs.AudienceProtocolMapperState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.openid.AudienceProtocolMapper
 * 
 * Allows for creating and managing audience protocol mappers within
 * Keycloak. This mapper was added in Keycloak v4.6.0.Final.
 * 
 * Audience protocol mappers allow you add audiences to the `aud` claim
 * within issued tokens. The audience can be a custom string, or it can be
 * mapped to the ID of a pre-existing client.
 * 
 * ### Example Usage (Client)
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
 * import com.pulumi.keycloak.openid.AudienceProtocolMapper;
 * import com.pulumi.keycloak.openid.AudienceProtocolMapperArgs;
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
 *             .enabled(true)
 *             .realm(&#34;my-realm&#34;)
 *             .build());
 * 
 *         var openidClient = new Client(&#34;openidClient&#34;, ClientArgs.builder()        
 *             .accessType(&#34;CONFIDENTIAL&#34;)
 *             .clientId(&#34;test-client&#34;)
 *             .enabled(true)
 *             .realmId(realm.id())
 *             .validRedirectUris(&#34;http://localhost:8080/openid-callback&#34;)
 *             .build());
 * 
 *         var audienceMapper = new AudienceProtocolMapper(&#34;audienceMapper&#34;, AudienceProtocolMapperArgs.builder()        
 *             .clientId(openidClient.id())
 *             .includedCustomAudience(&#34;foo&#34;)
 *             .realmId(realm.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Example Usage (Client Scope)
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
 * import com.pulumi.keycloak.openid.AudienceProtocolMapper;
 * import com.pulumi.keycloak.openid.AudienceProtocolMapperArgs;
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
 *             .enabled(true)
 *             .realm(&#34;my-realm&#34;)
 *             .build());
 * 
 *         var clientScope = new ClientScope(&#34;clientScope&#34;, ClientScopeArgs.builder()        
 *             .realmId(realm.id())
 *             .build());
 * 
 *         var audienceMapper = new AudienceProtocolMapper(&#34;audienceMapper&#34;, AudienceProtocolMapperArgs.builder()        
 *             .clientScopeId(clientScope.id())
 *             .includedCustomAudience(&#34;foo&#34;)
 *             .realmId(realm.id())
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
 * - `realm_id` - (Required) The realm this protocol mapper exists within.
 * - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
 * - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the GUI.
 * - `included_client_audience` - (Required if `included_custom_audience` is not specified) A client ID to include within the token&#39;s `aud` claim.
 * - `included_custom_audience` - (Required if `included_client_audience` is not specified) A custom audience to include within the token&#39;s `aud` claim.
 * - `add_to_id_token` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
 * - `add_to_access_token` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
 * 
 * ### Import
 * 
 * Protocol mappers can be imported using one of the following formats:
 * - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
 * - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
 * 
 * Example:
 * 
 */
@ResourceType(type="keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper")
public class AudienceProtocolMapper extends com.pulumi.resources.CustomResource {
    /**
     * Indicates if this claim should be added to the access token.
     * 
     */
    @Export(name="addToAccessToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToAccessToken;

    /**
     * @return Indicates if this claim should be added to the access token.
     * 
     */
    public Output<Optional<Boolean>> addToAccessToken() {
        return Codegen.optional(this.addToAccessToken);
    }
    /**
     * Indicates if this claim should be added to the id token.
     * 
     */
    @Export(name="addToIdToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToIdToken;

    /**
     * @return Indicates if this claim should be added to the id token.
     * 
     */
    public Output<Optional<Boolean>> addToIdToken() {
        return Codegen.optional(this.addToIdToken);
    }
    /**
     * The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
     * 
     */
    @Export(name="clientScopeId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientScopeId;

    /**
     * @return The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
     * 
     */
    public Output<Optional<String>> clientScopeId() {
        return Codegen.optional(this.clientScopeId);
    }
    /**
     * A client ID to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
     * 
     */
    @Export(name="includedClientAudience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> includedClientAudience;

    /**
     * @return A client ID to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
     * 
     */
    public Output<Optional<String>> includedClientAudience() {
        return Codegen.optional(this.includedClientAudience);
    }
    /**
     * A custom audience to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
     * 
     */
    @Export(name="includedCustomAudience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> includedCustomAudience;

    /**
     * @return A custom audience to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
     * 
     */
    public Output<Optional<String>> includedCustomAudience() {
        return Codegen.optional(this.includedCustomAudience);
    }
    /**
     * A human-friendly name that will appear in the Keycloak console.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A human-friendly name that will appear in the Keycloak console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The realm id where the associated client or client scope exists.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm id where the associated client or client scope exists.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AudienceProtocolMapper(String name) {
        this(name, AudienceProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AudienceProtocolMapper(String name, AudienceProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AudienceProtocolMapper(String name, AudienceProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper", name, args == null ? AudienceProtocolMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AudienceProtocolMapper(String name, Output<String> id, @Nullable AudienceProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper", name, state, makeResourceOptions(options, id));
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
    public static AudienceProtocolMapper get(String name, Output<String> id, @Nullable AudienceProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AudienceProtocolMapper(name, id, state, options);
    }
}
