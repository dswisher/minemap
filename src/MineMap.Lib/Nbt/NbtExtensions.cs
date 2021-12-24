// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.IO;
using System.Linq;

namespace MineMap.Lib.Nbt
{
    public static class NbtExtensions
    {
        public static void Dump(this NbtCompound root, TextWriter writer = null)
        {
            if (writer == null)
            {
                writer = Console.Out;
            }

            Dump(writer, root, "Root", 0);
        }


        private static void Dump(TextWriter writer, NbtCompound item, string name, int depth)
        {
            writer.WriteLine("{0} {1}: compound", Spaces(depth), name);

            foreach (var child in item.OrderBy(x => x.Key))
            {
                if (child.Value is NbtCompound comVal)
                {
                    Dump(writer, comVal, child.Key, depth + 1);
                }
                else if (child.Value is NbtByte byteVal)
                {
                    writer.WriteLine("{0} {1} (byte): {2}", Spaces(depth + 1), child.Key, byteVal.Value);
                }
                else if (child.Value is NbtInt intVal)
                {
                    writer.WriteLine("{0} {1} (int): {2}", Spaces(depth + 1), child.Key, intVal.Value);
                }
                else if (child.Value is NbtLong longVal)
                {
                    writer.WriteLine("{0} {1} (long): {2}", Spaces(depth + 1), child.Key, longVal.Value);
                }
                else if (child.Value is NbtString strVal)
                {
                    writer.WriteLine("{0} {1} (str): '{2}'", Spaces(depth + 1), child.Key, strVal.Value);
                }
                else if (child.Value is NbtList listVal)
                {
                    writer.WriteLine("{0} {1} (list): {2} items of type {3}", Spaces(depth + 1), child.Key, listVal.Value.Length, listVal.ChildType);
                }
                else
                {
                    writer.WriteLine("{0} {1} ({2}): TYPE TBD", Spaces(depth + 1), child.Key, child.Value.TagType.ToString().ToLower());
                }
            }
        }


        private static string Spaces(int depth)
        {
            return new string(' ', depth * 3);
        }
    }
}
