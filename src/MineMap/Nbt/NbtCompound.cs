// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System.Collections.Generic;

namespace MineMap.Nbt
{
    public class NbtCompound : NbtTag, IEnumerable<KeyValuePair<string, NbtTag>>
    {
        private readonly Dictionary<string, NbtTag> tags = new Dictionary<string, NbtTag>();

        public NbtCompound()
        {
        }


        public override NbtTagType TagType => NbtTagType.Compound;


        public NbtTag this[string key]
        {
            get
            {
                return tags[key];
            }
        }


        public override NbtCompound AsCompound()
        {
            return this;
        }


        public bool ContainsKey(string key)
        {
            return tags.ContainsKey(key);
        }


        public void Add(string key, NbtTag tag)
        {
            tags.Add(key, tag);
        }


        public IEnumerator<KeyValuePair<string, NbtTag>> GetEnumerator()
        {
            return tags.GetEnumerator();
        }


        System.Collections.IEnumerator System.Collections.IEnumerable.GetEnumerator()
        {
            return tags.GetEnumerator();
        }
    }
}
