namespace DayFive;
public class AlmanacMap 
{
    private readonly Dictionary<double, double> _map =  new();
    private readonly AlmanacMap _next;
    public double GetDestination(double source) 
    {
        if (_map.TryGetValue(source, out double destination)) 
        {
            return destination;
        }
        return source;
    } 

    public void EnterMapEntries(double destination, double source, double range)
    {
        Console.WriteLine("Mapping entries, range: {0}", range);
        for (double i = 0; i <= range; i++) 
        {
            double key = source + i;
            double value = destination + i;

            _map.Remove(key);
            _map.Add(key, value);
        }
        Console.WriteLine("Mapped entries");
    }
}