package storagecluster

import (
	"context"
	"reflect"

	ocsv1 "github.com/openshift/ocs-operator/pkg/apis/ocs/v1"
	statusutil "github.com/openshift/ocs-operator/pkg/controller/util"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ReconcileExternalStorageCluster reconciles the external Storage Cluster
func (r *ReconcileStorageCluster) ReconcileExternalStorageCluster(sc *ocsv1.StorageCluster) (reconcile.Result, error) {
	reqLogger := r.reqLogger.WithValues("Request.Namespace", sc.Namespace, "Request.Name", sc.Name)
	reqLogger.Info("Reconciling External StorageCluster")
	var err error

	// Add conditions if there are none
	if sc.Status.Conditions == nil {
		reason := ocsv1.ReconcileInit
		message := "Initializing External StorageCluster"
		statusutil.SetProgressingCondition(&sc.Status.Conditions, reason, message)
		err = r.client.Status().Update(context.TODO(), sc)
		if err != nil {
			reqLogger.Error(err, "Failed to add conditions to status")
			return reconcile.Result{}, err
		}
	}

	externalCephCluster := newExternalCephCluster(sc, r.cephImage)
	// Set StorageCluster instance as the owner and controller
	if err := controllerutil.SetControllerReference(sc, externalCephCluster, r.scheme); err != nil {
		return reconcile.Result{}, err
	}
	// Check if this CephCluster already exists
	found := &cephv1.CephCluster{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: externalCephCluster.Name, Namespace: externalCephCluster.Namespace}, found)
	if err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Creating External CephCluster")
			if err := r.client.Create(context.TODO(), externalCephCluster); err != nil {
				return reconcile.Result{}, err
			}
		}
		return reconcile.Result{}, err
	}
	// Update the CephCluster if it is not in the desired state
	if !reflect.DeepEqual(externalCephCluster.Spec, found.Spec) {
		reqLogger.Info("Updating spec for External CephCluster")
		sc.Status.Phase = string(found.Status.State)
		err = r.client.Status().Update(context.TODO(), sc)
		if err != nil {
			reqLogger.Error(err, "Failed to add conditions to status")
			return reconcile.Result{}, err
		}
		found.Spec = externalCephCluster.Spec
		if err := r.client.Update(context.TODO(), found); err != nil {
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}

func newExternalCephCluster(sc *ocsv1.StorageCluster, cephImage string) *cephv1.CephCluster {
	labels := map[string]string{
		"app": sc.Name,
	}
	externalCephCluster := &cephv1.CephCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      generateNameForExternalCephCluster(sc),
			Namespace: sc.Namespace,
			Labels:    labels,
		},
		Spec: cephv1.ClusterSpec{
			External: cephv1.ExternalSpec{
				Enable: true,
			},
			DataDirHostPath: "/var/lib/rook",
			CephVersion: cephv1.CephVersionSpec{
				Image: cephImage,
			},
		},
	}
	return externalCephCluster
}
