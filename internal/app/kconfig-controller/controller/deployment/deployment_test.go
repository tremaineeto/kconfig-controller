package deployment

import (
	"reflect"
	"testing"

	"github.com/gbraxton/kconfig/pkg/apis/kconfigcontroller/v1alpha1"
	kcfake "github.com/gbraxton/kconfig/pkg/client/clientset/versioned/fake"
	kcinformers "github.com/gbraxton/kconfig/pkg/client/informers/externalversions"
	testutil "github.com/gbraxton/kconfig/test/util"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	stdinformers "k8s.io/client-go/informers"
	stdfake "k8s.io/client-go/kubernetes/fake"
	core "k8s.io/client-go/testing"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
)

var (
	alwaysReady = func() bool { return true }
)

type fixture struct {
	t *testing.T

	stdclient *stdfake.Clientset
	kcclient  *kcfake.Clientset

	// Objects to put in the store
	deploymentLister []*appsv1.Deployment
	kbindingLister   []*v1alpha1.KconfigBinding

	// Actions expected to happen on the client. Objects from here are also
	// preloaded into NewSimpleFake.
	stdactions []core.Action
	kcactions  []core.Action
	stdobjects []runtime.Object
	kcobjects  []runtime.Object
}

func newFixture(t *testing.T) *fixture {
	f := &fixture{}
	f.t = t
	f.stdobjects = []runtime.Object{}
	f.kcobjects = []runtime.Object{}
	return f
}

func (f *fixture) newController() (*Controller, kcinformers.SharedInformerFactory, stdinformers.SharedInformerFactory, error) {
	f.stdclient = stdfake.NewSimpleClientset(f.stdobjects...)
	f.kcclient = kcfake.NewSimpleClientset(f.kcobjects...)

	stdinformers := stdinformers.NewSharedInformerFactory(f.stdclient, 0)
	kcinformers := kcinformers.NewSharedInformerFactory(f.kcclient, 0)

	c := NewController(f.stdclient, f.kcclient, stdinformers.Apps().V1().Deployments(), kcinformers.Kconfigcontroller().V1alpha1().KconfigBindings())
	c.recorder = &record.FakeRecorder{}
	c.deploymentsSynced = alwaysReady
	c.kconfigBindingsSynced = alwaysReady
	for _, d := range f.deploymentLister {
		stdinformers.Apps().V1().Deployments().Informer().GetIndexer().Add(d)
	}
	for _, kcb := range f.kbindingLister {
		kcinformers.Kconfigcontroller().V1alpha1().KconfigBindings().Informer().GetIndexer().Add(kcb)
	}
	return c, kcinformers, stdinformers, nil
}

func (f *fixture) runExpectError(deploymentName string, startInformers bool) {
	f.runSync(deploymentName, startInformers, true)
}

func (f *fixture) run(deploymentName string) {
	f.runSync(deploymentName, true, false)
}

func (f *fixture) runSync(deploymentName string, startInformers bool, expectError bool) {
	c, kcinformers, stdinformers, err := f.newController()
	if err != nil {
		f.t.Fatalf("error creating Kconfig controller: %v", err)
	}
	if startInformers {
		stopCh := make(chan struct{})
		defer close(stopCh)
		kcinformers.Start(stopCh)
		stdinformers.Start(stopCh)
	}

	err = c.syncHandler(deploymentName)
	if !expectError && err != nil {
		f.t.Errorf("error syncing kconfig: %v", err)
	} else if expectError && err == nil {
		f.t.Error("expected error syncing kconfig, got nil")
	}

	f.checkActions()
}

// runDelete calls deleteKconfig instead of syncHandler
func (f *fixture) runDelete(obj interface{}) {
	startInformers := true

	c, kcinformers, stdinformers, err := f.newController()
	if err != nil {
		f.t.Fatalf("error creating Deployment controller: %v", err)
	}
	if startInformers {
		stopCh := make(chan struct{})
		defer close(stopCh)
		kcinformers.Start(stopCh)
		stdinformers.Start(stopCh)
	}

	c.deleteHandler(obj)

	f.checkActions()
}

func (f *fixture) checkActions() {
	stdactions := filterStdInformerActions(f.stdclient.Actions())
	for i, stdaction := range stdactions {
		if len(f.stdactions) < i+1 {
			f.t.Errorf("%d unexpected actions: %+v", len(stdactions)-len(f.stdactions), stdactions[i:])
			break
		}

		expectedAction := f.stdactions[i]
		if !(expectedAction.Matches(stdaction.GetVerb(), stdaction.GetResource().Resource) && stdaction.GetSubresource() == expectedAction.GetSubresource()) {
			f.t.Errorf("Expected\n\t%#v\ngot\n\t%#v", expectedAction, stdaction)
			continue
		}
		if !f.actionObjectsMatch(expectedAction, stdaction) {
			continue
		}
	}
	if len(f.stdactions) > len(stdactions) {
		f.t.Errorf("%d additional expected actions:%+v", len(f.stdactions)-len(stdactions), f.stdactions[len(stdactions):])
	}

	kcactions := filterKcInformerActions(f.kcclient.Actions())
	for i, kcaction := range kcactions {
		if len(f.kcactions) < i+1 {
			f.t.Errorf("%d unexpected actions: %+v", len(kcactions)-len(f.kcactions), kcactions[i:])
			break
		}

		expectedAction := f.kcactions[i]
		if !(expectedAction.Matches(kcaction.GetVerb(), kcaction.GetResource().Resource) && kcaction.GetSubresource() == expectedAction.GetSubresource()) {
			f.t.Errorf("Expected\n\t%#v\ngot\n\t%#v", expectedAction, kcaction)
			continue
		}
		if !f.actionObjectsMatch(expectedAction, kcaction) {
			continue
		}
	}
	if len(f.kcactions) > len(kcactions) {
		f.t.Errorf("%d additional expected actions:%+v", len(f.kcactions)-len(kcactions), f.kcactions[len(kcactions):])
	}
}

