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
            // root.Dump();
            var sections = root["sections"].AsList();

            // var num = 20;
            // var section = sections.Value[num].AsCompound();
            // Console.WriteLine("Section[{0}]:", num);
            // section.Dump();

            for (var i = 0; i < sections.Value.Length; i++)
            {
                var section = sections.Value[i].AsCompound();
                var y = section["Y"].AsByte().Value;

                Console.WriteLine($"  section[{i}].Y = {y}");
            }



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
