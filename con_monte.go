package main
import (
	π___python__Γrandom "__python__/random"
	π___python__Γsys "__python__/sys"
	π___python__Γthreading "__python__/threading"
	πg "grumpy"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ßMyThread := πg.InternStr("MyThread")
	ßRandom := πg.InternStr("Random")
	ßThread := πg.InternStr("Thread")
	ß__init__ := πg.InternStr("__init__")
	ß__main__ := πg.InternStr("__main__")
	ß__metaclass__ := πg.InternStr("__metaclass__")
	ß__module__ := πg.InternStr("__module__")
	ß__name__ := πg.InternStr("__name__")
	ßappend := πg.InternStr("append")
	ßargv := πg.InternStr("argv")
	ßc := πg.InternStr("c")
	ßi := πg.InternStr("i")
	ßint := πg.InternStr("int")
	ßjoin := πg.InternStr("join")
	ßnum := πg.InternStr("num")
	ßpara := πg.InternStr("para")
	ßrandom := πg.InternStr("random")
	ßrun := πg.InternStr("run")
	ßstart := πg.InternStr("start")
	ßsuper := πg.InternStr("super")
	ßsys := πg.InternStr("sys")
	ßt := πg.InternStr("t")
	ßthreading := πg.InternStr("threading")
	ßthreads := πg.InternStr("threads")
	ßxrange := πg.InternStr("xrange")
	var πTemp001 *πg.Object
	_ = πTemp001
	var πTemp002 []*πg.Object
	_ = πTemp002
	var πTemp003 *πg.Dict
	_ = πTemp003
	var πTemp004 *πg.Object
	_ = πTemp004
	var πTemp005 *πg.Object
	_ = πTemp005
	var πTemp006 *πg.Object
	_ = πTemp006
	var πTemp007 bool
	_ = πTemp007
	var πTemp008 *πg.Object
	_ = πTemp008
	var πTemp009 *πg.Object
	_ = πTemp009
	var πE *πg.BaseException; _ = πE
	for ; πF.State() >= 0; πF.PopCheckpoint() {
		switch πF.State() {
		case 0:
		default: panic("unexpected function state")
		}
		// line 3: import threading
		πF.SetLineno(3)
		if πTemp002, πE = πg.ImportModule(πF, "threading", []*πg.Code{π___python__Γthreading.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 4: import random
		πF.SetLineno(4)
		if πTemp002, πE = πg.ImportModule(πF, "random", []*πg.Code{π___python__Γrandom.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßrandom.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 5: import sys
		πF.SetLineno(5)
		if πTemp002, πE = πg.ImportModule(πF, "sys", []*πg.Code{π___python__Γsys.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 7: class MyThread(threading.Thread):
		πF.SetLineno(7)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp005, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
			continue
		}
		if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßThread, nil); πE != nil {
			continue
		}
		πTemp002[0] = πTemp006
		πTemp003 = πg.NewDict()
		if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
			continue
		}
		_, πE = πg.NewCode("MyThread", "con_monte.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp003
			_ = πClass
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 []πg.Param
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 8: def __init__(self):
				πF.SetLineno(8)
				πTemp002 = make([]πg.Param, 1)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp001 = πg.NewFunction(πg.NewCode("__init__", "con_monte.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 9: super(MyThread, self).__init__()
						πF.SetLineno(9)
						πTemp001 = πF.MakeArgs(2)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßMyThread); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						πTemp001[1] = µself
						if πTemp002, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						// line 10: self.c = 0
						πF.SetLineno(10)
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßc, πTemp002); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 12: def run(self):
				πF.SetLineno(12)
				πTemp002 = make([]πg.Param, 1)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp003 = πg.NewFunction(πg.NewCode("run", "con_monte.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µr *πg.Object = πg.UnboundLocal; _ = µr
					var µc *πg.Object = πg.UnboundLocal; _ = µc
					var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
					var µx *πg.Object = πg.UnboundLocal; _ = µx
					var µy *πg.Object = πg.UnboundLocal; _ = µy
					var πTemp001 *πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 []*πg.Object
					_ = πTemp003
					var πTemp004 *πg.Object
					_ = πTemp004
					var πTemp005 *πg.Object
					_ = πTemp005
					var πTemp006 *πg.Object
					_ = πTemp006
					var πTemp007 *πg.Object
					_ = πTemp007
					var πTemp008 *πg.Object
					_ = πTemp008
					var πTemp009 bool
					_ = πTemp009
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 13: r = random.Random()
						πF.SetLineno(13)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßRandom, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						µr = πTemp001
						// line 14: c = 0
						πF.SetLineno(14)
						µc = πg.NewInt(0).ToObject()
						// line 15: for _ in xrange(num):
						πF.SetLineno(15)
						πTemp003 = πF.MakeArgs(1)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßnum); πE != nil {
							continue
						}
						πTemp003[0] = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp003)
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
							continue
						}
					Label1:
						if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
							goto Label2
						}
						µ_ = πTemp004
						// line 16: x = r.random()
						πF.SetLineno(16)
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µr, ßrandom, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						µx = πTemp006
						// line 17: y = r.random()
						πF.SetLineno(17)
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µr, ßrandom, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						µy = πTemp006
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.Mul(πF, µx, µx); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.Mul(πF, µy, µy); πE != nil {
							continue
						}
						if πTemp006, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
							continue
						}
						if πTemp005, πE = πg.LE(πF, πTemp006, πg.NewFloat(1.0).ToObject()); πE != nil {
							continue
						}
						if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
							continue
						}
						if πTemp009 {
							goto Label3
						}
						goto Label4
						// line 19: if x * x + y * y <= 1.0:
						πF.SetLineno(19)
					Label3:
						// line 20: c += 1
						πF.SetLineno(20)
						if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.IAdd(πF, µc, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µc = πTemp005
						goto Label4
					Label4:
						goto Label1
						goto Label2
					Label2:
						// line 21: self.c = c
						πF.SetLineno(21)
						if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µc); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßc, πTemp001); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp003); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			continue
		}
		if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp004 == nil {
			πTemp004 = πg.TypeType.ToObject()
		}
		if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MyThread").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßMyThread.ToObject(), πTemp005); πE != nil {
			continue
		}
		if πTemp004, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp001, πE = πg.Eq(πF, πTemp004, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
			continue
		}
		if πTemp007 {
			goto Label1
		}
		goto Label2
		// line 24: if __name__ == "__main__":
		πF.SetLineno(24)
	Label1:
		// line 25: num = int(sys.argv[1])
		πF.SetLineno(25)
		πTemp002 = πF.MakeArgs(1)
		πTemp001 = πg.NewInt(1).ToObject()
		if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßargv, nil); πE != nil {
			continue
		}
		if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
			continue
		}
		πTemp002[0] = πTemp004
		if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
			continue
		}
		if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßnum.ToObject(), πTemp004); πE != nil {
			continue
		}
		// line 26: para = int(sys.argv[2])
		πF.SetLineno(26)
		πTemp002 = πF.MakeArgs(1)
		πTemp001 = πg.NewInt(2).ToObject()
		if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßargv, nil); πE != nil {
			continue
		}
		if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
			continue
		}
		πTemp002[0] = πTemp004
		if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
			continue
		}
		if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßpara.ToObject(), πTemp004); πE != nil {
			continue
		}
		// line 28: threads = []
		πF.SetLineno(28)
		πTemp002 = make([]*πg.Object, 0)
		πTemp001 = πg.NewList(πTemp002...).ToObject()
		if πE = πF.Globals().SetItem(πF, ßthreads.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 29: for i in xrange(para):
		πF.SetLineno(29)
		πTemp002 = πF.MakeArgs(1)
		if πTemp001, πE = πg.ResolveGlobal(πF, ßpara); πE != nil {
			continue
		}
		πTemp002[0] = πTemp001
		if πTemp001, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
			continue
		}
		if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
		// line 30: t = MyThread()
		πF.SetLineno(30)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßMyThread); πE != nil {
			continue
		}
		if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßt.ToObject(), πTemp008); πE != nil {
			continue
		}
		// line 31: t.start()
		πF.SetLineno(31)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
			continue
		}
		if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßstart, nil); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp008.Call(πF, nil, nil); πE != nil {
			continue
		}
		// line 32: threads.append(t)
		πF.SetLineno(32)
		πTemp002 = πF.MakeArgs(1)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
			continue
		}
		πTemp002[0] = πTemp006
		if πTemp006, πE = πg.ResolveGlobal(πF, ßthreads); πE != nil {
			continue
		}
		if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßappend, nil); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		goto Label3
		goto Label4
	Label4:
		// line 34: c = 0
		πF.SetLineno(34)
		if πE = πF.Globals().SetItem(πF, ßc.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
			continue
		}
		// line 35: for t in threads:
		πF.SetLineno(35)
		if πTemp001, πE = πg.ResolveGlobal(πF, ßthreads); πE != nil {
			continue
		}
		if πTemp004, πE = πg.Iter(πF, πTemp001); πE != nil {
			continue
		}
	Label5:
		if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
			goto Label6
		}
		if πE = πF.Globals().SetItem(πF, ßt.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 36: t.join()
		πF.SetLineno(36)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
			continue
		}
		if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßjoin, nil); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp008.Call(πF, nil, nil); πE != nil {
			continue
		}
		// line 37: c += t.c
		πF.SetLineno(37)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßc); πE != nil {
			continue
		}
		if πTemp008, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
			continue
		}
		if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßc, nil); πE != nil {
			continue
		}
		if πTemp008, πE = πg.IAdd(πF, πTemp006, πTemp009); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßc.ToObject(), πTemp008); πE != nil {
			continue
		}
		goto Label5
		goto Label6
	Label6:
		// line 39: print 4.0*c/(num*para)
		πF.SetLineno(39)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp005, πE = πg.ResolveGlobal(πF, ßc); πE != nil {
			continue
		}
		if πTemp004, πE = πg.Mul(πF, πg.NewFloat(4.0).ToObject(), πTemp005); πE != nil {
			continue
		}
		if πTemp006, πE = πg.ResolveGlobal(πF, ßnum); πE != nil {
			continue
		}
		if πTemp008, πE = πg.ResolveGlobal(πF, ßpara); πE != nil {
			continue
		}
		if πTemp005, πE = πg.Mul(πF, πTemp006, πTemp008); πE != nil {
			continue
		}
		if πTemp001, πE = πg.Div(πF, πTemp004, πTemp005); πE != nil {
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
	Code = πg.NewCode("<module>", "con_monte.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
