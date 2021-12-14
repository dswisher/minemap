// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Nbt
{
    public class NbtLong : NbtTag
    {
        public NbtLong(long val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Long;

        public long Value { get; private set; }


        public override NbtLong AsLong()
        {
            return this;
        }
    }
}
