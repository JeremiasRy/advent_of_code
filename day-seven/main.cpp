#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>
#include "hand.h"

int main(int argc, char *argv[])
{
    std::ifstream file(argv[1]);
    if (!file.is_open())
    {
        std::cout << "Cant open file: " << argv[0];
        return 1;
    }

    std::string line;
    std::vector<Hand> hands;

    while (std::getline(file, line))
    {
        hands.push_back(Hand(line));
    }

    std::sort(hands.begin(), hands.end());
    int rank = 1;
    int result = 0;
    for (auto &&hand : hands)
    {
        std::cout << "Rank: " << rank << " Card: " << hand.hand << " type: " << hand.type << "\n";
        result += hand.GetBid() * rank;
        rank++;
    };
    std::cout << "Result: " << result << "\n";
    return 0;
};