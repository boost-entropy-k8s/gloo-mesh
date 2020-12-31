// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	certificates_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2"
	certificates_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2/sets"

	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "k8s.io/api/core/v1"
)

type InputSnapshotManualBuilder struct {
	name string

	certificatesMeshGlooSoloIov1Alpha2IssuedCertificates  certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet
	certificatesMeshGlooSoloIov1Alpha2CertificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet
	certificatesMeshGlooSoloIov1Alpha2PodBounceDirectives certificates_mesh_gloo_solo_io_v1alpha2_sets.PodBounceDirectiveSet

	v1Secrets v1_sets.SecretSet
	v1Pods    v1_sets.PodSet
}

func NewInputSnapshotManualBuilder(name string) *InputSnapshotManualBuilder {
	return &InputSnapshotManualBuilder{
		name: name,

		certificatesMeshGlooSoloIov1Alpha2IssuedCertificates:  certificates_mesh_gloo_solo_io_v1alpha2_sets.NewIssuedCertificateSet(),
		certificatesMeshGlooSoloIov1Alpha2CertificateRequests: certificates_mesh_gloo_solo_io_v1alpha2_sets.NewCertificateRequestSet(),
		certificatesMeshGlooSoloIov1Alpha2PodBounceDirectives: certificates_mesh_gloo_solo_io_v1alpha2_sets.NewPodBounceDirectiveSet(),

		v1Secrets: v1_sets.NewSecretSet(),
		v1Pods:    v1_sets.NewPodSet(),
	}
}

func (i *InputSnapshotManualBuilder) Build() Snapshot {
	return NewSnapshot(
		i.name,

		i.certificatesMeshGlooSoloIov1Alpha2IssuedCertificates,
		i.certificatesMeshGlooSoloIov1Alpha2CertificateRequests,
		i.certificatesMeshGlooSoloIov1Alpha2PodBounceDirectives,

		i.v1Secrets,
		i.v1Pods,
	)
}
func (i *InputSnapshotManualBuilder) AddCertificatesMeshGlooSoloIov1Alpha2IssuedCertificates(certificatesMeshGlooSoloIov1Alpha2IssuedCertificates []*certificates_mesh_gloo_solo_io_v1alpha2.IssuedCertificate) *InputSnapshotManualBuilder {
	i.certificatesMeshGlooSoloIov1Alpha2IssuedCertificates.Insert(certificatesMeshGlooSoloIov1Alpha2IssuedCertificates...)
	return i
}
func (i *InputSnapshotManualBuilder) AddCertificatesMeshGlooSoloIov1Alpha2CertificateRequests(certificatesMeshGlooSoloIov1Alpha2CertificateRequests []*certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequest) *InputSnapshotManualBuilder {
	i.certificatesMeshGlooSoloIov1Alpha2CertificateRequests.Insert(certificatesMeshGlooSoloIov1Alpha2CertificateRequests...)
	return i
}
func (i *InputSnapshotManualBuilder) AddCertificatesMeshGlooSoloIov1Alpha2PodBounceDirectives(certificatesMeshGlooSoloIov1Alpha2PodBounceDirectives []*certificates_mesh_gloo_solo_io_v1alpha2.PodBounceDirective) *InputSnapshotManualBuilder {
	i.certificatesMeshGlooSoloIov1Alpha2PodBounceDirectives.Insert(certificatesMeshGlooSoloIov1Alpha2PodBounceDirectives...)
	return i
}
func (i *InputSnapshotManualBuilder) AddV1Secrets(v1Secrets []*v1.Secret) *InputSnapshotManualBuilder {
	i.v1Secrets.Insert(v1Secrets...)
	return i
}
func (i *InputSnapshotManualBuilder) AddV1Pods(v1Pods []*v1.Pod) *InputSnapshotManualBuilder {
	i.v1Pods.Insert(v1Pods...)
	return i
}
