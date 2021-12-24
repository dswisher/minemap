// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Nbt
{
    public class NbtDouble : NbtTag
    {
        public NbtDouble(double val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Double;

        public double Value { get; private set; }
    }
}
