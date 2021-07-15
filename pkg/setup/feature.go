package setup

import (
	"github.com/rivo/tview"
	"installer/pkg/constants"
	"installer/pkg/util"
	"strconv"
	"strings"
)

type Features struct {
	*tview.Form
}

var Changed []Feature

func NewFeatures() *Features {

	//item := &tview.NewGrid().new
	features := &Features{
		Form: tview.NewForm(),
	}

	var result []Feature
	featureMap := NewFeatureMap()

	for i := 107; i <= constants.GateVersion; i++ {
		result = append(result, featureMap[i]...)
	}

	for r := range result {
		features.Form.AddCheckbox(result[r].Name, result[r].DefaultValue, func(checked bool) {
			for c := range result {
				if result[c].Name == result[r].Name && result[c].DefaultValue != checked {
					Changed = append(Changed, Feature{
						Name:         result[c].Name,
						DefaultValue: checked,
					})
				}
			}
		})
		//fmt.Print(result[r].name + " ")
		//fmt.Println(result[r].defaultValue)
	}

	features.SetBorder(true).SetTitle("Feature info").SetTitleAlign(tview.AlignCenter)
	features.SetItemPadding(0).SetBorderPadding(0, 0, 0, 10)

	//info.SetText("docker page  " + StructureToJSON(*s))

	return features
}

//func GetFeatureChange(f []Feature) []Feature{
//	return Changed
//}

type Feature struct {
	Name         string
	DefaultValue bool
}

func NewFeatureMap() map[int][]Feature {
	m := make(map[int][]Feature)

	m[107] = append(m[107], Feature{Name: "AllowExtTrafficLocalEndpoints", DefaultValue: true})

	m[108] = append(m[108], Feature{Name: "DynamicVolumeProvisioning", DefaultValue: true})

	m[110] = append(m[110], Feature{Name: "VolumeSubpath", DefaultValue: true})

	m[111] = append(m[111], Feature{Name: "StorageObjectInUseProtection", DefaultValue: true})
	m[111] = append(m[111], Feature{Name: "SupportIPVSProxyMode", DefaultValue: true})

	m[112] = append(m[112], Feature{Name: "AdvancedAuditing", DefaultValue: true})
	m[112] = append(m[112], Feature{Name: "MountPropagation", DefaultValue: true})

	m[113] = append(m[113], Feature{Name: "CSIPersistentVolume", DefaultValue: true})
	m[113] = append(m[113], Feature{Name: "GCERegionalPersistentDisk", DefaultValue: true})
	m[113] = append(m[113], Feature{Name: "KubeletPluginsWatcher", DefaultValue: true})
	m[113] = append(m[113], Feature{Name: "VolumeScheduling", DefaultValue: true})

	m[114] = append(m[114], Feature{Name: "HugePages", DefaultValue: true})
	m[114] = append(m[114], Feature{Name: "CustomPodDNS", DefaultValue: true})
	m[114] = append(m[114], Feature{Name: "PersistentLocalVolumes", DefaultValue: true})
	m[114] = append(m[114], Feature{Name: "PodPriority", DefaultValue: true})
	m[114] = append(m[114], Feature{Name: "PodReadinessGates", DefaultValue: true})

	m[116] = append(m[116], Feature{Name: "CustomResourcePublishOpenAPI", DefaultValue: true})
	m[116] = append(m[116], Feature{Name: "CustomResourceSubresources", DefaultValue: true})
	m[116] = append(m[116], Feature{Name: "CustomResourceValidation", DefaultValue: true})
	m[116] = append(m[116], Feature{Name: "CustomResourceWebhookConversion", DefaultValue: true})

	m[117] = append(m[117], Feature{Name: "AttachVolumeLimit", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "CustomResourceDefaulting", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "NodeLease", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "PodShareProcessNamespace", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "ResourceQuotaScopeSelectors", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "ScheduleDaemonSetPods", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "ServiceLoadBalancerFinalizer", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "TaintNodesByCondition", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "VolumeSubpathEnvExpansion", DefaultValue: true})
	m[117] = append(m[117], Feature{Name: "WatchBookmark", DefaultValue: true})

	m[118] = append(m[118], Feature{Name: "BlockVolume", DefaultValue: true})
	m[118] = append(m[118], Feature{Name: "CSIBlockVolume", DefaultValue: true})
	m[118] = append(m[118], Feature{Name: "ExternalPolicyForExternalIP", DefaultValue: true})
	m[118] = append(m[118], Feature{Name: "TaintBasedEvictions", DefaultValue: true})
	m[118] = append(m[118], Feature{Name: "VolumePVCDataSource", DefaultValue: true})
	m[118] = append(m[118], Feature{Name: "WindowsGMSA", DefaultValue: true})
	m[118] = append(m[118], Feature{Name: "WindowsRunAsUserName", DefaultValue: true})

	m[119] = append(m[119], Feature{Name: "DryRun", DefaultValue: true})
	m[119] = append(m[119], Feature{Name: "EvenPodsSpread", DefaultValue: true})
	m[119] = append(m[119], Feature{Name: "RotateKubeletClientCertificate", DefaultValue: true})

	m[120] = append(m[120], Feature{Name: "ExecProbeTimeout", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "RuntimeClass", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "SCTPSupport", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "ServiceAppProtocol", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "StartupProbe", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "SupportNodePidsLimit", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "SupportPodPidsLimit", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "TokenRequest", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "TokenRequestProjection", DefaultValue: true})
	m[120] = append(m[120], Feature{Name: "VolumeSnapshotDataSource", DefaultValue: true})

	m[121] = append(m[121], Feature{Name: "CRIContainerLogRotation", DefaultValue: true})
	m[121] = append(m[121], Feature{Name: "EndpointSlice", DefaultValue: true})
	m[121] = append(m[121], Feature{Name: "EndpointSliceNodeName", DefaultValue: true})
	m[121] = append(m[121], Feature{Name: "LegacyNodeRoleBehavior", DefaultValue: false})
	m[121] = append(m[121], Feature{Name: "NodeDisruptionExclusion", DefaultValue: true})
	m[121] = append(m[121], Feature{Name: "PodDisruptionBudget", DefaultValue: true})
	m[121] = append(m[121], Feature{Name: "RootCAConfigMap", DefaultValue: true})
	m[121] = append(m[121], Feature{Name: "ServiceAccountIssuerDiscovery", DefaultValue: true})
	m[121] = append(m[121], Feature{Name: "ServiceNodeExclusion", DefaultValue: true})

	return m
}

func GetFeatureGates(f []Feature) string {
	result := ""
	for i := 0; i < len(f); i++ {
		result = util.StringAppend(result, f[i].Name+"="+strconv.FormatBool(f[i].DefaultValue)+",")
	}
	return strings.TrimRight(result, ",")
}
