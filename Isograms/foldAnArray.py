# https://www.codewars.com/kata/57ea70aa5500adfe8a000110/python

def help(array):
    ans = []

    if len(array)%2==0:
        mid = (len(array) / 2) - 1
    else:
        mid = int(len(array) / 2)
    
    for i, v in enumerate(array):
        if i == mid and len(array)%2!=0:
            ans.append(v)
            break
        elif i == mid:
            ans.append(array[i] + array[-(i+1)])
            break
        ans.append(array[i] + array[-(i+1)])
    return ans


def fold_array(array, runs):
    for _ in range(runs):
        array = help(array)
    return array