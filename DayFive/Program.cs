using System.Security.Cryptography;
using System.Text.RegularExpressions;
using DayFive;

var commandLineArgs = Environment.GetCommandLineArgs();
int[] seeds = Array.Empty<int>();
Almanac currentlyParsing = Almanac.None;
Dictionary<Almanac, AlmanacMap> almanacs = new();
if (commandLineArgs.Length != 2) 
{
    Console.WriteLine("Usage: dotnet run input.txt");
    return;
}
try 
{
    using (var sr = new StreamReader(commandLineArgs[1])) 
    {
        bool firstRow = true;
        while (!sr.EndOfStream) 
        {
            var line = sr.ReadLine();
            if (line != null) 
            {
                if (line.Length < 1) 
                {
                    continue;
                }
                if (firstRow) 
                {
                    
                    var seedsStr = line.Split(":")[1];
                    seeds = DigitRegex().Matches(seedsStr).Select(match => int.Parse(match.Value)).ToArray();
                    firstRow = false;
                    continue;
                }
                if (line.Contains("map:")) 
                {
                    currentlyParsing = ParseLineToAlmanac(line);
                    almanacs.Add(currentlyParsing, new AlmanacMap());
                    continue;
                }
                if (currentlyParsing == Almanac.None) 
                {
                    Console.WriteLine("You shouldn't reach this part before setting a Almanac. Check input.");
                    return;
                }
                int[] mapValues = DigitRegex().Matches(line).Select(match => int.Parse(match.Value)).ToArray();
                almanacs[currentlyParsing].EnterMapEntries(mapValues[0], mapValues[1], mapValues[2]);
            } 
        }
    };
    foreach (int seed in seeds) {
        var destination = almanacs[Almanac.SeedToSoil].GetDestination(seed);
        Console.WriteLine("seed: {0}, soil: {1}", seed, destination);
    }
} catch 
{
    Console.WriteLine("can't open {0}", commandLineArgs[1]);
}

partial class Program
{
    [GeneratedRegex("\\d+")]
    private static partial Regex DigitRegex();

    enum Almanac
    {
        SeedToSoil,
        SoilToFertilizer,
        FertilizerToWater,
        WaterToLight,
        LightToTemperature,
        TemperatureToHumidity,
        HumidityToLocation,
        None
    }

    static Almanac ParseLineToAlmanac(string line) 
    {
        return line.Split(" map:")[0] switch
        {
            "seed-to-soil" => Almanac.SeedToSoil,
            "soil-to-fertilizer" => Almanac.SoilToFertilizer,
            "fertilizer-to-water" => Almanac.FertilizerToWater,
            "water-to-light" => Almanac.WaterToLight,
            "light-to-temperature" => Almanac.LightToTemperature,
            "temperature-to-humidity" => Almanac.TemperatureToHumidity,
            "humidity-to-location" => Almanac.HumidityToLocation,
            _ => throw new ArgumentException("Line: `{0}` didn't contain almanac expression"),
        };
    }
}