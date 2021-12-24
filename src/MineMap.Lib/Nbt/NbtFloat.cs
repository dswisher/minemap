// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Nbt
{
    public class NbtFloat : NbtTag
    {
        public NbtFloat(float val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.Float;


        public float Value { get; private set; }
    }
}
