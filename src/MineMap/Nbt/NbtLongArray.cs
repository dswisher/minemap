
namespace MineMap.Nbt
{
    public class NbtLongArray : NbtTag
    {
        public NbtLongArray(long[] val)
        {
            Value = val;
        }


        public long[] Value { get; private set; }
    }
}
