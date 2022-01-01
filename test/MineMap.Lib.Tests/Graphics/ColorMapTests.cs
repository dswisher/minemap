// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using FluentAssertions;
using MineMap.Lib.Graphics;
using Xunit;

namespace MineMap.Lib.Tests.Graphics
{
    public class ColorMapTests
    {
        private readonly ColorMap map = new ColorMap();

        [Theory]
        [InlineData(0, 128)]
        public void HandlesSimpleCase(int val, byte red)
        {
            // Arrange
            map.AddSample(0);
            map.AddSample(128);

            // Act
            var color = map.GetColor(val);

            // Assert
            color.R.Should().Be(red);
            color.G.Should().Be(0);
            color.B.Should().Be(0);
        }
    }
}
