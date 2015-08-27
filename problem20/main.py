def qsum(n):
    ret = 0
    while n > 0:
        ret += n % 10
        n = n // 10
    return ret

def factorial(n):
    ret = 1
    while n > 0:
        ret *= n
        n -= 1
    return ret

print(qsum(123))
print(factorial(10))
print(qsum(factorial(100)))
