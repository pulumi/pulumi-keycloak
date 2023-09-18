// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.GenericClientProtocolMapperArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.GenericClientProtocolMapperState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * !&gt; **WARNING:** This resource is deprecated and will be removed in the next major version. Please use `keycloak.GenericProtocolMapper` instead.
 * 
 * Allows for creating and managing protocol mappers for both types of clients (openid-connect and saml) within Keycloak.
 * 
 * There are two uses cases for using this resource:
 * * If you implemented a custom protocol mapper, this resource can be used to configure it
 * * If the provider doesn&#39;t support a particular protocol mapper, this resource can be used instead.
 * 
 * Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
 * Therefore, if possible, a specific mapper should be used.
 * 
 * ## Example Usage
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
 * import com.pulumi.keycloak.GenericClientProtocolMapper;
 * import com.pulumi.keycloak.GenericClientProtocolMapperArgs;
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
 *             .clientId(&#34;test-client&#34;)
 *             .build());
 * 
 *         var samlHardcodeAttributeMapper = new GenericClientProtocolMapper(&#34;samlHardcodeAttributeMapper&#34;, GenericClientProtocolMapperArgs.builder()        
 *             .realmId(realm.id())
 *             .clientId(samlClient.id())
 *             .protocol(&#34;saml&#34;)
 *             .protocolMapper(&#34;saml-hardcode-attribute-mapper&#34;)
 *             .config(Map.ofEntries(
 *                 Map.entry(&#34;attribute.name&#34;, &#34;name&#34;),
 *                 Map.entry(&#34;attribute.nameformat&#34;, &#34;Basic&#34;),
 *                 Map.entry(&#34;attribute.value&#34;, &#34;value&#34;),
 *                 Map.entry(&#34;friendly.name&#34;, &#34;display name&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Protocol mappers can be imported using the following format`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper saml_hardcode_attribute_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 * 
 */
@ResourceType(type="keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper")
public class GenericClientProtocolMapper extends com.pulumi.resources.CustomResource {
    /**
     * The client this protocol mapper is attached to.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The client this protocol mapper is attached to.
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
    public GenericClientProtocolMapper(String name) {
        this(name, GenericClientProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GenericClientProtocolMapper(String name, GenericClientProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GenericClientProtocolMapper(String name, GenericClientProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper", name, args == null ? GenericClientProtocolMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GenericClientProtocolMapper(String name, Output<String> id, @Nullable GenericClientProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper", name, state, makeResourceOptions(options, id));
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
    public static GenericClientProtocolMapper get(String name, Output<String> id, @Nullable GenericClientProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GenericClientProtocolMapper(name, id, state, options);
    }
}
