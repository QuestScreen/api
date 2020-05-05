package module

import (
	"time"

	"github.com/QuestScreen/api/common"
	"github.com/QuestScreen/api/group"
	"github.com/QuestScreen/api/render"
	"github.com/QuestScreen/api/resources"
	"github.com/QuestScreen/api/server"
	"gopkg.in/yaml.v3"
)

// the interfaces declared in this file are to be implemented by
// a QuestScreen module, which is provided by a plugin.

// PureEndpoint is an endpoint of a module for the HTTP server.
// It takes POST requests on the path specified by Module.
type PureEndpoint interface {
	// Post handles a POST request. Two return values are expected:
	//
	// The first return value will be serialized to JSON and sent back to the
	// client. If it's nil, nothing will be sent and the client will get a
	// 204 No Content.
	//
	// The second return value will be handed over to the module's InitTransition
	// which will be called in the OpenGL thread after Post has returned.
	// For thread safety, that value should be constructed from scratch and not be
	// a pointer into the ModuleState object.
	//
	// If an error is returned, InitTransition will not be called and both return
	// values will be ignored. The server will the respond according to the cause
	// of the returned error.
	Post(payload []byte) (interface{}, interface{}, common.SendableError)
}

// IDEndpoint is an endpoint of a module for the HTTP server.
// It takes POST requests on the path specified by Module, with an
// additional URL path item interpreted as ID.
type IDEndpoint interface {
	// Post works analoguous to ModulePureEndpoint.Post, but gets the id from the
	// request URL path as additional parameter.
	Post(id string, payload []byte) (interface{}, interface{}, common.SendableError)
}

// State describes the state of a module. It is written to and loaded
// from a group's state.yaml.
//
// All funcs are expected to be called in the server thread.
type State interface {
	server.SerializableItem
	// CreateRendererData generates a data object that contains all required data
	// for the Renderer to rebuild its state. The returned data object will
	// be handed over to the renderer's RebuildState. For thread safety, it should
	// not be a pointer into the ModuleState object.
	CreateRendererData() interface{}
}

// HeroAwareState is an interface that must be implemented by module
// states if they work with heroes. It lets the application send messages to the
// state when the list of heroes changes.
type HeroAwareState interface {
	HeroListChanged(heroes group.HeroList, action group.HeroChangeAction, heroIndex int)
}

// PureEndpointProvider is a State extension for modules whose
// Module defines one or more pure endpoints in its EndpointPaths.
type PureEndpointProvider interface {
	// PureEndpoint returns the pure endpoint defined at the given index of the
	// Module's EndpointPaths slice. This should be a cheap getter as it
	// will be called for every request on one of the module's pure endpoints.
	PureEndpoint(index int) PureEndpoint
}

// IDEndpointProvider is a State extension for modules whose
// Module defines one or more id endpoints in its EndpointPaths.
type IDEndpointProvider interface {
	// IDEndpoint returns the id endpoint defined at the given index of the
	// Module's EndpointPaths slice. This should be a cheap getter as it
	// will be called for every request on one of the module's id endpoints.
	IDEndpoint(index int) IDEndpoint
}

