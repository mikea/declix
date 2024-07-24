package filesystem

import "mikea/declix/interfaces"

var _ interfaces.Resource = &FileImpl{}
var _ interfaces.Resource = &DirImpl{}
