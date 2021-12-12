
namespace MineMap.Nbt
{
    public class NbtInt : NbtTag
    {
        public NbtInt(int val)
        {
            Value = val;
        }


        public int Value { get; private set; }
    }
}
