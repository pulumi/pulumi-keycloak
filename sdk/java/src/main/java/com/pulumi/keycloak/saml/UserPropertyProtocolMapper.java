// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.saml.UserPropertyProtocolMapperArgs;
import com.pulumi.keycloak.saml.inputs.UserPropertyProtocolMapperState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.saml.UserPropertyProtocolMapper
 * 
 * Allows for creating and managing user property protocol mappers for
 * SAML clients within Keycloak.
 * 
 * SAML user property protocol mappers allow you to map properties of the Keycloak
 * user model to an attribute in a SAML assertion. Protocol mappers
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
 * import com.pulumi.keycloak.saml.UserPropertyProtocolMapper;
 * import com.pulumi.keycloak.saml.UserPropertyProtocolMapperArgs;
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
 *         var samlUserPropertyMapper = new UserPropertyProtocolMapper("samlUserPropertyMapper", UserPropertyProtocolMapperArgs.builder()
 *             .realmId(test.id())
 *             .clientId(samlClient.id())
 *             .name("email-user-property-mapper")
 *             .userProperty("email")
 *             .samlAttributeName("email")
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
 * - `user_property` - (Required) The property of the Keycloak user model to map.
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
@ResourceType(type="keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper")
public class UserPropertyProtocolMapper extends com.pulumi.resources.CustomResource {
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
    @Export(name="userProperty", refs={String.class}, tree="[0]")
    private Output<String> userProperty;

    public Output<String> userProperty() {
        return this.userProperty;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserPropertyProtocolMapper(java.lang.String name) {
        this(name, UserPropertyProtocolMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserPropertyProtocolMapper(java.lang.String name, UserPropertyProtocolMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserPropertyProtocolMapper(java.lang.String name, UserPropertyProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private UserPropertyProtocolMapper(java.lang.String name, Output<java.lang.String> id, @Nullable UserPropertyProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, state, makeResourceOptions(options, id), false);
    }

    private static UserPropertyProtocolMapperArgs makeArgs(UserPropertyProtocolMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UserPropertyProtocolMapperArgs.Empty : args;
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
    public static UserPropertyProtocolMapper get(java.lang.String name, Output<java.lang.String> id, @Nullable UserPropertyProtocolMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserPropertyProtocolMapper(name, id, state, options);
    }
}
