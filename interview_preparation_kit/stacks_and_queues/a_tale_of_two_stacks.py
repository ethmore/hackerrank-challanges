queue = []
t = int(input())
for line in range(t):
    values = map(int, input().split())
    values = list(values)

    if values[0] == 1:
        queue.append(values[1])        
    elif values[0] == 2:
        queue.pop(0) 
    else:
        print(queue[0])