// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gardener

import (
	"crypto/x509"
	"reflect"
	"strings"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"

	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
)

const (
	// SeedNamespaceNamePrefix is the prefix used for seed namespaces.
	SeedNamespaceNamePrefix = "seed-"
)

// ComputeGardenNamespace returns the name of the namespace belonging to the given seed in the Garden cluster.
func ComputeGardenNamespace(seedName string) string {
	return SeedNamespaceNamePrefix + seedName
}

// ComputeSeedName computes the name of the seed out of the seed namespace in the Garden cluster.
func ComputeSeedName(seedNamespaceName string) string {
	seedName := strings.TrimPrefix(seedNamespaceName, SeedNamespaceNamePrefix)
	if seedName == seedNamespaceName {
		return ""
	}
	return seedName
}

// IsSeedClientCert returns true when the given CSR and usages match the requirements for a client certificate for a
// seed.
func IsSeedClientCert(x509cr *x509.CertificateRequest, usages []certificatesv1beta1.KeyUsage) bool {
	if !reflect.DeepEqual([]string{v1beta1constants.SeedsGroup}, x509cr.Subject.Organization) {
		return false
	}

	if (len(x509cr.DNSNames) > 0) || (len(x509cr.EmailAddresses) > 0) || (len(x509cr.IPAddresses) > 0) {
		return false
	}

	if !hasExactUsages(usages, []certificatesv1beta1.KeyUsage{
		certificatesv1beta1.UsageKeyEncipherment,
		certificatesv1beta1.UsageDigitalSignature,
		certificatesv1beta1.UsageClientAuth,
	}) {
		return false
	}

	return strings.HasPrefix(x509cr.Subject.CommonName, v1beta1constants.SeedUserNamePrefix)
}

func hasExactUsages(usages, requiredUsages []certificatesv1beta1.KeyUsage) bool {
	if len(requiredUsages) != len(usages) {
		return false
	}

	usageMap := map[certificatesv1beta1.KeyUsage]struct{}{}
	for _, u := range usages {
		usageMap[u] = struct{}{}
	}

	for _, u := range requiredUsages {
		if _, ok := usageMap[u]; !ok {
			return false
		}
	}

	return true
}
