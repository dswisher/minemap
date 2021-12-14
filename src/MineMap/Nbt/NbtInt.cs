// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Nbt
{
    public class NbtInt : NbtTag
    {
        public NbtInt(int val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Int;

        public int Value { get; private set; }


        public override NbtInt AsInt()
        {
            return this;
        }
    }
}
