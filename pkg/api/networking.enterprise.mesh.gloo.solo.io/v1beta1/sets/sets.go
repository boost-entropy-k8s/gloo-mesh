// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1beta1sets

import (
	networking_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type WasmDeploymentSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	// Insert a resource into the set.
	Insert(wasmDeployment ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(wasmDeploymentSet WasmDeploymentSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(wasmDeployment ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(wasmDeployment ezkube.ResourceId)
	// Return the union with the provided set
	Union(set WasmDeploymentSet) WasmDeploymentSet
	// Return the difference with the provided set
	Difference(set WasmDeploymentSet) WasmDeploymentSet
	// Return the intersection with the provided set
	Intersection(set WasmDeploymentSet) WasmDeploymentSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another WasmDeploymentSet
	Delta(newSet WasmDeploymentSet) sksets.ResourceDelta
}

func makeGenericWasmDeploymentSet(wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range wasmDeploymentList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type wasmDeploymentSet struct {
	set sksets.ResourceSet
}

func NewWasmDeploymentSet(wasmDeploymentList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) WasmDeploymentSet {
	return &wasmDeploymentSet{set: makeGenericWasmDeploymentSet(wasmDeploymentList)}
}

func NewWasmDeploymentSetFromList(wasmDeploymentList *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeploymentList) WasmDeploymentSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, 0, len(wasmDeploymentList.Items))
	for idx := range wasmDeploymentList.Items {
		list = append(list, &wasmDeploymentList.Items[idx])
	}
	return &wasmDeploymentSet{set: makeGenericWasmDeploymentSet(list)}
}

func (s *wasmDeploymentSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *wasmDeploymentSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
		})
	}

	objs := s.Generic().List(genericFilters...)
	wasmDeploymentList := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, 0, len(objs))
	for _, obj := range objs {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
	}
	return wasmDeploymentList
}

func (s *wasmDeploymentSet) UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
		})
	}

	var wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
	}
	return wasmDeploymentList
}

func (s *wasmDeploymentSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
	}
	return newMap
}

func (s *wasmDeploymentSet) Insert(
	wasmDeploymentList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range wasmDeploymentList {
		s.Generic().Insert(obj)
	}
}

func (s *wasmDeploymentSet) Has(wasmDeployment ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(wasmDeployment)
}

func (s *wasmDeploymentSet) Equal(
	wasmDeploymentSet WasmDeploymentSet,
) bool {
	if s == nil {
		return wasmDeploymentSet == nil
	}
	return s.Generic().Equal(wasmDeploymentSet.Generic())
}

func (s *wasmDeploymentSet) Delete(WasmDeployment ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(WasmDeployment)
}

func (s *wasmDeploymentSet) Union(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return set
	}
	return NewWasmDeploymentSet(append(s.List(), set.List()...)...)
}

func (s *wasmDeploymentSet) Difference(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &wasmDeploymentSet{set: newSet}
}

func (s *wasmDeploymentSet) Intersection(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	for _, obj := range newSet.List() {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
	}
	return NewWasmDeploymentSet(wasmDeploymentList...)
}

func (s *wasmDeploymentSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find WasmDeployment %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment), nil
}

func (s *wasmDeploymentSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *wasmDeploymentSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *wasmDeploymentSet) Delta(newSet WasmDeploymentSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type VirtualDestinationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	// Insert a resource into the set.
	Insert(virtualDestination ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualDestinationSet VirtualDestinationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualDestination ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualDestination ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualDestinationSet) VirtualDestinationSet
	// Return the difference with the provided set
	Difference(set VirtualDestinationSet) VirtualDestinationSet
	// Return the intersection with the provided set
	Intersection(set VirtualDestinationSet) VirtualDestinationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualDestinationSet
	Delta(newSet VirtualDestinationSet) sksets.ResourceDelta
}

func makeGenericVirtualDestinationSet(virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualDestinationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualDestinationSet struct {
	set sksets.ResourceSet
}

func NewVirtualDestinationSet(virtualDestinationList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) VirtualDestinationSet {
	return &virtualDestinationSet{set: makeGenericVirtualDestinationSet(virtualDestinationList)}
}

func NewVirtualDestinationSetFromList(virtualDestinationList *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestinationList) VirtualDestinationSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, 0, len(virtualDestinationList.Items))
	for idx := range virtualDestinationList.Items {
		list = append(list, &virtualDestinationList.Items[idx])
	}
	return &virtualDestinationSet{set: makeGenericVirtualDestinationSet(list)}
}

func (s *virtualDestinationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualDestinationSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
		})
	}

	objs := s.Generic().List(genericFilters...)
	virtualDestinationList := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, 0, len(objs))
	for _, obj := range objs {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
	}
	return virtualDestinationList
}

func (s *virtualDestinationSet) UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
		})
	}

	var virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
	}
	return virtualDestinationList
}

func (s *virtualDestinationSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
	}
	return newMap
}

func (s *virtualDestinationSet) Insert(
	virtualDestinationList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualDestinationList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualDestinationSet) Has(virtualDestination ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualDestination)
}

func (s *virtualDestinationSet) Equal(
	virtualDestinationSet VirtualDestinationSet,
) bool {
	if s == nil {
		return virtualDestinationSet == nil
	}
	return s.Generic().Equal(virtualDestinationSet.Generic())
}

func (s *virtualDestinationSet) Delete(VirtualDestination ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualDestination)
}

func (s *virtualDestinationSet) Union(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return set
	}
	return NewVirtualDestinationSet(append(s.List(), set.List()...)...)
}

func (s *virtualDestinationSet) Difference(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualDestinationSet{set: newSet}
}

func (s *virtualDestinationSet) Intersection(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	for _, obj := range newSet.List() {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
	}
	return NewVirtualDestinationSet(virtualDestinationList...)
}

func (s *virtualDestinationSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualDestination %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination), nil
}

func (s *virtualDestinationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualDestinationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualDestinationSet) Delta(newSet VirtualDestinationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}
