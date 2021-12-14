
namespace MineMap.Nbt
{
    public class NbtLongArray : NbtTag
    {
        public NbtLongArray(long[] val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.LongArray;

        public long[] Value { get; private set; }
    }
}
