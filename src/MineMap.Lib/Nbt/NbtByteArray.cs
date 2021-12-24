// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Nbt
{
    public class NbtByteArray : NbtTag
    {
        public NbtByteArray(byte[] val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.ByteArray;

        public byte[] Value { get; private set; }
    }
}
