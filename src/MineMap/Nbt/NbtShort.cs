
namespace MineMap.Nbt
{
    public class NbtShort : NbtTag
    {
        public NbtShort(short val)
        {
            Value = val;
        }


        public short Value { get; private set; }
    }
}
