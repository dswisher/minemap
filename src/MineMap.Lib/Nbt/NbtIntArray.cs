// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Nbt
{
    public class NbtIntArray : NbtTag
    {
        public NbtIntArray(int[] val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.IntArray;

        public int[] Value { get; private set; }
    }
}
