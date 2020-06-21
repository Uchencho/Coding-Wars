def validBraces(string):
    if len(string) % 2 != 0:
        return False
    elif len(string) == 1:
        return False

    c = [line for line in string]
    the_map = {
        "[" : "]",
        "{" : "}",
        "(" : ")",
    }

    for i in string:
        if i in the_map.keys() and the_map[i] not in c:
            return False
        elif i in the_map.keys() and c.count(i) != c.count(the_map[i]):
            return False

    a_key = c[int(len(c) / 2) - 1]
    if a_key not in the_map.keys():
        return False
    if the_map[a_key] != c[int(len(c) / 2)]:
        return False
            
    return True

if __name__ == '__main__':
    print(validBraces("(){}[]")) 
    print(validBraces("(}"))
    print(validBraces("[({})](]"))
    print(validBraces("([{}])"))
    print(validBraces("[(])"))