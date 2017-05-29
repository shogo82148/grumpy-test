package main
import (
	π___python__Γsys "__python__/sys"
	πg "grumpy"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ß__main__ := πg.InternStr("__main__")
	ß__name__ := πg.InternStr("__name__")
	ßa := πg.InternStr("a")
	ßargv := πg.InternStr("argv")
	ßb := πg.InternStr("b")
	ßc := πg.InternStr("c")
	ßint := πg.InternStr("int")
	ßsys := πg.InternStr("sys")
	ßtak := πg.InternStr("tak")
	var πTemp001 *πg.Object
	_ = πTemp001
	var πTemp002 []*πg.Object
	_ = πTemp002
	var πTemp003 []πg.Param
	_ = πTemp003
	var πTemp004 *πg.Object
	_ = πTemp004
	var πTemp005 *πg.Object
	_ = πTemp005
	var πTemp006 bool
	_ = πTemp006
	var πTemp007 *πg.Object
	_ = πTemp007
	var πTemp008 *πg.Object
	_ = πTemp008
	var πTemp009 []*πg.Object
	_ = πTemp009
	var πE *πg.BaseException; _ = πE
	for ; πF.State() >= 0; πF.PopCheckpoint() {
		switch πF.State() {
		case 0:
		default: panic("unexpected function state")
		}
		// line 3: import sys
		πF.SetLineno(3)
		if πTemp002, πE = πg.ImportModule(πF, "sys", []*πg.Code{π___python__Γsys.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 5: def tak(x, y, z):
		πF.SetLineno(5)
		πTemp003 = make([]πg.Param, 3)
		πTemp003[0] = πg.Param{Name: "x", Def: nil}
		πTemp003[1] = πg.Param{Name: "y", Def: nil}
		πTemp003[2] = πg.Param{Name: "z", Def: nil}
		πTemp001 = πg.NewFunction(πg.NewCode("tak", "take.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µx *πg.Object = πArgs[0]; _ = µx
			var µy *πg.Object = πArgs[1]; _ = µy
			var µz *πg.Object = πArgs[2]; _ = µz
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 bool
			_ = πTemp002
			var πTemp003 []*πg.Object
			_ = πTemp003
			var πTemp004 []*πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.LE(πF, µx, µy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
					continue
				}
				if πTemp002 {
					goto Label1
				}
				goto Label2
				// line 6: if x <= y:
				πF.SetLineno(6)
			Label1:
				// line 7: return y
				πF.SetLineno(7)
				if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
					continue
				}
				return µy, nil
				goto Label3
			Label2:
				// line 9: return tak(tak((x-1), y , z), tak((y-1), z , x), tak((z-1) , x, y))
				πF.SetLineno(9)
				πTemp003 = πF.MakeArgs(3)
				πTemp004 = πF.MakeArgs(3)
				if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.Sub(πF, µx, πg.NewInt(1).ToObject()); πE != nil {
					continue
				}
				πTemp004[0] = πTemp001
				if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
					continue
				}
				πTemp004[1] = µy
				if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
					continue
				}
				πTemp004[2] = µz
				if πTemp001, πE = πg.ResolveGlobal(πF, ßtak); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[0] = πTemp005
				πTemp004 = πF.MakeArgs(3)
				if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.Sub(πF, µy, πg.NewInt(1).ToObject()); πE != nil {
					continue
				}
				πTemp004[0] = πTemp001
				if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
					continue
				}
				πTemp004[1] = µz
				if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
					continue
				}
				πTemp004[2] = µx
				if πTemp001, πE = πg.ResolveGlobal(πF, ßtak); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[1] = πTemp005
				πTemp004 = πF.MakeArgs(3)
				if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.Sub(πF, µz, πg.NewInt(1).ToObject()); πE != nil {
					continue
				}
				πTemp004[0] = πTemp001
				if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
					continue
				}
				πTemp004[1] = µx
				if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
					continue
				}
				πTemp004[2] = µy
				if πTemp001, πE = πg.ResolveGlobal(πF, ßtak); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[2] = πTemp005
				if πTemp001, πE = πg.ResolveGlobal(πF, ßtak); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				return πTemp005, nil
				goto Label3
			Label3:
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßtak.ToObject(), πTemp001); πE != nil {
			continue
		}
		if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp004, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
			continue
		}
		if πTemp006 {
			goto Label1
		}
		goto Label2
		// line 11: if __name__=="__main__":
		πF.SetLineno(11)
	Label1:
		// line 12: a = int(sys.argv[1])
		πF.SetLineno(12)
		πTemp002 = πF.MakeArgs(1)
		πTemp004 = πg.NewInt(1).ToObject()
		if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßargv, nil); πE != nil {
			continue
		}
		if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
			continue
		}
		πTemp002[0] = πTemp005
		if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
			continue
		}
		if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßa.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 13: b = int(sys.argv[2])
		πF.SetLineno(13)
		πTemp002 = πF.MakeArgs(1)
		πTemp004 = πg.NewInt(2).ToObject()
		if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßargv, nil); πE != nil {
			continue
		}
		if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
			continue
		}
		πTemp002[0] = πTemp005
		if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
			continue
		}
		if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßb.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 14: c = int(sys.argv[3])
		πF.SetLineno(14)
		πTemp002 = πF.MakeArgs(1)
		πTemp004 = πg.NewInt(3).ToObject()
		if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßargv, nil); πE != nil {
			continue
		}
		if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
			continue
		}
		πTemp002[0] = πTemp005
		if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
			continue
		}
		if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßc.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 15: print tak(a, b, c)
		πF.SetLineno(15)
		πTemp002 = make([]*πg.Object, 1)
		πTemp009 = πF.MakeArgs(3)
		if πTemp004, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
			continue
		}
		πTemp009[0] = πTemp004
		if πTemp004, πE = πg.ResolveGlobal(πF, ßb); πE != nil {
			continue
		}
		πTemp009[1] = πTemp004
		if πTemp004, πE = πg.ResolveGlobal(πF, ßc); πE != nil {
			continue
		}
		πTemp009[2] = πTemp004
		if πTemp004, πE = πg.ResolveGlobal(πF, ßtak); πE != nil {
			continue
		}
		if πTemp005, πE = πTemp004.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		πTemp002[0] = πTemp005
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
	Code = πg.NewCode("<module>", "take.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
