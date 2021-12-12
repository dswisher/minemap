
using System;
using System.IO;
using System.Text;

namespace MineMap.Nbt
{
    public class NbtReader : IDisposable
    {
        private Stream stream;


        public NbtReader(Stream stream)
        {
            this.stream = stream;
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
            var name = ReadString();

            Console.WriteLine(" -> tag type {0}, name '{1}'", tagType, name);

            // Based on the type, create the tag and finish parsing it
            return ParseValue(tagType);
        }


        public void Dispose()
        {
            if (stream != null)
            {
                stream.Close();
                stream = null;
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

                case NbtTagType.IntArray:
                    return ParseIntArrayValue();

                default:
                    throw new NotImplementedException($"Tag type {tagType} is not yet implemented.");
            }
        }


        private NbtByte ParseByteValue()
        {
            return new NbtByte(ReadByte());
        }


        private NbtShort ParseShortValue()
        {
            return new NbtShort(ReadShort());
        }


        private NbtInt ParseIntValue()
        {
            return new NbtInt(ReadInt());
        }


        private NbtLong ParseLongValue()
        {
            return new NbtLong(ReadLong());
        }


        private NbtFloat ParseFloatValue()
        {
            return new NbtFloat(ReadFloat());
        }


        private NbtDouble ParseDoubleValue()
        {
            return new NbtDouble(ReadDouble());
        }


        private NbtString ParseStringValue()
        {
            return new NbtString(ReadString());
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

                var name = ReadString();

                Console.WriteLine(" -> tag type {0}, name '{1}'", childType, name);

                var val = ParseValue(childType);

                // TODO - add the child onto the parent
            }
        }


        private NbtList ParseListValue()
        {
            var parent = new NbtList();

            var childType = ReadTagType();
            var count = ReadInt();

            for (var i = 0; i < count; i++)
            {
                var child = ParseValue(childType);

                // TODO - add child to the list
            }

            return parent;
        }


        private NbtIntArray ParseIntArrayValue()
        {
            var count = ReadInt();
            var array = new int[count];

            for (var i = 0; i < count; i++)
            {
                array[i] = ReadInt();
            }

            return new NbtIntArray(array);
        }


        private NbtTagType ReadTagType()
        {
            return (NbtTagType)ReadByte();
        }


        private byte ReadByte()
        {
            var val = stream.ReadByte();

            if (val == -1)
            {
                // TODO - better exception
                throw new Exception("EOF");
            }

            return (byte)val;
        }


        private short ReadShort()
        {
            var bytes = new byte[2];

            stream.Read(bytes, 0, 2);

            if (BitConverter.IsLittleEndian)
            {
                Array.Reverse(bytes);
            }

            return BitConverter.ToInt16(bytes, 0);
        }


        private ushort ReadUnsignedShort()
        {
            var bytes = new byte[2];

            stream.Read(bytes, 0, 2);

            if (BitConverter.IsLittleEndian)
            {
                Array.Reverse(bytes);
            }

            return BitConverter.ToUInt16(bytes, 0);
        }


        private int ReadInt()
        {
            var bytes = new byte[4];

            stream.Read(bytes, 0, 4);

            if (BitConverter.IsLittleEndian)
            {
                Array.Reverse(bytes);
            }

            return BitConverter.ToInt32(bytes, 0);
        }


        private long ReadLong()
        {
            var bytes = new byte[8];

            stream.Read(bytes, 0, 8);

            if (BitConverter.IsLittleEndian)
            {
                Array.Reverse(bytes);
            }

            return BitConverter.ToInt64(bytes, 0);
        }


        private float ReadFloat()
        {
            var bytes = new byte[4];

            stream.Read(bytes, 0, 4);

            if (BitConverter.IsLittleEndian)
            {
                Array.Reverse(bytes);
            }

            return BitConverter.ToSingle(bytes, 0);
        }


        private double ReadDouble()
        {
            var bytes = new byte[8];

            stream.Read(bytes, 0, 8);

            if (BitConverter.IsLittleEndian)
            {
                Array.Reverse(bytes);
            }

            return BitConverter.ToDouble(bytes, 0);
        }


        private string ReadString()
        {
            var nameLen = ReadUnsignedShort();

            return ReadString(nameLen);
        }


        private string ReadString(int len)
        {
            // Special case - if the len is zero, return empty string
            if (len == 0)
            {
                return string.Empty;
            }

            // Read the string
            var strBytes = new byte[len];

            stream.Read(strBytes, 0, len);

            return Encoding.UTF8.GetString(strBytes);
        }
    }
}
