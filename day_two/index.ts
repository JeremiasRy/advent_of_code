import * as fs from "fs";
import { exit } from "process";

const CUBES_IN_BAG = new Map<string, number>();
CUBES_IN_BAG.set("red", 12);
CUBES_IN_BAG.set("green", 13);
CUBES_IN_BAG.set("blue", 14);
Object.freeze(CUBES_IN_BAG);

const gameIdentifierPrefix = /^(.*?)(\d+).*:/;

fs.readFile("input.txt", "utf-8", (err, data) => {
  const games = data.split(/\r?\n/);
  console.log(
    "Result:",
    games.reduce((acc, game) => acc + solveGame(game), 0)
  );
  exit();
});

function solveGame(game: string): number {
  const gameId = game.match(/^(.*?)(\d+).*:/)?.at(2);
  const smallestNumbers = new Map();
  if (!gameId) {
    return 0;
  }
  game = game.replace(gameIdentifierPrefix, "");

  const pulls = game.split(";");
  for (const pull of pulls) {
    const dices = pull.split(",");

    for (const dice of dices) {
      for (const [color, _] of CUBES_IN_BAG) {
        if (dice.includes(color)) {
          const amount = dice.match(/\d+/);
          const parsed = parseInt(amount?.[0]!);
          if (smallestNumbers.has(color)) {
            if (smallestNumbers.get(color) < parsed) {
              smallestNumbers.set(color, parsed);
            }
          } else {
            smallestNumbers.set(color, parsed);
          }
        }
      }
    }
  }
  return [...smallestNumbers.values()].reduce((acc, val) => acc * val);
}