// Module describes a module that has a persistable state, a renderer that
// paints something to the display, and client-side HTML/JS to control the
// module.
type Module struct {
	// Name is the human-readable name of the module.
	Name string
	// ID is a unique string, used for identifying the module inside
	// HTTP URLs and in the filesystem. Therefore, the ID is restricted to ASCII
	// letters, digits, and the symbols `.,-_`
	ID string
	// ResourceCollections lists selectors for resource collections of this
	// module. The maximum resources.CollectionIndex available to this module is
	// len(ResourceCollections()) - 1.
	ResourceCollections []resources.Selector
	// EndpointPaths defines a list of API endpoints for the client to change this
	// module's state and trigger animations.
	//
	// The endpoints will be queryable at
	//
	//     /state/<module-id>/<endpoint-path>[/<entity-id>]
	//
	// If a path ends with a `/`, it will take the additional <entity-id>
	// parameter. At most one path may be empty, in which cause it will be
	// queryable at
	//
	//     /state/<module-id>
	//
	// At most one path may be `"/"`, in which case it will be queryable at
	//
	//     /state/<module-id>/<entity-id>
	//
	// If the `"/"` path exists, the only other path that may exist is the empty
	// path.
	//
	// If at least one path not ending with `/` exists, the module's state must
	// implement PureEndpointProvider, and if at least one path ending with `/`
	// exists, the module's state must implement IDEndpointProvider.
	EndpointPaths []string
	// DefaultConfig is a configuration object with default values.
	//
	// This value defines the type of this module's configuration. Its type must
	// be a pointer to a struct in which each field is a pointer to an item
	// implementing ConfigItem.
	//
	// Generally, a value of this type may have any of its fields set to nil,
	// meaning that it should inherit the value from a previous level. This is
	// for scene, group, system and base config (in that order). However, the
	// default config must only have non-nil values since it defines the fallback
	// if the whole path up from scene config to base config does not define any
	// value for a certain item.
	DefaultConfig interface{}
	// CreateRenderer creates the renderer object. This func will only be called
	// once during app initialization, making the renderer a singleton object.
	//
	// CreateRenderer should only initialize the bare minimum of the renderer's
	// data;
	// RebuildState will be issued to the renderer before the first Render() call
	// to fully initialize its state.
	CreateRenderer func(backend *render.Renderer,
		ms server.MessageSender) (Renderer, error)
	// CreateState will be called in the server thread. It shall create a
	// ModuleState for the module created by CreateModule.
	//
	// input will reflect the Persisted layout of the serialized state as
	// generated by the state's SerializableView method.
	// It may be nil in which case the state will be created with default values.
	//
	// Communication between ModuleState and Module will be done via the state's
	// HandleAction and CreateModuleData methods which create data, and the
	// module's InitTransition and RebuildState methods which consume that data.
	//
	// If the module accesses a group's heroes, its state must additionally
	// implement HeroAwareModuleState.
	CreateState func(input *yaml.Node, ctx server.Context,
		ms server.MessageSender) (State, error)
}

// Renderer describes the renderer of a module.
// This object belongs with the OpenGL thread.
type Renderer interface {
	// Descriptor shall return the Module this renderer belongs to.
	Descriptor() *Module
	// Rebuild will be called after any action that requires rebuilding the
	// renderer's state, such as a scene, config or group change or after startup.
	// For scene and group change, `data` contains data generated by the
	// ModuleState's CreateRendererData; it will be nil for pure config changes.
	//
	// `config` will always be non-nil since the config will also change for each
	// scene and group change, since the config will be recalculated. Its value
	// is the merged config object with the type of Module.DefaultConfig.
	// All values in the `config` object will be set to non-nil values.
	//
	// A call to RebuildState will always immediately be followed by a call to
	// Render.
	Rebuild(ctx render.Context, data interface{}, config interface{})
	// InitTransition will be called after the current ModuleState has been
	// modified via HandleAction.
	// data contains the data generated by HandleAction.
	//
	// The return value is the duration of the transition initiated by this call.
	// For that duration, the render thread will continuously call
	// TransitionStep and Render. After the time has passed,
	// FinishTransition and Render will be called to render the final state.
	//
	// if 0 is returned, TransitionStep will never be called; if a negative
	// value is returned, neither FinishTransition nor Render will be
	// called.
	InitTransition(ctx render.Context, data interface{}) time.Duration
	// TransitionStep should update the renderer's current state while
	// transitioning. A call to TransitionStep() will always immediately be
	// followed by a call to Render().
	//
	// The given elapsed time is guaranteed to always be smaller than what was
	// returned by InitTransition().
	TransitionStep(ctx render.Context, elapsed time.Duration)
	// FinishTransition() is for cleanup after a transition and for preparing the
	// final state. It will be called exactly once for each call to
	// InitTransition() that returned a non-negative value.
	//
	// A call to FinishTransition() will always immediately be followed by a call
	// to Render().
	FinishTransition(ctx render.Context)
	// Render renders the Module's current state.
	Render(ctx render.Context)
}
