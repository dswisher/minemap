// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using MineMap.Lib.Files;

namespace MineMap.Lib.Nbt
{
    public class NbtReader : IDisposable
    {
        private StreamWrapper wrapper;


        public NbtReader(StreamWrapper wrapper)
        {
            this.wrapper = wrapper;
        }


        public NbtTag ReadTag()
        {
            var tagType = ReadTagType();

            // TagEnd is a special case. It is the only tag that lacks a name and date.
            if (tagType == NbtTagType.End)
            {
                return new NbtEnd();
            }

            // Read the name
            var name = wrapper.ReadString();

            // Based on the type, create the tag and finish parsing it
            return ParseValue(tagType);
        }


        public void Dispose()
        {
            if (wrapper != null)
            {
                wrapper.Dispose();
                wrapper = null;
            }
        }


        private NbtTag ParseValue(NbtTagType tagType)
        {
            switch (tagType)
            {
                case NbtTagType.Byte:
                    return ParseByteValue();

                case NbtTagType.Short:
                    return ParseShortValue();

                case NbtTagType.Int:
                    return ParseIntValue();

                case NbtTagType.Long:
                    return ParseLongValue();

                case NbtTagType.Float:
                    return ParseFloatValue();

                case NbtTagType.Double:
                    return ParseDoubleValue();

                case NbtTagType.String:
                    return ParseStringValue();

                case NbtTagType.List:
                    return ParseListValue();

                case NbtTagType.Compound:
                    return ParseCompoundValue();

                case NbtTagType.ByteArray:
                    return ParseByteArrayValue();

                case NbtTagType.IntArray:
                    return ParseIntArrayValue();

                case NbtTagType.LongArray:
                    return ParseLongArrayValue();

                default:
                    throw new NotImplementedException($"Tag type {tagType} is not yet implemented.");
            }
        }


        private NbtByte ParseByteValue()
        {
            return new NbtByte(wrapper.ReadByte());
        }


        private NbtShort ParseShortValue()
        {
            return new NbtShort(wrapper.ReadShort());
        }


        private NbtInt ParseIntValue()
        {
            return new NbtInt(wrapper.ReadInt());
        }


        private NbtLong ParseLongValue()
        {
            return new NbtLong(wrapper.ReadLong());
        }


        private NbtFloat ParseFloatValue()
        {
            return new NbtFloat(wrapper.ReadFloat());
        }


        private NbtDouble ParseDoubleValue()
        {
            return new NbtDouble(wrapper.ReadDouble());
        }


        private NbtString ParseStringValue()
        {
            return new NbtString(wrapper.ReadString());
        }


        private NbtCompound ParseCompoundValue()
        {
            var parent = new NbtCompound();

            while (true)
            {
                var childType = ReadTagType();

                if (childType == NbtTagType.End)
                {
                    return parent;
                }

                var name = wrapper.ReadString();

                var val = ParseValue(childType);

                parent.Add(name, val);
            }
        }


        private NbtList ParseListValue()
        {
            var childType = ReadTagType();
            var count = wrapper.ReadInt();

            var parent = new NbtList(childType, count);

            for (var i = 0; i < count; i++)
            {
                var child = ParseValue(childType);

                parent.Value[i] = child;
            }

            return parent;
        }


        private NbtByteArray ParseByteArrayValue()
        {
            var count = wrapper.ReadInt();
            var array = new byte[count];

            for (var i = 0; i < count; i++)
            {
                array[i] = wrapper.ReadByte();
            }

            return new NbtByteArray(array);
        }


        private NbtIntArray ParseIntArrayValue()
        {
            var count = wrapper.ReadInt();
            var array = new int[count];

            for (var i = 0; i < count; i++)
            {
                array[i] = wrapper.ReadInt();
            }

            return new NbtIntArray(array);
        }


        private NbtLongArray ParseLongArrayValue()
        {
            var count = wrapper.ReadInt();
            var array = new long[count];

            for (var i = 0; i < count; i++)
            {
                array[i] = wrapper.ReadLong();
            }

            return new NbtLongArray(array);
        }


        private NbtTagType ReadTagType()
        {
            return (NbtTagType)wrapper.ReadByte();
        }
    }
}
