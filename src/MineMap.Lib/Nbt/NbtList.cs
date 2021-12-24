// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Nbt
{
    public class NbtList : NbtTag
    {
        public NbtList(NbtTagType childType, int count)
        {
            ChildType = childType;
            Value = new NbtTag[count];
        }


        public override NbtTagType TagType => NbtTagType.List;


        public NbtTagType ChildType { get; private set; }
        public NbtTag[] Value { get; private set; }


        public override NbtList AsList()
        {
            return this;
        }
    }
}
