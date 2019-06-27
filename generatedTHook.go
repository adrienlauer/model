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
// Implementation(s) of THook
// ----------------------------------------------------

//THookOnHookHolder is the struct containing the Hook in order to implement THook
type THookOnHookHolder struct {
	h Hook
}

//CreateTHookForHook returns an holder of Hook implementing THook
func CreateTHookForHook(o Hook) THookOnHookHolder {
	return THookOnHookHolder{
		h: o,
	}
}

//After returns the references of tasks to run after the hooked action
func (r THookOnHookHolder) After() []TTaskRef {
	result := make([]TTaskRef, 0, 0)
	for _, val := range r.h.After {
		result = append(result, CreateTTaskRefForTaskRef(val))
	}
	return result

}

//Before returns the references of tasks to run before the hooked action
func (r THookOnHookHolder) Before() []TTaskRef {
	result := make([]TTaskRef, 0, 0)
	for _, val := range r.h.Before {
		result = append(result, CreateTTaskRefForTaskRef(val))
	}
	return result

}
