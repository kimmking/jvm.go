package awt

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_cgl(cgl_initCGL, "initCGL", "()Z")
}

func _cgl(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/java2d/opengl/CGLGraphicsConfig", name, desc, method)
}

func cgl_initCGL(frame *rtda.Frame) {
	//TODO
	frame.PushBoolean(true)
}
