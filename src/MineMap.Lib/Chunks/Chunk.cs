// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.IO;

using MineMap.Lib.Files;
using MineMap.Lib.Nbt;

namespace MineMap.Lib.Chunks
{
    public class Chunk
    {
        private readonly NbtCompound root;

        private Chunk(NbtCompound root)
        {
            this.root = root;

#if false
            X = root["xPos"].AsInt().Value;
            Y = root["yPos"].AsInt().Value;
            Z = root["zPos"].AsInt().Value;

            // TODO! HACK! Remove this debug code
            Console.WriteLine("Chunk: X={0}, Y={1}, Z={2}", X, Y, Z);

            // root.Dump();

#if true
            var sections = root["sections"].AsList();

            foreach (var stag in sections.Value)
            {
                var section = stag.AsCompound();
                var sy = section["Y"].AsByte().Value;

                Console.WriteLine("Section, Y={0}", sy);
            }

            // var num = 8;
            // var section = sections.Value[num].AsCompound();
            // Console.WriteLine("Section[{0}]:", num);
            // section.Dump();
#endif
#endif
        }


        public int X
        {
            get
            {
                return root["xPos"].AsInt().Value;
            }
        }


        public int Y
        {
            get
            {
                return root["yPos"].AsInt().Value;
            }
        }


        public int Z
        {
            get
            {
                return root["zPos"].AsInt().Value;
            }
        }


        public long InhabitedTime
        {
            get
            {
                if (root.ContainsKey("InhabitedTime"))
                {
                    return root["InhabitedTime"].AsLong().Value;
                }
                else
                {
                    return 0;
                }
            }
        }


        public static Chunk LoadFrom(Stream stream)
        {
            using (var wrapper = new StreamWrapper(stream))
            using (var reader = new NbtReader(wrapper))
            {
                var rootTag = reader.ReadTag().AsCompound();

                return new Chunk(rootTag);
            }
        }
    }
}
