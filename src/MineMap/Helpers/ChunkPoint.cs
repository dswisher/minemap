
namespace MineMap.Helpers
{
    public class ChunkPoint
    {
        public ChunkPoint(int x, int y, int z)
        {
            X = x;
            Y = y;
            Z = z;
        }


        public int X { get; set; }
        public int Y { get; set; }
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
            return $"({X}, {Y}, {Z})";
        }
    }
}
