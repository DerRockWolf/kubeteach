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

package controllers

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kubeteachv1alpha1 "github.com/dergeberl/kubeteach/api/v1alpha1"
)

// ExerciseSetReconciler reconciles a ExerciseSet object
type ExerciseSetReconciler struct {
	client.Client
	Log         logr.Logger
	Scheme      *runtime.Scheme
	RequeueTime time.Duration
}

//+kubebuilder:rbac:groups=kubeteach.geberl.io,resources=taskdefinitions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubeteach.geberl.io,resources=exercisesets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubeteach.geberl.io,resources=exercisesets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kubeteach.geberl.io,resources=exercisesets/finalizers,verbs=update

// Reconcile handles reconcile of an ExersiceSet
func (r *ExerciseSetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("exerciseset", req.NamespacedName)

	var exerciseSet kubeteachv1alpha1.ExerciseSet
	err := r.Client.Get(ctx, req.NamespacedName, &exerciseSet)
	if err != nil {
		// ignore ExerciseSet that dose not exists
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var newExerciseSetStatus kubeteachv1alpha1.ExerciseSetStatus

	for _, taskDefinition := range exerciseSet.Spec.TaskDefinitions {
		var taskDefinitionObject kubeteachv1alpha1.TaskDefinition
		err = r.Client.Get(ctx, client.ObjectKey{Name: taskDefinition.Name, Namespace: req.Namespace}, &taskDefinitionObject)
		if err != nil {
			if client.IgnoreNotFound(err) != nil {
				return ctrl.Result{}, err
			}
			// create taskDefinition
			taskDefinitionObject = kubeteachv1alpha1.TaskDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name:      taskDefinition.Name,
					Namespace: req.Namespace,
					OwnerReferences: []metav1.OwnerReference{{
						APIVersion: exerciseSet.APIVersion,
						Kind:       exerciseSet.Kind,
						Name:       exerciseSet.Name,
						UID:        exerciseSet.UID,
					}},
				},
				Spec: taskDefinition.TaskDefinitionSpec,
			}
			err = r.Client.Create(ctx, &taskDefinitionObject)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		// update TaskDefinition if needed
		if !reflect.DeepEqual(taskDefinitionObject.Spec, taskDefinition.TaskDefinitionSpec) {
			taskDefinitionObject.Spec = taskDefinition.TaskDefinitionSpec
			err = r.Client.Update(ctx, &taskDefinitionObject)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		// update OwnerReferences if needed
		if !reflect.DeepEqual(taskDefinitionObject.OwnerReferences, []metav1.OwnerReference{{
			APIVersion: exerciseSet.APIVersion,
			Kind:       exerciseSet.Kind,
			Name:       exerciseSet.Name,
			UID:        exerciseSet.UID,
		}}) {
			taskDefinitionObject.OwnerReferences = []metav1.OwnerReference{{
				APIVersion: exerciseSet.APIVersion,
				Kind:       exerciseSet.Kind,
				Name:       exerciseSet.Name,
				UID:        exerciseSet.UID,
			}}
			err = r.Client.Update(ctx, &taskDefinitionObject)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		// count total tasks
		newExerciseSetStatus.NumberOfTasks++

		// count tasks with state
		if taskDefinitionObject.Status.State != nil {
			switch *taskDefinitionObject.Status.State {
			case StateActive:
				newExerciseSetStatus.NumberOfActiveTasks++
			case StatePending:
				newExerciseSetStatus.NumberOfPendingTasks++
			case StateSuccessful:
				newExerciseSetStatus.NumberOfSuccessfulTasks++
			}
		} else {
			newExerciseSetStatus.NumberOfUnknownTasks++
		}

		// count total sum of points
		newExerciseSetStatus.PointsTotal += taskDefinition.TaskDefinitionSpec.Points

		// count tasks without points
		if taskDefinition.TaskDefinitionSpec.Points == 0 {
			newExerciseSetStatus.NumberOfTasksWithoutPoints++
		}

		// count points from successful tasks
		if taskDefinitionObject.Status.State != nil &&
			*taskDefinitionObject.Status.State == StateSuccessful {
			newExerciseSetStatus.PointsAchieved += taskDefinition.TaskDefinitionSpec.Points
		}
	}

	// update status if needed
	if !reflect.DeepEqual(exerciseSet.Status, newExerciseSetStatus) {
		patch := []byte(`{"status": {` +
			`"numberOfTasks": ` + fmt.Sprint(newExerciseSetStatus.NumberOfTasks) + `, ` +
			`"numberOfActiveTasks": ` + fmt.Sprint(newExerciseSetStatus.NumberOfActiveTasks) + `, ` +
			`"numberOfPendingTasks": ` + fmt.Sprint(newExerciseSetStatus.NumberOfPendingTasks) + `, ` +
			`"numberOfSuccessfulTasks": ` + fmt.Sprint(newExerciseSetStatus.NumberOfSuccessfulTasks) + `, ` +
			`"numberOfUnknownTasks": ` + fmt.Sprint(newExerciseSetStatus.NumberOfUnknownTasks) + `, ` +
			`"numberOfTasksWithoutPoints": ` + fmt.Sprint(newExerciseSetStatus.NumberOfTasksWithoutPoints) + `, ` +
			`"pointsTotal": ` + fmt.Sprint(newExerciseSetStatus.PointsTotal) + `, ` +
			`"pointsAchieved": ` + fmt.Sprint(newExerciseSetStatus.PointsAchieved) + `}}`)
		err = r.Client.Status().Patch(ctx, &exerciseSet, client.RawPatch(types.MergePatchType, patch))
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{RequeueAfter: r.RequeueTime}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExerciseSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubeteachv1alpha1.ExerciseSet{}).
		Complete(r)
}
