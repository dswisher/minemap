
namespace MineMap.Nbt
{
    public class NbtByteArray : NbtTag
    {
        public NbtByteArray(byte[] val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.ByteArray;

        public byte[] Value { get; private set; }
    }
}
