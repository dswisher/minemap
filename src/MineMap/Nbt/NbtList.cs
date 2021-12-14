
namespace MineMap.Nbt
{
    public class NbtList : NbtTag
    {
        public NbtList(NbtTagType childType, int count)
        {
            ChildType = childType;
            Value = new NbtTag[count];
        }


        public override NbtTagType TagType => NbtTagType.List;


        public NbtTagType ChildType { get; private set; }
        public NbtTag[] Value { get; private set; }


        public override NbtList AsList()
        {
            return this;
        }
    }
}
