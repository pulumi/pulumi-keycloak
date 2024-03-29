// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientScopeState extends com.pulumi.resources.ResourceArgs {

    public static final ClientScopeState Empty = new ClientScopeState();

    @Import(name="consentScreenText")
    private @Nullable Output<String> consentScreenText;

    public Optional<Output<String>> consentScreenText() {
        return Optional.ofNullable(this.consentScreenText);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="guiOrder")
    private @Nullable Output<Integer> guiOrder;

    public Optional<Output<Integer>> guiOrder() {
        return Optional.ofNullable(this.guiOrder);
    }

    @Import(name="includeInTokenScope")
    private @Nullable Output<Boolean> includeInTokenScope;

    public Optional<Output<Boolean>> includeInTokenScope() {
        return Optional.ofNullable(this.includeInTokenScope);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private ClientScopeState() {}

    private ClientScopeState(ClientScopeState $) {
        this.consentScreenText = $.consentScreenText;
        this.description = $.description;
        this.guiOrder = $.guiOrder;
        this.includeInTokenScope = $.includeInTokenScope;
        this.name = $.name;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientScopeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientScopeState $;

        public Builder() {
            $ = new ClientScopeState();
        }

        public Builder(ClientScopeState defaults) {
            $ = new ClientScopeState(Objects.requireNonNull(defaults));
        }

        public Builder consentScreenText(@Nullable Output<String> consentScreenText) {
            $.consentScreenText = consentScreenText;
            return this;
        }

        public Builder consentScreenText(String consentScreenText) {
            return consentScreenText(Output.of(consentScreenText));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder guiOrder(@Nullable Output<Integer> guiOrder) {
            $.guiOrder = guiOrder;
            return this;
        }

        public Builder guiOrder(Integer guiOrder) {
            return guiOrder(Output.of(guiOrder));
        }

        public Builder includeInTokenScope(@Nullable Output<Boolean> includeInTokenScope) {
            $.includeInTokenScope = includeInTokenScope;
            return this;
        }

        public Builder includeInTokenScope(Boolean includeInTokenScope) {
            return includeInTokenScope(Output.of(includeInTokenScope));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public ClientScopeState build() {
            return $;
        }
    }

}
