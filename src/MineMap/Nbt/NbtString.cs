
namespace MineMap.Nbt
{
    public class NbtString : NbtTag
    {
        public NbtString(string val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.String;

        public string Value { get; private set; }
    }
}
