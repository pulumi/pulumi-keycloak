// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class ClientGroupPolicyGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClientGroupPolicyGroupArgs Empty = new ClientGroupPolicyGroupArgs();

    @Import(name="extendChildren", required=true)
    private Output<Boolean> extendChildren;

    public Output<Boolean> extendChildren() {
        return this.extendChildren;
    }

    @Import(name="id", required=true)
    private Output<String> id;

    public Output<String> id() {
        return this.id;
    }

    @Import(name="path", required=true)
    private Output<String> path;

    public Output<String> path() {
        return this.path;
    }

    private ClientGroupPolicyGroupArgs() {}

    private ClientGroupPolicyGroupArgs(ClientGroupPolicyGroupArgs $) {
        this.extendChildren = $.extendChildren;
        this.id = $.id;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientGroupPolicyGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientGroupPolicyGroupArgs $;

        public Builder() {
            $ = new ClientGroupPolicyGroupArgs();
        }

        public Builder(ClientGroupPolicyGroupArgs defaults) {
            $ = new ClientGroupPolicyGroupArgs(Objects.requireNonNull(defaults));
        }

        public Builder extendChildren(Output<Boolean> extendChildren) {
            $.extendChildren = extendChildren;
            return this;
        }

        public Builder extendChildren(Boolean extendChildren) {
            return extendChildren(Output.of(extendChildren));
        }

        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        public Builder id(String id) {
            return id(Output.of(id));
        }

        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        public Builder path(String path) {
            return path(Output.of(path));
        }

        public ClientGroupPolicyGroupArgs build() {
            if ($.extendChildren == null) {
                throw new MissingRequiredPropertyException("ClientGroupPolicyGroupArgs", "extendChildren");
            }
            if ($.id == null) {
                throw new MissingRequiredPropertyException("ClientGroupPolicyGroupArgs", "id");
            }
            if ($.path == null) {
                throw new MissingRequiredPropertyException("ClientGroupPolicyGroupArgs", "path");
            }
            return $;
        }
    }

}
