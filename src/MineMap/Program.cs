
using System;
using System.IO;
using System.IO.Compression;

using MineMap.Chunks;
using MineMap.Helpers;
using MineMap.Models;
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

                // Read the level info
                var levelFile = "level.dat";

                Level level;

                Console.WriteLine("Reading {0}...", levelFile);
                using (var stream = new FileStream(Path.Join(worldDir, levelFile), FileMode.Open, FileAccess.Read))
                using (var gzipStream = new GZipStream(stream, CompressionMode.Decompress))
                using (var wrapper = new StreamWrapper(gzipStream))
                using (var reader = new NbtReader(wrapper))
                {
                    level = Level.Load(reader);
                }

                Console.WriteLine("World Spawn: ({0}, {1}, {2})", level.SpawnX, level.SpawnY, level.SpawnZ);
                Console.WriteLine("Player Spawn: ({0}, {1}, {2})", level.Player.SpawnX, level.Player.SpawnY, level.Player.SpawnZ);
                Console.WriteLine("Seed: {0}", level.Seed);

                var spawnPos = new WorldPoint(level.Player.SpawnX, level.Player.SpawnY, level.Player.SpawnZ);
                // var spawnPos = new WorldPoint(174, 67, -40);
                var chunkPos = spawnPos.ToChunk();
                // var chunkPos = new ChunkPoint(10, 4, -3);
                // var chunkPos = new ChunkPoint(10, 4, -3);
                var regionPos = chunkPos.ToRegion();

                Console.WriteLine("Block {0} in chunk {1} in region {2}.", spawnPos, chunkPos, regionPos);

                var chunkManager = new ChunkManager(worldDir);
                var chunk = chunkManager.GetChunk(chunkPos);

                Console.WriteLine("  xPos={0}, yPos={1}, zPos={2}", chunk.X, chunk.Y, chunk.Z);
            }
            catch (Exception ex)
            {
                Console.WriteLine("Unhandled exception in main!");
                Console.WriteLine(ex);
            }
        }
    }
}
