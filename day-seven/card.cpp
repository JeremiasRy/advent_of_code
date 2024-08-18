#include "card.h"
#include <iostream>
#include <map>
#include <string>

int Card::GetValue() const
{
    return Card::labelToValueMap.at(label);
}

void Card::InitValue(char c)
{
    label = c;
}

const std::map<char, int> Card::labelToValueMap = {
    {'2', 0},
    {'3', 1},
    {'4', 2},
    {'5', 3},
    {'6', 4},
    {'7', 5},
    {'8', 6},
    {'9', 7},
    {'T', 8},
    {'J', 9},
    {'Q', 10},
    {'K', 11},
    {'A', 12}};
