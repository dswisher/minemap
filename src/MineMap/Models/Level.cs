
using MineMap.Nbt;

namespace MineMap.Models
{
    public class Level
    {
        public static Level Load(NbtReader reader)
        {
            var tag = reader.ReadTag();

            return new Level(tag.AsCompound());
        }


        public Level(NbtCompound root)
        {
            var data = root["Data"].AsCompound();

            SpawnX = data["SpawnX"].AsInt().Value;
            SpawnY = data["SpawnY"].AsInt().Value;
            SpawnZ = data["SpawnZ"].AsInt().Value;

            // TODO - where is the seed? Doesn't seem to exist in my world?
            if (data.ContainsKey("RandomSeed"))
            {
                RandomSeed = data["RandomSeed"].AsLong().Value;
            }

            if (data.ContainsKey("WorldGenSettings"))
            {
                var settings = data["WorldGenSettings"].AsCompound();

                if (settings.ContainsKey("seed"))
                {
                    Seed = settings["seed"].AsLong().Value;
                }
            }

            Player = new Player(data["Player"].AsCompound());

            // TODO - pick out more props
            // TODO - use a "schema" or some such
        }


        public Player Player { get; private set; }

        public long RandomSeed { get; private set; }
        public long Seed { get; private set; }

        public int SpawnX { get; private set; }
        public int SpawnY { get; private set; }
        public int SpawnZ { get; private set; }
    }
}
