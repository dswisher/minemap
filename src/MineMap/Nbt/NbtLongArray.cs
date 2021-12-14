// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Nbt
{
    public class NbtLongArray : NbtTag
    {
        public NbtLongArray(long[] val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.LongArray;

        public long[] Value { get; private set; }
    }
}
