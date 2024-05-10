// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.saml.UserAttributeProtocolMapperArgs;
import com.pulumi.keycloak.saml.inputs.UserAttributeProtocolMapperState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.saml.UserAttributeProtocolMapper
 * 
 * Allows for creating and managing user attribute protocol mappers for
 * SAML clients within Keycloak.
 * 
 * SAML user attribute protocol mappers allow you to map custom attributes defined
 * for a user within Keycloak to an attribute in a SAML assertion. Protocol mappers
 * can be defined for a single client, or they can be defined for a client scope which
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
 * import com.pulumi.keycloak.saml.Client;
 * import com.pulumi.keycloak.saml.ClientArgs;
 * import com.pulumi.keycloak.saml.UserAttributeProtocolMapper;
 * import com.pulumi.keycloak.saml.UserAttributeProtocolMapperArgs;
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
 *             .realmId(test.id())
 *             .clientId("test-saml-client")
 *             .name("test-saml-client")
 *             .build());
 * 
 *         var samlUserAttributeMapper = new UserAttributeProtocolMapper("samlUserAttributeMapper", UserAttributeProtocolMapperArgs.builder()        
 *             .realmId(test.id())
 *             .clientId(samlClient.id())
 *             .name("displayname-user-attribute-mapper")
 *             .userAttribute("displayName")
 *             .samlAttributeName("displayName")
 *             .samlAttributeNameFormat("Unspecified")
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
 * - `client_id` - (Required if `client_scope_id` is not specified) The SAML client this protocol mapper is attached to.
 * - `client_scope_id` - (Required if `client_id` is not specified) The SAML client scope this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the GUI.
 * - `user_attribute` - (Required) The custom user attribute to map.
 * - `friendly_name` - (Optional) An optional human-friendly name for this attribute.
 * - `saml_attribute_name` - (Required) The name of the SAML attribute.
 * - `saml_attribute_name_format` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
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
@ResourceType(type="keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper")
public class UserAttributeProtocolMapper extends com.pulumi.resources.CustomResource {
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    @Export(name="clientScopeId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientScopeId;

    public Output<Optional<String>> clientScopeId() {
        return Codegen.optional(this.clientScopeId);
    }
    @Export(name="friendlyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> friendlyName;

    public Output<Optional<String>> friendlyName() {
        return Codegen.optional(this.friendlyName);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="samlAttributeName", refs={String.class}, tree="[0]")
    private Output<String> samlAttributeName;

    public Output<String> samlAttributeName() {
        return this.samlAttributeName;
    }
    @Export(name="samlAttributeNameFormat", refs={String.class}, tree="[0]")
    private Output<String> samlAttributeNameFormat;

    public Output<String> samlAttributeNameFormat() {
        return this.samlAttributeNameFormat;
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
    public UserAttributeProtocolMapper(String name) {
        this(name, UserAttributeProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserAttributeProtocolMapper(String name, UserAttributeProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserAttributeProtocolMapper(String name, UserAttributeProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, args == null ? UserAttributeProtocolMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserAttributeProtocolMapper(String name, Output<String> id, @Nullable UserAttributeProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, state, makeResourceOptions(options, id));
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
    public static UserAttributeProtocolMapper get(String name, Output<String> id, @Nullable UserAttributeProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserAttributeProtocolMapper(name, id, state, options);
    }
}
