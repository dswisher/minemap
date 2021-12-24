// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using MineMap.Lib.Nbt;

namespace MineMap.Lib.Chunks
{
    public class Chunk
    {
        private Chunk(NbtCompound root)
        {
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
        }


        public int X { get; private set; }
        public int Y { get; private set; }
        public int Z { get; private set; }


        public static Chunk LoadFrom(NbtCompound root)
        {
            return new Chunk(root);
        }
    }
}
