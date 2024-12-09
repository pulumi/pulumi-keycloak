// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.ClientScopeArgs;
import com.pulumi.keycloak.openid.inputs.ClientScopeState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing Keycloak client scopes that can be attached to clients that use the OpenID Connect protocol.
 * 
 * Client Scopes can be used to share common protocol and role mappings between multiple clients within a realm. They can also
 * be used by clients to conditionally request claims or roles for a user based on the OAuth 2.0 `scope` parameter.
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
 * import com.pulumi.keycloak.openid.ClientScope;
 * import com.pulumi.keycloak.openid.ClientScopeArgs;
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
 *         var openidClientScope = new ClientScope("openidClientScope", ClientScopeArgs.builder()
 *             .realmId(realm.id())
 *             .name("groups")
 *             .description("When requested, this scope will map a user's group memberships to a claim")
 *             .includeInTokenScope(true)
 *             .guiOrder(1)
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
 * Client scopes can be imported using the format `{{realm_id}}/{{client_scope_id}}`, where `client_scope_id` is the unique ID that Keycloak
 * 
 * assigns to the client scope upon creation. This value can be found in the URI when editing this client scope in the GUI, and is typically a GUID.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:openid/clientScope:ClientScope openid_client_scope my-realm/8e8f7fe1-df9b-40ed-bed3-4597aa0dac52
 * ```
 * 
 */
@ResourceType(type="keycloak:openid/clientScope:ClientScope")
public class ClientScope extends com.pulumi.resources.CustomResource {
    /**
     * When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
     * 
     */
    @Export(name="consentScreenText", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> consentScreenText;

    /**
     * @return When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
     * 
     */
    public Output<Optional<String>> consentScreenText() {
        return Codegen.optional(this.consentScreenText);
    }
    /**
     * The description of this client scope in the GUI.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of this client scope in the GUI.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specify order of the client scope in GUI (such as in Consent page) as integer.
     * 
     */
    @Export(name="guiOrder", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> guiOrder;

    /**
     * @return Specify order of the client scope in GUI (such as in Consent page) as integer.
     * 
     */
    public Output<Optional<Integer>> guiOrder() {
        return Codegen.optional(this.guiOrder);
    }
    /**
     * When `true`, the name of this client scope will be added to the access token property &#39;scope&#39; as well as to the Token Introspection Endpoint response.
     * 
     */
    @Export(name="includeInTokenScope", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> includeInTokenScope;

    /**
     * @return When `true`, the name of this client scope will be added to the access token property &#39;scope&#39; as well as to the Token Introspection Endpoint response.
     * 
     */
    public Output<Optional<Boolean>> includeInTokenScope() {
        return Codegen.optional(this.includeInTokenScope);
    }
    /**
     * The display name of this client scope in the GUI.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The display name of this client scope in the GUI.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The realm this client scope belongs to.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm this client scope belongs to.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClientScope(java.lang.String name) {
        this(name, ClientScopeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClientScope(java.lang.String name, ClientScopeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClientScope(java.lang.String name, ClientScopeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientScope:ClientScope", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ClientScope(java.lang.String name, Output<java.lang.String> id, @Nullable ClientScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientScope:ClientScope", name, state, makeResourceOptions(options, id), false);
    }

    private static ClientScopeArgs makeArgs(ClientScopeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ClientScopeArgs.Empty : args;
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
    public static ClientScope get(java.lang.String name, Output<java.lang.String> id, @Nullable ClientScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClientScope(name, id, state, options);
    }
}
