/*
Copyright 2022 The RamenDR authors.

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
	"fmt"

	"k8s.io/apimachinery/pkg/types"

	ramendrv1alpha1 "github.com/ramendr/ramen/api/v1alpha1"
)

func (v *VRGInstance) kubeObjectsProtect() error {
	if v.instance.Spec.KubeObjectProtection == nil {
		v.log.Info("Kube object protection disabled")

		return nil
	}

	errors := make([]error, 0, len(v.instance.Spec.S3Profiles))

	for _, s3ProfileName := range v.instance.Spec.S3Profiles {
		// TODO reuse objectStore kube objects from pv upload
		objectStore, err := v.reconciler.ObjStoreGetter.ObjectStore(
			v.ctx,
			v.reconciler.APIReader,
			s3ProfileName,
			v.namespacedName,
			v.log,
		)
		if err != nil {
			v.log.Error(err, "kube objects protect object store access", "profile", s3ProfileName)
			errors = append(errors, err)

			continue
		}

		err = v.processSubBackups(objectStore, s3ProfileName)
		if err != nil {
			errors := append(errors, err)

			return errors[0]
		}
	}

	return nil
}

func (v *VRGInstance) processSubBackups(objectStore ObjectStorer, s3ProfileName string) error {
	categories := v.getTypeSequenceSubBackups()

	var err error // declare for scope

	for _, captureInstance := range categories {
		includeList, excludeList := getTypeSequenceResourceList(captureInstance.IncludedResources,
			captureInstance.ExcludedResources)

		spec := captureInstance.DeepCopy() // don't modify spec with processed results
		spec.IncludedResources = includeList
		spec.ExcludedResources = excludeList

		sourceNamespacedName := types.NamespacedName{Name: v.instance.Name, Namespace: v.instance.Namespace}

		if err := kubeObjectsProtect(
			v.ctx,
			v.reconciler.Client,
			v.reconciler.APIReader,
			v.log,
			objectStore.AddressComponent1(),
			objectStore.AddressComponent2(),
			v.s3KeyPrefix(),
			sourceNamespacedName,
			captureInstance,
		); err != nil {
			v.log.Error(err, "kube object protect", "profile", s3ProfileName)
			// don't return error yet since it could be non-fatal (e.g. slow API server)
		}

		backupNamespacedName := getBackupNamespacedName(sourceNamespacedName.Name,
			sourceNamespacedName.Namespace, spec.Name, getSequenceNumber())

		// TODO: remove tight loop here
		backupComplete := false
		for !backupComplete {
			backupComplete, err = backupIsDone(v.ctx, v.reconciler.APIReader,
				objectWriter{v.reconciler.Client, v.ctx, v.log}, backupNamespacedName)

			if err != nil {
				backupComplete = true
			}
		}
	}

	return err
}

// return values: includedResources, excludedResources
func getTypeSequenceResourceList(toInclude, toExclude []string) ([]string, []string) {
	included := []string{"*"} // include everything
	excluded := []string{}    // exclude nothing

	if toInclude != nil {
		included = toInclude
	}

	if toExclude != nil {
		excluded = toExclude
	}

	return included, excluded
}

func (v *VRGInstance) getTypeSequenceSubBackups() []ramendrv1alpha1.ResourceCaptureInstance {
	if v.instance.Spec.KubeObjectProtection != nil &&
		v.instance.Spec.KubeObjectProtection.ResourceCaptureOrder != nil {
		return v.instance.Spec.KubeObjectProtection.ResourceCaptureOrder
	}

	// default case: no backup groups defined
	return getTypeSequenceDefaultBackup()
}

func (v *VRGInstance) getTypeSequenceSubRestores() []ramendrv1alpha1.ResourceRestoreInstance {
	if v.instance.Spec.KubeObjectProtection != nil &&
		v.instance.Spec.KubeObjectProtection.ResourceCaptureOrder != nil {
		return v.instance.Spec.KubeObjectProtection.ResourceRecoveryOrder
	}

	// default case: no backup groups defined
	return getTypeSequenceDefaultRestore()
}

func getTypeSequenceDefaultBackup() []ramendrv1alpha1.ResourceCaptureInstance {
	instance := make([]ramendrv1alpha1.ResourceCaptureInstance, 1)

	includedResources := make([]string, 0)
	includedResources[0] = "*"

	instance[0].Name = "everything"
	instance[0].IncludeClusterScopedResources = false
	instance[0].IncludedResources = includedResources

	return instance
}

func getTypeSequenceDefaultRestore() []ramendrv1alpha1.ResourceRestoreInstance {
	instance := make([]ramendrv1alpha1.ResourceRestoreInstance, 1)

	includedResources := make([]string, 0)
	includedResources[0] = "*"

	instance[0].IncludeClusterScopedResources = true // TODO: needed for default restore case?
	instance[0].IncludedResources = includedResources

	return instance
}

func getBackupNamespacedName(vrg, namespace, backupName string, sequenceNumber int) types.NamespacedName {
	name := fmt.Sprintf("%s-%s-%s-%d", vrg, namespace, backupName, sequenceNumber)
	namespacedName := types.NamespacedName{
		Name:      name,
		Namespace: namespace, // note: must create backups in the same namespace as Velero/OADP
	}

	return namespacedName
}

func (v *VRGInstance) kubeObjectsRecover(objectStore ObjectStorer) error {
	if v.instance.Spec.KubeObjectProtection == nil {
		v.log.Info("Kube object recovery disabled")

		return nil
	}

	categories := v.getTypeSequenceSubRestores()

	var err error // declare for scope

	for restoreIndex, restoreInstance := range categories {
		includeList, excludeList := getTypeSequenceResourceList(restoreInstance.IncludedResources,
			restoreInstance.ExcludedResources)

		spec := restoreInstance.DeepCopy() // don't modify spec with processed results
		spec.IncludedResources = includeList
		spec.ExcludedResources = excludeList

		sourceNamespacedName := types.NamespacedName{Name: v.instance.Name, Namespace: v.instance.Namespace}

		if err := kubeObjectsRecover(
			v.ctx,
			v.reconciler.Client,
			v.reconciler.APIReader,
			v.log,
			objectStore.AddressComponent1(),
			objectStore.AddressComponent2(),
			v.s3KeyPrefix(),
			// TODO query source namespace from velero backup kube object in s3 store
			sourceNamespacedName,
			v.instance.Namespace,
			*spec,
			restoreIndex,
		); err != nil {
			v.log.Error(err, "kube object recover")
			// don't return error yet since it could be non-fatal (e.g. slow API server)
		}

		// for backupName, user specifies name. for restoreName, it's generated by system.
		restoreNamespacedName := getBackupNamespacedName(sourceNamespacedName.Name,
			sourceNamespacedName.Namespace, fmt.Sprintf("%d", restoreIndex), getSequenceNumber())

		// can only create backup/restores in Velero/OADP namespace
		restoreNamespacedName.Namespace = VeleroNamespaceNameDefault

		// TODO: remove tight loop here
		restoreComplete := false
		for !restoreComplete {
			restoreComplete, err = restoreIsDone(v.ctx, v.reconciler.APIReader,
				objectWriter{v.reconciler.Client, v.ctx, v.log}, restoreNamespacedName)

			if err != nil {
				restoreComplete = true
			}
		}
	}

	return err
}
