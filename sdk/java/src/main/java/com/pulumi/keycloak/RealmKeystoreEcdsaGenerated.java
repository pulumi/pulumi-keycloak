// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.RealmKeystoreEcdsaGeneratedArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.RealmKeystoreEcdsaGeneratedState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing `acdsa_generated` Realm keystores within Keycloak.
 * 
 * A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
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
 * import com.pulumi.keycloak.RealmKeystoreEcdsaGenerated;
 * import com.pulumi.keycloak.RealmKeystoreEcdsaGeneratedArgs;
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
 *             .build());
 * 
 *         var keystoreEcdsaGenerated = new RealmKeystoreEcdsaGenerated("keystoreEcdsaGenerated", RealmKeystoreEcdsaGeneratedArgs.builder()
 *             .name("my-ecdsa-generated-key")
 *             .realmId(realm.id())
 *             .enabled(true)
 *             .active(true)
 *             .priority(100)
 *             .ellipticCurveKey("P-256")
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
 * Realm keys can be imported using realm name and keystore id, you can find it in web UI.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:index/realmKeystoreEcdsaGenerated:RealmKeystoreEcdsaGenerated keystore_ecdsa_generated my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
 * ```
 * 
 */
@ResourceType(type="keycloak:index/realmKeystoreEcdsaGenerated:RealmKeystoreEcdsaGenerated")
public class RealmKeystoreEcdsaGenerated extends com.pulumi.resources.CustomResource {
    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> active;

    /**
     * @return When `false`, key in not used for signing. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> active() {
        return Codegen.optional(this.active);
    }
    /**
     * Elliptic Curve used in ECDSA. Defaults to `P-256`.
     * 
     */
    @Export(name="ellipticCurveKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ellipticCurveKey;

    /**
     * @return Elliptic Curve used in ECDSA. Defaults to `P-256`.
     * 
     */
    public Output<Optional<String>> ellipticCurveKey() {
        return Codegen.optional(this.ellipticCurveKey);
    }
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When `false`, key is not accessible in this realm. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Display name of provider when linked in admin console.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name of provider when linked in admin console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Priority for the provider. Defaults to `0`
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return Priority for the provider. Defaults to `0`
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The realm this keystore exists in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm this keystore exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RealmKeystoreEcdsaGenerated(java.lang.String name) {
        this(name, RealmKeystoreEcdsaGeneratedArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RealmKeystoreEcdsaGenerated(java.lang.String name, RealmKeystoreEcdsaGeneratedArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RealmKeystoreEcdsaGenerated(java.lang.String name, RealmKeystoreEcdsaGeneratedArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmKeystoreEcdsaGenerated:RealmKeystoreEcdsaGenerated", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RealmKeystoreEcdsaGenerated(java.lang.String name, Output<java.lang.String> id, @Nullable RealmKeystoreEcdsaGeneratedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmKeystoreEcdsaGenerated:RealmKeystoreEcdsaGenerated", name, state, makeResourceOptions(options, id), false);
    }

    private static RealmKeystoreEcdsaGeneratedArgs makeArgs(RealmKeystoreEcdsaGeneratedArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RealmKeystoreEcdsaGeneratedArgs.Empty : args;
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
    public static RealmKeystoreEcdsaGenerated get(java.lang.String name, Output<java.lang.String> id, @Nullable RealmKeystoreEcdsaGeneratedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RealmKeystoreEcdsaGenerated(name, id, state, options);
    }
}
