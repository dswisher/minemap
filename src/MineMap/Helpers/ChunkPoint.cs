
namespace MineMap.Helpers
{
    public class ChunkPoint
    {
        public int X { get; set; }
        public int Z { get; set; }


        public RegionPoint ToRegion()
        {
            return new RegionPoint
            {
                X = X / 32,
                Z = Z / 32
            };
        }


        public override string ToString()
        {
            return $"({X}, {Z})";
        }
    }
}
