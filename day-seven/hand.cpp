#include "hand.h"
#include <sstream>
#include <iostream>
#include <set>

Hand::Hand(const std::string &str)
{
    std::array<std::string, 2> handBidVec = splitInputStr(str);
    for (size_t i = 0; i < handBidVec[0].length(); i++)
    {
        hand[i].InitValue(handBidVec[0][i]);
    }
    std::string handStr = handBidVec[0];
    bid = stoi(handBidVec[1]);
    type = Type::HIGH_CARD;
    std::map<char, int> distinctCards;
    for (char c : handStr)
    {
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

    for (size_t i = 0; i < hand.size(); i++)
    {
        values[i] = hand[i].GetValue();
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