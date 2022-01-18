package runner

// handle the result
func (r *Runner) HandleResult() (err error, rets []JSONResult) {

	for hostIP, ports := range r.scanner.ScanResults.IPPorts {
		dt, err := r.scanner.IPRanger.GetHostsByIP(hostIP)
		if err != nil {
			continue
		}

		for _, host := range dt {
			if host == "ip" {
				host = hostIP
			}
			// gologger.Info().Msgf("Found %d ports on host %s (%s)\n", len(ports), host, hostIP)

			data := JSONResult{IP: hostIP}
			if host != hostIP {
				data.Host = host
			}
			for port := range ports {
				data.Port = port

				rets = append(rets, data)

				//b, marshallErr := json.Marshal(data)
				//if marshallErr != nil {
				//	continue
				//}
				// gologger.Silent().Msgf("%s\n", string(b))
			}
		}

	}
	return
}
