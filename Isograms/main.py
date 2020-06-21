"""

An isogram is a word that has no repeating letters, consecutive or non-consecutive. Implement a function that determines whether a string that contains only letters is an isogram. Assume the empty string is an isogram. Ignore letter case.

is_isogram("Dermatoglyphics" ) == true
is_isogram("aba" ) == false
is_isogram("moOse" ) == false # -- ignore letter case

"""




def is_isogram(string):
    for i in string.lower():
        if string.lower().count(i) > 1: return False
    return True

def is_isogram_better(string):
    return len(string) == len(set(string.lower()))


if __name__ == "__main__":
    print(is_isogram("moOse"))