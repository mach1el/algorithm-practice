from collections import Counter
input()
socks, pairs = Counter(map(int,input().strip().split())), 0
for s in socks: pairs += socks[s]//2
print(pairs)