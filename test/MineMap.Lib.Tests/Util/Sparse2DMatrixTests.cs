// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using FluentAssertions;
using MineMap.Lib.Util;
using Xunit;

namespace MineMap.Lib.Tests.Util
{
    public class Sparse2DMatrixTests
    {
        private readonly Sparse2DMatrix<int, int, int> matrix = new Sparse2DMatrix<int, int, int>();

        [Theory]
        [InlineData(0, 0, 0)]
        [InlineData(1, 1, 1)]
        public void CanGetAndSetSingleValue(int x, int y, int val)
        {
            // Arrange
            matrix[x, y] = val;

            // Act and assert
            matrix[x, y].Should().Be(val);
        }


        [Fact]
        public void CanGetAndSetMultipleValues()
        {
            // Arrange
            matrix[-1, -1] = -1;
            matrix[1, 1] = 1;
            matrix[2, 2] = 2;
            matrix[3, 3] = 3;

            // Act and assert
            matrix[-1, -1].Should().Be(-1);
            matrix[1, 1].Should().Be(1);
            matrix[2, 2].Should().Be(2);
            matrix[3, 3].Should().Be(3);

            matrix[4, 4].Should().Be(0);
        }
    }
}
