// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.UserAttributeProtocolMapperArgs;
import com.pulumi.keycloak.openid.inputs.UserAttributeProtocolMapperState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.openid.UserAttributeProtocolMapper
 * 
 * Allows for creating and managing user attribute protocol mappers within
 * Keycloak.
 * 
 * User attribute protocol mappers allow you to map custom attributes defined
 * for a user within Keycloak to a claim in a token. Protocol mappers can be
 * defined for a single client, or they can be defined for a client scope which
 * can be shared between multiple different clients.
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
 * import com.pulumi.keycloak.openid.UserAttributeProtocolMapper;
 * import com.pulumi.keycloak.openid.UserAttributeProtocolMapperArgs;
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
 *         var userAttributeMapper = new UserAttributeProtocolMapper("userAttributeMapper", UserAttributeProtocolMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(openidClient.id())
 *             .name("test-mapper")
 *             .userAttribute("foo")
 *             .claimName("bar")
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
 * import com.pulumi.keycloak.openid.UserAttributeProtocolMapper;
 * import com.pulumi.keycloak.openid.UserAttributeProtocolMapperArgs;
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
 *         var userAttributeMapper = new UserAttributeProtocolMapper("userAttributeMapper", UserAttributeProtocolMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientScopeId(clientScope.id())
 *             .name("test-mapper")
 *             .userAttribute("foo")
 *             .claimName("bar")
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
 * - `user_attribute` - (Required) The custom user attribute to map a claim for.
 * - `claim_name` - (Required) The name of the claim to insert into a token.
 * - `claim_value_type` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
 * - `multivalued` - (Optional) Indicates whether this attribute is a single value or an array of values. Defaults to `false`.
 * - `add_to_id_token` - (Optional) Indicates if the attribute should be added as a claim to the id token. Defaults to `true`.
 * - `add_to_access_token` - (Optional) Indicates if the attribute should be added as a claim to the access token. Defaults to `true`.
 * - `add_to_userinfo` - (Optional) Indicates if the attribute should be added as a claim to the UserInfo response body. Defaults to `true`.
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
@ResourceType(type="keycloak:openid/userAttributeProtocolMapper:UserAttributeProtocolMapper")
public class UserAttributeProtocolMapper extends com.pulumi.resources.CustomResource {
    /**
     * Indicates if the attribute should be a claim in the access token.
     * 
     */
    @Export(name="addToAccessToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToAccessToken;

    /**
     * @return Indicates if the attribute should be a claim in the access token.
     * 
     */
    public Output<Optional<Boolean>> addToAccessToken() {
        return Codegen.optional(this.addToAccessToken);
    }
    /**
     * Indicates if the attribute should be a claim in the id token.
     * 
     */
    @Export(name="addToIdToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToIdToken;

    /**
     * @return Indicates if the attribute should be a claim in the id token.
     * 
     */
    public Output<Optional<Boolean>> addToIdToken() {
        return Codegen.optional(this.addToIdToken);
    }
    /**
     * Indicates if the attribute should appear in the userinfo response body.
     * 
     */
    @Export(name="addToUserinfo", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addToUserinfo;

    /**
     * @return Indicates if the attribute should appear in the userinfo response body.
     * 
     */
    public Output<Optional<Boolean>> addToUserinfo() {
        return Codegen.optional(this.addToUserinfo);
    }
    /**
     * Indicates if attribute values should be aggregated within the group attributes
     * 
     */
    @Export(name="aggregateAttributes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> aggregateAttributes;

    /**
     * @return Indicates if attribute values should be aggregated within the group attributes
     * 
     */
    public Output<Optional<Boolean>> aggregateAttributes() {
        return Codegen.optional(this.aggregateAttributes);
    }
    @Export(name="claimName", refs={String.class}, tree="[0]")
    private Output<String> claimName;

    public Output<String> claimName() {
        return this.claimName;
    }
    /**
     * Claim type used when serializing tokens.
     * 
     */
    @Export(name="claimValueType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> claimValueType;

    /**
     * @return Claim type used when serializing tokens.
     * 
     */
    public Output<Optional<String>> claimValueType() {
        return Codegen.optional(this.claimValueType);
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
     * Indicates whether this attribute is a single value or an array of values.
     * 
     */
    @Export(name="multivalued", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> multivalued;

    /**
     * @return Indicates whether this attribute is a single value or an array of values.
     * 
     */
    public Output<Optional<Boolean>> multivalued() {
        return Codegen.optional(this.multivalued);
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
    @Export(name="userAttribute", refs={String.class}, tree="[0]")
    private Output<String> userAttribute;

    public Output<String> userAttribute() {
        return this.userAttribute;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserAttributeProtocolMapper(java.lang.String name) {
        this(name, UserAttributeProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserAttributeProtocolMapper(java.lang.String name, UserAttributeProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserAttributeProtocolMapper(java.lang.String name, UserAttributeProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private UserAttributeProtocolMapper(java.lang.String name, Output<java.lang.String> id, @Nullable UserAttributeProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, state, makeResourceOptions(options, id), false);
    }

    private static UserAttributeProtocolMapperArgs makeArgs(UserAttributeProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UserAttributeProtocolMapperArgs.Empty : args;
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
    public static UserAttributeProtocolMapper get(java.lang.String name, Output<java.lang.String> id, @Nullable UserAttributeProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserAttributeProtocolMapper(name, id, state, options);
    }
}
