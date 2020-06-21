# Problem statement

"""
The new "Avengers" movie has just been released! There are a lot of people at the cinema box office standing in a huge line. Each of them has a single 100, 50 or 25 dollar bill. An "Avengers" ticket costs 25 dollars.

Vasya is currently working as a clerk. He wants to sell a ticket to every single person in this line.

Can Vasya sell a ticket to every person and give change if he initially has no money and sells the tickets strictly in the order people queue?

Return YES, if Vasya can sell a ticket to every person and give change with the bills he has at hand at that moment. Otherwise return NO.

Examples:
tickets([25, 25, 50]) # => YES 
tickets([25, 100]) # => NO. Vasya will not have enough money to give change to 100 dollars
tickets([25, 25, 50, 50, 100]) # => NO. Vasya will not have the right bills to give 75 dollars of change (you can't make two bills of 25 from one of 50)
"""

def tickets(people):
    if people[0] != 25:
        return "NO"

    the_map = {
        25 : 25,
        50 : -25,
        100 : -75
    }
    bal = 0
    for i in people:
        bal = bal + the_map[i]
        if bal < 0:
            return "NO"
    return "YES"

if __name__ == "__main__":
    print(tickets([25, 25, 50]))