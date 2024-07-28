namespace DayFive;
public class Map 
{
    readonly double _destination, _source, _range;

    public bool FitsToMap(double source, out double destination) {
        destination = 0;
        if (source >= _source && source < _source + _range) 
        {
            destination = _destination + source - _source;
            return true;
        }
        return false;
    }
    public Map(double destination, double source, double range)
    {
        _destination = destination;
        _source = source;
        _range = range;
    }
    public override string ToString()
    {
        return $"_destination: {_destination}, _source, {_source}, _range {_range}";
    }
}