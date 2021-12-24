// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Nbt
{
    public class NbtShort : NbtTag
    {
        public NbtShort(short val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Short;

        public short Value { get; private set; }
    }
}
