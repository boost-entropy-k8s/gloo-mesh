// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

    networking_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"

    "github.com/pkg/errors"
    "github.com/solo-io/skv2/pkg/events"
    "sigs.k8s.io/controller-runtime/pkg/manager"
    "sigs.k8s.io/controller-runtime/pkg/predicate"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

// Handle events for the WasmDeployment Resource
// DEPRECATED: Prefer reconciler pattern.
type WasmDeploymentEventHandler interface {
    CreateWasmDeployment(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error
    UpdateWasmDeployment(old, new *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error
    DeleteWasmDeployment(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error
    GenericWasmDeployment(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error
}

type WasmDeploymentEventHandlerFuncs struct {
    OnCreate  func(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error
    OnUpdate  func(old, new *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error
    OnDelete  func(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error
    OnGeneric func(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error
}

func (f *WasmDeploymentEventHandlerFuncs) CreateWasmDeployment(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *WasmDeploymentEventHandlerFuncs) DeleteWasmDeployment(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *WasmDeploymentEventHandlerFuncs) UpdateWasmDeployment(objOld, objNew *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *WasmDeploymentEventHandlerFuncs) GenericWasmDeployment(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type WasmDeploymentEventWatcher interface {
    AddEventHandler(ctx context.Context, h WasmDeploymentEventHandler, predicates ...predicate.Predicate) error
}

type wasmDeploymentEventWatcher struct {
    watcher events.EventWatcher
}

func NewWasmDeploymentEventWatcher(name string, mgr manager.Manager) WasmDeploymentEventWatcher {
    return &wasmDeploymentEventWatcher{
        watcher: events.NewWatcher(name, mgr, &networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment{}),
    }
}

func (c *wasmDeploymentEventWatcher) AddEventHandler(ctx context.Context, h WasmDeploymentEventHandler, predicates ...predicate.Predicate) error {
	handler := genericWasmDeploymentHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericWasmDeploymentHandler implements a generic events.EventHandler
type genericWasmDeploymentHandler struct {
    handler WasmDeploymentEventHandler
}

func (h genericWasmDeploymentHandler) Create(object client.Object) error {
    obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
    if !ok {
        return errors.Errorf("internal error: WasmDeployment handler received event for %T", object)
    }
    return h.handler.CreateWasmDeployment(obj)
}

func (h genericWasmDeploymentHandler) Delete(object client.Object) error {
    obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
    if !ok {
        return errors.Errorf("internal error: WasmDeployment handler received event for %T", object)
    }
    return h.handler.DeleteWasmDeployment(obj)
}

func (h genericWasmDeploymentHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
    if !ok {
        return errors.Errorf("internal error: WasmDeployment handler received event for %T", old)
    }
    objNew, ok := new.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
    if !ok {
        return errors.Errorf("internal error: WasmDeployment handler received event for %T", new)
    }
    return h.handler.UpdateWasmDeployment(objOld, objNew)
}

func (h genericWasmDeploymentHandler) Generic(object client.Object) error {
    obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
    if !ok {
        return errors.Errorf("internal error: WasmDeployment handler received event for %T", object)
    }
    return h.handler.GenericWasmDeployment(obj)
}

// Handle events for the VirtualDestination Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualDestinationEventHandler interface {
    CreateVirtualDestination(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error
    UpdateVirtualDestination(old, new *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error
    DeleteVirtualDestination(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error
    GenericVirtualDestination(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error
}

type VirtualDestinationEventHandlerFuncs struct {
    OnCreate  func(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error
    OnUpdate  func(old, new *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error
    OnDelete  func(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error
    OnGeneric func(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error
}

func (f *VirtualDestinationEventHandlerFuncs) CreateVirtualDestination(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *VirtualDestinationEventHandlerFuncs) DeleteVirtualDestination(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *VirtualDestinationEventHandlerFuncs) UpdateVirtualDestination(objOld, objNew *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *VirtualDestinationEventHandlerFuncs) GenericVirtualDestination(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type VirtualDestinationEventWatcher interface {
    AddEventHandler(ctx context.Context, h VirtualDestinationEventHandler, predicates ...predicate.Predicate) error
}

type virtualDestinationEventWatcher struct {
    watcher events.EventWatcher
}

func NewVirtualDestinationEventWatcher(name string, mgr manager.Manager) VirtualDestinationEventWatcher {
    return &virtualDestinationEventWatcher{
        watcher: events.NewWatcher(name, mgr, &networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination{}),
    }
}

func (c *virtualDestinationEventWatcher) AddEventHandler(ctx context.Context, h VirtualDestinationEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualDestinationHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericVirtualDestinationHandler implements a generic events.EventHandler
type genericVirtualDestinationHandler struct {
    handler VirtualDestinationEventHandler
}

func (h genericVirtualDestinationHandler) Create(object client.Object) error {
    obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
    if !ok {
        return errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
    }
    return h.handler.CreateVirtualDestination(obj)
}

func (h genericVirtualDestinationHandler) Delete(object client.Object) error {
    obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
    if !ok {
        return errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
    }
    return h.handler.DeleteVirtualDestination(obj)
}

func (h genericVirtualDestinationHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
    if !ok {
        return errors.Errorf("internal error: VirtualDestination handler received event for %T", old)
    }
    objNew, ok := new.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
    if !ok {
        return errors.Errorf("internal error: VirtualDestination handler received event for %T", new)
    }
    return h.handler.UpdateVirtualDestination(objOld, objNew)
}

func (h genericVirtualDestinationHandler) Generic(object client.Object) error {
    obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
    if !ok {
        return errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
    }
    return h.handler.GenericVirtualDestination(obj)
}
