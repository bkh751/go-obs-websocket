package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// GetStreamingStatusRequest : Get current streaming and recording status.
// Since obs-websocket version: 0.3.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstreamingstatus
type GetStreamingStatusRequest _request

// NewGetStreamingStatusRequest returns a new GetStreamingStatusRequest.
func NewGetStreamingStatusRequest() GetStreamingStatusRequest {
	return GetStreamingStatusRequest{MessageID: getMessageID(), RequestType: "GetStreamingStatus"}
}

// ID returns the request's message ID.
func (r GetStreamingStatusRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r GetStreamingStatusRequest) Type() string { return r.RequestType }

// GetStreamingStatusResponse : Response for GetStreamingStatusRequest.
// Since obs-websocket version: 0.3.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstreamingstatus
type GetStreamingStatusResponse struct {
	// Current streaming status.
	// Required: Yes.
	Streaming bool `json:"streaming"`
	// Current recording status.
	// Required: Yes.
	Recording bool `json:"recording"`
	// Time elapsed since streaming started (only present if currently streaming).
	// Required: No.
	StreamTimecode string `json:"stream-timecode"`
	// Time elapsed since recording started (only present if currently recording).
	// Required: No.
	RecTimecode string `json:"rec-timecode"`
	// Always false.
	// Retrocompatibility with OBSRemote.
	// Required: Yes.
	PreviewOnly bool `json:"preview-only"`
	_response   `json:",squash"`
}

// ID returns the response's message ID.
func (r GetStreamingStatusResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r GetStreamingStatusResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r GetStreamingStatusResponse) Err() string { return r.Error }

// StartStopStreamingRequest : Toggle streaming on or off.
// Since obs-websocket version: 0.3.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstopstreaming
type StartStopStreamingRequest _request

// NewStartStopStreamingRequest returns a new StartStopStreamingRequest.
func NewStartStopStreamingRequest() StartStopStreamingRequest {
	return StartStopStreamingRequest{MessageID: getMessageID(), RequestType: "StartStopStreaming"}
}

// ID returns the request's message ID.
func (r StartStopStreamingRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r StartStopStreamingRequest) Type() string { return r.RequestType }

// StartStopStreamingResponse : Response for StartStopStreamingRequest.
// Since obs-websocket version: 0.3.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstopstreaming
type StartStopStreamingResponse _response

// ID returns the response's message ID.
func (r StartStopStreamingResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r StartStopStreamingResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r StartStopStreamingResponse) Err() string { return r.Error }

// StartStreamingRequest : Start streaming.
// Will return an `error` if streaming is already active.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstreaming
type StartStreamingRequest struct {
	// Special stream configuration.
	// Please note: these won't be saved to OBS' configuration.
	// Required: No.
	Stream map[string]interface{} `json:"stream"`
	// If specified ensures the type of stream matches the given type (usually 'rtmp_custom' or 'rtmp_common').
	// If the currently configured stream type does not match the given stream type, all settings must be specified in the `settings` object or an error will occur when starting the stream.
	// Required: No.
	StreamType string `json:"stream.type"`
	// Adds the given object parameters as encoded query string parameters to the 'key' of the RTMP stream.
	// Used to pass data to the RTMP service about the streaming.
	// May be any String, Numeric, or Boolean field.
	// Required: No.
	StreamMetadata map[string]interface{} `json:"stream.metadata"`
	// Settings for the stream.
	// Required: No.
	StreamSettings map[string]interface{} `json:"stream.settings"`
	// The publish URL.
	// Required: No.
	StreamSettingsServer string `json:"stream.settings.server"`
	// The publish key of the stream.
	// Required: No.
	StreamSettingsKey string `json:"stream.settings.key"`
	// Indicates whether authentication should be used when connecting to the streaming server.
	// Required: No.
	StreamSettingsUseAuth bool `json:"stream.settings.use-auth"`
	// If authentication is enabled, the username for the streaming server.
	// Ignored if `use-auth` is not set to `true`.
	// Required: No.
	StreamSettingsUsername string `json:"stream.settings.username"`
	// If authentication is enabled, the password for the streaming server.
	// Ignored if `use-auth` is not set to `true`.
	// Required: No.
	StreamSettingsPassword string `json:"stream.settings.password"`
	_request               `json:",squash"`
}

// NewStartStreamingRequest returns a new StartStreamingRequest.
func NewStartStreamingRequest(
	stream map[string]interface{},
	streamType string,
	streamMetadata map[string]interface{},
	streamSettings map[string]interface{},
	streamSettingsServer string,
	streamSettingsKey string,
	streamSettingsUseAuth bool,
	streamSettingsUsername string,
	streamSettingsPassword string,
) StartStreamingRequest {
	return StartStreamingRequest{
		stream,
		streamType,
		streamMetadata,
		streamSettings,
		streamSettingsServer,
		streamSettingsKey,
		streamSettingsUseAuth,
		streamSettingsUsername,
		streamSettingsPassword,
		_request{
			MessageID:   getMessageID(),
			RequestType: "StartStreaming",
		},
	}

}

// ID returns the request's message ID.
func (r StartStreamingRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r StartStreamingRequest) Type() string { return r.RequestType }

// StartStreamingResponse : Response for StartStreamingRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstreaming
type StartStreamingResponse _response

// ID returns the response's message ID.
func (r StartStreamingResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r StartStreamingResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r StartStreamingResponse) Err() string { return r.Error }

// StopStreamingRequest : Stop streaming.
// Will return an `error` if streaming is not active.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#stopstreaming
type StopStreamingRequest _request

