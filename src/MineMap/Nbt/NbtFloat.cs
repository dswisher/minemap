
namespace MineMap.Nbt
{
    public class NbtFloat : NbtTag
    {
        public NbtFloat(float val)
        {
            Value = val;
        }


        public float Value { get; private set; }
    }
}
