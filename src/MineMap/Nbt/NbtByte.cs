// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Nbt
{
    public class NbtByte : NbtTag
    {
        public NbtByte(byte val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Byte;

        public byte Value { get; private set; }


        public override NbtByte AsByte()
        {
            return this;
        }
    }
}
