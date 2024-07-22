namespace DayFive;
public class AlmanacMap 
{
    private readonly Dictionary<int, int> _map =  new();
    private readonly AlmanacMap _next;
    public int GetDestination(int source) 
    {
        if (_map.TryGetValue(source, out int destination)) 
        {
            return destination;
        }
        return source;
    } 

    public void EnterMapEntries(int destination, int source, int range)
    {
        for (int i = 0; i <= range; i++) 
        {
            int key = source + i;
            int value = destination + i;

            _map.Remove(key);
            _map.Add(key, value);
        }
    }
}