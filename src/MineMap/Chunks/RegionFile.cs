
using System;
using System.IO;

using ICSharpCode.SharpZipLib.Zip.Compression.Streams;
using MineMap.Helpers;

namespace MineMap.Chunks
{
    public class RegionFile : IDisposable
    {
        private const int SectorBytes = 4096;
        private const int IndexSize = SectorBytes / 4;

        private const int CompressionGzip = 1;
        private const int CompressionZlib = 2;

        private readonly int[] offsets = new int[IndexSize];

        private StreamWrapper wrapper;


        public static RegionFile Load(string regionDir, RegionPoint pos)
        {
            var path = Path.Combine(regionDir, $"r.{pos.X}.{pos.Z}.mca");

            return new RegionFile(path);
        }


        private RegionFile(string path)
        {
            var stream = new FileStream(path, FileMode.Open, FileAccess.Read);
            wrapper = new StreamWrapper(stream);

            for (var i = 0; i < IndexSize; i++)
            {
                offsets[i] = wrapper.ReadInt();
            }
        }


        public bool HasChunk(ChunkPoint pos)
        {
            var offset = GetOffset(pos);

            return offset != 0;
        }


        public Stream GetChunkStream(ChunkPoint pos)
        {
            var offset = GetOffset(pos);

            if (offset == 0)
            {
                // TODO - custom exception
                throw new Exception($"Chunk {pos} does not exist.");
            }

            var sectorNumber = offset >> 8;
            var chunkSectors = offset & 0xFF;

            wrapper.Seek(sectorNumber * SectorBytes, SeekOrigin.Begin);

            var chunkByteLen = wrapper.ReadInt();

            if (chunkByteLen > chunkSectors * SectorBytes)
            {
                // TODO - custom exception
                throw new Exception($"Chunk {pos} has invalid length, chunkSectors={chunkSectors}, chunkByteLen={chunkByteLen}.");
            }

            var compressionType = wrapper.ReadByte();

            if (compressionType == CompressionGzip)
            {
                // TODO - gzip
                throw new NotImplementedException($"Gzip compression is not yet handled, for chunk {pos}.");
            }
            else if (compressionType == CompressionZlib)
            {
                var chunkBytes = wrapper.ReadBytes(chunkByteLen - 1);

                return new InflaterInputStream(new MemoryStream(chunkBytes));
            }
            else
            {
                // TODO - custom exception
                throw new Exception($"Chunk {pos} has invalid compression type: {compressionType}.");
            }
        }


        public void Dispose()
        {
            if (wrapper != null)
            {
                wrapper.Dispose();
                wrapper = null;
            }
        }


        private int GetOffset (ChunkPoint pos)
        {
            var indexPos = (pos.X & 0x1f) + (pos.Z & 0x1f) * 32;

            return offsets[indexPos];
        }
    }
}
