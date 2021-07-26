package main
//
//import (
//	"context"
//	"errors"
//	"flag"
//	"fmt"
//	"installer/pkg/manage"
//	"installer/pkg/util"
//	v1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/labels"
//	"k8s.io/client-go/discovery"
//	"k8s.io/client-go/kubernetes"
//	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
//	"k8s.io/client-go/rest"
//	"k8s.io/client-go/tools/clientcmd"
//	"k8s.io/client-go/util/homedir"
//	metricsapi "k8s.io/metrics/pkg/apis/metrics"
//	metricsV1beta1api "k8s.io/metrics/pkg/apis/metrics/v1beta1"
//	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"
//	"path/filepath"
//	"sort"
//	"strconv"
//)
//
//var Kubeconfig *string
//
//func configCheck() (bool, string, *kubernetes.Clientset, *rest.Config) {
//	var kubeconfig *string
//	var clientset *kubernetes.Clientset
//
//	if Kubeconfig == nil {
//		if home := homedir.HomeDir(); home != "" {
//			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
//		} else {
//			kubeconfig = flag.String("kubeconfig", "", "config")
//		}
//		flag.Parse()
//		Kubeconfig = kubeconfig
//	}
//
//	config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
//	if err != nil {
//		return false, "build config error, check input file path", nil, nil
//	}
//
//	clientset, err = kubernetes.NewForConfig(config)
//	if err != nil {
//		return false, "create clientset error, check input file available", nil, nil
//	}
//
//	return true, "", clientset, config
//}
//
//type TopNodeOptions struct {
//	ResourceName       string
//	Selector           string
//	SortBy             string
//	UseProtocolBuffers bool
//
//	NodeClient      corev1client.CoreV1Interface
//	DiscoveryClient discovery.DiscoveryInterface
//	MetricsClient   metricsclientset.Interface
//}
//
//const GroupName = "metrics.k8s.io"
//
//var supportedMetricsAPIVersions = []string{"v1beta1"}
//
//func SupportedMetricsAPIVersionAvailable(discoveredAPIGroups *metav1.APIGroupList) bool {
//	for _, discoveredAPIGroup := range discoveredAPIGroups.Groups {
//		if discoveredAPIGroup.Name != GroupName {
//			continue
//		}
//		for _, version := range discoveredAPIGroup.Versions {
//			for _, supportedVersion := range supportedMetricsAPIVersions {
//				if version.Version == supportedVersion {
//					return true
//				}
//			}
//		}
//	}
//	return false
//}
//
//func getNodeMetricsFromMetricsAPI(metricsClient metricsclientset.Interface, resourceName string, selector labels.Selector) (*metricsapi.NodeMetricsList, error) {
//	var err error
//	versionedMetrics := &metricsV1beta1api.NodeMetricsList{}
//	mc := metricsClient.MetricsV1beta1()
//	nm := mc.NodeMetricses()
//	if resourceName != "" {
//		m, err := nm.Get(context.TODO(), resourceName, metav1.GetOptions{})
//		if err != nil {
//			return nil, err
//		}
//		versionedMetrics.Items = []metricsV1beta1api.NodeMetrics{*m}
//	} else {
//		versionedMetrics, err = nm.List(context.TODO(), metav1.ListOptions{LabelSelector: selector.String()})
//		if err != nil {
//			return nil, err
//		}
//	}
//	metrics := &metricsapi.NodeMetricsList{}
//	err = metricsV1beta1api.Convert_v1beta1_NodeMetricsList_To_metrics_NodeMetricsList(versionedMetrics, metrics, nil)
//	if err != nil {
//		return nil, err
//	}
//	return metrics, nil
//}
//
//type NodeMetricsSorter struct {
//	metrics []metricsapi.NodeMetrics
//	sortBy  string
//}
//
//func (n *NodeMetricsSorter) Len() int {
//	return len(n.metrics)
//}
//
//func (n *NodeMetricsSorter) Swap(i, j int) {
//	n.metrics[i], n.metrics[j] = n.metrics[j], n.metrics[i]
//}
//
//func (n *NodeMetricsSorter) Less(i, j int) bool {
//	switch n.sortBy {
//	case "cpu":
//		return n.metrics[i].Usage.Cpu().MilliValue() > n.metrics[j].Usage.Cpu().MilliValue()
//	case "memory":
//		return n.metrics[i].Usage.Memory().Value() > n.metrics[j].Usage.Memory().Value()
//	default:
//		return n.metrics[i].Name < n.metrics[j].Name
//	}
//}
//
//func NewNodeMetricsSorter(metrics []metricsapi.NodeMetrics, sortBy string) *NodeMetricsSorter {
//	return &NodeMetricsSorter{
//		metrics: metrics,
//		sortBy:  sortBy,
//	}
//}
//func PrintNodeMetrics(metrics []metricsapi.NodeMetrics, availableResources map[string]v1.ResourceList, sortBy string) error {
//	if len(metrics) == 0 {
//		return nil
//	}
//
//	sort.Sort(NewNodeMetricsSorter(metrics, sortBy))
//
//	var nodeStatus []manage.NodeStatusOutPutStructure
//
//	var usage v1.ResourceList
//	for _, m := range metrics {
//
//		m.Usage.DeepCopyInto(&usage)
//
//		var name, cpu, cpuPer, memory, memoryPer string
//		name = m.Name
//		for _, res := range MeasuredResources {
//			quantity := usage[res]
//			switch res {
//			case v1.ResourceCPU:
//				cpu = strconv.FormatInt(quantity.MilliValue(), 10) + "m"
//				if available, found := availableResources[m.Name][res]; found {
//					fraction := float64(quantity.MilliValue()) / float64(available.MilliValue()) * 100
//					cpuPer = util.BuildAsciiMeterCurrentTotal(int64(fraction), 100, 15) + strconv.FormatInt(int64(fraction), 10) + "%"
//				}
//			case v1.ResourceMemory:
//				memory = strconv.FormatInt(quantity.Value()/(1024*1024), 10) + "Mi"
//				if available, found := availableResources[m.Name][res]; found {
//					fraction := float64(quantity.MilliValue()) / float64(available.MilliValue()) * 100
//					memoryPer = util.BuildAsciiMeterCurrentTotal(int64(fraction), 100, 15) + strconv.FormatInt(int64(fraction), 10) + "%"
//				}
//			default:
//				fmt.Printf("%v", quantity.Value())
//			}
//		}
//		nodeStatus = append(nodeStatus, manage.NodeStatusOutPutStructure{
//			Name:      name,
//			CPU:       cpu,
//			CPUPer:    cpuPer,
//			MEMORY:    memory,
//			MEMORYPer: memoryPer,
//		})
//		delete(availableResources, m.Name)
//	}
//
//	for i := 0; i < len(nodeStatus); i++ {
//		fmt.Print(nodeStatus[i].Name + nodeStatus[i].CPUPer + "\n" + nodeStatus[i].MEMORYPer)
//	}
//	return nil
//}
//
//var MeasuredResources = []v1.ResourceName{
//	v1.ResourceCPU,
//	v1.ResourceMemory,
//}
//
//func main() {
//	b, s, clientSet, config := configCheck()
//
//	if !b {
//		fmt.Print(s)
//		return
//	}
//
//	metricsClient, err := metricsclientset.NewForConfig(config)
//	if err != nil {
//		fmt.Print("error")
//	}
//	o := &TopNodeOptions{
//		UseProtocolBuffers: true,
//		NodeClient:         clientSet.CoreV1(),
//		DiscoveryClient:    clientSet.DiscoveryClient,
//		MetricsClient:      metricsClient,
//	}
//
//	_ = o.RunTopNode()
//}
//
//func (o TopNodeOptions) RunTopNode() error {
//	var err error
//	selector := labels.Everything()
//	if len(o.Selector) > 0 {
//		selector, err = labels.Parse(o.Selector)
//		if err != nil {
//			return err
//		}
//	}
//
//	apiGroups, err := o.DiscoveryClient.ServerGroups()
//	if err != nil {
//		return err
//	}
//
//	metricsAPIAvailable := SupportedMetricsAPIVersionAvailable(apiGroups)
//
//	if !metricsAPIAvailable {
//		return errors.New("metrics API not available.")
//	}
//
//	metrics, err := getNodeMetricsFromMetricsAPI(o.MetricsClient, o.ResourceName, selector)
//	if err != nil {
//		return err
//	}
//
//	if len(metrics.Items) == 0 {
//		return errors.New("metrics not available yet")
//	}
//
//	var nodes []v1.Node
//	if len(o.ResourceName) > 0 {
//		node, err := o.NodeClient.Nodes().Get(context.TODO(), o.ResourceName, metav1.GetOptions{})
//		if err != nil {
//			return err
//		}
//		nodes = append(nodes, *node)
//	} else {
//		nodeList, err := o.NodeClient.Nodes().List(context.TODO(), metav1.ListOptions{
//			LabelSelector: selector.String(),
//		})
//		if err != nil {
//			return err
//		}
//		nodes = append(nodes, nodeList.Items...)
//	}
//
//	allocatable := make(map[string]v1.ResourceList)
//
//	for _, n := range nodes {
//		allocatable[n.Name] = n.Status.Allocatable
//	}
//
//	return PrintNodeMetrics(metrics.Items, allocatable, o.SortBy)
//}
