// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.RealmKeystoreRsaArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.RealmKeystoreRsaState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing `rsa` Realm keystores within Keycloak.
 * 
 * A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.RealmKeystoreRsa;
 * import com.pulumi.keycloak.RealmKeystoreRsaArgs;
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
 *             .build());
 * 
 *         var keystoreRsa = new RealmKeystoreRsa(&#34;keystoreRsa&#34;, RealmKeystoreRsaArgs.builder()        
 *             .realmId(realm.id())
 *             .enabled(true)
 *             .active(true)
 *             .privateKey(&#34;&lt;your rsa private key&gt;&#34;)
 *             .certificate(&#34;&lt;your certificate&gt;&#34;)
 *             .priority(100)
 *             .algorithm(&#34;RS256&#34;)
 *             .keystoreSize(2048)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:index/realmKeystoreRsa:RealmKeystoreRsa keystore_rsa my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
 * ```
 * 
 */
@ResourceType(type="keycloak:index/realmKeystoreRsa:RealmKeystoreRsa")
public class RealmKeystoreRsa extends com.pulumi.resources.CustomResource {
    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     * 
     */
    @Export(name="active", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> active;

    /**
     * @return When `false`, key in not used for signing. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> active() {
        return Codegen.optional(this.active);
    }
    /**
     * Intended algorithm for the key. Defaults to `RS256`
     * 
     */
    @Export(name="algorithm", type=String.class, parameters={})
    private Output</* @Nullable */ String> algorithm;

    /**
     * @return Intended algorithm for the key. Defaults to `RS256`
     * 
     */
    public Output<Optional<String>> algorithm() {
        return Codegen.optional(this.algorithm);
    }
    /**
     * X509 Certificate encoded in PEM format.
     * 
     */
    @Export(name="certificate", type=String.class, parameters={})
    private Output<String> certificate;

    /**
     * @return X509 Certificate encoded in PEM format.
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
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
    @Export(name="name", type=String.class, parameters={})
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
    @Export(name="priority", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return Priority for the provider. Defaults to `0`
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * Private RSA Key encoded in PEM format.
     * 
     */
    @Export(name="privateKey", type=String.class, parameters={})
    private Output<String> privateKey;

    /**
     * @return Private RSA Key encoded in PEM format.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * The realm this keystore exists in.
     * 
     */
    @Export(name="realmId", type=String.class, parameters={})
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
    public RealmKeystoreRsa(String name) {
        this(name, RealmKeystoreRsaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RealmKeystoreRsa(String name, RealmKeystoreRsaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RealmKeystoreRsa(String name, RealmKeystoreRsaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmKeystoreRsa:RealmKeystoreRsa", name, args == null ? RealmKeystoreRsaArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RealmKeystoreRsa(String name, Output<String> id, @Nullable RealmKeystoreRsaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmKeystoreRsa:RealmKeystoreRsa", name, state, makeResourceOptions(options, id));
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
    public static RealmKeystoreRsa get(String name, Output<String> id, @Nullable RealmKeystoreRsaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RealmKeystoreRsa(name, id, state, options);
    }
}