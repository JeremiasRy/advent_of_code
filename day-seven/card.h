#ifndef CARD_H
#define CARD_H
#include <map>

class Card
{
public:
    char label;
    static const std::map<char, int> labelToValueMap;
    int GetValue() const;
    void InitValue(char c);
};

#endif