package naabu_plus

import (
	"github.com/deenrookie/naabu-plus/pkg/runner"
	"github.com/projectdiscovery/gologger"
)

func PortScan(host string, threads int, ports string) (rets []runner.JSONResult, err error) {
	options := runner.ParseOptions()

	options.Host = host
	options.Threads = threads
	options.TopPorts = ports

	naabuRunner, err := runner.NewRunner(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	err, rets = naabuRunner.RunEnumeration()
	return
}

//func main() {
//	_, data := PortScan("dns.d33n.cn", 10, "top-1000")
//	fmt.Println(data)
//}