// NewStopStreamingRequest returns a new StopStreamingRequest.
func NewStopStreamingRequest() StopStreamingRequest {
	return StopStreamingRequest{MessageID: getMessageID(), RequestType: "StopStreaming"}
}

// ID returns the request's message ID.
func (r StopStreamingRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r StopStreamingRequest) Type() string { return r.RequestType }

// StopStreamingResponse : Response for StopStreamingRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#stopstreaming
type StopStreamingResponse _response

// ID returns the response's message ID.
func (r StopStreamingResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r StopStreamingResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r StopStreamingResponse) Err() string { return r.Error }

// SetStreamSettingsRequest : Sets one or more attributes of the current streaming server settings
// Any options not passed will remain unchanged
// Returns the updated settings in response
// If 'type' is different than the current streaming service type, all settings are required
// Returns the full settings of the stream (the same as GetStreamSettings).
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setstreamsettings
type SetStreamSettingsRequest struct {
	// The type of streaming service configuration, usually `rtmp_custom` or `rtmp_common`.
	// Required: Yes.
	StreamType string `json:"type"`
	// The actual settings of the stream.
	// Required: Yes.
	Settings map[string]interface{} `json:"settings"`
	// The publish URL.
	// Required: No.
	SettingsServer string `json:"settings.server"`
	// The publish key.
	// Required: No.
	SettingsKey string `json:"settings.key"`
	// Indicates whether authentication should be used when connecting to the streaming server.
	// Required: No.
	SettingsUseAuth bool `json:"settings.use-auth"`
	// The username for the streaming service.
	// Required: No.
	SettingsUsername string `json:"settings.username"`
	// The password for the streaming service.
	// Required: No.
	SettingsPassword string `json:"settings.password"`
	// Persist the settings to disk.
	// Required: Yes.
	Save     bool `json:"save"`
	_request `json:",squash"`
}

// NewSetStreamSettingsRequest returns a new SetStreamSettingsRequest.
func NewSetStreamSettingsRequest(
	_type string,
	settings map[string]interface{},
	settingsServer string,
	settingsKey string,
	settingsUseAuth bool,
	settingsUsername string,
	settingsPassword string,
	save bool,
) SetStreamSettingsRequest {
	return SetStreamSettingsRequest{
		_type,
		settings,
		settingsServer,
		settingsKey,
		settingsUseAuth,
		settingsUsername,
		settingsPassword,
		save,
		_request{
			MessageID:   getMessageID(),
			RequestType: "SetStreamSettings",
		},
	}

}

// ID returns the request's message ID.
func (r SetStreamSettingsRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetStreamSettingsRequest) Type() string { return r.RequestType }

// SetStreamSettingsResponse : Response for SetStreamSettingsRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setstreamsettings
type SetStreamSettingsResponse _response

// ID returns the response's message ID.
func (r SetStreamSettingsResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetStreamSettingsResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetStreamSettingsResponse) Err() string { return r.Error }

// GetStreamSettingsRequest : Get the current streaming server settings.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstreamsettings
type GetStreamSettingsRequest _request

// NewGetStreamSettingsRequest returns a new GetStreamSettingsRequest.
func NewGetStreamSettingsRequest() GetStreamSettingsRequest {
	return GetStreamSettingsRequest{MessageID: getMessageID(), RequestType: "GetStreamSettings"}
}

// ID returns the request's message ID.
func (r GetStreamSettingsRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r GetStreamSettingsRequest) Type() string { return r.RequestType }

// GetStreamSettingsResponse : Response for GetStreamSettingsRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstreamsettings
type GetStreamSettingsResponse struct {
	// The type of streaming service configuration.
	// Possible values: 'rtmp_custom' or 'rtmp_common'.
	// Required: Yes.
	Type string `json:"type"`
	// Stream settings object.
	// Required: Yes.
	Settings map[string]interface{} `json:"settings"`
	// The publish URL.
	// Required: Yes.
	SettingsServer string `json:"settings.server"`
	// The publish key of the stream.
	// Required: Yes.
	SettingsKey string `json:"settings.key"`
	// Indicates whether audentication should be used when connecting to the streaming server.
	// Required: Yes.
	SettingsUseAuth bool `json:"settings.use-auth"`
	// The username to use when accessing the streaming server.
	// Only present if `use-auth` is `true`.
	// Required: Yes.
	SettingsUsername string `json:"settings.username"`
	// The password to use when accessing the streaming server.
	// Only present if `use-auth` is `true`.
	// Required: Yes.
	SettingsPassword string `json:"settings.password"`
	_response        `json:",squash"`
}

// ID returns the response's message ID.
func (r GetStreamSettingsResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r GetStreamSettingsResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r GetStreamSettingsResponse) Err() string { return r.Error }

// SaveStreamSettingsRequest : Save the current streaming server settings to disk.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#savestreamsettings
type SaveStreamSettingsRequest _request

// NewSaveStreamSettingsRequest returns a new SaveStreamSettingsRequest.
func NewSaveStreamSettingsRequest() SaveStreamSettingsRequest {
	return SaveStreamSettingsRequest{MessageID: getMessageID(), RequestType: "SaveStreamSettings"}
}

// ID returns the request's message ID.
func (r SaveStreamSettingsRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SaveStreamSettingsRequest) Type() string { return r.RequestType }

// SaveStreamSettingsResponse : Response for SaveStreamSettingsRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#savestreamsettings
type SaveStreamSettingsResponse _response

// ID returns the response's message ID.
func (r SaveStreamSettingsResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SaveStreamSettingsResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SaveStreamSettingsResponse) Err() string { return r.Error }
