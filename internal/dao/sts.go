package dao

import (
	"context"
	"errors"
	"fmt"

	"github.com/derailed/k9s/internal/client"
	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubectl/pkg/polymorphichelpers"
)

var (
	_ Accessor    = (*StatefulSet)(nil)
	_ Nuker       = (*StatefulSet)(nil)
	_ Loggable    = (*StatefulSet)(nil)
	_ Restartable = (*StatefulSet)(nil)
	_ Scalable    = (*StatefulSet)(nil)
	_ Controller  = (*StatefulSet)(nil)
)

// StatefulSet represents a K8s sts.
type StatefulSet struct {
	Resource
}

// IsHappy check for happy sts.
func (s *StatefulSet) IsHappy(sts appsv1.StatefulSet) bool {
	return sts.Status.Replicas == sts.Status.ReadyReplicas
}

// Scale a StatefulSet.
func (s *StatefulSet) Scale(ctx context.Context, path string, replicas int32) error {
	ns, n := client.Namespaced(path)
	auth, err := s.Client().CanI(ns, "apps/v1/statefulsets:scale", []string{client.GetVerb, client.UpdateVerb})
	if err != nil {
		return err
	}
	if !auth {
		return fmt.Errorf("user is not authorized to scale statefulsets")
	}

	dial, err := s.Client().Dial()
	if err != nil {
		return err
	}
	scale, err := dial.AppsV1().StatefulSets(ns).GetScale(ctx, n, metav1.GetOptions{})
	if err != nil {
		return err
	}
	scale.Spec.Replicas = replicas
	_, err = dial.AppsV1().StatefulSets(ns).UpdateScale(ctx, n, scale, metav1.UpdateOptions{})

	return err
}

// Restart a StatefulSet rollout.
func (s *StatefulSet) Restart(ctx context.Context, path string) error {
	sts, err := s.getStatefulSet(path)
	if err != nil {
		return err
	}

	ns, _ := client.Namespaced(path)
	auth, err := s.Client().CanI(ns, "apps/v1/statefulsets", []string{client.PatchVerb})
	if err != nil {
		return err
	}
	if !auth {
		return fmt.Errorf("user is not authorized to update statefulsets")
	}

	update, err := polymorphichelpers.ObjectRestarterFn(sts)
	if err != nil {
		return err
	}

	dial, err := s.Client().Dial()
	if err != nil {
		return err
	}
	_, err = dial.AppsV1().StatefulSets(sts.Namespace).Patch(
		ctx,
		sts.Name,
		types.StrategicMergePatchType,
		update,
		metav1.PatchOptions{},
	)
	return err
}

// TailLogs tail logs for all pods represented by this StatefulSet.
func (s *StatefulSet) TailLogs(ctx context.Context, c LogChan, opts LogOptions) error {
	sts, err := s.getStatefulSet(opts.Path)
	if err != nil {
		return errors.New("expecting StatefulSet resource")
	}
	if sts.Spec.Selector == nil || len(sts.Spec.Selector.MatchLabels) == 0 {
		return fmt.Errorf("No valid selector found on StatefulSet %s", opts.Path)
	}

	return podLogs(ctx, c, sts.Spec.Selector.MatchLabels, opts)
}

// Pod returns a pod victim by name.
func (s *StatefulSet) Pod(fqn string) (string, error) {
	sts, err := s.getStatefulSet(fqn)
	if err != nil {
		return "", err
	}

	return podFromSelector(s.Factory, sts.Namespace, sts.Spec.Selector.MatchLabels)
}

func (s *StatefulSet) getStatefulSet(fqn string) (*appsv1.StatefulSet, error) {
	o, err := s.Factory.Get(s.gvr.String(), fqn, true, labels.Everything())
	if err != nil {
		return nil, err
	}

	var sts appsv1.StatefulSet
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &sts)
	if err != nil {
		return nil, errors.New("expecting Service resource")
	}

	return &sts, nil
}

// ScanSA scans for serviceaccount refs.
func (s *StatefulSet) ScanSA(ctx context.Context, fqn string, wait bool) (Refs, error) {
	ns, n := client.Namespaced(fqn)
	oo, err := s.Factory.List(s.GVR(), ns, wait, labels.Everything())
	if err != nil {
		return nil, err
	}

	refs := make(Refs, 0, len(oo))
	for _, o := range oo {
		var sts appsv1.StatefulSet
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &sts)
		if err != nil {
			return nil, errors.New("expecting StatefulSet resource")
		}
		if sts.Spec.Template.Spec.ServiceAccountName == n {
			refs = append(refs, Ref{
				GVR: s.GVR(),
				FQN: client.FQN(sts.Namespace, sts.Name),
			})
		}
	}

	return refs, nil
}

// Scan scans for cluster resource refs.
func (s *StatefulSet) Scan(ctx context.Context, gvr, fqn string, wait bool) (Refs, error) {
	ns, n := client.Namespaced(fqn)
	oo, err := s.Factory.List(s.GVR(), ns, wait, labels.Everything())
	if err != nil {
		return nil, err
	}

	refs := make(Refs, 0, len(oo))
	for _, o := range oo {
		var sts appsv1.StatefulSet
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &sts)
		if err != nil {
			return nil, errors.New("expecting StatefulSet resource")
		}
		switch gvr {
		case "v1/configmaps":
			if !hasConfigMap(&sts.Spec.Template.Spec, n) {
				continue
			}
			refs = append(refs, Ref{
				GVR: s.GVR(),
				FQN: client.FQN(sts.Namespace, sts.Name),
			})
		case "v1/secrets":
			found, err := hasSecret(s.Factory, &sts.Spec.Template.Spec, sts.Namespace, n, wait)
			if err != nil {
				log.Warn().Err(err).Msgf("locate secret %q", fqn)
				continue
			}
			if !found {
				continue
			}
			refs = append(refs, Ref{
				GVR: s.GVR(),
				FQN: client.FQN(sts.Namespace, sts.Name),
			})
		}
	}

	return refs, nil
}
