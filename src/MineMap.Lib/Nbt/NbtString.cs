// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Nbt
{
    public class NbtString : NbtTag
    {
        public NbtString(string val)
        {
            Value = val;
        }


        public override NbtTagType TagType => NbtTagType.String;

        public string Value { get; private set; }


        public override NbtString AsString()
        {
            return this;
        }
    }
}
