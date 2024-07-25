package impl

import (
	"fmt"
	"mikea/declix/interfaces"
	"mikea/declix/resources"

	_ "mikea/declix/resources/apt"
	_ "mikea/declix/resources/dpkg"
	_ "mikea/declix/resources/filesystem"
	_ "mikea/declix/resources/systemd"
	_ "mikea/declix/resources/users"
)

func CreateResource(r resources.Resource) interfaces.Resource {
	if res, ok := r.(interfaces.Resource); ok {
		return res
	} else {
		panic(fmt.Sprintf("%#v does not implement interfaces.Resource", r))
	}
}
