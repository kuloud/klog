package main

import (
	log "github.com/kuloud/klog"
)

func main() {
	log.V("vTag", "hello", "kuloud")
	log.Vf("vfTag", "%s, just like %s said.", "hello", "kuloud")
	log.D("dTag", "hello", "kuloud")
	log.Df("dfTag", "%s, just like %s said. %d", "hello", "kuloud", 1)
	log.I("iTag", "hello", "kuloud")
	log.If("ifTag", "%s, just like %s said. %t", "hello", "kuloud", true)
	log.W("wTag", "hello", "kuloud")
	log.Wf("wfTag", "%s, just like %s said. %s", "hello", "kuloud", "Debug")
	log.E("eTag", "hello", "kuloud")
	log.Ef("efTag", "%s, just like you said.")
}
