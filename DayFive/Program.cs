using System.Text.RegularExpressions;
using DayFive;

var commandLineArgs = Environment.GetCommandLineArgs();
double[][] seeds = Array.Empty<double[]>();
AlmanacMap map = new();
object lockObj = new();
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
                    var seedsDigits = DigitRegex().Matches(seedsStr).Select(match => double.Parse(match.Value)).ToArray();
                    seeds = new double[seedsDigits.Length / 2][];
                    var outerIndex = 0;
                    for (int i = 0; i < seedsDigits.Length; i += 2) 
                    {
                        seeds[outerIndex] = new double[2];
                        Array.Copy(seedsDigits[i..], seeds[outerIndex], 2);
                        outerIndex++;
                    }
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
    int threadCount = 1;

    Parallel.ForEach(seeds, seedRange  => {
        string loopId = "";
        lock (lockObj) 
        {
            Console.WriteLine($"Loop {threadCount}. \nStarting parsing...");
            loopId = $"loop {threadCount}";
            threadCount++;
        }   
        
        for (double seed = seedRange[0]; seed < seedRange[0] + seedRange[1]; seed++) 
        {
            lock (lockObj) 
            {
                lowest = Math.Min(lowest, map.GetDestination(seed));
            }
        }
        Console.WriteLine("Finished {0}.", loopId);
        Console.WriteLine();
    }); 

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