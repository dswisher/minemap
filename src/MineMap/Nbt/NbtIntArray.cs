
namespace MineMap.Nbt
{
    public class NbtIntArray : NbtTag
    {
        public NbtIntArray(int[] val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.IntArray;

        public int[] Value { get; private set; }
    }
}
