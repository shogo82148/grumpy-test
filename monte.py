#coding:utf-8
# モンテカルロ法 Pure Python版
import random
import sys

if __name__=="__main__":
    num = int(sys.argv[1])
    c = 0

    for i in xrange(num):
        x = random.random()
        y = random.random()

        if x * x + y * y <= 1.0:
            c += 1

    print 4.0*c/num
