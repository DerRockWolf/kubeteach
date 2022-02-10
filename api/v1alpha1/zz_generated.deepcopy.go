//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 Maximilian Geberl.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExerciseSet) DeepCopyInto(out *ExerciseSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExerciseSet.
func (in *ExerciseSet) DeepCopy() *ExerciseSet {
	if in == nil {
		return nil
	}
	out := new(ExerciseSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExerciseSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExerciseSetList) DeepCopyInto(out *ExerciseSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExerciseSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExerciseSetList.
func (in *ExerciseSetList) DeepCopy() *ExerciseSetList {
	if in == nil {
		return nil
	}
	out := new(ExerciseSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExerciseSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExerciseSetSpec) DeepCopyInto(out *ExerciseSetSpec) {
	*out = *in
	if in.TaskDefinitions != nil {
		in, out := &in.TaskDefinitions, &out.TaskDefinitions
		*out = make([]ExerciseSetSpecTaskDefinitions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExerciseSetSpec.
func (in *ExerciseSetSpec) DeepCopy() *ExerciseSetSpec {
	if in == nil {
		return nil
	}
	out := new(ExerciseSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExerciseSetSpecTaskDefinitions) DeepCopyInto(out *ExerciseSetSpecTaskDefinitions) {
	*out = *in
	in.TaskDefinitionSpec.DeepCopyInto(&out.TaskDefinitionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExerciseSetSpecTaskDefinitions.
func (in *ExerciseSetSpecTaskDefinitions) DeepCopy() *ExerciseSetSpecTaskDefinitions {
	if in == nil {
		return nil
	}
	out := new(ExerciseSetSpecTaskDefinitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExerciseSetStatus) DeepCopyInto(out *ExerciseSetStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExerciseSetStatus.
func (in *ExerciseSetStatus) DeepCopy() *ExerciseSetStatus {
	if in == nil {
		return nil
	}
	out := new(ExerciseSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceCondition) DeepCopyInto(out *ResourceCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceCondition.
func (in *ResourceCondition) DeepCopy() *ResourceCondition {
	if in == nil {
		return nil
	}
	out := new(ResourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Task) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskCondition) DeepCopyInto(out *TaskCondition) {
	*out = *in
	if in.ResourceCondition != nil {
		in, out := &in.ResourceCondition, &out.ResourceCondition
		*out = make([]ResourceCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskCondition.
func (in *TaskCondition) DeepCopy() *TaskCondition {
	if in == nil {
		return nil
	}
	out := new(TaskCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskDefinition) DeepCopyInto(out *TaskDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskDefinition.
func (in *TaskDefinition) DeepCopy() *TaskDefinition {
	if in == nil {
		return nil
	}
	out := new(TaskDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskDefinitionList) DeepCopyInto(out *TaskDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TaskDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskDefinitionList.
func (in *TaskDefinitionList) DeepCopy() *TaskDefinitionList {
	if in == nil {
		return nil
	}
	out := new(TaskDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskDefinitionSpec) DeepCopyInto(out *TaskDefinitionSpec) {
	*out = *in
	out.TaskSpec = in.TaskSpec
	if in.TaskConditions != nil {
		in, out := &in.TaskConditions, &out.TaskConditions
		*out = make([]TaskCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequiredTaskName != nil {
		in, out := &in.RequiredTaskName, &out.RequiredTaskName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskDefinitionSpec.
func (in *TaskDefinitionSpec) DeepCopy() *TaskDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(TaskDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskDefinitionStatus) DeepCopyInto(out *TaskDefinitionStatus) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskDefinitionStatus.
func (in *TaskDefinitionStatus) DeepCopy() *TaskDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(TaskDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskList) DeepCopyInto(out *TaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Task, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskList.
func (in *TaskList) DeepCopy() *TaskList {
	if in == nil {
		return nil
	}
	out := new(TaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}
