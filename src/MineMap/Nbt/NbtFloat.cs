
namespace MineMap.Nbt
{
    public class NbtFloat : NbtTag
    {
        public NbtFloat(float val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Float;


        public float Value { get; private set; }
    }
}
