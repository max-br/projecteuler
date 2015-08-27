def fib(n):
    f1 = 0
    f2 = 1
    while n > 1:
        tmp = f2
        f2 = f2+f1
        f1 = tmp
        n -= 1
    return f2

def cnt_digits(n):
    cnt = 0
    while n > 0:
        n = n // 10
        cnt += 1
    return cnt

digitcnt = 0
i = 0
while digitcnt < 1000:
    i += 1
    f = fib(i)
    digitcnt = cnt_digits(f)

print(i, fib(i),cnt_digits(fib(i)))
