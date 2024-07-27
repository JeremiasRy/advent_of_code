using System.Text.Json.Serialization;

namespace DayFive;

public class AlmanacMap 
{
    public Stack<Map> Maps = new();
    public virtual double GetDestination(double source) 
    {
        foreach (var map in Maps) 
        {
            if (map.FitsToMap(source, out double destination)) 
            {
                return destination;
            }

        }
        return source;
    }
    public void AddToMap(double destination, double source, double range) 
    {
        Maps.Push(new Map(destination, source, range));
    }
}




