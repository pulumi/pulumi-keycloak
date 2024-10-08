// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.ClientServiceAccountRoleArgs;
import com.pulumi.keycloak.openid.inputs.ClientServiceAccountRoleState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Allows for assigning client roles to the service account of an openid client.
 * You need to set `service_accounts_enabled` to `true` for the openid client that should be assigned the role.
 * 
 * If you&#39;d like to attach realm roles to a service account, please use the `keycloak.openid.ClientServiceAccountRealmRole`
 * resource.
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
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
 * import com.pulumi.keycloak.openid.ClientServiceAccountRole;
 * import com.pulumi.keycloak.openid.ClientServiceAccountRoleArgs;
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
 *         // client1 provides a role to other clients
 *         var client1 = new Client("client1", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .name("client1")
 *             .build());
 * 
 *         var client1Role = new Role("client1Role", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .clientId(client1.id())
 *             .name("my-client1-role")
 *             .description("A role that client1 provides")
 *             .build());
 * 
 *         // client2 is assigned the role of client1
 *         var client2 = new Client("client2", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .name("client2")
 *             .serviceAccountsEnabled(true)
 *             .build());
 * 
 *         var client2ServiceAccountRole = new ClientServiceAccountRole("client2ServiceAccountRole", ClientServiceAccountRoleArgs.builder()
 *             .realmId(realm.id())
 *             .serviceAccountUserId(client2.serviceAccountUserId())
 *             .clientId(client1.id())
 *             .role(client1Role.name())
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
 * This resource can be imported using the format `{{realmId}}/{{serviceAccountUserId}}/{{clientId}}/{{roleId}}`.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole client2_service_account_role my-realm/489ba513-1ceb-49ba-ae0b-1ab1f5099ebf/baf01820-0f8b-4494-9be2-fb3bc8a397a4/c7230ab7-8e4e-4135-995d-e81b50696ad8
 * ```
 * 
 */
@ResourceType(type="keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole")
public class ClientServiceAccountRole extends com.pulumi.resources.CustomResource {
    /**
     * The id of the client that provides the role.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    /**
     * @return The id of the client that provides the role.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * The realm the clients and roles belong to.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm the clients and roles belong to.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }
    /**
     * The name of the role that is assigned.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The name of the role that is assigned.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * The id of the service account that is assigned the role (the service account of the client that &#34;consumes&#34; the role).
     * 
     */
    @Export(name="serviceAccountUserId", refs={String.class}, tree="[0]")
    private Output<String> serviceAccountUserId;

    /**
     * @return The id of the service account that is assigned the role (the service account of the client that &#34;consumes&#34; the role).
     * 
     */
    public Output<String> serviceAccountUserId() {
        return this.serviceAccountUserId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClientServiceAccountRole(java.lang.String name) {
        this(name, ClientServiceAccountRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClientServiceAccountRole(java.lang.String name, ClientServiceAccountRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClientServiceAccountRole(java.lang.String name, ClientServiceAccountRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ClientServiceAccountRole(java.lang.String name, Output<java.lang.String> id, @Nullable ClientServiceAccountRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole", name, state, makeResourceOptions(options, id), false);
    }

    private static ClientServiceAccountRoleArgs makeArgs(ClientServiceAccountRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ClientServiceAccountRoleArgs.Empty : args;
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
    public static ClientServiceAccountRole get(java.lang.String name, Output<java.lang.String> id, @Nullable ClientServiceAccountRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClientServiceAccountRole(name, id, state, options);
    }
}
