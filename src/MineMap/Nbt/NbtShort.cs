
namespace MineMap.Nbt
{
    public class NbtShort : NbtTag
    {
        public NbtShort(short val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Short;

        public short Value { get; private set; }
    }
}
