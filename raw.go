// Copyright 2026 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tcell

import "time"

// EventRaw is emitted when the input processor recognises a raw escape
// sequence that was registered via Screen.RegisterRawSeq. It carries the
// matched sequence verbatim so callers can key off the exact bytes.
//
// Applications can register arbitrary escape sequences on Screen and
// react to terminal-specific or extension key codes that tcell's built-in
// parser does not decode on its own.
type EventRaw struct {
	t   time.Time
	esc string
}

// When returns the time when this event was created.
func (ev *EventRaw) When() time.Time {
	return ev.t
}

// EscSeq returns the raw escape sequence bytes that produced this event.
func (ev *EventRaw) EscSeq() string {
	return ev.esc
}

// NewEventRaw returns a new EventRaw carrying the given escape sequence.
func NewEventRaw(code string) *EventRaw {
	return &EventRaw{t: time.Now(), esc: code}
}
