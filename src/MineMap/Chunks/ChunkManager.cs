
using System.Collections.Generic;
using System.IO;
using System.Linq;

using MineMap.Helpers;
using MineMap.Nbt;

namespace MineMap.Chunks
{
    public class ChunkManager
    {
        // TODO - implement a real LRU cache
        private readonly Dictionary<RegionPoint, RegionFile> regionCache = new Dictionary<RegionPoint, RegionFile>();
        private readonly string regionDir;

        public ChunkManager(string worldDir)
        {
            regionDir = Path.Combine(worldDir, "region");
        }


        public Chunk GetChunk(ChunkPoint pos)
        {
            var regionPos = pos.ToRegion();

            RegionFile regionFile;
            if (regionCache.ContainsKey(regionPos))
            {
                regionFile = regionCache[regionPos];
            }
            else
            {
                regionFile = RegionFile.Load(regionDir, regionPos);

                if (regionCache.Count > 3)
                {
                    var doomedKey = regionCache.Keys.First();
                    regionCache.Remove(doomedKey);
                }

                regionCache.Add(regionPos, regionFile);
            }

            using (var chunkStream = regionFile.GetChunkStream(pos))
            using (var wrapper = new StreamWrapper(chunkStream))
            using (var reader = new NbtReader(wrapper))
            {
                var rootTag = reader.ReadTag().AsCompound();

                rootTag.Dump();

                // return Chunk.LoadFrom(rootTag);
                return new Chunk();
            }
        }
    }
}
