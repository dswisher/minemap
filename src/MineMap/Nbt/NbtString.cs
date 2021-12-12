
namespace MineMap.Nbt
{
    public class NbtString : NbtTag
    {
        public NbtString(string val)
        {
            Value = val;
        }


        public string Value { get; private set; }
    }
}
