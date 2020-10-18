package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// GetStreamingStatusRequest : Get current streaming and recording status.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getstreamingstatus
type GetStreamingStatusRequest struct {
	_request `json:",squash"`
	response chan GetStreamingStatusResponse
}

// NewGetStreamingStatusRequest returns a new GetStreamingStatusRequest.
func NewGetStreamingStatusRequest() GetStreamingStatusRequest {
	return GetStreamingStatusRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "GetStreamingStatus",
			err:   make(chan error, 1),
		},
		make(chan GetStreamingStatusResponse, 1),
	}
}

// Send sends the request.
func (r *GetStreamingStatusRequest) Send(c *Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetStreamingStatusResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetStreamingStatusRequest) Receive() (GetStreamingStatusResponse, error) {
	if !r.sent {
		return GetStreamingStatusResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetStreamingStatusResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetStreamingStatusResponse{}, err
		case <-time.After(receiveTimeout):
			return GetStreamingStatusResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetStreamingStatusRequest) SendReceive(c *Client) (GetStreamingStatusResponse, error) {
	if err := r.Send(c); err != nil {
		return GetStreamingStatusResponse{}, err
	}
	return r.Receive()
}

// GetStreamingStatusResponse : Response for GetStreamingStatusRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getstreamingstatus
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

// StartStopStreamingRequest : Toggle streaming on or off.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startstopstreaming
type StartStopStreamingRequest struct {
	_request `json:",squash"`
	response chan StartStopStreamingResponse
}

// NewStartStopStreamingRequest returns a new StartStopStreamingRequest.
func NewStartStopStreamingRequest() StartStopStreamingRequest {
	return StartStopStreamingRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "StartStopStreaming",
			err:   make(chan error, 1),
		},
		make(chan StartStopStreamingResponse, 1),
	}
}

// Send sends the request.
func (r *StartStopStreamingRequest) Send(c *Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp StartStopStreamingResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r StartStopStreamingRequest) Receive() (StartStopStreamingResponse, error) {
	if !r.sent {
		return StartStopStreamingResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartStopStreamingResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartStopStreamingResponse{}, err
		case <-time.After(receiveTimeout):
			return StartStopStreamingResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StartStopStreamingRequest) SendReceive(c *Client) (StartStopStreamingResponse, error) {
	if err := r.Send(c); err != nil {
		return StartStopStreamingResponse{}, err
	}
	return r.Receive()
}

// StartStopStreamingResponse : Response for StartStopStreamingRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startstopstreaming
type StartStopStreamingResponse struct {
	_response `json:",squash"`
}

// StartStreamingRequest : Start streaming.
// Will return an `error` if streaming is already active.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startstreaming
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
	response               chan StartStreamingResponse
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
			ID_:   GetMessageID(),
			Type_: "StartStreaming",
			err:   make(chan error, 1),
		},
		make(chan StartStreamingResponse, 1),
	}
}

// Send sends the request.
func (r *StartStreamingRequest) Send(c *Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp StartStreamingResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r StartStreamingRequest) Receive() (StartStreamingResponse, error) {
	if !r.sent {
		return StartStreamingResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartStreamingResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartStreamingResponse{}, err
		case <-time.After(receiveTimeout):
			return StartStreamingResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StartStreamingRequest) SendReceive(c *Client) (StartStreamingResponse, error) {
	if err := r.Send(c); err != nil {
		return StartStreamingResponse{}, err
	}
	return r.Receive()
}

// StartStreamingResponse : Response for StartStreamingRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startstreaming
type StartStreamingResponse struct {
	_response `json:",squash"`
}

// StopStreamingRequest : Stop streaming.
// Will return an `error` if streaming is not active.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stopstreaming
type StopStreamingRequest struct {
	_request `json:",squash"`
	response chan StopStreamingResponse
}

// NewStopStreamingRequest returns a new StopStreamingRequest.
func NewStopStreamingRequest() StopStreamingRequest {
	return StopStreamingRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "StopStreaming",
			err:   make(chan error, 1),
		},
		make(chan StopStreamingResponse, 1),
	}
}

// Send sends the request.
func (r *StopStreamingRequest) Send(c *Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp StopStreamingResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r StopStreamingRequest) Receive() (StopStreamingResponse, error) {
	if !r.sent {
		return StopStreamingResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopStreamingResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopStreamingResponse{}, err
		case <-time.After(receiveTimeout):
			return StopStreamingResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StopStreamingRequest) SendReceive(c *Client) (StopStreamingResponse, error) {
	if err := r.Send(c); err != nil {
		return StopStreamingResponse{}, err
	}
	return r.Receive()
}

// StopStreamingResponse : Response for StopStreamingRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stopstreaming
type StopStreamingResponse struct {
	_response `json:",squash"`
}

// SetStreamSettingsRequest : Sets one or more attributes of the current streaming server settings
// Any options not passed will remain unchanged
// Returns the updated settings in response
// If 'type' is different than the current streaming service type, all settings are required
// Returns the full settings of the stream (the same as GetStreamSettings).
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setstreamsettings
type SetStreamSettingsRequest struct {
	// The type of streaming service configuration, usually `rtmp_custom` or `rtmp_common`.
	// Required: Yes.
	Type_ string `json:"type"`
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
	response chan SetStreamSettingsResponse
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
			ID_:   GetMessageID(),
			Type_: "SetStreamSettings",
			err:   make(chan error, 1),
		},
		make(chan SetStreamSettingsResponse, 1),
	}
}

