package cluster

import (
	"context"
	"errors"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/log"
	"github.com/openshift/openshift-azure/pkg/util/wait"
)

var deploymentWhitelist = []struct {
	Name      string
	Namespace string
}{
	{
		Name:      "docker-registry",
		Namespace: "default",
	},
	{
		Name:      "router",
		Namespace: "default",
	},
	{
		Name:      "registry-console",
		Namespace: "default",
	},
	{
		Name:      "customer-admin-controller",
		Namespace: "openshift-infra",
	},
	{
		Name:      "asb",
		Namespace: "openshift-ansible-service-broker",
	},
	{
		Name:      "webconsole",
		Namespace: "openshift-web-console",
	},
	{
		Name:      "console",
		Namespace: "openshift-console",
	},
	{
		Name:      "cluster-monitoring-operator",
		Namespace: "openshift-monitoring",
	},
}

var daemonsetWhitelist = []struct {
	Name      string
	Namespace string
}{
	{
		Name:      "sync",
		Namespace: "openshift-node",
	},
	{
		Name:      "ovs",
		Namespace: "openshift-sdn",
	},
	{
		Name:      "sdn",
		Namespace: "openshift-sdn",
	},
	{
		Name:      "apiserver",
		Namespace: "kube-service-catalog",
	},
	{
		Name:      "controller-manager",
		Namespace: "kube-service-catalog",
	},
	{
		Name:      "apiserver",
		Namespace: "openshift-template-service-broker",
	},
}

var statefulsetWhitelist = []struct {
	Name      string
	Namespace string
}{
	{
		Name:      "bootstrap-autoapprover",
		Namespace: "openshift-infra",
	},
}

type computerName string

func (computerName computerName) toKubernetes() string {
	return strings.ToLower(string(computerName))
}

func (u *simpleUpgrader) waitForNodes(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	for _, role := range []api.AgentPoolProfileRole{api.AgentPoolProfileRoleMaster, api.AgentPoolProfileRoleInfra, api.AgentPoolProfileRoleCompute} {
		vms, err := u.listVMs(ctx, cs, role)
		if err != nil {
			return err
		}
		for _, vm := range vms {
			computerName := computerName(*vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
			log.Infof("waiting for %s to be ready", computerName)
			err = u.waitForReady(ctx, cs, role, computerName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *simpleUpgrader) WaitForInfraServices(ctx context.Context, cs *api.OpenShiftManagedCluster) *api.PluginError {
	for _, app := range daemonsetWhitelist {
		log.Infof("checking daemonset %s/%s", app.Namespace, app.Name)

		err := wait.PollImmediateUntil(time.Second, func() (bool, error) {
			ds, err := u.kubeclient.AppsV1().DaemonSets(app.Namespace).Get(app.Name, metav1.GetOptions{})
			switch {
			case kerrors.IsNotFound(err):
				return false, nil
			case err == nil:
				return ds.Status.DesiredNumberScheduled == ds.Status.CurrentNumberScheduled &&
					ds.Status.DesiredNumberScheduled == ds.Status.NumberReady &&
					ds.Status.DesiredNumberScheduled == ds.Status.UpdatedNumberScheduled &&
					ds.Generation == ds.Status.ObservedGeneration, nil
			default:
				return false, err
			}
		}, ctx.Done())
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepWaitForInfraDaemonSets}
		}
	}

	for _, app := range statefulsetWhitelist {
		log.Infof("checking statefulset %s/%s", app.Namespace, app.Name)

		err := wait.PollImmediateUntil(time.Second, func() (bool, error) {
			sts, err := u.kubeclient.AppsV1().StatefulSets(app.Namespace).Get(app.Name, metav1.GetOptions{})
			switch {
			case kerrors.IsNotFound(err):
				return false, nil
			case err == nil:
				specReplicas := int32(1)
				specReplicas = *sts.Spec.Replicas

				return specReplicas == sts.Status.Replicas &&
					specReplicas == sts.Status.ReadyReplicas &&
					specReplicas == sts.Status.CurrentReplicas &&
					specReplicas == sts.Status.UpdatedReplicas &&
					sts.Generation == sts.Status.ObservedGeneration, nil
			default:
				return false, err
			}
		}, ctx.Done())
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepWaitForInfraStatefulSets}
		}
	}

	for _, app := range deploymentWhitelist {
		log.Infof("checking deployment %s/%s", app.Namespace, app.Name)

		err := wait.PollImmediateUntil(time.Second, func() (bool, error) {
			d, err := u.kubeclient.AppsV1().Deployments(app.Namespace).Get(app.Name, metav1.GetOptions{})
			switch {
			case kerrors.IsNotFound(err):
				return false, nil
			case err == nil:
				specReplicas := int32(1)
				if d.Spec.Replicas != nil {
					specReplicas = *d.Spec.Replicas
				}

				return specReplicas == d.Status.Replicas &&
					specReplicas == d.Status.ReadyReplicas &&
					specReplicas == d.Status.AvailableReplicas &&
					specReplicas == d.Status.UpdatedReplicas &&
					d.Generation == d.Status.ObservedGeneration, nil
			default:
				return false, err
			}
		}, ctx.Done())
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepWaitForInfraDeployments}
		}
	}

	return nil
}

