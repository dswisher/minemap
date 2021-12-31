// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.IO;
using System.Text.RegularExpressions;

using ICSharpCode.SharpZipLib.Zip.Compression.Streams;
using MineMap.Lib.Chunks;
using MineMap.Lib.Util;

namespace MineMap.Lib.Files
{
    public class Region : IDisposable
    {
        private const int SectorBytes = 4096;
        private const int IndexSize = SectorBytes / 4;

        private const int CompressionGzip = 1;
        private const int CompressionZlib = 2;

        private static readonly Regex NameRegex = new Regex(@"r\.([^\.]*)\.([^\.]*)\.mca");

        private readonly int[] offsets = new int[IndexSize];

        private StreamWrapper wrapper;

        public Region(string path)
        {
            // Pick apart the file name to extract the X and Z values
            var name = Path.GetFileName(path);
            var match = NameRegex.Match(name);

            if (!match.Success)
            {
                // TODO - throw custom exception
                throw new Exception($"File name {name} is not a valid region file name.");
            }

            X = int.Parse(match.Groups[1].Value);
            Z = int.Parse(match.Groups[2].Value);

            // Populate the offset list by parsing the file header. Note that the stream and wrapper
            // are cleaned up when this class is disposed.
            var stream = new FileStream(path, FileMode.Open, FileAccess.Read);
            wrapper = new StreamWrapper(stream);

            for (var i = 0; i < IndexSize; i++)
            {
                offsets[i] = wrapper.ReadInt();
            }
        }


        public int X { get; private set; }
        public int Z { get; private set; }


        public bool HasChunk(ChunkPoint pos)
        {
            var offset = GetOffset(pos);

            return offset != 0;
        }


        public Chunk GetChunk(ChunkPoint pos)
        {
            return Chunk.LoadFrom(GetChunkStream(pos));
        }


        public void Dispose()
        {
            if (wrapper != null)
            {
                wrapper.Dispose();
                wrapper = null;
            }
        }


        private Stream GetChunkStream(ChunkPoint pos)
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


        private int GetOffset(ChunkPoint pos)
        {
            // TODO - convert coordinate to chunk coords

            var indexPos = (pos.X & 0x1f) + ((pos.Z & 0x1f) * 32);

            return offsets[indexPos];
        }
    }
}
