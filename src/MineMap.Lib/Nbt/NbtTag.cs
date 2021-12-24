// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

namespace MineMap.Lib.Nbt
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

        public virtual NbtString AsString()
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
