package impl

import (
	"fmt"
	"mikea/declix/impl/apt_package"
	"mikea/declix/impl/file"
	"mikea/declix/interfaces"
	"mikea/declix/pkl"
)

func CreateResource(expected pkl.Resource) interfaces.Resource {
	switch v := expected.(type) {
	case pkl.File:
		return file.New(v)
	case pkl.Package:
		return apt_package.New(v)
	default:
		panic(fmt.Sprintf("unexpected pkl.Resource: %#v", v))
	}
}
