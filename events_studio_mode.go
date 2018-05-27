package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// PreviewSceneChangedEvent : The selected preview scene has changed (only available in Studio Mode). Since: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#previewscenechanged
type PreviewSceneChangedEvent struct {
	SceneName string      `json:"scene-name"` // Name of the scene being previewed.
	Sources   interface{} `json:"sources"`    // List of sources composing the scene. Same specification as [`GetCurrentScene`](#getcurrentscene). TODO: Unknown type (Source|Array).
	_event
}

// Type returns the event's update type.
func (e PreviewSceneChangedEvent) Type() string { return e.UpdateType }

// StreamTC returns the event's stream timecode.
func (e PreviewSceneChangedEvent) StreamTC() string { return e.StreamTimecode }

// RecTC returns the event's recording timecode.
func (e PreviewSceneChangedEvent) RecTC() string { return e.RecTimecode }

// StudioModeSwitchedEvent : Studio Mode has been enabled or disabled. Since: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#studiomodeswitched
type StudioModeSwitchedEvent struct {
	NewState bool `json:"new-state"` // The new enabled state of Studio Mode.
	_event
}

// Type returns the event's update type.
func (e StudioModeSwitchedEvent) Type() string { return e.UpdateType }

// StreamTC returns the event's stream timecode.
func (e StudioModeSwitchedEvent) StreamTC() string { return e.StreamTimecode }

// RecTC returns the event's recording timecode.
func (e StudioModeSwitchedEvent) RecTC() string { return e.RecTimecode }