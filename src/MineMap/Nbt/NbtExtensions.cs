
using System;
using System.IO;
using System.Linq;

namespace MineMap.Nbt
{
    public static class NbtExtensions
    {
        public static void Dump(this NbtCompound root, TextWriter writer = null)
        {
            if (writer == null)
            {
                writer = Console.Out;
            }

            writer.WriteLine("Root");

            foreach (var child in root.OrderBy(x => x.Key))
            {
                writer.WriteLine("   -> {0}", child.Key);
            }
        }
    }
}