// actionObjectsMatch Assumes expectedAction and action have already had their
// verbs matched. Always returns true if params aren't Create or Update actions.
func (f *fixture) actionObjectsMatch(expectedAction, action core.Action) bool {
	if expectedCreateAction, ok := expectedAction.(core.CreateAction); ok {
		createAction, _ := action.(core.CreateAction)
		if !reflect.DeepEqual(expectedCreateAction.GetObject(), createAction.GetObject()) {
			f.t.Errorf("Expected\n\t%+v\ngot\n\t%+v", expectedCreateAction.GetObject(), createAction.GetObject())
			return false
		}
		return true
	}
	if expectedUpdateAction, ok := expectedAction.(core.UpdateAction); ok {
		updateAction, _ := action.(core.UpdateAction)
		if !reflect.DeepEqual(expectedUpdateAction.GetObject(), updateAction.GetObject()) {
			f.t.Errorf("Expected\n\t%+v\ngot\n\t%+v", expectedUpdateAction.GetObject(), updateAction.GetObject())
			return false
		}
		return true
	}
	return true
}

func filterStdInformerActions(actions []core.Action) []core.Action {
	ret := []core.Action{}
	for _, action := range actions {
		if len(action.GetNamespace()) == 0 &&
			(action.Matches("watch", "deployments") ||
				action.Matches("list", "deployments")) {
			continue
		}
		ret = append(ret, action)
	}

	return ret
}

func filterKcInformerActions(actions []core.Action) []core.Action {
	ret := []core.Action{}
	for _, action := range actions {
		if len(action.GetNamespace()) == 0 &&
			(action.Matches("watch", "kconfigbindings") ||
				action.Matches("list", "kconfigbindings")) {
			continue
		}
		ret = append(ret, action)
	}
	return ret
}

func (f *fixture) expectCreateKconfigBindingAction(k *v1alpha1.KconfigBinding) {
	resource := schema.GroupVersionResource{
		Group:    v1alpha1.SchemeGroupVersion.Group,
		Version:  v1alpha1.SchemeGroupVersion.Version,
		Resource: "kconfigbindings",
	}
	action := core.NewCreateAction(resource, k.Namespace, k)
	f.kcactions = append(f.kcactions, action)
}

func (f *fixture) expectUpdateKconfigBindingAction(k *v1alpha1.KconfigBinding) {
	resource := schema.GroupVersionResource{
		Group:    v1alpha1.SchemeGroupVersion.Group,
		Version:  v1alpha1.SchemeGroupVersion.Version,
		Resource: "kconfigbindings",
	}
	action := core.NewUpdateAction(resource, k.Namespace, k)
	f.kcactions = append(f.kcactions, action)
}

func (f *fixture) expectDeleteKconfigBindingAction(k *v1alpha1.KconfigBinding) {
	resource := schema.GroupVersionResource{
		Group:    v1alpha1.SchemeGroupVersion.Group,
		Version:  v1alpha1.SchemeGroupVersion.Version,
		Resource: "kconfigbindings",
	}
	action := core.NewDeleteAction(resource, k.Namespace, k.Name)
	f.kcactions = append(f.kcactions, action)
}

func TestNewDeploymentCreatesKconfigBinding(t *testing.T) {
	f := newFixture(t)

	d := testutil.Deployment()
	expectedKcbCreate := testutil.KconfigBinding()
	expectedKcbUpdate := testutil.KconfigBinding()

	f.deploymentLister = append(f.deploymentLister, &d)
	f.stdobjects = append(f.stdobjects, &d)

	f.expectCreateKconfigBindingAction(&expectedKcbCreate)
	f.expectUpdateKconfigBindingAction(&expectedKcbUpdate)

	key, _ := cache.MetaNamespaceKeyFunc(&d.ObjectMeta)
	f.run(key)
}

func TestDeleteDeploymentDeletesKconfigBinding(t *testing.T) {
	f := newFixture(t)

	d := testutil.Deployment()
	now := metav1.Now()
	d.ObjectMeta.DeletionTimestamp = &now
	kcb := testutil.KconfigBinding()

	f.kbindingLister = append(f.kbindingLister, &kcb)
	f.kcobjects = append(f.kcobjects, &kcb)

	f.expectDeleteKconfigBindingAction(&kcb)

	key, _ := cache.MetaNamespaceKeyFunc(&d.ObjectMeta)
	f.runDelete(cache.DeletedFinalStateUnknown{Key: key, Obj: &d})
}
