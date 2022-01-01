// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Graphics
{
    public class Transform2D
    {
        private readonly Transform1D xTransform = new Transform1D();
        private readonly Transform1D yTransform = new Transform1D();

        public void InputRange(int minX, int maxX, int minY, int maxY)
        {
            xTransform.InputRange(minX, maxX);
            yTransform.InputRange(minY, maxY);
        }


        public void OutputRange(int minX, int maxX, int minY, int maxY)
        {
            xTransform.OutputRange(minX, maxX);
            yTransform.OutputRange(minY, maxY);
        }


        public long TransformX(long x)
        {
            return xTransform.Transform(x);
        }


        public long TransformY(long y)
        {
            return yTransform.Transform(y);
        }
    }
}
