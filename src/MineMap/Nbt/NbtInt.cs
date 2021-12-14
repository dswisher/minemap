
namespace MineMap.Nbt
{
    public class NbtInt : NbtTag
    {
        public NbtInt(int val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Int;


        public override NbtInt AsInt()
        {
            return this;
        }


        public int Value { get; private set; }
    }
}
