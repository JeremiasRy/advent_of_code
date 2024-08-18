#ifndef HAND_H
#define HAND_H

#include "card.h"
#include <string>

class Hand
{

private:
    std::array<Card, 5> hand;
    int bid;
    std::array<std::string, 2> splitInputStr(const std::string &str);

public:
    enum Type
    {
        HIGH_CARD,
        ONE_PAIR,
        TWO_PAIR,
        THREE_OF_A_KIND,
        FULL_HOUSE,
        FOUR_OF_A_KIND,
        FIVE_OF_A_KIND
    };
    Hand::Type type;
    std::array<int, 5> values;
    Hand(const std::string &str);
    int GetBid();
    friend bool operator<(const Hand &a, const Hand &b);
};

#endif // HAND_H