#coding:utf-8
# モンテカルロ法 Python+Go 並列版
import threading
import random
import sys
from __go__.time import Now
from __go__.math.rand import New, NewSource


class MyThread(threading.Thread):
    def __init__(self):
        super(MyThread, self).__init__()
        self.c = 0

    def run(self):
        r = New(NewSource(Now().UnixNano()))
        c = 0
        for _ in xrange(num):
            x = r.Float64()
            y = r.Float64()

            if x * x + y * y <= 1.0:
                c += 1
        self.c = c


if __name__ == "__main__":
    num = int(sys.argv[1])
    para = int(sys.argv[2])

    threads = []
    for i in xrange(para):
        t = MyThread()
        t.start()
        threads.append(t)

    c = 0
    for t in threads:
        t.join()
        c += t.c

    print 4.0*c/(num*para)
