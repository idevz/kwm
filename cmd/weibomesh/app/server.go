package app

import (
	"github.com/idevz/kwm/cmd/weibomesh/signals"
	"github.com/spf13/cobra"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"time"

	"k8s.io/client-go/tools/clientcmd"
	informers "github.com/idevz/kwm/pkg/client/informers/externalversions"
	clientset "github.com/idevz/kwm/pkg/client/clientset/versioned"
)
var (
	kubeconfig string
	masterURL string
)
func NewWeiboMeshCommand()  *cobra.Command{
	cmd := &cobra.Command{
		Use:"WeiboMesh",
		Long:`WeiboMesh with k8s`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run()
		},
	}
	cmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	cmd.PersistentFlags().StringVar(&masterURL, "masterURL", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	return cmd
}

func Run() error {
	klog.InitFlags(nil)

	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err !=nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	kwmClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kwm clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	kwmInformerFactory := informers.NewSharedInformerFactory(kwmClient, time.Second*30)

	controller := NewController(
		kubeClient,kwmClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		kwmInformerFactory.Weibomesh().V1().WeibMeshes()
		)

	return nil
}

//func init()  {
//}