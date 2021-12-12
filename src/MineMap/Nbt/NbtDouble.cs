
namespace MineMap.Nbt
{
    public class NbtDouble : NbtTag
    {
        public NbtDouble(double val)
        {
            Value = val;
        }


        public double Value { get; private set; }
    }
}
