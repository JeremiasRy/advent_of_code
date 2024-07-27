using System.Text.RegularExpressions;
using DayFive;

var commandLineArgs = Environment.GetCommandLineArgs();
double[] seeds = Array.Empty<double>();
AlmanacMap map = new();
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
        bool initialMap = true;
        while (!sr.EndOfStream) 
        {
            var line = sr.ReadLine();
            if (line != null) 
            {
                if (firstRow) 
                {
                    var seedsStr = line.Split(":")[1];
                    seeds = DigitRegex().Matches(seedsStr).Select(match => double.Parse(match.Value)).ToArray();
                    firstRow = false;
                    continue;
                }
                if (line.Length < 1) 
                {
                    continue;
                }
                if (line.Contains("map:")) 
                {
                    if (initialMap) 
                    {
                        initialMap = false;
                        continue;
                    }
                    map = new AlmanacMapDecorator(map);
                    continue;
                }
                var mapValues = DigitRegex().Matches(line).Select(match => double.Parse(match.Value)).ToArray();
                map.AddToMap(mapValues[0], mapValues[1], mapValues[2]);
            } 
        }
    };
    double lowest = double.MaxValue;
    foreach (double seed in seeds) 
    {
        var destination = map.GetDestination(seed);
        if (destination < lowest) 
        {
            lowest = destination;
        }
    }
    Console.WriteLine("Lowest: {0}", lowest);
} catch (Exception e)
{
    Console.WriteLine(e);
    Console.WriteLine("Can't open file {0}", commandLineArgs[1]);
}

partial class Program
{
    [GeneratedRegex("\\d+")]
    private static partial Regex DigitRegex();
}