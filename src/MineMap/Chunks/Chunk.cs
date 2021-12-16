// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using MineMap.Nbt;

namespace MineMap.Chunks
{
    public class Chunk
    {
        private Chunk(NbtCompound root)
        {
            // TODO! HACK! Remove this debug code
            var sections = root["sections"].AsList();

            var num = 8;
            var section = sections.Value[num].AsCompound();
            Console.WriteLine("Section[{0}]:", num);
            section.Dump();


            X = root["xPos"].AsInt().Value;
            Y = root["yPos"].AsInt().Value;
            Z = root["zPos"].AsInt().Value;
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
