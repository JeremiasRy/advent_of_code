#include "hand.h"
#include <sstream>
#include <iostream>
#include <set>

const std::map<char, int> Hand::charToValueMap = {
    {'2', 0},
    {'3', 1},
    {'4', 2},
    {'5', 3},
    {'6', 4},
    {'7', 5},
    {'8', 6},
    {'9', 7},
    {'T', 8},
    {'J', -1},
    {'Q', 10},
    {'K', 11},
    {'A', 12}};

Hand::Hand(const std::string &str)
{
    std::array<std::string, 2> handBidVec = splitInputStr(str);
    std::string handStr = handBidVec[0];
    bid = stoi(handBidVec[1]);
    type = Type::HIGH_CARD;
    hand = str;
    std::map<char, int> distinctCards;
    bool hasJoker = false;

    for (char c : handStr)
    {
        if (!hasJoker)
        {
            hasJoker = c == 'J';
        }
        distinctCards[c]++;
    }

    if (distinctCards.size() == 1)
    {
        type = Hand::FIVE_OF_A_KIND;
    }
    else
    {
        for (const auto &pair : distinctCards)
        {
            if (pair.second == 2)
            {
                if (type == Type::THREE_OF_A_KIND)
                {
                    type = Type::FULL_HOUSE;
                    break;
                }
                if (type == Type::ONE_PAIR)
                {
                    type = Type::TWO_PAIR;
                    break;
                }
                type = Type::ONE_PAIR;
                continue;
            }
            if (pair.second == 3)
            {
                if (type == Type::ONE_PAIR)
                {

                    type = Type::FULL_HOUSE;
                    break;
                }
                type = Type::THREE_OF_A_KIND;
                continue;
            }
            if (pair.second == 4)
            {
                type = Type::FOUR_OF_A_KIND;
                continue;
            }
        }
    }
    for (size_t i = 0; i < handStr.size(); i++)
    {
        values[i] = charToValueMap.at(handStr[i]);
    }

    if (hasJoker)
    {
        int jokers = distinctCards.at('J');
        if (jokers == 4)
        {
            type = Type::FIVE_OF_A_KIND;
        }
        else if (jokers == 3)
        {
            switch (type)
            {
            case Type::THREE_OF_A_KIND:
                type = Type::FOUR_OF_A_KIND;
                break;
            case Type::FULL_HOUSE:
                type = Type::FIVE_OF_A_KIND;
            }
        }
        else if (jokers == 2)
        {
            switch (type)
            {
            case Type::HIGH_CARD:
                type = Type::THREE_OF_A_KIND;
                break;
            case Type::ONE_PAIR:
                type = Type::THREE_OF_A_KIND;
                break;
            case Type::TWO_PAIR:
                type = Type::FOUR_OF_A_KIND;
                break;
            case Type::THREE_OF_A_KIND:
            case Type::FULL_HOUSE:
                type = Type::FIVE_OF_A_KIND;
                break;
            }
        }
        else if (jokers == 1)
        {
            switch (type)
            {
            case Type::HIGH_CARD:
                type = Type::ONE_PAIR;
                break;
            case Type::ONE_PAIR:
                type = Type::THREE_OF_A_KIND;
                break;
            case Type::TWO_PAIR:
                type = Type::FULL_HOUSE;
                break;
            case Type::THREE_OF_A_KIND:
                type = Type::FOUR_OF_A_KIND;
                break;
            case Type::FOUR_OF_A_KIND:
                type = Type::FIVE_OF_A_KIND;
                break;
            }
        }
    }
}

int Hand::GetBid()
{
    return bid;
}

bool operator<(const Hand &a, const Hand &b)
{
    if (a.type == b.type)
    {
        int i = 0;
        for (auto &value : a.values)
        {
            if (value == b.values[i])
            {
                i++;
                continue;
            }
            return value < b.values[i];
        }
    }
    return a.type < b.type;
}

std::array<std::string, 2> Hand::splitInputStr(const std::string &str)
{
    std::array<std::string, 2> result;
    std::string token;
    std::stringstream ss(str);
    int i = 0;
    while (std::getline(ss, token, ' ') && i < 2)
    {
        result[i] = token;
        i++;
    }

    return result;
}