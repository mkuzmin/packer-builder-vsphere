package main

import (
	"github.com/mitchellh/multistep"
	"github.com/hashicorp/packer/packer"
	"fmt"
	"github.com/vmware/govmomi/vim25/mo"
	"context"
	"github.com/vmware/govmomi/object"
)

type StepCreateSnapshot struct{
	createSnapshot bool
	vmName         string
}

func (s *StepCreateSnapshot) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	vmSrc := state.Get("vmSrc").(*object.VirtualMachine)
	ctx := state.Get("ctx").(context.Context)

	if s.createSnapshot {
		ui.Say("creating snapshot...")

		_, err := vmSrc.CreateSnapshot(ctx, s.vmName, "", true, true)
		if err != nil {
			state.Put("error", err)
			return multistep.ActionHalt
		}
		//var info *types.TaskInfo
		//info, err = task.WaitForResult(ctx, nil)
		//if err != nil {
		//	state.Put("error", err)
		//	return multistep.ActionHalt
		//}
		//snapshot := info.Result.(types.ManagedObjectReference)
		ui.Say("done")
	}

	return multistep.ActionContinue
}

func (s *StepCreateSnapshot) Cleanup(state multistep.StateBag) {}
