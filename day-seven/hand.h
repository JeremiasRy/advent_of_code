#ifndef HAND_H
#define HAND_H

#include <string>
#include <map>

class Hand
{

private:
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

    static const std::map<char, int> charToValueMap;

    Hand(const std::string &str);

    std::string hand;
    Hand::Type type;
    std::array<int, 5> values;
    int GetBid();
    friend bool operator<(const Hand &a, const Hand &b);
};

#endif // HAND_H