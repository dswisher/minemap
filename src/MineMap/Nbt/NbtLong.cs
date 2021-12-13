
namespace MineMap.Nbt
{
    public class NbtLong : NbtTag
    {
        public NbtLong(long val)
        {
            Value = val;
        }


        public override NbtLong AsLong()
        {
            return this;
        }


        public long Value { get; private set; }
    }
}
