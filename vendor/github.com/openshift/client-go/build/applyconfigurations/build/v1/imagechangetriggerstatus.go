// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ImageChangeTriggerStatusApplyConfiguration represents an declarative configuration of the ImageChangeTriggerStatus type for use
// with apply.
type ImageChangeTriggerStatusApplyConfiguration struct {
	LastTriggeredImageID *string                                    `json:"lastTriggeredImageID,omitempty"`
	From                 *ImageStreamTagReferenceApplyConfiguration `json:"from,omitempty"`
	LastTriggerTime      *metav1.Time                               `json:"lastTriggerTime,omitempty"`
}

// ImageChangeTriggerStatusApplyConfiguration constructs an declarative configuration of the ImageChangeTriggerStatus type for use with
// apply.
func ImageChangeTriggerStatus() *ImageChangeTriggerStatusApplyConfiguration {
	return &ImageChangeTriggerStatusApplyConfiguration{}
}

// WithLastTriggeredImageID sets the LastTriggeredImageID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTriggeredImageID field is set to the value of the last call.
func (b *ImageChangeTriggerStatusApplyConfiguration) WithLastTriggeredImageID(value string) *ImageChangeTriggerStatusApplyConfiguration {
	b.LastTriggeredImageID = &value
	return b
}

// WithFrom sets the From field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the From field is set to the value of the last call.
func (b *ImageChangeTriggerStatusApplyConfiguration) WithFrom(value *ImageStreamTagReferenceApplyConfiguration) *ImageChangeTriggerStatusApplyConfiguration {
	b.From = value
	return b
}

// WithLastTriggerTime sets the LastTriggerTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTriggerTime field is set to the value of the last call.
func (b *ImageChangeTriggerStatusApplyConfiguration) WithLastTriggerTime(value metav1.Time) *ImageChangeTriggerStatusApplyConfiguration {
	b.LastTriggerTime = &value
	return b
}