package main
import (
	π___python__Γrandom "__python__/random"
	π___python__Γsys "__python__/sys"
	π___python__Γthreading "__python__/threading"
	πg "grumpy"
	π_mathΓrand "math/rand"
	π_os "os"
	π_reflect "reflect"
	π_time "time"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ßFloat64 := πg.InternStr("Float64")
	ßMyThread := πg.InternStr("MyThread")
	ßNew := πg.InternStr("New")
	ßNewSource := πg.InternStr("NewSource")
	ßNow := πg.InternStr("Now")
	ßThread := πg.InternStr("Thread")
	ßUnixNano := πg.InternStr("UnixNano")
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
	var πTemp003 map[string]*πg.Object
	_ = πTemp003
	var πTemp004 *πg.Object
	_ = πTemp004
	var πTemp005 *πg.Dict
	_ = πTemp005
	var πTemp006 *πg.Object
	_ = πTemp006
	var πTemp007 *πg.Object
	_ = πTemp007
	var πTemp008 bool
	_ = πTemp008
	var πTemp009 *πg.Object
	_ = πTemp009
	var πTemp010 *πg.Object
	_ = πTemp010
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
		// line 6: from __go__.time import Now
		πF.SetLineno(6)
		πTemp003 = map[string]*πg.Object{}
		if πTemp004, πE = πg.WrapNative(πF, π_reflect.ValueOf(π_time.Now)); πE != nil {
			continue
		}
		πTemp003["Now"] = πTemp004
		if πTemp001, πE = πg.ImportNativeModule(πF, "time", πTemp003); πE != nil {
			continue
		}
		if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßNow, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßNow.ToObject(), πTemp004); πE != nil {
			continue
		}
		// line 7: from __go__.math.rand import New, NewSource
		πF.SetLineno(7)
		πTemp003 = map[string]*πg.Object{}
		if πTemp004, πE = πg.WrapNative(πF, π_reflect.ValueOf(π_mathΓrand.New)); πE != nil {
			continue
		}
		πTemp003["New"] = πTemp004
		if πTemp004, πE = πg.WrapNative(πF, π_reflect.ValueOf(π_mathΓrand.NewSource)); πE != nil {
			continue
		}
		πTemp003["NewSource"] = πTemp004
		if πTemp001, πE = πg.ImportNativeModule(πF, "math.rand", πTemp003); πE != nil {
			continue
		}
		if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßNew, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßNew.ToObject(), πTemp004); πE != nil {
			continue
		}
		if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßNewSource, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßNewSource.ToObject(), πTemp004); πE != nil {
			continue
		}
		// line 9: class MyThread(threading.Thread):
		πF.SetLineno(9)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
			continue
		}
		if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßThread, nil); πE != nil {
			continue
		}
		πTemp002[0] = πTemp007
		πTemp005 = πg.NewDict()
		if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
			continue
		}
		_, πE = πg.NewCode("MyThread", "con_monte_go.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp005
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
				// line 10: def __init__(self):
				πF.SetLineno(10)
				πTemp002 = make([]πg.Param, 1)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp001 = πg.NewFunction(πg.NewCode("__init__", "con_monte_go.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						// line 11: super(MyThread, self).__init__()
						πF.SetLineno(11)
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
						// line 12: self.c = 0
						πF.SetLineno(12)
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
				// line 14: def run(self):
				πF.SetLineno(14)
				πTemp002 = make([]πg.Param, 1)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp003 = πg.NewFunction(πg.NewCode("run", "con_monte_go.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µr *πg.Object = πg.UnboundLocal; _ = µr
					var µc *πg.Object = πg.UnboundLocal; _ = µc
					var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
					var µx *πg.Object = πg.UnboundLocal; _ = µx
					var µy *πg.Object = πg.UnboundLocal; _ = µy
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 []*πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
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
					var πTemp009 *πg.Object
					_ = πTemp009
					var πTemp010 bool
					_ = πTemp010
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 15: r = New(NewSource(Now().UnixNano()))
						πF.SetLineno(15)
						πTemp001 = πF.MakeArgs(1)
						πTemp002 = πF.MakeArgs(1)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßNow); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßUnixNano, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp002[0] = πTemp004
						if πTemp003, πE = πg.ResolveGlobal(πF, ßNewSource); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp001[0] = πTemp004
						if πTemp003, πE = πg.ResolveGlobal(πF, ßNew); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µr = πTemp004
						// line 16: c = 0
						πF.SetLineno(16)
						µc = πg.NewInt(0).ToObject()
						// line 17: for _ in xrange(num):
						πF.SetLineno(17)
						πTemp001 = πF.MakeArgs(1)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßnum); πE != nil {
							continue
						}
						πTemp001[0] = πTemp003
						if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
							continue
						}
					Label1:
						if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µ_ = πTemp005
						// line 18: x = r.Float64()
						πF.SetLineno(18)
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, µr, ßFloat64, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
							continue
						}
						µx = πTemp007
						// line 19: y = r.Float64()
						πF.SetLineno(19)
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, µr, ßFloat64, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
							continue
						}
						µy = πTemp007
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.Mul(πF, µx, µx); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.Mul(πF, µy, µy); πE != nil {
							continue
						}
						if πTemp007, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
							continue
						}
						if πTemp006, πE = πg.LE(πF, πTemp007, πg.NewFloat(1.0).ToObject()); πE != nil {
							continue
						}
						if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if πTemp010 {
							goto Label3
						}
						goto Label4
						// line 21: if x * x + y * y <= 1.0:
						πF.SetLineno(21)
					Label3:
						// line 22: c += 1
						πF.SetLineno(22)
						if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.IAdd(πF, µc, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µc = πTemp006
						goto Label4
					Label4:
						goto Label1
						goto Label2
					Label2:
						// line 23: self.c = c
						πF.SetLineno(23)
						if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µc); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßc, πTemp003); πE != nil {
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
		if πTemp004, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp004 == nil {
			πTemp004 = πg.TypeType.ToObject()
		}
		if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MyThread").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßMyThread.ToObject(), πTemp006); πE != nil {
			continue
		}
		if πTemp004, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp001, πE = πg.Eq(πF, πTemp004, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
			continue
		}
		if πTemp008 {
			goto Label1
		}
		goto Label2
		// line 26: if __name__ == "__main__":
		πF.SetLineno(26)
	Label1:
		// line 27: num = int(sys.argv[1])
		πF.SetLineno(27)
		πTemp002 = πF.MakeArgs(1)
		πTemp001 = πg.NewInt(1).ToObject()
		if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßargv, nil); πE != nil {
			continue
		}
		if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
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
		// line 28: para = int(sys.argv[2])
		πF.SetLineno(28)
		πTemp002 = πF.MakeArgs(1)
		πTemp001 = πg.NewInt(2).ToObject()
		if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßargv, nil); πE != nil {
			continue
		}
		if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
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
		// line 30: threads = []
		πF.SetLineno(30)
		πTemp002 = make([]*πg.Object, 0)
		πTemp001 = πg.NewList(πTemp002...).ToObject()
		if πE = πF.Globals().SetItem(πF, ßthreads.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 31: for i in xrange(para):
		πF.SetLineno(31)
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
		if πTemp006, πE = πg.Next(πF, πTemp001); πE != nil {
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
		if πE = πF.Globals().SetItem(πF, ßi.ToObject(), πTemp006); πE != nil {
			continue
		}
		// line 32: t = MyThread()
		πF.SetLineno(32)
		if πTemp007, πE = πg.ResolveGlobal(πF, ßMyThread); πE != nil {
			continue
		}
		if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßt.ToObject(), πTemp009); πE != nil {
			continue
		}
		// line 33: t.start()
		πF.SetLineno(33)
		if πTemp007, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
			continue
		}
		if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßstart, nil); πE != nil {
			continue
		}
		if πTemp007, πE = πTemp009.Call(πF, nil, nil); πE != nil {
			continue
		}
		// line 34: threads.append(t)
		πF.SetLineno(34)
		πTemp002 = πF.MakeArgs(1)
		if πTemp007, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
			continue
		}
		πTemp002[0] = πTemp007
		if πTemp007, πE = πg.ResolveGlobal(πF, ßthreads); πE != nil {
			continue
		}
		if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßappend, nil); πE != nil {
			continue
		}
		if πTemp007, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		goto Label3
		goto Label4
	Label4:
		// line 36: c = 0
		πF.SetLineno(36)
		if πE = πF.Globals().SetItem(πF, ßc.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
			continue
		}
		// line 37: for t in threads:
		πF.SetLineno(37)
		if πTemp001, πE = πg.ResolveGlobal(πF, ßthreads); πE != nil {
			continue
		}
		if πTemp004, πE = πg.Iter(πF, πTemp001); πE != nil {
			continue
		}
	Label5:
		if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
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
		if πE = πF.Globals().SetItem(πF, ßt.ToObject(), πTemp006); πE != nil {
			continue
		}
		// line 38: t.join()
		πF.SetLineno(38)
		if πTemp007, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
			continue
		}
		if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßjoin, nil); πE != nil {
			continue
		}
		if πTemp007, πE = πTemp009.Call(πF, nil, nil); πE != nil {
			continue
		}
		// line 39: c += t.c
		πF.SetLineno(39)
		if πTemp007, πE = πg.ResolveGlobal(πF, ßc); πE != nil {
			continue
		}
		if πTemp009, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
			continue
		}
		if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßc, nil); πE != nil {
			continue
		}
		if πTemp009, πE = πg.IAdd(πF, πTemp007, πTemp010); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßc.ToObject(), πTemp009); πE != nil {
			continue
		}
		goto Label5
		goto Label6
	Label6:
		// line 41: print 4.0*c/(num*para)
		πF.SetLineno(41)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßc); πE != nil {
			continue
		}
		if πTemp004, πE = πg.Mul(πF, πg.NewFloat(4.0).ToObject(), πTemp006); πE != nil {
			continue
		}
		if πTemp007, πE = πg.ResolveGlobal(πF, ßnum); πE != nil {
			continue
		}
		if πTemp009, πE = πg.ResolveGlobal(πF, ßpara); πE != nil {
			continue
		}
		if πTemp006, πE = πg.Mul(πF, πTemp007, πTemp009); πE != nil {
			continue
		}
		if πTemp001, πE = πg.Div(πF, πTemp004, πTemp006); πE != nil {
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
	Code = πg.NewCode("<module>", "con_monte_go.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
