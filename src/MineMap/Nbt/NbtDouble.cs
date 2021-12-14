
namespace MineMap.Nbt
{
    public class NbtDouble : NbtTag
    {
        public NbtDouble(double val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Double;

        public double Value { get; private set; }
    }
}
