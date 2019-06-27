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
// Implementation(s) of TBase
// ----------------------------------------------------

//TBaseOnBaseHolder is the struct containing the Base in order to implement TBase
type TBaseOnBaseHolder struct {
	h Base
}

//CreateTBaseForBase returns an holder of Base implementing TBase
func CreateTBaseForBase(o Base) TBaseOnBaseHolder {
	return TBaseOnBaseHolder{
		h: o,
	}
}

//URL returns the url where the base refers
func (r TBaseOnBaseHolder) URL() TURL {
	return CreateTURLForEkUrl(r.h.Url)
}
