package options

import (
	goflag "flag"

	"github.com/caicloud/clientset/kubernetes"
	"github.com/caicloud/loadbalancer-controller/cmd/controller/app/config"
	lbconfig "github.com/caicloud/loadbalancer-controller/pkg/config"
	"github.com/spf13/pflag"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

const (
	// UserAgent is the userAgent name when starting loadbalancer controller.
	UserAgent = "loadbalancer-controller"
)

// Options is the main context object for the admission controller.
type Options struct {
	Master     string
	Kubeconfig string
	Cfg        lbconfig.Configuration
}

// NewOptions creates a new AddmissionOptions with a default config.
func NewOptions() *Options {
	return &Options{}
}

// Flags returns flags for admission controller
func (s *Options) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("options", pflag.ExitOnError)

	s.Cfg.AddFlags(fs)

	fs.StringVar(&s.Kubeconfig, "kubeconfig", s.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	fs.StringVar(&s.Master, "master", s.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig).")

	// init log
	gofs := goflag.NewFlagSet("klog", goflag.ExitOnError)
	klog.InitFlags(gofs)

	fs.AddGoFlagSet(gofs)

	return fs
}

// Config return a controller config objective
func (s *Options) Config() (*config.Config, error) {
	kubeconfig, err := clientcmd.BuildConfigFromFlags(s.Master, s.Kubeconfig)
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(restclient.AddUserAgent(kubeconfig, UserAgent))
	if err != nil {
		return nil, err
	}

	s.Cfg.Client = client
	c := &config.Config{
		Cfg: s.Cfg,
	}

	return c, nil
}
