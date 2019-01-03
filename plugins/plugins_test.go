package plugins_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Microsoft/windows-container-networking/plugins"
	cniSkel "github.com/containernetworking/cni/pkg/skel"
)

func generateBaseCmdArgs() skel.CmdArgs {
	f := skel.CmdArgs{}
	return f
}

var _ = Describe("Plugins", func() {

})
