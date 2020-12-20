package network

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func SetLatency(t, device string) error {
	exist, err := LatenlyRuleExist(device)
	if err != nil {
		log.Fatalln("show qdisc error")
		return err
	}
	var cmd *exec.Cmd
	if exist {
		cmd = exec.Command("tc", "qdisc", "change", "dev", device, "root", "netem", "delay", t)
	} else {
		cmd = exec.Command("tc", "qdisc", "add", "dev", device, "root", "netem", "delay", t)
	}
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Println(fmt.Sprintf("err: tc args %v", cmd.Args))
		return err
	}
	log.Println("set tc rule success")
	return nil
}

func LatenlyRuleExist(device string) (bool, error) {
	cmd := exec.Command("tc", "qdisc", "show", "dev", device)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	if strings.HasPrefix(string(output), fmt.Sprintf("qdisc netem")) {
		return true, nil
	}
	return false, nil
}
