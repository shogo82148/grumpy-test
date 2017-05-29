#coding:utf-8
# 竹内関数 Pure Python 並列版
import sys
import threading

def tak(x, y, z):
    if x <= y:
        return y
    else:
        return tak(tak((x-1), y , z), tak((y-1), z , x), tak((z-1) , x, y))

class MyThread(threading.Thread):
    def __init__(self, a, b, c):
        super(MyThread, self).__init__()
        self.a = a
        self.b = b
        self.c = c
        self.result = 0

    def run(self):
        self.result = tak(self.a, self.b, self.c)

def main():
    a = int(sys.argv[1])
    b = int(sys.argv[2])
    c = int(sys.argv[3])
    para = int(sys.argv[4])

    threads = []
    for i in xrange(para):
        t = MyThread(a, b, c)
        t.start()
        threads.append(t)

    for t in threads:
        t.join()
        print t.result

if __name__=="__main__":
    main()
