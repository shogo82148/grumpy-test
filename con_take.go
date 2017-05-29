package main
import (
	π___python__Γsys "__python__/sys"
	π___python__Γthreading "__python__/threading"
	πg "grumpy"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ßMyThread := πg.InternStr("MyThread")
	ßThread := πg.InternStr("Thread")
	ß__init__ := πg.InternStr("__init__")
	ß__main__ := πg.InternStr("__main__")
	ß__metaclass__ := πg.InternStr("__metaclass__")
	ß__module__ := πg.InternStr("__module__")
	ß__name__ := πg.InternStr("__name__")
	ßa := πg.InternStr("a")
	ßappend := πg.InternStr("append")
	ßargv := πg.InternStr("argv")
	ßb := πg.InternStr("b")
	ßc := πg.InternStr("c")
	ßint := πg.InternStr("int")
	ßjoin := πg.InternStr("join")
	ßmain := πg.InternStr("main")
	ßresult := πg.InternStr("result")
	ßrun := πg.InternStr("run")
	ßstart := πg.InternStr("start")
	ßsuper := πg.InternStr("super")
	ßsys := πg.InternStr("sys")
	ßtak := πg.InternStr("tak")
	ßthreading := πg.InternStr("threading")
	ßxrange := πg.InternStr("xrange")
	var πTemp001 *πg.Object
	_ = πTemp001
	var πTemp002 []*πg.Object
	_ = πTemp002
	var πTemp003 []πg.Param
	_ = πTemp003
	var πTemp004 *πg.Dict
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
		// line 3: import sys
		πF.SetLineno(3)
		if πTemp002, πE = πg.ImportModule(πF, "sys", []*πg.Code{π___python__Γsys.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 4: import threading
		πF.SetLineno(4)
		if πTemp002, πE = πg.ImportModule(πF, "threading", []*πg.Code{π___python__Γthreading.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 6: def tak(x, y, z):
		πF.SetLineno(6)
		πTemp003 = make([]πg.Param, 3)
		πTemp003[0] = πg.Param{Name: "x", Def: nil}
		πTemp003[1] = πg.Param{Name: "y", Def: nil}
		πTemp003[2] = πg.Param{Name: "z", Def: nil}
		πTemp001 = πg.NewFunction(πg.NewCode("tak", "con_take.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				// line 7: if x <= y:
				πF.SetLineno(7)
			Label1:
				// line 8: return y
				πF.SetLineno(8)
				if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
					continue
				}
				return µy, nil
				goto Label3
			Label2:
				// line 10: return tak(tak((x-1), y , z), tak((y-1), z , x), tak((z-1) , x, y))
				πF.SetLineno(10)
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
		// line 12: class MyThread(threading.Thread):
		πF.SetLineno(12)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp007, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
			continue
		}
		if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßThread, nil); πE != nil {
			continue
		}
		πTemp002[0] = πTemp008
		πTemp004 = πg.NewDict()
		if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
			continue
		}
		_, πE = πg.NewCode("MyThread", "con_take.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp004
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
				// line 13: def __init__(self, a, b, c):
				πF.SetLineno(13)
				πTemp002 = make([]πg.Param, 4)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp002[1] = πg.Param{Name: "a", Def: nil}
				πTemp002[2] = πg.Param{Name: "b", Def: nil}
				πTemp002[3] = πg.Param{Name: "c", Def: nil}
				πTemp001 = πg.NewFunction(πg.NewCode("__init__", "con_take.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µa *πg.Object = πArgs[1]; _ = µa
					var µb *πg.Object = πArgs[2]; _ = µb
					var µc *πg.Object = πArgs[3]; _ = µc
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
						// line 14: super(MyThread, self).__init__()
						πF.SetLineno(14)
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
						// line 15: self.a = a
						πF.SetLineno(15)
						if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µa); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßa, πTemp002); πE != nil {
							continue
						}
						// line 16: self.b = b
						πF.SetLineno(16)
						if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µb); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßb, πTemp002); πE != nil {
							continue
						}
						// line 17: self.c = c
						πF.SetLineno(17)
						if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µc); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßc, πTemp002); πE != nil {
							continue
						}
						// line 18: self.result = 0
						πF.SetLineno(18)
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßresult, πTemp002); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 20: def run(self):
				πF.SetLineno(20)
				πTemp002 = make([]πg.Param, 1)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp003 = πg.NewFunction(πg.NewCode("run", "con_take.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						// line 21: self.result = tak(self.a, self.b, self.c)
						πF.SetLineno(21)
						πTemp001 = πF.MakeArgs(3)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßc, nil); πE != nil {
							continue
						}
						πTemp001[2] = πTemp002
						if πTemp002, πE = πg.ResolveGlobal(πF, ßtak); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßresult, πTemp002); πE != nil {
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
		if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp006 == nil {
			πTemp006 = πg.TypeType.ToObject()
		}
		if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("MyThread").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßMyThread.ToObject(), πTemp007); πE != nil {
			continue
		}
		// line 23: def main():
		πF.SetLineno(23)
		πTemp003 = make([]πg.Param, 0)
		πTemp005 = πg.NewFunction(πg.NewCode("main", "con_take.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µa *πg.Object = πg.UnboundLocal; _ = µa
			var µb *πg.Object = πg.UnboundLocal; _ = µb
			var µc *πg.Object = πg.UnboundLocal; _ = µc
			var µpara *πg.Object = πg.UnboundLocal; _ = µpara
			var µthreads *πg.Object = πg.UnboundLocal; _ = µthreads
			var µi *πg.Object = πg.UnboundLocal; _ = µi
			var µt *πg.Object = πg.UnboundLocal; _ = µt
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 24: a = int(sys.argv[1])
				πF.SetLineno(24)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πg.NewInt(1).ToObject()
				if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßargv, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
					continue
				}
				πTemp001[0] = πTemp003
				if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				µa = πTemp003
				// line 25: b = int(sys.argv[2])
				πF.SetLineno(25)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πg.NewInt(2).ToObject()
				if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßargv, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
					continue
				}
				πTemp001[0] = πTemp003
				if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				µb = πTemp003
				// line 26: c = int(sys.argv[3])
				πF.SetLineno(26)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πg.NewInt(3).ToObject()
				if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßargv, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
					continue
				}
				πTemp001[0] = πTemp003
				if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				µc = πTemp003
				// line 27: para = int(sys.argv[4])
				πF.SetLineno(27)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πg.NewInt(4).ToObject()
				if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßargv, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
					continue
				}
				πTemp001[0] = πTemp003
				if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				µpara = πTemp003
				// line 29: threads = []
				πF.SetLineno(29)
				πTemp001 = make([]*πg.Object, 0)
				πTemp002 = πg.NewList(πTemp001...).ToObject()
				µthreads = πTemp002
				// line 30: for i in xrange(para):
				πF.SetLineno(30)
				πTemp001 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
					continue
				}
				πTemp001[0] = µpara
				if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
					continue
				}
			Label1:
				if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
				µi = πTemp004
				// line 31: t = MyThread(a, b, c)
				πF.SetLineno(31)
				πTemp001 = πF.MakeArgs(3)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp001[0] = µa
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp001[1] = µb
				if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
					continue
				}
				πTemp001[2] = µc
				if πTemp005, πE = πg.ResolveGlobal(πF, ßMyThread); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				µt = πTemp006
				// line 32: t.start()
				πF.SetLineno(32)
				if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
					continue
				}
				// line 33: threads.append(t)
				πF.SetLineno(33)
				πTemp001 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
					continue
				}
				πTemp001[0] = µt
				if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, µthreads, ßappend, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				goto Label1
				goto Label2
			Label2:
				// line 35: for t in threads:
				πF.SetLineno(35)
				if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
					continue
				}
				if πTemp002, πE = πg.Iter(πF, µthreads); πE != nil {
					continue
				}
			Label3:
				if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
				µt = πTemp003
				// line 36: t.join()
				πF.SetLineno(36)
				if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				// line 37: print t.result
				πF.SetLineno(37)
				πTemp001 = make([]*πg.Object, 1)
				if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, µt, ßresult, nil); πE != nil {
					continue
				}
				πTemp001[0] = πTemp004
				if πE = πg.Print(πF, πTemp001, true); πE != nil {
					continue
				}
				goto Label3
				goto Label4
			Label4:
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßmain.ToObject(), πTemp005); πE != nil {
			continue
		}
		if πTemp007, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp006, πE = πg.Eq(πF, πTemp007, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
			continue
		}
		if πTemp009 {
			goto Label1
		}
		goto Label2
		// line 39: if __name__=="__main__":
		πF.SetLineno(39)
	Label1:
		// line 40: main()
		πF.SetLineno(40)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßmain); πE != nil {
			continue
		}
		if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
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
	Code = πg.NewCode("<module>", "con_take.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
