
using System;
using System.IO;
using System.IO.Compression;

using MineMap.Nbt;

namespace MineMap
{
    public class Program
    {
        public static void Main(string[] args)
        {
            try
            {
                var worldDir = "/Users/swisherd/Library/Application Support/minecraft/saves/New World";
                var levelFile = "level.dat";

                Console.WriteLine("Opening {0}...", levelFile);

                using (var stream = new FileStream(Path.Join(worldDir, levelFile), FileMode.Open, FileAccess.Read))
                using (var gzipStream = new GZipStream(stream, CompressionMode.Decompress))
                using (var reader = new NbtReader(gzipStream))
                {
                    var tag = reader.ReadTag();

                    // TODO - do something with the data
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine("Unhandled exception in main!");
                Console.WriteLine(ex);
            }
        }
    }
}
