// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.IO;
using System.Text;

namespace MineMap.Helpers
{
    public class StreamWrapper : IDisposable
    {
        public StreamWrapper(Stream stream)
        {
            Stream = stream;
        }


        public Stream Stream { get; private set; }


        public void Dispose()
        {
            if (Stream != null)
            {
                Stream.Close();
                Stream = null;
            }
        }


        public long Seek(long offset, SeekOrigin origin)
        {
            return Stream.Seek(offset, origin);
        }


        public byte[] ReadBytes(int len)
        {
            var bytes = new byte[len];

            Stream.Read(bytes, 0, len);

            return bytes;
        }


        public byte ReadByte()
        {
            var val = Stream.ReadByte();

            if (val == -1)
            {
                // TODO - better exception
                throw new Exception("EOF");
            }

            return (byte)val;
        }


        public short ReadShort()
        {
            return ReadSimpleType<short>(2, x => BitConverter.ToInt16(x, 0));
        }


        public ushort ReadUnsignedShort()
        {
            return ReadSimpleType<ushort>(2, x => BitConverter.ToUInt16(x, 0));
        }


        public int ReadInt()
        {
            return ReadSimpleType<int>(4, x => BitConverter.ToInt32(x, 0));
        }


        public long ReadLong()
        {
            return ReadSimpleType<long>(8, x => BitConverter.ToInt64(x, 0));
        }


        public float ReadFloat()
        {
            return ReadSimpleType<float>(4, x => BitConverter.ToSingle(x, 0));
        }


        public double ReadDouble()
        {
            return ReadSimpleType<double>(8, x => BitConverter.ToDouble(x, 0));
        }


        public string ReadString()
        {
            var nameLen = ReadUnsignedShort();

            return ReadString(nameLen);
        }


        public string ReadString(int len)
        {
            // Special case - if the len is zero, return empty string
            if (len == 0)
            {
                return string.Empty;
            }

            // Read the string
            var strBytes = new byte[len];

            Stream.Read(strBytes, 0, len);

            return Encoding.UTF8.GetString(strBytes);
        }


        private T ReadSimpleType<T>(int len, Func<byte[], T> converter)
        {
            var bytes = new byte[len];

            Stream.Read(bytes, 0, len);

            if (BitConverter.IsLittleEndian)
            {
                Array.Reverse(bytes);
            }

            return converter(bytes);
        }
    }
}