func (u *simpleUpgrader) waitForReady(ctx context.Context, cs *api.OpenShiftManagedCluster, role api.AgentPoolProfileRole, computerName computerName) error {
	switch role {
	case api.AgentPoolProfileRoleMaster:
		return u.masterWaitForReady(ctx, cs, computerName)
	case api.AgentPoolProfileRoleInfra, api.AgentPoolProfileRoleCompute:
		return u.nodeWaitForReady(ctx, cs, computerName)
	default:
		return errors.New("unrecognised role")
	}
}

func (u *simpleUpgrader) masterWaitForReady(ctx context.Context, cs *api.OpenShiftManagedCluster, computerName computerName) error {
	return wait.PollImmediateUntil(time.Second, func() (bool, error) {
		return u.masterIsReady(computerName)
	}, ctx.Done())
}

func (u *simpleUpgrader) masterIsReady(computerName computerName) (bool, error) {
	ready, err := u.nodeIsReady(computerName)
	if !ready || err != nil {
		return ready, err
	}

	etcdPod, err := u.kubeclient.CoreV1().Pods("kube-system").Get("master-etcd-"+computerName.toKubernetes(), metav1.GetOptions{})
	switch {
	case err == nil:
	case kerrors.IsNotFound(err):
		return false, nil
	default:
		return false, err
	}

	apiPod, err := u.kubeclient.CoreV1().Pods("kube-system").Get("master-api-"+computerName.toKubernetes(), metav1.GetOptions{})
	switch {
	case err == nil:
	case kerrors.IsNotFound(err):
		return false, nil
	default:
		return false, err
	}

	cmPod, err := u.kubeclient.CoreV1().Pods("kube-system").Get("controllers-"+computerName.toKubernetes(), metav1.GetOptions{})
	switch {
	case err == nil:
	case kerrors.IsNotFound(err):
		return false, nil
	default:
		return false, err
	}

	return isPodReady(etcdPod) && isPodReady(apiPod) && isPodReady(cmPod), nil
}

func isPodReady(pod *corev1.Pod) bool {
	for _, c := range pod.Status.Conditions {
		if c.Type == corev1.PodReady {
			return c.Status == corev1.ConditionTrue
		}
	}
	return false
}

func (u *simpleUpgrader) nodeWaitForReady(ctx context.Context, cs *api.OpenShiftManagedCluster, computerName computerName) error {
	err := wait.PollImmediateUntil(time.Second, func() (bool, error) {
		return u.nodeIsReady(computerName)
	}, ctx.Done())
	if err != nil {
		return err
	}

	return u.setUnschedulable(ctx, computerName, false)
}

func (u *simpleUpgrader) nodeIsReady(computerName computerName) (bool, error) {
	node, err := u.kubeclient.CoreV1().Nodes().Get(computerName.toKubernetes(), metav1.GetOptions{})
	switch {
	case err == nil:
	case kerrors.IsNotFound(err):
		return false, nil
	default:
		return false, err
	}

	for _, c := range node.Status.Conditions {
		if c.Type == corev1.NodeReady {
			return c.Status == corev1.ConditionTrue, nil
		}
	}
	return false, nil
}