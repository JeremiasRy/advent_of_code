namespace DayFive;
public class AlmanacMapDecorator : AlmanacMap 
{
    readonly AlmanacMap _almanacMap;
    public override double GetDestination(double source) 
    {
        source = _almanacMap.GetDestination(source);
        foreach (var map in Maps) 
        {
            if (map.FitsToMap(source, out double destination)) 
            {
                return destination;
            }
        }
        return source;
    }

    public AlmanacMapDecorator(AlmanacMap almanacMap)
    {
        _almanacMap = almanacMap;
    }
}