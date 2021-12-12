
namespace MineMap.Nbt
{
    public class NbtIntArray : NbtTag
    {
        public NbtIntArray(int[] val)
        {
            Value = val;
        }


        public int[] Value { get; private set; }
    }
}
