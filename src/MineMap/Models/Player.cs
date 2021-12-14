// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using MineMap.Nbt;

namespace MineMap.Models
{
    public class Player
    {
        public Player(NbtCompound root)
        {
            SpawnX = root["SpawnX"].AsInt().Value;
            SpawnY = root["SpawnY"].AsInt().Value;
            SpawnZ = root["SpawnZ"].AsInt().Value;
        }


        public int SpawnX { get; private set; }
        public int SpawnY { get; private set; }
        public int SpawnZ { get; private set; }
    }
}
