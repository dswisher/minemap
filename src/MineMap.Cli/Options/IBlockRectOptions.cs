// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Cli.Options
{
    public interface IBlockRectOptions
    {
        public int Width { get; }
        public int Height { get; }
        public int CenterX { get; }
        public int CenterZ { get; }
    }
}
