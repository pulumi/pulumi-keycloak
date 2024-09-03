// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.FullNameProtocolMapperArgs;
import com.pulumi.keycloak.openid.inputs.FullNameProtocolMapperState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * # keycloak.openid.FullNameProtocolMapper
 * 
 * Allows for creating and managing full name protocol mappers within
 * Keycloak.
 * 
 * Full name protocol mappers allow you to map a user&#39;s first and last name
 * to the OpenID Connect `name` claim in a token. Protocol mappers can be defined
 * for a single client, or they can be defined for a client scope which can
 * be shared between multiple different clients.
 * 
 * ### Example Usage (Client)
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
 * import com.pulumi.keycloak.openid.FullNameProtocolMapper;
 * import com.pulumi.keycloak.openid.FullNameProtocolMapperArgs;
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
 *         var openidClient = new Client("openidClient", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId("test-client")
 *             .name("test client")
 *             .enabled(true)
 *             .accessType("CONFIDENTIAL")
 *             .validRedirectUris("http://localhost:8080/openid-callback")
 *             .build());
 * 
 *         var fullNameMapper = new FullNameProtocolMapper("fullNameMapper", FullNameProtocolMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(openidClient.id())
 *             .name("full-name-mapper")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Example Usage (Client Scope)
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
 * import com.pulumi.keycloak.openid.FullNameProtocolMapper;
 * import com.pulumi.keycloak.openid.FullNameProtocolMapperArgs;
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
 *             .name("test-client-scope")
 *             .build());
 * 
 *         var fullNameMapper = new FullNameProtocolMapper("fullNameMapper", FullNameProtocolMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientScopeId(clientScope.id())
 *             .name("full-name-mapper")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
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
 * - `add_to_id_token` - (Optional) Indicates if the user&#39;s full name should be added as a claim to the id token. Defaults to `true`.
 * - `add_to_access_token` - (Optional) Indicates if the user&#39;s full name should be added as a claim to the access token. Defaults to `true`.
 * - `add_to_userinfo` - (Optional) Indicates if the user&#39;s full name should be added as a claim to the UserInfo response body. Defaults to `true`.
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
@ResourceType(type="keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper")
public class FullNameProtocolMapper extends com.pulumi.resources.CustomResource {
    @Export(name="addToAccessToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToAccessToken;

    public Output<Optional<Boolean>> addToAccessToken() {
        return Codegen.optional(this.addToAccessToken);
    }
    @Export(name="addToIdToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToIdToken;

    public Output<Optional<Boolean>> addToIdToken() {
        return Codegen.optional(this.addToIdToken);
    }
    @Export(name="addToUserinfo", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToUserinfo;

    public Output<Optional<Boolean>> addToUserinfo() {
        return Codegen.optional(this.addToUserinfo);
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
    public FullNameProtocolMapper(java.lang.String name) {
        this(name, FullNameProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FullNameProtocolMapper(java.lang.String name, FullNameProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FullNameProtocolMapper(java.lang.String name, FullNameProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private FullNameProtocolMapper(java.lang.String name, Output<java.lang.String> id, @Nullable FullNameProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper", name, state, makeResourceOptions(options, id), false);
    }

    private static FullNameProtocolMapperArgs makeArgs(FullNameProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FullNameProtocolMapperArgs.Empty : args;
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
    public static FullNameProtocolMapper get(java.lang.String name, Output<java.lang.String> id, @Nullable FullNameProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FullNameProtocolMapper(name, id, state, options);
    }
}
