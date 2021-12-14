
using System;

namespace MineMap.Nbt
{
    public abstract class NbtTag
    {
        public abstract NbtTagType TagType { get; }

        public virtual NbtByte AsByte()
        {
            throw new InvalidCastException();
        }

        public virtual NbtInt AsInt()
        {
            throw new InvalidCastException();
        }

        public virtual NbtLong AsLong()
        {
            throw new InvalidCastException();
        }

        public virtual NbtList AsList()
        {
            throw new InvalidCastException();
        }

        public virtual NbtCompound AsCompound()
        {
            throw new InvalidCastException();
        }
    }
}
