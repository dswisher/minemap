// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using MineMap.Cli.Helpers;
using MineMap.Cli.Options;
using MineMap.Lib.Files;
using MineMap.Lib.Graphics;
using MineMap.Lib.Util;

using SixLabors.ImageSharp;
using SixLabors.ImageSharp.PixelFormats;

namespace MineMap.Cli.Commands
{
    public class TimeMapCommand
    {
        public void Run(TimeMapOptions options)
        {
            // Set up the world, which will be used to find the files in the world.
            var world = options.GetWorld();

            // To properly scale the histogram, all chunks must be scanned. A 2D sparse matrix is used
            // to store the inhabited times for each chunk. While that is being populated, the histogram
            // is also created.
            var matrix = new Sparse2DMatrix<int, int, long>();
            var colorMap = new ColorMap();

            foreach (var regionPath in world.ListRegionPaths())
            {
                var region = new Region(regionPath);

                Console.WriteLine("Region {0}, X={1}, Z={2}:", regionPath, region.X, region.Z);

                var found = 0;
                for (var x = 0; x < 32; x++)
                {
                    for (var z = 0; z < 32; z++)
                    {
                        var pt = new ChunkPoint((region.X * 32) + x, (region.Z * 32) + z);

                        if (region.HasChunk(pt))
                        {
                            found += 1;

                            var chunk = region.GetChunk(pt);
                            var time = chunk.InhabitedTime;

                            matrix[pt.X, pt.Z] = time;

                            colorMap.AddSample(time);
                        }
                    }
                }

                Console.WriteLine("   {0} chunks", found);
            }

            // Set up the transformation from chunk coords to pixel coords
            var dx = matrix.MaxX - matrix.MinX;
            var dy = matrix.MaxY - matrix.MinY;

            var transform = new Transform2D();

            transform.InputRange(matrix.MinX, matrix.MaxX, matrix.MinY, matrix.MaxY);
            transform.OutputRange(0, dx - 1, 0, dy - 1);

            // Use the data to create the image
            using (var image = new Image<Rgba32>(dx * options.ChunkSize, dy * options.ChunkSize))
            {
                for (var x = matrix.MinX; x < matrix.MaxX; x++)
                {
                    for (var y = matrix.MinY; y < matrix.MaxY; y++)
                    {
                        var val = matrix[x, y];

                        var black = new Rgba32(0, 0, 0, 255);
                        Rgba32 color;
                        if (val == 0)
                        {
                            color = black;
                        }
                        else
                        {
                            color = colorMap.GetColor(val);
                        }

                        var px = options.ChunkSize * (int)transform.TransformX(x);
                        var py = options.ChunkSize * (int)transform.TransformY(y);

                        for (var xx = 0; xx < options.ChunkSize; xx++)
                        {
                            for (var yy = 0; yy < options.ChunkSize; yy++)
                            {
                                image[px + xx, py + yy] = color;
                            }
                        }
                    }
                }

                image.SaveAsPng(options.OutputPath);
            }

            Console.WriteLine("Wrote image to {0}.", options.OutputPath);
        }
    }
}