// Send sends the request.
func (r *SetStreamSettingsRequest) Send(c *Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp SetStreamSettingsResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r SetStreamSettingsRequest) Receive() (SetStreamSettingsResponse, error) {
	if !r.sent {
		return SetStreamSettingsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetStreamSettingsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetStreamSettingsResponse{}, err
		case <-time.After(receiveTimeout):
			return SetStreamSettingsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetStreamSettingsRequest) SendReceive(c *Client) (SetStreamSettingsResponse, error) {
	if err := r.Send(c); err != nil {
		return SetStreamSettingsResponse{}, err
	}
	return r.Receive()
}

// SetStreamSettingsResponse : Response for SetStreamSettingsRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setstreamsettings
type SetStreamSettingsResponse struct {
	_response `json:",squash"`
}

// GetStreamSettingsRequest : Get the current streaming server settings.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getstreamsettings
type GetStreamSettingsRequest struct {
	_request `json:",squash"`
	response chan GetStreamSettingsResponse
}

// NewGetStreamSettingsRequest returns a new GetStreamSettingsRequest.
func NewGetStreamSettingsRequest() GetStreamSettingsRequest {
	return GetStreamSettingsRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "GetStreamSettings",
			err:   make(chan error, 1),
		},
		make(chan GetStreamSettingsResponse, 1),
	}
}

// Send sends the request.
func (r *GetStreamSettingsRequest) Send(c *Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetStreamSettingsResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetStreamSettingsRequest) Receive() (GetStreamSettingsResponse, error) {
	if !r.sent {
		return GetStreamSettingsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetStreamSettingsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetStreamSettingsResponse{}, err
		case <-time.After(receiveTimeout):
			return GetStreamSettingsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetStreamSettingsRequest) SendReceive(c *Client) (GetStreamSettingsResponse, error) {
	if err := r.Send(c); err != nil {
		return GetStreamSettingsResponse{}, err
	}
	return r.Receive()
}

// GetStreamSettingsResponse : Response for GetStreamSettingsRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getstreamsettings
type GetStreamSettingsResponse struct {
	// The type of streaming service configuration.
	// Possible values: 'rtmp_custom' or 'rtmp_common'.
	// Required: Yes.
	Type_ string `json:"type"`
	// Stream settings object.
	// Required: Yes.
	Settings map[string]interface{} `json:"settings"`
	// The publish URL.
	// Required: Yes.
	SettingsServer string `json:"settings.server"`
	// The publish key of the stream.
	// Required: Yes.
	SettingsKey string `json:"settings.key"`
	// Indicates whether authentication should be used when connecting to the streaming server.
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

// SaveStreamSettingsRequest : Save the current streaming server settings to disk.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#savestreamsettings
type SaveStreamSettingsRequest struct {
	_request `json:",squash"`
	response chan SaveStreamSettingsResponse
}

// NewSaveStreamSettingsRequest returns a new SaveStreamSettingsRequest.
func NewSaveStreamSettingsRequest() SaveStreamSettingsRequest {
	return SaveStreamSettingsRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "SaveStreamSettings",
			err:   make(chan error, 1),
		},
		make(chan SaveStreamSettingsResponse, 1),
	}
}

// Send sends the request.
func (r *SaveStreamSettingsRequest) Send(c *Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp SaveStreamSettingsResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r SaveStreamSettingsRequest) Receive() (SaveStreamSettingsResponse, error) {
	if !r.sent {
		return SaveStreamSettingsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SaveStreamSettingsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SaveStreamSettingsResponse{}, err
		case <-time.After(receiveTimeout):
			return SaveStreamSettingsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SaveStreamSettingsRequest) SendReceive(c *Client) (SaveStreamSettingsResponse, error) {
	if err := r.Send(c); err != nil {
		return SaveStreamSettingsResponse{}, err
	}
	return r.Receive()
}

// SaveStreamSettingsResponse : Response for SaveStreamSettingsRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#savestreamsettings
type SaveStreamSettingsResponse struct {
	_response `json:",squash"`
}

// SendCaptionsRequest : Send the provided text as embedded CEA-608 caption data.
// As of OBS Studio 23.1, captions are not yet available on Linux.
//
// Since obs-websocket version: 4.6.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sendcaptions
type SendCaptionsRequest struct {
	// Captions text.
	// Required: Yes.
	Text     string `json:"text"`
	_request `json:",squash"`
	response chan SendCaptionsResponse
}

// NewSendCaptionsRequest returns a new SendCaptionsRequest.
func NewSendCaptionsRequest(text string) SendCaptionsRequest {
	return SendCaptionsRequest{
		text,
		_request{
			ID_:   GetMessageID(),
			Type_: "SendCaptions",
			err:   make(chan error, 1),
		},
		make(chan SendCaptionsResponse, 1),
	}
}

// Send sends the request.
func (r *SendCaptionsRequest) Send(c *Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp SendCaptionsResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r SendCaptionsRequest) Receive() (SendCaptionsResponse, error) {
	if !r.sent {
		return SendCaptionsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SendCaptionsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SendCaptionsResponse{}, err
		case <-time.After(receiveTimeout):
			return SendCaptionsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SendCaptionsRequest) SendReceive(c *Client) (SendCaptionsResponse, error) {
	if err := r.Send(c); err != nil {
		return SendCaptionsResponse{}, err
	}
	return r.Receive()
}

// SendCaptionsResponse : Response for SendCaptionsRequest.
//
// Since obs-websocket version: 4.6.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sendcaptions
type SendCaptionsResponse struct {
	_response `json:",squash"`
}
