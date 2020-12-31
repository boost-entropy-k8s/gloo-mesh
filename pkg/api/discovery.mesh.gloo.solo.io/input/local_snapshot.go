// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./local_snapshot.go -destination mocks/local_snapshot.go

// The Input LocalSnapshot contains the set of all:
// * SettingsMeshGlooSoloIov1Alpha2Settings
// read from a given cluster or set of clusters, across all namespaces.
//
// A snapshot can be constructed from either a single Manager (for a single cluster)
// or a ClusterWatcher (for multiple clusters) using the LocalSnapshotBuilder.
//
// Resources in a MultiCluster snapshot will have their ClusterName set to the
// name of the cluster from which the resource was read.

package input

import (
	"context"
	"encoding/json"

	"github.com/solo-io/skv2/pkg/verifier"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/hashicorp/go-multierror"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	settings_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2"
	settings_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/sets"
)

// the snapshot of input resources consumed by translation
type LocalSnapshot interface {

	// return the set of input SettingsMeshGlooSoloIov1Alpha2Settings
	SettingsMeshGlooSoloIov1Alpha2Settings() settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet
	// update the status of all input objects which support
	// the Status subresource (in the local cluster)
	SyncStatuses(ctx context.Context, c client.Client, opts LocalSyncStatusOptions) error
	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

// options for syncing input object statuses
type LocalSyncStatusOptions struct {

	// sync status of SettingsMeshGlooSoloIov1Alpha2Settings objects
	SettingsMeshGlooSoloIov1Alpha2Settings bool
}

type snapshotLocal struct {
	name string

	settingsMeshGlooSoloIov1Alpha2Settings settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet
}

func NewLocalSnapshot(
	name string,

	settingsMeshGlooSoloIov1Alpha2Settings settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet,

) LocalSnapshot {
	return &snapshotLocal{
		name: name,

		settingsMeshGlooSoloIov1Alpha2Settings: settingsMeshGlooSoloIov1Alpha2Settings,
	}
}

func (s snapshotLocal) SettingsMeshGlooSoloIov1Alpha2Settings() settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet {
	return s.settingsMeshGlooSoloIov1Alpha2Settings
}

func (s snapshotLocal) SyncStatuses(ctx context.Context, c client.Client, opts LocalSyncStatusOptions) error {
	var errs error

	if opts.SettingsMeshGlooSoloIov1Alpha2Settings {
		for _, obj := range s.SettingsMeshGlooSoloIov1Alpha2Settings().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	return errs
}

func (s snapshotLocal) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["settingsMeshGlooSoloIov1Alpha2Settings"] = s.settingsMeshGlooSoloIov1Alpha2Settings.List()
	return json.Marshal(snapshotMap)
}

// builds the input snapshot from API Clients.
type LocalBuilder interface {
	BuildSnapshot(ctx context.Context, name string, opts LocalBuildOptions) (LocalSnapshot, error)
}

// Options for building a snapshot
type LocalBuildOptions struct {

	// List options for composing a snapshot from SettingsMeshGlooSoloIov1Alpha2Settings
	SettingsMeshGlooSoloIov1Alpha2Settings ResourceLocalBuildOptions
}

// Options for reading resources of a given type
type ResourceLocalBuildOptions struct {

	// List options for composing a snapshot from a resource type
	ListOptions []client.ListOption

	// If provided, ensure the resource has been verified before adding it to snapshots
	Verifier verifier.ServerResourceVerifier
}

// build a snapshot from resources in a single cluster
type singleClusterLocalBuilder struct {
	mgr manager.Manager
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewSingleClusterLocalBuilder(
	mgr manager.Manager,
) LocalBuilder {
	return &singleClusterLocalBuilder{
		mgr: mgr,
	}
}

func (b *singleClusterLocalBuilder) BuildSnapshot(ctx context.Context, name string, opts LocalBuildOptions) (LocalSnapshot, error) {

	settingsMeshGlooSoloIov1Alpha2Settings := settings_mesh_gloo_solo_io_v1alpha2_sets.NewSettingsSet()

	var errs error

	if err := b.insertSettingsMeshGlooSoloIov1Alpha2Settings(ctx, settingsMeshGlooSoloIov1Alpha2Settings, opts.SettingsMeshGlooSoloIov1Alpha2Settings); err != nil {
		errs = multierror.Append(errs, err)
	}

	outputSnap := NewLocalSnapshot(
		name,

		settingsMeshGlooSoloIov1Alpha2Settings,
	)

	return outputSnap, errs
}

func (b *singleClusterLocalBuilder) insertSettingsMeshGlooSoloIov1Alpha2Settings(ctx context.Context, settingsMeshGlooSoloIov1Alpha2Settings settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet, opts ResourceLocalBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "settings.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "Settings",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	settingsMeshGlooSoloIov1Alpha2SettingsList, err := settings_mesh_gloo_solo_io_v1alpha2.NewSettingsClient(b.mgr.GetClient()).ListSettings(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range settingsMeshGlooSoloIov1Alpha2SettingsList.Items {
		item := item // pike
		settingsMeshGlooSoloIov1Alpha2Settings.Insert(&item)
	}

	return nil
}
