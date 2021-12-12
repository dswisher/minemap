
namespace MineMap.Nbt
{
    public class NbtLong : NbtTag
    {
        public NbtLong(long val)
        {
            Value = val;
        }


        public long Value { get; private set; }
    }
}
