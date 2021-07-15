package main

import "fmt"

//
//import (
//	"bytes"
//	"os/exec"
//)
//
func main() {
	s1 := "123"

	a(&s1)

	fmt.Print(s1)

}

func a(s *string) {
	*s = *s + "456"
}

//	//version := "A"
//	//execShell("sed 's/a/" + version + "/g' test.txt")
//	//
//	//execShell("sed -i '/#SET_MASTER_SERVER_HERE/a " + version + "' ha.cfg")
//
//	var ip string
//	var card string
//
//	ip = "192.168.48.205"
//	card = "ens3"
//	execShell("sh add_ansible_host.sh " + "netCardChange " + ip)
//	execShell("ansible-playbook 1.yaml -e netCard=" + card)
//	//
//	//
//	//var s string
//	//s = "123342%%"
//	//
//	//fmt.Println(isDigitWithPercent(s))
//
//
//	//someString := "one    two   three four "
//	//
//	//words := strings.Fields(someString)
//	//
//	//fmt.Println(words, len(words))
//
//
//	//fmt.Print(regexp.MustCompile("[0-9]+").FindAllString("//docker:19.3.15", -1))
//	//
//	//
//	//str := regexp.MustCompile("[0-9]+").FindAllString("//docker:19.3.15", -1)
//	//
//	////result := "docker-ce-"docker-ce-19.03.15
//	//var buffer bytes.Buffer
//	//buffer.WriteString("docker-ce-")
//	//for i:=0;i< len(str);i++{
//	//	num, _ := strconv.Atoi(str[i])
//	//	buffer.WriteString(fmt.Sprintf("%02d", num))
//	//	buffer.WriteString(".")
//	//}
//	//
//	//fmt.Print(buffer.String()[0:len(buffer.String())-1])
//
//
//
//
//
//
//
//
//
//
//
//
//	//var kubeconfig *string
//	//if home := homedir.HomeDir(); home != "" {
//	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
//	//	fmt.Println("test done")
//	//} else {
//	//	fmt.Println("test upp")
//	//	kubeconfig = flag.String("kubeconfig", "", "config")
//	//}
//	//flag.Parse()
//	//
//	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//	//
//	//clientset, err := kubernetes.NewForConfig(config)
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//	//
//	//
//	//
//	//clientset.CoreV1().ComponentStatuses()
//	//
//	//nodes, err := clientset.CoreV1().Nodes().List(context.TODO(),metav1.ListOptions{})
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//for i := 0; i < len(nodes.Items); i++{
//	////	fmt.Println(nodes.Items[i].Name + string(nodes.Items[i].Status.Conditions[i].Type))
//	//
//	//	node := nodes.Items[i]
//	//	fmt.Printf("Name: %s \n", node.Name)
//	//
//	//	conditionMap := make(map[v1.NodeConditionType]*v1.NodeCondition)
//	//	NodeAllConditions := []v1.NodeConditionType{v1.NodeReady}
//	//	for i := range node.Status.Conditions {
//	//		cond := node.Status.Conditions[i]
//	//		conditionMap[cond.Type] = &cond
//	//	}
//	//	var status []string
//	//	for _, validCondition := range NodeAllConditions {
//	//		if condition, ok := conditionMap[validCondition]; ok {
//	//			if condition.Status == v1.ConditionTrue {
//	//				status = append(status, string(condition.Type))
//	//			} else {
//	//				status = append(status, "Not"+string(condition.Type))
//	//			}
//	//		}
//	//	}
//	//	if len(status) == 0 {
//	//		status = append(status, "Unknown")
//	//	}
//	//	if node.Spec.Unschedulable {
//	//		status = append(status, "SchedulingDisabled")
//	//	}
//	//	fmt.Printf("status: %s \n", strings.Join(status, ","))
//	//
//	//	roles := strings.Join(findNodeRoles(&node), ",")
//	//	if len(roles) == 0 {
//	//		roles = "<none>"
//	//	}
//	//	fmt.Printf("roles: %s %s %s \n", roles, translateTimestampSince(node.CreationTimestamp), node.Status.NodeInfo.KubeletVersion)
//	//
//	//
//	//
//	//}
//
//
//	//
//	//var s string
//	//s = ""
//	//fmt.Print(strings.TrimRight(s, ","))
//
//
//
//
//}
//
//func execShell(s string) string {
//	cmd := exec.Command("/bin/bash", "-c", s)
//	//	cmd := exec.Command("sh", "-c", s)
//	var out bytes.Buffer
//	cmd.Stdout = &out
//	_ = cmd.Run()
//	//	ch <- out.String()
//	return out.String()
//}
////
////func writeToKubeInit(path string, s string) {
////	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777) //linux 路径
////	if err != nil {
////		fmt.Printf("open err%s", err)
////		return
////	}
////	defer f.Close()
////	_, err2 := f.WriteString(s)
////	if err2 != nil {
////		fmt.Printf("write err:\n%s", err2)
////		return
////	}
////}
////
////
////
////func stringAppend(s1 string, s2 string) string {
////	var build strings.Builder
////	build.WriteString(s1)
////	build.WriteString(s2)
////	return build.String()
////}
////
////func Command(cmd string) error {
////	//c := exec.Command("cmd", "/C", cmd) 	// windows
////	c := exec.Command("bash", "-c", cmd)  // mac or linux
////	stdout, err := c.StdoutPipe()
////	if err != nil {
////		return err
////	}
////	var wg sync.WaitGroup
////	wg.Add(1)
////	go func() {
////		defer wg.Done()
////		reader := bufio.NewReader(stdout)
////		for {
////			_, err := reader.ReadString('\n')
////			if err != nil || err == io.EOF {
////				return
////			}
////
////
////		}
////	}()
////	err = c.Start()
////	wg.Wait()
////	return err
////}
//
////func translateTimestampSince(timestamp metav1.Time) string {
////	if timestamp.IsZero() {
////		return "<unknown>"
////	}
////
////	return duration.HumanDuration(time.Since(timestamp.Time))
////}
//
//
////var labelNodeRolePrefix = "node-role.kubernetes.io/"
////var nodeLabelRole = "kubernetes.io/role"
////func findNodeRoles(node *v1.Node) []string {
////	roles := sets.NewString()
////	for k, v := range node.Labels {
////		switch {
////		case strings.HasPrefix(k, labelNodeRolePrefix):
////			if role := strings.TrimPrefix(k, labelNodeRolePrefix); len(role) > 0 {
////				roles.Insert(role)
////			}
////
////		case k == nodeLabelRole && v != "":
////			roles.Insert(v)
////		}
////	}
////	return roles.List()
////}
//
////func isDigit(str string) bool {
////	for _, x := range []rune(str) {
////		if !unicode.IsDigit(x) {
////			return false
////		}
////	}
////	return true
////}
////
////
////func isDigitWithStorage(str string) bool {
////	if(isDigit(str[0:len(str)-2])) == false{
////		return false
////	}
////	if(str[len(str)-2:]) == "Gi" || (str[len(str)-2:]) == "Mi"{
////		return true
////	}
////	return false
////}
////
////func isDigitWithPercent(str string) bool {
////	if(isDigit(str[0:len(str)-1])) == false{
////		return false
////	}
////	if(str[len(str)-1:]) == "%"{
////		return true
////	}
////	return false
////}
////
////func readFile(s string){
////	f, err := os.Open(s)
////	if err != nil {
////		fmt.Println(err.Error())
////	}
////	buf := bufio.NewReader(f)
////	for {
////		//遇到\n结束读取
////		b, err := buf.ReadBytes('\n')
////		if err != nil {
////			if err == io.EOF {
////				break
////			}
////			fmt.Println(err.Error())
////		}
////		fmt.Println(string(b))
////	}
////}
