package model

//*****************************************************************************
//
//     _         _           ____                           _           _
//    / \  _   _| |_ ___    / ___| ___ _ __   ___ _ __ __ _| |_ ___  __| |
//   / _ \| | | | __/ _ \  | |  _ / _ \ '_ \ / _ \ '__/ _` | __/ _ \/ _` |
//  / ___ \ |_| | || (_) | | |_| |  __/ | | |  __/ | | (_| | ||  __/ (_| |
// /_/   \_\__,_|\__\___/   \____|\___|_| |_|\___|_|  \__,_|\__\___|\__,_|
//
// This file is autogenerated by "go generate .". Do not modify.
//
//*****************************************************************************

// ----------------------------------------------------
// Implementation(s) of TNodeSet
// ----------------------------------------------------

//TNodeSetOnNodeSetHolder is the struct containing the NodeSet in order to implement TNodeSet
type TNodeSetOnNodeSetHolder struct {
	h NodeSet
}

//CreateTNodeSetForNodeSet returns an holder of NodeSet implementing TNodeSet
func CreateTNodeSetForNodeSet(o NodeSet) TNodeSetOnNodeSetHolder {
	return TNodeSetOnNodeSetHolder{
		h: o,
	}
}

//Name returns the name of the node set
func (r TNodeSetOnNodeSetHolder) Name() string {
	return r.h.Name
}

//Instances returns the number of nodes to create for this node set
func (r TNodeSetOnNodeSetHolder) Instances() int {
	return r.h.Instances
}

//Orchestrator returns the reference on the orchestrator managing the node
func (r TNodeSetOnNodeSetHolder) Orchestrator() TOrchestratorRef {
	return CreateTOrchestratorRefForOrchestratorRef(r.h.Orchestrator)
}

//Provider returns the reference on the provider wherein the node should be deployed
func (r TNodeSetOnNodeSetHolder) Provider() TProviderRef {
	return CreateTProviderRefForProviderRef(r.h.Provider)
}

//HasHooks returns true if the node has hooks
func (r TNodeSetOnNodeSetHolder) HasHooks() bool {
	return r.h.Hooks.HasTasks()
}

//Hooks returns the node hooks
func (r TNodeSetOnNodeSetHolder) Hooks() TNodeHook {
	return CreateTNodeHookForNodeHook(r.h.Hooks)
}

//HasProvisionHooks returns true if the node has hooks while provisioning
func (r TNodeSetOnNodeSetHolder) HasProvisionHooks() bool {
	return r.h.Hooks.Provision.HasTasks()
}

//HasDestroyHooks returns true if the node has hooks while destroying
func (r TNodeSetOnNodeSetHolder) HasDestroyHooks() bool {
	return r.h.Hooks.Destroy.HasTasks()
}

//HasLabels returns true if the node has defined labels
func (r TNodeSetOnNodeSetHolder) HasLabels() bool {
	return len(r.h.Labels) > 0
}

//Labels returns the node labels
func (r TNodeSetOnNodeSetHolder) Labels() map[string]string {
	return r.h.Labels
}
