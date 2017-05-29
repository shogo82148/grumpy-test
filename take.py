#coding:utf-8
# 竹内関数 Pure Python版
import sys

def tak(x, y, z):
    if x <= y:
        return y
    else:
        return tak(tak((x-1), y , z), tak((y-1), z , x), tak((z-1) , x, y))

if __name__=="__main__":
    a = int(sys.argv[1])
    b = int(sys.argv[2])
    c = int(sys.argv[3])
    print tak(a, b, c)
