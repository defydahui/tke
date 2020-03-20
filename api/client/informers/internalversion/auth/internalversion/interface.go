/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	internalinterfaces "tkestack.io/tke/api/client/informers/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// APIKeys returns a APIKeyInformer.
	APIKeys() APIKeyInformer
	// APISigningKeys returns a APISigningKeyInformer.
	APISigningKeys() APISigningKeyInformer
	// Categories returns a CategoryInformer.
	Categories() CategoryInformer
	// Clients returns a ClientInformer.
	Clients() ClientInformer
	// ConfigMaps returns a ConfigMapInformer.
	ConfigMaps() ConfigMapInformer
	// Groups returns a GroupInformer.
	Groups() GroupInformer
	// IdentityProviders returns a IdentityProviderInformer.
	IdentityProviders() IdentityProviderInformer
	// LocalGroups returns a LocalGroupInformer.
	LocalGroups() LocalGroupInformer
	// LocalIdentities returns a LocalIdentityInformer.
	LocalIdentities() LocalIdentityInformer
	// Policies returns a PolicyInformer.
	Policies() PolicyInformer
	// ProjectPolicyBindings returns a ProjectPolicyBindingInformer.
	ProjectPolicyBindings() ProjectPolicyBindingInformer
	// ProjectRoles returns a ProjectRoleInformer.
	ProjectRoles() ProjectRoleInformer
	// Roles returns a RoleInformer.
	Roles() RoleInformer
	// Rules returns a RuleInformer.
	Rules() RuleInformer
	// Users returns a UserInformer.
	Users() UserInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// APIKeys returns a APIKeyInformer.
func (v *version) APIKeys() APIKeyInformer {
	return &aPIKeyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// APISigningKeys returns a APISigningKeyInformer.
func (v *version) APISigningKeys() APISigningKeyInformer {
	return &aPISigningKeyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Categories returns a CategoryInformer.
func (v *version) Categories() CategoryInformer {
	return &categoryInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Clients returns a ClientInformer.
func (v *version) Clients() ClientInformer {
	return &clientInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ConfigMaps returns a ConfigMapInformer.
func (v *version) ConfigMaps() ConfigMapInformer {
	return &configMapInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Groups returns a GroupInformer.
func (v *version) Groups() GroupInformer {
	return &groupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// IdentityProviders returns a IdentityProviderInformer.
func (v *version) IdentityProviders() IdentityProviderInformer {
	return &identityProviderInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalGroups returns a LocalGroupInformer.
func (v *version) LocalGroups() LocalGroupInformer {
	return &localGroupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalIdentities returns a LocalIdentityInformer.
func (v *version) LocalIdentities() LocalIdentityInformer {
	return &localIdentityInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Policies returns a PolicyInformer.
func (v *version) Policies() PolicyInformer {
	return &policyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ProjectPolicyBindings returns a ProjectPolicyBindingInformer.
func (v *version) ProjectPolicyBindings() ProjectPolicyBindingInformer {
	return &projectPolicyBindingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ProjectRoles returns a ProjectRoleInformer.
func (v *version) ProjectRoles() ProjectRoleInformer {
	return &projectRoleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Roles returns a RoleInformer.
func (v *version) Roles() RoleInformer {
	return &roleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Rules returns a RuleInformer.
func (v *version) Rules() RuleInformer {
	return &ruleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Users returns a UserInformer.
func (v *version) Users() UserInformer {
	return &userInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
