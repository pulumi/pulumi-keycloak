// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.GenericProtocolMapperArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.GenericProtocolMapperState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing protocol mappers for both types of clients (openid-connect and saml) within Keycloak.
 * 
 * There are two uses cases for using this resource:
 * * If you implemented a custom protocol mapper, this resource can be used to configure it
 * * If the provider doesn&#39;t support a particular protocol mapper, this resource can be used instead.
 * 
 * Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
 * Therefore, if possible, a specific mapper should be used instead.
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
 * import com.pulumi.keycloak.saml.Client;
 * import com.pulumi.keycloak.saml.ClientArgs;
 * import com.pulumi.keycloak.GenericProtocolMapper;
 * import com.pulumi.keycloak.GenericProtocolMapperArgs;
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
 *         var samlClient = new Client("samlClient", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId("test-client")
 *             .build());
 * 
 *         var samlHardcodeAttributeMapper = new GenericProtocolMapper("samlHardcodeAttributeMapper", GenericProtocolMapperArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(samlClient.id())
 *             .name("test-mapper")
 *             .protocol("saml")
 *             .protocolMapper("saml-hardcode-attribute-mapper")
 *             .config(Map.ofEntries(
 *                 Map.entry("attribute.name", "name"),
 *                 Map.entry("attribute.nameformat", "Basic"),
 *                 Map.entry("attribute.value", "value"),
 *                 Map.entry("friendly.name", "display name")
 *             ))
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
 * Protocol mappers can be imported using the following format: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:index/genericProtocolMapper:GenericProtocolMapper saml_hardcode_attribute_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 * 
 */
@ResourceType(type="keycloak:index/genericProtocolMapper:GenericProtocolMapper")
public class GenericProtocolMapper extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the client this protocol mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The ID of the client this protocol mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The ID of the client scope this protocol mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
     * 
     */
    @Export(name="clientScopeId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientScopeId;

    /**
     * @return The ID of the client scope this protocol mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
     * 
     */
    public Output<Optional<String>> clientScopeId() {
        return Codegen.optional(this.clientScopeId);
    }
    /**
     * A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
     * 
     */
    @Export(name="config", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> config;

    /**
     * @return A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
     * 
     */
    public Output<Map<String,Object>> config() {
        return this.config;
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
     * The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
     * 
     */
    @Export(name="protocolMapper", refs={String.class}, tree="[0]")
    private Output<String> protocolMapper;

    /**
     * @return The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
     * 
     */
    public Output<String> protocolMapper() {
        return this.protocolMapper;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GenericProtocolMapper(java.lang.String name) {
        this(name, GenericProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GenericProtocolMapper(java.lang.String name, GenericProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GenericProtocolMapper(java.lang.String name, GenericProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/genericProtocolMapper:GenericProtocolMapper", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GenericProtocolMapper(java.lang.String name, Output<java.lang.String> id, @Nullable GenericProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/genericProtocolMapper:GenericProtocolMapper", name, state, makeResourceOptions(options, id), false);
    }

    private static GenericProtocolMapperArgs makeArgs(GenericProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GenericProtocolMapperArgs.Empty : args;
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
    public static GenericProtocolMapper get(java.lang.String name, Output<java.lang.String> id, @Nullable GenericProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GenericProtocolMapper(name, id, state, options);
    }
}
