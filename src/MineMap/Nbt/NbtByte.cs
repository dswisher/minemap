
namespace MineMap.Nbt
{
    public class NbtByte : NbtTag
    {
        public NbtByte(byte val)
        {
            Value = val;
        }


        public byte Value { get; private set; }
    }
}
