/*
Copyright 2021.

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

package hyperNode

import (
	"context"
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	clusterv1 "computedriver/pkg/hyperNode/api/v1"
	"computedriver/pkg/hyperNode/controllers"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

var kubeConfig string
var config string

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(clusterv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
	kubeConfig = filepath.Join("/etc/kubernetes/kubelet.conf")
}

func Run() {
	// var metricsAddr string
	// var enableLeaderElection bool
	// var probeAddr string
	// flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	// flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	// flag.BoolVar(&enableLeaderElection, "leader-elect", false,
	// 	"Enable leader election for controller manager. "+
	// 		"Enabling this will ensure there is only one active controller manager.")
	// opts := zap.Options{
	// 	Development: true,
	// }
	// opts.BindFlags(flag.CommandLine)
	// flag.Parse()
	//
	// ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	mgr, err := ctrl.NewManager(config, ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: "0",
		//Port:                   9443,
		//HealthProbeBindAddress: probeAddr,
		LeaderElection:   false,
		LeaderElectionID: "663e3d41.pml.com.cn",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.HypernodeReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Hypernode")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	// if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
	// 	setupLog.Error(err, "unable to set up health check")
	// 	os.Exit(1)
	// }
	// if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
	// 	setupLog.Error(err, "unable to set up ready check")
	// 	os.Exit(1)
	// }

	setupLog.Info("starting manager")

	ctx := context.TODO()
	ctxContext, _ := context.WithCancel(ctx)
	if err := mgr.Start(ctxContext); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
