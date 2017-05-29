package main
import (
	π___python__Γrandom "__python__/random"
	π___python__Γsys "__python__/sys"
	πg "grumpy"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ß__main__ := πg.InternStr("__main__")
	ß__name__ := πg.InternStr("__name__")
	ßargv := πg.InternStr("argv")
	ßc := πg.InternStr("c")
	ßi := πg.InternStr("i")
	ßint := πg.InternStr("int")
	ßnum := πg.InternStr("num")
	ßrandom := πg.InternStr("random")
	ßsys := πg.InternStr("sys")
	ßx := πg.InternStr("x")
	ßxrange := πg.InternStr("xrange")
	ßy := πg.InternStr("y")
	var πTemp001 *πg.Object
	_ = πTemp001
	var πTemp002 []*πg.Object
	_ = πTemp002
	var πTemp003 *πg.Object
	_ = πTemp003
	var πTemp004 bool
	_ = πTemp004
	var πTemp005 *πg.Object
	_ = πTemp005
	var πTemp006 *πg.Object
	_ = πTemp006
	var πTemp007 *πg.Object
	_ = πTemp007
	var πTemp008 *πg.Object
	_ = πTemp008
	var πTemp009 *πg.Object
	_ = πTemp009
	var πTemp010 *πg.Object
	_ = πTemp010
	var πTemp011 *πg.Object
	_ = πTemp011
	var πE *πg.BaseException; _ = πE
	for ; πF.State() >= 0; πF.PopCheckpoint() {
		switch πF.State() {
		case 0:
		default: panic("unexpected function state")
		}
		// line 3: import random
		πF.SetLineno(3)
		if πTemp002, πE = πg.ImportModule(πF, "random", []*πg.Code{π___python__Γrandom.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßrandom.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 4: import sys
		πF.SetLineno(4)
		if πTemp002, πE = πg.ImportModule(πF, "sys", []*πg.Code{π___python__Γsys.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
			continue
		}
		if πTemp003, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp001, πE = πg.Eq(πF, πTemp003, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
			continue
		}
		if πTemp004 {
			goto Label1
		}
		goto Label2
		// line 6: if __name__=="__main__":
		πF.SetLineno(6)
	Label1:
		// line 7: num = int(sys.argv[1])
		πF.SetLineno(7)
		πTemp002 = πF.MakeArgs(1)
		πTemp001 = πg.NewInt(1).ToObject()
		if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßargv, nil); πE != nil {
			continue
		}
		if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
			continue
		}
		πTemp002[0] = πTemp003
		if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
			continue
		}
		if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßnum.ToObject(), πTemp003); πE != nil {
			continue
		}
		// line 8: c = 0
		πF.SetLineno(8)
		if πE = πF.Globals().SetItem(πF, ßc.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
			continue
		}
		// line 10: for i in xrange(num):
		πF.SetLineno(10)
		πTemp002 = πF.MakeArgs(1)
		if πTemp001, πE = πg.ResolveGlobal(πF, ßnum); πE != nil {
			continue
		}
		πTemp002[0] = πTemp001
		if πTemp001, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
			continue
		}
		if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
			continue
		}
	Label3:
		if πTemp005, πE = πg.Next(πF, πTemp001); πE != nil {
			isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
			if exc != nil {
				πE = exc
				continue
			}
			if !isStop {
				continue
			}
			πE = nil
			πF.RestoreExc(nil, nil)
			goto Label4
		}
		if πE = πF.Globals().SetItem(πF, ßi.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 11: x = random.random()
		πF.SetLineno(11)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
			continue
		}
		if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßrandom, nil); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßx.ToObject(), πTemp006); πE != nil {
			continue
		}
		// line 12: y = random.random()
		πF.SetLineno(12)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
			continue
		}
		if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßrandom, nil); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßy.ToObject(), πTemp006); πE != nil {
			continue
		}
		if πTemp009, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
			continue
		}
		if πTemp010, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
			continue
		}
		if πTemp008, πE = πg.Mul(πF, πTemp009, πTemp010); πE != nil {
			continue
		}
		if πTemp010, πE = πg.ResolveGlobal(πF, ßy); πE != nil {
			continue
		}
		if πTemp011, πE = πg.ResolveGlobal(πF, ßy); πE != nil {
			continue
		}
		if πTemp009, πE = πg.Mul(πF, πTemp010, πTemp011); πE != nil {
			continue
		}
		if πTemp007, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
			continue
		}
		if πTemp006, πE = πg.LE(πF, πTemp007, πg.NewFloat(1.0).ToObject()); πE != nil {
			continue
		}
		if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
			continue
		}
		if πTemp004 {
			goto Label5
		}
		goto Label6
		// line 14: if x * x + y * y <= 1.0:
		πF.SetLineno(14)
	Label5:
		// line 15: c += 1
		πF.SetLineno(15)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßc); πE != nil {
			continue
		}
		if πTemp007, πE = πg.IAdd(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßc.ToObject(), πTemp007); πE != nil {
			continue
		}
		goto Label6
	Label6:
		goto Label3
		goto Label4
	Label4:
		// line 17: print 4.0*c/num
		πF.SetLineno(17)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp005, πE = πg.ResolveGlobal(πF, ßc); πE != nil {
			continue
		}
		if πTemp003, πE = πg.Mul(πF, πg.NewFloat(4.0).ToObject(), πTemp005); πE != nil {
			continue
		}
		if πTemp005, πE = πg.ResolveGlobal(πF, ßnum); πE != nil {
			continue
		}
		if πTemp001, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
			continue
		}
		πTemp002[0] = πTemp001
		if πE = πg.Print(πF, πTemp002, true); πE != nil {
			continue
		}
		goto Label2
	Label2:
		return nil, nil
	}
	return nil, πE
}
var Code *πg.Code
func main() {
	Code = πg.NewCode("<module>", "monte.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
