package execcmd

import (
	"io/ioutil"
	"k8s.io/klog/v2"
	"os/exec"
)

func ExecCommand(strCommand string) string {
	cmd := exec.Command("/bin/bash", "-c", strCommand)

	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		klog.Errorf("Execute failed when Start: %s", err.Error())
		return ""
	}

	outBytes, _ := ioutil.ReadAll(stdout)
	err := stdout.Close()
	if err != nil {
		klog.Errorf("stdout.Close() error: %s", err.Error())
		return ""
	}

	if err := cmd.Wait(); err != nil {
		klog.Errorf("Execute failed when Wait: %s", err.Error())
		return ""
	}
	return string(outBytes)
}
