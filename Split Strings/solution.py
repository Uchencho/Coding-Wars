#https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/python

"""
Complete the solution so that it splits the string into pairs of two characters. 
If the string contains an odd number of characters then it should replace the missing 
second character of the final pair with an underscore ('_').

Examples:

solution('abc') # should return ['ab', 'c_']
solution('abcdef') # should return ['ab', 'cd', 'ef']
"""


def solution(s):
    #Final list of strings to be returned
    ans  = []

    #Loop through each letter and check if it is the second one
    for i, v in enumerate(s):
        if (i + 1) % 2 == 0:
            #If it is, append it with its predecessor
            ans.append(s[(i-1)] + v)

    #After the loop, check if it needs and underscore at the end
    if len(s) > 0 and (len(s) % 2) != 0:
        ans.append(s[-1]+"_")    
    return ans


if __name__ == "__main__":
    print(solution(''))
    print(solution("abc"))
    print(solution('abcdef'))
    print(solution("asdfads"))
    print(solution("x"))