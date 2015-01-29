package instructions

import (
    //"log"
    //. "jvmgo/any"
    //"jvmgo/native"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Invoke instance method; dispatch based on class
type invokevirtual struct {Index16Instruction}
func (self *invokevirtual) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*rtc.ConstantMethodref)
    ref := stack.Top(cMethodRef.ArgCount())
    if ref == nil {
        panic("NPE")
    }

    method := cMethodRef.VirtualMethod(ref.(*rtc.Obj))
    if method.IsNative() {
        nativeMethod := method.NativeMethod().(func(*rtda.Frame))
        nativeMethod(frame)
    } else {
        newFrame := thread.NewFrame(method)
        thread.PushFrame(newFrame)
        passArgs(stack, newFrame.LocalVars(), method.ArgCount() + 1)
    }
}
