def qsum(n):
        ret = 0
        while n > 0:
            ret += n % 10
            n = n // 10
        return ret

print(qsum(2**1000))
